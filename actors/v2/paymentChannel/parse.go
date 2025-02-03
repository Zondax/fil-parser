package paymentchannel

import (
	"github.com/zondax/fil-parser/parser"
)

func (p *PaymentChannel) Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt) (map[string]interface{}, error) {
	switch txType {
	case parser.MethodSend:
		// return p.parseSend(msg), nil
	case parser.MethodConstructor:
		return p.Constructor(network, height, msg.Params)
	case parser.MethodUpdateChannelState:
		return p.UpdateChannelState(network, height, msg.Params)
	case parser.MethodSettle, parser.MethodCollect:
		// return p.emptyParamsAndReturn()
	case parser.UnknownStr:
		// return p.unknownMetadata(msg.Params, msgRct.Return)
	}
	return map[string]interface{}{}, parser.ErrUnknownMethod
}
