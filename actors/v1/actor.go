package actors

import (
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/actors"
	actormetrics "github.com/zondax/fil-parser/actors/metrics"
	logger2 "github.com/zondax/fil-parser/logger"
	"github.com/zondax/fil-parser/metrics"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/fil-parser/types"
	"go.uber.org/zap"
)

type ActorParser struct {
	helper  *helper.Helper
	logger  *zap.Logger
	metrics *actormetrics.ActorsMetricsClient
}

func NewActorParser(helper *helper.Helper, logger *zap.Logger, metrics metrics.MetricsClient) actors.ActorParserInterface {
	return &ActorParser{
		helper:  helper,
		logger:  logger2.GetSafeLogger(logger),
		metrics: actormetrics.NewClient(metrics),
	}
}

func (p *ActorParser) GetMetadata(txType string, msg *parser.LotusMessage, mainMsgCid cid.Cid, msgRct *parser.LotusMessageReceipt,
	height int64, key filTypes.TipSetKey) (string, map[string]interface{}, *types.AddressInfo, error) {
	metadata := make(map[string]interface{})
	if msg == nil {
		return "", metadata, nil, nil
	}

	actor, err := p.helper.GetActorNameFromAddress(msg.To, height, key)
	if err != nil {
		return "", metadata, nil, err
	}

	var addressInfo *types.AddressInfo
	switch actor {
	case manifest.InitKey:
		metadata, addressInfo, err = p.ParseInit(txType, msg, msgRct)
	case manifest.CronKey:
		metadata, err = p.ParseCron(txType, msg, msgRct)
	case manifest.AccountKey:
		metadata, err = p.ParseAccount(txType, msg, msgRct)
	case manifest.PowerKey:
		metadata, addressInfo, err = p.ParseStoragepower(txType, msg, msgRct)
	case manifest.MinerKey:
		metadata, err = p.ParseStorageminer(txType, msg, msgRct)
	case manifest.MarketKey:
		metadata, err = p.ParseStoragemarket(txType, msg, msgRct)
	case manifest.PaychKey:
		metadata, err = p.ParsePaymentchannel(txType, msg, msgRct)
	case manifest.MultisigKey:
		metadata, err = p.ParseMultisig(txType, msg, msgRct, height, key)
	case manifest.RewardKey:
		metadata, err = p.ParseReward(txType, msg, msgRct)
	case manifest.VerifregKey:
		metadata, err = p.ParseVerifiedRegistry(txType, msg, msgRct)
	case manifest.EvmKey:
		metadata, err = p.ParseEvm(txType, msg, msgRct)
	case manifest.EamKey:
		metadata, addressInfo, err = p.ParseEam(txType, msg, msgRct, mainMsgCid)
	case manifest.DatacapKey:
		metadata, err = p.ParseDatacap(txType, msg, msgRct)
	case manifest.EthAccountKey:
		metadata, err = p.ParseEthAccount(txType, msg, msgRct)
	case manifest.PlaceholderKey:
		metadata, err = p.ParsePlaceholder(txType, msg, msgRct)
	default:
		err = parser.ErrNotValidActor
	}

	return actor, metadata, addressInfo, err
}
