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

func TestMarket(t *testing.T) {
	market := &market.Market{}
	testFns := map[string]testFn{
		"PublishStorageDeals":       market.PublishStorageDealsParams,
		"VerifyDealsForActivation":  market.VerifyDealsForActivationParams,
		"ActivateDeals":             market.ActivateDealsParams,
		"ComputeDataCommitment":     market.ComputeDataCommitmentParams,
		"GetBalance":                market.GetBalanceParams,
		"GetDealDataCommitment":     market.GetDealDataCommitmentParams,
		"GetDealClient":             market.GetDealClientParams,
		"GetDealProvider":           market.GetDealProviderParams,
		"GetDealLabel":              market.GetDealLabelParams,
		"GetDealTerm":               market.GetDealTermParams,
		"GetDealTotalPrice":         market.GetDealTotalPriceParams,
		"GetDealClientCollateral":   market.GetDealClientCollateralParams,
		"GetDealProviderCollateral": market.GetDealProviderCollateralParams,
		"GetDealVerified":           market.GetDealVerifiedParams,
		"GetDealActivation":         market.GetDealActivationParams,
	}
	for name, fn := range testFns {
		t.Run(name, func(t *testing.T) {
			tests, err := tools.LoadTestData[map[string]any](network, name, expected)
			require.NoError(t, err)
			runTest(t, fn, tests)
		})
	}
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
				market := &market.Market{}
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
				market := &market.Market{}
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
				market := &market.Market{}
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
