package miner

import (
	"math/big"

	"github.com/zondax/fil-parser/tools"

	legacyv1 "github.com/filecoin-project/specs-actors/actors/builtin/miner"
	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/miner"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/miner"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/miner"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/miner"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/miner"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/miner"

	builtinv1 "github.com/filecoin-project/specs-actors/actors/builtin"
	builtinv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	builtinv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin"
	builtinv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
	builtinv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin"
	builtinv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin"
	builtinv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin"
	builtinv8 "github.com/filecoin-project/specs-actors/v8/actors/builtin"

	builtin "github.com/filecoin-project/go-state-types/builtin"
)

var lockedRewardFactorNum = map[string]*big.Int{
	tools.V0.String(): legacyv1.LockTargetFactorNum.Int,
	tools.V1.String(): legacyv1.LockTargetFactorNum.Int,
	tools.V2.String(): legacyv1.LockTargetFactorNum.Int,
	tools.V3.String(): legacyv1.LockTargetFactorNum.Int,

	tools.V4.String(): legacyv2.LockedRewardFactorNumV6.Int,
	tools.V5.String(): legacyv2.LockedRewardFactorNumV6.Int,
	tools.V6.String(): legacyv2.LockedRewardFactorNumV6.Int,
	tools.V7.String(): legacyv2.LockedRewardFactorNumV6.Int,
	tools.V8.String(): legacyv2.LockedRewardFactorNumV6.Int,
	tools.V9.String(): legacyv2.LockedRewardFactorNumV6.Int,

	tools.V10.String(): legacyv3.LockedRewardFactorNum.Int,
	tools.V11.String(): legacyv3.LockedRewardFactorNum.Int,

	tools.V12.String(): legacyv4.LockedRewardFactorNum.Int,
	tools.V13.String(): legacyv5.LockedRewardFactorNum.Int,
	tools.V14.String(): legacyv6.LockedRewardFactorNum.Int,
	tools.V15.String(): legacyv7.LockedRewardFactorNum.Int,

	// golang-impl.: LockedRewardFactor = (75/100) * FILCirculatingSupply(t)
	// rust-lang: LockedRewardFactor = (3/4) * FILCirculatingSupply(t)
	tools.V16.String(): big.NewInt(75),
	tools.V17.String(): big.NewInt(75),
	tools.V18.String(): big.NewInt(75),
	tools.V19.String(): big.NewInt(75),
	tools.V20.String(): big.NewInt(75),
	tools.V21.String(): big.NewInt(75),
	tools.V22.String(): big.NewInt(75),
	tools.V23.String(): big.NewInt(75),
	tools.V24.String(): big.NewInt(75),
	tools.V25.String(): big.NewInt(75),
	tools.V26.String(): big.NewInt(75),
	tools.V27.String(): big.NewInt(75),
}

var lockedRewardFactorDenom = map[string]*big.Int{
	tools.V0.String(): legacyv1.LockTargetFactorDenom.Int,
	tools.V1.String(): legacyv1.LockTargetFactorDenom.Int,
	tools.V2.String(): legacyv1.LockTargetFactorDenom.Int,
	tools.V3.String(): legacyv1.LockTargetFactorDenom.Int,

	tools.V4.String(): legacyv2.LockedRewardFactorDenomV6.Int,
	tools.V5.String(): legacyv2.LockedRewardFactorDenomV6.Int,
	tools.V6.String(): legacyv2.LockedRewardFactorDenomV6.Int,
	tools.V7.String(): legacyv2.LockedRewardFactorDenomV6.Int,
	tools.V8.String(): legacyv2.LockedRewardFactorDenomV6.Int,
	tools.V9.String(): legacyv2.LockedRewardFactorDenomV6.Int,

	tools.V10.String(): legacyv3.LockedRewardFactorDenom.Int,
	tools.V11.String(): legacyv3.LockedRewardFactorDenom.Int,

	tools.V12.String(): legacyv4.LockedRewardFactorDenom.Int,
	tools.V13.String(): legacyv5.LockedRewardFactorDenom.Int,
	tools.V14.String(): legacyv6.LockedRewardFactorDenom.Int,
	tools.V15.String(): legacyv7.LockedRewardFactorDenom.Int,

	// golang-impl.: LockedRewardFactor = (75/100) * FILCirculatingSupply(t)
	// rust-lang: LockedRewardFactor = (3/4) * FILCirculatingSupply(t)
	tools.V16.String(): big.NewInt(100),
	tools.V17.String(): big.NewInt(100),
	tools.V18.String(): big.NewInt(100),
	tools.V19.String(): big.NewInt(100),
	tools.V20.String(): big.NewInt(100),
	tools.V21.String(): big.NewInt(100),
	tools.V22.String(): big.NewInt(100),
	tools.V23.String(): big.NewInt(100),
	tools.V24.String(): big.NewInt(100),
	tools.V25.String(): big.NewInt(100),
	tools.V26.String(): big.NewInt(100),
	tools.V27.String(): big.NewInt(100),
}

