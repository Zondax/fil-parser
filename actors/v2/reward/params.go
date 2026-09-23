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
	rewardv17 "github.com/filecoin-project/go-state-types/builtin/v17/reward"
	rewardv18 "github.com/filecoin-project/go-state-types/builtin/v18/reward"
	rewardv19 "github.com/filecoin-project/go-state-types/builtin/v19/reward"
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
	"github.com/zondax/fil-parser/actors/v2/reward/types"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

// All methods can be found in the Actor.Exports method in
// the correct version package for "github.com/filecoin-project/specs-actors/actors/builtin/reward"

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
	tools.V26.String(): func() cbg.CBORUnmarshaler { return new(rewardv16.AwardBlockRewardParams) },
	tools.V27.String(): func() cbg.CBORUnmarshaler { return new(rewardv17.AwardBlockRewardParams) },
	tools.V28.String(): func() cbg.CBORUnmarshaler { return new(rewardv18.AwardBlockRewardParams) },
	tools.V29.String(): func() cbg.CBORUnmarshaler { return new(rewardv19.AwardBlockRewardParams) },
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
	tools.V26.String(): func() cbg.CBORUnmarshaler { return new(rewardv16.ThisEpochRewardReturn) },
	tools.V27.String(): func() cbg.CBORUnmarshaler { return new(rewardv17.ThisEpochRewardReturn) },
	tools.V28.String(): func() cbg.CBORUnmarshaler { return new(rewardv18.ThisEpochRewardReturn) },
	tools.V29.String(): func() cbg.CBORUnmarshaler { return new(rewardv19.ThisEpochRewardReturn) },
}

var constructorParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return new(abi.StoragePower) },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(abi.StoragePower) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(abi.StoragePower) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(abi.StoragePower) },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(abi.StoragePower) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(abi.StoragePower) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(abi.StoragePower) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(abi.StoragePower) },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(abi.StoragePower) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(abi.StoragePower) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(abi.StoragePower) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(abi.StoragePower) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(abi.StoragePower) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(abi.StoragePower) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(abi.StoragePower) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(abi.StoragePower) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(abi.StoragePower) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(abi.StoragePower) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(abi.StoragePower) },

	// From V19 the storagePower is in a struct that is not present in the go-state-types package but present
	// in the rust builtin-actors package.
	// https://github.com/filecoin-project/builtin-actors/blob/cd9ac2bb0afcca7a59465e57cee6569e69070d7a/actors/reward/src/lib.rs#L54
	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(types.ConstructorParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(types.ConstructorParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(types.ConstructorParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(types.ConstructorParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(types.ConstructorParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(types.ConstructorParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(types.ConstructorParams) },
	tools.V26.String(): func() cbg.CBORUnmarshaler { return new(types.ConstructorParams) },
	tools.V27.String(): func() cbg.CBORUnmarshaler { return new(types.ConstructorParams) },
	tools.V28.String(): func() cbg.CBORUnmarshaler { return new(types.ConstructorParams) },
	tools.V29.String(): func() cbg.CBORUnmarshaler { return new(types.ConstructorParams) },
}

// SetWeightRecordsExported was added in builtin-actors v19 (NV29 Solstice), so V29 is the
// first and only supported version.
var setWeightRecordsExportedParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V29.String(): func() cbg.CBORUnmarshaler { return new(rewardv19.SetWeightRecordsParams) },
}

// StepWeightRecordsExported was added in builtin-actors v19 (NV29 Solstice), so V29 is the
// first and only supported version.
var stepWeightRecordsExportedParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V29.String(): func() cbg.CBORUnmarshaler { return new(rewardv19.StepWeightRecordsParams) },
}

// RegisterStreamExported was added in builtin-actors v19 (NV29 Solstice), so V29 is the
// first and only supported version.
var registerStreamExportedParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V29.String(): func() cbg.CBORUnmarshaler { return new(rewardv19.RegisterStreamParams) },
}

// RemoveStreamExported was added in builtin-actors v19 (NV29 Solstice), so V29 is the
// first and only supported version.
var removeStreamExportedParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V29.String(): func() cbg.CBORUnmarshaler { return new(rewardv19.RemoveStreamParams) },
}

// SetDistributionExported was added in builtin-actors v19 (NV29 Solstice), so V29 is the
// first and only supported version.
var setDistributionExportedParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V29.String(): func() cbg.CBORUnmarshaler { return new(rewardv19.SetDistributionParams) },
}

// SetSharesExported was added in builtin-actors v19 (NV29 Solstice), so V29 is the
// first and only supported version.
var setSharesExportedParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V29.String(): func() cbg.CBORUnmarshaler { return new(rewardv19.SetSharesParams) },
}

// ReplaceAddressExported was added in builtin-actors v19 (NV29 Solstice), so V29 is the
// first and only supported version.
var replaceAddressExportedParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V29.String(): func() cbg.CBORUnmarshaler { return new(rewardv19.ReplaceAddressParams) },
}

// ReplaceAddressReturn is a repr(u8) enum encoded as a CBOR unsigned integer
// (AddressReplaced=0, OldAddressNotInLedger=1), added in builtin-actors v19.0.1 and
// go-state-types v0.19.1, which gives it its own UnmarshalCBOR.
var replaceAddressExportedReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V29.String(): func() cbg.CBORUnmarshaler { return new(rewardv19.ReplaceAddressReturn) },
}

// CancelPendingExported was added in builtin-actors v19 (NV29 Solstice), so V29 is the
// first and only supported version.
var cancelPendingExportedParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V29.String(): func() cbg.CBORUnmarshaler { return new(rewardv19.CancelPendingParams) },
}

// ClaimExported was added in builtin-actors v19 (NV29 Solstice), so V29 is the
// first and only supported version.
var claimExportedParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V29.String(): func() cbg.CBORUnmarshaler { return new(rewardv19.ClaimParams) },
}

var claimExportedReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V29.String(): func() cbg.CBORUnmarshaler { return new(rewardv19.ClaimReturn) },
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
	case rewardv17.AwardBlockRewardParams:
		return p.Miner.String()
	case rewardv18.AwardBlockRewardParams:
		return p.Miner.String()
	case rewardv19.AwardBlockRewardParams:
		return p.Miner.String()
	}
	return ""
}
