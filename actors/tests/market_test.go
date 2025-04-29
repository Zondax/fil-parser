package actortest

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/actors"
	actorsV1 "github.com/zondax/fil-parser/actors/v1"
	actorsV2 "github.com/zondax/fil-parser/actors/v2"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

var marketWithParamsOrReturnTests = []struct {
	name   string
	txType string

	key string
}{
	{
		name:   "Add Balance",
		txType: parser.MethodAddBalance,
		key:    parser.ParamsKey,
	},
	{
		name:   "Add Balance Exported",
		txType: parser.MethodAddBalanceExported,
		key:    parser.ParamsKey,
	},
	{
		name:   "On Miner Sector Terminate",
		txType: parser.MethodOnMinerSectorsTerminate,
		key:    parser.ParamsKey,
	},
}

var marketWithParamsAndReturnTests = []struct {
	name    string
	txType  string
	version string
}{
	{
		name:   "Publish Storage Deals",
		txType: parser.MethodPublishStorageDeals,
		// version: tools.V20.String(),
	},
	{
		name:   "Publish Storage Deals Exported",
		txType: parser.MethodPublishStorageDealsExported,
	},
	{
		name:   "Verify Deals For Activation",
		txType: parser.MethodVerifyDealsForActivation,
	},
	{
		name:   "Compute Data Commitment",
		txType: parser.MethodComputeDataCommitment,
	},
	{
		name:   "Get Deal Activation",
		txType: parser.MethodGetDealActivation,
	},
	{
		name:   "Activate Deals",
		txType: parser.MethodActivateDeals,
	},
	{
		name:   "Withdraw Balance",
		txType: parser.MethodWithdrawBalance,
	},
	{
		name:   "Withdraw Balance Exported",
		txType: parser.MethodWithdrawBalanceExported,
	},
	{
		name:   "Get Balance",
		txType: parser.MethodGetBalance,
	},
	{
		name:   "Get Deal Data Commitment",
		txType: parser.MethodGetDealDataCommitment,
	},
	{
		name:   "Get Deal Client Exported",
		txType: parser.MethodGetDealClient,
	},
	{
		name:   "Get Deal Provided Exported",
		txType: parser.MethodGetDealProvider,
	},
	{
		name:   "Get Deal Label",
		txType: parser.MethodGetDealLabel,
	},
	{
		name:   "Get Deal Term",
		txType: parser.MethodGetDealTerm,
	},
	{
		name:   "Get Deal Total Price",
		txType: parser.MethodGetDealTotalPrice,
	},
	{
		name:   "Get Deal Client Collateral",
		txType: parser.MethodGetDealClientCollateral,
	},
	{
		name:   "Get Deal Provider Collateral",
		txType: parser.MethodGetDealProviderCollateral,
	},
	{
		name:   "Get Deal Verified",
		txType: parser.MethodGetDealVerified,
	},
}

func TestActorParserV1_MarketWithParamsOrReturn(t *testing.T) {
	p := getActorParser(actorsV1.NewActorParser).(*actorsV1.ActorParser)

	for _, tt := range marketWithParamsOrReturnTests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, err := loadFile(manifest.MarketKey, tt.txType, tt.key)
			require.NoError(t, err)
			require.NotNil(t, rawParams)

			got, err := p.ParseStoragemarket(tt.txType, &parser.LotusMessage{
				Params: rawParams,
			}, &parser.LotusMessageReceipt{
				Return: nil,
			})
			require.NoError(t, err)
			require.NotNil(t, got)
			require.Contains(t, got, tt.key, fmt.Sprintf("%s could no be found in metadata", tt.key))
			require.NotNil(t, got[tt.key])
		})
	}
}

func TestActorParserV1_MarketWithParamsAndReturn(t *testing.T) {
	p := getActorParser(actorsV1.NewActorParser).(*actorsV1.ActorParser)

	for _, tt := range marketWithParamsAndReturnTests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, rawReturn, err := getParamsAndReturn(manifest.MarketKey, tt.txType)
			require.NoError(t, err)
			require.NotNil(t, rawParams)
			require.NotNil(t, rawReturn)

			got, err := p.ParseStoragemarket(tt.txType, &parser.LotusMessage{
				Params: rawParams,
			}, &parser.LotusMessageReceipt{
				Return: rawReturn,
			})
			require.NoError(t, err)
			require.NotNil(t, got)
			require.Contains(t, got, parser.ParamsKey, "Params could no be found in metadata")
			require.NotNil(t, got[parser.ParamsKey])
			require.Contains(t, got, parser.ReturnKey, "Return could no be found in metadata")
			require.NotNil(t, got[parser.ReturnKey])
		})
	}
}

func TestActorParserV2_MarketWithParamsOrReturn(t *testing.T) {
	p := getActorParser(actorsV2.NewActorParser).(*actorsV2.ActorParser)
	actor, err := p.GetActor(manifest.MarketKey)
	require.NoError(t, err)
	require.NotNil(t, actor)

	for _, tt := range marketWithParamsOrReturnTests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, err := loadFile(manifest.MarketKey, tt.txType, tt.key)
			require.NoError(t, err)
			require.NotNil(t, rawParams)

			got, _, err := actor.Parse(context.Background(), network, tools.LatestVersion(network).Height(), tt.txType, &parser.LotusMessage{
				Params: rawParams,
			}, &parser.LotusMessageReceipt{
				Return: nil,
			}, cid.Undef, filTypes.EmptyTSK)
			require.NoError(t, err)
			require.NotNil(t, got)
			require.Contains(t, got, tt.key, fmt.Sprintf("%s could no be found in metadata", tt.key))
			require.NotNil(t, got[tt.key])
		})
	}
}

func TestActorParserV2_MarketWithParamsAndReturn(t *testing.T) {
	p := getActorParser(actorsV2.NewActorParser).(*actorsV2.ActorParser)
	actor, err := p.GetActor(manifest.MarketKey)
	require.NoError(t, err)
	require.NotNil(t, actor)

	for _, tt := range marketWithParamsAndReturnTests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, rawReturn, err := getParamsAndReturn(manifest.MarketKey, tt.txType)
			require.NoError(t, err)
			require.NotNil(t, rawParams)
			require.NotNil(t, rawReturn)

			got, _, err := actor.Parse(context.Background(), network, tools.V20.Height(), tt.txType, &parser.LotusMessage{
				Params: rawParams,
			}, &parser.LotusMessageReceipt{
				Return: rawReturn,
			}, cid.Undef, filTypes.EmptyTSK)

			if errors.Is(err, actors.ErrInvalidHeightForMethod) {
				t.Skipf("skipping %s because of unsupported height", tt.name)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, got)
			require.Contains(t, got, parser.ParamsKey, "Params could no be found in metadata")
			require.NotNil(t, got[parser.ParamsKey])
			require.Contains(t, got, parser.ReturnKey, "Return could no be found in metadata")
			require.NotNil(t, got[parser.ReturnKey])
		})
	}
}
