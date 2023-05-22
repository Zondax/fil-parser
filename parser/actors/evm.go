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

func ParseEvm(txType string, msg *parser.LotusMessage, msgCid cid.Cid, msgRct *parser.LotusMessageReceipt, ethLogs []types.EthLog) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	switch txType {
	case parser.MethodConstructor:
		return evmConstructor(msg.Params)
	case parser.MethodResurrect: // TODO: not tested
		return resurrect(msg.Params)
	case parser.MethodInvokeContract, parser.MethodInvokeContractReadOnly:
		metadata[parser.ParamsKey] = parser.EthPrefix + hex.EncodeToString(msg.Params)
		metadata[parser.ReturnKey] = parser.EthPrefix + hex.EncodeToString(msgRct.Return)
		logs, err := searchEthLogs(ethLogs, msgCid.String())
		if err != nil {
			return metadata, err
		}
		metadata[parser.EthLogsKey] = logs
	case parser.MethodInvokeContractDelegate:
		return invokeContractDelegate(msg.Params, msgRct.Return)
	case parser.MethodGetBytecode:
		return getByteCode(msgRct.Return)
	case parser.MethodGetBytecodeHash: // TODO: not tested
		return getByteCodeHash(msgRct.Return)
	case parser.MethodGetStorageAt: // TODO: not tested
		return getStorageAt(msg.Params, msgRct.Return)
	case parser.UnknownStr:
		return unknownMetadata(msg.Params, msgRct.Return)
	}
	return metadata, nil
}

func resurrect(raw []byte) (map[string]interface{}, error) {
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

func invokeContractDelegate(rawParams, rawReturn []byte) (map[string]interface{}, error) {
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

func getByteCode(raw []byte) (map[string]interface{}, error) {
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

func getByteCodeHash(raw []byte) (map[string]interface{}, error) {
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

func getStorageAt(rawParams, rawReturn []byte) (map[string]interface{}, error) {
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

func evmConstructor(raw []byte) (map[string]interface{}, error) {
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
