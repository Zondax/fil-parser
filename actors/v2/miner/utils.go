package miner

import (
	"math/big"

	"github.com/filecoin-project/go-state-types/builtin"
	"github.com/zondax/fil-parser/tools"

	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/miner"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/miner"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/miner"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/miner"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/miner"

	builtinv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	builtinv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin"
	builtinv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
	builtinv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin"
	builtinv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin"
	builtinv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin"
)

func LockedRewardFactorNum(network string, height int64) *big.Int {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsAfter(tools.V15)...):
		return legacyv7.LockedRewardFactorNum.Int
	case tools.V14.IsSupported(network, height):
		return legacyv6.LockedRewardFactorNum.Int
	case tools.V13.IsSupported(network, height):
		return legacyv5.LockedRewardFactorNum.Int
	case tools.V12.IsSupported(network, height):
		return legacyv4.LockedRewardFactorNum.Int
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V11)...):
		return legacyv3.LockedRewardFactorNum.Int
	}
	return nil
}

func LockedRewardFactorDenom(network string, height int64) *big.Int {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsAfter(tools.V15)...):
		return legacyv7.LockedRewardFactorDenom.Int
	case tools.V14.IsSupported(network, height):
		return legacyv6.LockedRewardFactorDenom.Int
	case tools.V13.IsSupported(network, height):
		return legacyv5.LockedRewardFactorDenom.Int
	case tools.V12.IsSupported(network, height):
		return legacyv4.LockedRewardFactorDenom.Int
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V11)...):
		return legacyv3.LockedRewardFactorDenom.Int
	}
	return nil
}

func VerifiedDealWeightMultiplier(network string, height int64) *big.Int {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsAfter(tools.V16)...):
		return builtin.VerifiedDealWeightMultiplier.Int
	case tools.V15.IsSupported(network, height):
		return builtinv7.VerifiedDealWeightMultiplier.Int
	case tools.V14.IsSupported(network, height):
		return builtinv6.VerifiedDealWeightMultiplier.Int
	case tools.V13.IsSupported(network, height):
		return builtinv5.VerifiedDealWeightMultiplier.Int
	case tools.V12.IsSupported(network, height):
		return builtinv4.VerifiedDealWeightMultiplier.Int
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return builtinv3.VerifiedDealWeightMultiplier.Int
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V9)...):
		return builtinv2.VerifiedDealWeightMultiplier.Int
	}
	return nil
}

func QualityBaseMultiplier(network string, height int64) *big.Int {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsAfter(tools.V16)...):
		return builtin.QualityBaseMultiplier.Int
	case tools.V15.IsSupported(network, height):
		return builtinv7.QualityBaseMultiplier.Int
	case tools.V14.IsSupported(network, height):
		return builtinv6.QualityBaseMultiplier.Int
	case tools.V13.IsSupported(network, height):
		return builtinv5.QualityBaseMultiplier.Int
	case tools.V12.IsSupported(network, height):
		return builtinv4.QualityBaseMultiplier.Int
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return builtinv3.QualityBaseMultiplier.Int
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V9)...):
		return builtinv2.QualityBaseMultiplier.Int
	}
	return nil
}
