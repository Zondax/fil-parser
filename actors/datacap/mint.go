package datacap

import (
	"bytes"

	"github.com/filecoin-project/go-state-types/builtin/v10/datacap"
	"github.com/zondax/fil-parser/parser"
)

type (
	mintParams = unmarshalCBOR
	mintReturn = unmarshalCBOR
)

func mintExported(raw, rawReturn []byte, params mintParams, r mintReturn) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params

	reader = bytes.NewReader(rawReturn)
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r
	return metadata, nil
}

func mintExportedv11(raw, rawReturn []byte) (map[string]interface{}, error) {
	return mintExported(raw, rawReturn, &datacap.MintParams{}, &datacap.MintReturn{})
}

func mintExportedv14(raw, rawReturn []byte) (map[string]interface{}, error) {
	return mintExported(raw, rawReturn, &datacap.MintParams{}, &datacap.MintReturn{})
}
