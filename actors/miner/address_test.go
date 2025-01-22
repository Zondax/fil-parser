package miner_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/actors/miner"
	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
	"github.com/zondax/fil-parser/tools"
)

type testFn func(height int64, rawReturn []byte) (map[string]interface{}, error)

type test struct {
	name     string
	version  string
	url      string
	height   int64
	expected map[string]any
}

func TestChangeMultiaddrs(t *testing.T) {
	tests := []test{}
	runTest(t, miner.ChangeMultiaddrs, tests)
}

func TestChangePeerID(t *testing.T) {
	tests := []test{}
	runTest(t, miner.ChangePeerID, tests)
}

func TestChangeWorkerAddress(t *testing.T) {
	tests := []test{}
	runTest(t, miner.ChangeWorkerAddress, tests)
}

func TestChangeOwnerAddress(t *testing.T) {
	tests := []test{}
	runTest(t, miner.ChangeOwnerAddress, tests)
}

func TestGetOwner(t *testing.T) {
	tests := []test{}
	runTest(t, miner.GetOwner, tests)
}

func TestGetPeerID(t *testing.T) {
	tests := []test{}
	runTest(t, miner.GetPeerID, tests)
}

func TestGetMultiaddrs(t *testing.T) {
	tests := []test{}
	runTest(t, miner.GetMultiaddrs, tests)
}

func TestIsControllingAddressExported(t *testing.T) {
	tests := []test{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.height, tt.version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}

				result, err := miner.IsControllingAddressExported(tt.height, trace.Msg.Params, trace.MsgRct.Return)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.expected))

			}
		})
	}
}

func runTest(t *testing.T, fn testFn, tests []test) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.height, tt.version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}

				result, err := fn(tt.height, trace.MsgRct.Return)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.expected))

			}
		})
	}
}
