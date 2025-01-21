package power

import (
	"bytes"

	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
)

func parse[T powerParams, R powerReturn](msg *parser.LotusMessage, raw, rawReturn []byte, customReturn bool) (map[string]interface{}, *types.AddressInfo, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var constructor T
	err := constructor.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, nil, err
	}

	metadata[parser.ParamsKey] = constructor
	if !customReturn {
		return metadata, nil, nil
	}

	reader = bytes.NewReader(rawReturn)
	var r R
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, nil, err
	}
	createdActor := getAddressInfo(r, msg)
	metadata[parser.ReturnKey] = createdActor
	return metadata, createdActor, nil
}
