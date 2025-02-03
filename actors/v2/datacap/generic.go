package datacap

import (
	"bytes"

	"github.com/zondax/fil-parser/parser"
)

func parse[T datacapParams, R datacapReturn](raw, rawReturn []byte, customReturn bool, params T, r R) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)

	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params

	if !customReturn {
		return metadata, nil
	}

	reader = bytes.NewReader(rawReturn)

	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r
	return metadata, nil
}
