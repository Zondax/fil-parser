package eam

import (
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
)

func (p *Eam) Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, msgCid cid.Cid) (map[string]interface{}, *types.AddressInfo, error) {
	metadata := make(map[string]interface{})
	var err error
	switch txType {
	case parser.MethodConstructor:
		// metadata, err = p.emptyParamsAndReturn()
	case parser.MethodCreate:
		return p.Create(network, height, msg.Params, msgRct.Return, msgCid)
	case parser.MethodCreate2:
		return p.Create2(network, height, msg.Params, msgRct.Return, msgCid)
	case parser.MethodCreateExternal:
		return p.CreateExternal(network, height, msg.Params, msgRct.Return, msgCid)
	case parser.UnknownStr:
		// metadata, err = p.unknownMetadata(msg.Params, msgRct.Return)
	default:
		err = parser.ErrUnknownMethod
	}
	return metadata, nil, err
}

func (p *Eam) TransactionTypes() []string {
	return []string{
		parser.MethodConstructor,
		parser.MethodCreate,
		parser.MethodCreate2,
		parser.MethodCreateExternal,
	}
}
