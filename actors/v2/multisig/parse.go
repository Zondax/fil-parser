package multisig

import (
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
)

type Msig struct{}

func (p *Msig) Name() string {
	return manifest.MultisigKey
}

/*
Still needs to parse:

	Receive
*/
func (p *Msig) Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, _ cid.Cid, key filTypes.TipSetKey) (map[string]interface{}, *types.AddressInfo, error) {
	return map[string]interface{}{}, nil, parser.ErrUnknownMethod
}

func (p *Msig) TransactionTypes() map[string]any {
	return map[string]any{
		// parser.MethodSend: p.Send,
	}
}
