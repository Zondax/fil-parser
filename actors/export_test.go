package actors

import (
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	logger2 "github.com/zondax/golem/pkg/logger"

	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/zondax/fil-parser/actors/cache"
	"github.com/zondax/fil-parser/actors/cache/impl/common"
	"github.com/zondax/fil-parser/parser"
	helper2 "github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/fil-parser/types"
	"github.com/zondax/golem/pkg/metrics"
	"github.com/zondax/golem/pkg/zcache"

	"github.com/filecoin-project/lotus/api/client"
	rosettaFilecoinLib "github.com/zondax/rosetta-filecoin-lib"
)

const (
	testUrl  = "https://api.zondax.ch/fil/node/mainnet/rpc/v1"
	dataPath = "../data/actors"
)

func loadFile(actor, txType, name string) ([]byte, error) {
	item, err := os.ReadFile(fmt.Sprintf("%s/%s/%s/%s", dataPath, actor, txType, name))
	if err != nil {
		return nil, err
	}
	return item, nil
}

func getActorParser() *ActorParser {
	lotusClient, _, err := client.NewFullNodeRPCV1(context.Background(), testUrl, http.Header{})
	if err != nil {
		return nil
	}
	logger := logger2.NewLogger()

	actorsCache, err := cache.SetupActorsCache(common.DataSource{
		Node: lotusClient,
		Config: common.DataSourceConfig{
			Cache: &zcache.CombinedConfig{
				IsRemoteBestEffort: true,
				GlobalLogger:       logger,
				GlobalMetricServer: metrics.NewTaskMetrics("", "3000", "test"),
			},
		},
	}, logger)

	if err != nil {
		return nil
	}

	lib := rosettaFilecoinLib.NewRosettaConstructionFilecoin(lotusClient)
	helper := helper2.NewHelper(lib, actorsCache, lotusClient, logger)

	return NewActorParser(helper, logger)
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
