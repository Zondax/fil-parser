package paymentchannel_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	paymentchannel "github.com/zondax/fil-parser/actors/paymentChannel"
	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
	"github.com/zondax/fil-parser/tools"
)

type testFn func(height int64, raw []byte) (map[string]interface{}, error)
type test struct {
	name     string
	height   int64
	version  string
	expected map[string]interface{}
}

func TestPaymentChannelConstructor(t *testing.T) {
	tests := []test{}
	runTest(t, paymentchannel.PaymentChannelConstructor, tests)
}

func TestUpdateChannelState(t *testing.T) {
	tests := []test{}
	runTest(t, paymentchannel.UpdateChannelState, tests)
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
