package parser

import (
	"github.com/zondax/fil-parser/parser/V23"

	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/zondax/fil-parser/parser/V22"
	"github.com/zondax/fil-parser/types"
	rosettaFilecoinLib "github.com/zondax/rosetta-filecoin-lib"
)

type IParser interface {
	ParseTransactions(traces any, tipSet *filTypes.TipSet, ethLogs []types.EthLog) ([]*types.Transaction, *types.AddressInfoMap, error)
}

type Parser struct {
	parserV22 IParser
	parserV23 IParser
}

func NewParser(lib *rosettaFilecoinLib.RosettaConstructionFilecoin) *Parser {
	return &Parser{
		parserV22: V22.NewParserV22(lib),
		parserV23: V23.NewParserV23(lib),
	}
}

func (p *Parser) ParseTransactions(traces any, tipSet *filTypes.TipSet, ethLogs []types.EthLog) ([]*types.Transaction, *types.AddressInfoMap, error) {
	// Parse traces according to its inner version
	// TODO: check how we could check the version of the traces. One idea is that every trace file would have a 'version' field
	// that we can read and then decide which parser to use
	switch version {
	case "v22":
		return p.parserV22.ParseTransactions(traces, tipSet, ethLogs)
	case "v23":
		return p.parserV23.ParseTransactions(traces, tipSet, ethLogs)
	}

	return nil, nil, ErrInvalidVersion
}
