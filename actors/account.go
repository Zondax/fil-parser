package actors

import (
	"github.com/zondax/fil-parser/actors/account"
	"github.com/zondax/fil-parser/parser"
)

func (p *ActorParser) ParseAccount(height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt) (map[string]interface{}, error) {
	switch txType {
	case parser.MethodSend:
		return p.parseSend(msg), nil
	case parser.MethodConstructor:
		return p.parseConstructor(msg.Params)
	case parser.MethodPubkeyAddress:
		return account.PubkeyAddress(msg.Params, msgRct.Return)
	case parser.MethodAuthenticateMessage:
		return account.AuthenticateMessage(height, msg.Params, msgRct.Return)
	case parser.UnknownStr:
		return p.unknownMetadata(msg.Params, msgRct.Return)
	}
	return map[string]interface{}{}, parser.ErrUnknownMethod
}
