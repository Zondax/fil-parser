package reward

import (
	"context"
	"fmt"

	"github.com/zondax/golem/pkg/logger"

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
	rewardv16 "github.com/filecoin-project/go-state-types/builtin/v16/reward"
	rewardv8 "github.com/filecoin-project/go-state-types/builtin/v8/reward"
	rewardv9 "github.com/filecoin-project/go-state-types/builtin/v9/reward"

	actor_tools "github.com/zondax/fil-parser/actors/v2/tools"

	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

type Reward struct {
	logger *logger.Logger
}

func New(logger *logger.Logger) *Reward {
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

func (r *Reward) Methods(_ context.Context, network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error) {
	switch {
	// all legacy version
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V15)...):
		return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
			legacyBuiltin.MethodsReward.Constructor: {
				Name:   parser.MethodConstructor,
				Method: actor_tools.ParseConstructor,
			},
			legacyBuiltin.MethodsReward.AwardBlockReward: {
				Name:   parser.MethodAwardBlockReward,
				Method: r.AwardBlockReward,
			},
			legacyBuiltin.MethodsReward.ThisEpochReward: {
				Name:   parser.MethodThisEpochReward,
				Method: r.ThisEpochReward,
			},
			legacyBuiltin.MethodsReward.UpdateNetworkKPI: {
				Name:   parser.MethodUpdateNetworkKPI,
				Method: r.UpdateNetworkKPI,
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
	case tools.V25.IsSupported(network, height):
		return rewardv16.Methods, nil
	default:
		return nil, fmt.Errorf("%w: %d", actor_tools.ErrUnsupportedHeight, height)
	}
}

func (*Reward) AwardBlockReward(network string, height int64, raw []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := awardBlockRewardParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actor_tools.ErrUnsupportedHeight, height)
	}

	return parse(raw, params(), parser.ParamsKey)
}

func (*Reward) UpdateNetworkKPI(network string, height int64, raw []byte) (map[string]interface{}, error) {
	return parse(raw, &abi.StoragePower{}, parser.ParamsKey)
}

func (*Reward) ThisEpochReward(network string, height int64, raw []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	returns, ok := thisEpochRewardReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actor_tools.ErrUnsupportedHeight, height)
	}

	return parse(raw, returns(), parser.ReturnKey)
}
