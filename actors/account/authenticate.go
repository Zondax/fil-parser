package account

import (
	"bytes"

	accountv11 "github.com/filecoin-project/go-state-types/builtin/v11/account"
	accountv14 "github.com/filecoin-project/go-state-types/builtin/v14/account"
	typegen "github.com/whyrusleeping/cbor-gen"
	"github.com/zondax/fil-parser/parser"
)

type authenticateMessageParams interface {
	UnmarshalCBOR(r *bytes.Reader) error
}

func authenticateMessage(raw, rawReturn []byte, params authenticateMessageParams) (map[string]interface{}, error) {
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
func authenticateMessagev11(raw, rawReturn []byte) (map[string]interface{}, error) {
	params := &accountv11.AuthenticateMessageParams{}
	return authenticateMessage(raw, rawReturn, params)
}

func authenticateMessagev14(raw, rawReturn []byte) (map[string]interface{}, error) {
	params := &accountv14.AuthenticateMessageParams{}
	return authenticateMessage(raw, rawReturn, params)
}
