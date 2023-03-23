package parser

import (
	"bytes"
	"encoding/hex"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin/v10/evm"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"

	"github.com/zondax/fil-parser/types"
)

func (p *Parser) parseEvm(txType string, msg *filTypes.Message, msgCid cid.Cid, msgRct *filTypes.MessageReceipt, ethLogs []types.EthLog) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	switch txType {
	case MethodConstructor:
		return p.evmConstructor(msg.Params)
	case MethodResurrect: // TODO: not tested
		return p.resurrect(msg.Params)
	case MethodInvokeContract, MethodInvokeContractReadOnly:
		metadata[ParamsKey] = ethPrefix + hex.EncodeToString(msg.Params)
		metadata[ReturnKey] = ethPrefix + hex.EncodeToString(msgRct.Return)
		logs, err := searchEthLogs(ethLogs, msgCid.String())
		if err != nil {
			return metadata, err
		}
		metadata[ethLogsKey] = logs
	case MethodInvokeContractDelegate: // TODO: not tested
		return p.invokeContractDelegate(msg.Params, msgRct.Return)
	case MethodGetBytecode: // TODO: not tested
		return p.getByteCode(msgRct.Return)
	case MethodGetBytecodeHash: // TODO: not tested
		return p.getByteCodeHash(msgRct.Return)
	case MethodGetStorageAt: // TODO: not tested
		return p.getStorageAt(msg.Params, msgRct.Return)
	case UnknownStr:
		return p.unknownMetadata(msg.Params, msgRct.Return)
	}
	return metadata, nil
}

func (p *Parser) resurrect(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params evm.ResurrectParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) invokeContractDelegate(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params evm.DelegateCallParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var r abi.CborBytes
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = r
	return metadata, nil
}

func (p *Parser) getByteCode(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var r evm.GetBytecodeReturn
	err := r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = r
	return metadata, nil
}

func (p *Parser) getByteCodeHash(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var r abi.CborBytes
	err := r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = r
	return metadata, nil
}

func (p *Parser) getStorageAt(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params evm.GetStorageAtParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var r abi.CborBytes
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = r
	return metadata, nil
}

func (p *Parser) evmConstructor(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params evm.ConstructorParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
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
