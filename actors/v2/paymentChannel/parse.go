package paymentchannel

import (
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
)

func (p *PaymentChannel) Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, _ cid.Cid) (map[string]interface{}, *types.AddressInfo, error) {
	switch txType {
	case parser.MethodSend:
		// return p.parseSend(msg), nil
	case parser.MethodConstructor:
		resp, err := p.Constructor(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodUpdateChannelState:
		resp, err := p.UpdateChannelState(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodSettle, parser.MethodCollect:
		// return p.emptyParamsAndReturn()
	case parser.UnknownStr:
		// return p.unknownMetadata(msg.Params, msgRct.Return)
	}
	return map[string]interface{}{}, nil, parser.ErrUnknownMethod
}

func (p *PaymentChannel) TransactionTypes() []string {
	return []string{
		parser.MethodSend,
		parser.MethodConstructor,
		parser.MethodUpdateChannelState,
		parser.MethodSettle,
		parser.MethodCollect,
	}
}
