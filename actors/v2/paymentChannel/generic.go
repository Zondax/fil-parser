package paymentChannel

import (
	"bytes"

	"github.com/zondax/fil-parser/parser"
)

func parse[T paymentChannelParams](raw []byte, constructor T) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	err := constructor.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = constructor
	return metadata, nil
}
