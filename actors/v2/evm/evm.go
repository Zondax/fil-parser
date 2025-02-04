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
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

type Evm struct{}

func (e *Evm) Name() string {
	return manifest.EvmKey
}

func (*Evm) InvokeContract(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
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

func (*Evm) Resurrect(network string, height int64, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parse(raw, nil, false, &evmv15.ResurrectParams{}, &evmv15.ResurrectParams{})
	case tools.V23.IsSupported(network, height):
		return parse(raw, nil, false, &evmv14.ResurrectParams{}, &evmv14.ResurrectParams{})
	case tools.V22.IsSupported(network, height):
		return parse(raw, nil, false, &evmv13.ResurrectParams{}, &evmv13.ResurrectParams{})
	case tools.V21.IsSupported(network, height):
		return parse(raw, nil, false, &evmv12.ResurrectParams{}, &evmv12.ResurrectParams{})
	case tools.AnyIsSupported(network, height, tools.V20, tools.V19):
		return parse(raw, nil, false, &evmv11.ResurrectParams{}, &evmv11.ResurrectParams{})
	case tools.V18.IsSupported(network, height):
		return parse(raw, nil, false, &evmv10.ResurrectParams{}, &evmv10.ResurrectParams{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V17)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Evm) InvokeContractDelegate(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parse(rawParams, rawReturn, true, &evmv15.DelegateCallParams{}, &abi.CborBytes{})
	case tools.V23.IsSupported(network, height):
		return parse(rawParams, rawReturn, true, &evmv14.DelegateCallParams{}, &abi.CborBytes{})
	case tools.V22.IsSupported(network, height):
		return parse(rawParams, rawReturn, true, &evmv13.DelegateCallParams{}, &abi.CborBytes{})
	case tools.V21.IsSupported(network, height):
		return parse(rawParams, rawReturn, true, &evmv12.DelegateCallParams{}, &abi.CborBytes{})
	case tools.AnyIsSupported(network, height, tools.V20, tools.V19):
		return parse(rawParams, rawReturn, true, &evmv11.DelegateCallParams{}, &abi.CborBytes{})
	case tools.V18.IsSupported(network, height):
		return parse(rawParams, rawReturn, true, &evmv10.DelegateCallParams{}, &abi.CborBytes{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V17)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Evm) GetBytecode(network string, height int64, raw []byte) (map[string]interface{}, error) {
	var data map[string]interface{}
	var err error
	switch {
	case tools.V24.IsSupported(network, height):
		data, err = parse(raw, nil, false, &evmv15.GetBytecodeReturn{}, &evmv15.GetBytecodeReturn{})
	case tools.V23.IsSupported(network, height):
		data, err = parse(raw, nil, false, &evmv14.GetBytecodeReturn{}, &evmv14.GetBytecodeReturn{})
	case tools.V22.IsSupported(network, height):
		data, err = parse(raw, nil, false, &evmv13.GetBytecodeReturn{}, &evmv13.GetBytecodeReturn{})
	case tools.V21.IsSupported(network, height):
		data, err = parse(raw, nil, false, &evmv12.GetBytecodeReturn{}, &evmv12.GetBytecodeReturn{})
	case tools.AnyIsSupported(network, height, tools.V20, tools.V19):
		data, err = parse(raw, nil, false, &evmv11.GetBytecodeReturn{}, &evmv11.GetBytecodeReturn{})
	case tools.V18.IsSupported(network, height):
		data, err = parse(raw, nil, false, &evmv10.GetBytecodeReturn{}, &evmv10.GetBytecodeReturn{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V17)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	default:
		err = fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	if data != nil {
		data[parser.ReturnKey] = data[parser.ParamsKey]
	}
	return data, err
}

func (*Evm) GetBytecodeHash(network string, height int64, raw []byte) (map[string]interface{}, error) {
	var data map[string]interface{}
	var err error
	switch {
	case tools.V24.IsSupported(network, height):
		data, err = parse(raw, nil, false, &abi.CborBytes{}, &abi.CborBytes{})
	case tools.V23.IsSupported(network, height):
		data, err = parse(raw, nil, false, &abi.CborBytes{}, &abi.CborBytes{})
	case tools.V22.IsSupported(network, height):
		data, err = parse(raw, nil, false, &abi.CborBytes{}, &abi.CborBytes{})
	case tools.V21.IsSupported(network, height):
		data, err = parse(raw, nil, false, &abi.CborBytes{}, &abi.CborBytes{})
	case tools.AnyIsSupported(network, height, tools.V20, tools.V19):
		data, err = parse(raw, nil, false, &abi.CborBytes{}, &abi.CborBytes{})
	case tools.V18.IsSupported(network, height):
		data, err = parse(raw, nil, false, &abi.CborBytes{}, &abi.CborBytes{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V17)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	default:
		err = fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	if data != nil {
		data[parser.ReturnKey] = data[parser.ParamsKey]
	}

	return data, err
}

func (*Evm) Constructor(network string, height int64, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parse(raw, nil, false, &evmv15.ConstructorParams{}, &evmv15.ConstructorParams{})
	case tools.V23.IsSupported(network, height):
		return parse(raw, nil, false, &evmv14.ConstructorParams{}, &evmv14.ConstructorParams{})
	case tools.V22.IsSupported(network, height):
		return parse(raw, nil, false, &evmv13.ConstructorParams{}, &evmv13.ConstructorParams{})
	case tools.V21.IsSupported(network, height):
		return parse(raw, nil, false, &evmv12.ConstructorParams{}, &evmv12.ConstructorParams{})
	case tools.AnyIsSupported(network, height, tools.V20, tools.V19):
		return parse(raw, nil, false, &evmv11.ConstructorParams{}, &evmv11.ConstructorParams{})
	case tools.V18.IsSupported(network, height):
		return parse(raw, nil, false, &evmv10.ConstructorParams{}, &evmv10.ConstructorParams{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V17)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Evm) GetStorageAt(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parse(rawParams, rawReturn, true, &evmv15.GetStorageAtParams{}, &abi.CborBytes{})
	case tools.V23.IsSupported(network, height):
		return parse(rawParams, rawReturn, true, &evmv14.GetStorageAtParams{}, &abi.CborBytes{})
	case tools.V22.IsSupported(network, height):
		return parse(rawParams, rawReturn, true, &evmv13.GetStorageAtParams{}, &abi.CborBytes{})
	case tools.V21.IsSupported(network, height):
		return parse(rawParams, rawReturn, true, &evmv12.GetStorageAtParams{}, &abi.CborBytes{})
	case tools.AnyIsSupported(network, height, tools.V20, tools.V19):
		return parse(rawParams, rawReturn, true, &evmv11.GetStorageAtParams{}, &abi.CborBytes{})
	case tools.V18.IsSupported(network, height):
		return parse(rawParams, rawReturn, true, &evmv10.GetStorageAtParams{}, &abi.CborBytes{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V17)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}
