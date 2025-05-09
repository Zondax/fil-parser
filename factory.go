package fil_parser

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/zondax/fil-parser/metrics"
	"github.com/zondax/golem/pkg/logger"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/go-state-types/manifest"
	types2 "github.com/filecoin-project/lotus/chain/types"
	"github.com/google/uuid"
	"github.com/zondax/fil-parser/actors/cache"
	"github.com/zondax/fil-parser/actors/cache/impl/common"
	logger2 "github.com/zondax/fil-parser/logger"
	"github.com/zondax/fil-parser/parser"
	helper2 "github.com/zondax/fil-parser/parser/helper"
	v1 "github.com/zondax/fil-parser/parser/v1"
	v2 "github.com/zondax/fil-parser/parser/v2"
	"github.com/zondax/fil-parser/tools"
	multisigTools "github.com/zondax/fil-parser/tools/multisig"
	"github.com/zondax/fil-parser/types"
	rosettaFilecoinLib "github.com/zondax/rosetta-filecoin-lib"
)

var (
	errUnknownImpl    = errors.New("unknown implementation")
	errUnknownVersion = errors.New("unknown trace version")
)

type FilecoinParser struct {
	parserV1 Parser
	parserV2 Parser
	Helper   *helper2.Helper
	logger   *logger.Logger
	network  string
}

type Parser interface {
	Version() string
	NodeVersionsSupported() []string
	ParseTransactions(ctx context.Context, txsData types.TxsData) (*types.TxsParsedResult, error)
	ParseNativeEvents(ctx context.Context, eventsData types.EventsData) (*types.EventsParsedResult, error)
	ParseMultisigEvents(ctx context.Context, multisigTxs []*types.Transaction, tipsetCid string, tipsetKey types2.TipSetKey) (*types.MultisigEvents, error)
	ParseMinerEvents(ctx context.Context, txs []*types.Transaction, tipsetCid string, tipsetKey types2.TipSetKey) (*types.MinerEvents, error)
	ParseEthLogs(ctx context.Context, eventsData types.EventsData) (*types.EventsParsedResult, error)
	GetBaseFee(traces []byte, tipset *types.ExtendedTipSet) (uint64, error)
	IsNodeVersionSupported(ver string) bool
}

func NewFilecoinParser(lib *rosettaFilecoinLib.RosettaConstructionFilecoin, cacheSource common.DataSource, logger *logger.Logger, opts ...Option) (*FilecoinParser, error) {
	defaultOpts := FilecoinParserOptions{
		metrics: metrics.NewNoopMetricsClient(),
		config: parser.Config{
			FeesAsColumn:                  false,
			ConsolidateRobustAddress:      false,
			RobustAddressBestEffort:       false,
			NodeMaxRetries:                3,
			NodeMaxWaitBeforeRetrySeconds: 1,
			NodeRetryStrategy:             "linear",
		},
	}
	for _, opt := range opts {
		opt(&defaultOpts)
	}

	logger = logger2.GetSafeLogger(logger)
	actorsCache, err := cache.SetupActorsCache(cacheSource, logger, defaultOpts.metrics, defaultOpts.backoff)
	if err != nil {
		logger.Errorf("could not setup actors cache: %v", err)
		return nil, err
	}

	helper := helper2.NewHelper(lib, actorsCache, cacheSource.Node, logger, defaultOpts.metrics)
	if helper == nil {
		return nil, errors.New("helper is nil")
	}

	network, err := helper.GetFilecoinNodeClient().StateNetworkName(context.Background())
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	parserV1 := v1.NewParser(helper, logger, defaultOpts.metrics, defaultOpts.config)
	parserV2 := v2.NewParser(helper, logger, defaultOpts.metrics, defaultOpts.config)

	return &FilecoinParser{
		parserV1: parserV1,
		parserV2: parserV2,
		Helper:   helper,
		logger:   logger,
		network:  tools.ParseRawNetworkName(string(network)),
	}, nil
}

