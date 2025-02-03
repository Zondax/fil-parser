package evm

import (
	"bytes"

	"github.com/zondax/fil-parser/parser"
)

func parse[T evmParams, R evmReturn](rawParams, rawReturn []byte, customReturn bool, params T, r R) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	if len(rawParams) > 0 {
		reader := bytes.NewReader(rawParams)
		err := params.UnmarshalCBOR(reader)
		if err != nil {
			return metadata, err
		}
		metadata[parser.ParamsKey] = params
	}
	if !customReturn {
		return metadata, nil
	}
	if len(rawReturn) > 0 {
		reader := bytes.NewReader(rawReturn)
		err := r.UnmarshalCBOR(reader)
		if err != nil {
			return metadata, err
		}
		metadata[parser.ReturnKey] = r
	}
	return metadata, nil
}
