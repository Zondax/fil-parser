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

func TestActorParserV1_Cron(t *testing.T) {
	p := getActorParser(actorsV1.NewActorParser).(*actorsV1.ActorParser)

	rawParams, err := loadFile(manifest.CronKey, parser.MethodConstructor, parser.ParamsKey)
	require.NoError(t, err)
	require.NotNil(t, rawParams)

	got, err := p.ParseCron(parser.MethodConstructor, &parser.LotusMessage{
		Params: rawParams,
	}, &parser.LotusMessageReceipt{
		Return: nil,
	})
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Contains(t, got, parser.ParamsKey, fmt.Sprintf("%s could no be found in metadata", parser.ParamsKey))
	require.NotNil(t, got[parser.ParamsKey])
}

func TestActorParserV2_Cron(t *testing.T) {
	p := getActorParser(actorsV2.NewActorParser).(*actorsV2.ActorParser)
	actor, err := p.GetActor(manifest.CronKey, &metrics.ActorsMetricsClient{MetricsClient: metrics2.NewNoopMetricsClient()})
	require.NoError(t, err)
	require.NotNil(t, actor)

	rawParams, err := loadFile(manifest.CronKey, parser.MethodConstructor, parser.ParamsKey)
	require.NoError(t, err)
	require.NotNil(t, rawParams)

	got, _, err := actor.Parse(context.Background(), network, tools.LatestVersion(network).Height(), parser.MethodConstructor, &parser.LotusMessage{
		Params: rawParams,
	}, &parser.LotusMessageReceipt{
		Return: nil,
	}, cid.Undef, filTypes.EmptyTSK)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Contains(t, got, parser.ParamsKey, fmt.Sprintf("%s could no be found in metadata", parser.ParamsKey))
	require.NotNil(t, got[parser.ParamsKey])
}
