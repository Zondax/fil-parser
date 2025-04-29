package actortest

import (
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/zondax/fil-parser/metrics"
	"net/http"
	"os"

	"github.com/filecoin-project/lotus/api/client"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/actors/cache"
	"github.com/zondax/fil-parser/actors/cache/impl/common"
	"github.com/zondax/fil-parser/parser"
	helper2 "github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/fil-parser/types"
	"github.com/zondax/golem/pkg/logger"
	rosettaFilecoinLib "github.com/zondax/rosetta-filecoin-lib"
)

const (
	network  = "mainnet"
	dataPath = "../../data/actors"
	testUrl  = "https://api.zondax.ch/fil/node/mainnet/rpc/v1"
	height   = int64(0)
)

func getActorParser(actorParserFn any) actors.ActorParserInterface {
	lotusClient, _, err := client.NewFullNodeRPCV1(context.Background(), testUrl, http.Header{})
	if err != nil {
		return nil
	}
	actorsCache, err := cache.SetupActorsCache(common.DataSource{
		Node: lotusClient,
	}, nil, metrics.NewNoopMetricsClient())

	if err != nil {
		return nil
	}

	lib := rosettaFilecoinLib.NewRosettaConstructionFilecoin(lotusClient)
	helper := helper2.NewHelper(lib, actorsCache, lotusClient, nil, metrics.NewNoopMetricsClient())
	gLogger := logger.NewDevelopmentLogger()
	switch fn := actorParserFn.(type) {
	case func(*helper2.Helper, *logger.Logger) actors.ActorParserInterface:
		return fn(helper, gLogger)
	case func(string, *helper2.Helper, *logger.Logger) actors.ActorParserInterface:
		return fn(network, helper, gLogger)
	case func(*helper2.Helper, *logger.Logger, metrics.MetricsClient) actors.ActorParserInterface:
		return fn(helper, gLogger, metrics.NewNoopMetricsClient())
	case func(string, *helper2.Helper, *logger.Logger, metrics.MetricsClient) actors.ActorParserInterface:
		return fn(network, helper, gLogger, metrics.NewNoopMetricsClient())
	default:
		panic(fmt.Sprintf("invalid actor parser function: %T", actorParserFn))
	}
}

func getParamsAndReturn(actor, txType string) ([]byte, []byte, error) {
	rawParams, err := loadFile(actor, txType, parser.ParamsKey)
	if err != nil {
		return nil, nil, err
	}
	rawReturn, err := loadFile(actor, txType, parser.ReturnKey)
	if err != nil {
		return nil, nil, err
	}
	return rawParams, rawReturn, nil
}

func loadFile(actor, txType, name string) ([]byte, error) {
	item, err := os.ReadFile(fmt.Sprintf("%s/%s/%s/%s", dataPath, actor, txType, name))
	if err != nil {
		return nil, err
	}
	return item, nil
}

func deserializeMessage(actor, txType string) (*parser.LotusMessage, error) {
	file, err := os.Open(fmt.Sprintf("%s/%s/%s/%s", dataPath, actor, txType, "Message"))
	if err != nil {
		return nil, err
	}
	decoder := gob.NewDecoder(file)
	message := &parser.LotusMessage{}
	err = decoder.Decode(message)
	return message, err
}

func deserializeTipset(actor, txType string) (t filTypes.TipSet, err error) {
	file, err := os.Open(fmt.Sprintf("%s/%s/%s/%s", dataPath, actor, txType, "Tipset"))
	if err != nil {
		return
	}
	err = t.UnmarshalCBOR(file)
	return
}

func getEthLogs(actor, txType string) ([]types.EthLog, error) {
	file, err := os.ReadFile(fmt.Sprintf("%s/%s/%s/%s", dataPath, actor, txType, parser.EthLogsKey))
	if err != nil {
		return nil, err
	}
	var ethLogs []types.EthLog
	err = json.Unmarshal(file, &ethLogs)
	return ethLogs, err
}
