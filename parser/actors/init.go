package actors

import (
	"bytes"
	"encoding/base64"
	"github.com/zondax/fil-parser/parser"
	"strings"

	"github.com/filecoin-project/go-address"
	builtinInit "github.com/filecoin-project/go-state-types/builtin/v11/init"
	finit "github.com/filecoin-project/go-state-types/builtin/v11/init"
	filInit "github.com/filecoin-project/specs-actors/actors/builtin/init"

	"github.com/zondax/fil-parser/types"
)

func ParseInit(txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt) (map[string]interface{}, error) {
	switch txType {
	case parser.MethodSend:
		return parseSend(msg), nil
	case parser.MethodConstructor:
		return initConstructor(msg.Params)
	case parser.MethodExec:
		return parseExec(msg, msgRct)
	case parser.MethodExec4:
		return parseExec4(msg, msgRct)
	case parser.UnknownStr:
		return unknownMetadata(msg.Params, msgRct.Return)
	}
	return map[string]interface{}{}, parser.ErrUnknownMethod
}

func initConstructor(raw []byte) (map[string]interface{}, error) {
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

func parseExec(msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(msg.Params)
	var params filInit.ExecParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = parser.ExecParams{
		CodeCid:           params.CodeCID.String(),
		ConstructorParams: base64.StdEncoding.EncodeToString(params.ConstructorParams),
	}

	reader = bytes.NewReader(msgRct.Return)
	var r finit.ExecReturn
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	createdActorName, err := p.Lib.BuiltinActors.GetActorNameFromCid(params.CodeCID)
	if err != nil {
		return metadata, err
	}
	createdActor := types.AddressInfo{
		Short:          r.IDAddress.String(),
		Robust:         r.RobustAddress.String(),
		ActorCid:       params.CodeCID,
		ActorType:      parseExecActor(createdActorName),
		CreationTxHash: msg.Cid.String(),
	}
	appendToAddresses(createdActor)
	metadata[parser.ReturnKey] = createdActor
	return metadata, nil
}

func parseExec4(msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(msg.Params)
	var params finit.Exec4Params
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	subAddress, _ := address.NewFromBytes(params.SubAddress)
	metadata[parser.ParamsKey] = parser.Exec4Params{
		CodeCid:           params.CodeCID.String(),
		ConstructorParams: base64.StdEncoding.EncodeToString(params.ConstructorParams),
		SubAddress:        subAddress.String(),
	}

	createdActorName, err := p.Lib.BuiltinActors.GetActorNameFromCid(params.CodeCID)
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
		ActorType:      parseExecActor(createdActorName),
		CreationTxHash: msg.Cid.String(),
	}
	metadata[parser.ReturnKey] = createdActor
	appendToAddresses(createdActor)
	return metadata, nil
}

func parseExecActor(actor string) string {
	s := strings.Split(actor, "/")
	if len(s) < 1 {
		return actor
	}
	return s[len(s)-1]
}
