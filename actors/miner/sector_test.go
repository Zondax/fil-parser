package miner_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/actors/miner"
	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
	"github.com/zondax/fil-parser/tools"
)

func TestExtendSectorExpiration2(t *testing.T) {
	tests := []test{}
	runTest(t, miner.ExtendSectorExpiration2, tests)
}

func TestPreCommitSector(t *testing.T) {
	tests := []test{}
	runTest(t, miner.PreCommitSector, tests)
}

func TestProveCommitSector(t *testing.T) {
	tests := []test{}
	runTest(t, miner.ProveCommitSector, tests)
}

func TestSubmitWindowedPoSt(t *testing.T) {
	tests := []test{}
	runTest(t, miner.SubmitWindowedPoSt, tests)
}

func TestConfirmSectorProofsValid(t *testing.T) {
	tests := []test{}
	runTest(t, miner.ConfirmSectorProofsValid, tests)
}

func TestCheckSectorProven(t *testing.T) {
	tests := []test{}
	runTest(t, miner.CheckSectorProven, tests)
}

func TestExtendSectorExpiration(t *testing.T) {
	tests := []test{}
	runTest(t, miner.ExtendSectorExpiration, tests)
}

func TestCompactSectorNumbers(t *testing.T) {
	tests := []test{}
	runTest(t, miner.CompactSectorNumbers, tests)
}

func TestCompactPartitions(t *testing.T) {
	tests := []test{}
	runTest(t, miner.CompactPartitions, tests)
}

func TestPreCommitSectorBatch(t *testing.T) {
	tests := []test{}
	runTest(t, miner.PreCommitSectorBatch, tests)
}

func TestGetSectorSize(t *testing.T) {
	tests := []test{}
	runTest(t, miner.GetSectorSize, tests)
}

func TestProveCommitSectors3(t *testing.T) {
	tests := []test{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.height, tt.version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}

				result, err := miner.ProveCommitSectors3(tt.height, trace.Msg.Params, trace.MsgRct.Return)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.expected))

			}
		})
	}
}
