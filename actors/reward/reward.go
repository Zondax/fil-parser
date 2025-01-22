package reward

import (
	"github.com/filecoin-project/go-state-types/abi"
	rewardv10 "github.com/filecoin-project/go-state-types/builtin/v10/reward"
	rewardv11 "github.com/filecoin-project/go-state-types/builtin/v11/reward"
	rewardv12 "github.com/filecoin-project/go-state-types/builtin/v12/reward"
	rewardv13 "github.com/filecoin-project/go-state-types/builtin/v13/reward"
	rewardv14 "github.com/filecoin-project/go-state-types/builtin/v14/reward"
	rewardv15 "github.com/filecoin-project/go-state-types/builtin/v15/reward"
	rewardv8 "github.com/filecoin-project/go-state-types/builtin/v8/reward"
	rewardv9 "github.com/filecoin-project/go-state-types/builtin/v9/reward"
	"github.com/zondax/fil-parser/tools"
)

func RewardConstructor(network string, height int64, raw []byte) (map[string]interface{}, error) {
	return parse[*abi.StoragePower](raw)
}

func AwardBlockReward(network string, height int64, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(network, height):
		return parse[*rewardv8.AwardBlockRewardParams](raw)
	case tools.V9.IsSupported(network, height):
		return parse[*rewardv9.AwardBlockRewardParams](raw)
	case tools.V10.IsSupported(network, height):
		return parse[*rewardv10.AwardBlockRewardParams](raw)
	case tools.V11.IsSupported(network, height):
		return parse[*rewardv11.AwardBlockRewardParams](raw)
	case tools.V12.IsSupported(network, height):
		return parse[*rewardv12.AwardBlockRewardParams](raw)
	case tools.V13.IsSupported(network, height):
		return parse[*rewardv13.AwardBlockRewardParams](raw)
	case tools.V14.IsSupported(network, height):
		return parse[*rewardv14.AwardBlockRewardParams](raw)
	case tools.V15.IsSupported(network, height):
		return parse[*rewardv15.AwardBlockRewardParams](raw)
	}
	return nil, nil
}

func UpdateNetworkKPI(network string, height int64, raw []byte) (map[string]interface{}, error) {
	return parse[*abi.StoragePower](raw)
}

func ThisEpochReward(network string, height int64, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(network, height):
		return parse[*rewardv8.ThisEpochRewardReturn](raw)
	case tools.V9.IsSupported(network, height):
		return parse[*rewardv9.ThisEpochRewardReturn](raw)
	case tools.V10.IsSupported(network, height):
		return parse[*rewardv10.ThisEpochRewardReturn](raw)
	case tools.V11.IsSupported(network, height):
		return parse[*rewardv11.ThisEpochRewardReturn](raw)
	case tools.V12.IsSupported(network, height):
		return parse[*rewardv12.ThisEpochRewardReturn](raw)
	case tools.V13.IsSupported(network, height):
		return parse[*rewardv13.ThisEpochRewardReturn](raw)
	case tools.V14.IsSupported(network, height):
		return parse[*rewardv14.ThisEpochRewardReturn](raw)
	case tools.V15.IsSupported(network, height):
		return parse[*rewardv15.ThisEpochRewardReturn](raw)
	}
	return nil, nil
}
