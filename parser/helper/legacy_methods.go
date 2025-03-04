package helper

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/zondax/fil-parser/parser"

	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"
	legacyBuiltin "github.com/filecoin-project/specs-actors/actors/builtin"
)

func legacyMethods(actorName string) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error) {
	actorMethods := map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
		abi.MethodNum(0): {
			Name: "Constructor",
		},
		abi.MethodNum(1): {
			Name: "Send",
		},
	}
	switch actorName {
	case manifest.AccountKey:
		actorMethods = accountMethods()
	case manifest.InitKey:
		actorMethods = initMethods()
	case manifest.CronKey:
		actorMethods = cronMethods()
	case manifest.RewardKey:
		actorMethods = rewardMethods()
	case manifest.MultisigKey:
		actorMethods = multisigMethods()
	case manifest.PaychKey:
		actorMethods = paychMethods()
	case manifest.MarketKey:
		actorMethods = marketMethods()
	case manifest.PowerKey:
		actorMethods = powerMethods()
	case manifest.MinerKey:
		actorMethods = minerMethods()
	case manifest.VerifregKey:
		actorMethods = verifiedRegistryMethods()
	case manifest.SystemKey:
		actorMethods = systemMethods()
	default:
		return nil, parser.ErrNotKnownActor
	}
	return actorMethods, nil
}

func accountMethods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
		legacyBuiltin.MethodsAccount.Constructor: {
			Name: "Constructor",
		},
		legacyBuiltin.MethodsAccount.PubkeyAddress: {
			Name: "PubkeyAddress",
		},
	}
}

func initMethods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
		legacyBuiltin.MethodsInit.Constructor: {
			Name: "Constructor",
		},
		legacyBuiltin.MethodsInit.Exec: {
			Name: "Exec",
		},
	}
}

func cronMethods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
		legacyBuiltin.MethodsCron.Constructor: {
			Name: "Constructor",
		},
		legacyBuiltin.MethodsCron.EpochTick: {
			Name: "EpochTick",
		},
	}
}

func rewardMethods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
		legacyBuiltin.MethodsReward.Constructor: {
			Name: "Constructor",
		},
		legacyBuiltin.MethodsReward.AwardBlockReward: {
			Name: "AwardBlockReward",
		},
		legacyBuiltin.MethodsReward.ThisEpochReward: {
			Name: "ThisEpochReward",
		},
		legacyBuiltin.MethodsReward.UpdateNetworkKPI: {
			Name: "UpdateNetworkKPI",
		},
	}
}

func multisigMethods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
		legacyBuiltin.MethodsMultisig.Constructor: {
			Name: "Constructor",
		},
		legacyBuiltin.MethodsMultisig.Propose: {
			Name: "Propose",
		},
		legacyBuiltin.MethodsMultisig.Approve: {
			Name: "Approve",
		},
		legacyBuiltin.MethodsMultisig.Cancel: {
			Name: "Cancel",
		},
		legacyBuiltin.MethodsMultisig.AddSigner: {
			Name: "AddSigner",
		},
		legacyBuiltin.MethodsMultisig.RemoveSigner: {
			Name: "RemoveSigner",
		},
		legacyBuiltin.MethodsMultisig.SwapSigner: {
			Name: "SwapSigner",
		},
		legacyBuiltin.MethodsMultisig.ChangeNumApprovalsThreshold: {
			Name: "ChangeNumApprovalsThreshold",
		},
		legacyBuiltin.MethodsMultisig.LockBalance: {
			Name: "LockBalance",
		},
	}
}

func paychMethods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
		legacyBuiltin.MethodsPaych.Constructor: {
			Name: "Constructor",
		},
		legacyBuiltin.MethodsPaych.UpdateChannelState: {
			Name: "UpdateChannelState",
		},
		legacyBuiltin.MethodsPaych.Settle: {
			Name: "Settle",
		},
		legacyBuiltin.MethodsPaych.Collect: {
			Name: "Collect",
		},
	}
}

func marketMethods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
		legacyBuiltin.MethodsMarket.Constructor: {
			Name: "Constructor",
		},
		legacyBuiltin.MethodsMarket.AddBalance: {
			Name: "AddBalance",
		},
		legacyBuiltin.MethodsMarket.WithdrawBalance: {
			Name: "WithdrawBalance",
		},
		legacyBuiltin.MethodsMarket.PublishStorageDeals: {
			Name: "PublishStorageDeals",
		},
		legacyBuiltin.MethodsMarket.VerifyDealsForActivation: {
			Name: "VerifyDealsForActivation",
		},
		legacyBuiltin.MethodsMarket.ActivateDeals: {
			Name: "ActivateDeals",
		},
		legacyBuiltin.MethodsMarket.OnMinerSectorsTerminate: {
			Name: "OnMinerSectorsTerminate",
		},
		legacyBuiltin.MethodsMarket.ComputeDataCommitment: {
			Name: "ComputeDataCommitment",
		},
		legacyBuiltin.MethodsMarket.CronTick: {
			Name: "CronTick",
		},
	}
}

