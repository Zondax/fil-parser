package parser

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/filecoin-project/lotus/api"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	rosettaFilecoinLib "github.com/zondax/rosetta-filecoin-lib"
	"go.uber.org/zap"

	"github.com/zondax/fil-parser/database"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
)

type Parser struct {
	lib       *rosettaFilecoinLib.RosettaConstructionFilecoin
	addresses types.AddressInfoMap
}

func NewParser(lib *rosettaFilecoinLib.RosettaConstructionFilecoin) *Parser {
	return &Parser{
		lib:       lib,
		addresses: types.NewAddressInfoMap(),
	}
}

func (p *Parser) ParseTransactions(traces []*api.InvocResult, tipSet *filTypes.TipSet, ethLogs []types.EthLog) ([]*types.Transaction, *types.AddressInfoMap, error) {
	var transactions []*types.Transaction
	p.addresses = types.NewAddressInfoMap()
	tipsetKey := tipSet.Key()
	blockHash, err := tools.BuildTipSetKeyHash(tipsetKey)
	if err != nil {
		return nil, nil, errBlockHash
	}
	for _, trace := range traces {
		if !hasMessage(trace) {
			continue
		}

		// Main transaction
		transaction, err := p.parseTrace(trace.ExecutionTrace, trace.MsgCid, tipSet, ethLogs, *blockHash, tipsetKey)
		if err != nil {
			continue
		}
		transactions = append(transactions, transaction)

		// Subcalls
		subTxs := p.parseSubTxs(trace.ExecutionTrace.Subcalls, trace.MsgCid, tipSet, ethLogs, *blockHash,
			trace.Msg.Cid().String(), tipsetKey, 0)
		if len(subTxs) > 0 {
			transactions = append(transactions, subTxs...)
		}

		// Fees
		if trace.GasCost.TotalCost.Uint64() > 0 {
			feeTx := p.feesTransactions(trace, tipSet.Blocks()[0].Miner.String(), transaction.TxHash, *blockHash,
				transaction.TxType, uint64(tipSet.Height()), tipSet.MinTimestamp())
			transactions = append(transactions, feeTx)
		}
	}

	return transactions, &p.addresses, nil
}

func (p *Parser) parseSubTxs(subTxs []filTypes.ExecutionTrace, mainMsgCid cid.Cid, tipSet *filTypes.TipSet, ethLogs []types.EthLog, blockHash, txHash string,
	key filTypes.TipSetKey, level uint16) (txs []*types.Transaction) {

	level++
	for _, subTx := range subTxs {
		subTransaction, err := p.parseTrace(subTx, mainMsgCid, tipSet, ethLogs, blockHash, key)
		if err != nil {
			continue
		}
		subTransaction.Level = level
		txs = append(txs, subTransaction)
		txs = append(txs, p.parseSubTxs(subTx.Subcalls, mainMsgCid, tipSet, ethLogs, blockHash, txHash, key, level)...)
	}
	return
}

func (p *Parser) parseTrace(trace filTypes.ExecutionTrace, msgCid cid.Cid, tipSet *filTypes.TipSet, ethLogs []types.EthLog, blockHash string,
	key filTypes.TipSetKey) (*types.Transaction, error) {
	txType, err := p.GetMethodName(trace.Msg, int64(tipSet.Height()), key)
	if err != nil {
		zap.S().Errorf("Error when trying to get method name in tx cid'%s': %v", msgCid.String(), err)
		txType = UnknownStr
	}
	if err == nil && txType == UnknownStr {
		zap.S().Errorf("Could not get method name in transaction '%s'", msgCid.String())
	}

	metadata, mErr := p.getMetadata(txType, trace.Msg, msgCid, trace.MsgRct, int64(tipSet.Height()), key, ethLogs)
	if mErr != nil {
		zap.S().Warnf("Could not get metadata for transaction in height %s of type '%s': %s", tipSet.Height().String(), txType, mErr.Error())
	}
	if trace.Error != "" {
		metadata["Error"] = trace.Error
	}
	params := parseParams(metadata)
	jsonMetadata, _ := json.Marshal(metadata)
	txReturn := parseReturn(metadata)

	p.appendAddressInfo(trace.Msg, int64(tipSet.Height()), key)

	return &types.Transaction{
		BasicBlockData: types.BasicBlockData{
			Height: uint64(tipSet.Height()),
			Hash:   blockHash,
		},
		Level:       0,
		TxTimestamp: p.getTimestamp(tipSet.MinTimestamp()),
		TxHash:      msgCid.String(),
		TxFrom:      trace.Msg.From.String(),
		TxTo:        trace.Msg.To.String(),
		Amount:      trace.Msg.Value.Int,
		GasUsed:     trace.MsgRct.GasUsed,
		Status:      getStatus(trace.MsgRct.ExitCode.String()),
		TxType:      txType,
		TxMetadata:  string(jsonMetadata),
		TxParams:    fmt.Sprintf("%v", params),
		TxReturn:    txReturn,
	}, nil
}

