package reward

import (
	"bytes"
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
