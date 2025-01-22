package paymentchannel_test

import (
	_ "embed"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	paymentchannel "github.com/zondax/fil-parser/actors/paymentChannel"
	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
	"github.com/zondax/fil-parser/tools"
)

//go:embed expected.json
var expectedData []byte
var expected map[string]any

func TestMain(m *testing.M) {
	if err := json.Unmarshal(expectedData, &expected); err != nil {
		panic(err)
	}
	m.Run()
}

type testFn func(height int64, raw []byte) (map[string]interface{}, error)
type test struct {
	name     string
	height   int64
	version  string
	expected map[string]interface{}
}

func TestPaymentChannelConstructor(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("PaymentChannelConstructor", expected)
	require.NoError(t, err)
	runTest(t, paymentchannel.PaymentChannelConstructor, tests)
}

func TestUpdateChannelState(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("UpdateChannelState", expected)
	require.NoError(t, err)
	runTest(t, paymentchannel.UpdateChannelState, tests)
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
