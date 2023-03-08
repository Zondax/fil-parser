package parser

import (
	"bytes"
	"encoding/hex"
	"github.com/filecoin-project/go-state-types/builtin/v10/eam"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/types"
)

func (p *Parser) parseEam(txType string, msg *filTypes.Message, msgRct *filTypes.MessageReceipt, msgCid cid.Cid, ethLogs []types.EthLog) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	switch txType {
	case MethodConstructor:
	case MethodCreate:
		return p.parseCreate(msg.Params, msgRct.Return, msgCid)
	case MethodCreate2:
		return p.parseCreate2(msg.Params, msgRct.Return, msgCid)
	case MethodCreateExternal:
		return p.parseCreateExternal(msg, msgRct, msgCid, ethLogs)
	}
	return metadata, nil
}

func (p *Parser) parseCreate(rawParams, rawReturn []byte, msgCid cid.Cid) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})

	reader := bytes.NewReader(rawParams)
	var params eam.CreateParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params

	reader = bytes.NewReader(rawReturn)
	var createReturn eam.CreateReturn
	err = createReturn.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = eamCreate{
		ActorId:       createReturn.ActorID,
		RobustAddress: createReturn.RobustAddress,
		EthAddress:    ethPrefix + hex.EncodeToString(createReturn.EthAddress[:]),
	}
	p.appendEamAddress(eam.Return(createReturn))

	ethHash, err := ethtypes.EthHashFromCid(msgCid)
	if err != nil {
		return metadata, err
	}
	metadata[ethHashKey] = ethHash.String()

	return metadata, nil
}

func (p *Parser) parseCreate2(rawParams, rawReturn []byte, msgCid cid.Cid) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})

	reader := bytes.NewReader(rawParams)
	var params eam.Create2Params
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params

	reader = bytes.NewReader(rawReturn)
	var createReturn eam.Create2Return
	err = createReturn.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = eamCreate{
		ActorId:       createReturn.ActorID,
		RobustAddress: createReturn.RobustAddress,
		EthAddress:    ethPrefix + hex.EncodeToString(createReturn.EthAddress[:]),
	}
	p.appendEamAddress(eam.Return(createReturn))

	ethHash, err := ethtypes.EthHashFromCid(msgCid)
	if err != nil {
		return metadata, err
	}
	metadata[ethHashKey] = ethHash.String()

	return metadata, nil
}

func (p *Parser) parseCreateExternal(msg *filTypes.Message, msgRct *filTypes.MessageReceipt, msgCid cid.Cid, ethLogs []types.EthLog) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	metadata[ParamsKey] = ethPrefix + hex.EncodeToString(msg.Params[3:]) // TODO

	reader := bytes.NewReader(msgRct.Return)
	var createExternalReturn eam.CreateExternalReturn
	err := createExternalReturn.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = eamCreate{
		ActorId:       createExternalReturn.ActorID,
		RobustAddress: createExternalReturn.RobustAddress,
		EthAddress:    ethPrefix + hex.EncodeToString(createExternalReturn.EthAddress[:]),
	}
	p.appendEamAddress(eam.Return(createExternalReturn))

	ethHash, err := ethtypes.EthHashFromCid(msgCid)
	if err != nil {
		return metadata, err
	}
	metadata[ethHashKey] = ethHash.String()

	// TODO: still needs to get the name
	res := make([]types.EthLog, 0)
	for _, log := range ethLogs {
		if log[addressKey] == ethPrefix+hex.EncodeToString(createExternalReturn.EthAddress[:]) {
			res = append(res, log)
		}
	}
	metadata[ethLogsKey] = res

	return metadata, nil
}

// TODO: add cid and short?
func (p *Parser) appendEamAddress(r eam.Return) {
	p.appendToAddresses(types.AddressInfo{
		Robust:    r.RobustAddress.String(),
		ActorType: "evm",
	})
}
