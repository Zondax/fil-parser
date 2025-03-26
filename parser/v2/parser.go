package v2

import (
	"cmp"
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"slices"
	"strings"

	"github.com/zondax/fil-parser/actors"
	logger2 "github.com/zondax/fil-parser/logger"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/parser/helper"
	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
	"github.com/zondax/fil-parser/tools"
	eventTools "github.com/zondax/fil-parser/tools/events"
	multisigTools "github.com/zondax/fil-parser/tools/multisig"
	"github.com/zondax/fil-parser/types"

	"github.com/bytedance/sonic"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/google/uuid"
	"github.com/ipfs/go-cid"
	"go.uber.org/zap"
)

const Version = "v2"

var NodeVersionsSupported = []string{"v1.23", "v1.24", "v1.25", "v1.26", "v1.27", "v1.28", "v1.29", "v1.30", "v1.31", "v1.32"}

type Parser struct {
	actorParser            *actors.ActorParser
	addresses              *types.AddressInfoMap
	txCidEquivalents       []types.TxCidTranslation
	helper                 *helper.Helper
	logger                 *zap.Logger
	multisigEventGenerator multisigTools.EventGenerator
}

func NewParser(helper *helper.Helper, logger *zap.Logger) *Parser {
	return &Parser{
		actorParser:            actors.NewActorParser(helper, logger),
		addresses:              types.NewAddressInfoMap(),
		helper:                 helper,
		logger:                 logger2.GetSafeLogger(logger),
		multisigEventGenerator: multisigTools.NewEventGenerator(helper, logger2.GetSafeLogger(logger)),
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

func (p *Parser) ParseTransactions(_ context.Context, txsData types.TxsData) (*types.TxsParsedResult, error) {
	// Unmarshal into vComputeState
	computeState := &typesV2.ComputeStateOutputV2{}
	err := sonic.UnmarshalString(string(txsData.Traces), &computeState)
	if err != nil {
		p.logger.Sugar().Error(err)
		return nil, errors.New("could not decode")
	}

	var transactions []*types.Transaction
	p.addresses = types.NewAddressInfoMap()
	p.txCidEquivalents = make([]types.TxCidTranslation, 0)

	if err != nil {
		return nil, parser.ErrBlockHash
	}
	for _, trace := range computeState.Trace {
		if trace.Msg == nil {
			continue
		}

		// Main transaction
		transaction, err := p.parseTrace(trace.ExecutionTrace, trace.MsgCid, txsData.Tipset, uuid.Nil.String())
		if err != nil {
			continue
		}

		// We only set the gas usage for the main transaction.
		// If we need the gas usage of all sub-txs, we need to also parse GasCharges (today is very inefficient)
		transaction.GasUsed = trace.GasCost.GasUsed.Uint64()

		transactions = append(transactions, transaction)

		// Only process sub-calls if the parent call was successfully executed
		if trace.ExecutionTrace.MsgRct.ExitCode.IsSuccess() {
			subTxs := p.parseSubTxs(trace.ExecutionTrace.Subcalls, trace.MsgCid, txsData.Tipset, txsData.EthLogs,
				trace.Msg.Cid().String(), transaction.Id, 0)
			if len(subTxs) > 0 {
				transactions = append(transactions, subTxs...)
			}
		}

		// Fees
		if trace.GasCost.TotalCost.Uint64() > 0 {
			feeTx := p.feesTransactions(trace, txsData.Tipset, transaction.TxType, transaction.Id)
			transactions = append(transactions, feeTx)
		}

		// TxCid <-> TxHash
		txHash, err := parser.TranslateTxCidToTxHash(p.helper.GetFilecoinNodeClient(), trace.MsgCid)
		if err == nil && txHash != "" {
			p.txCidEquivalents = append(p.txCidEquivalents, types.TxCidTranslation{TxCid: trace.MsgCid.String(), TxHash: txHash})
		}
	}

	transactions = tools.SetNodeMetadata(transactions, txsData.Metadata, Version)

	// Clear this cache when we finish processing a tipset.
	// Bad addresses in this tipset might be valid in the next one
	p.helper.GetActorsCache().ClearBadAddressCache()

	return &types.TxsParsedResult{
		Txs:       transactions,
		Addresses: p.addresses,
		TxCids:    p.txCidEquivalents,
	}, nil
}

func (p *Parser) ParseNativeEvents(_ context.Context, eventsData types.EventsData) (*types.EventsParsedResult, error) {
	var parsed []*types.Event
	nativeEventsTotal, evmEventsTotal := 0, 0
	for idx, nativeLog := range eventsData.NativeLog {
		// #nosec G115
		event, err := eventTools.ParseNativeLog(eventsData.Tipset, nativeLog, uint64(idx))
		if err != nil {
			return nil, err
		}

		if event.Type == types.EventTypeEVM {
			evmEventsTotal++
		} else if event.Type == types.EventTypeNative {
			nativeEventsTotal++
		}

		parsed = append(parsed, event)
	}

	parsed = tools.SetNodeMetadata(parsed, eventsData.Metadata, Version)

	return &types.EventsParsedResult{EVMEvents: evmEventsTotal, NativeEvents: nativeEventsTotal, ParsedEvents: parsed}, nil
}

func (p *Parser) ParseEthLogs(_ context.Context, eventsData types.EventsData) (*types.EventsParsedResult, error) {
	var parsed []*types.Event
	// sort the events by the TransactionIndex ASC and the logIndex ASC
	slices.SortFunc(eventsData.EthLogs, func(a, b types.EthLog) int {
		return cmp.Or(
			cmp.Compare(a.TransactionIndex, b.TransactionIndex),
			cmp.Compare(a.LogIndex, b.LogIndex),
		)
	})

	for idx, ethLog := range eventsData.EthLogs {
		// #nosec G115
		event, err := eventTools.ParseEthLog(eventsData.Tipset, ethLog, p.helper, uint64(idx))
		if err != nil {
			zap.S().Errorf("error retrieving selector_sig for hash: %s err: %s", event.SelectorID, err)
		}

		parsed = append(parsed, event)
	}

	parsed = tools.SetNodeMetadata(parsed, eventsData.Metadata, Version)

	return &types.EventsParsedResult{EVMEvents: len(parsed), ParsedEvents: parsed}, nil
}

func (p *Parser) ParseMultisigEvents(ctx context.Context, multisigTxs []*types.Transaction, tipsetCid string, tipsetKey filTypes.TipSetKey) (*types.MultisigEvents, error) {
	return p.multisigEventGenerator.GenerateMultisigEvents(ctx, multisigTxs, tipsetCid, tipsetKey)
}

func (p *Parser) GetBaseFee(traces []byte, tipset *types.ExtendedTipSet) (uint64, error) {
	// Unmarshal into vComputeState
	computeState := &typesV2.ComputeStateOutputV2{}
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

func (p *Parser) parseSubTxs(subTxs []typesV2.ExecutionTraceV2, mainMsgCid cid.Cid, tipSet *types.ExtendedTipSet, ethLogs []types.EthLog, txHash string,
	parentId string, level uint16) (txs []*types.Transaction) {
	level++
	for _, subTx := range subTxs {
		subTransaction, err := p.parseTrace(subTx, mainMsgCid, tipSet, parentId)
		if err != nil {
			continue
		}

		subTransaction.Level = level
		txs = append(txs, subTransaction)
		txs = append(txs, p.parseSubTxs(subTx.Subcalls, mainMsgCid, tipSet, ethLogs, txHash, subTransaction.Id, level)...)
	}
	return
}

func (p *Parser) parseTrace(trace typesV2.ExecutionTraceV2, mainMsgCid cid.Cid, tipset *types.ExtendedTipSet, parentId string) (*types.Transaction, error) {
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
		p.logger.Sugar().Errorf("Could not get method name in transaction '%s'", mainMsgCid.String())
	}

	metadata, addressInfo, mErr := p.actorParser.GetMetadata(txType, &parser.LotusMessage{
		To:     trace.Msg.To,
		From:   trace.Msg.From,
		Method: trace.Msg.Method,
		Cid:    mainMsgCid,
		Params: trace.Msg.Params,
	}, mainMsgCid, &parser.LotusMessageReceipt{
		ExitCode: trace.MsgRct.ExitCode,
		Return:   trace.MsgRct.Return,
	}, int64(tipset.Height()), tipset.Key())

	if mErr != nil {
		p.logger.Sugar().Warnf("Could not get metadata for transaction in height %s of type '%s': %s", tipset.Height().String(), txType, mErr.Error())
	}
	if addressInfo != nil {
		parser.AppendToAddressesMap(p.addresses, addressInfo)
	}
	if trace.MsgRct.ExitCode.IsError() {
		metadata["Error"] = trace.MsgRct.ExitCode.Error()
	}

	jsonMetadata, _ := json.Marshal(metadata)

	p.appendAddressInfo(&parser.LotusMessage{
		To:     trace.Msg.To,
		From:   trace.Msg.From,
		Method: trace.Msg.Method,
		Cid:    mainMsgCid,
		Params: trace.Msg.Params,
	}, tipset.Key())

	appTools := tools.Tools{Logger: p.logger}
	blockCid, err := appTools.GetBlockCidFromMsgCid(mainMsgCid.String(), txType, metadata, tipset)
	if err != nil {
		p.logger.Sugar().Errorf("Error when trying to get block cid from message, txType '%s': %v", txType, err)
	}

	msgCid, err := tools.BuildCidFromMessageTrace(trace.Msg, mainMsgCid.String())
	if err != nil {
		p.logger.Sugar().Errorf("Error when trying to build message cid in tx cid'%s': %v", mainMsgCid.String(), err)
	}

	tipsetCid := tipset.GetCidString()
	messageUuid := tools.BuildMessageId(tipsetCid, blockCid, mainMsgCid.String(), msgCid, parentId)

	return &types.Transaction{
		TxBasicBlockData: types.TxBasicBlockData{
			BasicBlockData: types.BasicBlockData{
				// #nosec G115
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
	}, nil
}

func (p *Parser) feesTransactions(msg *typesV2.InvocResultV2, tipset *types.ExtendedTipSet, txType, parentTxId string) *types.Transaction {
	timestamp := parser.GetTimestamp(tipset.MinTimestamp())
	appTools := tools.Tools{Logger: p.logger}
	blockCid, err := appTools.GetBlockCidFromMsgCid(msg.MsgCid.String(), txType, nil, tipset)
	if err != nil {
		p.logger.Sugar().Errorf("Error when trying to get block cid from message, txType '%s': %v", txType, err)
	}

	minerAddress, err := tipset.GetBlockMiner(blockCid)
	if err != nil {
		p.logger.Sugar().Errorf("Error when trying to get miner address from block cid '%s': %v", blockCid, err)
	}

	feesMetadata := parser.FeesMetadata{
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

	metadata, _ := json.Marshal(feesMetadata)
	feeID := tools.BuildFeeId(tipset.GetCidString(), blockCid, msg.MsgCid.String())

	return &types.Transaction{
		TxBasicBlockData: types.TxBasicBlockData{
			BasicBlockData: types.BasicBlockData{
				// #nosec G115
				Height:    uint64(tipset.Height()),
				TipsetCid: tipset.GetCidString(),
			},
			BlockCid: blockCid,
		},
		Id:          feeID,
		ParentId:    parentTxId,
		TxTimestamp: timestamp,
		TxCid:       msg.MsgCid.String(),
		TxFrom:      msg.Msg.From.String(),
		Amount:      msg.GasCost.TotalCost.Int,
		Status:      "Ok",
		TxType:      parser.TotalFeeOp,
		TxMetadata:  string(metadata),
	}
}

func (p *Parser) appendAddressInfo(msg *parser.LotusMessage, key filTypes.TipSetKey) {
	if msg == nil {
		return
	}
	fromAdd := p.helper.GetActorAddressInfo(msg.From, key)
	toAdd := p.helper.GetActorAddressInfo(msg.To, key)
	parser.AppendToAddressesMap(p.addresses, fromAdd, toAdd)
}
