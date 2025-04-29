package actortest

import (
	"context"
	"testing"

	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/stretchr/testify/require"
	actorsV1 "github.com/zondax/fil-parser/actors/v1"
	actorsV2 "github.com/zondax/fil-parser/actors/v2"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

var eamTests = []struct {
	name   string
	txType string
}{
	{
		name:   "Create",
		txType: parser.MethodCreate,
	},
	{
		name:   "Create2",
		txType: parser.MethodCreate2,
	},
	{
		name:   "Create External",
		txType: parser.MethodCreateExternal,
	},
}

func TestActorParserV1_EamCreates(t *testing.T) {
	p := getActorParser(actorsV1.NewActorParser).(*actorsV1.ActorParser)

	for _, tt := range eamTests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, rawReturn, err := getParamsAndReturn(manifest.EamKey, tt.txType)
			require.NoError(t, err)
			require.NotNil(t, rawParams)
			require.NotNil(t, rawReturn)

			msg, err := deserializeMessage(manifest.EamKey, tt.txType)
			require.NoError(t, err)
			require.NotNil(t, msg)
			got, _, err := p.ParseEam(tt.txType, msg, &parser.LotusMessageReceipt{
				Return: rawReturn,
			}, msg.Cid)
			require.NoError(t, err)
			require.NotNil(t, got)
			require.Contains(t, got, parser.ParamsKey, "Params could no be found in metadata")
			require.NotNil(t, got[parser.ParamsKey])
			require.Contains(t, got, parser.ReturnKey, "Return could no be found in metadata")
			require.NotNil(t, got[parser.ReturnKey])
		})
	}
}

func TestActorParserV2_EamCreates(t *testing.T) {
	p := getActorParser(actorsV2.NewActorParser).(*actorsV2.ActorParser)
	actor, err := p.GetActor(manifest.EamKey)
	require.NoError(t, err)
	require.NotNil(t, actor)

	for _, tt := range eamTests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, rawReturn, err := getParamsAndReturn(manifest.EamKey, tt.txType)
			require.NoError(t, err)
			require.NotNil(t, rawParams)
			require.NotNil(t, rawReturn)

			msg, err := deserializeMessage(manifest.EamKey, tt.txType)
			require.NoError(t, err)
			require.NotNil(t, msg)

			got, _, err := actor.Parse(context.Background(), network, tools.LatestVersion(network).Height(), tt.txType, msg, &parser.LotusMessageReceipt{
				Return: rawReturn,
			}, msg.Cid, filTypes.EmptyTSK)
			require.NoError(t, err)
			require.NotNil(t, got)
			require.Contains(t, got, parser.ParamsKey, "Params could no be found in metadata")
			require.NotNil(t, got[parser.ParamsKey])
			require.Contains(t, got, parser.ReturnKey, "Return could no be found in metadata")
			require.NotNil(t, got[parser.ReturnKey])
		})
	}
}
