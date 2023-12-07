package actors

import (
	"github.com/zondax/fil-parser/parser"
)

func (p *ActorParser) ParseEthAccount(_ string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt) (map[string]interface{}, error) {
	// ethaccount can only receive tokens by Send or InvokeEVM methods
	return p.parseEthAccountAny(msg.Params, msgRct.Return)
}

func (p *ActorParser) parseEthAccountAny(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	metadata[parser.ParamsKey] = rawParams
	metadata[parser.ReturnKey] = rawReturn

	return metadata, nil
}
