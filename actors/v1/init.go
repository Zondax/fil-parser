package actors

import (
	"bytes"
	"encoding/base64"
	"strings"

	"github.com/zondax/fil-parser/parser"

	"github.com/filecoin-project/go-address"
	builtinInit "github.com/filecoin-project/go-state-types/builtin/v11/init"
	finit "github.com/filecoin-project/go-state-types/builtin/v11/init"
	filInit "github.com/filecoin-project/specs-actors/actors/builtin/init"

	"github.com/zondax/fil-parser/types"
)

func (p *ActorParser) ParseInit(txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt) (map[string]interface{}, *types.AddressInfo, error) {
	var err error
	metadata := make(map[string]interface{})
	switch txType {
	case parser.MethodSend:
		metadata, err = p.parseSend(msg), nil
	case parser.MethodConstructor:
		metadata, err = p.initConstructor(msg.Params)
	case parser.MethodExec:
		return p.parseExec(msg, msgRct.Return)
	case parser.MethodExec4:
		return p.parseExec4(msg, msgRct.Return)
	case parser.UnknownStr:
		metadata, err = p.unknownMetadata(msg.Params, msgRct.Return)
	default:
		err = parser.ErrUnknownMethod
	}
	return metadata, nil, err
}

func (p *ActorParser) initConstructor(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var constructor builtinInit.ConstructorParams
	err := constructor.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = constructor
	return metadata, nil
}

func (p *ActorParser) parseExec(msg *parser.LotusMessage, rawReturn []byte) (map[string]interface{}, *types.AddressInfo, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(msg.Params)
	var params filInit.ExecParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, nil, err
	}
	metadata[parser.ParamsKey] = parser.ExecParams{
		CodeCid:           params.CodeCID.String(),
		ConstructorParams: base64.StdEncoding.EncodeToString(params.ConstructorParams),
	}

	reader = bytes.NewReader(rawReturn)
	var r finit.ExecReturn
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, nil, err
	}
	createdActorName, err := p.helper.GetFilecoinLib().BuiltinActors.GetActorNameFromCid(params.CodeCID)
	if err != nil {
		return metadata, nil, err
	}
	createdActor := &types.AddressInfo{
		Short:         r.IDAddress.String(),
		Robust:        r.RobustAddress.String(),
		ActorCid:      params.CodeCID.String(),
		ActorType:     parseExecActor(createdActorName),
		CreationTxCid: msg.Cid.String(),
	}
	metadata[parser.ReturnKey] = createdActor

	p.helper.GetActorsCache().StoreAddressInfoAddress(*createdActor)

	return metadata, createdActor, nil
}

func (p *ActorParser) parseExec4(msg *parser.LotusMessage, rawReturn []byte) (map[string]interface{}, *types.AddressInfo, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(msg.Params)
	var params finit.Exec4Params
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, nil, err
	}
	subAddress, _ := address.NewFromBytes(params.SubAddress)
	metadata[parser.ParamsKey] = parser.Exec4Params{
		CodeCid:           params.CodeCID.String(),
		ConstructorParams: base64.StdEncoding.EncodeToString(params.ConstructorParams),
		SubAddress:        subAddress.String(),
	}

	createdActorName, err := p.helper.GetFilecoinLib().BuiltinActors.GetActorNameFromCid(params.CodeCID)
	if err != nil {
		return metadata, nil, err
	}
	reader = bytes.NewReader(rawReturn)
	var r finit.Exec4Return
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, nil, err
	}

	createdActor := &types.AddressInfo{
		Short:         r.IDAddress.String(),
		Robust:        r.RobustAddress.String(),
		ActorCid:      params.CodeCID.String(),
		ActorType:     parseExecActor(createdActorName),
		CreationTxCid: msg.Cid.String(),
	}
	metadata[parser.ReturnKey] = createdActor

	p.helper.GetActorsCache().StoreAddressInfoAddress(*createdActor)

	return metadata, createdActor, nil
}

func parseExecActor(actor string) string {
	s := strings.Split(actor, "/")
	if len(s) < 1 {
		return actor
	}
	return s[len(s)-1]
}
