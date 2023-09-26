package fil_parser

import (
	"errors"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/big"
	types2 "github.com/filecoin-project/lotus/chain/types"
	"github.com/google/uuid"
	"github.com/zondax/fil-parser/actors/cache"
	"github.com/zondax/fil-parser/actors/cache/impl/common"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/parser/V22"
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
	parserV22 Parser
	parserV23 Parser
	Helper    *helper2.Helper
}

type Parser interface {
	Version() string
	ParseTransactions(traces []byte, tipSet *types.ExtendedTipSet, ethLogs []types.EthLog) ([]*types.Transaction, *types.AddressInfoMap, error)
	GetBaseFee(traces []byte) (uint64, error)
}

func NewFilecoinParser(lib *rosettaFilecoinLib.RosettaConstructionFilecoin, cacheSource common.DataSource) (*FilecoinParser, error) {
	actorsCache, err := cache.SetupActorsCache(cacheSource)
	if err != nil {
		zap.S().Errorf("could not setup actors cache: %v", err)
		return nil, err
	}

	helper := helper2.NewHelper(lib, actorsCache)
	parserV22 := V22.NewParserV22(helper)
	parserV23 := V23.NewParserV23(helper)

	return &FilecoinParser{
		parserV22: parserV22,
		parserV23: parserV23,
		Helper:    helper,
	}, nil
}

func (p *FilecoinParser) ParseTransactions(traces []byte, tipSet *types.ExtendedTipSet, ethLogs []types.EthLog, metadata *types.BlockMetadata) ([]*types.Transaction, *types.AddressInfoMap, error) {
	version := detectTraceVersion(*metadata)
	if version == "" {
		return nil, nil, errUnknownVersion
	}

	var txs []*types.Transaction
	var addrs *types.AddressInfoMap
	var err error

	switch version {
	case V22.Version:
		txs, addrs, err = p.parserV22.ParseTransactions(traces, tipSet, ethLogs)
	case V23.Version:
		txs, addrs, err = p.parserV23.ParseTransactions(traces, tipSet, ethLogs)
	default:
		zap.S().Errorf("[parser] implementation not supported: %s", version)
		return nil, nil, errUnknownImpl
	}

	if err != nil {
		return nil, nil, err
	}

	return p.FilterDuplicated(txs), addrs, nil
}

func detectTraceVersion(metadata types.BlockMetadata) string {
	switch metadata.NodeMajorMinorVersion {
	case V22.Version, "": // The empty string is for backwards compatibility with older traces versions
		return V22.Version
	case V23.Version:
		return V23.Version
	default:
		zap.S().Errorf("[parser] unsupported node version: %s", metadata.NodeFullVersion)
		return ""
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
	version := detectTraceVersion(metadata)
	if version == "" {
		return 0, errUnknownVersion
	}

	switch version {
	case V22.Version:
		return p.parserV22.GetBaseFee(traces)
	case V23.Version:
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
		actorCode, _ := p.Helper.GetActorsCache().GetActorCode(filAdd, types2.EmptyTSK)
		actorName := p.Helper.GetActorNameFromAddress(filAdd, 0, types2.EmptyTSK)

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
