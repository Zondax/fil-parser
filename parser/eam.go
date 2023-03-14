package parser

import (
	"bytes"
	"encoding/hex"
	"strconv"

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
		return p.parseCreateExternal(msg, msgRct, msgCid)
	}
	return metadata, nil
}

func (p *Parser) parseEamReturn(rawReturn []byte) (cr eam.CreateReturn, err error) {
	reader := bytes.NewReader(rawReturn)
	err = cr.UnmarshalCBOR(reader)
	return cr, err
}

func (p *Parser) newEamCreate(r eam.CreateReturn) eamCreateReturn {
	return eamCreateReturn{
		ActorId:       r.ActorID,
		RobustAddress: r.RobustAddress,
		EthAddress:    ethPrefix + hex.EncodeToString(r.EthAddress[:]),
	}
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

	createReturn, err := p.parseEamReturn(rawReturn)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = p.newEamCreate(createReturn)
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

	createReturn, err := p.parseEamReturn(rawReturn)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = p.newEamCreate(createReturn)
	p.appendEamAddress(eam.Return(createReturn))

	ethHash, err := ethtypes.EthHashFromCid(msgCid)
	if err != nil {
		return metadata, err
	}
	metadata[ethHashKey] = ethHash.String()

	return metadata, nil
}

func (p *Parser) parseCreateExternal(msg *filTypes.Message, msgRct *filTypes.MessageReceipt, msgCid cid.Cid) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	metadata[ParamsKey] = ethPrefix + hex.EncodeToString(msg.Params[3:]) // TODO

	createExternalReturn, err := p.parseEamReturn(msgRct.Return)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = p.newEamCreate(createExternalReturn)
	p.appendEamAddress(eam.Return(createExternalReturn))

	ethHash, err := ethtypes.EthHashFromCid(msgCid)
	if err != nil {
		return metadata, err
	}
	metadata[ethHashKey] = ethHash.String()

	return metadata, nil
}

func (p *Parser) appendEamAddress(r eam.Return) {
	p.appendToAddresses(types.AddressInfo{
		Short:     filPrefix + strconv.FormatUint(r.ActorID, 10),
		Robust:    r.RobustAddress.String(),
		ActorType: "evm",
	})
}