func powerMethods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
		legacyBuiltin.MethodsPower.Constructor: {
			Name: "Constructor",
		},
		legacyBuiltin.MethodsPower.CreateMiner: {
			Name: "CreateMiner",
		},
		legacyBuiltin.MethodsPower.UpdateClaimedPower: {
			Name: "UpdateClaimedPower",
		},
		legacyBuiltin.MethodsPower.EnrollCronEvent: {
			Name: "EnrollCronEvent",
		},
		legacyBuiltin.MethodsPower.OnEpochTickEnd: {
			Name: "OnEpochTickEnd",
		},
		legacyBuiltin.MethodsPower.UpdatePledgeTotal: {
			Name: "UpdatePledgeTotal",
		},
		legacyBuiltin.MethodsPower.OnConsensusFault: {
			Name: "OnConsensusFault",
		},
		legacyBuiltin.MethodsPower.SubmitPoRepForBulkVerify: {
			Name: "SubmitPoRepForBulkVerify",
		},
		legacyBuiltin.MethodsPower.CurrentTotalPower: {
			Name: "CurrentTotalPower",
		},
	}
}

func minerMethods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
		legacyBuiltin.MethodsMiner.Constructor: {
			Name: "Constructor",
		},
		legacyBuiltin.MethodsMiner.ControlAddresses: {
			Name: "ControlAddresses",
		},
		legacyBuiltin.MethodsMiner.ChangeWorkerAddress: {
			Name: "ChangeWorkerAddress",
		},
		legacyBuiltin.MethodsMiner.ChangePeerID: {
			Name: "ChangePeerID",
		},
		legacyBuiltin.MethodsMiner.SubmitWindowedPoSt: {
			Name: "SubmitWindowedPoSt",
		},
		legacyBuiltin.MethodsMiner.PreCommitSector: {
			Name: "PreCommitSector",
		},
		legacyBuiltin.MethodsMiner.ProveCommitSector: {
			Name: "ProveCommitSector",
		},
		nonLegacyBuiltin.MethodsMiner.ExtendSectorExpiration: {
			Name: "ExtendSectorExpiration",
		},
		legacyBuiltin.MethodsMiner.TerminateSectors: {
			Name: "TerminateSectors",
		},
		legacyBuiltin.MethodsMiner.DeclareFaults: {
			Name: "DeclareFaults",
		},
		legacyBuiltin.MethodsMiner.DeclareFaultsRecovered: {
			Name: "DeclareFaultsRecovered",
		},
		legacyBuiltin.MethodsMiner.OnDeferredCronEvent: {
			Name: "OnDeferredCronEvent",
		},
		legacyBuiltin.MethodsMiner.CheckSectorProven: {
			Name: "CheckSectorProven",
		},
		legacyBuiltin.MethodsMiner.AddLockedFund: {
			Name: "AddLockedFund",
		},
		legacyBuiltin.MethodsMiner.ReportConsensusFault: {
			Name: "ReportConsensusFault",
		},
		legacyBuiltin.MethodsMiner.WithdrawBalance: {
			Name: "WithdrawBalance",
		},
	}
}

func verifiedRegistryMethods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
		legacyBuiltin.MethodsVerifiedRegistry.Constructor: {
			Name: "Constructor",
		},
		legacyBuiltin.MethodsVerifiedRegistry.AddVerifier: {
			Name: "AddVerifier",
		},
		legacyBuiltin.MethodsVerifiedRegistry.RemoveVerifier: {
			Name: "RemoveVerifier",
		},
		legacyBuiltin.MethodsVerifiedRegistry.AddVerifiedClient: {
			Name: "AddVerifiedClient",
		},
		legacyBuiltin.MethodsVerifiedRegistry.UseBytes: {
			Name: "UseBytes",
		},
		legacyBuiltin.MethodsVerifiedRegistry.RestoreBytes: {
			Name: "RestoreBytes",
		},
	}
}

func systemMethods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
		abi.MethodNum(0): {
			Name: "Constructor",
		},
	}
}
