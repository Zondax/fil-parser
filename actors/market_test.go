package actors

import (
	"fmt"
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/parser"
	"testing"
)

func TestActorParser_marketWithParamsOrReturn(t *testing.T) {
	p := getActorParser()
	tests := []struct {
		name   string
		txType string
		f      func([]byte) (map[string]interface{}, error)
		key    string
	}{
		{
			name:   "Add Balance",
			txType: parser.MethodAddBalance,
			f:      p.addBalance,
			key:    parser.ParamsKey,
		},
		{
			name:   "Add Balance Exported",
			txType: parser.MethodAddBalanceExported,
			f:      p.addBalance,
			key:    parser.ParamsKey,
		},
		{
			name:   "On Miner Sector Terminate",
			txType: parser.MethodOnMinerSectorsTerminate,
			f:      p.onMinerSectorsTerminate,
			key:    parser.ParamsKey,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, err := loadFile(manifest.MarketKey, tt.txType, tt.key)
			require.NoError(t, err)
			require.NotNil(t, rawParams)

			got, err := tt.f(rawParams)
			require.NoError(t, err)
			require.NotNil(t, got)
			require.Contains(t, got, tt.key, fmt.Sprintf("%s could no be found in metadata", tt.key))
			require.NotNil(t, got[tt.key])
		})
	}
}

func TestActorParser_marketWithParamsAndReturn(t *testing.T) {
	p := getActorParser()
	tests := []struct {
		name   string
		txType string
		f      func([]byte, []byte) (map[string]interface{}, error)
	}{
		{
			name:   "Publish Storage Deals",
			txType: parser.MethodPublishStorageDeals,
			f:      p.publishStorageDeals,
		},
		{
			name:   "Publish Storage Deals Exported",
			txType: parser.MethodPublishStorageDealsExported,
			f:      p.publishStorageDeals,
		},
		{
			name:   "Verify Deals For Activation",
			txType: parser.MethodVerifyDealsForActivation,
			f:      p.verifyDealsForActivation,
		},
		/*
			{ // TODO: cbor input had wrong number of fields
				name:   "Compute Data Commitment",
				txType: parser.MethodComputeDataCommitment,
				f:      p.computeDataCommitment,
			},
		*/
		{
			name:   "Get Deal Activation",
			txType: parser.MethodGetDealActivation,
			f:      p.getDealActivation,
		},
		{
			name:   "Activate Deals",
			txType: parser.MethodActivateDeals,
			f:      p.activateDeals,
		},
		{
			name:   "Withdraw Balance",
			txType: parser.MethodWithdrawBalance,
			f:      p.withdrawBalance,
		},
		{
			name:   "Withdraw Balance Exported",
			txType: parser.MethodWithdrawBalanceExported,
			f:      p.withdrawBalance,
		},
		{
			name:   "Get Balance",
			txType: parser.MethodGetBalance,
			f:      p.getBalance,
		},
		/*
			{
				name:   "Get Deal Data Commitment",
				txType: parser.MethodGetDealDataCommitment,
				f:      p.getDealDataCommitment,
			},
		*/
		{
			name:   "Get Deal Client Exported",
			txType: parser.MethodGetDealClient,
			f:      p.getDealClient,
		},
		{
			name:   "Get Deal Provided Exported",
			txType: parser.MethodGetDealProvider,
			f:      p.getDealProvider,
		},
		/*
			{
				name:   "Get Deal Label",
				txType: parser.MethodGetDealLabel,
				f:      p.getDealProvider,
			},
		*/
		{
			name:   "Get Deal Term",
			txType: parser.MethodGetDealTerm,
			f:      p.getDealTerm,
		},
		{
			name:   "Get Deal Total Price",
			txType: parser.MethodGetDealTotalPrice,
			f:      p.getDealTotalPrice,
		},
		{
			name:   "Get Deal Client Collateral",
			txType: parser.MethodGetDealClientCollateral,
			f:      p.getDealClientCollateral,
		},
		{
			name:   "Get Deal Provider Collateral",
			txType: parser.MethodGetDealProviderCollateral,
			f:      p.getDealProviderCollateral,
		},
		{
			name:   "Get Deal Verified",
			txType: parser.MethodGetDealVerified,
			f:      p.getDealVerified,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, rawReturn, err := getParmasAndReturn(manifest.MarketKey, tt.txType)
			require.NoError(t, err)
			require.NotNil(t, rawParams)
			require.NotNil(t, rawReturn)

			got, err := tt.f(rawParams, rawReturn)
			require.NoError(t, err)
			require.NotNil(t, got)
			require.Contains(t, got, parser.ParamsKey, "Params could no be found in metadata")
			require.NotNil(t, got[parser.ParamsKey])
			require.Contains(t, got, parser.ReturnKey, "Return could no be found in metadata")
			require.NotNil(t, got[parser.ReturnKey])
		})
	}
}
