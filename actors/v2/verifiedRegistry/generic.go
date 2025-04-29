package verifiedRegistry

import (
	"bytes"
	"fmt"

	"github.com/zondax/fil-parser/parser"
)

func parse[T verifiedRegistryParams, R verifiedRegistryReturn](raw, rawReturn []byte, customReturn bool, params T, r R) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	if len(raw) > 0 {
		reader := bytes.NewReader(raw)
		err := params.UnmarshalCBOR(reader)
		if err != nil {
			return metadata, fmt.Errorf("error unmarshaling params: %w", err)
		}
		metadata[parser.ParamsKey] = params
	}
	metadata[parser.ParamsKey] = params
	if !customReturn {
		return metadata, nil
	}
	if len(rawReturn) > 0 {
		reader := bytes.NewReader(rawReturn)
		err := r.UnmarshalCBOR(reader)
		if err != nil {
			return metadata, fmt.Errorf("error unmarshaling return: %w", err)
		}
		metadata[parser.ReturnKey] = r
	}
	return metadata, nil
}
