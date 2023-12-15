package actors

import (
	"bytes"
	"encoding/hex"
	"github.com/zondax/fil-parser/parser"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin/v11/evm"
	"github.com/ipfs/go-cid"

	"github.com/zondax/fil-parser/types"
)

func (p *ActorParser) ParseEvm(txType string, msg *parser.LotusMessage, msgCid cid.Cid, msgRct *parser.LotusMessageReceipt, ethLogs []types.EthLog) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	switch txType {
	case parser.MethodConstructor:
		return p.evmConstructor(msg.Params)
	case parser.MethodResurrect: // TODO: not tested
		return p.resurrect(msg.Params)
	case parser.MethodInvokeContract, parser.MethodInvokeContractReadOnly:
		return p.invokeContract(msg.Params, msgRct.Return, msgCid, ethLogs)
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

func (p *ActorParser) invokeContract(rawParams, rawReturn []byte, msgCid cid.Cid, ethLogs []types.EthLog) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params abi.CborBytes
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params

	metadata[parser.ReturnKey] = parser.EthPrefix + hex.EncodeToString(rawReturn)
	logs, err := searchEthLogs(ethLogs, msgCid.String())
	if err != nil {
		return metadata, err
	}
	metadata[parser.EthLogsKey] = logs

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

func searchEthLogs(logs []types.EthLog, msgCid string) ([]types.EthLog, error) {
	res := make([]types.EthLog, 0)
	for _, log := range logs {
		if log.TransactionCid == msgCid {
			res = append(res, log)
		}
	}
	return res, nil
}
