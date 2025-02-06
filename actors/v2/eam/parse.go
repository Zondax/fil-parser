package eam

import (
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
)

func (p *Eam) Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, msgCid cid.Cid, _ filTypes.TipSetKey) (map[string]interface{}, *types.AddressInfo, error) {
	metadata := make(map[string]interface{})
	var err error
	switch txType {
	case parser.MethodConstructor:
		resp, err := actors.ParseEmptyParamsAndReturn()
		return resp, nil, err
	case parser.MethodCreate:
		return p.Create(network, height, msg.Params, msgRct.Return, msgCid)
	case parser.MethodCreate2:
		return p.Create2(network, height, msg.Params, msgRct.Return, msgCid)
	case parser.MethodCreateExternal:
		return p.CreateExternal(network, height, msg.Params, msgRct.Return, msgCid)
	case parser.UnknownStr:
		resp, err := actors.ParseUnknownMetadata(msg.Params, msgRct.Return)
		return resp, nil, err
	default:
		err = parser.ErrUnknownMethod
	}
	return metadata, nil, err
}

func (p *Eam) TransactionTypes() map[string]any {
	return map[string]any{
		parser.MethodConstructor:    actors.ParseEmptyParamsAndReturn,
		parser.MethodCreate:         p.Create,
		parser.MethodCreate2:        p.Create2,
		parser.MethodCreateExternal: p.CreateExternal,
	}
}
