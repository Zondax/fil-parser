package init

import (
	"context"

	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	actor_tools "github.com/zondax/fil-parser/actors/v2/tools"

	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
)

func (i *Init) Parse(_ context.Context, network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, _ cid.Cid, _ filTypes.TipSetKey) (map[string]interface{}, *types.AddressInfo, error) {
	var err error
	metadata := make(map[string]interface{})
	switch txType {
	case parser.MethodSend:
		resp := actor_tools.ParseSend(msg)
		return resp, nil, nil
	case parser.MethodConstructor:
		resp, err := i.Constructor(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodExec:
		return i.Exec(network, height, msg, msgRct.Return)
	case parser.MethodExec4:
		return i.Exec4(network, height, msg, msgRct.Return)
	case parser.UnknownStr:
		resp, err := actor_tools.ParseUnknownMetadata(msg.Params, msgRct.Return)
		return resp, nil, err
	default:
		err = parser.ErrUnknownMethod
	}
	return metadata, nil, err
}

func (i *Init) TransactionTypes() map[string]any {
	return map[string]any{
		parser.MethodSend:        actor_tools.ParseSend,
		parser.MethodConstructor: i.Constructor,
		parser.MethodExec:        i.Exec,
		parser.MethodExec4:       i.Exec4,
	}
}
