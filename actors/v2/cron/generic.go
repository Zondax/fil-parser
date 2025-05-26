package cron

import (
	"bytes"

	cbg "github.com/whyrusleeping/cbor-gen"
	"github.com/zondax/fil-parser/parser"
)

func cronConstructorLegacy[P cbg.CBORUnmarshaler](raw []byte, params P) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}
