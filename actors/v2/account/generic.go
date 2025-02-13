package account

import (
	"bytes"
	"fmt"

	"github.com/zondax/fil-parser/parser"
)

func authenticateMessageGeneric[P authenticateMessageParams, R authenticateMessageReturn](raw, rawReturn []byte, params P, r R) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		fmt.Println("error unmarshalling params", len(raw))
		return metadata, fmt.Errorf("error unmarshalling params: %w", err)
	}
	metadata[parser.ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, fmt.Errorf("error unmarshalling return: %w", err)
	}
	metadata[parser.ReturnKey] = r
	return metadata, nil
}
