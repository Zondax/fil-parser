package reward

import (
	"bytes"

	"github.com/zondax/fil-parser/parser"
)

func parse[T rewardParams](raw []byte, params T, key string) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[key] = params
	return metadata, nil
}

// parseWithReturn decodes params and, when the method has a non-empty return, the return value
// too. NV29 (Solstice) added nine FRC-42 reward methods; only ClaimExported returns a value, so
// the rest go through parse() above. Mirrors actors/v2/miner/generic.go's parseGeneric.
func parseWithReturn[T rewardParams, R rewardParams](rawParams, rawReturn []byte, params T, r R) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	if err := params.UnmarshalCBOR(reader); err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	if len(rawReturn) > 0 {
		reader = bytes.NewReader(rawReturn)
		if err := r.UnmarshalCBOR(reader); err != nil {
			return metadata, err
		}
		metadata[parser.ReturnKey] = r
	}
	return metadata, nil
}
