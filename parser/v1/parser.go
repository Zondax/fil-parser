package v1

import (
	"encoding/json"
	"errors"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/api"
	"math/big"
	"strings"

	"github.com/bytedance/sonic"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/google/uuid"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/actors"
	logger2 "github.com/zondax/fil-parser/logger"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/parser/helper"
	typesV1 "github.com/zondax/fil-parser/parser/v1/types"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
	"go.uber.org/zap"
)

const Version = "v1"

var NodeVersionsSupported = []string{"v1.21", "v1.22"}

type Parser struct {
	actorParser *actors.ActorParser
	addresses   *types.AddressInfoMap
	helper      *helper.Helper
	logger      *zap.Logger
}

func NewParser(helper *helper.Helper, logger *zap.Logger) *Parser {
	return &Parser{
		actorParser: actors.NewActorParser(helper, logger),
		addresses:   types.NewAddressInfoMap(),
		helper:      helper,
		logger:      logger2.GetSafeLogger(logger),
	}
}

func (p *Parser) Version() string {
	return Version
}

func (p *Parser) NodeVersionsSupported() []string {
	return NodeVersionsSupported
}

func (p *Parser) IsNodeVersionSupported(ver string) bool {
	for _, i := range NodeVersionsSupported {
		if strings.EqualFold(i, ver) {
			return true
		}
	}

	return false
}

func (p *Parser) ParseTransactions(traces []byte, tipset *types.ExtendedTipSet, ethLogs []types.EthLog, metadata types.BlockMetadata) ([]*types.Transaction, *types.AddressInfoMap, error) {
	// Unmarshal into vComputeState
	computeState := &typesV1.ComputeStateOutputV1{}
	err := sonic.UnmarshalString(string(traces), &computeState)
	if err != nil {
		p.logger.Sugar().Error(err)
		return nil, nil, errors.New("could not decode")
	}

	appTools := tools.Tools{Logger: p.logger}
	var transactions []*types.Transaction
	p.addresses = types.NewAddressInfoMap()
	tipsetKey := tipset.Key()
	tipsetCid := tipset.GetCidString()

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
			}, int64(tipset.Height()), tipsetKey)

			blockCid, err := appTools.GetBlockCidFromMsgCid(trace.MsgCid.String(), txType, nil, tipset)
			if err != nil {
				p.logger.Sugar().Errorf("Error when trying to get block cid from message,txType '%s': %v", txType, err)
			}
			messageUuid := tools.BuildMessageId(tipsetCid, blockCid, trace.MsgCid.String(), trace.Msg.Cid().String(), uuid.Nil.String())

			badTx := &types.Transaction{
				TxBasicBlockData: types.TxBasicBlockData{
					BasicBlockData: types.BasicBlockData{
						Height:    uint64(tipset.Height()),
						TipsetCid: tipsetCid,
					},
					BlockCid: blockCid,
				},
				Id:          messageUuid,
				ParentId:    uuid.Nil.String(),
				TxCid:       trace.MsgCid.String(),
				TxFrom:      trace.Msg.From.String(),
				TxTo:        trace.Msg.To.String(),
				TxType:      txType,
				Amount:      trace.Msg.Value.Int,
				GasUsed:     uint64(trace.MsgRct.GasUsed),
				Status:      parser.GetExitCodeStatus(trace.MsgRct.ExitCode),
				TxMetadata:  trace.Error,
				TxTimestamp: parser.GetTimestamp(tipset.MinTimestamp()),
			}

			transactions = append(transactions, badTx)
			continue
		}

		// Main transaction
		transaction, err := p.parseTrace(trace.ExecutionTrace, trace.MsgCid, tipset, ethLogs, trace.GasCost, uuid.Nil.String())
		if err != nil {
			continue
		}
		transaction.GasUsed = trace.GasCost.GasUsed.Uint64()
		transactions = append(transactions, transaction)

		// Only process sub-calls if the parent call was successfully executed
		if trace.ExecutionTrace.MsgRct.ExitCode.IsSuccess() {
			subTxs := p.parseSubTxs(trace.ExecutionTrace.Subcalls, trace.MsgCid, tipset, ethLogs,
				trace.Msg.Cid().String(), transaction.Id, 0)
			if len(subTxs) > 0 {
				transactions = append(transactions, subTxs...)
			}
		}
	}

	transactions = tools.SetNodeMetadataOnTxs(transactions, metadata, Version)

	// Clear this cache when we finish processing a tipset.
	// Bad addresses in this tipset might be valid in the next one
	p.helper.GetActorsCache().ClearBadAddressCache()
	return transactions, p.addresses, nil
}

func (p *Parser) GetBaseFee(traces []byte, tipset *types.ExtendedTipSet) (uint64, error) {
	// Unmarshal into vComputeState
	computeState := &typesV1.ComputeStateOutputV1{}
	if err := sonic.UnmarshalString(string(traces), &computeState); err != nil {
		p.logger.Sugar().Error(err)
		return 0, errors.New("could not decode")
	}

	baseFee := big.NewInt(0)
	found := false
	for _, trace := range computeState.Trace {
		baseFeeBurn := trace.GasCost.BaseFeeBurn
		gasUsed := trace.GasCost.GasUsed
		if gasUsed.IsZero() {
			continue
		}

		found = true
		baseFee.Div(baseFeeBurn.Int, gasUsed.Int)
		break
	}

	if !found {
		return parser.GetParentBaseFeeByHeight(tipset, p.logger)
	}

	return baseFee.Uint64(), nil
}

