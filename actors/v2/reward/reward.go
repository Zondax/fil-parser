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
	"github.com/filecoin-project/go-state-types/manifest"
	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/reward"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/reward"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/reward"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/reward"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/reward"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/reward"
	"github.com/zondax/fil-parser/tools"
)

type Reward struct{}

func (r *Reward) Name() string {
	return manifest.RewardKey
}

func (*Reward) Constructor(network string, height int64, raw []byte) (map[string]interface{}, error) {
	return parse(raw, &abi.StoragePower{})
}

func (*Reward) AwardBlockReward(network string, height int64, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parse(raw, &rewardv15.AwardBlockRewardParams{})
	case tools.V23.IsSupported(network, height):
		return parse(raw, &rewardv14.AwardBlockRewardParams{})
	case tools.V22.IsSupported(network, height):
		return parse(raw, &rewardv13.AwardBlockRewardParams{})
	case tools.V21.IsSupported(network, height):
		return parse(raw, &rewardv12.AwardBlockRewardParams{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parse(raw, &rewardv11.AwardBlockRewardParams{})
	case tools.V18.IsSupported(network, height):
		return parse(raw, &rewardv10.AwardBlockRewardParams{})
	case tools.V17.IsSupported(network, height):
		return parse(raw, &rewardv9.AwardBlockRewardParams{})
	case tools.V16.IsSupported(network, height):
		return parse(raw, &rewardv8.AwardBlockRewardParams{})
	case tools.V15.IsSupported(network, height):
		return parse(raw, &legacyv7.AwardBlockRewardParams{})
	case tools.V14.IsSupported(network, height):
		return parse(raw, &legacyv6.AwardBlockRewardParams{})
	case tools.V13.IsSupported(network, height):
		return parse(raw, &legacyv5.AwardBlockRewardParams{})
	case tools.V12.IsSupported(network, height):
		return parse(raw, &legacyv4.AwardBlockRewardParams{})
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parse(raw, &legacyv3.AwardBlockRewardParams{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V9)...):
		return parse(raw, &legacyv2.AwardBlockRewardParams{})
	}
	return nil, nil
}

func (*Reward) UpdateNetworkKPI(network string, height int64, raw []byte) (map[string]interface{}, error) {
	return parse(raw, &abi.StoragePower{})
}

func (*Reward) ThisEpochReward(network string, height int64, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parse(raw, &rewardv15.ThisEpochRewardReturn{})
	case tools.V23.IsSupported(network, height):
		return parse(raw, &rewardv14.ThisEpochRewardReturn{})
	case tools.V22.IsSupported(network, height):
		return parse(raw, &rewardv13.ThisEpochRewardReturn{})
	case tools.V21.IsSupported(network, height):
		return parse(raw, &rewardv12.ThisEpochRewardReturn{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parse(raw, &rewardv11.ThisEpochRewardReturn{})
	case tools.V18.IsSupported(network, height):
		return parse(raw, &rewardv10.ThisEpochRewardReturn{})
	case tools.V17.IsSupported(network, height):
		return parse(raw, &rewardv9.ThisEpochRewardReturn{})
	case tools.V16.IsSupported(network, height):
		return parse(raw, &rewardv8.ThisEpochRewardReturn{})
	case tools.V15.IsSupported(network, height):
		return parse(raw, &legacyv7.ThisEpochRewardReturn{})
	case tools.V14.IsSupported(network, height):
		return parse(raw, &legacyv6.ThisEpochRewardReturn{})
	case tools.V13.IsSupported(network, height):
		return parse(raw, &legacyv5.ThisEpochRewardReturn{})
	case tools.V12.IsSupported(network, height):
		return parse(raw, &legacyv4.ThisEpochRewardReturn{})
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parse(raw, &legacyv3.ThisEpochRewardReturn{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V9)...):
		return parse(raw, &legacyv2.ThisEpochRewardReturn{})
	}
	return nil, nil
}
