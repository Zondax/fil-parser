package market_test

import (
	_ "embed"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/actors/market"
	"github.com/zondax/fil-parser/tools"

	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
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
	os.Exit(m.Run())
}

type testFn func(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error)

func TestPublishStorageDeals(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "PublishStorageDealsParams", expected)
	require.NoError(t, err)

	runTest(t, market.PublishStorageDealsParams, tests)
}

func TestVerifyDealsForActivation(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "VerifyDealsForActivationParams", expected)
	require.NoError(t, err)

	runTest(t, market.VerifyDealsForActivationParams, tests)
}

func TestActivateDeals(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "ActivateDealsParams", expected)
	require.NoError(t, err)

	runTest(t, market.ActivateDealsParams, tests)
}

func TestComputeDataCommitment(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "ComputeDataCommitmentParams", expected)
	require.NoError(t, err)

	runTest(t, market.ComputeDataCommitmentParams, tests)
}

func TestGetBalance(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "GetBalanceParams", expected)
	require.NoError(t, err)

	runTest(t, market.GetBalanceParams, tests)
}

func TestGetDealDataCommitment(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "GetDealDataCommitmentParams", expected)
	require.NoError(t, err)

	runTest(t, market.GetDealDataCommitmentParams, tests)
}

func TestGetDealClient(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "GetDealClientParams", expected)
	require.NoError(t, err)

	runTest(t, market.GetDealClientParams, tests)
}

func TestGetDealProvider(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "GetDealProviderParams", expected)
	require.NoError(t, err)

	runTest(t, market.GetDealProviderParams, tests)
}

func TestGetDealLabel(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "GetDealLabelParams", expected)
	require.NoError(t, err)

	runTest(t, market.GetDealLabelParams, tests)
}

func TestGetDealTerm(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "GetDealTermParams", expected)
	require.NoError(t, err)

	runTest(t, market.GetDealTermParams, tests)
}

func TestGetDealTotalPrice(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "GetDealTotalPriceParams", expected)
	require.NoError(t, err)

	runTest(t, market.GetDealTotalPriceParams, tests)
}

func TestGetDealClientCollateral(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "GetDealClientCollateralParams", expected)
	require.NoError(t, err)

	runTest(t, market.GetDealClientCollateralParams, tests)
}

func TestGetDealProviderCollateral(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "GetDealProviderCollateralParams", expected)
	require.NoError(t, err)

	runTest(t, market.GetDealProviderCollateralParams, tests)
}

func TestGetDealVerified(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "GetDealVerifiedParams", expected)
	require.NoError(t, err)

	runTest(t, market.GetDealVerifiedParams, tests)
}

func TestGetDealActivation(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "GetDealActivationParams", expected)
	require.NoError(t, err)

	runTest(t, market.GetDealActivationParams, tests)
}

func TestDealProviderCollateral(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "GetDealProviderCollateralParams", expected)
	require.NoError(t, err)

	runTest(t, market.GetDealProviderCollateralParams, tests)
}

func TestGetDealVerifiedParams(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "GetDealVerifiedParams", expected)
	require.NoError(t, err)

	runTest(t, market.GetDealVerifiedParams, tests)
}

func TestGetDealActivationParams(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "GetDealActivationParams", expected)
	require.NoError(t, err)

	runTest(t, market.GetDealActivationParams, tests)
}

func TestOnMinerSectorsTerminate(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "OnMinerSectorsTerminateParams", expected)
	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.Height, tt.Version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}

				result, err := market.OnMinerSectorsTerminateParams(tt.Network, tt.Height, trace.Msg.Params)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.Expected))
			}
		})
	}
}

func TestParseAddBalance(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "ParseAddBalance", expected)
	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.Height, tt.Version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}

				result, err := market.ParseAddBalance(tt.Network, tt.Height, trace.Msg.Params)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.Expected))
			}
		})
	}
}

func TestParseWithdrawBalance(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "ParseWithdrawBalance", expected)
	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.Height, tt.Version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}

				result, err := market.ParseWithdrawBalance(tt.Network, tt.Height, trace.Msg.Params)
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

				result, err := fn(tt.Network, tt.Height, trace.Msg.Params, trace.MsgRct.Return)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.Expected))
			}
		})
	}
}
