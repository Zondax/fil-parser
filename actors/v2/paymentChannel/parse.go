package paymentChannel

import (
	"context"

	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
)

func (p *PaymentChannel) Parse(_ context.Context, network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, _ cid.Cid, _ filTypes.TipSetKey) (map[string]interface{}, *types.AddressInfo, error) {
	switch txType {
	case parser.MethodSend:
		resp := actors.ParseSend(msg)
		return resp, nil, nil
	case parser.MethodConstructor:
		resp, err := p.Constructor(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodUpdateChannelState:
		resp, err := p.UpdateChannelState(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodSettle, parser.MethodCollect:
		resp, err := actors.ParseEmptyParamsAndReturn()
		return resp, nil, err
	case parser.UnknownStr:
		resp, err := actors.ParseUnknownMetadata(msg.Params, msgRct.Return)
		return resp, nil, err
	}
	return map[string]interface{}{}, nil, parser.ErrUnknownMethod
}

func (p *PaymentChannel) TransactionTypes() map[string]any {
	return map[string]any{
		parser.MethodSend:               actors.ParseSend,
		parser.MethodConstructor:        p.Constructor,
		parser.MethodUpdateChannelState: p.UpdateChannelState,
		parser.MethodSettle:             actors.ParseEmptyParamsAndReturn,
		parser.MethodCollect:            actors.ParseEmptyParamsAndReturn,
	}
}
