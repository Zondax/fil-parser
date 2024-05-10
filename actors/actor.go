package actors

import (
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/fil-parser/types"
	"github.com/zondax/golem/pkg/logger"
)

type ActorParser struct {
	helper *helper.Helper
	logger *logger.Logger
}

func NewActorParser(helper *helper.Helper, logger *logger.Logger) *ActorParser {
	return &ActorParser{
		helper: helper,
		logger: logger,
	}
}

func (p *ActorParser) GetMetadata(txType string, msg *parser.LotusMessage, mainMsgCid cid.Cid, msgRct *parser.LotusMessageReceipt,
	height int64, key filTypes.TipSetKey) (map[string]interface{}, *types.AddressInfo, error) {
	metadata := make(map[string]interface{})
	if msg == nil {
		return metadata, nil, nil
	}

	actor, err := p.helper.GetActorNameFromAddress(msg.To, height, key)
	if err != nil {
		return metadata, nil, err
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
	return metadata, addressInfo, err
}
