package datacap

import (
	"bytes"

	"github.com/filecoin-project/go-state-types/builtin/v10/datacap"
	"github.com/zondax/fil-parser/parser"
)

type (
	transferParams     = unmarshalCBOR
	transferReturn     = unmarshalCBOR
	transferFromParams = unmarshalCBOR
	transferFromReturn = unmarshalCBOR
)

func transferExported(raw, rawReturn []byte, params transferParams, r transferReturn) (map[string]interface{}, error) {
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

func transferExportedv11(raw, rawReturn []byte) (map[string]interface{}, error) {
	return transferExported(raw, rawReturn, &datacap.TransferParams{}, &datacap.TransferReturn{})
}

func transferExportedv14(raw, rawReturn []byte) (map[string]interface{}, error) {
	return transferExported(raw, rawReturn, &datacap.TransferParams{}, &datacap.TransferReturn{})
}

func transferFromExported(raw, rawReturn []byte, params transferFromParams, r transferFromReturn) (map[string]interface{}, error) {
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

func transferFromExportedv11(raw, rawReturn []byte) (map[string]interface{}, error) {
	return transferFromExported(raw, rawReturn, &datacap.TransferFromParams{}, &datacap.TransferFromReturn{})
}

func transferFromExportedv14(raw, rawReturn []byte) (map[string]interface{}, error) {
	return transferFromExported(raw, rawReturn, &datacap.TransferFromParams{}, &datacap.TransferFromReturn{})
}
