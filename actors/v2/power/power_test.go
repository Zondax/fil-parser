package power_test

import (
	"testing"

	powerv10 "github.com/filecoin-project/go-state-types/builtin/v10/power"
	powerv11 "github.com/filecoin-project/go-state-types/builtin/v11/power"
	powerv12 "github.com/filecoin-project/go-state-types/builtin/v12/power"
	powerv13 "github.com/filecoin-project/go-state-types/builtin/v13/power"
	powerv14 "github.com/filecoin-project/go-state-types/builtin/v14/power"
	powerv15 "github.com/filecoin-project/go-state-types/builtin/v15/power"
	powerv8 "github.com/filecoin-project/go-state-types/builtin/v8/power"
	powerv9 "github.com/filecoin-project/go-state-types/builtin/v9/power"
	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/power"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/power"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/power"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/power"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/power"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/power"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	v2 "github.com/zondax/fil-parser/actors/v2"
	"github.com/zondax/fil-parser/actors/v2/power"
	"github.com/zondax/fil-parser/parser"
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

type testFn func(network string, msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error)

func TestPower(t *testing.T) {
	power := &power.Power{}
	testFns := map[string]testFn{
		"CurrentTotalPower":        power.CurrentTotalPower,
		"SubmitPoRepForBulkVerify": power.SubmitPoRepForBulkVerify,
		// "CreateMiner":              power.CreateMinerExported,
		"EnrollCronEvent":     power.EnrollCronEvent,
		"UpdateClaimedPower":  power.UpdateClaimedPower,
		"UpdatePledgeTotal":   power.UpdatePledgeTotal,
		"NetworkRawPower":     power.NetworkRawPowerExported,
		"MinerRawPower":       power.MinerRawPowerExported,
		"MinerCount":          power.MinerCountExported,
		"MinerConsensusCount": power.MinerConsensusCountExported,
	}
	for name, fn := range testFns {
		t.Run(name, func(t *testing.T) {
			tests, err := tools.LoadTestData[map[string]any](network, name, expected)
			require.NoError(t, err)
			runTest(t, fn, tests)
		})
	}
}

func TestPowerConstructor(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "PowerConstructor", expected)
	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.Height, tt.Version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}
				lotusMsg := &parser.LotusMessage{
					To:     trace.Msg.To,
					From:   trace.Msg.From,
					Method: trace.Msg.Method,
				}
				power := &power.Power{}
				result, err := power.Constructor(tt.Network, tt.Height, lotusMsg, trace.Msg.Params)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.Expected))
			}
		})
	}
}

func TestMethodCoverage(t *testing.T) {
	power := &power.Power{}

	actorVersions := []any{
		legacyv2.Actor{},
		legacyv3.Actor{},
		legacyv4.Actor{},
		legacyv5.Actor{},
		legacyv6.Actor{},
		legacyv7.Actor{},
		powerv8.Methods,
		powerv9.Methods,
		powerv10.Methods,
		powerv11.Methods,
		powerv12.Methods,
		powerv13.Methods,
		powerv14.Methods,
		powerv15.Methods,
	}

	missingMethods := v2.MissingMethods(power, actorVersions)
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
				lotusMsg := &parser.LotusMessage{
					To:     trace.Msg.To,
					From:   trace.Msg.From,
					Method: trace.Msg.Method,
				}
				result, err := fn(tt.Network, lotusMsg, tt.Height, trace.Msg.Params, trace.MsgRct.Return)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.Expected))
			}
		})
	}
}
