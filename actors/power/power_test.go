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

func TestMain(m *testing.M) {
	if err := json.Unmarshal(expectedData, &expected); err != nil {
		panic(err)
	}
	m.Run()
}

type testFn func(msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error)

func TestCurrentTotalPower(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("CurrentTotalPower", expected)
	require.NoError(t, err)
	runTest(t, power.CurrentTotalPower, tests)
}

func TestSubmitPoRepForBulkVerify(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("SubmitPoRepForBulkVerify", expected)
	require.NoError(t, err)
	runTest(t, power.SubmitPoRepForBulkVerify, tests)
}

func TestCreateMiner(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("CreateMiner", expected)
	require.NoError(t, err)
	runTest(t, power.CreateMiner, tests)
}

func TestEnrollCronEvent(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("EnrollCronEvent", expected)
	require.NoError(t, err)
	runTest(t, power.EnrollCronEvent, tests)
}

func TestUpdateClaimedPower(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("UpdateClaimedPower", expected)
	require.NoError(t, err)
	runTest(t, power.UpdateClaimedPower, tests)
}

func TestUpdatePledgeTotal(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("UpdatePledgeTotal", expected)
	require.NoError(t, err)
	runTest(t, power.UpdatePledgeTotal, tests)
}

func TestNetworkRawPower(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("NetworkRawPower", expected)
	require.NoError(t, err)
	runTest(t, power.NetworkRawPower, tests)
}

func TestMinerRawPower(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("MinerRawPower", expected)
	require.NoError(t, err)
	runTest(t, power.MinerRawPower, tests)
}

func TestMinerCount(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("MinerCount", expected)
	require.NoError(t, err)
	runTest(t, power.MinerCount, tests)
}

func TestMinerConsensusCount(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("MinerConsensusCount", expected)
	require.NoError(t, err)
	runTest(t, power.MinerConsensusCount, tests)
}

func TestPowerConstructor(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("PowerConstructor", expected)
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
				result, err := power.PowerConstructor(tt.Height, lotusMsg, trace.Msg.Params)
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
				result, err := fn(lotusMsg, tt.Height, trace.Msg.Params, trace.MsgRct.Return)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.Expected))
			}
		})
	}
}