func NewFilecoinParserWithActorV2(lib *rosettaFilecoinLib.RosettaConstructionFilecoin, cacheSource common.DataSource, logger *logger.Logger, opts ...Option) (*FilecoinParser, error) {
	defaultOpts := FilecoinParserOptions{
		metrics: metrics.NewNoopMetricsClient(),
		config: parser.Config{
			FeesAsColumn:                  false,
			ConsolidateRobustAddress:      false,
			RobustAddressBestEffort:       false,
			NodeMaxRetries:                3,
			NodeMaxWaitBeforeRetrySeconds: 1,
			NodeRetryStrategy:             "linear",
		},
	}
	for _, opt := range opts {
		opt(&defaultOpts)
	}

	logger = logger2.GetSafeLogger(logger)
	actorsCache, err := cache.SetupActorsCache(cacheSource, logger, defaultOpts.metrics, defaultOpts.backoff)
	if err != nil {
		logger.Errorf("could not setup actors cache: %v", err)
		return nil, err
	}

	helper := helper2.NewHelper(lib, actorsCache, cacheSource.Node, logger, defaultOpts.metrics)
	if helper == nil {
		return nil, errors.New("helper is nil")
	}
	network, err := helper.GetFilecoinNodeClient().StateNetworkName(context.Background())
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	networkName := tools.ParseRawNetworkName(string(network))

	var parserV1 Parser
	var parserV2 Parser

	parserV1 = v1.NewActorsV2Parser(networkName, helper, logger, defaultOpts.metrics, defaultOpts.config)
	if networkName == tools.CalibrationNetwork {
		// trace files already use executiontracev2 because of a resync and calibration resets
		// so we need to use the new parser regardless of the height
		parserV1 = v2.NewActorsV2Parser(networkName, helper, logger, defaultOpts.metrics, defaultOpts.config)
	}

	parserV2 = v2.NewActorsV2Parser(networkName, helper, logger, defaultOpts.metrics, defaultOpts.config)

	return &FilecoinParser{
		parserV1: parserV1,
		parserV2: parserV2,
		Helper:   helper,
		logger:   logger,
		network:  networkName,
	}, nil
}

func (p *FilecoinParser) ParseTransactions(ctx context.Context, txsData types.TxsData) (*types.TxsParsedResult, error) {
	parserVersion, err := p.translateParserVersionFromMetadata(txsData.Metadata)
	if err != nil {
		return nil, errUnknownVersion
	}

	var parsedResult *types.TxsParsedResult

	p.logger.Debugf("trace files node version: [%s] - parser to use: [%s]", txsData.Metadata.NodeMajorMinorVersion, parserVersion)
	switch parserVersion {
	case v1.Version:
		parsedResult, err = p.parserV1.ParseTransactions(ctx, txsData)
	case v2.Version:
		parsedResult, err = p.parserV2.ParseTransactions(ctx, txsData)
	default:
		p.logger.Errorf("[parser] implementation not supported: %s", parserVersion)
		return nil, errUnknownImpl
	}

	if err != nil {
		return nil, err
	}

	parsedResult.Txs = p.FilterDuplicated(parsedResult.Txs)

	return parsedResult, nil
}

func (p *FilecoinParser) ParseNativeEvents(ctx context.Context, eventsData types.EventsData) (*types.EventsParsedResult, error) {
	parserVersion, err := p.translateParserVersionFromMetadata(eventsData.Metadata)
	if err != nil {
		return nil, errUnknownVersion
	}

	var parsedResult *types.EventsParsedResult

	p.logger.Debugf("trace files node version: [%s] - parser to use: [%s]", eventsData.Metadata.NodeMajorMinorVersion, parserVersion)
	switch parserVersion {
	case v1.Version, v2.Version:
		parsedResult, err = p.parserV2.ParseNativeEvents(ctx, eventsData)
	default:
		p.logger.Errorf("[parser] implementation not supported: %s", parserVersion)
		return nil, errUnknownImpl
	}

	if err != nil {
		return nil, err
	}

	return parsedResult, nil
}