func (p *Parser) parseSubTxs(subTxs []typesV1.ExecutionTraceV1, mainMsgCid cid.Cid, tipSet *types.ExtendedTipSet, ethLogs []types.EthLog, txHash string,
	parentId string, level uint16) (txs []*types.Transaction) {
	level++
	gasCost := api.MsgGasCost{TotalCost: abi.NewTokenAmount(0)} // It is necessary to avoid adding fees to transactions with level != 0

	for _, subTx := range subTxs {
		subTransaction, err := p.parseTrace(subTx, mainMsgCid, tipSet, ethLogs, gasCost, parentId)
		if err != nil {
			continue
		}

		subTransaction.Level = level
		txs = append(txs, subTransaction)
		txs = append(txs, p.parseSubTxs(subTx.Subcalls, mainMsgCid, tipSet, ethLogs, txHash, subTransaction.Id, level)...)
	}
	return
}

func (p *Parser) parseTrace(trace typesV1.ExecutionTraceV1, mainMsgCid cid.Cid, tipset *types.ExtendedTipSet, ethLogs []types.EthLog, gasCost api.MsgGasCost, parentId string) (*types.Transaction, error) {
	txType, err := p.helper.GetMethodName(&parser.LotusMessage{
		To:     trace.Msg.To,
		From:   trace.Msg.From,
		Method: trace.Msg.Method,
	}, int64(tipset.Height()), tipset.Key())

	if err != nil {
		p.logger.Sugar().Errorf("Error when trying to get method name in tx cid'%s': %v", mainMsgCid.String(), err)
		txType = parser.UnknownStr
	}
	if err == nil && txType == parser.UnknownStr {
		p.logger.Sugar().Errorf("Could not get method name in transaction '%s'", trace.Msg.Cid().String())
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
	}, int64(tipset.Height()), tipset.Key(), ethLogs)

	if mErr != nil {
		p.logger.Sugar().Warnf("Could not get metadata for transaction in height %s of type '%s': %s", tipset.Height().String(), txType, mErr.Error())
	}
	if addressInfo != nil {
		parser.AppendToAddressesMap(p.addresses, addressInfo)
	}
	if trace.MsgRct.ExitCode.IsError() {
		metadata["Error"] = trace.MsgRct.ExitCode.Error()
	}

	tipsetCid := tipset.GetCidString()
	jsonMetadata, _ := json.Marshal(metadata)

	p.appendAddressInfo(trace.Msg, tipset.Key())

	appTools := tools.Tools{Logger: p.logger}
	blockCid, err := appTools.GetBlockCidFromMsgCid(mainMsgCid.String(), txType, metadata, tipset)
	if err != nil {
		p.logger.Sugar().Errorf("Error when trying to get block cid from message, txType '%s': %v", txType, err)
	}

	messageUuid := tools.BuildMessageId(tipsetCid, blockCid, mainMsgCid.String(), trace.Msg.Cid().String(), parentId)

	tx := &types.Transaction{
		TxBasicBlockData: types.TxBasicBlockData{
			BasicBlockData: types.BasicBlockData{
				Height:    uint64(tipset.Height()),
				TipsetCid: tipsetCid,
			},
			BlockCid: blockCid,
		},
		ParentId:    parentId,
		Id:          messageUuid,
		TxTimestamp: parser.GetTimestamp(tipset.MinTimestamp()),
		TxCid:       mainMsgCid.String(),
		TxFrom:      trace.Msg.From.String(),
		TxTo:        trace.Msg.To.String(),
		Amount:      trace.Msg.Value.Int,
		Status:      parser.GetExitCodeStatus(trace.MsgRct.ExitCode),
		TxType:      txType,
		TxMetadata:  string(jsonMetadata),
	}

	if gasCost.TotalCost.Uint64() > 0 {
		transactionFeesJson := p.calculateTransactionFees(gasCost, tipset, blockCid)
		tx.FeeData = string(transactionFeesJson)
	}
	return tx, nil
}

func hasMessage(trace *typesV1.InvocResultV1) bool {
	return trace.Msg != nil
}

func hasExecutionTrace(trace *typesV1.InvocResultV1) bool {
	// check if this execution trace is valid
	if trace.ExecutionTrace.Msg == nil || trace.ExecutionTrace.MsgRct == nil {
		// this is an invalid message
		return false
	}
	return true
}

func (p *Parser) appendAddressInfo(msg *filTypes.Message, key filTypes.TipSetKey) {
	if msg == nil {
		return
	}
	fromAdd := p.helper.GetActorAddressInfo(msg.From, key)
	toAdd := p.helper.GetActorAddressInfo(msg.To, key)
	parser.AppendToAddressesMap(p.addresses, fromAdd, toAdd)
}

func (p *Parser) calculateTransactionFees(gasCost api.MsgGasCost, tipset *types.ExtendedTipSet, blockCid string) []byte {
	minerAddress, err := tipset.GetBlockMiner(blockCid)
	if err != nil {
		p.logger.Sugar().Errorf("Error when trying to get miner address from block cid '%s': %v", blockCid, err)
	}

	feeData := parser.FeeData{
		FeesMetadata: parser.FeesMetadata{
			MinerFee: parser.MinerFee{
				MinerAddress: minerAddress,
				Amount:       gasCost.MinerTip.String(),
			},
			OverEstimationBurnFee: parser.OverEstimationBurnFee{
				BurnAddress: parser.BurnAddress,
				Amount:      gasCost.OverEstimationBurn.String(),
			},
			BurnFee: parser.BurnFee{
				BurnAddress: parser.BurnAddress,
				Amount:      gasCost.BaseFeeBurn.String(),
			},
		},
		Amount: gasCost.TotalCost.Int,
	}

	data, _ := json.Marshal(feeData)

	return data
}
