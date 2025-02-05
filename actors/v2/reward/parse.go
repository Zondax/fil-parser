package reward

import (
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
)

func (p *Reward) Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, _ cid.Cid, _ filTypes.TipSetKey) (map[string]interface{}, *types.AddressInfo, error) {
	switch txType {
	case parser.MethodSend:
		// return p.parseSend(msg), nil
	case parser.MethodConstructor:
		resp, err := p.Constructor(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodAwardBlockReward:
		resp, err := p.AwardBlockReward(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodUpdateNetworkKPI:
		resp, err := p.UpdateNetworkKPI(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodThisEpochReward:
		resp, err := p.ThisEpochReward(network, height, msgRct.Return)
		return resp, nil, err
	case parser.UnknownStr:
		// return p.unknownMetadata(msg.Params, msgRct.Return)
	}
	return map[string]interface{}{}, nil, parser.ErrUnknownMethod
}

func (p *Reward) TransactionTypes() map[string]any {
	return map[string]any{
		parser.MethodSend:             nil,
		parser.MethodConstructor:      p.Constructor,
		parser.MethodAwardBlockReward: p.AwardBlockReward,
		parser.MethodUpdateNetworkKPI: p.UpdateNetworkKPI,
		parser.MethodThisEpochReward:  p.ThisEpochReward,
	}
}
