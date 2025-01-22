package reward_test

import (
	_ "embed"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/actors/reward"
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

type testFn func(network string, height int64, raw []byte) (map[string]interface{}, error)

func TestRewardConstructor(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "RewardConstructor", expected)
	require.NoError(t, err)
	runTest(t, reward.RewardConstructor, tests)
}

func TestAwardBlockReward(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "AwardBlockReward", expected)
	require.NoError(t, err)
	runTest(t, reward.AwardBlockReward, tests)
}

func TestUpdateNetworkKPI(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "UpdateNetworkKPI", expected)
	require.NoError(t, err)
	runTest(t, reward.UpdateNetworkKPI, tests)
}

func TestThisEpochReward(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "ThisEpochReward", expected)
	require.NoError(t, err)
	runTest(t, reward.ThisEpochReward, tests)
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
