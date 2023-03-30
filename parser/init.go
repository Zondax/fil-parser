package parser

import (
	"bytes"
	"encoding/base64"

	"github.com/filecoin-project/go-address"
	builtinInit "github.com/filecoin-project/go-state-types/builtin/v11/init"
	finit "github.com/filecoin-project/go-state-types/builtin/v11/init"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	filInit "github.com/filecoin-project/specs-actors/actors/builtin/init"

	"github.com/zondax/fil-parser/types"
)

func (p *Parser) parseInit(txType string, msg *filTypes.Message, msgRct *filTypes.MessageReceipt) (map[string]interface{}, error) {
	switch txType {
	case MethodSend:
		return p.parseSend(msg), nil
	case MethodConstructor:
		return p.initConstructor(msg.Params)
	case MethodExec:
		return p.parseExec(msg, msgRct)
	case MethodExec4:
		return p.parseExec4(msg, msgRct)
	case UnknownStr:
		return p.unknownMetadata(msg.Params, msgRct.Return)
	}
	return map[string]interface{}{}, errUnknownMethod
}

func (p *Parser) initConstructor(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var constructor builtinInit.ConstructorParams
	err := constructor.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = constructor
	return metadata, nil
}

func (p *Parser) parseExec(msg *filTypes.Message, msgRct *filTypes.MessageReceipt) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(msg.Params)
	var params filInit.ExecParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = execParams{
		CodeCid:           params.CodeCID.String(),
		ConstructorParams: base64.StdEncoding.EncodeToString(params.ConstructorParams),
	}

	reader = bytes.NewReader(msgRct.Return)
	var r finit.ExecReturn
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	createdActorName, err := p.lib.BuiltinActors.GetActorNameFromCid(params.CodeCID)
	if err != nil {
		return metadata, err
	}
	createdActor := types.AddressInfo{
		Short:          r.IDAddress.String(),
		Robust:         r.RobustAddress.String(),
		ActorCid:       params.CodeCID,
		ActorType:      createdActorName,
		CreationTxHash: msg.Cid().String(),
	}
	p.appendToAddresses(createdActor)
	metadata[ReturnKey] = createdActor
	return metadata, nil
}

func (p *Parser) parseExec4(msg *filTypes.Message, msgRct *filTypes.MessageReceipt) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(msg.Params)
	var params finit.Exec4Params
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	subAddress, _ := address.NewFromBytes(params.SubAddress)
	metadata[ParamsKey] = exec4Params{
		CodeCid:           params.CodeCID.String(),
		ConstructorParams: base64.StdEncoding.EncodeToString(params.ConstructorParams),
		SubAddress:        subAddress.String(),
	}

	createdActorName, err := p.lib.BuiltinActors.GetActorNameFromCid(params.CodeCID)
	if err != nil {
		return metadata, err
	}
	var createdActor types.AddressInfo
	reader = bytes.NewReader(msgRct.Return)
	var r finit.Exec4Return
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}

	createdActor = types.AddressInfo{
		Short:          r.IDAddress.String(),
		Robust:         r.RobustAddress.String(),
		ActorCid:       params.CodeCID,
		ActorType:      createdActorName,
		CreationTxHash: msg.Cid().String(),
	}
	metadata[ReturnKey] = createdActor
	p.appendToAddresses(createdActor)
	return metadata, nil
}
