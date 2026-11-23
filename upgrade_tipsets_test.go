package fil_parser

import (
	"encoding/json"
	"testing"

	"github.com/filecoin-project/lotus/api"
	"github.com/stretchr/testify/require"

	"github.com/zondax/fil-parser/parser"
	v2 "github.com/zondax/fil-parser/parser/v2"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
)

// The tipset AT a network-upgrade height is executed by the old actors (the state migration runs
// before the next tipset's messages; StateNetworkVersion is the old version there). These tests
// parse real upgrade tipsets and check the parser decodes them with the old actors' types and
// method names (tools.ExecutedVersionFromHeight).
//
// Old heights reference actors the light node no longer has, so they run against the archival node.
func parseUpgradeTipset(t *testing.T, height string) []*types.Transaction {
	t.Helper()
	ds := getCacheDataSource("mainnet", mainnetStableNodeUrl)
	p, err := NewFilecoinParserWithActorV2(getLib(mainnetStableNodeUrl), ds, gLogger)
	require.NoError(t, err)
	tipset, err := readTipset(height)
	require.NoError(t, err)
	traces, err := readGzFile(tracesFilename(height))
	require.NoError(t, err)
	parsed, err := p.ParseTransactions(t.Context(), []api.FullNode{ds.Node}, types.TxsData{
		Tipset:    tipset,
		Traces:    traces,
		Metadata:  types.BlockMetadata{NodeInfo: types.NodeInfo{NodeMajorMinorVersion: v2.NodeVersionsSupported[len(v2.NodeVersionsSupported)-1]}},
		Canonical: true,
	})
	require.NoError(t, err)
	return parsed.Txs
}

// Mainnet 3,469,380 is the NV21 (Watermelon) upgrade height: its tipset ran actors v11, whose
// market method 6 is ActivateDeals (actors v12 renamed it BatchActivateDeals).
func TestParser_UpgradeTipsetNV21UsesV11Methods(t *testing.T) {
	require.Equal(t, tools.V20.String(), tools.ExecutedVersionFromHeight(tools.MainnetNetwork, 3469380).String())
	txs := parseUpgradeTipset(t, "3469380")
	activate := 0
	for _, tx := range txs {
		require.NotEqual(t, parser.MethodBatchActivateDeals, tx.TxType, "tx %s named with the v12 method map", tx.TxCid)
		if tx.TxType == parser.MethodActivateDeals {
			activate++
		}
	}
	require.Positive(t, activate)
}

// Mainnet 3,855,360 is the NV22 (Dragon) upgrade height: its tipset ran actors v12, whose
// VerifyDealsForActivation SectorDeals has no sector number. Decoding it with the v13 type showed a
// spurious "SectorNumber":0.
func TestParser_UpgradeTipsetNV22UsesV12Types(t *testing.T) {
	txs := parseUpgradeTipset(t, "3855360")
	verify := 0
	for _, tx := range txs {
		if tx.TxType != parser.MethodVerifyDealsForActivation {
			continue
		}
		verify++
		var md struct {
			Params struct {
				Sectors []map[string]json.RawMessage
			}
		}
		require.NoError(t, json.Unmarshal([]byte(tx.TxMetadata), &md))
		require.NotEmpty(t, md.Params.Sectors)
		for _, s := range md.Params.Sectors {
			require.NotContains(t, s, "SectorNumber", "tx %s decoded with the v13 SectorDeals", tx.TxCid)
		}
	}
	require.Positive(t, verify)
}
