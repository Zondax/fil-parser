package miner

import (
	"bytes"

	"github.com/zondax/fil-parser/parser"
)

func parseGeneric[T minerParam, R minerReturn](rawParams, rawReturn []byte, customReturn bool) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params T
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	if !customReturn {
		return metadata, nil
	}
	reader = bytes.NewReader(rawReturn)
	var r R
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r
	return metadata, nil
}