var verifiedDealWeightMultiplier = map[string]*big.Int{
	tools.V0.String(): builtinv1.VerifiedDealWeightMultiplier.Int,
	tools.V1.String(): builtinv1.VerifiedDealWeightMultiplier.Int,
	tools.V2.String(): builtinv1.VerifiedDealWeightMultiplier.Int,
	tools.V3.String(): builtinv1.VerifiedDealWeightMultiplier.Int,

	tools.V4.String(): builtinv2.VerifiedDealWeightMultiplier.Int,
	tools.V5.String(): builtinv2.VerifiedDealWeightMultiplier.Int,
	tools.V6.String(): builtinv2.VerifiedDealWeightMultiplier.Int,
	tools.V7.String(): builtinv2.VerifiedDealWeightMultiplier.Int,
	tools.V8.String(): builtinv2.VerifiedDealWeightMultiplier.Int,
	tools.V9.String(): builtinv2.VerifiedDealWeightMultiplier.Int,

	tools.V10.String(): builtinv3.VerifiedDealWeightMultiplier.Int,
	tools.V11.String(): builtinv3.VerifiedDealWeightMultiplier.Int,

	tools.V12.String(): builtinv4.VerifiedDealWeightMultiplier.Int,
	tools.V13.String(): builtinv5.VerifiedDealWeightMultiplier.Int,
	tools.V14.String(): builtinv6.VerifiedDealWeightMultiplier.Int,
	tools.V15.String(): builtinv7.VerifiedDealWeightMultiplier.Int,
	tools.V16.String(): builtinv8.VerifiedDealWeightMultiplier.Int,

	tools.V17.String(): builtin.VerifiedDealWeightMultiplier.Int,
	tools.V18.String(): builtin.VerifiedDealWeightMultiplier.Int,
	tools.V19.String(): builtin.VerifiedDealWeightMultiplier.Int,
	tools.V20.String(): builtin.VerifiedDealWeightMultiplier.Int,
	tools.V21.String(): builtin.VerifiedDealWeightMultiplier.Int,
	tools.V22.String(): builtin.VerifiedDealWeightMultiplier.Int,
	tools.V23.String(): builtin.VerifiedDealWeightMultiplier.Int,
	tools.V24.String(): builtin.VerifiedDealWeightMultiplier.Int,
	tools.V25.String(): builtin.VerifiedDealWeightMultiplier.Int,
	tools.V26.String(): builtin.VerifiedDealWeightMultiplier.Int,
	tools.V27.String(): builtin.VerifiedDealWeightMultiplier.Int,
}

var qualityBaseMultiplier = map[string]*big.Int{
	tools.V0.String(): builtinv1.QualityBaseMultiplier.Int,
	tools.V1.String(): builtinv1.QualityBaseMultiplier.Int,
	tools.V2.String(): builtinv1.QualityBaseMultiplier.Int,
	tools.V3.String(): builtinv1.QualityBaseMultiplier.Int,

	tools.V4.String(): builtinv2.QualityBaseMultiplier.Int,
	tools.V5.String(): builtinv2.QualityBaseMultiplier.Int,
	tools.V6.String(): builtinv2.QualityBaseMultiplier.Int,
	tools.V7.String(): builtinv2.QualityBaseMultiplier.Int,
	tools.V8.String(): builtinv2.QualityBaseMultiplier.Int,
	tools.V9.String(): builtinv2.QualityBaseMultiplier.Int,

	tools.V10.String(): builtinv3.QualityBaseMultiplier.Int,
	tools.V11.String(): builtinv3.QualityBaseMultiplier.Int,

	tools.V12.String(): builtinv4.QualityBaseMultiplier.Int,
	tools.V13.String(): builtinv5.QualityBaseMultiplier.Int,
	tools.V14.String(): builtinv6.QualityBaseMultiplier.Int,
	tools.V15.String(): builtinv7.QualityBaseMultiplier.Int,
	tools.V16.String(): builtinv8.QualityBaseMultiplier.Int,

	tools.V17.String(): builtin.QualityBaseMultiplier.Int,
	tools.V18.String(): builtin.QualityBaseMultiplier.Int,
	tools.V19.String(): builtin.QualityBaseMultiplier.Int,
	tools.V20.String(): builtin.QualityBaseMultiplier.Int,
	tools.V21.String(): builtin.QualityBaseMultiplier.Int,
	tools.V22.String(): builtin.QualityBaseMultiplier.Int,
	tools.V23.String(): builtin.QualityBaseMultiplier.Int,
	tools.V24.String(): builtin.QualityBaseMultiplier.Int,
	tools.V25.String(): builtin.QualityBaseMultiplier.Int,
	tools.V26.String(): builtin.QualityBaseMultiplier.Int,
	tools.V27.String(): builtin.QualityBaseMultiplier.Int,
}

func LockedRewardFactorNum(network string, height int64) *big.Int {
	version := tools.VersionFromHeight(network, height)
	lockedRewardFactorNum, ok := lockedRewardFactorNum[version.String()]
	if !ok {
		return big.NewInt(0)
	}
	return lockedRewardFactorNum
}

func LockedRewardFactorDenom(network string, height int64) *big.Int {
	version := tools.VersionFromHeight(network, height)
	lockedRewardFactorDenom, ok := lockedRewardFactorDenom[version.String()]
	if !ok {
		return big.NewInt(1) // prevent divide by zero
	}
	return lockedRewardFactorDenom
}

func VerifiedDealWeightMultiplier(network string, height int64) *big.Int {
	version := tools.VersionFromHeight(network, height)
	verifiedDealWeightMultiplier, ok := verifiedDealWeightMultiplier[version.String()]
	if !ok {
		return big.NewInt(0)
	}
	return verifiedDealWeightMultiplier
}

func QualityBaseMultiplier(network string, height int64) *big.Int {
	version := tools.VersionFromHeight(network, height)
	qualityBaseMultiplier, ok := qualityBaseMultiplier[version.String()]
	if !ok {
		return big.NewInt(0)
	}
	return qualityBaseMultiplier
}
