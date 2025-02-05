package ethaccount

import (
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
)

type EthAccount struct{}

func (e *EthAccount) Name() string {
	return manifest.EthAccountKey
}

func (e *EthAccount) Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, mainMsgCid cid.Cid, key filTypes.TipSetKey) (map[string]interface{}, *types.AddressInfo, error) {
	var resp map[string]interface{}
	var err error
	switch txType {
	case parser.MethodConstructor:
		resp, err = e.Constructor()
	default:
		resp, err = e.parseEthAccountAny(msg.Params, msgRct.Return)
	}

	return resp, nil, err
}

func (e *EthAccount) TransactionTypes() map[string]any {
	return map[string]any{
		parser.MethodConstructor: e.Constructor,
	}
}

func (e *EthAccount) Constructor() (map[string]interface{}, error) {
	return e.parseEthAccountAny(nil, nil)
}

func (e *EthAccount) parseEthAccountAny(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	metadata[parser.ParamsKey] = rawParams
	metadata[parser.ReturnKey] = rawReturn

	return metadata, nil
}
