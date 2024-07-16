package v11

import (
	"bytes"

	"github.com/filecoin-project/go-state-types/builtin/v8/multisig"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/zondax/fil-parser/parser"
)

func ParseMultisig(txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, height int64, key filTypes.TipSetKey) (map[string]interface{}, error) {
	switch txType {
	case parser.MethodConstructor:
		return msigConstructor(msg.Params)
	case parser.MethodChangeNumApprovalsThreshold:
		return changeNumApprovalsThreshold(msg.Params)
	case parser.MethodLockBalance:
		return lockBalance(msg.Params)
		//// TODO: the rest of the methods
	}
	return map[string]interface{}{}, parser.ErrUnknownMethod
}

func msigConstructor(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var proposeParams multisig.ConstructorParams
	err := proposeParams.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = proposeParams
	return metadata, nil
}

func changeNumApprovalsThreshold(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	var params multisig.ChangeNumApprovalsThresholdParams
	reader := bytes.NewReader(raw)
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func lockBalance(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	var params multisig.LockBalanceParams
	reader := bytes.NewReader(raw)
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

// the rest of code...
