package account

import (
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/zondax/fil-parser/parser"
)

type Account struct{}

func (a *Account) Name() string {
	return manifest.AccountKey
}

func (a *Account) Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt) (map[string]interface{}, error) {
	switch txType {
	case parser.MethodSend:
		// return a.parseSend(msg), nil
	case parser.MethodConstructor:
		// return p.parseConstructor(msg.Params)
	case parser.MethodPubkeyAddress:
		return a.PubkeyAddress(network, msg.Params, msgRct.Return)
	case parser.MethodAuthenticateMessage:
		return a.AuthenticateMessage(network, height, msg.Params, msgRct.Return)
	case parser.UnknownStr:
		// return p.unknownMetadata(msg.Params, msgRct.Return)
	}
	return map[string]interface{}{}, parser.ErrUnknownMethod
}
