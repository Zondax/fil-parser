package reward

import (
	"github.com/filecoin-project/go-state-types/abi"
	rewardv11 "github.com/filecoin-project/go-state-types/builtin/v11/reward"
	rewardv8 "github.com/filecoin-project/go-state-types/builtin/v8/reward"
)

func RewardConstructor(height int64, raw []byte) (map[string]interface{}, error) {
	return parse[*abi.StoragePower](raw)
}

func AwardBlockReward(height int64, raw []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		return parse[*rewardv8.AwardBlockRewardParams](raw)
	case 11:
		return parse[*rewardv11.AwardBlockRewardParams](raw)
	}
	return nil, nil
}

func UpdateNetworkKPI(height int64, raw []byte) (map[string]interface{}, error) {
	return parse[*abi.StoragePower](raw)
}

func ThisEpochReward(height int64, raw []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		return parse[*rewardv8.ThisEpochRewardReturn](raw)
	case 11:
		return parse[*rewardv11.ThisEpochRewardReturn](raw)
	}
	return nil, nil
}
