package miner_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/actors/miner"
	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
	"github.com/zondax/fil-parser/tools"
)

type testFn func(network string, height int64, rawReturn []byte) (map[string]interface{}, error)

func TestChangeMultiaddrs(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("ChangeMultiaddrs", expectedData)
	require.NoError(t, err)
	runTest(t, miner.ChangeMultiaddrs, tests)
}

func TestChangePeerID(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("ChangePeerID", expectedData)
	require.NoError(t, err)
	runTest(t, miner.ChangePeerID, tests)
}

func TestChangeWorkerAddress(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("ChangeWorkerAddress", expectedData)
	require.NoError(t, err)
	runTest(t, miner.ChangeWorkerAddress, tests)
}

func TestChangeOwnerAddress(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("ChangeOwnerAddress", expectedData)
	require.NoError(t, err)
	runTest(t, miner.ChangeOwnerAddress, tests)
}

func TestGetOwner(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("GetOwner", expectedData)
	require.NoError(t, err)
	runTest(t, miner.GetOwner, tests)
}

func TestGetPeerID(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("GetPeerID", expectedData)
	require.NoError(t, err)
	runTest(t, miner.GetPeerID, tests)
}

func TestGetMultiaddrs(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("GetMultiaddrs", expectedData)
	require.NoError(t, err)
	runTest(t, miner.GetMultiaddrs, tests)
}

func TestIsControllingAddressExported(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("IsControllingAddressExported", expectedData)
	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.Height, tt.Version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}

				result, err := miner.IsControllingAddressExported(tt.Network, tt.Height, trace.Msg.Params, trace.MsgRct.Return)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.Expected))

			}
		})
	}
}

func runTest(t *testing.T, fn testFn, tests []tools.TestCase[map[string]any]) {
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.Height, tt.Version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}

				result, err := fn(tt.Network, tt.Height, trace.MsgRct.Return)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.Expected))

			}
		})
	}
}
