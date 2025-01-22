package miner_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/actors/miner"
	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
	"github.com/zondax/fil-parser/tools"
)

func TestExtendSectorExpiration2(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("ExtendSectorExpiration2", expectedData)
	require.NoError(t, err)
	runTest(t, miner.ExtendSectorExpiration2, tests)
}

func TestPreCommitSector(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("PreCommitSector", expectedData)
	require.NoError(t, err)
	runTest(t, miner.PreCommitSector, tests)
}

func TestProveCommitSector(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("ProveCommitSector", expectedData)
	require.NoError(t, err)
	runTest(t, miner.ProveCommitSector, tests)
}

func TestSubmitWindowedPoSt(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("SubmitWindowedPoSt", expectedData)
	require.NoError(t, err)
	runTest(t, miner.SubmitWindowedPoSt, tests)
}

func TestConfirmSectorProofsValid(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("ConfirmSectorProofsValid", expectedData)
	require.NoError(t, err)
	runTest(t, miner.ConfirmSectorProofsValid, tests)
}

func TestCheckSectorProven(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("CheckSectorProven", expectedData)
	require.NoError(t, err)
	runTest(t, miner.CheckSectorProven, tests)
}

func TestExtendSectorExpiration(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("ExtendSectorExpiration", expectedData)
	require.NoError(t, err)
	runTest(t, miner.ExtendSectorExpiration, tests)
}

func TestCompactSectorNumbers(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("CompactSectorNumbers", expectedData)
	require.NoError(t, err)
	runTest(t, miner.CompactSectorNumbers, tests)
}

func TestCompactPartitions(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("CompactPartitions", expectedData)
	require.NoError(t, err)
	runTest(t, miner.CompactPartitions, tests)
}

func TestPreCommitSectorBatch(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("PreCommitSectorBatch", expectedData)
	require.NoError(t, err)
	runTest(t, miner.PreCommitSectorBatch, tests)
}

func TestGetSectorSize(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("GetSectorSize", expectedData)
	require.NoError(t, err)
	runTest(t, miner.GetSectorSize, tests)
}

func TestProveCommitSectors3(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("ProveCommitSectors3", expectedData)
	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.Height, tt.Version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}

				result, err := miner.ProveCommitSectors3(tt.Height, trace.Msg.Params, trace.MsgRct.Return)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.Expected))

			}
		})
	}
}
