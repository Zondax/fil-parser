package fil_parser

import (
	"errors"
	"github.com/zondax/fil-parser/actors/cache"
	"github.com/zondax/fil-parser/actors/cache/impl/common"
	"github.com/zondax/fil-parser/parser/V22"
	"github.com/zondax/fil-parser/parser/V23"
	helper2 "github.com/zondax/fil-parser/parser/helper"
	"go.uber.org/zap"

	rosettaFilecoinLib "github.com/zondax/rosetta-filecoin-lib"

	"github.com/zondax/fil-parser/types"
)

var (
	errUnknownImpl = errors.New("unknown implementation")
)

type FilecoinParser struct {
	ParserV22 Parser
	ParserV23 Parser
}

type Parser interface {
	Version() string
	ParseTransactions(traces []byte, tipSet *types.ExtendedTipSet, ethLogs []types.EthLog) ([]*types.Transaction, types.AddressInfoMap, error)
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
		ParserV22: parserV22,
		ParserV23: parserV23,
	}, nil
}

func (p *FilecoinParser) ParseTransactions(traces []byte, tipSet *types.ExtendedTipSet, ethLogs []types.EthLog, version string) ([]*types.Transaction, types.AddressInfoMap, error) {
	switch version {
	case V22.Version:
		return p.ParserV22.ParseTransactions(traces, tipSet, ethLogs)
	case V23.Version:
		return p.ParserV23.ParseTransactions(traces, tipSet, ethLogs)
	}
	return nil, nil, errUnknownImpl
}
