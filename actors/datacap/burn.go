package datacap

import (
	"bytes"

	"github.com/filecoin-project/go-state-types/builtin/v10/datacap"
	"github.com/zondax/fil-parser/parser"
)

type (
	burnParams     = unmarshalCBOR
	burnReturn     = unmarshalCBOR
	destroyParams  = unmarshalCBOR
	burnFromParams = unmarshalCBOR
	burnFromReturn = unmarshalCBOR
)

func burnExported(raw, rawReturn []byte, params burnParams, r burnReturn) (map[string]interface{}, error) {
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

func burnExportedv11(raw, rawReturn []byte) (map[string]interface{}, error) {
	return burnExported(raw, rawReturn, &datacap.BurnParams{}, &datacap.BurnReturn{})
}

func burnExportedv14(raw, rawReturn []byte) (map[string]interface{}, error) {
	return burnExported(raw, rawReturn, &datacap.BurnParams{}, &datacap.BurnReturn{})
}

func destroyExported(raw, rawReturn []byte, params destroyParams, r burnReturn) (map[string]interface{}, error) {
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

func destroyExportedv11(raw, rawReturn []byte) (map[string]interface{}, error) {
	return destroyExported(raw, rawReturn, &datacap.DestroyParams{}, &datacap.BurnReturn{})
}

func destroyExportedv14(raw, rawReturn []byte) (map[string]interface{}, error) {
	return destroyExported(raw, rawReturn, &datacap.DestroyParams{}, &datacap.BurnReturn{})
}

func burnFromExported(raw, rawReturn []byte, params burnFromParams, r burnFromReturn) (map[string]interface{}, error) {
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

func burnFromExportedv11(raw, rawReturn []byte) (map[string]interface{}, error) {
	return burnFromExported(raw, rawReturn, &datacap.BurnFromParams{}, &datacap.BurnFromReturn{})
}

func burnFromExportedv14(raw, rawReturn []byte) (map[string]interface{}, error) {
	return burnFromExported(raw, rawReturn, &datacap.BurnFromParams{}, &datacap.BurnFromReturn{})
}
