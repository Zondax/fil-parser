package reward

import (
	"reflect"

	"github.com/filecoin-project/go-state-types/abi"
	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"
	rewardv10 "github.com/filecoin-project/go-state-types/builtin/v10/reward"
	rewardv11 "github.com/filecoin-project/go-state-types/builtin/v11/reward"
	rewardv12 "github.com/filecoin-project/go-state-types/builtin/v12/reward"
	rewardv13 "github.com/filecoin-project/go-state-types/builtin/v13/reward"
	rewardv14 "github.com/filecoin-project/go-state-types/builtin/v14/reward"
	rewardv15 "github.com/filecoin-project/go-state-types/builtin/v15/reward"
	rewardv16 "github.com/filecoin-project/go-state-types/builtin/v16/reward"
	rewardv8 "github.com/filecoin-project/go-state-types/builtin/v8/reward"
	rewardv9 "github.com/filecoin-project/go-state-types/builtin/v9/reward"
	legacyv1 "github.com/filecoin-project/specs-actors/actors/builtin/reward"
	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/reward"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/reward"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/reward"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/reward"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/reward"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/reward"
	cbg "github.com/whyrusleeping/cbor-gen"
	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

func v1Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	r := &Reward{}
	return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
		1: {
			Name:   parser.MethodConstructor,
			Method: actors.ParseConstructor,
		},
		2: {
			Name:   parser.MethodAwardBlockReward,
			Method: r.AwardBlockReward,
		},
		3: {
			Name:   parser.MethodThisEpochReward,
			Method: r.ThisEpochReward,
		},
		4: {
			Name:   parser.MethodUpdateNetworkKPI,
			Method: r.UpdateNetworkKPI,
		},
	}
}

func v2Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return v1Methods()
}

func v3Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return v1Methods()
}

func v4Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return v1Methods()
}

func v5Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return v1Methods()
}

func v6Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return v1Methods()
}
func v7Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return v1Methods()
}

var awardBlockRewardParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.AwardBlockRewardParams) },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.AwardBlockRewardParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.AwardBlockRewardParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.AwardBlockRewardParams) },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.AwardBlockRewardParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.AwardBlockRewardParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.AwardBlockRewardParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.AwardBlockRewardParams) },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.AwardBlockRewardParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.AwardBlockRewardParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.AwardBlockRewardParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.AwardBlockRewardParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.AwardBlockRewardParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.AwardBlockRewardParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.AwardBlockRewardParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.AwardBlockRewardParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(rewardv8.AwardBlockRewardParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(rewardv9.AwardBlockRewardParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(rewardv10.AwardBlockRewardParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(rewardv11.AwardBlockRewardParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(rewardv11.AwardBlockRewardParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(rewardv12.AwardBlockRewardParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(rewardv13.AwardBlockRewardParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(rewardv14.AwardBlockRewardParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(rewardv15.AwardBlockRewardParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(rewardv16.AwardBlockRewardParams) },
}

var thisEpochRewardReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ThisEpochRewardReturn) },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ThisEpochRewardReturn) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ThisEpochRewardReturn) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ThisEpochRewardReturn) },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ThisEpochRewardReturn) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ThisEpochRewardReturn) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ThisEpochRewardReturn) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ThisEpochRewardReturn) },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ThisEpochRewardReturn) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ThisEpochRewardReturn) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.ThisEpochRewardReturn) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.ThisEpochRewardReturn) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.ThisEpochRewardReturn) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.ThisEpochRewardReturn) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.ThisEpochRewardReturn) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.ThisEpochRewardReturn) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(rewardv8.ThisEpochRewardReturn) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(rewardv9.ThisEpochRewardReturn) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(rewardv10.ThisEpochRewardReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(rewardv11.ThisEpochRewardReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(rewardv11.ThisEpochRewardReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(rewardv12.ThisEpochRewardReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(rewardv13.ThisEpochRewardReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(rewardv14.ThisEpochRewardReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(rewardv15.ThisEpochRewardReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(rewardv16.ThisEpochRewardReturn) },
}

func GetMinerFromAwardBlockRewardParams(params any) string {
	if params == nil {
		return ""
	}

	var reward any

	t := reflect.TypeOf(params)
	switch t.Kind() {
	case reflect.Ptr:
		val := reflect.ValueOf(params).Elem()
		if val.IsValid() && val.CanInterface() {
			reward = val.Interface()
		}
	case reflect.Struct:
		reward = params
	}

	switch p := reward.(type) {
	case legacyv1.AwardBlockRewardParams:
		return p.Miner.String()
	// Duplicate cases , they will cause a compile time error.
	// case legacyv2.AwardBlockRewardParams:
	// 	return p.Miner.String()
	// case legacyv3.AwardBlockRewardParams:
	// 	return p.Miner.String()
	// case legacyv4.AwardBlockRewardParams:
	// 	return p.Miner.String()
	// case legacyv5.AwardBlockRewardParams:
	// 	return p.Miner.String()
	// case legacyv6.AwardBlockRewardParams:
	// 	return p.Miner.String()
	// case legacyv7.AwardBlockRewardParams:
	// 	return p.Miner.String()
	// case rewardv8.AwardBlockRewardParams:
	// 	return p.Miner.String()
	case rewardv8.AwardBlockRewardParams:
		return p.Miner.String()
	case rewardv9.AwardBlockRewardParams:
		return p.Miner.String()
	case rewardv10.AwardBlockRewardParams:
		return p.Miner.String()
	case rewardv11.AwardBlockRewardParams:
		return p.Miner.String()
	case rewardv12.AwardBlockRewardParams:
		return p.Miner.String()
	case rewardv13.AwardBlockRewardParams:
		return p.Miner.String()
	case rewardv14.AwardBlockRewardParams:
		return p.Miner.String()
	case rewardv15.AwardBlockRewardParams:
		return p.Miner.String()
	case rewardv16.AwardBlockRewardParams:
		return p.Miner.String()
	}
	return ""
}
