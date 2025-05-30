package multisig

import (
	"bytes"

	"github.com/zondax/fil-parser/parser"
)

func parseCBOR[T multisigParams](raw, rawReturn []byte, params, ret T) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	if len(raw) > 0 {
		reader := bytes.NewReader(raw)
		err := params.UnmarshalCBOR(reader)
		if err != nil {
			return map[string]interface{}{}, err
		}
		metadata[parser.ParamsKey] = params
	}
	if len(rawReturn) > 0 {
		reader := bytes.NewReader(rawReturn)
		err := ret.UnmarshalCBOR(reader)
		if err != nil {
			return map[string]interface{}{}, err
		}
		metadata[parser.ReturnKey] = ret
	}
	return metadata, nil
}
