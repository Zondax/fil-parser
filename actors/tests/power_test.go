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

var powerWithParamsOrReturnTests = []struct {
	name   string
	txType string
	key    string
}{
	{
		name:   "Constructor",
		txType: parser.MethodConstructor,
		key:    parser.ParamsKey,
	},
	{
		name:   "Current Total Power",
		txType: parser.MethodCurrentTotalPower,
		key:    parser.ReturnKey,
	},
	{
		name:   "Enroll Cron Event",
		txType: parser.MethodEnrollCronEvent,
		key:    parser.ParamsKey,
	},
	{
		name:   "Submit PoRep For Bulk Verify",
		txType: parser.MethodSubmitPoRepForBulkVerify,
		key:    parser.ParamsKey,
	},
	{
		name:   "Update Claimed Power",
		txType: parser.MethodUpdateClaimedPower,
		key:    parser.ParamsKey,
	},
	{
		name:   "Update Pledge Total",
		txType: parser.MethodUpdatePledgeTotal,
		key:    parser.ParamsKey,
	},
	{
		name:   "Network Raw Power Exported",
		txType: parser.MethodNetworkRawPowerExported,
		key:    parser.ReturnKey,
	},
	{
		name:   "Miner Count Exported",
		txType: parser.MethodMinerCountExported,
		key:    parser.ReturnKey,
	},
	{
		name:   "Miner Consensus Count Exported",
		txType: parser.MethodMinerConsensusCountExported,
		key:    parser.ReturnKey,
	},
}

var powerWithParamsAndReturnTests = []struct {
	name   string
	txType string
}{
	{
		name:   "Miner Raw Power Exported",
		txType: parser.MethodMinerRawPowerExported,
	},
}
var powerCreateMinerTests = []struct {
	name   string
	method string
}{
	{
		name:   "Create Miner",
		method: parser.MethodCreateMiner,
	},
	{
		name:   "Create Miner Exported",
		method: parser.MethodCreateMinerExported,
	},
}

func TestActorParserV1_PowerWithParamsOrReturn(t *testing.T) {
	p := getActorParser(actorsV1.NewActorParser).(*actorsV1.ActorParser)

	for _, tt := range powerWithParamsOrReturnTests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, err := loadFile(manifest.PowerKey, tt.txType, tt.key)
			require.NoError(t, err)
			require.NotNil(t, rawParams)

			msg := &parser.LotusMessage{}
			msgRct := &parser.LotusMessageReceipt{}

			if tt.key == parser.ReturnKey {
				msgRct.Return = rawParams
			} else {
				msg.Params = rawParams
			}

			got, _, err := p.ParseStoragepower(tt.txType, msg, msgRct)
			require.NoError(t, err)
			require.NotNil(t, got)
			require.Contains(t, got, tt.key, fmt.Sprintf("%s could no be found in metadata", tt.key))
			require.NotNil(t, got[tt.key])
		})
	}
}

func TestActorParserV1_PowerWithParamsAndReturn(t *testing.T) {
	p := getActorParser(actorsV1.NewActorParser).(*actorsV1.ActorParser)

	for _, tt := range powerWithParamsAndReturnTests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, rawReturn, err := getParamsAndReturn(manifest.PowerKey, tt.txType)
			require.NoError(t, err)
			require.NotNil(t, rawParams)
			require.NotNil(t, rawReturn)

			got, _, err := p.ParseStoragepower(tt.txType, &parser.LotusMessage{
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

func TestActorParserV1_ParseCreateMiner(t *testing.T) {
	p := getActorParser(actorsV1.NewActorParser).(*actorsV1.ActorParser)

	for _, tt := range powerCreateMinerTests {
		t.Run(tt.name, func(t *testing.T) {
			rawReturn, err := loadFile(manifest.PowerKey, tt.method, parser.ReturnKey)
			require.NoError(t, err)
			require.NotNil(t, rawReturn)

			msg, err := deserializeMessage(manifest.PowerKey, tt.method)
			require.NoError(t, err)
			require.NotNil(t, msg)

			got, addr, err := p.ParseStoragepower(tt.method, msg, &parser.LotusMessageReceipt{
				Return: rawReturn,
			})
			require.NoError(t, err)
			require.NotNil(t, got)
			require.NotNil(t, addr)
			require.Contains(t, got, parser.ReturnKey)
		})
	}
}

func TestActorParserV2_PowerWithParamsOrReturn(t *testing.T) {
	p := getActorParser(actors.NewActorParser).(*actors.ActorParser)
	actor, err := p.GetActor(manifest.PowerKey, &metrics.ActorsMetricsClient{MetricsClient: metrics2.NewNoopMetricsClient()})
	require.NoError(t, err)
	require.NotNil(t, actor)

	for _, tt := range powerWithParamsOrReturnTests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, err := loadFile(manifest.PowerKey, tt.txType, tt.key)
			require.NoError(t, err)
			require.NotNil(t, rawParams)

			msg := &parser.LotusMessage{}
			msgRct := &parser.LotusMessageReceipt{}

			if tt.key == parser.ReturnKey {
				msgRct.Return = rawParams
			} else {
				msg.Params = rawParams
			}

			got, _, err := actor.Parse(context.Background(), network, tools.V20.Height(), tt.txType, msg, msgRct, cid.Undef, filTypes.EmptyTSK)
			require.NoError(t, err)
			require.NotNil(t, got)
			require.Contains(t, got, tt.key, fmt.Sprintf("%s could no be found in metadata", tt.key))
			require.NotNil(t, got[tt.key])
		})
	}
}

func TestActorParserV2_PowerWithParamsAndReturn(t *testing.T) {
	p := getActorParser(actors.NewActorParser).(*actors.ActorParser)
	actor, err := p.GetActor(manifest.PowerKey, &metrics.ActorsMetricsClient{MetricsClient: metrics2.NewNoopMetricsClient()})
	require.NoError(t, err)
	require.NotNil(t, actor)

	for _, tt := range powerWithParamsAndReturnTests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, rawReturn, err := getParamsAndReturn(manifest.PowerKey, tt.txType)
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

func TestActorParserV2_ParseCreateMiner(t *testing.T) {
	p := getActorParser(actors.NewActorParser).(*actors.ActorParser)
	actor, err := p.GetActor(manifest.PowerKey, &metrics.ActorsMetricsClient{MetricsClient: metrics2.NewNoopMetricsClient()})
	require.NoError(t, err)
	require.NotNil(t, actor)

	for _, tt := range powerCreateMinerTests {
		t.Run(tt.name, func(t *testing.T) {
			rawReturn, err := loadFile(manifest.PowerKey, tt.method, parser.ReturnKey)
			require.NoError(t, err)
			require.NotNil(t, rawReturn)

			msg, err := deserializeMessage(manifest.PowerKey, tt.method)
			require.NoError(t, err)
			require.NotNil(t, msg)

			got, addr, err := actor.Parse(context.Background(), network, tools.LatestVersion(network).Height(), tt.method, msg, &parser.LotusMessageReceipt{
				Return: rawReturn,
			}, cid.Undef, filTypes.EmptyTSK)

			require.NoError(t, err)
			require.NotNil(t, got)
			require.NotNil(t, addr)
			require.Contains(t, got, parser.ReturnKey)
		})
	}
}