func (p *FilecoinParser) ParseEthLogs(ctx context.Context, eventsData types.EventsData) (*types.EventsParsedResult, error) {
	parserVersion, err := p.translateParserVersionFromMetadata(eventsData.Metadata)
	if err != nil {
		return nil, errUnknownVersion
	}

	var parsedResult *types.EventsParsedResult

	p.logger.Debugf("trace files node version: [%s] - parser to use: [%s]", eventsData.Metadata.NodeMajorMinorVersion, parserVersion)
	switch parserVersion {
	case v1.Version, v2.Version:
		parsedResult, err = p.parserV2.ParseEthLogs(ctx, eventsData)
	default:
		p.logger.Errorf("[parser] implementation not supported: %s", parserVersion)
		return nil, errUnknownImpl
	}

	if err != nil {
		return nil, err
	}

	return parsedResult, nil
}

func (p *FilecoinParser) ParseMultisigEvents(ctx context.Context, txs []*types.Transaction, tipsetCid string, tipsetKey types2.TipSetKey) (*types.MultisigEvents, error) {
	multisigTxs, err := p.Helper.FilterTxsByActorType(ctx, txs, manifest.MultisigKey, tipsetKey)
	if err != nil {
		return nil, err
	}
	return p.parserV2.ParseMultisigEvents(ctx, multisigTxs, tipsetCid, tipsetKey)
}

func (p *FilecoinParser) ParseMinerEvents(ctx context.Context, txs []*types.Transaction, tipsetCid string, tipsetKey types2.TipSetKey) (*types.MinerEvents, error) {
	return p.parserV2.ParseMinerEvents(ctx, txs, tipsetCid, tipsetKey)
}

func (p *FilecoinParser) translateParserVersionFromMetadata(metadata types.BlockMetadata) (string, error) {
	switch {
	// The empty string is for backwards compatibility with older traces versions
	case p.parserV1.IsNodeVersionSupported(metadata.NodeMajorMinorVersion), metadata.NodeMajorMinorVersion == "":
		return v1.Version, nil
	case p.parserV2.IsNodeVersionSupported(metadata.NodeMajorMinorVersion):
		return v2.Version, nil
	default:
		p.logger.Errorf("[parser] unsupported node version: %s", metadata.NodeFullVersion)
		return "", fmt.Errorf("node version not supported %s", metadata.NodeFullVersion)
	}
}

func (p *FilecoinParser) FilterDuplicated(txs []*types.Transaction) []*types.Transaction {
	idsFound := make(map[string]bool)
	filteredTxs := make([]*types.Transaction, 0)

	for _, tx := range txs {
		if _, found := idsFound[tx.Id]; !found {
			idsFound[tx.Id] = true
			filteredTxs = append(filteredTxs, tx)
		}
	}

	return filteredTxs
}

func (p *FilecoinParser) GetBaseFee(traces []byte, metadata types.BlockMetadata, tipset *types.ExtendedTipSet) (uint64, error) {
	parserVersion, err := p.translateParserVersionFromMetadata(metadata)
	if err != nil {
		return 0, errUnknownVersion
	}

	p.logger.Debugf("trace files node version: [%s] - parser to use: [%s]", metadata.NodeMajorMinorVersion, parserVersion)
	switch parserVersion {
	case v1.Version:
		return p.parserV1.GetBaseFee(traces, tipset)
	case v2.Version:
		return p.parserV2.GetBaseFee(traces, tipset)
	}

	return 0, errUnknownImpl
}

