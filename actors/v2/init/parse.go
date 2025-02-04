package init

import (
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
)

func (i *Init) Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, _ cid.Cid) (map[string]interface{}, *types.AddressInfo, error) {
	var err error
	metadata := make(map[string]interface{})
	switch txType {
	case parser.MethodSend:
		// metadata, err = i.Send(msg), nil
	case parser.MethodConstructor:
		metadata, err = i.Constructor(network, height, msg.Params)
	case parser.MethodExec:
		return i.Exec(network, height, msg, msgRct.Return)
	case parser.MethodExec4:
		return i.Exec4(network, height, msg, msgRct.Return)
	case parser.UnknownStr:
		// metadata, err = p.unknownMetadata(msg.Params, msgRct.Return)
	default:
		err = parser.ErrUnknownMethod
	}
	return metadata, nil, err
}

func (i *Init) TransactionTypes() []string {
	return []string{
		parser.MethodSend,
		parser.MethodConstructor,
		parser.MethodExec,
		parser.MethodExec4,
	}
}
