package V22

import (
	"encoding/json"
	"errors"
	"github.com/zondax/fil-parser/actors"
	"strings"
	"time"

	"github.com/bytedance/sonic"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	rosettaFilecoinLib "github.com/zondax/rosetta-filecoin-lib"
	"go.uber.org/zap"

	"github.com/zondax/fil-parser/parser"
	typesv22 "github.com/zondax/fil-parser/parser/V22/types"
	"github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
)

const (
	Version = "v22"
)

type Parser struct {
	actorParser *actors.ActorParser
	addresses   types.AddressInfoMap
	helper      *helper.Helper
}

func NewParserV22(lib *rosettaFilecoinLib.RosettaConstructionFilecoin, helper *helper.Helper) *Parser {
	return &Parser{
		actorParser: actors.NewActorParser(lib, helper),
		addresses:   types.NewAddressInfoMap(),
		helper:      helper,
	}
}

func (p *Parser) Version() string {
	return Version
}

func (p *Parser) ParseTransactions(traces []byte, tipSet *types.ExtendedTipSet, ethLogs []types.EthLog) ([]*types.Transaction, types.AddressInfoMap, error) {
	// Unmarshal into vComputeState
	computeState := &typesv22.ComputeStateOutputV22{}
	err := sonic.UnmarshalString(string(traces), &computeState)
	if err != nil {
		zap.S().Error(err)
		return nil, nil, errors.New("could not decode")
	}

	var transactions []*types.Transaction
	p.addresses = types.NewAddressInfoMap()
	tipsetKey := tipSet.Key()
	tipsetHash, err := tools.BuildTipSetKeyHash(tipsetKey)
	if err != nil {
		return nil, nil, parser.ErrBlockHash
	}

	for _, trace := range computeState.Trace {
		if !hasMessage(trace) {
			continue
		}

		// TODO find a way to not having this special case handled outside func parseTrace
		if ok := hasExecutionTrace(trace); !ok {
			// Create tx
			txType, _ := p.helper.GetMethodName(&parser.LotusMessage{
				To:     trace.Msg.To,
				From:   trace.Msg.From,
				Method: trace.Msg.Method,
			}, int64(tipSet.Height()), tipsetKey)

			blockCid, err := tools.GetBlockCidFromMsgCid(trace.MsgCid.String(), txType, nil, tipSet)
			if err != nil {
				zap.S().Errorf("Error when trying to get block cid from message,txType '%s': %v", txType, err)
			}
			messageUuid := tools.BuildMessageHash(*tipsetHash, blockCid, trace.Msg.Cid().String())

			badTx := &types.Transaction{
				BasicBlockData: types.BasicBlockData{
					Height:    uint64(tipSet.Height()),
					TipsetCid: *tipsetHash,
					BlockCid:  blockCid,
					Canonical: false,
				},
				Id:          messageUuid,
				TxHash:      trace.MsgCid.String(),
				TxFrom:      trace.Msg.From.String(),
				TxTo:        trace.Msg.To.String(),
				TxType:      txType,
				Amount:      trace.Msg.Value.Int,
				GasUsed:     trace.MsgRct.GasUsed,
				Status:      getStatus(trace.MsgRct.ExitCode.String()),
				TxMetadata:  trace.Error,
				TxTimestamp: p.getTimestamp(tipSet.MinTimestamp()),
			}

			transactions = append(transactions, badTx)
			continue
		}

		// Main transaction
		transaction, err := p.parseTrace(trace.ExecutionTrace, trace.MsgCid, tipSet, ethLogs, *tipsetHash, tipsetKey)
		if err != nil {
			continue
		}
		transaction.GasUsed = trace.GasCost.GasUsed.Int64()
		transactions = append(transactions, transaction)

		// Only process sub-calls if the parent call was successfully executed
		if trace.ExecutionTrace.MsgRct.ExitCode.IsSuccess() {
			subTxs := p.parseSubTxs(trace.ExecutionTrace.Subcalls, trace.MsgCid, tipSet, ethLogs, *tipsetHash,
				trace.Msg.Cid().String(), tipsetKey)
			if len(subTxs) > 0 {
				transactions = append(transactions, subTxs...)
			}
		}

		// Fees
		if trace.GasCost.TotalCost.Uint64() > 0 {
			feeTx := p.feesTransactions(trace, tipSet.Blocks()[0].Miner.String(), transaction.TxHash, *tipsetHash,
				transaction.TxType, uint64(tipSet.Height()), tipSet.MinTimestamp())
			transactions = append(transactions, feeTx)
		}
	}

	return transactions, p.addresses, nil
}

