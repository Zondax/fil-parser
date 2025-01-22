package power_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/actors/power"
	"github.com/zondax/fil-parser/parser"
	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
	"github.com/zondax/fil-parser/tools"
)

type testFn func(msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error)

type test struct {
	name     string
	height   int64
	version  string
	expected map[string]interface{}
}

func TestCurrentTotalPower(t *testing.T) {
	tests := []test{}
	runTest(t, power.CurrentTotalPower, tests)
}

func TestSubmitPoRepForBulkVerify(t *testing.T) {
	tests := []test{}
	runTest(t, power.SubmitPoRepForBulkVerify, tests)
}

func TestCreateMiner(t *testing.T) {
	tests := []test{}
	runTest(t, power.CreateMiner, tests)
}

func TestEnrollCronEvent(t *testing.T) {
	tests := []test{}
	runTest(t, power.EnrollCronEvent, tests)
}

func TestUpdateClaimedPower(t *testing.T) {
	tests := []test{}
	runTest(t, power.UpdateClaimedPower, tests)
}

func TestUpdatePledgeTotal(t *testing.T) {
	tests := []test{}
	runTest(t, power.UpdatePledgeTotal, tests)
}

func TestNetworkRawPower(t *testing.T) {
	tests := []test{}
	runTest(t, power.NetworkRawPower, tests)
}

func TestMinerRawPower(t *testing.T) {
	tests := []test{}
	runTest(t, power.MinerRawPower, tests)
}

func TestMinerCount(t *testing.T) {
	tests := []test{}
	runTest(t, power.MinerCount, tests)
}

func TestMinerConsensusCount(t *testing.T) {
	tests := []test{}
	runTest(t, power.MinerConsensusCount, tests)
}

func TestPowerConstructor(t *testing.T) {
	tests := []test{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.height, tt.version)
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
				result, err := power.PowerConstructor(tt.height, lotusMsg, trace.Msg.Params)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.expected))
			}
		})
	}
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
				lotusMsg := &parser.LotusMessage{
					To:     trace.Msg.To,
					From:   trace.Msg.From,
					Method: trace.Msg.Method,
				}
				result, err := fn(lotusMsg, tt.height, trace.Msg.Params, trace.MsgRct.Return)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.expected))
			}
		})
	}
}
