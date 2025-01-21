package multisig

import (
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
	"github.com/zondax/fil-parser/tools"
)

func MsigConstructor(height int64, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(height):
		return parse[*multisig8.ConstructorParams, []byte](raw, cborUnmarshaller[*multisig8.ConstructorParams])
	case tools.V9.IsSupported(height):
		return parse[*multisig9.ConstructorParams, []byte](raw, cborUnmarshaller[*multisig9.ConstructorParams])
	case tools.V10.IsSupported(height):
		return parse[*multisig10.ConstructorParams, []byte](raw, cborUnmarshaller[*multisig10.ConstructorParams])
	case tools.V11.IsSupported(height):
		return parse[*multisig11.ConstructorParams, []byte](raw, cborUnmarshaller[*multisig11.ConstructorParams])
	case tools.V12.IsSupported(height):
		return parse[*multisig12.ConstructorParams, []byte](raw, cborUnmarshaller[*multisig12.ConstructorParams])
	case tools.V13.IsSupported(height):
		return parse[*multisig13.ConstructorParams, []byte](raw, cborUnmarshaller[*multisig13.ConstructorParams])
	case tools.V14.IsSupported(height):
		return parse[*multisig14.ConstructorParams, []byte](raw, cborUnmarshaller[*multisig14.ConstructorParams])
	case tools.V15.IsSupported(height):
		return parse[*multisig15.ConstructorParams, []byte](raw, cborUnmarshaller[*multisig15.ConstructorParams])
	}
	return map[string]interface{}{}, nil
}

func MsigParams(msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, parser parseFn) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(height):
		return parseWithMsigParser[*multisig8.ConstructorParams, *multisig8.ConstructorParams](msg, height, key, parser, nil, jsonUnmarshaller[*multisig8.ConstructorParams], false)
	case tools.V9.IsSupported(height):
		return parseWithMsigParser[*multisig9.ConstructorParams, *multisig9.ConstructorParams](msg, height, key, parser, nil, jsonUnmarshaller[*multisig9.ConstructorParams], false)
	case tools.V10.IsSupported(height):
		return parseWithMsigParser[*multisig10.ConstructorParams, *multisig10.ConstructorParams](msg, height, key, parser, nil, jsonUnmarshaller[*multisig10.ConstructorParams], false)
	case tools.V11.IsSupported(height):
		return parseWithMsigParser[*multisig11.ConstructorParams, *multisig11.ConstructorParams](msg, height, key, parser, nil, jsonUnmarshaller[*multisig11.ConstructorParams], false)
	case tools.V12.IsSupported(height):
		return parseWithMsigParser[*multisig12.ConstructorParams, *multisig12.ConstructorParams](msg, height, key, parser, nil, jsonUnmarshaller[*multisig12.ConstructorParams], false)
	case tools.V13.IsSupported(height):
		return parseWithMsigParser[*multisig13.ConstructorParams, *multisig13.ConstructorParams](msg, height, key, parser, nil, jsonUnmarshaller[*multisig13.ConstructorParams], false)
	case tools.V14.IsSupported(height):
		return parseWithMsigParser[*multisig14.ConstructorParams, *multisig14.ConstructorParams](msg, height, key, parser, nil, jsonUnmarshaller[*multisig14.ConstructorParams], false)
	case tools.V15.IsSupported(height):
		return parseWithMsigParser[*multisig15.ConstructorParams, *multisig15.ConstructorParams](msg, height, key, parser, nil, jsonUnmarshaller[*multisig15.ConstructorParams], false)
	}
	return map[string]interface{}{}, nil
}

func Approve(msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, rawReturn []byte, parser parseFn) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(height):
		return parseWithMsigParser[*multisig8.ApproveReturn, *multisig8.ApproveReturn](msg, height, key, parser, rawReturn, cborUnmarshaller[*multisig8.ApproveReturn], true)
	case tools.V9.IsSupported(height):
		return parseWithMsigParser[*multisig9.ApproveReturn, *multisig9.ApproveReturn](msg, height, key, parser, rawReturn, cborUnmarshaller[*multisig9.ApproveReturn], true)
	case tools.V10.IsSupported(height):
		return parseWithMsigParser[*multisig10.ApproveReturn, *multisig10.ApproveReturn](msg, height, key, parser, rawReturn, cborUnmarshaller[*multisig10.ApproveReturn], true)
	case tools.V11.IsSupported(height):
		return parseWithMsigParser[*multisig11.ApproveReturn, *multisig11.ApproveReturn](msg, height, key, parser, rawReturn, cborUnmarshaller[*multisig11.ApproveReturn], true)
	case tools.V12.IsSupported(height):
		return parseWithMsigParser[*multisig12.ApproveReturn, *multisig12.ApproveReturn](msg, height, key, parser, rawReturn, cborUnmarshaller[*multisig12.ApproveReturn], true)
	case tools.V13.IsSupported(height):
		return parseWithMsigParser[*multisig13.ApproveReturn, *multisig13.ApproveReturn](msg, height, key, parser, rawReturn, cborUnmarshaller[*multisig13.ApproveReturn], true)
	case tools.V14.IsSupported(height):
		return parseWithMsigParser[*multisig14.ApproveReturn, *multisig14.ApproveReturn](msg, height, key, parser, rawReturn, cborUnmarshaller[*multisig14.ApproveReturn], true)
	case tools.V15.IsSupported(height):
		return parseWithMsigParser[*multisig15.ApproveReturn, *multisig15.ApproveReturn](msg, height, key, parser, rawReturn, cborUnmarshaller[*multisig15.ApproveReturn], true)
	}
	return map[string]interface{}{}, nil
}

