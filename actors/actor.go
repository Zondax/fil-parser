package actors

import (
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/fil-parser/types"
	rosettaFilecoinLib "github.com/zondax/rosetta-filecoin-lib"
	"go.uber.org/zap"
)

type ActorParser struct {
	lib    *rosettaFilecoinLib.RosettaConstructionFilecoin
	helper *helper.Helper
}

func NewActorParser(lib *rosettaFilecoinLib.RosettaConstructionFilecoin, helper *helper.Helper) *ActorParser {
	return &ActorParser{
		lib:    lib,
		helper: helper,
	}
}

func (p *ActorParser) GetMetadata(txType string, msg *parser.LotusMessage, mainMsgCid cid.Cid, msgRct *parser.LotusMessageReceipt,
	height int64, key filTypes.TipSetKey, ethLogs []types.EthLog) (map[string]interface{}, *types.AddressInfo, error) {
	metadata := make(map[string]interface{})
	if msg == nil {
		return metadata, nil, nil
	}

	actorCode, err := p.helper.GetActorsCache().GetActorCode(msg.To, key)
	if err != nil {
		return metadata, nil, err
	}

	c, err := cid.Parse(actorCode)
	if err != nil {
		zap.S().Errorf("Could not parse actor code: %v", err)
		return metadata, nil, err
	}

	actor, err := p.lib.BuiltinActors.GetActorNameFromCid(c)
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
	case manifest.EvmKey, manifest.EthAccountKey:
		metadata, err = p.ParseEvm(txType, msg, mainMsgCid, msgRct, ethLogs)
	case manifest.EamKey:
		metadata, addressInfo, err = p.ParseEam(txType, msg, msgRct, mainMsgCid)
	case manifest.DatacapKey:
		metadata, err = p.ParseDatacap(txType, msg, msgRct)
	case manifest.PlaceholderKey:
		err = nil // placeholder has no methods
	default:
		err = parser.ErrNotValidActor
	}
	return metadata, addressInfo, err
}
