package actortest

import (
	"context"
	"fmt"
	"testing"

	"github.com/zondax/fil-parser/actors/metrics"
	metrics2 "github.com/zondax/fil-parser/metrics"

	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
	actorsV1 "github.com/zondax/fil-parser/actors/v1"
	actorsV2 "github.com/zondax/fil-parser/actors/v2"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

var datacapWithParamsAndReturnTests = []struct {
	name   string
	txType string
}{
	{
		name:   "Mint Exported",
		txType: parser.MethodMintExported,
	},
	{
		name:   "Burn Exported",
		txType: parser.MethodBurnExported,
	},
	{
		name:   "Balance Exported",
		txType: parser.MethodBalanceExported,
	},
	{
		name:   "Transfer Exported",
		txType: parser.MethodTransferExported,
	},
	{
		name:   "Transfer From Exported",
		txType: parser.MethodTransferFromExported,
	},
	{
		name:   "Destroy Exported",
		txType: parser.MethodDestroyExported,
	},
	{
		name:   "Increase Allowance Exported",
		txType: parser.MethodIncreaseAllowanceExported,
	},
	{
		name:   "Decrease Allowance Exported",
		txType: parser.MethodDecreaseAllowanceExported,
	},
	{
		name:   "Revoke Allowance Exported",
		txType: parser.MethodRevokeAllowanceExported,
	},
	{
		name:   "Burn From Exported",
		txType: parser.MethodBurnFromExported,
	},
	{
		name:   "Allowance Exported",
		txType: parser.MethodAllowanceExported,
	},
}

var datacapWithParamsOrReturnTests = []struct {
	name   string
	txType string
	key    string
}{
	{
		name:   "Name Exported",
		txType: parser.MethodNameExported,
		key:    parser.ReturnKey,
	},
	{
		name:   "Symbol Exported",
		txType: parser.MethodSymbolExported,
		key:    parser.ReturnKey,
	},
	{
		name:   "Total Supply Exported",
		txType: parser.MethodTotalSupplyExported,
		key:    parser.ReturnKey,
	},
	{
		name:   "Granularity Exported",
		txType: parser.MethodGranularityExported,
		key:    parser.ReturnKey,
	},
}

func TestActorParserV1_DatacapWithParamsAndReturn(t *testing.T) {
	p := getActorParser(actorsV1.NewActorParser).(*actorsV1.ActorParser)

	for _, tt := range datacapWithParamsAndReturnTests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, rawReturn, err := getParamsAndReturn(manifest.DatacapKey, tt.txType)
			require.NoError(t, err)
			require.NotNil(t, rawParams)
			require.NotNil(t, rawReturn)

			got, err := p.ParseDatacap(tt.txType, &parser.LotusMessage{
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

func TestActorParserV1_DatacapWithParamsOrReturn(t *testing.T) {
	p := getActorParser(actorsV1.NewActorParser).(*actorsV1.ActorParser)

	for _, tt := range datacapWithParamsOrReturnTests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, err := loadFile(manifest.DatacapKey, tt.txType, tt.key)
			require.NoError(t, err)
			require.NotNil(t, rawParams)

			msg := &parser.LotusMessage{}
			msgRct := &parser.LotusMessageReceipt{}

			if tt.key == parser.ReturnKey {
				msgRct.Return = rawParams
			} else {
				msg.Params = rawParams
			}

			got, err := p.ParseDatacap(tt.txType, msg, msgRct)

			require.NoError(t, err)
			require.NotNil(t, got)
			require.Contains(t, got, tt.key, fmt.Sprintf("%s could no be found in metadata", tt.key))
			require.NotNil(t, got[tt.key])
		})
	}
}

func TestActorParserV2_DatacapWithParamsAndReturn(t *testing.T) {
	p := getActorParser(actorsV2.NewActorParser).(*actorsV2.ActorParser)
	actor, err := p.GetActor(manifest.DatacapKey, &metrics.ActorsMetricsClient{MetricsClient: metrics2.NewNoopMetricsClient()})
	require.NoError(t, err)
	require.NotNil(t, actor)

	for _, tt := range datacapWithParamsAndReturnTests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, rawReturn, err := getParamsAndReturn(manifest.DatacapKey, tt.txType)
			require.NoError(t, err)
			require.NotNil(t, rawParams)
			require.NotNil(t, rawReturn)

			got, _, err := actor.Parse(context.Background(), network, tools.LatestVersion(network).Height(), tt.txType, &parser.LotusMessage{
				Params: rawParams,
			}, &parser.LotusMessageReceipt{
				Return: rawReturn,
			}, cid.Undef, filTypes.EmptyTSK)
			require.NoError(t, err)
			require.NotNil(t, got)
			require.Contains(t, got, parser.ParamsKey, "Params could no be found in metadata")
			require.NotNil(t, got[parser.ParamsKey])
			require.Contains(t, got, parser.ReturnKey, "Return could no be found in metadata")
			require.NotNil(t, got[parser.ReturnKey])
		})
	}
}

func TestActorParserV2_DatacapWithParamsOrReturn(t *testing.T) {
	p := getActorParser(actorsV2.NewActorParser).(*actorsV2.ActorParser)
	actor, err := p.GetActor(manifest.DatacapKey, &metrics.ActorsMetricsClient{MetricsClient: metrics2.NewNoopMetricsClient()})
	require.NoError(t, err)
	require.NotNil(t, actor)

	for _, tt := range datacapWithParamsOrReturnTests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, err := loadFile(manifest.DatacapKey, tt.txType, tt.key)
			require.NoError(t, err)
			require.NotNil(t, rawParams)

			msg := &parser.LotusMessage{}
			msgRct := &parser.LotusMessageReceipt{}

			if tt.key == parser.ReturnKey {
				msgRct.Return = rawParams
			} else {
				msg.Params = rawParams
			}
			got, _, err := actor.Parse(context.Background(), network, tools.LatestVersion(network).Height(), tt.txType, msg, msgRct, cid.Undef, filTypes.EmptyTSK)

			require.NoError(t, err)
			require.NotNil(t, got)
			require.Contains(t, got, tt.key, fmt.Sprintf("%s could no be found in metadata", tt.key))
			require.NotNil(t, got[tt.key])
		})
	}
}