func Cancel(msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, rawReturn []byte, parser parseFn) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(height):
		return parseWithMsigParser[metadataWithCbor, metadataWithCbor](msg, height, key, parser, rawReturn, noopUnmarshaller[metadataWithCbor], false)
	case tools.V9.IsSupported(height):
		return parseWithMsigParser[metadataWithCbor, metadataWithCbor](msg, height, key, parser, rawReturn, noopUnmarshaller[metadataWithCbor], false)
	case tools.V10.IsSupported(height):
		return parseWithMsigParser[metadataWithCbor, metadataWithCbor](msg, height, key, parser, rawReturn, noopUnmarshaller[metadataWithCbor], false)
	case tools.V11.IsSupported(height):
		return parseWithMsigParser[metadataWithCbor, metadataWithCbor](msg, height, key, parser, rawReturn, noopUnmarshaller[metadataWithCbor], false)
	case tools.V12.IsSupported(height):
		return parseWithMsigParser[metadataWithCbor, metadataWithCbor](msg, height, key, parser, rawReturn, noopUnmarshaller[metadataWithCbor], false)
	case tools.V13.IsSupported(height):
		return parseWithMsigParser[metadataWithCbor, metadataWithCbor](msg, height, key, parser, rawReturn, noopUnmarshaller[metadataWithCbor], false)
	case tools.V14.IsSupported(height):
		return parseWithMsigParser[metadataWithCbor, metadataWithCbor](msg, height, key, parser, rawReturn, noopUnmarshaller[metadataWithCbor], false)
	case tools.V15.IsSupported(height):
		return parseWithMsigParser[metadataWithCbor, metadataWithCbor](msg, height, key, parser, rawReturn, noopUnmarshaller[metadataWithCbor], false)
	}
	return map[string]interface{}{}, nil
}

func RemoveSigner(msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, rawReturn []byte, parser parseFn) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(height):
		return parseWithMsigParser[metadataWithCbor, metadataWithCbor](msg, height, key, parser, rawReturn, noopUnmarshaller[metadataWithCbor], false)
	case tools.V9.IsSupported(height):
		return parseWithMsigParser[metadataWithCbor, metadataWithCbor](msg, height, key, parser, rawReturn, noopUnmarshaller[metadataWithCbor], false)
	case tools.V10.IsSupported(height):
		return parseWithMsigParser[metadataWithCbor, metadataWithCbor](msg, height, key, parser, rawReturn, noopUnmarshaller[metadataWithCbor], false)
	case tools.V11.IsSupported(height):
		return parseWithMsigParser[metadataWithCbor, metadataWithCbor](msg, height, key, parser, rawReturn, noopUnmarshaller[metadataWithCbor], false)
	case tools.V12.IsSupported(height):
		return parseWithMsigParser[metadataWithCbor, metadataWithCbor](msg, height, key, parser, rawReturn, noopUnmarshaller[metadataWithCbor], false)
	case tools.V13.IsSupported(height):
		return parseWithMsigParser[metadataWithCbor, metadataWithCbor](msg, height, key, parser, rawReturn, noopUnmarshaller[metadataWithCbor], false)
	case tools.V14.IsSupported(height):
		return parseWithMsigParser[metadataWithCbor, metadataWithCbor](msg, height, key, parser, rawReturn, noopUnmarshaller[metadataWithCbor], false)
	case tools.V15.IsSupported(height):
		return parseWithMsigParser[metadataWithCbor, metadataWithCbor](msg, height, key, parser, rawReturn, noopUnmarshaller[metadataWithCbor], false)
	}
	return map[string]interface{}{}, nil
}

