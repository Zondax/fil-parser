package reward

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/filecoin-project/go-state-types/abi"
	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/manifest"
	legacyBuiltin "github.com/filecoin-project/specs-actors/actors/builtin"

	rewardv10 "github.com/filecoin-project/go-state-types/builtin/v10/reward"
	rewardv11 "github.com/filecoin-project/go-state-types/builtin/v11/reward"
	rewardv12 "github.com/filecoin-project/go-state-types/builtin/v12/reward"
	rewardv13 "github.com/filecoin-project/go-state-types/builtin/v13/reward"
	rewardv14 "github.com/filecoin-project/go-state-types/builtin/v14/reward"
	rewardv15 "github.com/filecoin-project/go-state-types/builtin/v15/reward"
	rewardv8 "github.com/filecoin-project/go-state-types/builtin/v8/reward"
	rewardv9 "github.com/filecoin-project/go-state-types/builtin/v9/reward"

	legacyv1 "github.com/filecoin-project/specs-actors/actors/builtin/reward"
	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/reward"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/reward"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/reward"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/reward"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/reward"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/reward"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

type Reward struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) *Reward {
	return &Reward{
		logger: logger,
	}
}

func (r *Reward) Name() string {
	return manifest.RewardKey
}

func (*Reward) Constructor(network string, height int64, raw []byte) (map[string]interface{}, error) {
	return parse(raw, &abi.StoragePower{}, parser.ParamsKey)
}

func (*Reward) StartNetworkHeight() int64 {
	return tools.V1.Height()
}

func (*Reward) Methods(network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error) {
	switch {
	// all legacy version
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V15)...):
		return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
			legacyBuiltin.MethodsReward.Constructor: {
				Name: parser.MethodConstructor,
			},
			legacyBuiltin.MethodsReward.AwardBlockReward: {
				Name: parser.MethodAwardBlockReward,
			},
			legacyBuiltin.MethodsReward.ThisEpochReward: {
				Name: parser.MethodThisEpochReward,
			},
			legacyBuiltin.MethodsReward.UpdateNetworkKPI: {
				Name: parser.MethodUpdateNetworkKPI,
			},
		}, nil
	case tools.V16.IsSupported(network, height):
		return rewardv8.Methods, nil
	case tools.V17.IsSupported(network, height):
		return rewardv9.Methods, nil
	case tools.V18.IsSupported(network, height):
		return rewardv10.Methods, nil
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return rewardv11.Methods, nil
	case tools.V21.IsSupported(network, height):
		return rewardv12.Methods, nil
	case tools.V22.IsSupported(network, height):
		return rewardv13.Methods, nil
	case tools.V23.IsSupported(network, height):
		return rewardv14.Methods, nil
	case tools.V24.IsSupported(network, height):
		return rewardv15.Methods, nil
	default:
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
}

func (*Reward) AwardBlockReward(network string, height int64, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parse(raw, &rewardv15.AwardBlockRewardParams{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parse(raw, &rewardv14.AwardBlockRewardParams{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		return parse(raw, &rewardv13.AwardBlockRewardParams{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		return parse(raw, &rewardv12.AwardBlockRewardParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parse(raw, &rewardv11.AwardBlockRewardParams{}, parser.ParamsKey)
	case tools.V18.IsSupported(network, height):
		return parse(raw, &rewardv10.AwardBlockRewardParams{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		return parse(raw, &rewardv9.AwardBlockRewardParams{}, parser.ParamsKey)
	case tools.V16.IsSupported(network, height):
		return parse(raw, &rewardv8.AwardBlockRewardParams{}, parser.ParamsKey)
	case tools.V15.IsSupported(network, height):
		return parse(raw, &legacyv7.AwardBlockRewardParams{}, parser.ParamsKey)
	case tools.V14.IsSupported(network, height):
		return parse(raw, &legacyv6.AwardBlockRewardParams{}, parser.ParamsKey)
	case tools.V13.IsSupported(network, height):
		return parse(raw, &legacyv5.AwardBlockRewardParams{}, parser.ParamsKey)
	case tools.V12.IsSupported(network, height):
		return parse(raw, &legacyv4.AwardBlockRewardParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parse(raw, &legacyv3.AwardBlockRewardParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V8, tools.V9):
		return parse(raw, &legacyv2.AwardBlockRewardParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return parse(raw, &legacyv1.AwardBlockRewardParams{}, parser.ParamsKey)
	}
	return nil, nil
}

func (*Reward) UpdateNetworkKPI(network string, height int64, raw []byte) (map[string]interface{}, error) {
	return parse(raw, &abi.StoragePower{}, parser.ParamsKey)
}

func (*Reward) ThisEpochReward(network string, height int64, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parse(raw, &rewardv15.ThisEpochRewardReturn{}, parser.ReturnKey)
	case tools.V23.IsSupported(network, height):
		return parse(raw, &rewardv14.ThisEpochRewardReturn{}, parser.ReturnKey)
	case tools.V22.IsSupported(network, height):
		return parse(raw, &rewardv13.ThisEpochRewardReturn{}, parser.ReturnKey)
	case tools.V21.IsSupported(network, height):
		return parse(raw, &rewardv12.ThisEpochRewardReturn{}, parser.ReturnKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parse(raw, &rewardv11.ThisEpochRewardReturn{}, parser.ReturnKey)
	case tools.V18.IsSupported(network, height):
		return parse(raw, &rewardv10.ThisEpochRewardReturn{}, parser.ReturnKey)
	case tools.V17.IsSupported(network, height):
		return parse(raw, &rewardv9.ThisEpochRewardReturn{}, parser.ReturnKey)
	case tools.V16.IsSupported(network, height):
		return parse(raw, &rewardv8.ThisEpochRewardReturn{}, parser.ReturnKey)
	case tools.V15.IsSupported(network, height):
		return parse(raw, &legacyv7.ThisEpochRewardReturn{}, parser.ReturnKey)
	case tools.V14.IsSupported(network, height):
		return parse(raw, &legacyv6.ThisEpochRewardReturn{}, parser.ReturnKey)
	case tools.V13.IsSupported(network, height):
		return parse(raw, &legacyv5.ThisEpochRewardReturn{}, parser.ReturnKey)
	case tools.V12.IsSupported(network, height):
		return parse(raw, &legacyv4.ThisEpochRewardReturn{}, parser.ReturnKey)
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parse(raw, &legacyv3.ThisEpochRewardReturn{}, parser.ReturnKey)
	case tools.AnyIsSupported(network, height, tools.V8, tools.V9):
		return parse(raw, &legacyv2.ThisEpochRewardReturn{}, parser.ReturnKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return parse(raw, &legacyv1.ThisEpochRewardReturn{}, parser.ReturnKey)
	}
	return nil, nil
}
