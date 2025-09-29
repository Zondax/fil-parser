package actortest

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"

	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/mock"
	"github.com/zondax/fil-parser/metrics"
	"github.com/zondax/fil-parser/tools/mocks"

	"github.com/filecoin-project/go-state-types/manifest"
	filApiTypes "github.com/filecoin-project/lotus/api/types"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/node/modules/dtypes"
	"github.com/zondax/fil-parser/actors"
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
	actorCid := cid.MustParse("bafk2bzaceaclpbrhoqdruvsuqqgknvy2k5dywzmjoehk4uarce3uvt3w2rewu")
	lotusClient := &mocks.FullNode{}
	lotusClient.On("StateNetworkName", mock.Anything).Return(dtypes.NetworkName("calibrationnet"), nil)
	lotusClient.On("StateNetworkVersion", mock.Anything, mock.Anything).Return(filApiTypes.NetworkVersion(16), nil)
	lotusClient.On("StateActorCodeCIDs", mock.Anything, mock.Anything).Return(map[string]cid.Cid{
		manifest.MultisigKey: actorCid,
	}, nil)

	cache := &mocks.IActorsCache{}
	cache.On("StoreAddressInfo", mock.Anything).Return(nil)
	cache.On("GetActorCode", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(actorCid.String(), nil)
	cache.On("GetActorNameFromAddress", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(manifest.MultisigKey, nil)

	lib := rosettaFilecoinLib.NewRosettaConstructionFilecoin(lotusClient)
	helper := helper2.NewHelper(lib, cache, lotusClient, nil, metrics.NewNoopMetricsClient())
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
