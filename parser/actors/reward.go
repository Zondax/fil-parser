package actors

import (
	"bytes"
	"github.com/filecoin-project/go-state-types/builtin/v11/reward"
	"github.com/zondax/fil-parser/parser"

	"github.com/filecoin-project/go-state-types/abi"
)

func ParseReward(txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt) (map[string]interface{}, error) {
	switch txType {
	case parser.MethodSend:
		return parseSend(msg), nil
	case parser.MethodConstructor:
		return rewardConstructor(msg.Params)
	case parser.MethodAwardBlockReward:
		return awardBlockReward(msg.Params)
	case parser.MethodUpdateNetworkKPI:
		return updateNetworkKpi(msg.Params)
	case parser.MethodThisEpochReward:
		return thisEpochReward(msgRct.Return)
	case parser.UnknownStr:
		return unknownMetadata(msg.Params, msgRct.Return)
	}
	return map[string]interface{}{}, parser.ErrUnknownMethod
}

func rewardConstructor(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params abi.StoragePower
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func awardBlockReward(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var blockRewards reward.AwardBlockRewardParams
	err := blockRewards.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = blockRewards
	return metadata, nil
}

func updateNetworkKpi(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var blockRewards abi.StoragePower
	err := blockRewards.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = blockRewards
	return metadata, nil
}

func thisEpochReward(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var epochRewards reward.ThisEpochRewardReturn
	err := epochRewards.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = epochRewards
	return metadata, nil
}
