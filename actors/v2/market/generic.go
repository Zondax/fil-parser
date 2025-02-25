package market

import (
	"bytes"

	"github.com/zondax/fil-parser/parser"
)

func parseGeneric[T marketParam, R marketReturn](rawParams, rawReturn []byte, returnCustomParam bool, params T, r R) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}

	metadata[parser.ParamsKey] = params
	if !returnCustomParam {
		if addr := getAddressAsString(params); addr != "" {
			metadata[parser.ParamsKey] = addr
		}
		return metadata, nil
	}
	if len(rawReturn) > 0 {
		reader = bytes.NewReader(rawReturn)
		err = r.UnmarshalCBOR(reader)
		if err != nil {
			return metadata, err
		}
		metadata[parser.ReturnKey] = r
	}
	return metadata, nil
}