func (p *FilecoinParser) ParseGenesis(genesis *types.GenesisBalances, genesisTipset *types.ExtendedTipSet) ([]*types.Transaction, *types.AddressInfoMap) {
	genesisTxs := make([]*types.Transaction, 0)
	addresses := types.NewAddressInfoMap()
	genesisTimestamp := parser.GetTimestamp(genesisTipset.MinTimestamp())

	for _, balance := range genesis.Actors.All {
		if balance.Value.Balance == "0" {
			continue
		}

		filAdd, _ := address.NewFromString(balance.Key)
		shortAdd, _ := p.Helper.GetActorsCache().GetShortAddress(filAdd)
		robustAdd, _ := p.Helper.GetActorsCache().GetRobustAddress(filAdd)
		actorCode, _ := p.Helper.GetActorsCache().GetActorCode(filAdd, types2.EmptyTSK, false)
		actorName, _ := p.Helper.GetActorNameFromAddress(filAdd, 0, types2.EmptyTSK)

		addresses.Set(balance.Key, &types.AddressInfo{
			Short:     shortAdd,
			Robust:    robustAdd,
			ActorCid:  actorCode,
			ActorType: actorName,
		})
		amount, _ := big.FromString(balance.Value.Balance)

		tipsetCid := genesisTipset.GetCidString()
		txType := "Genesis"
		blockCid := genesisTipset.Key().String()
		blockCid = strings.ReplaceAll(blockCid, "{", "")
		blockCid = strings.ReplaceAll(blockCid, "}", "")
		genesisTxs = append(genesisTxs, &types.Transaction{
			TxBasicBlockData: types.TxBasicBlockData{
				BasicBlockData: types.BasicBlockData{
					Height:    0,
					TipsetCid: tipsetCid,
				},
				BlockCid: blockCid,
			},
			Id:          tools.BuildId(genesisTipset.Key().String(), balance.Key, balance.Value.Balance),
			ParentId:    uuid.Nil.String(),
			Level:       0,
			TxTimestamp: genesisTimestamp,
			TxTo:        balance.Key,
			Amount:      amount.Int,
			Status:      "Ok",
			TxType:      txType,
			TxMetadata:  "{}",
		})
	}

	return genesisTxs, addresses
}

func (p *FilecoinParser) ParseGenesisMultisig(ctx context.Context, genesis *types.GenesisBalances, genesisTipset *types.ExtendedTipSet) ([]*types.MultisigInfo, error) {
	var multisigInfos []*types.MultisigInfo
	for _, actor := range genesis.Actors.All {
		addrStr := actor.Key
		// parse address
		addr, err := address.NewFromString(addrStr)
		if err != nil {
			p.logger.Errorf("could not parse address: %s. err: %s", addrStr, err)
			continue
		}

		// get actor name from address
		actorName, err := p.Helper.GetActorNameFromAddress(addr, int64(parser.GenesisHeight), genesisTipset.Key())
		if err != nil {
			p.logger.Errorf("could not get actor name from address: %s. err: %s", addrStr, err)
			continue
		}

		// check if the address is a multisig address
		if !strings.Contains(actorName, manifest.MultisigKey) {
			continue
		}

		api := p.Helper.GetFilecoinNodeClient()
		metadata, err := multisigTools.GenerateGenesisMultisigData(ctx, api, addr, genesisTipset)
		if err != nil {
			return nil, fmt.Errorf("multisigTools.GenerateGenesisMultisigData(%s): %s", addrStr, err)
		}

		metadataJson, err := json.Marshal(metadata)
		if err != nil {
			return nil, fmt.Errorf("json.Marshal(): %s", err)
		}

		multisigInfo := &types.MultisigInfo{
			ID:              tools.BuildId(genesisTipset.GetCidString(), addrStr, fmt.Sprint(parser.GenesisHeight), "", parser.TxTypeGenesis),
			MultisigAddress: addrStr,
			Height:          parser.GenesisHeight,
			ActionType:      parser.MultisigConstructorMethod,
			Value:           string(metadataJson),

			// there is no signer as this is genesis
			Signer: "",
			// there are no transactions for the multisig addresses in the genesis block
			TxCid: "",
		}
		multisigInfos = append(multisigInfos, multisigInfo)

	}
	return multisigInfos, nil
}