func (p *Parser) feesTransactions(msg *api.InvocResult, minerAddress, txHash, blockHash, txType string, height uint64, timestamp uint64) *types.Transaction {
	ts := p.getTimestamp(timestamp)
	metadata := FeesMetadata{
		TxType: txType,
		MinerFee: MinerFee{
			MinerAddress: minerAddress,
			Amount:       msg.GasCost.MinerTip.String(),
		},
		OverEstimationBurnFee: OverEstimationBurnFee{
			BurnAddress: BurnAddress,
			Amount:      msg.GasCost.OverEstimationBurn.String(),
		},
		BurnFee: BurnFee{
			BurnAddress: BurnAddress,
			Amount:      msg.GasCost.BaseFeeBurn.String(),
		},
	}

	feeTx := p.newFeeTx(msg, txHash, blockHash, height, ts, metadata)
	return feeTx
}

func (p *Parser) newFeeTx(msg *api.InvocResult, txHash, blockHash string, height uint64,
	timestamp time.Time, feesMetadata FeesMetadata) *types.Transaction {
	metadata, _ := json.Marshal(feesMetadata)

	return &types.Transaction{
		BasicBlockData: types.BasicBlockData{
			Height: height,
			Hash:   blockHash,
		},
		TxTimestamp: timestamp,
		TxHash:      txHash,
		TxFrom:      msg.Msg.From.String(),
		Amount:      msg.GasCost.TotalCost.Int,
		Status:      "Ok",
		TxType:      TotalFeeOp,
		TxMetadata:  string(metadata),
	}

}

func hasMessage(trace *api.InvocResult) bool {
	return trace.Msg != nil
}

func getStatus(code string) string {
	status := strings.Split(code, "(")
	if len(status) == 2 {
		return status[0]
	}
	return code
}

func (p *Parser) getMetadata(txType string, msg *filTypes.Message, mainMsgCid cid.Cid, msgRct *filTypes.MessageReceipt,
	height int64, key filTypes.TipSetKey, ethLogs []types.EthLog) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	var err error
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
		return metadata, errNotValidActor
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

//  func getCastedAmount(amount string) decimal.Decimal {
//	  decimal.DivisionPrecision = 18
//	  parsed, err := decimal.NewFromString(amount)
//	  if err != nil {
//		  return decimal.Decimal{}
//	  }
//	  abs := parsed.Abs()
//	  divided := abs.DivRound(decimal.NewFromInt(1e+18), 18)
//
//	  return divided
//  }

func (p *Parser) appendAddressInfo(msg *filTypes.Message, height int64, key filTypes.TipSetKey) {
	fromAdd := p.getActorAddressInfo(msg.From, height, key)
	toAdd := p.getActorAddressInfo(msg.To, height, key)
	p.appendToAddresses(fromAdd, toAdd)
}

func (p *Parser) appendToAddresses(info ...types.AddressInfo) {
	if p.addresses == nil {
		return
	}
	for _, i := range info {
		if i.Robust != "" && i.Short != "" && i.Robust != i.Short {
			if _, ok := p.addresses[i.Short]; !ok {
				p.addresses[i.Short] = i
			}
		}
	}
}

func (p *Parser) getTimestamp(timestamp uint64) time.Time {
	blockTimeStamp := int64(timestamp) * 1000
	return time.Unix(blockTimeStamp/1000, blockTimeStamp%1000)
}