func ChangeNumApprovalsThreshold(msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, rawReturn []byte, parser parseFn) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(height):
		return parse[*multisig8.ChangeNumApprovalsThresholdParams, []byte](rawReturn, cborUnmarshaller[*multisig8.ChangeNumApprovalsThresholdParams])
	case tools.V9.IsSupported(height):
		return parse[*multisig9.ChangeNumApprovalsThresholdParams, []byte](rawReturn, cborUnmarshaller[*multisig9.ChangeNumApprovalsThresholdParams])
	case tools.V10.IsSupported(height):
		return parse[*multisig10.ChangeNumApprovalsThresholdParams, []byte](rawReturn, cborUnmarshaller[*multisig10.ChangeNumApprovalsThresholdParams])
	case tools.V11.IsSupported(height):
		return parse[*multisig11.ChangeNumApprovalsThresholdParams, []byte](rawReturn, cborUnmarshaller[*multisig11.ChangeNumApprovalsThresholdParams])
	case tools.V12.IsSupported(height):
		return parse[*multisig12.ChangeNumApprovalsThresholdParams, []byte](rawReturn, cborUnmarshaller[*multisig12.ChangeNumApprovalsThresholdParams])
	case tools.V13.IsSupported(height):
		return parse[*multisig13.ChangeNumApprovalsThresholdParams, []byte](rawReturn, cborUnmarshaller[*multisig13.ChangeNumApprovalsThresholdParams])
	case tools.V14.IsSupported(height):
		return parse[*multisig14.ChangeNumApprovalsThresholdParams, []byte](rawReturn, cborUnmarshaller[*multisig14.ChangeNumApprovalsThresholdParams])
	case tools.V15.IsSupported(height):
		return parse[*multisig15.ChangeNumApprovalsThresholdParams, []byte](rawReturn, cborUnmarshaller[*multisig15.ChangeNumApprovalsThresholdParams])
	}
	return map[string]interface{}{}, nil
}

func LockBalance(msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, rawReturn []byte, parser parseFn) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(height):
		return parse[*multisig8.LockBalanceParams, []byte](rawReturn, cborUnmarshaller[*multisig8.LockBalanceParams])
	case tools.V9.IsSupported(height):
		return parse[*multisig9.LockBalanceParams, []byte](rawReturn, cborUnmarshaller[*multisig9.LockBalanceParams])
	case tools.V10.IsSupported(height):
		return parse[*multisig10.LockBalanceParams, []byte](rawReturn, cborUnmarshaller[*multisig10.LockBalanceParams])
	case tools.V11.IsSupported(height):
		return parse[*multisig11.LockBalanceParams, []byte](rawReturn, cborUnmarshaller[*multisig11.LockBalanceParams])
	case tools.V12.IsSupported(height):
		return parse[*multisig12.LockBalanceParams, []byte](rawReturn, cborUnmarshaller[*multisig12.LockBalanceParams])
	case tools.V13.IsSupported(height):
		return parse[*multisig13.LockBalanceParams, []byte](rawReturn, cborUnmarshaller[*multisig13.LockBalanceParams])
	case tools.V14.IsSupported(height):
		return parse[*multisig14.LockBalanceParams, []byte](rawReturn, cborUnmarshaller[*multisig14.LockBalanceParams])
	case tools.V15.IsSupported(height):
		return parse[*multisig15.LockBalanceParams, []byte](rawReturn, cborUnmarshaller[*multisig15.LockBalanceParams])
	}
	return map[string]interface{}{}, nil
}

func UniversalReceiverHook(msg *parser.LotusMessage, height int64, key filTypes.TipSetKey, rawReturn []byte, parser parseFn) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(height):
		return parse[*abi.CborBytesTransparent, []byte](rawReturn, cborUnmarshaller[*abi.CborBytesTransparent])
	case tools.V9.IsSupported(height):
		return parse[*abi.CborBytesTransparent, []byte](rawReturn, cborUnmarshaller[*abi.CborBytesTransparent])
	case tools.V10.IsSupported(height):
		return parse[*abi.CborBytesTransparent, []byte](rawReturn, cborUnmarshaller[*abi.CborBytesTransparent])
	case tools.V11.IsSupported(height):
		return parse[*abi.CborBytesTransparent, []byte](rawReturn, cborUnmarshaller[*abi.CborBytesTransparent])
	case tools.V12.IsSupported(height):
		return parse[*abi.CborBytesTransparent, []byte](rawReturn, cborUnmarshaller[*abi.CborBytesTransparent])
	case tools.V13.IsSupported(height):
		return parse[*abi.CborBytesTransparent, []byte](rawReturn, cborUnmarshaller[*abi.CborBytesTransparent])
	case tools.V14.IsSupported(height):
		return parse[*abi.CborBytesTransparent, []byte](rawReturn, cborUnmarshaller[*abi.CborBytesTransparent])
	case tools.V15.IsSupported(height):
		return parse[*abi.CborBytesTransparent, []byte](rawReturn, cborUnmarshaller[*abi.CborBytesTransparent])
	}
	return map[string]interface{}{}, nil
}
