package power

import (
	"bytes"

	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
)

func parse[T powerParams, R powerReturn](msg *parser.LotusMessage, raw, rawReturn []byte, customReturn bool, params T, r R) (map[string]interface{}, *types.AddressInfo, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, nil, err
	}

	metadata[parser.ParamsKey] = params
	if !customReturn {
		return metadata, nil, nil
	}

	var createdActor *types.AddressInfo
	if len(rawReturn) > 0 {
		reader = bytes.NewReader(rawReturn)
		err = r.UnmarshalCBOR(reader)
		if err != nil {
			return metadata, nil, err
		}
		createdActor = getAddressInfo(r, msg)
		metadata[parser.ReturnKey] = createdActor
	}
	return metadata, createdActor, nil
}
