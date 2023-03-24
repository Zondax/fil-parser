package parser

import (
	"bytes"
	"encoding/base64"

	"github.com/filecoin-project/go-state-types/builtin/v11/account"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	typegen "github.com/whyrusleeping/cbor-gen"
)

func (p *Parser) parseAccount(txType string, msg *filTypes.Message, msgRct *filTypes.MessageReceipt) (map[string]interface{}, error) {
	switch txType {
	case MethodSend:
		return p.parseSend(msg), nil
	case MethodPubkeyAddress:
		return p.pubkeyAddress(msg.Params)
	case MethodAuthenticateMessage:
		return p.authenticateMessage(msg.Params, msgRct.Return)
	case UnknownStr:
		return p.unknownMetadata(msg.Params, msgRct.Return)
	}
	return map[string]interface{}{}, errUnknownMethod
}

func (p *Parser) pubkeyAddress(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	metadata[ParamsKey] = base64.StdEncoding.EncodeToString(raw)
	return metadata, nil
}

func (p *Parser) authenticateMessage(raw, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params account.AuthenticateMessageParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var expiredReturn typegen.CborBool
	err = expiredReturn.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = expiredReturn
	return metadata, nil
}
