package fil_parser

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/filecoin-project/lotus/api"
	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/parser"
	v2 "github.com/zondax/fil-parser/parser/v2"
	"github.com/zondax/fil-parser/types"
)

// mainnetStableNodeUrl is an archival mainnet node, for heights whose actors the light node no
// longer has.
const mainnetStableNodeUrl = "https://node-fil-mainnet-stable.zondax.ch/rpc/v1"

func parseFixtureDeals(t *testing.T, p *FilecoinParser, height string, traces []byte) ([]*types.Transaction, *types.DealsEvents) {
	node := mainnetCacheDataSource.Node
	tipset, err := readTipset(height)
	require.NoError(t, err)
	ethlogs, err := readEthLogs(height)
	require.NoError(t, err)
	parsed, err := p.ParseTransactions(t.Context(), []api.FullNode{node}, types.TxsData{
		EthLogs:   ethlogs,
		Tipset:    tipset,
		Traces:    traces,
		Metadata:  types.BlockMetadata{NodeInfo: types.NodeInfo{NodeMajorMinorVersion: v2.NodeVersionsSupported[len(v2.NodeVersionsSupported)-1]}},
		Canonical: true,
	})
	require.NoError(t, err)
	events, err := p.ParseDealsEvents(t.Context(), []api.FullNode{node}, parsed.Txs, tipset.GetCidString(), tipset.Key())
	require.NoError(t, err)
	return parsed.Txs, events
}

// A successful OnMinerSectorsTerminate sub-call under a message that failed emits nothing. The real
// trace of mainnet 6,539,228 (one TerminateSectors message, one market call) is parsed as is, then
// with the message's exit code set to an error and the sub-call's left at 0, so the parser's own
// status propagation (main status / sub-call status) decides.
func TestParser_TerminationUnderRevertedParent(t *testing.T) {
	const height = "6539228"
	p, err := NewFilecoinParserWithActorV2(getLib(nodeUrl), mainnetCacheDataSource, gLogger)
	require.NoError(t, err)
	traces, err := readGzFile(tracesFilename(height))
	require.NoError(t, err)

	_, events := parseFixtureDeals(t, p, height, traces)
	require.Len(t, events.DealsTerminations, 1)

	var compute map[string]interface{}
	require.NoError(t, json.Unmarshal(traces, &compute))
	reverted := 0
	for _, raw := range compute["Trace"].([]interface{}) {
		trace := raw.(map[string]interface{})
		execution := trace["ExecutionTrace"].(map[string]interface{})
		calls, _ := execution["Subcalls"].([]interface{})
		terminates := false
		for _, rawCall := range calls {
			msg := rawCall.(map[string]interface{})["Msg"].(map[string]interface{})
			if msg["To"] == "f05" && msg["Method"] == float64(7) {
				terminates = true
			}
		}
		if !terminates {
			continue
		}
		// USR_ILLEGAL_STATE on the message, in both receipts the parser reads
		trace["MsgRct"].(map[string]interface{})["ExitCode"] = 16
		execution["MsgRct"].(map[string]interface{})["ExitCode"] = 16
		reverted++
	}
	require.Equal(t, 1, reverted)
	mutated, err := json.Marshal(compute)
	require.NoError(t, err)

	txs, events := parseFixtureDeals(t, p, height, mutated)
	found := false
	for _, tx := range txs {
		if tx.TxType == parser.MethodOnMinerSectorsTerminate {
			found = true
			require.NotEqual(t, "Ok", tx.Status)
		}
	}
	require.True(t, found)
	require.Empty(t, events.DealsTerminations)
}

// With ConsolidateRobustAddress on, as in production, the miner gets the same address in the rows
// that use the miner as the caller of the market (terminations, activations: tx.TxFrom of the
// sub-call) and in the rows that use it as the receiver of a message (pre-commit and SnapDeals
// mappings: tx.TxTo), so the join on actor_address matches.
func TestParser_DealsAddressesConsolidated(t *testing.T) {
	p, err := NewFilecoinParserWithActorV2(getLib(nodeUrl), mainnetCacheDataSource, gLogger,
		WithConfig(parser.Config{ConsolidateRobustAddress: true}))
	require.NoError(t, err)

	t.Run("SnapDeals mapping and activation, 3573062", func(t *testing.T) {
		const height = "3573062"
		traces, err := readGzFile(tracesFilename(height))
		require.NoError(t, err)
		_, events := parseFixtureDeals(t, p, height, traces)

		activations := map[string]map[uint64]string{}
		for _, a := range events.DealsActivations {
			if activations[a.TxCid] == nil {
				activations[a.TxCid] = map[uint64]string{}
			}
			activations[a.TxCid][a.DealID] = a.ActorAddress
		}
		snap := 0
		for _, m := range events.DealsSectorMappings {
			require.Truef(t, strings.HasPrefix(m.ActorAddress, "f2"), "mapping address %q is not robust", m.ActorAddress)
			if m.ActionType != parser.MethodProveReplicaUpdates {
				continue
			}
			snap++
			activated, ok := activations[m.TxCid][m.DealID]
			require.True(t, ok)
			require.Equal(t, activated, m.ActorAddress)
		}
		require.Equal(t, 2, snap)
	})

	t.Run("termination and the miner's own message, 6541321", func(t *testing.T) {
		const height = "6541321"
		traces, err := readGzFile(tracesFilename(height))
		require.NoError(t, err)
		txs, events := parseFixtureDeals(t, p, height, traces)

		byID := map[string]*types.Transaction{}
		for _, tx := range txs {
			byID[tx.Id] = tx
		}
		// the miner as the receiver of TerminateSectors (tx.TxTo) and as the caller of the market
		// (tx.TxFrom of the OnMinerSectorsTerminate sub-call)
		minerOfCall := map[string]string{}
		for _, tx := range txs {
			if tx.TxType != parser.MethodOnMinerSectorsTerminate {
				continue
			}
			parent := byID[tx.ParentId]
			require.NotNil(t, parent)
			require.Equal(t, parent.TxTo, tx.TxFrom)
			minerOfCall[tx.TxFrom] = parent.TxTo
		}
		require.Len(t, minerOfCall, 3)
		require.Len(t, events.DealsTerminations, 156)
		for _, r := range events.DealsTerminations {
			require.Truef(t, strings.HasPrefix(r.ActorAddress, "f2"), "termination address %q is not robust", r.ActorAddress)
			require.Equal(t, minerOfCall[r.ActorAddress], r.ActorAddress)
		}
	})
}

