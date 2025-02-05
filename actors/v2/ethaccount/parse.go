package ethaccount

import (
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
)

type EthAccount struct{}

func (e *EthAccount) Name() string {
	return manifest.EthAccountKey
}

func (e *EthAccount) Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, mainMsgCid cid.Cid, key filTypes.TipSetKey) (map[string]interface{}, *types.AddressInfo, error) {
	return map[string]interface{}{}, nil, parser.ErrUnknownMethod
}

func (e *EthAccount) TransactionTypes() map[string]any {
	return map[string]any{
		// parser.MethodSend: e.Send,
	}
}
