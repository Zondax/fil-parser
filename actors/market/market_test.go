package market_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/actors/market"
	"github.com/zondax/fil-parser/tools"

	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
)

type testFn func(height int64, rawParams, rawReturn []byte) (map[string]interface{}, error)

type test struct {
	name     string
	version  string
	url      string
	height   int64
	expected map[string]any
}

func TestPublishStorageDeals(t *testing.T) {
	tests := []test{}

	runTest(t, market.PublishStorageDealsParams, tests)
}

func TestVerifyDealsForActivation(t *testing.T) {
	tests := []test{}
	runTest(t, market.VerifyDealsForActivationParams, tests)
}

func TestActivateDeals(t *testing.T) {
	tests := []test{}
	runTest(t, market.ActivateDealsParams, tests)
}

func TestComputeDataCommitment(t *testing.T) {
	tests := []test{}
	runTest(t, market.ComputeDataCommitmentParams, tests)
}

func TestGetBalance(t *testing.T) {
	tests := []test{}
	runTest(t, market.GetBalanceParams, tests)
}

func TestGetDealDataCommitment(t *testing.T) {
	tests := []test{}
	runTest(t, market.GetDealDataCommitmentParams, tests)
}

func TestGetDealClient(t *testing.T) {
	tests := []test{}
	runTest(t, market.GetDealClientParams, tests)
}

func TestGetDealProvider(t *testing.T) {
	tests := []test{}
	runTest(t, market.GetDealProviderParams, tests)
}

func TestGetDealLabel(t *testing.T) {
	tests := []test{}
	runTest(t, market.GetDealLabelParams, tests)
}

func TestGetDealTerm(t *testing.T) {
	tests := []test{}
	runTest(t, market.GetDealTermParams, tests)
}

func TestGetDealTotalPrice(t *testing.T) {
	tests := []test{}
	runTest(t, market.GetDealTotalPriceParams, tests)
}

func TestGetDealClientCollateral(t *testing.T) {
	tests := []test{}
	runTest(t, market.GetDealClientCollateralParams, tests)
}

func TestGetDealProviderCollateral(t *testing.T) {
	tests := []test{}
	runTest(t, market.GetDealProviderCollateralParams, tests)
}

func TestGetDealVerified(t *testing.T) {
	tests := []test{}
	runTest(t, market.GetDealVerifiedParams, tests)
}

func TestGetDealActivation(t *testing.T) {
	tests := []test{}
	runTest(t, market.GetDealActivationParams, tests)
}

func TestDealProviderCollateral(t *testing.T) {
	tests := []test{}
	runTest(t, market.GetDealProviderCollateralParams, tests)
}

func TestGetDealVerifiedParams(t *testing.T) {
	tests := []test{}
	runTest(t, market.GetDealVerifiedParams, tests)
}

func TestGetDealActivationParams(t *testing.T) {
	tests := []test{}
	runTest(t, market.GetDealActivationParams, tests)
}

func TestOnMinerSectorsTerminate(t *testing.T) {
	tests := []test{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.height, tt.version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}

				result, err := market.OnMinerSectorsTerminateParams(tt.height, trace.Msg.Params)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.expected))
			}
		})
	}
}

func TestParseAddBalance(t *testing.T) {
	tests := []test{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.height, tt.version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}

				result, err := market.ParseAddBalance(tt.height, trace.Msg.Params)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.expected))
			}
		})
	}
}

func TestParseWithdrawBalance(t *testing.T) {
	tests := []test{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.height, tt.version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}

				result, err := market.ParseWithdrawBalance(tt.height, trace.Msg.Params)
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

				result, err := fn(tt.height, trace.Msg.Params, trace.MsgRct.Return)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.expected))
			}
		})
	}
}
