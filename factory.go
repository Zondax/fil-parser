package fil_parser

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/zondax/fil-parser/actors"
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
	GetConfig() parser.Config
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
	postGenesisActors := parser.MainnetPostGenesisActors
	if p.network == tools.CalibrationNetwork {
		postGenesisActors = parser.CalibrationPostGenesisActors
	}

	genesisTxs := make([]*types.Transaction, 0)
	addresses := types.NewAddressInfoMap()
	genesisTimestamp := parser.GetTimestamp(genesisTipset.MinTimestamp())

	for _, actorInfo := range postGenesisActors {
		decKey, err := base64.StdEncoding.DecodeString(actorInfo[1])
		if err != nil {
			p.logger.Errorf("Error while decoding tipsetKey: %s. err: %s", actorInfo[1], err)
			continue
		}

		tipsetKey, err := types2.TipSetKeyFromBytes(decKey)
		if err != nil {
			p.logger.Errorf("genesis could not get tipset key: %s. err: %s", actorInfo[1], err)
			continue
		}

		addressInfo, err := getGenesisAddressInfo(actorInfo[0], tipsetKey, p.Helper)
		if err != nil {
			p.logger.Errorf("genesis could not get address info: %s. err: %s", actorInfo[0], err)
		} else {
			parser.AppendToAddressesMap(addresses, addressInfo)
		}
	}

	for _, balance := range genesis.Actors.All {
		addressInfo, err := getGenesisAddressInfo(balance.Key, genesisTipset.Key(), p.Helper)
		if err != nil {
			p.logger.Errorf("genesis could not get address info: %s. err: %s", balance.Key, err)
		} else {
			parser.AppendToAddressesMap(addresses, addressInfo)
		}

		if balance.Value.Balance == "0" {
			continue
		}

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
			TxFrom:      parser.TxFromGenesis,
			Amount:      amount.Int,
			Status:      "Ok",
			TxCid:       tipsetCid,
			TxType:      txType,
			TxMetadata:  "{}",
		})

	}

	return genesisTxs, addresses
}

func (p *FilecoinParser) ParseGenesisMultisig(ctx context.Context, genesis *types.GenesisBalances, genesisTipset *types.ExtendedTipSet) ([]*types.MultisigInfo, error) {
	var multisigInfos []*types.MultisigInfo

	for _, actor := range genesis.Actors.All {
		addressInfo, err := getGenesisAddressInfo(actor.Key, genesisTipset.Key(), p.Helper)
		if err != nil {
			p.logger.Errorf("multisig genesis could not get address info: %s. err: %s", actor.Key, err)
			continue
		}
		// actorName already parsed in getAddressInfo
		actorName := addressInfo.ActorType

		// check if the address is a multisig address
		if !strings.Contains(actorName, manifest.MultisigKey) {
			continue
		}
		addr, _ := address.NewFromString(actor.Key)

		api := p.Helper.GetFilecoinNodeClient()
		metadata, err := multisigTools.GenerateGenesisMultisigData(ctx, api, addr, genesisTipset)
		if err != nil {
			return nil, fmt.Errorf("multisigTools.GenerateGenesisMultisigData(%s): %s", actor.Key, err)
		}

		metadataJson, err := json.Marshal(metadata)
		if err != nil {
			return nil, fmt.Errorf("json.Marshal(): %s", err)
		}

		tipsetCid := genesisTipset.GetCidString()
		multisigInfo := &types.MultisigInfo{
			ID:              tools.BuildId(genesisTipset.GetCidString(), actor.Key, fmt.Sprint(parser.GenesisHeight), "", parser.TxTypeGenesis),
			MultisigAddress: actor.Key,
			Height:          parser.GenesisHeight,
			ActionType:      parser.MultisigConstructorMethod,
			Value:           string(metadataJson),
			Signer:          parser.TxFromGenesis,
			TxCid:           tipsetCid,
		}
		multisigInfos = append(multisigInfos, multisigInfo)

	}
	return multisigInfos, nil
}

