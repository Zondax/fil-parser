package reward

import (
	"bytes"

	"github.com/zondax/fil-parser/parser"
)

func parse[T rewardParams](raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params T
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}
