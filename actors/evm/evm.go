package evm

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/filecoin-project/go-state-types/abi"
	evmv10 "github.com/filecoin-project/go-state-types/builtin/v10/evm"
	evmv11 "github.com/filecoin-project/go-state-types/builtin/v11/evm"
	evmv12 "github.com/filecoin-project/go-state-types/builtin/v12/evm"
	evmv13 "github.com/filecoin-project/go-state-types/builtin/v13/evm"
	evmv14 "github.com/filecoin-project/go-state-types/builtin/v14/evm"
	evmv15 "github.com/filecoin-project/go-state-types/builtin/v15/evm"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

func InvokeContract(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	metadata[parser.ParamsKey] = parser.EthPrefix + hex.EncodeToString(rawParams)
	metadata[parser.ReturnKey] = parser.EthPrefix + hex.EncodeToString(rawReturn)

	var params abi.CborBytes
	if err := params.UnmarshalCBOR(reader); err != nil {
		return metadata, fmt.Errorf("error deserializing rawParams: %s - hex data: %s", err.Error(), hex.EncodeToString(rawParams))
	}

	if reader.Len() == 0 { // This means that the reader has processed all the bytes
		metadata[parser.ParamsKey] = parser.EthPrefix + hex.EncodeToString(params)
	}

	reader = bytes.NewReader(rawReturn)
	var returnValue abi.CborBytes
	if err := returnValue.UnmarshalCBOR(reader); err != nil {
		return metadata, fmt.Errorf("error deserializing rawReturn: %s - hex data: %s", err.Error(), hex.EncodeToString(rawReturn))
	}

	if reader.Len() == 0 { // This means that the reader has processed all the bytes
		metadata[parser.ReturnKey] = parser.EthPrefix + hex.EncodeToString(returnValue)
	}

	return metadata, nil
}

