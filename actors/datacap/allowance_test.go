package datacap_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/actors/datacap"
	"github.com/zondax/fil-parser/tools"

	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
)

type testFn func(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error)

func TestIncreaseAllowance(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "IncreaseAllowance", expected)
	require.NoError(t, err)

	runTest(t, datacap.IncreaseAllowance, tests)
}

func TestDecreaseAllowance(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "DecreaseAllowance", expected)
	require.NoError(t, err)

	runTest(t, datacap.DecreaseAllowance, tests)
}

func TestRevokeAllowance(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "RevokeAllowance", expected)
	require.NoError(t, err)

	runTest(t, datacap.RevokeAllowance, tests)
}

func TestGetAllowance(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "GetAllowance", expected)
	require.NoError(t, err)

	runTest(t, datacap.GetAllowance, tests)
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
				result, err := fn(tt.Network, tt.Height, trace.Msg.Params, trace.MsgRct.Return)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.Expected))
			}
		})
	}
}
