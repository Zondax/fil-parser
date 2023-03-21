package parser

import (
	"bytes"

	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/specs-actors/actors/builtin/reward"
)

func (p *Parser) parseReward(txType string, msg *filTypes.Message, msgRct *filTypes.MessageReceipt) (map[string]interface{}, error) {
	switch txType {
	case MethodSend:
		return p.parseSend(msg), nil
	case MethodAwardBlockReward:
		return p.awardBlockReward(msg.Params)
	case MethodUpdateNetworkKPI:
		return p.updateNerworkKpi(msg.Params)
	case MethodThisEpochReward:
		return p.thisEpochReward(msgRct.Return)
	case UnknownStr:
		return p.unkmownMetadata(msg.Params, msgRct.Return)
	}
	return map[string]interface{}{}, errUnknownMethod
}

func (p *Parser) awardBlockReward(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var blockRewards reward.AwardBlockRewardParams
	err := blockRewards.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = blockRewards
	return metadata, nil
}

func (p *Parser) updateNerworkKpi(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var blockRewards reward.State
	err := blockRewards.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = blockRewards
	return metadata, nil
}

func (p *Parser) thisEpochReward(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var epochRewards reward.ThisEpochRewardReturn
	err := epochRewards.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = epochRewards
	return metadata, nil
}
