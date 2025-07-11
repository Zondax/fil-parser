package v2

import (
	"cmp"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"slices"
	"strings"

	"github.com/bytedance/sonic"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/exitcode"
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
	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
	"github.com/zondax/fil-parser/tools"
	eventTools "github.com/zondax/fil-parser/tools/events"
	minerTools "github.com/zondax/fil-parser/tools/miner"
	multisigTools "github.com/zondax/fil-parser/tools/multisig"
	"github.com/zondax/fil-parser/types"
	"github.com/zondax/golem/pkg/logger"
	golemBackoff "github.com/zondax/golem/pkg/zhttpclient/backoff"
)

const Version = "v2"

var NodeVersionsSupported = []string{"v1.23", "v1.24", "v1.25", "v1.26", "v1.27", "v1.28", "v1.29", "v1.30", "v1.31", "v1.32"}

type Parser struct {
	network                string
	actorParser            actors.ActorParserInterface
	addresses              *types.AddressInfoMap
	txCidEquivalents       []types.TxCidTranslation
	helper                 *helper.Helper
	logger                 *logger.Logger
	multisigEventGenerator multisigTools.EventGenerator
	minerEventGenerator    minerTools.EventGenerator
	metrics                *parsermetrics.ParserMetricsClient
	actorsCacheMetrics     *cacheMetrics.ActorsCacheMetricsClient
	backoff                *golemBackoff.BackOff
	config                 parser.Config
}

func NewParser(helper *helper.Helper, logger *logger.Logger, metrics metrics.MetricsClient, backoff *golemBackoff.BackOff, config parser.Config) *Parser {
	network, err := helper.GetFilecoinNodeClient().StateNetworkName(context.Background())
	if err != nil {
		logger.Fatal(err.Error())
		return nil
	}
	networkName := tools.ParseRawNetworkName(string(network))
	p := &Parser{
		network:                networkName,
		actorParser:            actorsV1.NewActorParser(helper, logger, metrics),
		addresses:              types.NewAddressInfoMap(),
		helper:                 helper,
		logger:                 logger2.GetSafeLogger(logger),
		multisigEventGenerator: multisigTools.NewEventGenerator(helper, logger2.GetSafeLogger(logger), metrics, config),
		minerEventGenerator:    minerTools.NewEventGenerator(helper, logger2.GetSafeLogger(logger), metrics, config),
		metrics:                parsermetrics.NewClient(metrics, "parserV2"),
		actorsCacheMetrics:     cacheMetrics.NewClient(metrics, "actorsCache"),
		config:                 config,
		backoff:                backoff,
	}

	return p
}

