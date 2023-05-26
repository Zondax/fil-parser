package fil_parser

import (
	"errors"
	"github.com/zondax/fil-parser/parser/V22"
	"github.com/zondax/fil-parser/parser/V23"

	filTypes "github.com/filecoin-project/lotus/chain/types"
	rosettaFilecoinLib "github.com/zondax/rosetta-filecoin-lib"

	"github.com/zondax/fil-parser/types"
)

type IParser interface {
	Version() string
	ParseTransactions(traces any, tipSet *filTypes.TipSet, ethLogs []types.EthLog) ([]*types.Transaction, types.AddressInfoMap, error)
}

func NewParser(lib *rosettaFilecoinLib.RosettaConstructionFilecoin, version string) (IParser, error) {
	switch version {
	case "v22":
		return V22.NewParserV22(lib), nil
	case "v23":
		return V23.NewParserV23(lib), nil
	}
	return nil, errors.New("unknown implementation")
}