func (p *FilecoinParser) ParseBlocksInfo(ctx context.Context, trace []byte, metadata types.BlockMetadata, tipset *types.ExtendedTipSet) (*types.BlocksTimestamp, *types.AddressInfoMap, error) {
	addresses := types.NewAddressInfoMap()
	nodeFullVersion := parser.UnknownStr
	nodeMajorMinorVersion := parser.UnknownStr
	if metadata.NodeFullVersion != "" {
		nodeFullVersion = metadata.NodeFullVersion
	}
	if metadata.NodeMajorMinorVersion != "" {
		nodeMajorMinorVersion = metadata.NodeMajorMinorVersion
	}

	if len(tipset.Blocks()) == 0 {
		p.logger.Debugf("found a tipset with no blocks at height '%d'", tipset.Height())

		tipsetId := tools.BuildTipsetId(fmt.Sprintf("%d-%s", tipset.Height(), tipset.GetCidString()))

		return &types.BlocksTimestamp{
			TipsetBasicBlockData: types.TipsetBasicBlockData{
				BasicBlockData: types.BasicBlockData{
					// #nosec G115
					Height:    uint64(tipset.Height()),
					TipsetCid: tipset.GetCidString(),
				},
				BlocksCid: []string{},
				NodeInfo: types.NodeInfo{
					NodeFullVersion:       nodeFullVersion,
					NodeMajorMinorVersion: nodeMajorMinorVersion,
				},
			},
			Id:              tipsetId,
			ParentTipsetCid: "",
			Timestamp:       time.Unix(0, 0),
			BaseFee:         0,
			BlocksInfo:      "[]",
		}, addresses, nil
	}

	tipsetId := tools.BuildTipsetId(tipset.GetCidString())

	baseFee, err := p.GetBaseFee(trace, metadata, tipset)
	if err != nil {
		// p.metrics.UpdateProcessedBlockTotalMetricFailure(parsermetrics.ErrorTypeGetBaseFee, false)
		p.logger.Errorf("error getting base fee: %w", err)
	}

	minTs := tipset.Blocks()[0].Timestamp
	for _, bh := range tipset.Blocks()[1:] {
		if bh.Timestamp < minTs {
			minTs = bh.Timestamp
		}
	}

	blocksInfo := make([]types.BlockInfo, 0, len(tipset.Blocks()))
	consolidateAddrs := p.parserV2.GetConfig().ConsolidateRobustAddress
	bestEffort := p.parserV2.GetConfig().RobustAddressBestEffort

	for _, block := range tipset.Blocks() {
		minerAddr := block.Miner.String()
		if consolidateAddrs {
			consolidatedMinerAddr, err := actors.ConsolidateToRobustAddress(block.Miner, p.Helper, p.logger, bestEffort)
			if err != nil {
				p.logger.Errorf("error consolidating miner address: %s. err: %s", block.Miner.String(), err)
			}
			minerAddr = consolidatedMinerAddr
		}
		blocksInfo = append(blocksInfo, types.BlockInfo{
			BlockCid: block.Cid().String(),
			Miner:    minerAddr,
		})

		addressInfo := p.Helper.GetActorAddressInfo(block.Miner, tipset.Key(), block.Height)
		parser.AppendToAddressesMap(addresses, addressInfo)
	}
	blocksBlob, _ := json.Marshal(blocksInfo)

	blockTimeStamp := int64(minTs) * 1000 //nolint:gosec,G115 // Allowing integer overflow conversion

	blocksCid := tools.GetBlocksCidByString(tipset.Key().String())
	return &types.BlocksTimestamp{
		TipsetBasicBlockData: types.TipsetBasicBlockData{
			BasicBlockData: types.BasicBlockData{
				// #nosec G115
				Height:    uint64(tipset.Height()),
				TipsetCid: tipset.GetCidString(),
			},
			BlocksCid: blocksCid,
			NodeInfo: types.NodeInfo{
				NodeFullVersion:       nodeFullVersion,
				NodeMajorMinorVersion: nodeMajorMinorVersion,
			},
		},
		Id:              tipsetId,
		ParentTipsetCid: tipset.GetParentCidString(),
		Timestamp:       time.Unix(blockTimeStamp/1000, blockTimeStamp%1000),
		BaseFee:         baseFee,
		BlocksInfo:      string(blocksBlob),
	}, addresses, nil

}

func getGenesisAddressInfo(addrStr string, tipsetKey types2.TipSetKey, helper *helper2.Helper) (*types.AddressInfo, error) {
	filAdd, err := address.NewFromString(addrStr)
	if err != nil {
		return nil, fmt.Errorf("could not parse address: %s. err: %s", addrStr, err)
	}

	shortAdd, err := helper.GetActorsCache().GetShortAddress(filAdd)
	if err != nil {
		return nil, fmt.Errorf("could not get short address: %s. err: %s", addrStr, err)
	}
	robustAdd, err := helper.GetActorsCache().GetRobustAddress(filAdd)
	if err != nil {
		return nil, fmt.Errorf("could not get robust address: %s. err: %s", addrStr, err)
	}
	actorCode, err := helper.GetActorsCache().GetActorCode(filAdd, tipsetKey, false)
	if err != nil {
		return nil, fmt.Errorf("could not get actor code: %s. err: %s", addrStr, err)
	}
	_, actorName, err := helper.GetActorNameFromAddress(filAdd, 0, tipsetKey)
	if err != nil {
		return nil, fmt.Errorf("could not get actor name: %s. err: %s", addrStr, err)
	}

	return &types.AddressInfo{
		Short:    shortAdd,
		Robust:   robustAdd,
		ActorCid: actorCode,
		// genesis transactions do not have a creation_tx_cid ,
		// we use the tipset_cid in this case to enable users to find the genesis tipset from this address info.
		CreationTxCid: tipsetKey.String(),
		ActorType:     tools.ParseActorName(actorName),
		IsSystemActor: helper.IsSystemActor(filAdd) || helper.IsGenesisActor(filAdd),
	}, nil
}
