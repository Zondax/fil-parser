package actortest

import (
	"fmt"
	"github.com/zondax/fil-parser/actors/metrics"
	metrics2 "github.com/zondax/fil-parser/metrics"
	"testing"

	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
	actorsV1 "github.com/zondax/fil-parser/actors/v1"
	actorsV2 "github.com/zondax/fil-parser/actors/v2"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

var verifiedRegistryWithParamsOrReturnTests = []struct {
	name   string
	txType string
	key    string
}{
	{
		name:   "Add Verifier",
		txType: parser.MethodAddVerifier,
		key:    parser.ParamsKey,
	},
	{
		name:   "Add Verified Client",
		txType: parser.MethodAddVerifiedClient,
		key:    parser.ParamsKey,
	},
	{
		name:   "Add Verified Client Exported",
		txType: parser.MethodAddVerifiedClientExported,
		key:    parser.ParamsKey,
	},
	{
		name:   "Use Bytes",
		txType: parser.MethodUseBytes,
		key:    parser.ParamsKey,
	},
	{
		name:   "Restore Bytes",
		txType: parser.MethodRestoreBytes,
		key:    parser.ParamsKey,
	},
	{
		name:   "Remove Verified Client DataCap",
		txType: parser.MethodRemoveVerifiedClientDataCap,
		key:    parser.ParamsKey,
	},
	{
		name:   "Deprecated1",
		txType: parser.MethodVerifiedDeprecated1,
		key:    parser.ParamsKey,
	},
	{
		name:   "Deprecated2",
		txType: parser.MethodVerifiedDeprecated2,
		key:    parser.ParamsKey,
	},
}

var verifiedRegistryWithParamsAndReturnTests = []struct {
	name   string
	txType string
}{
	{
		name:   "Claim Allocations",
		txType: parser.MethodClaimAllocations,
	},
	{
		name:   "Extend Claim Terms",
		txType: parser.MethodExtendClaimTerms,
	},
	{
		name:   "Extend Claims Terms Exported",
		txType: parser.MethodExtendClaimTermsExported,
	},
	{
		name:   "Universal Receiver Hook",
		txType: parser.MethodMsigUniversalReceiverHook,
	},
	{
		name:   "Remove Expired Allocations",
		txType: parser.MethodRemoveExpiredAllocations,
	},
	{
		name:   "Remove Expired Allocations Exported",
		txType: parser.MethodRemoveExpiredAllocationsExported,
	},
	{
		name:   "Get Claims",
		txType: parser.MethodGetClaims,
	},
	{
		name:   "Get Claims Exported",
		txType: parser.MethodGetClaimsExported,
	},
	{
		name:   "Remove Expired Claims",
		txType: parser.MethodRemoveExpiredClaims,
	},
	{
		name:   "Remove Expired Claims Exported",
		txType: parser.MethodRemoveExpiredClaimsExported,
	},
}

func TestActorParserV1_VerifiedWithParamsOrReturn(t *testing.T) {
	p := getActorParser(actorsV1.NewActorParser).(*actorsV1.ActorParser)

	for _, tt := range verifiedRegistryWithParamsOrReturnTests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, err := loadFile(manifest.VerifregKey, tt.txType, tt.key)
			require.NoError(t, err)
			require.NotNil(t, rawParams)

			msg := &parser.LotusMessage{}
			msgRct := &parser.LotusMessageReceipt{}

			if tt.key == parser.ReturnKey {
				msgRct.Return = rawParams
			} else {
				msg.Params = rawParams
			}

			got, err := p.ParseVerifiedRegistry(tt.txType, msg, msgRct)
			require.NoError(t, err)
			require.NotNil(t, got)
			require.Contains(t, got, tt.key, fmt.Sprintf("%s could no be found in metadata", tt.key))
			require.NotNil(t, got[tt.key])
		})
	}
}

func TestActorParserV1_VerifiedWithParamsAndReturn(t *testing.T) {
	p := getActorParser(actorsV1.NewActorParser).(*actorsV1.ActorParser)

	for _, tt := range verifiedRegistryWithParamsAndReturnTests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, rawReturn, err := getParamsAndReturn(manifest.VerifregKey, tt.txType)
			require.NoError(t, err)
			require.NotNil(t, rawParams)
			require.NotNil(t, rawReturn)

			got, err := p.ParseVerifiedRegistry(tt.txType, &parser.LotusMessage{
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

func TestActorParserV2_VerifiedWithParamsOrReturn(t *testing.T) {
	p := getActorParser(actorsV2.NewActorParser).(*actorsV2.ActorParser)
	actor, err := p.GetActor(manifest.VerifregKey, &metrics.ActorsMetricsClient{MetricsClient: metrics2.UnimplementedMetricsClient{}})
	require.NoError(t, err)
	require.NotNil(t, actor)

	for _, tt := range verifiedRegistryWithParamsOrReturnTests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, err := loadFile(manifest.VerifregKey, tt.txType, tt.key)
			require.NoError(t, err)
			require.NotNil(t, rawParams)

			msg := &parser.LotusMessage{}
			msgRct := &parser.LotusMessageReceipt{}

			if tt.key == parser.ReturnKey {
				msgRct.Return = rawParams
			} else {
				msg.Params = rawParams
			}

			got, _, err := actor.Parse(network, tools.LatestVersion.Height(), tt.txType, msg, msgRct, cid.Undef, filTypes.EmptyTSK)

			require.NoError(t, err)
			require.NotNil(t, got)
			require.Contains(t, got, tt.key, fmt.Sprintf("%s could no be found in metadata", tt.key))
			require.NotNil(t, got[tt.key])
		})
	}
}

func TestActorParserV2_VerifiedWithParamsAndReturn(t *testing.T) {
	p := getActorParser(actorsV2.NewActorParser).(*actorsV2.ActorParser)
	actor, err := p.GetActor(manifest.VerifregKey, &metrics.ActorsMetricsClient{MetricsClient: metrics2.UnimplementedMetricsClient{}})
	require.NoError(t, err)
	require.NotNil(t, actor)

	for _, tt := range verifiedRegistryWithParamsAndReturnTests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, rawReturn, err := getParamsAndReturn(manifest.VerifregKey, tt.txType)
			require.NoError(t, err)
			require.NotNil(t, rawParams)
			require.NotNil(t, rawReturn)

			got, _, err := actor.Parse(network, tools.V20.Height(), tt.txType, &parser.LotusMessage{
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
