package reward

import (
	"github.com/zondax/fil-parser/parser"
)

func (p *Reward) Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt) (map[string]interface{}, error) {
	switch txType {
	case parser.MethodSend:
		// return p.parseSend(msg), nil
	case parser.MethodConstructor:
		return p.Constructor(network, height, msg.Params)
	case parser.MethodAwardBlockReward:
		return p.AwardBlockReward(network, height, msg.Params)
	case parser.MethodUpdateNetworkKPI:
		return p.UpdateNetworkKPI(network, height, msg.Params)
	case parser.MethodThisEpochReward:
		return p.ThisEpochReward(network, height, msgRct.Return)
	case parser.UnknownStr:
		// return p.unknownMetadata(msg.Params, msgRct.Return)
	}
	return map[string]interface{}{}, parser.ErrUnknownMethod
}
