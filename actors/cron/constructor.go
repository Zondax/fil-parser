package cron

import (
	"bytes"
	"io"

	cronv8 "github.com/filecoin-project/specs-actors/v8/actors/builtin/cron"

	"github.com/zondax/fil-parser/parser"
)

type cronConstructorParams interface {
	UnmarshalCBOR(r io.Reader) error
}

func cronConstructorv8(raw []byte) (map[string]interface{}, error) {
	params := &cronv8.ConstructorParams{}
	return cronConstructor(raw, params)
}

func cronConstructor(raw []byte, params cronConstructorParams) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

// func cronConstructorv11(raw []byte) (map[string]interface{}, error) {
// 	params := &cronv11.ConstructorParams{}
// 	return cronConstructor(raw, params)
// }
// func cronConstructorv14(raw []byte) (map[string]interface{}, error) {
// 	params := &cronv14.ConstructorParams{}
// 	return cronConstructor(raw, params)
// }
