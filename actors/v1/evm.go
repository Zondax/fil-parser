package actors

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/filecoin-project/go-state-types/manifest"

	"github.com/zondax/fil-parser/parser"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin/v11/evm"
)

func (p *ActorParser) ParseEvm(txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	switch txType {
	case parser.MethodConstructor:
		return p.evmConstructor(msg.Params)
	case parser.MethodResurrect: // TODO: not tested
		return p.resurrect(msg.Params)
	case parser.MethodInvokeContract, parser.MethodInvokeContractReadOnly:
		return p.invokeContract(msg.Params, msgRct.Return)
	case parser.MethodInvokeContractDelegate:
		return p.invokeContractDelegate(msg.Params, msgRct.Return)
	case parser.MethodGetBytecode:
		return p.getByteCode(msgRct.Return)
	case parser.MethodGetBytecodeHash: // TODO: not tested
		return p.getByteCodeHash(msgRct.Return)
	case parser.MethodGetStorageAt: // TODO: not tested
		return p.getStorageAt(msg.Params, msgRct.Return)
	case parser.UnknownStr:
		return p.unknownMetadata(msg.Params, msgRct.Return)
	}
	return metadata, nil
}

func (p *ActorParser) resurrect(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params evm.ResurrectParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func (p *ActorParser) invokeContract(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	metadata[parser.ParamsKey] = parser.EthPrefix + hex.EncodeToString(rawParams)
	metadata[parser.ReturnKey] = parser.EthPrefix + hex.EncodeToString(rawReturn)

	var params abi.CborBytes
	if err := params.UnmarshalCBOR(reader); err != nil {
		_ = p.metrics.UpdateActorMethodErrorMetric(manifest.EvmKey, parser.MethodInvokeContract)
		p.logger.Sugar().Warn(fmt.Sprintf("error deserializing rawParams: %s - hex data: %s", err.Error(), hex.EncodeToString(rawParams)))
	}

	if reader.Len() == 0 { // This means that the reader has processed all the bytes
		metadata[parser.ParamsKey] = parser.EthPrefix + hex.EncodeToString(params)
	}

	reader = bytes.NewReader(rawReturn)
	var returnValue abi.CborBytes
	if err := returnValue.UnmarshalCBOR(reader); err != nil {
		_ = p.metrics.UpdateActorMethodErrorMetric(manifest.EvmKey, parser.MethodInvokeContract)
		p.logger.Sugar().Warn(fmt.Sprintf("Error deserializing rawReturn: %s - hex data: %s", err.Error(), hex.EncodeToString(rawReturn)))
	}

	if reader.Len() == 0 { // This means that the reader has processed all the bytes
		metadata[parser.ReturnKey] = parser.EthPrefix + hex.EncodeToString(returnValue)
	}

	return metadata, nil
}

func (p *ActorParser) invokeContractDelegate(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params evm.DelegateCallParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var r abi.CborBytes
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r
	return metadata, nil
}

func (p *ActorParser) getByteCode(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var r evm.GetBytecodeReturn
	err := r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r
	return metadata, nil
}

func (p *ActorParser) getByteCodeHash(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var r abi.CborBytes
	err := r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r
	return metadata, nil
}

func (p *ActorParser) getStorageAt(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params evm.GetStorageAtParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var r abi.CborBytes
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r
	return metadata, nil
}

func (p *ActorParser) evmConstructor(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params evm.ConstructorParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}
