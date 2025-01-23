package power_test

import (
	_ "embed"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/actors/power"
	"github.com/zondax/fil-parser/parser"
	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
	"github.com/zondax/fil-parser/tools"
)

//go:embed expected.json
var expectedData []byte
var expected map[string]any

var network string

func TestMain(m *testing.M) {
	network = "mainnet"
	if err := json.Unmarshal(expectedData, &expected); err != nil {
		panic(err)
	}
	m.Run()
}

type testFn func(network string, msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error)

func TestPower(t *testing.T) {
	power := &power.Power{}
	testFns := map[string]testFn{
		"CurrentTotalPower":        power.CurrentTotalPower,
		"SubmitPoRepForBulkVerify": power.SubmitPoRepForBulkVerify,
		"CreateMiner":              power.CreateMiner,
		"EnrollCronEvent":          power.EnrollCronEvent,
		"UpdateClaimedPower":       power.UpdateClaimedPower,
		"UpdatePledgeTotal":        power.UpdatePledgeTotal,
		"NetworkRawPower":          power.NetworkRawPower,
		"MinerRawPower":            power.MinerRawPower,
		"MinerCount":               power.MinerCount,
		"MinerConsensusCount":      power.MinerConsensusCount,
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
				result, err := power.PowerConstructor(tt.Network, tt.Height, lotusMsg, trace.Msg.Params)
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
