package evm

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/filecoin-project/go-state-types/abi"
	evmv10 "github.com/filecoin-project/go-state-types/builtin/v10/evm"
	evmv11 "github.com/filecoin-project/go-state-types/builtin/v11/evm"
	evmv12 "github.com/filecoin-project/go-state-types/builtin/v12/evm"
	evmv13 "github.com/filecoin-project/go-state-types/builtin/v13/evm"
	evmv14 "github.com/filecoin-project/go-state-types/builtin/v14/evm"
	evmv15 "github.com/filecoin-project/go-state-types/builtin/v15/evm"
	"github.com/zondax/fil-parser/parser"
)

type evmParams interface {
	UnmarshalCBOR(io.Reader) error
}

type evmReturn interface {
	UnmarshalCBOR(io.Reader) error
}

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
	switch height {
	case 15:
		return getByteCode[*evmv15.ResurrectParams](parser.ParamsKey, raw)
	case 14:
		return getByteCode[*evmv14.ResurrectParams](parser.ParamsKey, raw)
	case 13:
		return getByteCode[*evmv13.ResurrectParams](parser.ParamsKey, raw)
	case 12:
		return getByteCode[*evmv12.ResurrectParams](parser.ParamsKey, raw)
	case 11:
		return getByteCode[*evmv11.ResurrectParams](parser.ParamsKey, raw)
	case 10:
		return getByteCode[*evmv10.ResurrectParams](parser.ParamsKey, raw)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func InvokeContractDelegate(height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return invokeContractDelegate[*evmv15.DelegateCallParams, *abi.CborBytes](rawParams, rawReturn)
	case 14:
		return invokeContractDelegate[*evmv14.DelegateCallParams, *abi.CborBytes](rawParams, rawReturn)
	case 13:
		return invokeContractDelegate[*evmv13.DelegateCallParams, *abi.CborBytes](rawParams, rawReturn)
	case 12:
		return invokeContractDelegate[*evmv12.DelegateCallParams, *abi.CborBytes](rawParams, rawReturn)
	case 11:
		return invokeContractDelegate[*evmv11.DelegateCallParams, *abi.CborBytes](rawParams, rawReturn)
	case 10:
		return invokeContractDelegate[*evmv10.DelegateCallParams, *abi.CborBytes](rawParams, rawReturn)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func GetByteCode(height int64, raw []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return getByteCode[*evmv15.GetBytecodeReturn](parser.ReturnKey, raw)
	case 14:
		return getByteCode[*evmv14.GetBytecodeReturn](parser.ReturnKey, raw)
	case 13:
		return getByteCode[*evmv13.GetBytecodeReturn](parser.ReturnKey, raw)
	case 12:
		return getByteCode[*evmv12.GetBytecodeReturn](parser.ReturnKey, raw)
	case 11:
		return getByteCode[*evmv11.GetBytecodeReturn](parser.ReturnKey, raw)
	case 10:
		return getByteCode[*evmv10.GetBytecodeReturn](parser.ReturnKey, raw)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func GetByteCodeHash(height int64, raw []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return getByteCode[*abi.CborBytes](parser.ReturnKey, raw)
	case 14:
		return getByteCode[*abi.CborBytes](parser.ReturnKey, raw)
	case 13:
		return getByteCode[*abi.CborBytes](parser.ReturnKey, raw)
	case 12:
		return getByteCode[*abi.CborBytes](parser.ReturnKey, raw)
	case 11:
		return getByteCode[*abi.CborBytes](parser.ReturnKey, raw)
	case 10:
		return getByteCode[*abi.CborBytes](parser.ReturnKey, raw)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func EVMConstructor(height int64, raw []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return getByteCode[*evmv15.ConstructorParams](parser.ParamsKey, raw)
	case 14:
		return getByteCode[*evmv14.ConstructorParams](parser.ParamsKey, raw)
	case 13:
		return getByteCode[*evmv13.ConstructorParams](parser.ParamsKey, raw)
	case 12:
		return getByteCode[*evmv12.ConstructorParams](parser.ParamsKey, raw)
	case 11:
		return getByteCode[*evmv11.ConstructorParams](parser.ParamsKey, raw)
	case 10:
		return getByteCode[*evmv10.ConstructorParams](parser.ParamsKey, raw)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func GetStorageAt(height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return getStorageAt[*evmv15.GetStorageAtParams, *abi.CborBytes](rawParams, rawReturn)
	case 14:
		return getStorageAt[*evmv14.GetStorageAtParams, *abi.CborBytes](rawParams, rawReturn)
	case 13:
		return getStorageAt[*evmv13.GetStorageAtParams, *abi.CborBytes](rawParams, rawReturn)
	case 12:
		return getStorageAt[*evmv12.GetStorageAtParams, *abi.CborBytes](rawParams, rawReturn)
	case 11:
		return getStorageAt[*evmv11.GetStorageAtParams, *abi.CborBytes](rawParams, rawReturn)
	case 10:
		return getStorageAt[*evmv10.GetStorageAtParams, *abi.CborBytes](rawParams, rawReturn)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func invokeContractDelegate[T evmParams, R evmReturn](rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params T
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var r R
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r
	return metadata, nil
}

func getByteCode[R evmReturn](key string, raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var r R
	err := r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[key] = r
	return metadata, nil
}

func getStorageAt[T evmParams, R evmReturn](rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params T
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var r R
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r
	return metadata, nil
}
