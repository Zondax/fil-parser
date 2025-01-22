package miner_test

import (
	_ "embed"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/actors/miner"
	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
	"github.com/zondax/fil-parser/tools"
)

//go:embed expected.json
var expected []byte
var expectedData map[string]any

func TestMain(m *testing.M) {
	if err := json.Unmarshal(expected, &expectedData); err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestDeclareFaults(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("DeclareFaults", expectedData)
	require.NoError(t, err)
	runTest(t, miner.DeclareFaults, tests)
}

func TestDeclareFaultsRecovered(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("DeclareFaultsRecovered", expectedData)
	require.NoError(t, err)
	runTest(t, miner.DeclareFaultsRecovered, tests)
}

func TestProveReplicaUpdates(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("ProveReplicaUpdates", expectedData)
	require.NoError(t, err)
	runTest(t, miner.ProveReplicaUpdates, tests)
}

func TestPreCommitSectorBatch2(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("PreCommitSectorBatch2", expectedData)
	require.NoError(t, err)
	runTest(t, miner.PreCommitSectorBatch2, tests)
}

func TestProveCommitAggregate(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("ProveCommitAggregate", expectedData)
	require.NoError(t, err)
	runTest(t, miner.ProveCommitAggregate, tests)
}

func TestDisputeWindowedPoSt(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("DisputeWindowedPoSt", expectedData)
	require.NoError(t, err)
	runTest(t, miner.DisputeWindowedPoSt, tests)
}

func TestReportConsensusFault(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("ReportConsensusFault", expectedData)
	require.NoError(t, err)
	runTest(t, miner.ReportConsensusFault, tests)
}

func TestChangeBeneficiary(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("ChangeBeneficiary", expectedData)
	require.NoError(t, err)
	runTest(t, miner.ChangeBeneficiary, tests)
}

func TestMinerConstructor(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("MinerConstructor", expectedData)
	require.NoError(t, err)
	runTest(t, miner.MinerConstructor, tests)
}

func TestApplyRewards(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("ApplyRewards", expectedData)
	require.NoError(t, err)
	runTest(t, miner.ApplyRewards, tests)
}

func TestOnDeferredCronEvent(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("OnDeferredCronEvent", expectedData)
	require.NoError(t, err)
	runTest(t, miner.OnDeferredCronEvent, tests)
}

func TestTerminateSectors(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("TerminateSectors", expectedData)
	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.Height, tt.Version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}

				result, err := miner.TerminateSectors(tt.Height, trace.Msg.Params, trace.MsgRct.Return)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.Expected))

			}
		})
	}
}

func TestProveReplicaUpdates2(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("ProveReplicaUpdates2", expectedData)
	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.Height, tt.Version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}

				result, err := miner.ProveReplicaUpdates2(tt.Height, trace.Msg.Params, trace.MsgRct.Return)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.Expected))

			}
		})
	}
}