func Resurrect(height int64, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		return parse[*evmv15.ResurrectParams, *evmv15.ResurrectParams](raw, nil, false)
	case tools.V14.IsSupported(height):
		return parse[*evmv14.ResurrectParams, *evmv14.ResurrectParams](raw, nil, false)
	case tools.V13.IsSupported(height):
		return parse[*evmv13.ResurrectParams, *evmv13.ResurrectParams](raw, nil, false)
	case tools.V12.IsSupported(height):
		return parse[*evmv12.ResurrectParams, *evmv12.ResurrectParams](raw, nil, false)
	case tools.V11.IsSupported(height):
		return parse[*evmv11.ResurrectParams, *evmv11.ResurrectParams](raw, nil, false)
	case tools.V10.IsSupported(height):
		return parse[*evmv10.ResurrectParams, *evmv10.ResurrectParams](raw, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func InvokeContractDelegate(height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		return parse[*evmv15.DelegateCallParams, *abi.CborBytes](rawParams, rawReturn, true)
	case tools.V14.IsSupported(height):
		return parse[*evmv14.DelegateCallParams, *abi.CborBytes](rawParams, rawReturn, true)
	case tools.V13.IsSupported(height):
		return parse[*evmv13.DelegateCallParams, *abi.CborBytes](rawParams, rawReturn, true)
	case tools.V12.IsSupported(height):
		return parse[*evmv12.DelegateCallParams, *abi.CborBytes](rawParams, rawReturn, true)
	case tools.V11.IsSupported(height):
		return parse[*evmv11.DelegateCallParams, *abi.CborBytes](rawParams, rawReturn, true)
	case tools.V10.IsSupported(height):
		return parse[*evmv10.DelegateCallParams, *abi.CborBytes](rawParams, rawReturn, true)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func GetByteCode(height int64, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		data, err := parse[*evmv15.GetBytecodeReturn, *evmv15.GetBytecodeReturn](raw, nil, false)
		if err != nil {
			return nil, err
		}
		// The return value is the same as the params
		data[parser.ReturnKey] = data[parser.ParamsKey]
		return data, nil
	case tools.V14.IsSupported(height):
		data, err := parse[*evmv14.GetBytecodeReturn, *evmv14.GetBytecodeReturn](raw, nil, false)
		if err != nil {
			return nil, err
		}
		// The return value is the same as the params
		data[parser.ReturnKey] = data[parser.ParamsKey]
		return data, nil
	case tools.V13.IsSupported(height):
		data, err := parse[*evmv13.GetBytecodeReturn, *evmv13.GetBytecodeReturn](raw, nil, false)
		if err != nil {
			return nil, err
		}
		// The return value is the same as the params
		data[parser.ReturnKey] = data[parser.ParamsKey]
		return data, nil
	case tools.V11.IsSupported(height):
		data, err := parse[*evmv11.GetBytecodeReturn, *evmv11.GetBytecodeReturn](raw, nil, false)
		if err != nil {
			return nil, err
		}
		// The return value is the same as the params
		data[parser.ReturnKey] = data[parser.ParamsKey]
		return data, nil
	case tools.V10.IsSupported(height):
		data, err := parse[*evmv10.GetBytecodeReturn, *evmv10.GetBytecodeReturn](raw, nil, false)
		if err != nil {
			return nil, err
		}
		// The return value is the same as the params
		data[parser.ReturnKey] = data[parser.ParamsKey]
		return data, nil
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func GetByteCodeHash(height int64, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		data, err := parse[*abi.CborBytes, *abi.CborBytes](raw, nil, false)
		if err != nil {
			return nil, err
		}
		// The return value is the same as the params
		data[parser.ReturnKey] = data[parser.ParamsKey]
		return data, nil
	case tools.V14.IsSupported(height):
		data, err := parse[*abi.CborBytes, *abi.CborBytes](raw, nil, false)
		if err != nil {
			return nil, err
		}
		// The return value is the same as the params
		data[parser.ReturnKey] = data[parser.ParamsKey]
		return data, nil
	case tools.V13.IsSupported(height):
		data, err := parse[*abi.CborBytes, *abi.CborBytes](raw, nil, false)
		if err != nil {
			return nil, err
		}
		// The return value is the same as the params
		data[parser.ReturnKey] = data[parser.ParamsKey]
		return data, nil
	case tools.V12.IsSupported(height):
		data, err := parse[*abi.CborBytes, *abi.CborBytes](raw, nil, false)
		if err != nil {
			return nil, err
		}
		// The return value is the same as the params
		data[parser.ReturnKey] = data[parser.ParamsKey]
		return data, nil
	case tools.V11.IsSupported(height):
		data, err := parse[*abi.CborBytes, *abi.CborBytes](raw, nil, false)
		if err != nil {
			return nil, err
		}
		// The return value is the same as the params
		data[parser.ReturnKey] = data[parser.ParamsKey]
		return data, nil
	case tools.V10.IsSupported(height):
		data, err := parse[*abi.CborBytes, *abi.CborBytes](raw, nil, false)
		if err != nil {
			return nil, err
		}
		// The return value is the same as the params
		data[parser.ReturnKey] = data[parser.ParamsKey]
		return data, nil
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func EVMConstructor(height int64, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		return parse[*evmv15.ConstructorParams, *evmv15.ConstructorParams](raw, nil, false)
	case tools.V14.IsSupported(height):
		return parse[*evmv14.ConstructorParams, *evmv14.ConstructorParams](raw, nil, false)
	case tools.V13.IsSupported(height):
		return parse[*evmv13.ConstructorParams, *evmv13.ConstructorParams](raw, nil, false)
	case tools.V12.IsSupported(height):
		return parse[*evmv12.ConstructorParams, *evmv12.ConstructorParams](raw, nil, false)
	case tools.V11.IsSupported(height):
		return parse[*evmv11.ConstructorParams, *evmv11.ConstructorParams](raw, nil, false)
	case tools.V10.IsSupported(height):
		return parse[*evmv10.ConstructorParams, *evmv10.ConstructorParams](raw, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func GetStorageAt(height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		return parse[*evmv15.GetStorageAtParams, *abi.CborBytes](rawParams, rawReturn, true)
	case tools.V14.IsSupported(height):
		return parse[*evmv14.GetStorageAtParams, *abi.CborBytes](rawParams, rawReturn, true)
	case tools.V13.IsSupported(height):
		return parse[*evmv13.GetStorageAtParams, *abi.CborBytes](rawParams, rawReturn, true)
	case tools.V12.IsSupported(height):
		return parse[*evmv12.GetStorageAtParams, *abi.CborBytes](rawParams, rawReturn, true)
	case tools.V11.IsSupported(height):
		return parse[*evmv11.GetStorageAtParams, *abi.CborBytes](rawParams, rawReturn, true)
	case tools.V10.IsSupported(height):
		return parse[*evmv10.GetStorageAtParams, *abi.CborBytes](rawParams, rawReturn, true)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}
