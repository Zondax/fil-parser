package fil_parser

import (
	"errors"
	"fmt"
	"strings"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/big"
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
	"github.com/zondax/fil-parser/types"
	rosettaFilecoinLib "github.com/zondax/rosetta-filecoin-lib"
	"go.uber.org/zap"
)

var (
	errUnknownImpl    = errors.New("unknown implementation")
	errUnknownVersion = errors.New("unknown trace version")
)

type FilecoinParser struct {
	parserV1 Parser
	parserV2 Parser
	Helper   *helper2.Helper
	logger   *zap.Logger
}

type Parser interface {
	Version() string
	NodeVersionsSupported() []string
	ParseTransactions(traces []byte, tipSet *types.ExtendedTipSet, ethLogs []types.EthLog, metadata types.BlockMetadata) ([]*types.Transaction, *types.AddressInfoMap, []types.TxCidTranslation, error)
	GetBaseFee(traces []byte, tipset *types.ExtendedTipSet) (uint64, error)
	IsNodeVersionSupported(ver string) bool
}

func NewFilecoinParser(lib *rosettaFilecoinLib.RosettaConstructionFilecoin, cacheSource common.DataSource, logger *zap.Logger) (*FilecoinParser, error) {
	logger = logger2.GetSafeLogger(logger)
	actorsCache, err := cache.SetupActorsCache(cacheSource, logger)
	if err != nil {
		logger.Sugar().Errorf("could not setup actors cache: %v", err)
		return nil, err
	}

	helper := helper2.NewHelper(lib, actorsCache, cacheSource.Node, logger)
	parserV1 := v1.NewParser(helper, logger)
	parserV2 := v2.NewParser(helper, logger)

	return &FilecoinParser{
		parserV1: parserV1,
		parserV2: parserV2,
		Helper:   helper,
		logger:   logger,
	}, nil
}

func (p *FilecoinParser) ParseTransactions(traces []byte, tipSet *types.ExtendedTipSet, ethLogs []types.EthLog, metadata types.BlockMetadata) ([]*types.Transaction, *types.AddressInfoMap, []types.TxCidTranslation, error) {
	parserVersion, err := p.translateParserVersionFromMetadata(metadata)
	if err != nil {
		return nil, nil, nil, errUnknownVersion
	}

	var txs []*types.Transaction
	var addrs *types.AddressInfoMap
	var txsCid []types.TxCidTranslation

	p.logger.Sugar().Debugf("trace files node version: [%s] - parser to use: [%s]", metadata.NodeMajorMinorVersion, parserVersion)
	switch parserVersion {
	case v1.Version:
		txs, addrs, txsCid, err = p.parserV1.ParseTransactions(traces, tipSet, ethLogs, metadata)
	case v2.Version:
		txs, addrs, txsCid, err = p.parserV2.ParseTransactions(traces, tipSet, ethLogs, metadata)
	default:
		p.logger.Sugar().Errorf("[parser] implementation not supported: %s", parserVersion)
		return nil, nil, nil, errUnknownImpl
	}

	if err != nil {
		return nil, nil, nil, err
	}

	return p.FilterDuplicated(txs), addrs, txsCid, nil
}

func (p *FilecoinParser) translateParserVersionFromMetadata(metadata types.BlockMetadata) (string, error) {
	switch {
	// The empty string is for backwards compatibility with older traces versions
	case p.parserV1.IsNodeVersionSupported(metadata.NodeMajorMinorVersion), metadata.NodeMajorMinorVersion == "":
		return v1.Version, nil
	case p.parserV2.IsNodeVersionSupported(metadata.NodeMajorMinorVersion):
		return v2.Version, nil
	default:
		p.logger.Sugar().Errorf("[parser] unsupported node version: %s", metadata.NodeFullVersion)
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

	p.logger.Sugar().Debugf("trace files node version: [%s] - parser to use: [%s]", metadata.NodeMajorMinorVersion, parserVersion)
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
		})
	}

	return genesisTxs, addresses
}
