package datacap

import (
	"bytes"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin/v10/datacap"
	"github.com/zondax/fil-parser/parser"
)

type (
	increaseAllowanceParams = unmarshalCBOR
	decreaseAllowanceParams = unmarshalCBOR
	revokeAllowanceParams   = unmarshalCBOR
	allowanceParams         = unmarshalCBOR
)

func increaseAllowanceExportedv11(raw, rawReturn []byte) (map[string]interface{}, error) {
	return increaseAllowanceExported(raw, rawReturn, &datacap.IncreaseAllowanceParams{})
}

func increaseAllowanceExportedv14(raw, rawReturn []byte) (map[string]interface{}, error) {
	return increaseAllowanceExported(raw, rawReturn, &datacap.IncreaseAllowanceParams{})
}

func decreaseAllowanceExportedv11(raw, rawReturn []byte) (map[string]interface{}, error) {
	return decreaseAllowanceExported(raw, rawReturn, &datacap.DecreaseAllowanceParams{})
}

func decreaseAllowanceExportedv14(raw, rawReturn []byte) (map[string]interface{}, error) {
	return decreaseAllowanceExported(raw, rawReturn, &datacap.DecreaseAllowanceParams{})
}

func revokeAllowanceExportedv11(raw, rawReturn []byte) (map[string]interface{}, error) {
	return revokeAllowanceExported(raw, rawReturn, &datacap.RevokeAllowanceParams{})
}

func revokeAllowanceExportedv14(raw, rawReturn []byte) (map[string]interface{}, error) {
	return revokeAllowanceExported(raw, rawReturn, &datacap.RevokeAllowanceParams{})
}

func allowanceExportedv11(raw, rawReturn []byte) (map[string]interface{}, error) {
	return allowanceExported(raw, rawReturn, &datacap.GetAllowanceParams{})
}

func allowanceExportedv14(raw, rawReturn []byte) (map[string]interface{}, error) {
	return allowanceExported(raw, rawReturn, &datacap.GetAllowanceParams{})
}

func increaseAllowanceExported(raw, rawReturn []byte, params increaseAllowanceParams) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params

	reader = bytes.NewReader(rawReturn)
	var r abi.TokenAmount
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r // TODO: .uint64()??
	return metadata, nil
}

func decreaseAllowanceExported(raw, rawReturn []byte, params decreaseAllowanceParams) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params

	reader = bytes.NewReader(rawReturn)
	var r abi.TokenAmount
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r // TODO: .uint64()??
	return metadata, nil
}

func revokeAllowanceExported(raw, rawReturn []byte, params revokeAllowanceParams) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params

	reader = bytes.NewReader(rawReturn)
	var r abi.TokenAmount
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r // TODO: .uint64()??
	return metadata, nil
}

func allowanceExported(raw, rawReturn []byte, params allowanceParams) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params

	reader = bytes.NewReader(rawReturn)
	var r abi.TokenAmount
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r // TODO: .uint64()??
	return metadata, nil
}
