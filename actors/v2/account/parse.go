package account

import (
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
)

type Account struct{}

func (a *Account) Name() string {
	return manifest.AccountKey
}

func (a *Account) Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, _ cid.Cid, _ filTypes.TipSetKey) (map[string]interface{}, *types.AddressInfo, error) {
	switch txType {
	case parser.MethodSend:
		// return a.parseSend(msg), nil
	case parser.MethodConstructor:
		// return p.parseConstructor(msg.Params)
	case parser.MethodPubkeyAddress:
		resp, err := a.PubkeyAddress(network, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodAuthenticateMessage:
		resp, err := a.AuthenticateMessage(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.UnknownStr:
		// return p.unknownMetadata(msg.Params, msgRct.Return)
	}
	return map[string]interface{}{}, nil, parser.ErrUnknownMethod
}

func (a *Account) TransactionTypes() map[string]any {
	return map[string]any{
		parser.MethodSend:                  nil,
		parser.MethodConstructor:           nil,
		parser.MethodPubkeyAddress:         a.PubkeyAddress,
		parser.MethodAuthenticateMessage:   a.AuthenticateMessage,
		parser.MethodUniversalReceiverHook: nil,
	}
}
