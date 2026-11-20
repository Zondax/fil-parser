package fil_parser

import (
	"bytes"
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

type minerSector struct {
	miner  string
	sector uint64
}

// terminationEvents returns the market's `deal-terminated` events (deal id -> provider) and the
// miners' `sector-terminated` events of a tipset.
func terminationEvents(t *testing.T, events []*filTypes.ActorEvent) (map[uint64]string, map[minerSector]bool) {
	market, err := address.NewIDAddress(5)
	require.NoError(t, err)

	dealsTerminated := map[uint64]string{}
	sectorsTerminated := map[minerSector]bool{}
	for _, event := range events {
		if event.Reverted {
			continue
		}
		var eventType string
		values := map[string]uint64{}
		for _, entry := range event.Entries {
			if entry.Key == "$type" {
				eventType, err = cbg.ReadString(cbg.NewCborReader(bytes.NewReader(entry.Value)))
				require.NoError(t, err)
				continue
			}
			maj, value, err := cbg.NewCborReader(bytes.NewReader(entry.Value)).ReadHeader()
			if err == nil && maj == cbg.MajUnsignedInt {
				values[entry.Key] = value
			}
		}
		switch eventType {
		case "deal-terminated":
			require.Equal(t, market, event.Emitter)
			provider, err := address.NewIDAddress(values["provider"])
			require.NoError(t, err)
			dealsTerminated[values["id"]] = provider.String()
		case "sector-terminated":
			sectorsTerminated[minerSector{miner: event.Emitter.String(), sector: values["sector"]}] = true
		}
	}
	return dealsTerminated, sectorsTerminated
}

// TestParser_DealTerminations parses whole mainnet tipsets with TerminateSectors messages (actors
// v16+, so OnMinerSectorsTerminate carries a sector bitfield) and checks the termination rows:
//   - every row is a terminated sector of its miner (`sector-terminated` event);
//   - every deal the market terminated (`deal-terminated` event) is in a terminated sector of its
//     provider. The deal -> sector mapping was read from the node (StateMarketStorageDeal at the
//     tipset, before its messages run) and is hard-coded below; deal_id is NULL on the rows, and
//     downstream joins them to deals through deals_sector_mappings.
//
// Not every row has a `deal-terminated` event: the market skips deals with end_epoch <= epoch, and
// the miner sends every terminated sector that holds data. At 6,541,321 only one of the three
// miners' sectors still had live deals.
func TestParser_DealTerminations(t *testing.T) {
	tests := []struct {
		height string
		// number of OnMinerSectorsTerminate calls and of rows (one per sector)
		calls, rows int
		// deal id -> provider and sector, from the node
		deals map[uint64]minerSector
	}{
		{
			// one TerminateSectors message, one sector, one live deal
			height: "6539228",
			calls:  1,
			rows:   1,
			deals: map[uint64]minerSector{
				129272826: {"f03644142", 46914},
			},
		},
		{
			// nine TerminateSectors messages from three miners; only f03645883's sectors had live deals
			height: "6541321",
			calls:  9,
			rows:   156,
			deals: map[uint64]minerSector{
				128427037: {"f03645883", 11382},
				128546624: {"f03645883", 12913},
				128640505: {"f03645883", 16755},
				129028318: {"f03645883", 35344},
				129465563: {"f03645883", 53294},
				129497568: {"f03645883", 56607},
				129603389: {"f03645883", 62239},
				129644668: {"f03645883", 65496},
				129741703: {"f03645883", 68901},
				130161405: {"f03645883", 84251},
				130247761: {"f03645883", 87059},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.height, func(t *testing.T) {
			p, err := NewFilecoinParserWithActorV2(getLib(nodeUrl), mainnetCacheDataSource, gLogger)
			require.NoError(t, err)
			node := mainnetCacheDataSource.Node

			tipset, err := readTipset(tt.height)
			require.NoError(t, err)
			ethlogs, err := readEthLogs(tt.height)
			require.NoError(t, err)
			traces, err := readGzFile(tracesFilename(tt.height))
			require.NoError(t, err)
			nativeLogs, err := readNativeLogs(tt.height)
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

			calls := 0
			for _, message := range events.DealsMessages {
				if message.ActionType == parser.MethodOnMinerSectorsTerminate {
					calls++
				}
			}
			require.Equal(t, tt.calls, calls)
			require.Len(t, events.DealsTerminations, tt.rows)

			dealsTerminated, sectorsTerminated := terminationEvents(t, nativeLogs)

			rows := map[minerSector]bool{}
			ids := map[string]bool{}
			for _, row := range events.DealsTerminations {
				require.Nil(t, row.DealID)
				require.NotNil(t, row.SectorNumber)
				// from actors v13 the epoch is the current epoch
				require.Equal(t, int64(tipset.Height()), row.TerminationEpoch)
				require.Equal(t, parser.MethodOnMinerSectorsTerminate, row.ActionType)
				require.Falsef(t, ids[row.ID], "duplicated id %s", row.ID)
				ids[row.ID] = true

				key := minerSector{miner: row.ActorAddress, sector: *row.SectorNumber}
				require.Falsef(t, rows[key], "sector %v twice", key)
				rows[key] = true
				require.Truef(t, sectorsTerminated[key], "sector %v has no sector-terminated event", key)
			}

			require.Len(t, dealsTerminated, len(tt.deals))
			for dealID, provider := range dealsTerminated {
				sector, ok := tt.deals[dealID]
				require.Truef(t, ok, "deal %d", dealID)
				require.Equal(t, sector.miner, provider)
				require.Truef(t, rows[sector], "deal %d: sector %v not terminated", dealID, sector)
			}
		})
	}
}
