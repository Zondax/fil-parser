package miner_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/actors/miner"
	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
	"github.com/zondax/fil-parser/tools"
)

func TestDeclareFaults(t *testing.T) {
	tests := []test{}
	runTest(t, miner.DeclareFaults, tests)
}

func TestDeclareFaultsRecovered(t *testing.T) {
	tests := []test{}
	runTest(t, miner.DeclareFaultsRecovered, tests)
}

func TestProveReplicaUpdates(t *testing.T) {
	tests := []test{}
	runTest(t, miner.ProveReplicaUpdates, tests)
}

func TestPreCommitSectorBatch2(t *testing.T) {
	tests := []test{}
	runTest(t, miner.PreCommitSectorBatch2, tests)
}

func TestProveCommitAggregate(t *testing.T) {
	tests := []test{}
	runTest(t, miner.ProveCommitAggregate, tests)
}

func TestDisputeWindowedPoSt(t *testing.T) {
	tests := []test{}
	runTest(t, miner.DisputeWindowedPoSt, tests)
}

func TestReportConsensusFault(t *testing.T) {
	tests := []test{}
	runTest(t, miner.ReportConsensusFault, tests)
}

func TestChangeBeneficiary(t *testing.T) {
	tests := []test{}
	runTest(t, miner.ChangeBeneficiary, tests)
}

func TestMinerConstructor(t *testing.T) {
	tests := []test{}
	runTest(t, miner.MinerConstructor, tests)
}

func TestApplyRewards(t *testing.T) {
	tests := []test{}
	runTest(t, miner.ApplyRewards, tests)
}

func TestOnDeferredCronEvent(t *testing.T) {
	tests := []test{}
	runTest(t, miner.OnDeferredCronEvent, tests)
}

func TestTerminateSectors(t *testing.T) {
	tests := []test{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.height, tt.version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}

				result, err := miner.TerminateSectors(tt.height, trace.Msg.Params, trace.MsgRct.Return)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.expected))

			}
		})
	}
}

func TestProveReplicaUpdates2(t *testing.T) {
	tests := []test{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.height, tt.version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}

				result, err := miner.ProveReplicaUpdates2(tt.height, trace.Msg.Params, trace.MsgRct.Return)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.expected))

			}
		})
	}
}
