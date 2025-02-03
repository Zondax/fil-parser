package multisig

import (
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/zondax/fil-parser/parser"
)

type Msig struct{}

func (p *Msig) Name() string {
	return manifest.MultisigKey
}

/*
Still needs to parse:

	Receive
*/
func (p *Msig) Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, key filTypes.TipSetKey) (map[string]interface{}, error) {
	return map[string]interface{}{}, parser.ErrUnknownMethod
}
