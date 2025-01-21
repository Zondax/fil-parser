package init

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"

	builtinInitv10 "github.com/filecoin-project/go-state-types/builtin/v10/init"
	builtinInitv11 "github.com/filecoin-project/go-state-types/builtin/v11/init"
	builtinInitv12 "github.com/filecoin-project/go-state-types/builtin/v12/init"
	builtinInitv13 "github.com/filecoin-project/go-state-types/builtin/v13/init"
	builtinInitv14 "github.com/filecoin-project/go-state-types/builtin/v14/init"
	builtinInitv15 "github.com/filecoin-project/go-state-types/builtin/v15/init"
	builtinInitv8 "github.com/filecoin-project/go-state-types/builtin/v8/init"
	builtinInitv9 "github.com/filecoin-project/go-state-types/builtin/v9/init"

	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
)

type constructorParams interface {
	UnmarshalCBOR(io.Reader) error
}

type execReturn interface {
	UnmarshalCBOR(io.Reader) error
}

func execParams(params constructorParams) parser.ExecParams {
	switch v := params.(type) {
	case *builtinInitv15.ExecParams:
		return parser.ExecParams{
			CodeCid:           v.CodeCID.String(),
			ConstructorParams: base64.StdEncoding.EncodeToString(v.ConstructorParams),
		}
	case *builtinInitv14.ExecParams:
		return parser.ExecParams{
			CodeCid:           v.CodeCID.String(),
			ConstructorParams: base64.StdEncoding.EncodeToString(v.ConstructorParams),
		}
	case *builtinInitv13.ExecParams:
		return parser.ExecParams{
			CodeCid:           v.CodeCID.String(),
			ConstructorParams: base64.StdEncoding.EncodeToString(v.ConstructorParams),
		}
	case *builtinInitv12.ExecParams:
		return parser.ExecParams{
			CodeCid:           v.CodeCID.String(),
			ConstructorParams: base64.StdEncoding.EncodeToString(v.ConstructorParams),
		}
	case *builtinInitv11.ExecParams:
		return parser.ExecParams{
			CodeCid:           v.CodeCID.String(),
			ConstructorParams: base64.StdEncoding.EncodeToString(v.ConstructorParams),
		}
	case *builtinInitv10.ExecParams:
		return parser.ExecParams{
			CodeCid:           v.CodeCID.String(),
			ConstructorParams: base64.StdEncoding.EncodeToString(v.ConstructorParams),
		}
	case *builtinInitv9.ExecParams:
		return parser.ExecParams{
			CodeCid:           v.CodeCID.String(),
			ConstructorParams: base64.StdEncoding.EncodeToString(v.ConstructorParams),
		}
	case *builtinInitv8.ExecParams:
		return parser.ExecParams{
			CodeCid:           v.CodeCID.String(),
			ConstructorParams: base64.StdEncoding.EncodeToString(v.ConstructorParams),
		}
	}
	return parser.ExecParams{}
}

func returnParams(msg *parser.LotusMessage, actorCID string, params execReturn) *types.AddressInfo {
	switch v := params.(type) {
	case *builtinInitv15.ExecReturn:
		return &types.AddressInfo{
			Short:         v.IDAddress.String(),
			Robust:        v.RobustAddress.String(),
			ActorCid:      actorCID,
			CreationTxCid: msg.Cid.String(),
		}
	case *builtinInitv14.ExecReturn:
		return &types.AddressInfo{
			Short:         v.IDAddress.String(),
			Robust:        v.RobustAddress.String(),
			ActorCid:      actorCID,
			CreationTxCid: msg.Cid.String(),
		}
	case *builtinInitv13.ExecReturn:
		return &types.AddressInfo{
			Short:         v.IDAddress.String(),
			Robust:        v.RobustAddress.String(),
			ActorCid:      actorCID,
			CreationTxCid: msg.Cid.String(),
		}
	case *builtinInitv12.ExecReturn:
		return &types.AddressInfo{
			Short:         v.IDAddress.String(),
			Robust:        v.RobustAddress.String(),
			ActorCid:      actorCID,
			CreationTxCid: msg.Cid.String(),
		}
	case *builtinInitv11.ExecReturn:
		return &types.AddressInfo{
			Short:         v.IDAddress.String(),
			Robust:        v.RobustAddress.String(),
			ActorCid:      actorCID,
			CreationTxCid: msg.Cid.String(),
		}
	case *builtinInitv10.ExecReturn:
		return &types.AddressInfo{
			Short:         v.IDAddress.String(),
			Robust:        v.RobustAddress.String(),
			ActorCid:      actorCID,
			CreationTxCid: msg.Cid.String(),
		}
	case *builtinInitv9.ExecReturn:
		return &types.AddressInfo{
			Short:         v.IDAddress.String(),
			Robust:        v.RobustAddress.String(),
			ActorCid:      actorCID,
			CreationTxCid: msg.Cid.String(),
		}
	case *builtinInitv8.ExecReturn:
		return &types.AddressInfo{
			Short:         v.IDAddress.String(),
			Robust:        v.RobustAddress.String(),
			ActorCid:      actorCID,
			CreationTxCid: msg.Cid.String(),
		}
	}
	return &types.AddressInfo{}
}

