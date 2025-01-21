package multisig

import (
	"bytes"
	"io"

	filTypes "github.com/filecoin-project/lotus/chain/types"

	"github.com/filecoin-project/go-state-types/abi"
	multisig10 "github.com/filecoin-project/go-state-types/builtin/v10/multisig"
	multisig11 "github.com/filecoin-project/go-state-types/builtin/v11/multisig"
	multisig12 "github.com/filecoin-project/go-state-types/builtin/v12/multisig"
	multisig13 "github.com/filecoin-project/go-state-types/builtin/v13/multisig"
	multisig14 "github.com/filecoin-project/go-state-types/builtin/v14/multisig"
	multisig15 "github.com/filecoin-project/go-state-types/builtin/v15/multisig"
	multisig8 "github.com/filecoin-project/go-state-types/builtin/v8/multisig"
	multisig9 "github.com/filecoin-project/go-state-types/builtin/v9/multisig"

	"github.com/zondax/fil-parser/parser"
)

func parseWithMsigParser[T multisigParams, R multisigReturn](msg *parser.LotusMessage,
	height int64,
	key filTypes.TipSetKey,
	fn parseFn,
	rawReturn []byte,
	unmarshaller func(io.Reader, any) error,
	customReturn bool,
) (map[string]interface{}, error) {

	metadata := make(map[string]interface{})
	params, err := fn(msg, height, key)
	if err != nil {
		return map[string]interface{}{}, err
	}
	metadata[parser.ParamsKey] = params

	if customReturn {
		var r R
		err = unmarshaller(bytes.NewReader(rawReturn), &r)
		if err != nil {
			return map[string]interface{}{}, err
		}
		metadata[parser.ReturnKey] = r
	}
	return metadata, nil

}

func parse[T multisigParams, P []byte | string](raw P, unmarshaller func(io.Reader, any) error) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	var params T
	rawBytes, err := toBytes(raw)
	if err != nil {
		return map[string]interface{}{}, err
	}
	reader := bytes.NewReader(rawBytes)
	err = unmarshaller(reader, &params)
	if err != nil {
		return map[string]interface{}{}, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func MsigConstructor(height int64, raw []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		return parse[*multisig8.ConstructorParams, []byte](raw, cborUnmarshaller[*multisig8.ConstructorParams])
	case 9:
		return parse[*multisig9.ConstructorParams, []byte](raw, cborUnmarshaller[*multisig9.ConstructorParams])
	case 10:
		return parse[*multisig10.ConstructorParams, []byte](raw, cborUnmarshaller[*multisig10.ConstructorParams])
	case 11:
		return parse[*multisig11.ConstructorParams, []byte](raw, cborUnmarshaller[*multisig11.ConstructorParams])
	case 12:
		return parse[*multisig12.ConstructorParams, []byte](raw, cborUnmarshaller[*multisig12.ConstructorParams])
	case 13:
		return parse[*multisig13.ConstructorParams, []byte](raw, cborUnmarshaller[*multisig13.ConstructorParams])
	case 14:
		return parse[*multisig14.ConstructorParams, []byte](raw, cborUnmarshaller[*multisig14.ConstructorParams])
	case 15:
		return parse[*multisig15.ConstructorParams, []byte](raw, cborUnmarshaller[*multisig15.ConstructorParams])
	}
	return map[string]interface{}{}, nil
}

func MsigParams(msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, parser parseFn) (map[string]interface{}, error) {
	switch height {
	case 8:
		return parseWithMsigParser[*multisig8.ConstructorParams, *multisig8.ConstructorParams](msg, height, key, parser, nil, jsonUnmarshaller[*multisig8.ConstructorParams], false)
	}
	return map[string]interface{}{}, nil
}

func Approve(msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, rawReturn []byte, parser parseFn) (map[string]interface{}, error) {
	switch height {
	case 8:
		return parseWithMsigParser[*multisig8.ApproveReturn, *multisig8.ApproveReturn](msg, height, key, parser, rawReturn, cborUnmarshaller[*multisig8.ApproveReturn], true)
	}
	return map[string]interface{}{}, nil
}

func Cancel(msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, rawReturn []byte, parser parseFn) (map[string]interface{}, error) {
	switch height {
	case 8:
		return parseWithMsigParser[metadataWithCbor, metadataWithCbor](msg, height, key, parser, rawReturn, noopUnmarshaller[metadataWithCbor], false)
	}
	return map[string]interface{}{}, nil
}

func RemoveSigner(msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, rawReturn []byte, parser parseFn) (map[string]interface{}, error) {
	switch height {
	case 8:
		return parseWithMsigParser[metadataWithCbor, metadataWithCbor](msg, height, key, parser, rawReturn, noopUnmarshaller[metadataWithCbor], false)
	}
	return map[string]interface{}{}, nil
}

func ChangeNumApprovalsThreshold(msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, rawReturn []byte, parser parseFn) (map[string]interface{}, error) {
	switch height {
	case 8:
		return parse[*multisig8.ChangeNumApprovalsThresholdParams, []byte](rawReturn, cborUnmarshaller[*multisig8.ChangeNumApprovalsThresholdParams])
	}
	return map[string]interface{}{}, nil
}

func LockBalance(msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, rawReturn []byte, parser parseFn) (map[string]interface{}, error) {
	switch height {
	case 8:
		return parse[*multisig8.LockBalanceParams, []byte](rawReturn, cborUnmarshaller[*multisig8.LockBalanceParams])
	}
	return map[string]interface{}{}, nil
}

func UniversalReceiverHook(msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, rawReturn []byte, parser parseFn) (map[string]interface{}, error) {
	switch height {
	case 8:
		return parse[*abi.CborBytesTransparent, []byte](rawReturn, cborUnmarshaller[*abi.CborBytesTransparent])
	}
	return map[string]interface{}{}, nil
}
