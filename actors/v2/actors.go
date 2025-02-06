package v2

import (
	"fmt"

	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/actors/v2/account"
	"github.com/zondax/fil-parser/actors/v2/cron"
	"github.com/zondax/fil-parser/actors/v2/datacap"
	"github.com/zondax/fil-parser/actors/v2/eam"
	"github.com/zondax/fil-parser/actors/v2/ethaccount"
	"github.com/zondax/fil-parser/actors/v2/evm"
	initActor "github.com/zondax/fil-parser/actors/v2/init"
	"github.com/zondax/fil-parser/actors/v2/market"
	"github.com/zondax/fil-parser/actors/v2/miner"
	"github.com/zondax/fil-parser/actors/v2/multisig"
	paymentchannel "github.com/zondax/fil-parser/actors/v2/paymentChannel"
	"github.com/zondax/fil-parser/actors/v2/placeholder"
	"github.com/zondax/fil-parser/actors/v2/power"
	"github.com/zondax/fil-parser/actors/v2/reward"
	"github.com/zondax/fil-parser/actors/v2/system"
	verifiedregistry "github.com/zondax/fil-parser/actors/v2/verifiedRegistry"
	logger2 "github.com/zondax/fil-parser/logger"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/fil-parser/types"
	"go.uber.org/zap"
)

type Actor interface {
	Name() string
	Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, mainMsgCid cid.Cid, key filTypes.TipSetKey) (map[string]interface{}, *types.AddressInfo, error)
	TransactionTypes() map[string]any
}

type ActorParser struct {
	helper *helper.Helper
	logger *zap.Logger
}

func NewActorParser(helper *helper.Helper, logger *zap.Logger) actors.ActorParserInterface {
	return &ActorParser{
		helper: helper,
		logger: logger2.GetSafeLogger(logger),
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
	network := ""

	actorParser, err := p.GetActor(actor)
	if err != nil {
		return nil, nil, parser.ErrNotValidActor
	}
	return actorParser.Parse(network, height, txType, msg, msgRct, mainMsgCid, key)
}

func (p *ActorParser) LatestSupportedVersion(actor string) (uint64, error) {
	keys := manifest.GetBuiltinActorsKeys(10)

	for _, key := range keys {
		if key == actor {
			return 10, nil
		}
	}
	return 0, nil
}

func (p *ActorParser) GetActor(actor string) (Actor, error) {
	switch actor {
	case manifest.AccountKey:
		return account.New(p.logger), nil
	case manifest.CronKey:
		return cron.New(p.logger), nil
	case manifest.DatacapKey:
		return datacap.New(p.logger), nil
	case manifest.EamKey:
		return eam.New(p.logger), nil
	case manifest.EthAccountKey:
		return ethaccount.New(p.logger), nil
	case manifest.EvmKey:
		return evm.New(p.logger), nil
	case manifest.InitKey:
		return initActor.New(p.logger), nil
	case manifest.MarketKey:
		return market.New(p.logger), nil
	case manifest.MinerKey:
		return miner.New(p.logger), nil
	case manifest.MultisigKey:
		return multisig.New(p.helper, p.logger), nil
	case manifest.PaychKey:
		return paymentchannel.New(p.logger), nil
	case manifest.PowerKey:
		return power.New(p.logger), nil
	case manifest.RewardKey:
		return reward.New(p.logger), nil
	case manifest.VerifregKey:
		return verifiedregistry.New(p.logger), nil
	case manifest.PlaceholderKey:
		return placeholder.New(p.logger), nil
	case manifest.SystemKey:
		return system.New(p.logger), nil
	default:
		return nil, fmt.Errorf("actor %s not found", actor)
	}
}
