package actors

import (
	"bytes"
	"encoding/base64"
	"github.com/zondax/fil-parser/parser"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/builtin/v11/account"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	typegen "github.com/whyrusleeping/cbor-gen"
)

func ParseAccount(txType string, msg *parser.LotusMessage, msgRct *filTypes.MessageReceipt) (map[string]interface{}, error) {
	switch txType {
	case parser.MethodSend:
		return parseSend(msg), nil
	case parser.MethodConstructor:
		return parseConstructor(msg.Params)
	case parser.MethodPubkeyAddress:
		return pubkeyAddress(msg.Params, msgRct.Return)
	case parser.MethodAuthenticateMessage:
		return authenticateMessage(msg.Params, msgRct.Return)
	case parser.UnknownStr:
		return unknownMetadata(msg.Params, msgRct.Return)
	}
	return map[string]interface{}{}, parser.ErrUnknownMethod
}

func pubkeyAddress(raw, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	metadata[parser.ParamsKey] = base64.StdEncoding.EncodeToString(raw)
	reader := bytes.NewReader(rawReturn)
	var r address.Address
	err := r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r.String()
	return metadata, nil
}

func authenticateMessage(raw, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params account.AuthenticateMessageParams
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
