package actortest

import (
	"context"
	"fmt"
	"testing"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/actors/metrics"
	metrics2 "github.com/zondax/fil-parser/metrics"

	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
	actorsV1 "github.com/zondax/fil-parser/actors/v1"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

var initWithParamsOrReturnTests = []struct {
	name   string
	txType string
	key    string
}{
	{
		name:   "Constructor",
		txType: parser.MethodConstructor,
		key:    parser.ParamsKey,
	},
}

var execTests = []struct {
	name   string
	txType string
}{
	{
		name:   "Exec",
		txType: parser.MethodExec,
	},
	{
		name:   "Exec4",
		txType: parser.MethodExec4,
	},
}

func TestActorParserV1_InitWithParamsOrReturn(t *testing.T) {
	p := getActorParser(actorsV1.NewActorParser).(*actorsV1.ActorParser)

	for _, tt := range initWithParamsOrReturnTests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, err := loadFile(manifest.InitKey, tt.txType, tt.key)
			require.NoError(t, err)
			require.NotNil(t, rawParams)

			got, _, err := p.ParseInit(tt.txType, &parser.LotusMessage{
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

func TestActorParserV1_Exec(t *testing.T) {
	p := getActorParser(actorsV1.NewActorParser).(*actorsV1.ActorParser)

	for _, tt := range execTests {
		t.Run(tt.name, func(t *testing.T) {
			rawReturn, err := loadFile(manifest.InitKey, tt.txType, parser.ReturnKey)
			require.NoError(t, err)
			msg, err := deserializeMessage(manifest.InitKey, tt.txType)
			require.NoError(t, err)
			require.NotNil(t, msg)
			got, addr, err := p.ParseInit(tt.txType, msg, &parser.LotusMessageReceipt{
				Return: rawReturn,
			})
			require.NoError(t, err)
			require.NotNil(t, got)
			require.NotNil(t, addr)
		})
	}
}

func TestActorParserV2_InitWithParamsOrReturn(t *testing.T) {
	p := getActorParser(actors.NewActorParser).(*actors.ActorParser)
	actor, err := p.GetActor(manifest.InitKey, &metrics.ActorsMetricsClient{MetricsClient: metrics2.NewNoopMetricsClient()})
	require.NoError(t, err)
	require.NotNil(t, actor)

	for _, tt := range initWithParamsOrReturnTests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, err := loadFile(manifest.InitKey, tt.txType, tt.key)
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

func TestActorParserV2_Exec(t *testing.T) {
	p := getActorParser(actors.NewActorParser).(*actors.ActorParser)
	actor, err := p.GetActor(manifest.InitKey, &metrics.ActorsMetricsClient{MetricsClient: metrics2.NewNoopMetricsClient()})
	require.NoError(t, err)
	require.NotNil(t, actor)

	for _, tt := range execTests {
		t.Run(tt.name, func(t *testing.T) {
			rawReturn, err := loadFile(manifest.InitKey, tt.txType, parser.ReturnKey)
			require.NoError(t, err)
			msg, err := deserializeMessage(manifest.InitKey, tt.txType)
			require.NoError(t, err)
			require.NotNil(t, msg)

			got, addr, err := actor.Parse(context.Background(), network, tools.LatestVersion(network).Height(), tt.txType, &parser.LotusMessage{
				Params: msg.Params,
			}, &parser.LotusMessageReceipt{
				Return: rawReturn,
			}, cid.Undef, filTypes.EmptyTSK)
			require.NoError(t, err)
			require.NotNil(t, got)
			require.NotNil(t, addr)
		})
	}
}
