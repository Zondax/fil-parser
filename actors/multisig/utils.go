package multisig

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/zondax/fil-parser/parser"
)

type multisigParams interface {
	UnmarshalCBOR(io.Reader) error
}

type multisigReturn interface {
	UnmarshalCBOR(io.Reader) error
}

type parseFn func(*parser.LotusMessage, int64, filTypes.TipSetKey) (string, error)

type metadataWithCbor map[string]interface{}

func (m metadataWithCbor) UnmarshalCBOR(reader io.Reader) error {
	return nil
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

func toBytes(raw any) ([]byte, error) {
	switch v := raw.(type) {
	case []byte:
		return v, nil
	case string:
		return []byte(v), nil
	}
	return nil, errors.New("invalid type")
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
