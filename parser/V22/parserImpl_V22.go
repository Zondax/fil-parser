package V22

import (
	"encoding/json"
	"fmt"
	"github.com/zondax/fil-parser/parser"
	typesv22 "github.com/zondax/fil-parser/parser/V22/types"
	"github.com/zondax/fil-parser/parser/actors"
	"github.com/zondax/fil-parser/parser/helper"
	rosettaFilecoinLib "github.com/zondax/rosetta-filecoin-lib"
	"strings"
	"time"

	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"go.uber.org/zap"

	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
)

type Parser struct {
	actorParser *actors.ActorParser
	addresses   types.AddressInfoMap
	helper      *helper.Helper
}

func NewParserV22(lib *rosettaFilecoinLib.RosettaConstructionFilecoin) parser.IParser {
	return &Parser{
		actorParser: actors.NewActorParser(lib),
		addresses:   types.NewAddressInfoMap(),
		helper:      helper.NewHelper(lib),
	}
}

func (p *Parser) ParseTransactions(traces interface{}, tipSet *filTypes.TipSet, ethLogs []types.EthLog) ([]*types.Transaction, *types.AddressInfoMap, error) {
	// cast to correct type
	tracesV22, ok := traces.([]*typesv22.InvocResultV22)
	if !ok {
		return nil, nil, parser.ErrInvalidType
	}

	var transactions []*types.Transaction
	p.addresses = types.NewAddressInfoMap()
	tipsetKey := tipSet.Key()
	blockHash, err := tools.BuildTipSetKeyHash(tipsetKey)
	if err != nil {
		return nil, nil, parser.ErrBlockHash
	}

	for _, trace := range tracesV22 {
		if !hasMessage(trace) {
			continue
		}

		if ok, badTx := hasExecutionTrace(trace); !ok {
			// Add missing fields and continue
			badTx.BasicBlockData = types.BasicBlockData{
				Height: uint64(tipSet.Height()),
				Hash:   *blockHash,
			}
			badTx.TxTimestamp = p.getTimestamp(tipSet.MinTimestamp())

			txType, err := p.helper.GetMethodName(&parser.LotusMessage{
				To:     trace.Msg.To,
				From:   trace.Msg.From,
				Method: trace.Msg.Method,
			}, int64(tipSet.Height()), tipsetKey)

			if err != nil {
				txType = parser.UnknownStr
			}
			badTx.TxType = txType
			continue
		}

		// Main transaction
		transaction, err := p.parseTrace(trace.ExecutionTrace, trace.MsgCid, tipSet, ethLogs, *blockHash, tipsetKey)
		if err != nil {
			continue
		}
		transactions = append(transactions, transaction)

		// Only process sub-calls if the parent call was successfully executed
		if trace.ExecutionTrace.MsgRct.ExitCode.IsSuccess() {
			subTxs := p.parseSubTxs(trace.ExecutionTrace.Subcalls, trace.MsgCid, tipSet, ethLogs, *blockHash,
				trace.Msg.Cid().String(), tipsetKey, 0)
			if len(subTxs) > 0 {
				transactions = append(transactions, subTxs...)
			}
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

func (p *Parser) parseSubTxs(subTxs []typesv22.ExecutionTraceV22, mainMsgCid cid.Cid, tipSet *filTypes.TipSet, ethLogs []types.EthLog, blockHash, txHash string,
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

func (p *Parser) parseTrace(trace typesv22.ExecutionTraceV22, msgCid cid.Cid, tipSet *filTypes.TipSet, ethLogs []types.EthLog, blockHash string,
	key filTypes.TipSetKey) (*types.Transaction, error) {
	txType, err := p.helper.GetMethodName(&parser.LotusMessage{
		To:     trace.Msg.To,
		From:   trace.Msg.From,
		Method: trace.Msg.Method,
	}, int64(tipSet.Height()), key)

	if err != nil {
		zap.S().Errorf("Error when trying to get method name in tx cid'%s': %v", msgCid.String(), err)
		txType = parser.UnknownStr
	}
	if err == nil && txType == parser.UnknownStr {
		zap.S().Errorf("Could not get method name in transaction '%s'", msgCid.String())
	}

	metadata, mErr := p.actorParser.GetMetadata(txType, &parser.LotusMessage{
		To:     trace.Msg.To,
		From:   trace.Msg.From,
		Method: trace.Msg.Method,
		Cid:    trace.Msg.Cid(),
		Params: trace.Msg.Params,
	}, msgCid, &parser.LotusMessageReceipt{
		ExitCode: trace.MsgRct.ExitCode,
		Return:   trace.MsgRct.Return,
	}, int64(tipSet.Height()), key, ethLogs)

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

func (p *Parser) feesTransactions(msg *typesv22.InvocResultV22, minerAddress, txHash, blockHash, txType string, height uint64, timestamp uint64) *types.Transaction {
	ts := p.getTimestamp(timestamp)
	metadata := parser.FeesMetadata{
		TxType: txType,
		MinerFee: parser.MinerFee{
			MinerAddress: minerAddress,
			Amount:       msg.GasCost.MinerTip.String(),
		},
		OverEstimationBurnFee: parser.OverEstimationBurnFee{
			BurnAddress: parser.BurnAddress,
			Amount:      msg.GasCost.OverEstimationBurn.String(),
		},
		BurnFee: parser.BurnFee{
			BurnAddress: parser.BurnAddress,
			Amount:      msg.GasCost.BaseFeeBurn.String(),
		},
	}

	feeTx := p.newFeeTx(msg, txHash, blockHash, height, ts, metadata)
	return feeTx
}

func (p *Parser) newFeeTx(msg *typesv22.InvocResultV22, txHash, blockHash string, height uint64,
	timestamp time.Time, feesMetadata parser.FeesMetadata) *types.Transaction {
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
		TxType:      parser.TotalFeeOp,
		TxMetadata:  string(metadata),
	}

}

func hasMessage(trace *typesv22.InvocResultV22) bool {
	return trace.Msg != nil
}

func hasExecutionTrace(trace *typesv22.InvocResultV22) (bool, *types.Transaction) {
	// check if this execution trace is valid
	if trace.ExecutionTrace.Msg == nil || trace.ExecutionTrace.MsgRct == nil {

		// this is an invalid message
		return false, &types.Transaction{
			Level:      0,
			TxHash:     trace.MsgCid.String(),
			TxFrom:     trace.Msg.From.String(),
			TxTo:       trace.Msg.To.String(),
			Amount:     trace.Msg.Value.Int,
			GasUsed:    trace.MsgRct.GasUsed,
			Status:     getStatus(trace.MsgRct.ExitCode.String()),
			TxType:     parser.UnknownStr,
			TxMetadata: trace.Error,
		}
	}
	return true, nil
}

func getStatus(code string) string {
	status := strings.Split(code, "(")
	if len(status) == 2 {
		return status[0]
	}
	return code
}

func parseParams(metadata map[string]interface{}) string {
	params, ok := metadata[parser.ParamsKey].(string)
	if ok && params != "" {
		return params
	}
	jsonMetadata, err := json.Marshal(metadata[parser.ParamsKey])
	if err == nil && string(jsonMetadata) != "null" {
		return string(jsonMetadata)
	}
	return ""
}

func parseReturn(metadata map[string]interface{}) string {
	params, ok := metadata[parser.ReturnKey].(string)
	if ok && params != "" {
		return params
	}
	jsonMetadata, err := json.Marshal(metadata[parser.ReturnKey])
	if err == nil && string(jsonMetadata) != "null" {
		return string(jsonMetadata)
	}
	return ""
}

func (p *Parser) appendAddressInfo(msg *filTypes.Message, height int64, key filTypes.TipSetKey) {
	if msg == nil {
		return
	}
	fromAdd := p.helper.GetActorAddressInfo(msg.From, height, key)
	toAdd := p.helper.GetActorAddressInfo(msg.To, height, key)
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
