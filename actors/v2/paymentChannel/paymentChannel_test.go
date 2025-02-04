package paymentchannel_test

import (
	"testing"

	paychv10 "github.com/filecoin-project/go-state-types/builtin/v10/paych"
	paychv11 "github.com/filecoin-project/go-state-types/builtin/v11/paych"
	paychv12 "github.com/filecoin-project/go-state-types/builtin/v12/paych"
	paychv13 "github.com/filecoin-project/go-state-types/builtin/v13/paych"
	paychv14 "github.com/filecoin-project/go-state-types/builtin/v14/paych"
	paychv15 "github.com/filecoin-project/go-state-types/builtin/v15/paych"
	paychv8 "github.com/filecoin-project/go-state-types/builtin/v8/paych"
	paychv9 "github.com/filecoin-project/go-state-types/builtin/v9/paych"
	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/paych"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/paych"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/paych"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/paych"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/paych"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/paych"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	v2 "github.com/zondax/fil-parser/actors/v2"
	paymentchannel "github.com/zondax/fil-parser/actors/v2/paymentChannel"
	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
	"github.com/zondax/fil-parser/tools"
)

var expectedData []byte
var expected map[string]any

var network string

func TestMain(m *testing.M) {
	network = "mainnet"
	// if err := json.Unmarshal(expectedData, &expected); err != nil {
	// 	panic(err)
	// }
	// var err error
	// expectedData, err = tools.ReadActorSnapshot()
	// if err != nil {
	// 	panic(err)
	// }
	m.Run()
}

type testFn func(network string, height int64, raw []byte) (map[string]interface{}, error)
type test struct {
	name     string
	height   int64
	version  string
	expected map[string]interface{}
}

func TestPaymentChannel(t *testing.T) {
	paymentchannel := &paymentchannel.PaymentChannel{}
	testFns := map[string]testFn{
		"PaymentChannelConstructor": paymentchannel.Constructor,
		"UpdateChannelState":        paymentchannel.UpdateChannelState,
	}
	for name, fn := range testFns {
		t.Run(name, func(t *testing.T) {
			tests, err := tools.LoadTestData[map[string]any](network, name, expected)
			require.NoError(t, err)
			runTest(t, fn, tests)
		})
	}
}

func TestMethodCoverage(t *testing.T) {
	pc := &paymentchannel.PaymentChannel{}

	actorVersions := []any{
		legacyv2.Actor{},
		legacyv3.Actor{},
		legacyv4.Actor{},
		legacyv5.Actor{},
		legacyv6.Actor{},
		legacyv7.Actor{},
		paychv8.Methods,
		paychv9.Methods,
		paychv10.Methods,
		paychv11.Methods,
		paychv12.Methods,
		paychv13.Methods,
		paychv14.Methods,
		paychv15.Methods,
	}

	missingMethods := v2.MissingMethods(pc, actorVersions)
	assert.Empty(t, missingMethods, "missing methods: %v", missingMethods)
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
