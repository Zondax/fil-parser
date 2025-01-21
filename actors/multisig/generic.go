package multisig

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/zondax/fil-parser/parser"
)

func parseWithMsigParser[T multisigParams, R multisigReturn](msg *parser.LotusMessage,
	height int64,
	key filTypes.TipSetKey,
	fn parseFn,
	rawReturn []byte,
	unmarshaller func(io.Reader, any) error,
	customReturn bool,
) (map[string]interface{}, error) {

	metadata := make(map[string]interface{})
	params, err := fn(msg, height, key)
	if err != nil {
		return map[string]interface{}{}, err
	}
	metadata[parser.ParamsKey] = params

	if customReturn {
		var r R
		err = unmarshaller(bytes.NewReader(rawReturn), &r)
		if err != nil {
			return map[string]interface{}{}, err
		}
		metadata[parser.ReturnKey] = r
	}
	return metadata, nil

}

func parse[T multisigParams, P []byte | string](raw P, unmarshaller func(io.Reader, any) error) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	var params T
	rawBytes, err := toBytes(raw)
	if err != nil {
		return map[string]interface{}{}, err
	}
	reader := bytes.NewReader(rawBytes)
	err = unmarshaller(reader, &params)
	if err != nil {
		return map[string]interface{}{}, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func getValue[T multisigParams](height int64, raw map[string]interface{}) (interface{}, error) {
	paramsRaw, ok := raw["Params"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("Params not found or not a map[string]interface{}")
	}

	var v T
	err := mapToStruct(paramsRaw, &v)
	if err != nil {
		return nil, err
	}
	return v, nil
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
