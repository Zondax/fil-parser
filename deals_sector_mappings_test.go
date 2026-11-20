package fil_parser

import (
	"bytes"
	"encoding/json"
	"sort"
	"testing"

	"github.com/filecoin-project/go-bitfield"
	miner12 "github.com/filecoin-project/go-state-types/builtin/v12/miner"
	"github.com/filecoin-project/lotus/api"
	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/parser"
	v2 "github.com/zondax/fil-parser/parser/v2"
	"github.com/zondax/fil-parser/types"
)

type sectorDeal struct {
	miner        string
	sector, deal uint64
}

// replicaUpdateMappingsFromTraces decodes every successful ProveReplicaUpdates (miner method 27)
// call of a fixture straight from the raw trace, independently of the parser: params with the
// actors v12 types (the CBOR tuple is the same from actors v7 to v12) and the return as a BitField.
// The expected mappings are the deals of the first update of each sector in the return.
func replicaUpdateMappingsFromTraces(t *testing.T, height string) (calls int, mappings []sectorDeal) {
	raw, err := readGzFile(tracesFilename(height))
	require.NoError(t, err)
	var trace struct {
		Trace []struct {
			ExecutionTrace json.RawMessage
		}
	}
	require.NoError(t, json.Unmarshal(raw, &trace))

	type call struct {
		Msg struct {
			To     string
			Method uint64
			Params []byte
		}
		MsgRct struct {
			ExitCode int
			Return   []byte
		}
		Subcalls []json.RawMessage
	}
	var walk func(json.RawMessage)
	walk = func(node json.RawMessage) {
		var c call
		require.NoError(t, json.Unmarshal(node, &c))
		if c.Msg.Method == 27 && c.MsgRct.ExitCode == 0 && len(c.MsgRct.Return) > 0 {
			calls++
			var params miner12.ProveReplicaUpdatesParams
			require.NoError(t, params.UnmarshalCBOR(bytes.NewReader(c.Msg.Params)))
			var updated bitfield.BitField
			require.NoError(t, updated.UnmarshalCBOR(bytes.NewReader(c.MsgRct.Return)))
			seen := map[uint64]bool{}
			for _, update := range params.Updates {
				sector := uint64(update.SectorID)
				if seen[sector] {
					continue
				}
				seen[sector] = true
				ok, err := updated.IsSet(sector)
				require.NoError(t, err)
				if !ok {
					continue
				}
				for _, deal := range update.Deals {
					mappings = append(mappings, sectorDeal{c.Msg.To, sector, uint64(deal)})
				}
			}
		}
		for _, sub := range c.Subcalls {
			walk(sub)
		}
	}
	for _, tr := range trace.Trace {
		walk(tr.ExecutionTrace)
	}
	return calls, mappings
}

func sortSectorDeals(s []sectorDeal) {
	sort.Slice(s, func(i, j int) bool {
		if s[i].miner != s[j].miner {
			return s[i].miner < s[j].miner
		}
		if s[i].sector != s[j].sector {
			return s[i].sector < s[j].sector
		}
		return s[i].deal < s[j].deal
	})
}

// TestParser_SectorDealMappings parses whole mainnet tipsets with pre-NV22 SnapDeals
// (ProveReplicaUpdates) and pre-commits and checks the sector -> deal mappings:
//   - the ProveReplicaUpdates rows equal the mappings decoded independently from the raw trace
//     (params and the returned BitField of updated sectors);
//   - every mapped SnapDeals deal was activated by the same message (its ActivateDeals /
//     BatchActivateDeals sub-call) by the same miner address, so the mapping joins the activation
//     rows and the termination rows on the same provider form;
//   - pre-commit rows are emitted for the pre-commits with deals.
func TestParser_SectorDealMappings(t *testing.T) {
	tests := []struct {
		height string
		// actors version of the height, for the test name only
		actors string
		// ProveReplicaUpdates calls in the tipset
		replicaUpdates int
	}{
		{"2907520", "v10", 4},
		{"3573062", "v12", 2},
		{"3573064", "v12", 2},
		{"3573066", "v12", 2},
	}

	for _, tt := range tests {
		t.Run(tt.height+" actors "+tt.actors, func(t *testing.T) {
			p, err := NewFilecoinParserWithActorV2(getLib(nodeUrl), mainnetCacheDataSource, gLogger)
			require.NoError(t, err)
			node := mainnetCacheDataSource.Node

			tipset, err := readTipset(tt.height)
			require.NoError(t, err)
			ethlogs, err := readEthLogs(tt.height)
			require.NoError(t, err)
			traces, err := readGzFile(tracesFilename(tt.height))
			require.NoError(t, err)

			parsedResult, err := p.ParseTransactions(t.Context(), []api.FullNode{node}, types.TxsData{
				EthLogs:   ethlogs,
				Tipset:    tipset,
				Traces:    traces,
				Metadata:  types.BlockMetadata{NodeInfo: types.NodeInfo{NodeMajorMinorVersion: v2.NodeVersionsSupported[len(v2.NodeVersionsSupported)-1]}},
				Canonical: true,
			})
			require.NoError(t, err)
			events, err := p.ParseDealsEvents(t.Context(), []api.FullNode{node}, parsedResult.Txs, tipset.GetCidString(), tipset.Key())
			require.NoError(t, err)

			calls, want := replicaUpdateMappingsFromTraces(t, tt.height)
			require.Equal(t, tt.replicaUpdates, calls)
			require.NotEmpty(t, want)

			activated := map[string]map[uint64]string{} // tx cid -> deal -> miner
			for _, a := range events.DealsActivations {
				if activated[a.TxCid] == nil {
					activated[a.TxCid] = map[uint64]string{}
				}
				activated[a.TxCid][a.DealID] = a.ActorAddress
			}

			got := []sectorDeal{}
			preCommits := 0
			ids := map[string]bool{}
			for _, m := range events.DealsSectorMappings {
				require.Falsef(t, ids[m.ID], "duplicated id %s", m.ID)
				ids[m.ID] = true
				switch m.ActionType {
				case parser.MethodProveReplicaUpdates:
					got = append(got, sectorDeal{m.ActorAddress, m.SectorNumber, m.DealID})
					miner, ok := activated[m.TxCid][m.DealID]
					require.Truef(t, ok, "deal %d mapped by %s was not activated by it", m.DealID, m.TxCid)
					require.Equal(t, miner, m.ActorAddress)
				case parser.MethodPreCommitSector, parser.MethodPreCommitSectorBatch, parser.MethodPreCommitSectorBatch2:
					preCommits++
				default:
					// before NV22 activations carry no sector number, so nothing is copied from them
					t.Fatalf("unexpected mapping source %s", m.ActionType)
				}
			}
			sortSectorDeals(want)
			sortSectorDeals(got)
			require.Equal(t, want, got)
			require.Positive(t, preCommits)
		})
	}
}
