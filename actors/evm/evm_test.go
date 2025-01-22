package evm_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/actors/evm"
	"github.com/zondax/fil-parser/tools"

	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
)

type testFn func(height int64, raw []byte) (map[string]interface{}, error)

type test struct {
	name     string
	version  string
	url      string
	height   int64
	expected map[string]any
}

func TestResurrect(t *testing.T) {
	tests := []test{}

	runTest(t, evm.Resurrect, tests)
}

func TestGetByteCode(t *testing.T) {
	tests := []test{}

	runTest(t, evm.GetByteCode, tests)
}

func TestGetByteCodeHash(t *testing.T) {
	tests := []test{}

	runTest(t, evm.GetByteCodeHash, tests)
}

func TestEVMConstructor(t *testing.T) {
	tests := []test{}

	runTest(t, evm.EVMConstructor, tests)
}

func TestGetStorageAt(t *testing.T) {
	tests := []test{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.height, tt.version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}
				result, err := evm.GetStorageAt(tt.height, trace.Msg.Params, trace.MsgRct.Return)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.expected))
			}
		})
	}
}

func TestInvokeContract(t *testing.T) {
	tests := []test{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.height, tt.version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}
				result, err := evm.InvokeContract(trace.Msg.Params, trace.MsgRct.Return)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.expected))
			}
		})
	}
}

func TestInvokeContractDelegate(t *testing.T) {
	tests := []test{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.height, tt.version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}
				result, err := evm.InvokeContractDelegate(tt.height, trace.Msg.Params, trace.MsgRct.Return)
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
				result, err := fn(tt.height, trace.Msg.Params)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.expected))
			}
		})
	}
}
