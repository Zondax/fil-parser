package init

import (
	"bytes"

	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
)

func initConstructor[T constructorParams](raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var constructor T
	err := constructor.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = constructor
	return metadata, nil
}

func parseExec[T constructorParams, R execReturn](msg *parser.LotusMessage, rawReturn []byte) (map[string]interface{}, *types.AddressInfo, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(msg.Params)
	var params T
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, nil, err
	}
	tmp := execParams(params)
	metadata[parser.ParamsKey] = tmp

	reader = bytes.NewReader(rawReturn)
	var r R
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, nil, err
	}

	createdActor := returnParams(msg, tmp.CodeCid, r)
	metadata[parser.ReturnKey] = createdActor
	return metadata, createdActor, nil
}
