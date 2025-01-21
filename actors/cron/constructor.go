package cron

import (
	"bytes"
	"io"

	cronv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/cron"
	cronv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/cron"
	cronv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/cron"
	cronv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/cron"
	cronv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/cron"
	cronv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/cron"
	cronv8 "github.com/filecoin-project/specs-actors/v8/actors/builtin/cron"

	"github.com/zondax/fil-parser/parser"
)

type cronConstructorParams interface {
	UnmarshalCBOR(r io.Reader) error
}

func cronConstructorv8(raw []byte) (map[string]interface{}, error) {
	return cronConstructorGeneric[*cronv8.ConstructorParams](raw, &cronv8.ConstructorParams{})
}

func cronConstructorv7(raw []byte) (map[string]interface{}, error) {
	return cronConstructorGeneric[*cronv7.ConstructorParams](raw, &cronv7.ConstructorParams{})
}

func cronConstructorv6(raw []byte) (map[string]interface{}, error) {
	return cronConstructorGeneric[*cronv6.ConstructorParams](raw, &cronv6.ConstructorParams{})
}

func cronConstructorv5(raw []byte) (map[string]interface{}, error) {
	return cronConstructorGeneric[*cronv5.ConstructorParams](raw, &cronv5.ConstructorParams{})
}

func cronConstructorv4(raw []byte) (map[string]interface{}, error) {
	return cronConstructorGeneric[*cronv4.ConstructorParams](raw, &cronv4.ConstructorParams{})
}

func cronConstructorv3(raw []byte) (map[string]interface{}, error) {
	return cronConstructorGeneric[*cronv3.ConstructorParams](raw, &cronv3.ConstructorParams{})
}

func cronConstructorv2(raw []byte) (map[string]interface{}, error) {
	return cronConstructorGeneric[*cronv2.ConstructorParams](raw, &cronv2.ConstructorParams{})
}

func cronConstructorGeneric[P cronConstructorParams](raw []byte, params P) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}
