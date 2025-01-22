package reward_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/actors/reward"
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

func TestRewardConstructor(t *testing.T) {
	tests := []test{}
	runTest(t, reward.RewardConstructor, tests)
}

func TestAwardBlockReward(t *testing.T) {
	tests := []test{}
	runTest(t, reward.AwardBlockReward, tests)
}

func TestUpdateNetworkKPI(t *testing.T) {
	tests := []test{}
	runTest(t, reward.UpdateNetworkKPI, tests)
}

func TestThisEpochReward(t *testing.T) {
	tests := []test{}
	runTest(t, reward.ThisEpochReward, tests)
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