func InitConstructor(height int64, raw []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return initConstructor[*builtinInitv15.ConstructorParams](raw)
	case 14:
		return initConstructor[*builtinInitv14.ConstructorParams](raw)
	case 13:
		return initConstructor[*builtinInitv13.ConstructorParams](raw)
	case 12:
		return initConstructor[*builtinInitv12.ConstructorParams](raw)
	case 11:
		return initConstructor[*builtinInitv11.ConstructorParams](raw)
	case 10:
		return initConstructor[*builtinInitv10.ConstructorParams](raw)
	case 9:
		return initConstructor[*builtinInitv9.ConstructorParams](raw)
	case 8:
		return initConstructor[*builtinInitv8.ConstructorParams](raw)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func ParseExec(height int64, msg *parser.LotusMessage, raw []byte) (map[string]interface{}, *types.AddressInfo, error) {
	switch height {
	case 15:
		return parseExec[*builtinInitv15.ExecParams, *builtinInitv15.ExecReturn](msg, raw)
	case 14:
		return parseExec[*builtinInitv14.ExecParams, *builtinInitv14.ExecReturn](msg, raw)
	case 13:
		return parseExec[*builtinInitv13.ExecParams, *builtinInitv13.ExecReturn](msg, raw)
	case 12:
		return parseExec[*builtinInitv12.ExecParams, *builtinInitv12.ExecReturn](msg, raw)
	case 11:
		return parseExec[*builtinInitv11.ExecParams, *builtinInitv11.ExecReturn](msg, raw)
	case 10:
		return parseExec[*builtinInitv10.ExecParams, *builtinInitv10.ExecReturn](msg, raw)
	case 9:
		return parseExec[*builtinInitv9.ExecParams, *builtinInitv9.ExecReturn](msg, raw)
	case 8:
		return parseExec[*builtinInitv8.ExecParams, *builtinInitv8.ExecReturn](msg, raw)
	}
	return nil, nil, fmt.Errorf("unsupported height: %d", height)
}

func ParseExec4(height int64, msg *parser.LotusMessage, raw []byte) (map[string]interface{}, *types.AddressInfo, error) {
	switch height {
	case 15:
		return parseExec[*builtinInitv15.Exec4Params, *builtinInitv15.Exec4Return](msg, raw)
	case 14:
		return parseExec[*builtinInitv14.Exec4Params, *builtinInitv14.Exec4Return](msg, raw)
	case 13:
		return parseExec[*builtinInitv13.Exec4Params, *builtinInitv13.Exec4Return](msg, raw)
	case 12:
		return parseExec[*builtinInitv12.Exec4Params, *builtinInitv12.Exec4Return](msg, raw)
	case 11:
		return parseExec[*builtinInitv11.Exec4Params, *builtinInitv11.Exec4Return](msg, raw)
	case 10:
		return parseExec[*builtinInitv10.Exec4Params, *builtinInitv10.Exec4Return](msg, raw)
	case 9:
		return nil, nil, fmt.Errorf("unsupported height: %d", height)
	case 8:
		return nil, nil, fmt.Errorf("unsupported height: %d", height)
	}
	return nil, nil, fmt.Errorf("unsupported height: %d", height)
}

func initConstructor[T constructorParams](raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var constructor T
	err := constructor.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = constructor
	return metadata, nil
}

func parseExec[T constructorParams, R execReturn](msg *parser.LotusMessage, rawReturn []byte) (map[string]interface{}, *types.AddressInfo, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(msg.Params)
	var params T
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, nil, err
	}
	tmp := execParams(params)
	metadata[parser.ParamsKey] = tmp

	reader = bytes.NewReader(rawReturn)
	var r R
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, nil, err
	}

	createdActor := returnParams(msg, tmp.CodeCid, r)
	metadata[parser.ReturnKey] = createdActor
	return metadata, createdActor, nil
}
