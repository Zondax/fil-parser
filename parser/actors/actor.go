package actors

import (
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/database"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
	rosettaFilecoinLib "github.com/zondax/rosetta-filecoin-lib"
)

type ActorParser struct {
	lib *rosettaFilecoinLib.RosettaConstructionFilecoin
}

func NewActorParser(lib *rosettaFilecoinLib.RosettaConstructionFilecoin) *ActorParser {
	return &ActorParser{
		lib: lib,
	}
}

func (p *ActorParser) GetMetadata(txType string, msg *parser.LotusMessage, mainMsgCid cid.Cid, msgRct *parser.LotusMessageReceipt,
	height int64, key filTypes.TipSetKey, ethLogs []types.EthLog) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	if msg == nil {
		return metadata, nil
	}

	actorCode, err := database.ActorsDB.GetActorCode(msg.To, height, key)
	if err != nil {
		return metadata, err
	}

	actor, err := p.lib.BuiltinActors.GetActorNameFromCid(actorCode)
	if err != nil {
		return metadata, err
	}

	switch actor {
	case manifest.InitKey:
		return ParseInit(txType, msg, msgRct)
	case manifest.CronKey:
		return ParseCron(txType, msg, msgRct)
	case manifest.AccountKey:
		return ParseAccount(txType, msg, msgRct)
	case manifest.PowerKey:
		return ParseStoragepower(txType, msg, msgRct)
	case manifest.MinerKey:
		return ParseStorageminer(txType, msg, msgRct)
	case manifest.MarketKey:
		return ParseStoragemarket(txType, msg, msgRct)
	case manifest.PaychKey:
		return ParsePaymentchannel(txType, msg, msgRct)
	case manifest.MultisigKey:
		return ParseMultisig(txType, msg, msgRct, height, key)
	case manifest.RewardKey:
		return ParseReward(txType, msg, msgRct)
	case manifest.VerifregKey:
		return ParseVerifiedRegistry(txType, msg, msgRct)
	case manifest.EvmKey, manifest.EthAccountKey:
		return ParseEvm(txType, msg, mainMsgCid, msgRct, ethLogs)
	case manifest.EamKey:
		return ParseEam(txType, msg, msgRct, mainMsgCid, ethLogs)
	case manifest.DatacapKey:
		return ParseDatacap(txType, msg, msgRct)
	case manifest.PlaceholderKey:
		return metadata, nil // placeholder has no methods
	default:
		return metadata, parser.ErrNotValidActor
	}
}