func (p *Parser) parseSubTxs(subTxs []typesv22.ExecutionTraceV22, mainMsgCid cid.Cid, tipSet *types.ExtendedTipSet, ethLogs []types.EthLog, tipsetHash, txHash string,
	key filTypes.TipSetKey) (txs []*types.Transaction) {

	for _, subTx := range subTxs {
		subTransaction, err := p.parseTrace(subTx, mainMsgCid, tipSet, ethLogs, tipsetHash, key)
		if err != nil {
			continue
		}

		txs = append(txs, subTransaction)
		txs = append(txs, p.parseSubTxs(subTx.Subcalls, mainMsgCid, tipSet, ethLogs, tipsetHash, txHash, key)...)
	}
	return
}

func (p *Parser) parseTrace(trace typesv22.ExecutionTraceV22, mainMsgCid cid.Cid, tipSet *types.ExtendedTipSet, ethLogs []types.EthLog, tipsetHash string,
	key filTypes.TipSetKey) (*types.Transaction, error) {

	txType, err := p.helper.GetMethodName(&parser.LotusMessage{
		To:     trace.Msg.To,
		From:   trace.Msg.From,
		Method: trace.Msg.Method,
	}, int64(tipSet.Height()), key)

	if err != nil {
		zap.S().Errorf("Error when trying to get method name in tx cid'%s': %v", mainMsgCid.String(), err)
		txType = parser.UnknownStr
	}
	if err == nil && txType == parser.UnknownStr {
		zap.S().Errorf("Could not get method name in transaction '%s'", trace.Msg.Cid().String())
	}

	metadata, addressInfo, mErr := p.actorParser.GetMetadata(txType, &parser.LotusMessage{
		To:     trace.Msg.To,
		From:   trace.Msg.From,
		Method: trace.Msg.Method,
		Cid:    trace.Msg.Cid(),
		Params: trace.Msg.Params,
	}, mainMsgCid, &parser.LotusMessageReceipt{
		ExitCode: trace.MsgRct.ExitCode,
		Return:   trace.MsgRct.Return,
	}, int64(tipSet.Height()), key, ethLogs)

	if mErr != nil {
		zap.S().Warnf("Could not get metadata for transaction in height %s of type '%s': %s", tipSet.Height().String(), txType, mErr.Error())
	}
	if addressInfo != nil {
		p.appendToAddresses(addressInfo)
	}
	if trace.MsgRct.ExitCode.IsError() {
		metadata["Error"] = trace.MsgRct.ExitCode.Error()
	}

	jsonMetadata, _ := json.Marshal(metadata)

	p.appendAddressInfo(trace.Msg, int64(tipSet.Height()), key)

	blockCid, err := tools.GetBlockCidFromMsgCid(mainMsgCid.String(), txType, metadata, tipSet)
	if err != nil {
		zap.S().Errorf("Error when trying to get block cid from message, txType '%s': %v", txType, err)
	}

	messageUuid := tools.BuildMessageHash(tipsetHash, blockCid, trace.Msg.Cid().String())

	return &types.Transaction{
		BasicBlockData: types.BasicBlockData{
			Height:    uint64(tipSet.Height()),
			TipsetCid: tipsetHash,
			BlockCid:  blockCid,
			Canonical: false,
		},
		Id:          messageUuid,
		TxTimestamp: p.getTimestamp(tipSet.MinTimestamp()),
		TxHash:      mainMsgCid.String(),
		TxFrom:      trace.Msg.From.String(),
		TxTo:        trace.Msg.To.String(),
		Amount:      trace.Msg.Value.Int,
		Status:      getStatus(trace.MsgRct.ExitCode.String()),
		TxType:      txType,
		TxMetadata:  string(jsonMetadata),
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
			Height:    height,
			TipsetCid: blockHash,
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

func hasExecutionTrace(trace *typesv22.InvocResultV22) bool {
	// check if this execution trace is valid
	if trace.ExecutionTrace.Msg == nil || trace.ExecutionTrace.MsgRct == nil {
		// this is an invalid message
		return false
	}
	return true
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

func (p *Parser) appendToAddresses(info ...*types.AddressInfo) {
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
