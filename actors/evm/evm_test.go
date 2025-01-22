package evm_test

import (
	_ "embed"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/actors/evm"
	"github.com/zondax/fil-parser/tools"

	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
)

//go:embed expected.json
var expectedData []byte
var expected map[string]any

func TestMain(m *testing.M) {
	if err := json.Unmarshal(expectedData, &expected); err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

type testFn func(height int64, raw []byte) (map[string]interface{}, error)

func TestResurrect(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("Resurrect", expected)
	require.NoError(t, err)

	runTest(t, evm.Resurrect, tests)
}

func TestGetByteCode(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("GetByteCode", expected)
	require.NoError(t, err)

	runTest(t, evm.GetByteCode, tests)
}

func TestGetByteCodeHash(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("GetByteCodeHash", expected)
	require.NoError(t, err)

	runTest(t, evm.GetByteCodeHash, tests)
}

func TestEVMConstructor(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("EVMConstructor", expected)
	require.NoError(t, err)

	runTest(t, evm.EVMConstructor, tests)
}

func TestGetStorageAt(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("GetStorageAt", expected)
	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.Height, tt.Version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}
				result, err := evm.GetStorageAt(tt.Height, trace.Msg.Params, trace.MsgRct.Return)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.Expected))
			}
		})
	}
}

func TestInvokeContract(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("InvokeContract", expected)
	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.Height, tt.Version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}
				result, err := evm.InvokeContract(trace.Msg.Params, trace.MsgRct.Return)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.Expected))
			}
		})
	}
}

func TestInvokeContractDelegate(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("InvokeContractDelegate", expected)
	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.Height, tt.Version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}
				result, err := evm.InvokeContractDelegate(tt.Height, trace.Msg.Params, trace.MsgRct.Return)
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
				result, err := fn(tt.Height, trace.Msg.Params)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.Expected))
			}
		})
	}
}
