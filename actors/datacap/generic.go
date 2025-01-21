package datacap

import (
	"bytes"

	"github.com/zondax/fil-parser/parser"
)

func parse[T datacapParams, R datacapReturn](raw, rawReturn []byte, customReturn bool) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
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
