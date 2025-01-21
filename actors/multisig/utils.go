package multisig

import (
	"encoding/json"
	"errors"
)

func toBytes(raw any) ([]byte, error) {
	switch v := raw.(type) {
	case []byte:
		return v, nil
	case string:
		return []byte(v), nil
	}
	return nil, errors.New("invalid type")
}

func mapToStruct(m map[string]interface{}, v interface{}) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}
