package eam

import (
	"context"

	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	actor_tools "github.com/zondax/fil-parser/actors/v2/tools"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
)

func (p *Eam) Parse(_ context.Context, network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, msgCid cid.Cid, _ filTypes.TipSetKey) (map[string]interface{}, *types.AddressInfo, error) {
	metadata := make(map[string]interface{})
	var err error
	switch txType {
	case parser.MethodSend:
		resp := actor_tools.ParseSend(msg)
		return resp, nil, nil
	case parser.MethodConstructor:
		resp, err := actor_tools.ParseEmptyParamsAndReturn()
		return resp, nil, err
	case parser.MethodCreate:
		return p.Create(network, height, msg.Params, msgRct.Return, msgCid)
	case parser.MethodCreate2:
		return p.Create2(network, height, msg.Params, msgRct.Return, msgCid)
	case parser.MethodCreateExternal:
		return p.CreateExternal(network, height, msg.Params, msgRct.Return, msgCid)
	case parser.UnknownStr:
		resp, err := actor_tools.ParseUnknownMetadata(msg.Params, msgRct.Return)
		return resp, nil, err
	default:
		err = parser.ErrUnknownMethod
	}
	return metadata, nil, err
}

func (p *Eam) TransactionTypes() map[string]any {
	return map[string]any{
		parser.MethodSend:           actor_tools.ParseSend,
		parser.MethodConstructor:    actor_tools.ParseEmptyParamsAndReturn,
		parser.MethodCreate:         p.Create,
		parser.MethodCreate2:        p.Create2,
		parser.MethodCreateExternal: p.CreateExternal,
	}
}
