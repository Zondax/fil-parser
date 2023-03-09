package parser

import (
	"bytes"
	"encoding/hex"
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
	case MethodInvokeContract, MethodInvokeContractReadOnly, MethodInvokeContractDelegate:
		metadata[ParamsKey] = ethPrefix + hex.EncodeToString(msg.Params)
		metadata[ReturnKey] = ethPrefix + hex.EncodeToString(msgRct.Return)
		logs, err := searchEthLogs(ethLogs, msgCid.String())
		if err != nil {
			return metadata, err
		}
		metadata[ethLogsKey] = logs
	case MethodGetBytecode:
	}
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
