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

func TestMain(m *testing.M) {
	if err := json.Unmarshal(expectedData, &expected); err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

type testFn func(height int64, rawParams, rawReturn []byte) (map[string]interface{}, error)

func TestPublishStorageDeals(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("PublishStorageDealsParams", expected)
	require.NoError(t, err)

	runTest(t, market.PublishStorageDealsParams, tests)
}

func TestVerifyDealsForActivation(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("VerifyDealsForActivationParams", expected)
	require.NoError(t, err)

	runTest(t, market.VerifyDealsForActivationParams, tests)
}

func TestActivateDeals(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("ActivateDealsParams", expected)
	require.NoError(t, err)

	runTest(t, market.ActivateDealsParams, tests)
}

func TestComputeDataCommitment(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("ComputeDataCommitmentParams", expected)
	require.NoError(t, err)

	runTest(t, market.ComputeDataCommitmentParams, tests)
}

func TestGetBalance(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("GetBalanceParams", expected)
	require.NoError(t, err)

	runTest(t, market.GetBalanceParams, tests)
}

func TestGetDealDataCommitment(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("GetDealDataCommitmentParams", expected)
	require.NoError(t, err)

	runTest(t, market.GetDealDataCommitmentParams, tests)
}

func TestGetDealClient(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("GetDealClientParams", expected)
	require.NoError(t, err)

	runTest(t, market.GetDealClientParams, tests)
}

func TestGetDealProvider(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("GetDealProviderParams", expected)
	require.NoError(t, err)

	runTest(t, market.GetDealProviderParams, tests)
}

func TestGetDealLabel(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("GetDealLabelParams", expected)
	require.NoError(t, err)

	runTest(t, market.GetDealLabelParams, tests)
}

func TestGetDealTerm(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("GetDealTermParams", expected)
	require.NoError(t, err)

	runTest(t, market.GetDealTermParams, tests)
}

func TestGetDealTotalPrice(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("GetDealTotalPriceParams", expected)
	require.NoError(t, err)

	runTest(t, market.GetDealTotalPriceParams, tests)
}

func TestGetDealClientCollateral(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("GetDealClientCollateralParams", expected)
	require.NoError(t, err)

	runTest(t, market.GetDealClientCollateralParams, tests)
}

func TestGetDealProviderCollateral(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("GetDealProviderCollateralParams", expected)
	require.NoError(t, err)

	runTest(t, market.GetDealProviderCollateralParams, tests)
}

func TestGetDealVerified(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("GetDealVerifiedParams", expected)
	require.NoError(t, err)

	runTest(t, market.GetDealVerifiedParams, tests)
}

func TestGetDealActivation(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("GetDealActivationParams", expected)
	require.NoError(t, err)

	runTest(t, market.GetDealActivationParams, tests)
}

func TestDealProviderCollateral(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("GetDealProviderCollateralParams", expected)
	require.NoError(t, err)

	runTest(t, market.GetDealProviderCollateralParams, tests)
}

func TestGetDealVerifiedParams(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("GetDealVerifiedParams", expected)
	require.NoError(t, err)

	runTest(t, market.GetDealVerifiedParams, tests)
}

func TestGetDealActivationParams(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("GetDealActivationParams", expected)
	require.NoError(t, err)

	runTest(t, market.GetDealActivationParams, tests)
}

func TestOnMinerSectorsTerminate(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("OnMinerSectorsTerminateParams", expected)
	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.Height, tt.Version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}

				result, err := market.OnMinerSectorsTerminateParams(tt.Height, trace.Msg.Params)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.Expected))
			}
		})
	}
}

func TestParseAddBalance(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("ParseAddBalance", expected)
	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.Height, tt.Version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}

				result, err := market.ParseAddBalance(tt.Height, trace.Msg.Params)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.Expected))
			}
		})
	}
}

func TestParseWithdrawBalance(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("ParseWithdrawBalance", expected)
	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.Height, tt.Version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}

				result, err := market.ParseWithdrawBalance(tt.Height, trace.Msg.Params)
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

				result, err := fn(tt.Height, trace.Msg.Params, trace.MsgRct.Return)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.Expected))
			}
		})
	}
}