func NewActorsV2Parser(network string, helper *helper.Helper, logger *logger.Logger, metrics metrics.MetricsClient, backoff *golemBackoff.BackOff, config parser.Config) *Parser {
	return &Parser{
		network:                network,
		actorParser:            actorsV2.NewActorParser(network, helper, logger, metrics),
		addresses:              types.NewAddressInfoMap(),
		helper:                 helper,
		logger:                 logger2.GetSafeLogger(logger),
		multisigEventGenerator: multisigTools.NewEventGenerator(helper, logger2.GetSafeLogger(logger), metrics, config),
		minerEventGenerator:    minerTools.NewEventGenerator(helper, logger2.GetSafeLogger(logger), metrics, config),
		metrics:                parsermetrics.NewClient(metrics, "parserV2"),
		actorsCacheMetrics:     cacheMetrics.NewClient(metrics, "actorsCache"),
		config:                 config,
		backoff:                backoff,
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
	computeState := &typesV2.ComputeStateOutputV2{}
	err := sonic.UnmarshalString(string(txsData.Traces), &computeState)
	if err != nil {
		p.logger.Error(err.Error())
		return nil, errors.New("could not decode")
	}

	var transactions []*types.Transaction
	p.addresses = types.NewAddressInfoMap()
	p.txCidEquivalents = make([]types.TxCidTranslation, 0)

	for _, trace := range computeState.Trace {
		if trace.Msg == nil {
			p.logger.Errorf("Trace without message: %s", trace.MsgCid.String())
			_ = p.metrics.UpdateTraceWithoutMessageMetric()
			continue
		}

		// Observed in the wild, the main tx data is not the same as the execution trace data.
		// For those cases, we want to observe it, and use the top-level tx info as fallback.
		// Ex: Height 501104
		// Ex: https://filfox.info/en/message/bafy2bzacecgh3w4qk3t53kg6skweh4isekscy22s4nfxdyc7zcbfb4ez6li5m?t=2
		if trace.ExecutionTrace.MsgRct.ExitCode != trace.MsgRct.ExitCode {
			p.logger.Errorf("ExecutionTrace.MsgRct.ExitCode != MsgRct.ExitCode for tx %s", trace.MsgCid.String())
			_ = p.metrics.UpdateMismatchExitCodeMetric()

			trace.ExecutionTrace.Msg = *LotusMsgToExecutionTraceMsg(trace.Msg)
			trace.ExecutionTrace.MsgRct = *LotusMsgRctToExecutionTraceMsgRct(trace.MsgRct)
			trace.ExecutionTrace.Error = trace.Error
			trace.ExecutionTrace.Duration = trace.Duration
		}

		// check the 1st execution
		systemExecution := p.helper.IsSystemActor(trace.ExecutionTrace.Msg.From) && p.helper.IsSystemActor(trace.ExecutionTrace.Msg.To)

		mainMsgCid := trace.MsgCid
		mainMsgExitCode := trace.MsgRct.ExitCode
		transaction, err := p.parseTrace(ctx, trace.ExecutionTrace, mainMsgCid, txsData.Tipset, uuid.Nil.String(), systemExecution, mainMsgExitCode)
		if err != nil {
			p.logger.Errorf("Error parsing trace for tx %s: %v", mainMsgCid, err)
			_ = p.metrics.UpdateParseTraceMetric()
			continue
		}

		// We only set the gas usage for the main transaction.
		// If we need the gas usage of all sub-txs, we need to also parse GasCharges (today is very inefficient)
		transaction.GasUsed = trace.GasCost.GasUsed.Uint64()

		transactions = append(transactions, transaction)

		// note: we are using the parent MsgRct.ExitCode not the ExecutionTrace.MsgRct.ExitCode
		subTxs := p.parseSubTxs(ctx, trace.ExecutionTrace.Subcalls, mainMsgCid, txsData.Tipset, txsData.EthLogs,
			trace.Msg.Cid().String(), transaction.Id, 0, systemExecution, mainMsgExitCode)
		if len(subTxs) > 0 {
			transactions = append(transactions, subTxs...)
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
			txHash, err := parser.TranslateTxCidToTxHash(p.helper.GetFilecoinNodeClient(), trace.MsgCid, p.actorsCacheMetrics, p.backoff)
			if err == nil && txHash != "" {
				p.txCidEquivalents = append(p.txCidEquivalents, types.TxCidTranslation{TxCid: trace.MsgCid.String(), TxHash: txHash})
			}
			if err != nil {
				_ = p.metrics.UpdateTranslateTxCidToTxHashMetric()
				p.logger.Warnf("Error when trying to translate tx cid to tx hash: %v", err)
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

func (p *Parser) ParseNativeEvents(_ context.Context, eventsData types.EventsData) (*types.EventsParsedResult, error) {
	var parsed []*types.Event
	nativeEventsTotal, evmEventsTotal := 0, 0
	for idx, nativeLog := range eventsData.NativeLog {
		// #nosec G115
		event, err := eventTools.ParseNativeLog(eventsData.Tipset, nativeLog, uint64(idx), p.logger)
		if err != nil {
			_ = p.metrics.UpdateParseNativeEventsLogsMetric()
			return nil, err
		}

		if event.Type == types.EventTypeEVM {
			evmEventsTotal++
		} else if event.Type == types.EventTypeNative {
			nativeEventsTotal++
		}

		if p.config.ConsolidateRobustAddress {
			eventAddr, err := address.NewFromString(event.Emitter)
			if err != nil {
				return nil, err
			}
			if consolidatedAddr, err := actors.ConsolidateToRobustAddress(eventAddr, p.helper, p.logger, p.config.RobustAddressBestEffort); err == nil {
				event.Emitter = consolidatedAddr
			}
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
			_ = p.metrics.UpdateParseEthLogMetric()
			p.logger.Errorf("error retrieving selector_sig for hash: %s err: %s", event.SelectorID, err)
		}

		// we don't consolidate eth addresses
		if p.config.ConsolidateRobustAddress && !strings.HasPrefix(event.Emitter, parser.EthPrefix) {
			eventAddr, err := address.NewFromString(event.Emitter)
			if err != nil {
				return nil, fmt.Errorf("error parsing emitter address: %s: %w", event.Emitter, err)
			}
			if consolidatedAddr, err := actors.ConsolidateToRobustAddress(eventAddr, p.helper, p.logger, p.config.RobustAddressBestEffort); err == nil {
				event.Emitter = consolidatedAddr
			}
		}

		parsed = append(parsed, event)
	}

	parsed = tools.SetNodeMetadata(parsed, eventsData.Metadata, Version)

	return &types.EventsParsedResult{EVMEvents: len(parsed), ParsedEvents: parsed}, nil
}

func (p *Parser) ParseMultisigEvents(ctx context.Context, multisigTxs []*types.Transaction, tipsetCid string, tipsetKey filTypes.TipSetKey) (*types.MultisigEvents, error) {
	return p.multisigEventGenerator.GenerateMultisigEvents(ctx, multisigTxs, tipsetCid, tipsetKey)
}

func (p *Parser) ParseMinerEvents(ctx context.Context, minerTxs []*types.Transaction, tipsetCid string, tipsetKey filTypes.TipSetKey) (*types.MinerEvents, error) {
	return p.minerEventGenerator.GenerateMinerEvents(ctx, minerTxs, tipsetCid, tipsetKey)
}

func (p *Parser) GetBaseFee(traces []byte, tipset *types.ExtendedTipSet) (uint64, error) {
	// Unmarshal into vComputeState
	computeState := &typesV2.ComputeStateOutputV2{}
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

func (p *Parser) parseSubTxs(ctx context.Context, subTxs []typesV2.ExecutionTraceV2, mainMsgCid cid.Cid, tipSet *types.ExtendedTipSet, ethLogs []types.EthLog, txHash string,
	parentId string, level uint16, systemExecution bool, mainExitCode exitcode.ExitCode) (txs []*types.Transaction) {
	level++
	for _, subTx := range subTxs {
		subTransaction, err := p.parseTrace(ctx, subTx, mainMsgCid, tipSet, parentId, systemExecution, mainExitCode)
		if err != nil {
			continue
		}

		subTransaction.Level = level
		txs = append(txs, subTransaction)
		txs = append(txs, p.parseSubTxs(ctx, subTx.Subcalls, mainMsgCid, tipSet, ethLogs, txHash, subTransaction.Id, level, systemExecution, mainExitCode)...)
	}
	return
}

func (p *Parser) parseTrace(ctx context.Context, trace typesV2.ExecutionTraceV2, mainMsgCid cid.Cid, tipset *types.ExtendedTipSet, parentId string, systemExecution bool, mainExitCode exitcode.ExitCode) (*types.Transaction, error) {
	mainFailedTx := mainExitCode.IsError()
	subcallFailedTx := trace.MsgRct.ExitCode.IsError()
	actorName, txType, err := p.getTxType(ctx, trace, mainMsgCid, tipset)
	if err != nil {
		txType = parser.UnknownStr
	}

	// The main tx may be successful, but the subcall tx is failed, so we don't need to update the method name error metric
	if !subcallFailedTx && (txType == parser.UnknownStr || err != nil) {
		_ = p.metrics.UpdateMethodNameErrorMetric(actorName, fmt.Sprint(trace.Msg.Method), !subcallFailedTx, !mainFailedTx)
		p.logger.Errorf("Could not get method name in transaction '%s': %s", mainMsgCid.String(), err)
	}

	actor, metadata, addressInfo, mErr := p.actorParser.GetMetadata(ctx, actorName, txType, &parser.LotusMessage{
		To:     trace.Msg.To,
		From:   trace.Msg.From,
		Method: trace.Msg.Method,
		Cid:    mainMsgCid,
		Params: trace.Msg.Params,
	}, mainMsgCid, &parser.LotusMessageReceipt{
		ExitCode: mainExitCode,
		Return:   trace.MsgRct.Return,
	}, int64(tipset.Height()), tipset.Key())

	// Whenever a tx fails, it could be for multiple reasons.
	// It can happen that the tx is unknown, the to address actor is not valid, or the params are wrong.
	// If that it is the case, we will fail to parse it too... so we don't want to count it as a failed tx.
	if mErr != nil && (!subcallFailedTx || !mainFailedTx) {
		_ = p.metrics.UpdateMetadataErrorMetric(actor, txType, !subcallFailedTx, !mainFailedTx)
		p.logger.Errorf("Could not get metadata for transaction in height %s of type '%s': %s", tipset.Height().String(), txType, mErr.Error())
	}

	// If the tx failed, we don't want to add the address info to the addresses map, as we could be adding bad relationships between short and robust..
	if !mainFailedTx && addressInfo != nil {
		parser.AppendToAddressesMap(p.addresses, addressInfo)
	}
	if len(metadata) == 0 {
		metadata = map[string]interface{}{
			parser.ParamsRawKey: trace.Msg.Params,
			parser.ReturnRawKey: trace.MsgRct.Return,
		}
	}

	// If the mainExitCode is an error, we want to add the error to the metadata (for the the corresponding tx, main or subcall)
	if subcallFailedTx {
		metadata[parser.ErrorKey] = trace.MsgRct.ExitCode.Error()
	}

	metadata[parser.MethodNumKey] = trace.Msg.Method.String()

	jsonMetadata, err := json.Marshal(metadata)
	if err != nil {
		_ = p.metrics.UpdateJsonMarshalMetric(parsermetrics.MetadataValue, txType)
	}

	p.appendAddressInfo(&parser.LotusMessage{
		To:     trace.Msg.To,
		From:   trace.Msg.From,
		Method: trace.Msg.Method,
		Cid:    mainMsgCid,
		Params: trace.Msg.Params,
	}, tipset.Key(), tipset.Height())

	var blockCid string
	if !systemExecution {
		blockCid, err = actorsV2.GetBlockCidFromMsgCid(mainMsgCid.String(), txType, metadata, tipset, p.logger)
		if err != nil {
			_ = p.metrics.UpdateBlockCidFromMsgCidMetric(txType)
			p.logger.Errorf("Error when trying to get block cid from message, txType '%s' cid '%s': %v", txType, mainMsgCid.String(), err)
		}
	}

	msgCid, err := tools.BuildCidFromMessageTrace(trace.Msg, mainMsgCid.String())
	if err != nil {
		_ = p.metrics.UpdateBuildCidFromMsgTraceMetric(txType)
		p.logger.Errorf("Error when trying to build message cid in tx cid'%s': %v", mainMsgCid.String(), err)
	}

	tipsetCid := tipset.GetCidString()
	messageUuid := tools.BuildMessageId(tipsetCid, blockCid, mainMsgCid.String(), msgCid, parentId)

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
		ParentId:      parentId,
		Id:            messageUuid,
		TxTimestamp:   parser.GetTimestamp(tipset.MinTimestamp()),
		TxCid:         mainMsgCid.String(),
		TxFrom:        txFrom,
		TxTo:          txTo,
		Amount:        trace.Msg.Value.Int,
		Status:        tools.GetExitCodeStatus(mainExitCode),
		SubcallStatus: tools.GetExitCodeStatus(trace.MsgRct.ExitCode),
		TxType:        txType,
		TxMetadata:    string(jsonMetadata),
	}, nil
}

func (p *Parser) feesTransactions(msg *typesV2.InvocResultV2, tipset *types.ExtendedTipSet, txType, parentTxId string, systemExecution bool) *types.Transaction {
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
		Id:            feeID,
		ParentId:      parentTxId,
		TxTimestamp:   timestamp,
		TxCid:         msg.MsgCid.String(),
		TxFrom:        msg.Msg.From.String(),
		TxTo:          parser.BurnAddress,
		Amount:        msg.GasCost.TotalCost.Int,
		Status:        tools.GetExitCodeStatus(exitcode.Ok),
		SubcallStatus: tools.GetExitCodeStatus(exitcode.Ok),
		TxType:        parser.TotalFeeOp,
		TxMetadata:    metadata,
	}
}

func (p *Parser) feesMetadata(msg *typesV2.InvocResultV2, tipset *types.ExtendedTipSet, txType, blockCid string, systemExecution bool) string {
	var minerAddress string
	var err error

	if !systemExecution && blockCid != "" {
		minerAddress, err = tipset.GetBlockMiner(blockCid)
		if err != nil {
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
			minerAddress = minerAddr.String()
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

func (p *Parser) appendAddressInfo(msg *parser.LotusMessage, key filTypes.TipSetKey, height abi.ChainEpoch) {
	if msg == nil {
		return
	}
	if msg.From != address.Undef {
		fromAdd := p.helper.GetActorAddressInfo(msg.From, key, height)
		parser.AppendToAddressesMap(p.addresses, fromAdd)
	}
	if msg.To != address.Undef {
		toAdd := p.helper.GetActorAddressInfo(msg.To, key, height)
		parser.AppendToAddressesMap(p.addresses, toAdd)
	}
}

func (p *Parser) getTxType(ctx context.Context, trace typesV2.ExecutionTraceV2, mainMsgCid cid.Cid, tipset *types.ExtendedTipSet) (actorName string, txType string, err error) {
	msg := &parser.LotusMessage{
		To:     trace.Msg.To,
		From:   trace.Msg.From,
		Method: trace.Msg.Method,
	}

	actorName, txType, err = p.getActorAndMethodName(ctx, trace, msg, mainMsgCid, tipset)
	if err != nil {
		p.logger.Errorf("Error when trying to get method name in tx cid'%s' using v2: %v", mainMsgCid.String(), err)
	}

	// fallback to depracated method
	if txType == parser.UnknownStr || txType == "" {
		//nolint:staticcheck // GetMethodName is deprecated, using v1 version for compatibility
		txType, err = p.helper.GetMethodName(msg, int64(tipset.Height()), tipset.Key())
		if err != nil {
			p.logger.Errorf("Error when trying to get method name in tx cid'%s' using v1: %v", mainMsgCid.String(), err)
			txType = parser.UnknownStr
		}
	}

	return actorName, txType, nil
}

func (p *Parser) getActorAndMethodName(ctx context.Context, trace typesV2.ExecutionTraceV2, msg *parser.LotusMessage, mainMsgCid cid.Cid, tipset *types.ExtendedTipSet) (actorName string, txType string, err error) {
	actorAddress := msg.To

	_, actorName, err = p.helper.GetActorInfoFromAddress(actorAddress, int64(tipset.Height()), tipset.Key())
	if err != nil || actorName == "" {
		p.logger.Warnf("Error when trying to get actor name in tx cid'%s': %v", mainMsgCid.String(), err)
		if trace.InvokedActor != nil {
			actorName, err = p.helper.GetActorNameFromCid(trace.InvokedActor.State.Code, int64(tipset.Height()))
			if err != nil {
				p.logger.Errorf("Error when trying to get actor name from cid in tx cid'%s' using invoked actor: %v", mainMsgCid.String(), err)
			}
		}
	}

	txType, err = actorsV2.GetMethodName(ctx, trace.Msg.Method, actorName, int64(tipset.Height()), p.network, p.helper, p.logger)
	if err != nil {
		txType = parser.UnknownStr
	}

	return actorName, txType, err
}
