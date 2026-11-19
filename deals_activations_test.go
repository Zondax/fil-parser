package fil_parser

import (
	"bytes"
	"sort"
	"testing"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/stretchr/testify/require"
	cbg "github.com/whyrusleeping/cbor-gen"
	"github.com/zondax/fil-parser/parser"
	v2 "github.com/zondax/fil-parser/parser/v2"
	"github.com/zondax/fil-parser/types"
)

// dealActivatedEventIDs returns the deal ids of the market actor's `deal-activated` events.
// The market emits exactly one per deal it activates, whatever the activation path
// (ActivateDeals/BatchActivateDeals or SectorContentChanged), so it is the oracle for the
// deal activations the parser must produce for a tipset.
func dealActivatedEventIDs(t *testing.T, events []*filTypes.ActorEvent) []uint64 {
	market, err := address.NewIDAddress(5)
	require.NoError(t, err)

	ids := []uint64{}
	for _, event := range events {
		if event.Reverted || event.Emitter.Payload() == nil || !bytes.Equal(event.Emitter.Payload(), market.Payload()) || event.Emitter.Protocol() != address.ID {
			continue
		}
		var eventType string
		var dealID *uint64
		for _, entry := range event.Entries {
			switch entry.Key {
			case "$type":
				eventType, err = cbg.ReadString(cbg.NewCborReader(bytes.NewReader(entry.Value)))
				require.NoError(t, err)
			case "id":
				maj, value, err := cbg.NewCborReader(bytes.NewReader(entry.Value)).ReadHeader()
				require.NoError(t, err)
				require.Equal(t, byte(cbg.MajUnsignedInt), maj)
				dealID = &value
			}
		}
		if eventType == "deal-activated" {
			require.NotNil(t, dealID)
			ids = append(ids, *dealID)
		}
	}
	return ids
}

// TestParser_DealActivationsMatchDealActivatedEvents parses whole tipsets and checks that the set of
// deal activations equals the set of `deal-activated` events the node reports for the tipset, and
// that no deal is emitted twice.
//
// Lotus indexes only events from explicit messages, not from implicit ones such as cron. Activations
// under cron (e.g. ConfirmSectorProofsValid -> BatchActivateDeals in actors v13) have no indexed
// event, so the fixtures are tipsets whose activations all come from explicit ProveCommitSectors3 /
// ProveReplicaUpdates3 messages.
func TestParser_DealActivationsMatchDealActivatedEvents(t *testing.T) {
	tests := []struct {
		name   string
		url    string
		height string
		// number of deal-activated events in the tipset
		activations int
		// number of activations coming from SectorContentChanged
		sectorContentChanged int
	}{
		{
			// one ProveReplicaUpdates3 (SnapDeals) message activating one deal
			name:                 "calibration ProveReplicaUpdates3",
			url:                  calibNextNodeUrl,
			height:               "4100979",
			activations:          1,
			sectorContentChanged: 1,
		},
		{
			// one ProveCommitSectors3 message activating 35 deals and one ProveReplicaUpdates3 activating 1
			name:                 "mainnet ProveCommitSectors3 and ProveReplicaUpdates3",
			url:                  nodeUrl,
			height:               "6559222",
			activations:          36,
			sectorContentChanged: 36,
		},
		{
			// one ProveCommitSectors3 message notifying the market of 32 sectors at once
			name:                 "mainnet ProveCommitSectors3 multi sector",
			url:                  nodeUrl,
			height:               "6559711",
			activations:          136,
			sectorContentChanged: 136,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var p *FilecoinParser
			var node api.FullNode
			var err error
			if tt.url == nodeUrl {
				p, err = NewFilecoinParserWithActorV2(getLib(tt.url), mainnetCacheDataSource, gLogger)
				node = mainnetCacheDataSource.Node
			} else {
				p, err = NewFilecoinParserWithActorV2(getLib(tt.url), calibNextNodeCacheDataSource, gLogger)
				node = calibNextNodeCacheDataSource.Node
			}
			require.NoError(t, err)

			tipset, err := readTipset(tt.height)
			require.NoError(t, err)
			ethlogs, err := readEthLogs(tt.height)
			require.NoError(t, err)
			traces, err := readGzFile(tracesFilename(tt.height))
			require.NoError(t, err)
			nativeLogs, err := readNativeLogs(tt.height)
			require.NoError(t, err)

			txsData := types.TxsData{
				EthLogs:   ethlogs,
				Tipset:    tipset,
				Traces:    traces,
				Metadata:  types.BlockMetadata{NodeInfo: types.NodeInfo{NodeMajorMinorVersion: v2.NodeVersionsSupported[len(v2.NodeVersionsSupported)-1]}},
				Canonical: true,
			}

			parsedResult, err := p.ParseTransactions(t.Context(), []api.FullNode{node}, txsData)
			require.NoError(t, err)

			events, err := p.ParseDealsEvents(t.Context(), []api.FullNode{node}, parsedResult.Txs, tipset.GetCidString(), tipset.Key())
			require.NoError(t, err)

			want := dealActivatedEventIDs(t, nativeLogs)
			require.Len(t, want, tt.activations)

			got := make([]uint64, 0, len(events.DealsActivations))
			seen := map[uint64]bool{}
			sectorContentChanged := 0
			for _, activation := range events.DealsActivations {
				require.Falsef(t, seen[activation.DealID], "deal %d activated twice", activation.DealID)
				seen[activation.DealID] = true
				got = append(got, activation.DealID)
				if activation.ActionType == parser.MethodSectorContentChanged {
					sectorContentChanged++
					require.NotNilf(t, activation.SectorNumber, "deal %d has no sector number", activation.DealID)
				}
			}
			sort.Slice(want, func(i, j int) bool { return want[i] < want[j] })
			sort.Slice(got, func(i, j int) bool { return got[i] < got[j] })
			require.Equal(t, want, got)
			require.Equal(t, tt.sectorContentChanged, sectorContentChanged)

			ids := map[string]bool{}
			for _, activation := range events.DealsActivations {
				require.Falsef(t, ids[activation.ID], "duplicated activation id %s", activation.ID)
				ids[activation.ID] = true
			}

			spaceDeals := map[uint64]bool{}
			for _, space := range events.DealsSpaceInfo {
				if space.ActionType != parser.MethodSectorContentChanged {
					continue
				}
				require.Falsef(t, spaceDeals[space.DealID], "deal %d has two space rows", space.DealID)
				spaceDeals[space.DealID] = true
				require.Positive(t, space.NonVerifiedDealSpace.Sign())
			}
			require.Len(t, spaceDeals, tt.sectorContentChanged)
		})
	}
}
