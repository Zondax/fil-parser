package multisig

import (
	"encoding/json"
	"fmt"
	"io"
)

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
