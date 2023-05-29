package fil_parser

import (
	"errors"
	"github.com/zondax/fil-parser/parser/V22"
	"github.com/zondax/fil-parser/parser/V23"

	filTypes "github.com/filecoin-project/lotus/chain/types"
	rosettaFilecoinLib "github.com/zondax/rosetta-filecoin-lib"

	"github.com/zondax/fil-parser/types"
)

var (
	errUnknownImpl = errors.New("unknown implementation")
)

type Parser interface {
	Version() string
	ParseTransactions(traces []byte, tipSet *filTypes.TipSet, ethLogs []types.EthLog) ([]*types.Transaction, types.AddressInfoMap, error)
}

func NewParser(lib *rosettaFilecoinLib.RosettaConstructionFilecoin, version string) (Parser, error) {
	switch version {
	case V22.Version:
		return V22.NewParserV22(lib), nil
	case V23.Version:
		return V23.NewParserV23(lib), nil
	}
	return nil, errUnknownImpl
}
