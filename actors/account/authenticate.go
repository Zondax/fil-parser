package account

import (
	"bytes"
	"io"

	accountv10 "github.com/filecoin-project/go-state-types/builtin/v10/account"
	accountv11 "github.com/filecoin-project/go-state-types/builtin/v11/account"
	accountv9 "github.com/filecoin-project/go-state-types/builtin/v11/account"
	accountv12 "github.com/filecoin-project/go-state-types/builtin/v12/account"
	accountv13 "github.com/filecoin-project/go-state-types/builtin/v13/account"
	accountv14 "github.com/filecoin-project/go-state-types/builtin/v14/account"
	accountv15 "github.com/filecoin-project/go-state-types/builtin/v15/account"

	typegen "github.com/whyrusleeping/cbor-gen"
	"github.com/zondax/fil-parser/parser"
)

type authenticateMessageParams interface {
	UnmarshalCBOR(r io.Reader) error
}

func authenticateMessageGeneric[P authenticateMessageParams](raw, rawReturn []byte, params P) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var expiredReturn typegen.CborBool
	err = expiredReturn.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = expiredReturn
	return metadata, nil
}

// Version-specific wrappers

func authenticateMessagev15(raw, rawReturn []byte) (map[string]interface{}, error) {
	return authenticateMessageGeneric[*accountv15.AuthenticateMessageParams](raw, rawReturn, &accountv15.AuthenticateMessageParams{})
}

func authenticateMessagev14(raw, rawReturn []byte) (map[string]interface{}, error) {
	return authenticateMessageGeneric[*accountv14.AuthenticateMessageParams](raw, rawReturn, &accountv14.AuthenticateMessageParams{})
}

func authenticateMessagev13(raw, rawReturn []byte) (map[string]interface{}, error) {
	return authenticateMessageGeneric[*accountv13.AuthenticateMessageParams](raw, rawReturn, &accountv13.AuthenticateMessageParams{})
}

func authenticateMessagev12(raw, rawReturn []byte) (map[string]interface{}, error) {
	return authenticateMessageGeneric[*accountv12.AuthenticateMessageParams](raw, rawReturn, &accountv12.AuthenticateMessageParams{})
}

func authenticateMessagev11(raw, rawReturn []byte) (map[string]interface{}, error) {
	return authenticateMessageGeneric[*accountv11.AuthenticateMessageParams](raw, rawReturn, &accountv11.AuthenticateMessageParams{})
}

func authenticateMessagev10(raw, rawReturn []byte) (map[string]interface{}, error) {
	return authenticateMessageGeneric[*accountv10.AuthenticateMessageParams](raw, rawReturn, &accountv10.AuthenticateMessageParams{})
}

func authenticateMessagev9(raw, rawReturn []byte) (map[string]interface{}, error) {
	return authenticateMessageGeneric[*accountv9.AuthenticateMessageParams](raw, rawReturn, &accountv9.AuthenticateMessageParams{})
}

// func authenticateMessagev8(raw, rawReturn []byte) (map[string]interface{}, error) {
// 	return authenticateMessageGeneric[*accountv8.AuthenticateMessageParams](raw, rawReturn, &accountv8.AuthenticateMessageParams{})
// }
