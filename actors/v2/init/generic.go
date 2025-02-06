package init

import (
	"bytes"

	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
)

func initConstructor[T constructorParams](raw []byte, constructor T) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	err := constructor.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = constructor
	return metadata, nil
}

func parseExec[T constructorParams, R execReturn](msg *parser.LotusMessage, rawReturn []byte, params T, r R) (map[string]interface{}, *types.AddressInfo, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(msg.Params)
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, nil, err
	}
	codeCid, tmp, err := execParams(params)
	if err != nil {
		return metadata, nil, err
	}
	metadata[parser.ParamsKey] = tmp

	var createdActor *types.AddressInfo
	if len(rawReturn) > 0 {
		reader = bytes.NewReader(rawReturn)
		err = r.UnmarshalCBOR(reader)
		if err != nil {
			return metadata, nil, err
		}
		createdActor = returnParams(msg, codeCid.String(), r)
		metadata[parser.ReturnKey] = createdActor
	}

	return metadata, createdActor, nil
}
