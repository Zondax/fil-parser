package v1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"

	"github.com/bytedance/sonic"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/google/uuid"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/actors"
	cacheMetrics "github.com/zondax/fil-parser/actors/cache/metrics"
	actorsV1 "github.com/zondax/fil-parser/actors/v1"
	actorsV2 "github.com/zondax/fil-parser/actors/v2"
	logger2 "github.com/zondax/fil-parser/logger"
	"github.com/zondax/fil-parser/metrics"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/parser/helper"
	parsermetrics "github.com/zondax/fil-parser/parser/metrics"
	typesV1 "github.com/zondax/fil-parser/parser/v1/types"
	"github.com/zondax/fil-parser/tools"
	multisigTools "github.com/zondax/fil-parser/tools/multisig"
	"github.com/zondax/fil-parser/types"
	"github.com/zondax/golem/pkg/logger"
)

const Version = "v1"

var NodeVersionsSupported = []string{"v1.21", "v1.22"}

type Parser struct {
	actorParser            actors.ActorParserInterface
	addresses              *types.AddressInfoMap
	txCidEquivalents       []types.TxCidTranslation
	helper                 *helper.Helper
	logger                 *logger.Logger
	multisigEventGenerator multisigTools.EventGenerator
	metrics                *parsermetrics.ParserMetricsClient
	actorsCacheMetrics     *cacheMetrics.ActorsCacheMetricsClient
	config                 parser.Config
	network                string
}

func NewParser(helper *helper.Helper, logger *logger.Logger, metrics metrics.MetricsClient, config parser.Config) *Parser {
	network, err := helper.GetFilecoinNodeClient().StateNetworkName(context.Background())
	if err != nil {
		logger.Fatal(err.Error())
		return nil
	}
	networkName := tools.ParseRawNetworkName(string(network))
	return &Parser{
		network:                networkName,
		actorParser:            actorsV1.NewActorParser(helper, logger, metrics),
		addresses:              types.NewAddressInfoMap(),
		helper:                 helper,
		logger:                 logger2.GetSafeLogger(logger),
		multisigEventGenerator: multisigTools.NewEventGenerator(helper, logger2.GetSafeLogger(logger), metrics),
		metrics:                parsermetrics.NewClient(metrics, "parserV1"),
		actorsCacheMetrics:     cacheMetrics.NewClient(metrics, "actorsCache"),
		config:                 config,
	}
}

func NewActorsV2Parser(network string, helper *helper.Helper, logger *logger.Logger, metrics metrics.MetricsClient, config parser.Config) *Parser {
	return &Parser{
		actorParser:            actorsV2.NewActorParser(network, helper, logger, metrics),
		addresses:              types.NewAddressInfoMap(),
		helper:                 helper,
		logger:                 logger2.GetSafeLogger(logger),
		multisigEventGenerator: multisigTools.NewEventGenerator(helper, logger2.GetSafeLogger(logger), metrics),
		metrics:                parsermetrics.NewClient(metrics, "parserV1"),
		actorsCacheMetrics:     cacheMetrics.NewClient(metrics, "actorsCache"),
		config:                 config,
	}
}

