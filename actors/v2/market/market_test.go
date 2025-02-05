package market_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/actors/v2/market"
	"github.com/zondax/fil-parser/tools"

	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
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
	os.Exit(m.Run())
}

type testFn func(network string, height int64, rawExported, rawReturn []byte) (map[string]interface{}, error)

func TestMarket(t *testing.T) {
	market := &market.Market{}
	testFns := map[string]testFn{
		"PublishStorageDealsExported":       market.PublishStorageDealsExported,
		"VerifyDealsForActivationExported":  market.VerifyDealsForActivationExported,
		"ActivateDealsExported":             market.ActivateDealsExported,
		"ComputeDataCommitmentExported":     market.ComputeDataCommitmentExported,
		"GetBalanceExported":                market.GetBalanceExported,
		"GetDealDataCommitmentExported":     market.GetDealDataCommitmentExported,
		"GetDealClientExported":             market.GetDealClientExported,
		"GetDealProviderExported":           market.GetDealProviderExported,
		"GetDealLabelExported":              market.GetDealLabelExported,
		"GetDealTermExported":               market.GetDealTermExported,
		"GetDealTotalPriceExported":         market.GetDealTotalPriceExported,
		"GetDealClientCollateralExported":   market.GetDealClientCollateralExported,
		"GetDealProviderCollateralExported": market.GetDealProviderCollateralExported,
		"GetDealVerifiedExported":           market.GetDealVerifiedExported,
		"GetDealActivationExported":         market.GetDealActivationExported,
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
				result, err := market.OnMinerSectorsTerminateExported(tt.Network, tt.Height, trace.Msg.Params)
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
				result, err := market.AddBalance(tt.Network, tt.Height, trace.Msg.Params)
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
				result, err := market.WithdrawBalance(tt.Network, tt.Height, trace.Msg.Params)
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
