package account

import (
	"bytes"
	"fmt"

	typegen "github.com/whyrusleeping/cbor-gen"
	"github.com/zondax/fil-parser/parser"
)

func authenticateMessageGeneric[P typegen.CBORUnmarshaler, R typegen.CBORUnmarshaler](raw, rawReturn []byte, params P, r R) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, fmt.Errorf("error unmarshalling params: %w", err)
	}
	metadata[parser.ParamsKey] = params
	if len(rawReturn) > 0 {
		reader = bytes.NewReader(rawReturn)
		err = r.UnmarshalCBOR(reader)
		if err != nil {
			return metadata, fmt.Errorf("error unmarshalling return: %w", err)
		}
		metadata[parser.ReturnKey] = r
	}
	return metadata, nil
}