func (p *Parser) GetConfig() parser.Config {
	return p.config
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

func (p *Parser) ParseTransactions(ctx context.Context, txsData types.TxsData) (*types.TxsParsedResult, error) {
	// Unmarshal into vComputeState
	computeState := &typesV1.ComputeStateOutputV1{}
	err := sonic.UnmarshalString(string(txsData.Traces), &computeState)
	if err != nil {
		p.logger.Error(err.Error())
		return nil, errors.New("could not decode")
	}

	var transactions []*types.Transaction
	p.addresses = types.NewAddressInfoMap()
	p.txCidEquivalents = make([]types.TxCidTranslation, 0)

	tipsetCid := txsData.Tipset.GetCidString()

	for _, trace := range computeState.Trace {
		if !hasMessage(trace) {
			continue
		}

		systemExecution := false
		// TODO find a way to not having this special case handled outside func parseTrace
		if ok := hasExecutionTrace(trace); !ok {
			if trace.Msg != nil {
				systemExecution = p.helper.IsSystemActor(trace.Msg.From) && p.helper.IsSystemActor(trace.Msg.To)
			}
			// Create tx
			actorName, txType, err := p.getTxType(ctx, trace.Msg.To, trace.Msg.From, trace.Msg.Method, trace.MsgCid, txsData.Tipset)
			if err != nil {
				p.logger.Errorf("Error when trying to get tx type: %v", err)
				_ = p.metrics.UpdateMethodNameErrorMetric(actorName, fmt.Sprint(trace.Msg.Method))
				continue
			}
			var blockCid string
			if !systemExecution {
				blockCid, err = actorsV2.GetBlockCidFromMsgCid(trace.MsgCid.String(), txType, nil, txsData.Tipset, p.logger)
				if err != nil {
					_ = p.metrics.UpdateBlockCidFromMsgCidMetric(txType)
					p.logger.Errorf("Error when trying to get block cid from message,txType '%s': cid '%s': %v", txType, trace.MsgCid.String(), err)
				}
			}
			messageUuid := tools.BuildMessageId(tipsetCid, blockCid, trace.MsgCid.String(), trace.Msg.Cid().String(), uuid.Nil.String())

			txFrom, txTo := p.getFromToRobustAddresses(trace.Msg.From, trace.Msg.To)

			badTx := &types.Transaction{
				TxBasicBlockData: types.TxBasicBlockData{
					BasicBlockData: types.BasicBlockData{
						// #nosec G115
						Height:    uint64(txsData.Tipset.Height()),
						TipsetCid: tipsetCid,
					},
					BlockCid: blockCid,
				},
				Id:       messageUuid,
				ParentId: uuid.Nil.String(),
				TxCid:    trace.MsgCid.String(),
				TxFrom:   txFrom,
				TxTo:     txTo,
				TxType:   txType,
				Amount:   trace.Msg.Value.Int,
				// #nosec G115
				GasUsed:     uint64(trace.MsgRct.GasUsed),
				Status:      parser.GetExitCodeStatus(trace.MsgRct.ExitCode),
				TxMetadata:  trace.Error,
				TxTimestamp: parser.GetTimestamp(txsData.Tipset.MinTimestamp()),
			}

			transactions = append(transactions, badTx)
			continue
		}
		systemExecution = p.helper.IsSystemActor(trace.ExecutionTrace.Msg.From) && p.helper.IsSystemActor(trace.ExecutionTrace.Msg.To)

		// Main transaction
		transaction, err := p.parseTrace(ctx, trace.ExecutionTrace, trace.MsgCid, txsData.Tipset, uuid.Nil.String(), systemExecution)
		if err != nil {
			continue
		}
		transaction.GasUsed = trace.GasCost.GasUsed.Uint64()
		transactions = append(transactions, transaction)

		// Only process sub-calls if the parent call was successfully executed
		if trace.ExecutionTrace.MsgRct.ExitCode.IsSuccess() {
			subTxs := p.parseSubTxs(ctx, trace.ExecutionTrace.Subcalls, trace.MsgCid, txsData.Tipset, txsData.EthLogs,
				trace.Msg.Cid().String(), transaction.Id, 0, systemExecution)
			if len(subTxs) > 0 {
				transactions = append(transactions, subTxs...)
			}
		}

		// Fees
		if trace.GasCost.TotalCost.Uint64() > 0 {
			feeTx := p.feesTransactions(trace, txsData.Tipset, transaction.TxType, transaction.Id, systemExecution)
			if p.config.FeesAsColumn {
				transaction.FeeData = feeTx.TxMetadata
			} else {
				transactions = append(transactions, feeTx)
			}
		}

		// TxCid <-> TxHash
		if int64(txsData.Tipset.Height()) >= p.config.TxCidTranslationStart {
			txHash, err := parser.TranslateTxCidToTxHash(p.helper.GetFilecoinNodeClient(), trace.MsgCid, p.actorsCacheMetrics)
			if err == nil && txHash != "" {
				p.txCidEquivalents = append(p.txCidEquivalents, types.TxCidTranslation{TxCid: trace.MsgCid.String(), TxHash: txHash})
			}
			if err != nil {
				p.logger.Warnf("Error when trying to translate tx cid to tx hash: %v", err)
				_ = p.metrics.UpdateTranslateTxCidToTxHashMetric()
			}
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

func (p *Parser) ParseMultisigEvents(ctx context.Context, multisigTxs []*types.Transaction, tipsetCid string, tipsetKey filTypes.TipSetKey) (*types.MultisigEvents, error) {
	return nil, errors.New("unimplimented")
}

func (p *Parser) ParseMinerEvents(ctx context.Context, minerTxs []*types.Transaction, tipsetCid string, tipsetKey filTypes.TipSetKey) (*types.MinerEvents, error) {
	return nil, errors.New("unimplimented")
}

func (p *Parser) ParseNativeEvents(_ context.Context, _ types.EventsData) (*types.EventsParsedResult, error) {
	return nil, errors.New("unimplimented")
}

func (p *Parser) ParseEthLogs(_ context.Context, _ types.EventsData) (*types.EventsParsedResult, error) {
	return nil, errors.New("unimplimented")
}

func (p *Parser) GetBaseFee(traces []byte, tipset *types.ExtendedTipSet) (uint64, error) {
	// Unmarshal into vComputeState
	computeState := &typesV1.ComputeStateOutputV1{}
	if err := sonic.UnmarshalString(string(traces), &computeState); err != nil {
		p.logger.Error(err.Error())
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

func (p *Parser) parseSubTxs(ctx context.Context, subTxs []typesV1.ExecutionTraceV1, mainMsgCid cid.Cid, tipSet *types.ExtendedTipSet, ethLogs []types.EthLog, txHash string,
	parentId string, level uint16, systemExecution bool) (txs []*types.Transaction) {
	level++
	for _, subTx := range subTxs {
		subTransaction, err := p.parseTrace(ctx, subTx, mainMsgCid, tipSet, parentId, systemExecution)
		if err != nil {
			continue
		}

		subTransaction.Level = level
		txs = append(txs, subTransaction)
		txs = append(txs, p.parseSubTxs(ctx, subTx.Subcalls, mainMsgCid, tipSet, ethLogs, txHash, subTransaction.Id, level, systemExecution)...)
	}
	return
}

func (p *Parser) parseTrace(ctx context.Context, trace typesV1.ExecutionTraceV1, mainMsgCid cid.Cid, tipset *types.ExtendedTipSet, parentId string, systemExecution bool) (*types.Transaction, error) {
	failedTx := trace.MsgRct.ExitCode.IsError()

	actorName, txType, err := p.getTxType(ctx, trace.Msg.To, trace.Msg.From, trace.Msg.Method, mainMsgCid, tipset)
	if err != nil {
		txType = parser.UnknownStr
	}

	if !failedTx && (txType == parser.UnknownStr || err != nil) {
		_ = p.metrics.UpdateMethodNameErrorMetric(actorName, fmt.Sprint(trace.Msg.Method))
		p.logger.Errorf("Could not get method name in transaction '%s' : method: %d height: %d err: %s", trace.Msg.Cid().String(), trace.Msg.Method, tipset.Height(), err)
	}
	actor, metadata, addressInfo, mErr := p.actorParser.GetMetadata(ctx, actorName, txType, &parser.LotusMessage{
		To:     trace.Msg.To,
		From:   trace.Msg.From,
		Method: trace.Msg.Method,
		Cid:    mainMsgCid,
		Params: trace.Msg.Params,
	}, mainMsgCid, &parser.LotusMessageReceipt{
		ExitCode: trace.MsgRct.ExitCode,
		Return:   trace.MsgRct.Return,
	}, int64(tipset.Height()), tipset.Key())

	if mErr != nil && !failedTx {
		_ = p.metrics.UpdateMetadataErrorMetric(actor, txType)
		p.logger.Warnf("Could not get metadata for transaction in height %s of type '%s': %s", tipset.Height().String(), txType, mErr.Error())
	}
	if trace.MsgRct.ExitCode.IsSuccess() && addressInfo != nil {
		parser.AppendToAddressesMap(p.addresses, addressInfo)
	}

	if metadata == nil {
		metadata = map[string]interface{}{
			parser.ParamsRawKey: trace.Msg.Params,
			parser.ReturnRawKey: trace.MsgRct.Return,
		}
	}

	if trace.MsgRct.ExitCode.IsError() {
		metadata[parser.ErrorKey] = trace.MsgRct.ExitCode.Error()
	}

	metadata[parser.MethodNumKey] = trace.Msg.Method.String()

	tipsetCid := tipset.GetCidString()
	jsonMetadata, err := json.Marshal(metadata)
	if err != nil {
		_ = p.metrics.UpdateJsonMarshalMetric(parsermetrics.MetadataValue, txType)
	}

	p.appendAddressInfo(trace.Msg, tipset.Key(), tipset.Height())

	var blockCid string
	if !systemExecution {
		blockCid, err = actorsV2.GetBlockCidFromMsgCid(mainMsgCid.String(), txType, metadata, tipset, p.logger)
		if err != nil {
			_ = p.metrics.UpdateBlockCidFromMsgCidMetric(txType)
			p.logger.Errorf("Error when trying to get block cid from message, txType '%s' cid '%s': %v", txType, mainMsgCid.String(), err)
		}
	}

	messageUuid := tools.BuildMessageId(tipsetCid, blockCid, mainMsgCid.String(), trace.Msg.Cid().String(), parentId)

	txFrom, txTo := p.getFromToRobustAddresses(trace.Msg.From, trace.Msg.To)

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
		TxFrom:      txFrom,
		TxTo:        txTo,
		Amount:      trace.Msg.Value.Int,
		Status:      parser.GetExitCodeStatus(trace.MsgRct.ExitCode),
		TxType:      txType,
		TxMetadata:  string(jsonMetadata),
	}, nil
}

func (p *Parser) feesTransactions(msg *typesV1.InvocResultV1, tipset *types.ExtendedTipSet, txType, parentTxId string, systemExecution bool) *types.Transaction {
	var blockCid string
	var err error

	timestamp := parser.GetTimestamp(tipset.MinTimestamp())
	if !systemExecution {
		blockCid, err = actorsV2.GetBlockCidFromMsgCid(msg.MsgCid.String(), txType, nil, tipset, p.logger)
		if err != nil {
			p.logger.Errorf("Error when trying to get block cid from message, txType '%s' cid '%s': %v", txType, msg.MsgCid.String(), err)
		}
	}

	metadata := p.feesMetadata(msg, tipset, txType, blockCid, systemExecution)

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
		TxMetadata:  metadata,
	}
}

func (p *Parser) feesMetadata(msg *typesV1.InvocResultV1, tipset *types.ExtendedTipSet, txType, blockCid string, systemExecution bool) string {
	var minerAddress string
	var err error
	if !systemExecution && blockCid != "" {
		minerAddress, err = tipset.GetBlockMiner(blockCid)
		if err != nil {
			// added a new error to avoid cardinality of GetBlockMiner error results which include cid
			_ = p.metrics.UpdateGetBlockMinerMetric(fmt.Sprint(uint64(msg.Msg.Method)), txType)
			p.logger.Errorf("Error when trying to get miner address from block cid '%s': %v", blockCid, err)
		}
	}

	if p.config.ConsolidateRobustAddress && minerAddress != "" {
		minerAddr, err := address.NewFromString(minerAddress)
		if err != nil {
			p.logger.Errorf("Error when trying to parse miner address: %v", err)
		}

		minerAddress, err = actors.ConsolidateToRobustAddress(minerAddr, p.helper, p.logger, p.config.RobustAddressBestEffort)
		if err != nil {
			p.logger.Errorf("Error when trying to consolidate miner address to robust: %v", err)
		}
	}

	if p.config.FeesAsColumn {
		txType = ""
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
		TotalCost: msg.GasCost.TotalCost.String(),
	}

	metadata, err := json.Marshal(feesMetadata)
	if err != nil {
		_ = p.metrics.UpdateJsonMarshalMetric(parsermetrics.FeesMetadataValue, txType)
	}

	return string(metadata)
}

func (p *Parser) getFromToRobustAddresses(from, to address.Address) (string, string) {
	var err error
	txFrom := from.String()
	txTo := to.String()
	if p.config.ConsolidateRobustAddress {
		txFrom, err = actors.ConsolidateToRobustAddress(from, p.helper, p.logger, p.config.RobustAddressBestEffort)
		if err != nil {
			txFrom = from.String()
			p.logger.Warnf("Could not consolidate robust address: %v", err)
		}
		txTo, err = actors.ConsolidateToRobustAddress(to, p.helper, p.logger, p.config.RobustAddressBestEffort)
		if err != nil {
			txTo = to.String()
			p.logger.Warnf("Could not consolidate robust address: %v", err)
		}
	}

	return txFrom, txTo
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

func (p *Parser) appendAddressInfo(msg *filTypes.Message, key filTypes.TipSetKey, height abi.ChainEpoch) {
	if msg == nil {
		return
	}
	fromAdd := p.helper.GetActorAddressInfo(msg.From, key, height)
	toAdd := p.helper.GetActorAddressInfo(msg.To, key, height)
	parser.AppendToAddressesMap(p.addresses, fromAdd, toAdd)
}

func (p *Parser) getTxType(ctx context.Context, to, from address.Address, method abi.MethodNum, mainMsgCid cid.Cid, tipset *types.ExtendedTipSet) (actorName string, txType string, err error) {
	msg := &parser.LotusMessage{
		To:     to,
		From:   from,
		Method: method,
	}
	_, actorName, err = p.helper.GetActorNameFromAddress(msg.To, int64(tipset.Height()), tipset.Key())
	if err != nil {
		p.logger.Errorf("Error when trying to get actor name in tx cid'%s': %v", mainMsgCid.String(), err)
	}

	txType, err = actorsV2.GetMethodName(ctx, msg.Method, actorName, int64(tipset.Height()), p.network, p.helper, p.logger)
	if err != nil {
		p.logger.Errorf("Error when trying to get method name in tx cid'%s' using v2: %v", mainMsgCid.String(), err)
		txType = parser.UnknownStr
	}

	return actorName, txType, err
}
