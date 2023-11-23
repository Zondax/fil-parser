package fil_parser

import (
	"errors"
	"fmt"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/big"
	types2 "github.com/filecoin-project/lotus/chain/types"
	"github.com/google/uuid"
	"github.com/zondax/fil-parser/actors/cache"
	"github.com/zondax/fil-parser/actors/cache/impl/common"
	logger2 "github.com/zondax/fil-parser/logger"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/parser/V23"
	helper2 "github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
	rosettaFilecoinLib "github.com/zondax/rosetta-filecoin-lib"
	"go.uber.org/zap"
	"strings"
)

var (
	errUnknownImpl    = errors.New("unknown implementation")
	errUnknownVersion = errors.New("unknown trace version")
)

type FilecoinParser struct {
	parserv1  Parser
	parserV23 Parser
	Helper    *helper2.Helper
	logger    *zap.Logger
}

type Parser interface {
	Version() string
	ParseTransactions(traces []byte, tipSet *types.ExtendedTipSet, ethLogs []types.EthLog) ([]*types.Transaction, *types.AddressInfoMap, error)
	GetBaseFee(traces []byte) (uint64, error)
	IsVersionCompatible(ver string) bool
}

func NewFilecoinParser(lib *rosettaFilecoinLib.RosettaConstructionFilecoin, cacheSource common.DataSource, logger *zap.Logger) (*FilecoinParser, error) {
	logger = logger2.GetSafeLogger(logger)
	actorsCache, err := cache.SetupActorsCache(cacheSource, logger)
	if err != nil {
		logger.Sugar().Errorf("could not setup actors cache: %v", err)
		return nil, err
	}

	helper := helper2.NewHelper(lib, actorsCache, logger)
	parserv1 := v1.NewParserv1(helper, logger)
	parserV23 := V23.NewParserV23(helper, logger)

	return &FilecoinParser{
		parserv1:  parserv1,
		parserV23: parserV23,
		Helper:    helper,
		logger:    logger,
	}, nil
}

func (p *FilecoinParser) ParseTransactions(traces []byte, tipSet *types.ExtendedTipSet, ethLogs []types.EthLog, metadata *types.BlockMetadata) ([]*types.Transaction, *types.AddressInfoMap, error) {
	parserVersion, err := p.translateParserVersionFromMetadata(*metadata)
	if err != nil {
		return nil, nil, errUnknownVersion
	}

	var txs []*types.Transaction
	var addrs *types.AddressInfoMap

	p.logger.Sugar().Debugf("node version found on trace files %s to parse transactions", parserVersion)
	switch parserVersion {
	case parser.ParserV1:
		txs, addrs, err = p.parserv1.ParseTransactions(traces, tipSet, ethLogs)
	case parser.ParserV2:
		txs, addrs, err = p.parserV23.ParseTransactions(traces, tipSet, ethLogs)
	default:
		p.logger.Sugar().Errorf("[parser] implementation not supported: %s", parserVersion)
		return nil, nil, errUnknownImpl
	}

	if err != nil {
		return nil, nil, err
	}

	return p.FilterDuplicated(txs), addrs, nil
}

func (p *FilecoinParser) translateParserVersionFromMetadata(metadata types.BlockMetadata) (string, error) {
	switch {
	// The empty string is for backwards compatibility with older traces versions
	case p.parserv1.IsVersionCompatible(metadata.NodeMajorMinorVersion), metadata.NodeMajorMinorVersion == "":
		return parser.ParserV1, nil
	case p.parserV23.IsVersionCompatible(metadata.NodeMajorMinorVersion):
		return parser.ParserV2, nil
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

func (p *FilecoinParser) GetBaseFee(traces []byte, metadata types.BlockMetadata) (uint64, error) {
	parserVersion, err := p.translateParserVersionFromMetadata(metadata)
	if err != nil {
		return 0, errUnknownVersion
	}

	p.logger.Sugar().Debugf("node version found on trace files %s to get base fee", parserVersion)
	switch parserVersion {
	case parser.ParserV1:
		return p.parserv1.GetBaseFee(traces)
	case parser.ParserV2:
		return p.parserV23.GetBaseFee(traces)
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

		tipsetCid := genesisTipset.Key().String()
		tipsetCid = strings.ReplaceAll(tipsetCid, "{", "")
		tipsetCid = strings.ReplaceAll(tipsetCid, "}", "")

		genesisTxs = append(genesisTxs, &types.Transaction{
			TxBasicBlockData: types.TxBasicBlockData{
				BasicBlockData: types.BasicBlockData{
					Height:    0,
					TipsetCid: tipsetCid,
				},
			},
			Id:          tools.BuildId(genesisTipset.Key().String(), balance.Key, balance.Value.Balance),
			ParentId:    uuid.Nil.String(),
			Level:       0,
			TxTimestamp: genesisTimestamp,
			TxTo:        balance.Key,
			Amount:      amount.Int,
			Status:      "Ok",
			TxType:      "Genesis",
		})
	}

	return genesisTxs, addresses
}
