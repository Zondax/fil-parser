package actors

import (
	"bytes"
	"context"
	"encoding/hex"

	"github.com/filecoin-project/go-state-types/abi"
	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"

	"github.com/filecoin-project/go-address"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
)

type ActorParserInterface interface {
	GetMetadata(ctx context.Context, txType string, msg *parser.LotusMessage, mainMsgCid cid.Cid, msgRct *parser.LotusMessageReceipt,
		height int64, key filTypes.TipSetKey) (string, map[string]interface{}, *types.AddressInfo, error)
}

type Actor interface {
	Name() string
	Parse(ctx context.Context, network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, mainMsgCid cid.Cid, key filTypes.TipSetKey) (map[string]interface{}, *types.AddressInfo, error)
	StartNetworkHeight() int64
	TransactionTypes() map[string]any
	Methods(ctx context.Context, network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error)
}

func ParseSend(msg *parser.LotusMessage) map[string]interface{} {
	metadata := make(map[string]interface{})
	metadata[parser.ParamsKey] = msg.Params
	return metadata
}

// ParseConstructor parse methods with format: *new(func(*address.Address) *abi.EmptyValue)
func ParseConstructor(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params address.Address
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params.String()
	return metadata, nil
}

func ParseUnknownMetadata(msgParams, msgReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	if len(msgParams) > 0 {
		metadata[parser.ParamsKey] = hex.EncodeToString(msgParams)
	}
	if len(msgReturn) > 0 {
		metadata[parser.ReturnKey] = hex.EncodeToString(msgReturn)
	}
	return metadata, nil
}

func ParseEmptyParamsAndReturn() (map[string]interface{}, error) {
	return make(map[string]interface{}), nil
}

func CopyMethods(methods ...map[abi.MethodNum]nonLegacyBuiltin.MethodMeta) map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	res := make(map[abi.MethodNum]nonLegacyBuiltin.MethodMeta)
	for _, method := range methods {
		for k, v := range method {
			res[k] = v
		}
	}

	return res
}
