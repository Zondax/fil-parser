package parser

import (
	"bytes"
	reward2 "github.com/filecoin-project/go-state-types/builtin/v11/reward"

	"github.com/filecoin-project/go-state-types/abi"
	filTypes "github.com/filecoin-project/lotus/chain/types"
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
		return p.unknownMetadata(msg.Params, msgRct.Return)
	}
	return map[string]interface{}{}, errUnknownMethod
}

func (p *Parser) awardBlockReward(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var blockRewards reward2.AwardBlockRewardParams
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
	var blockRewards abi.StoragePower
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
	var epochRewards reward2.ThisEpochRewardReturn
	err := epochRewards.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = epochRewards
	return metadata, nil
}
