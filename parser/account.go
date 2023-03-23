package parser

import (
	"encoding/base64"

	filTypes "github.com/filecoin-project/lotus/chain/types"
)

func (p *Parser) parseAccount(txType string, msg *filTypes.Message, msgRct *filTypes.MessageReceipt) (map[string]interface{}, error) {
	switch txType {
	case MethodSend:
		return p.parseSend(msg), nil
	case "PubkeyAddress":
		metadata := make(map[string]interface{})
		metadata[ParamsKey] = base64.StdEncoding.EncodeToString(msg.Params)
		return metadata, nil
	case UnknownStr:
		return p.unknownMetadata(msg.Params, msgRct.Return)
	}
	return map[string]interface{}{}, errUnknownMethod
}
