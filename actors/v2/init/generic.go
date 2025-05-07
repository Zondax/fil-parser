package init

import (
	"bytes"
	"fmt"

	"github.com/zondax/fil-parser/parser/helper"

	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"

	typegen "github.com/whyrusleeping/cbor-gen"
)

func initConstructor[T typegen.CBORUnmarshaler](raw []byte, constructor T) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	err := constructor.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = constructor
	return metadata, nil
}

func parseExec[T typegen.CBORUnmarshaler, R typegen.CBORUnmarshaler](msg *parser.LotusMessage, rawReturn []byte, params T, r R, h *helper.Helper) (map[string]interface{}, *types.AddressInfo, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(msg.Params)
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, nil, fmt.Errorf("error unmarshaling exec params: %w", err)
	}
	codeCid, tmp, err := setExecParams(params)
	if err != nil {
		return metadata, nil, fmt.Errorf("error parsing exec params: %w", err)
	}
	metadata[parser.ParamsKey] = tmp

	var createdActor *types.AddressInfo
	if len(rawReturn) > 0 {
		reader = bytes.NewReader(rawReturn)
		err = r.UnmarshalCBOR(reader)
		if err != nil {
			return metadata, nil, fmt.Errorf("error unmarshaling exec return: %w", err)
		}
		createdActor = setReturnParams(msg, codeCid.String(), r)
		metadata[parser.ReturnKey] = createdActor
	}

	return metadata, createdActor, nil
}
