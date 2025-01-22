package datacap_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/actors/datacap"
	"github.com/zondax/fil-parser/tools"

	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
)

type testFn func(height int64, raw, rawReturn []byte) (map[string]interface{}, error)

type test struct {
	name     string
	version  string
	url      string
	height   int64
	expected map[string]any
}

func TestIncreaseAllowance(t *testing.T) {
	tests := []test{}

	runTest(t, datacap.IncreaseAllowance, tests)
}

func TestDecreaseAllowance(t *testing.T) {
	tests := []test{}

	runTest(t, datacap.DecreaseAllowance, tests)
}

func TestRevokeAllowance(t *testing.T) {
	tests := []test{}

	runTest(t, datacap.RevokeAllowance, tests)
}

func TestGetAllowance(t *testing.T) {
	tests := []test{}

	runTest(t, datacap.GetAllowance, tests)
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
				result, err := fn(tt.height, trace.Msg.Params, trace.MsgRct.Return)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.expected))
			}
		})
	}
}
