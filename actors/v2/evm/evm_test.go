package evm_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/actors/v2/evm"
	"github.com/zondax/fil-parser/tools"

	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
)

var expectedData []byte
var expected map[string]any
var network string

func TestMain(m *testing.M) {
	network = "mainnet"
	// if err := json.Unmarshal(expectedData, &expected); err != nil {
	// 	panic(err)
	// }
	os.Exit(m.Run())
}

type testFn func(network string, height int64, raw []byte) (map[string]interface{}, error)

func TestEVM(t *testing.T) {
	evm := &evm.Evm{}
	testFns := map[string]testFn{
		"Resurrect":       evm.Resurrect,
		"GetByteCode":     evm.GetBytecode,
		"GetBytecodeHash": evm.GetBytecodeHash,
		"Constructor":     evm.Constructor,
	}
	for name, fn := range testFns {
		t.Run(name, func(t *testing.T) {
			tests, err := tools.LoadTestData[map[string]any](network, name, expected)
			require.NoError(t, err)
			runTest(t, fn, tests)
		})
	}
}

func TestGetStorageAt(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "GetStorageAt", expected)
	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.Height, tt.Version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}
				evm := &evm.Evm{}
				result, err := evm.GetStorageAt(tt.Network, tt.Height, trace.Msg.Params, trace.MsgRct.Return)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.Expected))
			}
		})
	}
}

func TestInvokeContract(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "InvokeContract", expected)
	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.Height, tt.Version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}
				evm := &evm.Evm{}
				result, err := evm.InvokeContract(tt.Network, tt.Height, trace.Msg.Params, trace.MsgRct.Return)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.Expected))
			}
		})
	}
}

func TestInvokeContractDelegate(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "InvokeContractDelegate", expected)
	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.Height, tt.Version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}
				evm := &evm.Evm{}
				result, err := evm.InvokeContractDelegate(tt.Network, tt.Height, trace.Msg.Params, trace.MsgRct.Return)
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
				result, err := fn(tt.Network, tt.Height, trace.Msg.Params)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.Expected))
			}
		})
	}
}