// The tipset AT the NV22 upgrade height (mainnet 3,855,360) still executes with actors v12: Lotus
// migrates the state while computing the next tipset (ApplyBlocks runs HandleStateForks for i in
// [parentEpoch, epoch)), and StateNetworkVersion is 21 there (tools.ExecutedVersionFromHeight is V21;
// tools.VersionFromHeight labels the height V22).
// The raw trace proves it: its BatchActivateDeals calls (cron ConfirmSectorProofsValid) carry the
// 3-field v12 SectorDeals {sector_type, sector_expiry, deal_ids}, without a sector number.
// The deals rows must not record a sector number for them (the decoded metadata shows
// "SectorNumber":0), must not copy them as mappings, and must map the tipset's pre-commits.
// Only the traces and the tipset are in the fixture: the node had no event index for the height.
func TestParser_DealsAtNV22UpgradeTipset(t *testing.T) {
	const height = "3855360"
	traces, err := readGzFile(tracesFilename(height))
	require.NoError(t, err)

	// independent check of the raw params: 3-field SectorDeals
	var compute struct {
		Trace []struct {
			ExecutionTrace json.RawMessage
		}
	}
	require.NoError(t, json.Unmarshal(traces, &compute))
	type call struct {
		Msg struct {
			To     string
			Method uint64
			Params []byte
		}
		Subcalls []json.RawMessage
	}
	batchCalls := 0
	var walk func(json.RawMessage)
	walk = func(node json.RawMessage) {
		var c call
		require.NoError(t, json.Unmarshal(node, &c))
		if c.Msg.To == "f05" && c.Msg.Method == 6 {
			batchCalls++
			// [ [ SectorDeals... ], compute_cid ]; SectorDeals header 0x83 = array of 3
			require.Equal(t, byte(0x82), c.Msg.Params[0])
			require.Equal(t, byte(0x83), c.Msg.Params[2], "v12 SectorDeals has 3 fields")
		}
		for _, s := range c.Subcalls {
			walk(s)
		}
	}
	for _, tr := range compute.Trace {
		walk(tr.ExecutionTrace)
	}
	require.Equal(t, 7, batchCalls)

	// The tipset has messages to miners that no longer exist; the light node only has the latest
	// state and cannot resolve them, so this height is parsed against the archival stable node.
	stable := getCacheDataSource("mainnet", mainnetStableNodeUrl)
	p, err := NewFilecoinParserWithActorV2(getLib(mainnetStableNodeUrl), stable, gLogger)
	require.NoError(t, err)
	node := stable.Node
	tipset, err := readTipset(height)
	require.NoError(t, err)
	parsed, err := p.ParseTransactions(t.Context(), []api.FullNode{node}, types.TxsData{
		Tipset:    tipset,
		Traces:    traces,
		Metadata:  types.BlockMetadata{NodeInfo: types.NodeInfo{NodeMajorMinorVersion: v2.NodeVersionsSupported[len(v2.NodeVersionsSupported)-1]}},
		Canonical: true,
	})
	require.NoError(t, err)
	events, err := p.ParseDealsEvents(t.Context(), []api.FullNode{node}, parsed.Txs, tipset.GetCidString(), tipset.Key())
	require.NoError(t, err)

	require.Len(t, events.DealsActivations, 13)
	for _, a := range events.DealsActivations {
		require.Equal(t, parser.MethodBatchActivateDeals, a.ActionType)
		require.Nilf(t, a.SectorNumber, "deal %d activated by actors v12 has a sector number", a.DealID)
	}
	preCommits := 0
	for _, m := range events.DealsSectorMappings {
		require.Equal(t, parser.MethodPreCommitSectorBatch2, m.ActionType)
		preCommits++
	}
	require.Positive(t, preCommits)
}
