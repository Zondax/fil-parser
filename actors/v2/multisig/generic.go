package multisig

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/zondax/fil-parser/parser"
)

func parseWithMsigParser[R multisigReturn](msg *parser.LotusMessage,
	height int64,
	key filTypes.TipSetKey,
	fn ParseFn,
	rawReturn []byte,
	unmarshaller func(io.Reader, any) error,
	customReturn bool,
	r R,
) (map[string]interface{}, error) {

	metadata := make(map[string]interface{})
	if msg == nil || msg.To.Empty() {
		return map[string]interface{}{}, fmt.Errorf("invalid message")
	}
	params, err := fn(msg, height, key)
	if err != nil {
		return map[string]interface{}{}, err
	}
	metadata[parser.ParamsKey] = params

	if customReturn {
		err = unmarshaller(bytes.NewReader(rawReturn), &r)
		if err != nil {
			return map[string]interface{}{}, err
		}
		metadata[parser.ReturnKey] = r
	}
	return metadata, nil

}

func parse[T multisigParams, P []byte | string](raw P, params T, unmarshaller func(io.Reader, any) error) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	rawBytes, err := toBytes(raw)
	if err != nil {
		return map[string]interface{}{}, err
	}
	reader := bytes.NewReader(rawBytes)
	err = unmarshaller(reader, params)
	if err != nil {
		return map[string]interface{}{}, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

// parseConstructorValue, parseLockBalanceValue, parseRemoveSignerValue, , parseSwapSignerValue, parseUniversalReceiverHookValu
func getValue[T multisigParams](height int64, raw map[string]interface{}, v T) (interface{}, error) {
	paramsRaw, ok := raw[parser.ParamsKey]
	if !ok {
		return nil, fmt.Errorf("Params not found")
	}
	switch p := paramsRaw.(type) {
	case map[string]interface{}:
		err := mapToStruct(p, v)
		if err != nil {
			return nil, err
		}
	case string:
		err := json.Unmarshal([]byte(p), v)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("Params not a map[string]interface{} or string")
	}

	return v, nil
}

func parseValue(raw string, v any) (interface{}, error) {
	err := json.Unmarshal([]byte(raw), &v)
	return v, err
}

func jsonUnmarshaller[T multisigParams](reader io.Reader, to any) error {
	err := json.NewDecoder(reader).Decode(to)
	if err != nil {
		return err
	}
	return nil
}

func cborUnmarshaller[T multisigParams](reader io.Reader, to any) error {
	return to.(T).UnmarshalCBOR(reader)
}

func noopUnmarshaller[T multisigParams](reader io.Reader, to any) error {
	return nil
}
