package parser

import (
	"encoding/json"

	"github.com/zondax/fil-parser/parser/V22"
	"strings"
	"time"

	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/database"
	"github.com/zondax/fil-parser/types"
	rosettaFilecoinLib "github.com/zondax/rosetta-filecoin-lib"
)

type IParser interface {
	ParseTransactions(traces any, tipSet *filTypes.TipSet, ethLogs []types.EthLog) ([]*types.Transaction, *types.AddressInfoMap, error)
}

type Parser struct {
	parserV22 IParser
	parserV23 IParser

	lib *rosettaFilecoinLib.RosettaConstructionFilecoin
}

func NewParser(lib *rosettaFilecoinLib.RosettaConstructionFilecoin) *Parser {
	return &Parser{
		parserV22: V22.NewParserV22(lib),
		parserV23: V23.NewParserV23(lib),
		lib:       lib,
	}
}

func (p *Parser) ParseTransactions(traces any, tipSet *filTypes.TipSet, ethLogs []types.EthLog) ([]*types.Transaction, *types.AddressInfoMap, error) {
	// Parse traces according to its inner version
	//p.parseTraceV23()
	//p.parsetraceV22()
}

func GetStatus(code string) string {
	status := strings.Split(code, "(")
	if len(status) == 2 {
		return status[0]
	}
	return code
}

func (p *Parser) GetMetadata(txType string, msg *LotusMessage, mainMsgCid cid.Cid, msgRct *filTypes.MessageReceipt,
	height int64, key filTypes.TipSetKey, ethLogs []types.EthLog) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	if msg == nil {
		return metadata, nil
	}
	var err error
	actorCode, err := database.ActorsDB.GetActorCode(msg.To, height, key)
	if err != nil {
		return metadata, err
	}
	actor, err := p.Lib.BuiltinActors.GetActorNameFromCid(actorCode)
	if err != nil {
		return metadata, err
	}
	switch actor {
	case manifest.InitKey:
		return p.parseInit(txType, msg, msgRct)
	case manifest.CronKey:
		return p.parseCron(txType, msg, msgRct)
	case manifest.AccountKey:
		return p.parseAccount(txType, msg, msgRct)
	case manifest.PowerKey:
		return p.parseStoragepower(txType, msg, msgRct)
	case manifest.MinerKey:
		return p.parseStorageminer(txType, msg, msgRct)
	case manifest.MarketKey:
		return p.parseStoragemarket(txType, msg, msgRct)
	case manifest.PaychKey:
		return p.parsePaymentchannel(txType, msg, msgRct)
	case manifest.MultisigKey:
		return p.parseMultisig(txType, msg, msgRct, height, key)
	case manifest.RewardKey:
		return p.parseReward(txType, msg, msgRct)
	case manifest.VerifregKey:
		return p.parseVerifiedRegistry(txType, msg, msgRct)
	case manifest.EvmKey, manifest.EthAccountKey:
		return p.parseEvm(txType, msg, mainMsgCid, msgRct, ethLogs)
	case manifest.EamKey:
		return p.parseEam(txType, msg, msgRct, mainMsgCid, ethLogs)
	case manifest.DatacapKey:
		return p.parseDatacap(txType, msg, msgRct)
	case manifest.PlaceholderKey:
		return metadata, nil // placeholder has no methods
	default:
		return metadata, ErrNotValidActor
	}
}

func parseParams(metadata map[string]interface{}) string {
	params, ok := metadata[ParamsKey].(string)
	if ok && params != "" {
		return params
	}
	jsonMetadata, err := json.Marshal(metadata[ParamsKey])
	if err == nil && string(jsonMetadata) != "null" {
		return string(jsonMetadata)
	}
	return ""
}

func parseReturn(metadata map[string]interface{}) string {
	params, ok := metadata[ReturnKey].(string)
	if ok && params != "" {
		return params
	}
	jsonMetadata, err := json.Marshal(metadata[ReturnKey])
	if err == nil && string(jsonMetadata) != "null" {
		return string(jsonMetadata)
	}
	return ""
}

func (p *Parser) appendAddressInfo(msg *filTypes.Message, height int64, key filTypes.TipSetKey) {
	if msg == nil {
		return
	}
	fromAdd := p.getActorAddressInfo(msg.From, height, key)
	toAdd := p.getActorAddressInfo(msg.To, height, key)
	p.appendToAddresses(fromAdd, toAdd)
}

func (p *Parser) appendToAddresses(info ...types.AddressInfo) {
	if p.Addresses == nil {
		return
	}
	for _, i := range info {
		if i.Robust != "" && i.Short != "" && i.Robust != i.Short {
			if _, ok := p.Addresses[i.Short]; !ok {
				p.Addresses[i.Short] = i
			}
		}
	}
}

func (p *Parser) getTimestamp(timestamp uint64) time.Time {
	blockTimeStamp := int64(timestamp) * 1000
	return time.Unix(blockTimeStamp/1000, blockTimeStamp%1000)
}
