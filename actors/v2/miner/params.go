package miner

import (
	miner10 "github.com/filecoin-project/go-state-types/builtin/v10/miner"
	miner11 "github.com/filecoin-project/go-state-types/builtin/v11/miner"
	miner12 "github.com/filecoin-project/go-state-types/builtin/v12/miner"
	miner13 "github.com/filecoin-project/go-state-types/builtin/v13/miner"
	miner14 "github.com/filecoin-project/go-state-types/builtin/v14/miner"
	miner15 "github.com/filecoin-project/go-state-types/builtin/v15/miner"
	miner16 "github.com/filecoin-project/go-state-types/builtin/v16/miner"
	miner8 "github.com/filecoin-project/go-state-types/builtin/v8/miner"
	miner9 "github.com/filecoin-project/go-state-types/builtin/v9/miner"
	legacyv1 "github.com/filecoin-project/specs-actors/actors/builtin/miner"
	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/miner"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/miner"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/miner"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/miner"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/miner"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/miner"
	cbg "github.com/whyrusleeping/cbor-gen"
	"github.com/zondax/fil-parser/tools"
)

func changeMultiaddrsParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.ChangeMultiaddrsParams{},

		tools.V8.String(): &legacyv2.ChangeMultiaddrsParams{},
		tools.V9.String(): &legacyv2.ChangeMultiaddrsParams{},

		tools.V10.String(): &legacyv3.ChangeMultiaddrsParams{},
		tools.V11.String(): &legacyv3.ChangeMultiaddrsParams{},

		tools.V12.String(): &legacyv4.ChangeMultiaddrsParams{},
		tools.V13.String(): &legacyv5.ChangeMultiaddrsParams{},
		tools.V14.String(): &legacyv6.ChangeMultiaddrsParams{},
		tools.V15.String(): &legacyv7.ChangeMultiaddrsParams{},
		tools.V16.String(): &miner8.ChangeMultiaddrsParams{},
		tools.V17.String(): &miner9.ChangeMultiaddrsParams{},
		tools.V18.String(): &miner10.ChangeMultiaddrsParams{},

		tools.V19.String(): &miner11.ChangeMultiaddrsParams{},
		tools.V20.String(): &miner11.ChangeMultiaddrsParams{},

		tools.V21.String(): &miner12.ChangeMultiaddrsParams{},
		tools.V22.String(): &miner13.ChangeMultiaddrsParams{},
		tools.V23.String(): &miner14.ChangeMultiaddrsParams{},
		tools.V24.String(): &miner15.ChangeMultiaddrsParams{},
		tools.V25.String(): &miner16.ChangeMultiaddrsParams{},
	}
}

func changePeerIDParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.ChangePeerIDParams{},

		tools.V8.String(): &legacyv2.ChangePeerIDParams{},
		tools.V9.String(): &legacyv2.ChangePeerIDParams{},

		tools.V10.String(): &legacyv3.ChangePeerIDParams{},
		tools.V11.String(): &legacyv3.ChangePeerIDParams{},

		tools.V12.String(): &legacyv4.ChangePeerIDParams{},
		tools.V13.String(): &legacyv5.ChangePeerIDParams{},
		tools.V14.String(): &legacyv6.ChangePeerIDParams{},
		tools.V15.String(): &legacyv7.ChangePeerIDParams{},
		tools.V16.String(): &miner8.ChangePeerIDParams{},
		tools.V17.String(): &miner9.ChangePeerIDParams{},
		tools.V18.String(): &miner10.ChangePeerIDParams{},

		tools.V19.String(): &miner11.ChangePeerIDParams{},
		tools.V20.String(): &miner11.ChangePeerIDParams{},

		tools.V21.String(): &miner12.ChangePeerIDParams{},
		tools.V22.String(): &miner13.ChangePeerIDParams{},
		tools.V23.String(): &miner14.ChangePeerIDParams{},
		tools.V24.String(): &miner15.ChangePeerIDParams{},
		tools.V25.String(): &miner16.ChangePeerIDParams{},
	}
}

func changeWorkerAddressParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.ChangeWorkerAddressParams{},

		tools.V8.String(): &legacyv2.ChangeWorkerAddressParams{},
		tools.V9.String(): &legacyv2.ChangeWorkerAddressParams{},

		tools.V10.String(): &legacyv3.ChangeWorkerAddressParams{},
		tools.V11.String(): &legacyv3.ChangeWorkerAddressParams{},

		tools.V12.String(): &legacyv4.ChangeWorkerAddressParams{},
		tools.V13.String(): &legacyv5.ChangeWorkerAddressParams{},
		tools.V14.String(): &legacyv6.ChangeWorkerAddressParams{},
		tools.V15.String(): &legacyv7.ChangeWorkerAddressParams{},
		tools.V16.String(): &miner8.ChangeWorkerAddressParams{},
		tools.V17.String(): &miner9.ChangeWorkerAddressParams{},
		tools.V18.String(): &miner10.ChangeWorkerAddressParams{},

		tools.V19.String(): &miner11.ChangeWorkerAddressParams{},
		tools.V20.String(): &miner11.ChangeWorkerAddressParams{},

		tools.V21.String(): &miner12.ChangeWorkerAddressParams{},
		tools.V22.String(): &miner13.ChangeWorkerAddressParams{},
		tools.V23.String(): &miner14.ChangeWorkerAddressParams{},
		tools.V24.String(): &miner15.ChangeWorkerAddressParams{},
		tools.V25.String(): &miner16.ChangeWorkerAddressParams{},
	}
}

func isControllingAddressParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V18.String(): &miner10.IsControllingAddressParams{},

		tools.V19.String(): &miner11.IsControllingAddressParams{},
		tools.V20.String(): &miner11.IsControllingAddressParams{},

		tools.V21.String(): &miner12.IsControllingAddressParams{},
		tools.V22.String(): &miner13.IsControllingAddressParams{},
		tools.V23.String(): &miner14.IsControllingAddressParams{},
		tools.V24.String(): &miner15.IsControllingAddressParams{},
		tools.V25.String(): &miner16.IsControllingAddressParams{},
	}
}

func isControllingAddressReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V18.String(): new(miner10.IsControllingAddressReturn),

		tools.V19.String(): new(miner11.IsControllingAddressReturn),
		tools.V20.String(): new(miner11.IsControllingAddressReturn),

		tools.V21.String(): new(miner12.IsControllingAddressReturn),
		tools.V22.String(): new(miner13.IsControllingAddressReturn),
		tools.V23.String(): new(miner14.IsControllingAddressReturn),
		tools.V24.String(): new(miner15.IsControllingAddressReturn),
		tools.V25.String(): new(miner16.IsControllingAddressReturn),
	}
}

func getOwnerReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V18.String(): &miner10.GetOwnerReturn{},

		tools.V19.String(): &miner11.GetOwnerReturn{},
		tools.V20.String(): &miner11.GetOwnerReturn{},

		tools.V21.String(): &miner12.GetOwnerReturn{},
		tools.V22.String(): &miner13.GetOwnerReturn{},
		tools.V23.String(): &miner14.GetOwnerReturn{},
		tools.V24.String(): &miner15.GetOwnerReturn{},
		tools.V25.String(): &miner16.GetOwnerReturn{},
	}
}

func getPeerIDReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V18.String(): &miner10.GetPeerIDReturn{},

		tools.V19.String(): &miner11.GetPeerIDReturn{},
		tools.V20.String(): &miner11.GetPeerIDReturn{},

		tools.V21.String(): &miner12.GetPeerIDReturn{},
		tools.V22.String(): &miner13.GetPeerIDReturn{},
		tools.V23.String(): &miner14.GetPeerIDReturn{},
		tools.V24.String(): &miner15.GetPeerIDReturn{},
		tools.V25.String(): &miner16.GetPeerIDReturn{},
	}
}

func getMultiAddrsReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V18.String(): &miner10.GetMultiAddrsReturn{},

		tools.V19.String(): &miner11.GetMultiAddrsReturn{},
		tools.V20.String(): &miner11.GetMultiAddrsReturn{},

		tools.V21.String(): &miner12.GetMultiAddrsReturn{},
		tools.V22.String(): &miner13.GetMultiAddrsReturn{},
		tools.V23.String(): &miner14.GetMultiAddrsReturn{},
		tools.V24.String(): &miner15.GetMultiAddrsReturn{},
		tools.V25.String(): &miner16.GetMultiAddrsReturn{},
	}
}

func getControlAddressesReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.GetControlAddressesReturn{},

		tools.V8.String(): &legacyv2.GetControlAddressesReturn{},
		tools.V9.String(): &legacyv2.GetControlAddressesReturn{},

		tools.V10.String(): &legacyv3.GetControlAddressesReturn{},
		tools.V11.String(): &legacyv3.GetControlAddressesReturn{},

		tools.V12.String(): &legacyv4.GetControlAddressesReturn{},
		tools.V13.String(): &legacyv5.GetControlAddressesReturn{},
		tools.V14.String(): &legacyv6.GetControlAddressesReturn{},
		tools.V15.String(): &legacyv7.GetControlAddressesReturn{},
		tools.V16.String(): &miner8.GetControlAddressesReturn{},
		tools.V17.String(): &miner9.GetControlAddressesReturn{},
		tools.V18.String(): &miner10.GetControlAddressesReturn{},

		tools.V19.String(): &miner11.GetControlAddressesReturn{},
		tools.V20.String(): &miner11.GetControlAddressesReturn{},

		tools.V21.String(): &miner12.GetControlAddressesReturn{},
		tools.V22.String(): &miner13.GetControlAddressesReturn{},
		tools.V23.String(): &miner14.GetControlAddressesReturn{},
		tools.V24.String(): &miner15.GetControlAddressesReturn{},
		tools.V25.String(): &miner16.GetControlAddressesReturn{},
	}
}

func getAvailableBalanceReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V18.String(): &miner10.GetAvailableBalanceReturn{},

		tools.V19.String(): &miner11.GetAvailableBalanceReturn{},
		tools.V20.String(): &miner11.GetAvailableBalanceReturn{},

		tools.V21.String(): &miner12.GetAvailableBalanceReturn{},
		tools.V22.String(): &miner13.GetAvailableBalanceReturn{},
		tools.V23.String(): &miner14.GetAvailableBalanceReturn{},
		tools.V24.String(): &miner15.GetAvailableBalanceReturn{},
		tools.V25.String(): &miner16.GetAvailableBalanceReturn{},
	}
}

func getVestingFundsReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V18.String(): &miner10.GetVestingFundsReturn{},

		tools.V19.String(): &miner11.GetVestingFundsReturn{},
		tools.V20.String(): &miner11.GetVestingFundsReturn{},

		tools.V21.String(): &miner12.GetVestingFundsReturn{},
		tools.V22.String(): &miner13.GetVestingFundsReturn{},
		tools.V23.String(): &miner14.GetVestingFundsReturn{},
		tools.V24.String(): &miner15.GetVestingFundsReturn{},
		tools.V25.String(): &miner16.GetVestingFundsReturn{},
	}
}

func getWithdrawBalanceParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.WithdrawBalanceParams{},

		tools.V8.String(): &legacyv2.WithdrawBalanceParams{},
		tools.V9.String(): &legacyv2.WithdrawBalanceParams{},

		tools.V10.String(): &legacyv3.WithdrawBalanceParams{},
		tools.V11.String(): &legacyv3.WithdrawBalanceParams{},

		tools.V12.String(): &legacyv4.WithdrawBalanceParams{},
		tools.V13.String(): &legacyv5.WithdrawBalanceParams{},
		tools.V14.String(): &legacyv6.WithdrawBalanceParams{},
		tools.V15.String(): &legacyv7.WithdrawBalanceParams{},
		tools.V16.String(): &miner8.WithdrawBalanceParams{},
		tools.V17.String(): &miner9.WithdrawBalanceParams{},
		tools.V18.String(): &miner10.WithdrawBalanceParams{},

		tools.V19.String(): &miner11.WithdrawBalanceParams{},
		tools.V20.String(): &miner11.WithdrawBalanceParams{},

		tools.V21.String(): &miner12.WithdrawBalanceParams{},
		tools.V22.String(): &miner13.WithdrawBalanceParams{},
		tools.V23.String(): &miner14.WithdrawBalanceParams{},
		tools.V24.String(): &miner15.WithdrawBalanceParams{},
		tools.V25.String(): &miner16.WithdrawBalanceParams{},
	}
}

func extendSectorExpiration2Params() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V17.String(): &miner9.ExtendSectorExpiration2Params{},

		tools.V18.String(): &miner10.ExtendSectorExpiration2Params{},
		tools.V19.String(): &miner11.ExtendSectorExpiration2Params{},

		tools.V20.String(): &miner11.ExtendSectorExpiration2Params{},
		tools.V21.String(): &miner12.ExtendSectorExpiration2Params{},
		tools.V22.String(): &miner13.ExtendSectorExpiration2Params{},
		tools.V23.String(): &miner14.ExtendSectorExpiration2Params{},
		tools.V24.String(): &miner15.ExtendSectorExpiration2Params{},
		tools.V25.String(): &miner16.ExtendSectorExpiration2Params{},
	}
}

func preCommitSectorParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.SectorPreCommitInfo{},

		tools.V8.String(): &legacyv2.SectorPreCommitInfo{},
		tools.V9.String(): &legacyv2.SectorPreCommitInfo{},

		tools.V10.String(): &legacyv3.SectorPreCommitInfo{},
		tools.V11.String(): &legacyv3.SectorPreCommitInfo{},

		tools.V12.String(): &legacyv4.SectorPreCommitInfo{},
		tools.V13.String(): &legacyv5.SectorPreCommitInfo{},
		tools.V14.String(): &legacyv6.SectorPreCommitInfo{},
		tools.V15.String(): &legacyv7.SectorPreCommitInfo{},
		tools.V16.String(): &miner8.SectorPreCommitInfo{},
		tools.V17.String(): &miner9.SectorPreCommitInfo{},
		tools.V18.String(): &miner10.SectorPreCommitInfo{},
		tools.V19.String(): &miner11.SectorPreCommitInfo{},
		tools.V20.String(): &miner11.SectorPreCommitInfo{},
		tools.V21.String(): &miner12.SectorPreCommitInfo{},
		tools.V22.String(): &miner13.SectorPreCommitInfo{},
		tools.V23.String(): &miner14.SectorPreCommitInfo{},
		tools.V24.String(): &miner15.SectorPreCommitInfo{},
		tools.V25.String(): &miner16.SectorPreCommitInfo{},
	}
}

func proveCommitSectorParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.ProveCommitSectorParams{},

		tools.V8.String(): &legacyv2.ProveCommitSectorParams{},
		tools.V9.String(): &legacyv2.ProveCommitSectorParams{},

		tools.V10.String(): &legacyv3.ProveCommitSectorParams{},
		tools.V11.String(): &legacyv3.ProveCommitSectorParams{},

		tools.V12.String(): &legacyv4.ProveCommitSectorParams{},
		tools.V13.String(): &legacyv5.ProveCommitSectorParams{},
		tools.V14.String(): &legacyv6.ProveCommitSectorParams{},
		tools.V15.String(): &legacyv7.ProveCommitSectorParams{},
		tools.V16.String(): &miner8.ProveCommitSectorParams{},
		tools.V17.String(): &miner9.ProveCommitSectorParams{},
		tools.V18.String(): &miner10.ProveCommitSectorParams{},

		tools.V19.String(): &miner11.ProveCommitSectorParams{},
		tools.V20.String(): &miner11.ProveCommitSectorParams{},

		tools.V21.String(): &miner12.ProveCommitSectorParams{},
		tools.V22.String(): &miner13.ProveCommitSectorParams{},
		tools.V23.String(): &miner14.ProveCommitSectorParams{},
		tools.V24.String(): &miner15.ProveCommitSectorParams{},
		tools.V25.String(): &miner16.ProveCommitSectorParams{},
	}
}

func proveCommitSectors3Params() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V22.String(): &miner13.ProveCommitSectors3Params{},
		tools.V23.String(): &miner14.ProveCommitSectors3Params{},
		tools.V24.String(): &miner15.ProveCommitSectors3Params{},
		tools.V25.String(): &miner16.ProveCommitSectors3Params{},
	}
}

func proveCommitSectors3Return() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V22.String(): &miner13.ProveCommitSectors3Return{},
		tools.V23.String(): &miner14.ProveCommitSectors3Return{},
		tools.V24.String(): &miner15.ProveCommitSectors3Return{},
		tools.V25.String(): &miner16.ProveCommitSectors3Return{},
	}
}

func internalSectorSetupForPresealParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V23.String(): &miner14.InternalSectorSetupForPresealParams{},
		tools.V24.String(): &miner15.InternalSectorSetupForPresealParams{},
		tools.V25.String(): &miner16.InternalSectorSetupForPresealParams{},
	}
}

func submitWindowedPoStParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.SubmitWindowedPoStParams{},

		tools.V8.String(): &legacyv2.SubmitWindowedPoStParams{},
		tools.V9.String(): &legacyv2.SubmitWindowedPoStParams{},

		tools.V10.String(): &legacyv3.SubmitWindowedPoStParams{},
		tools.V11.String(): &legacyv3.SubmitWindowedPoStParams{},

		tools.V12.String(): &legacyv4.SubmitWindowedPoStParams{},
		tools.V13.String(): &legacyv5.SubmitWindowedPoStParams{},
		tools.V14.String(): &legacyv6.SubmitWindowedPoStParams{},
		tools.V15.String(): &legacyv7.SubmitWindowedPoStParams{},
		tools.V16.String(): &miner8.SubmitWindowedPoStParams{},
		tools.V17.String(): &miner9.SubmitWindowedPoStParams{},
		tools.V18.String(): &miner10.SubmitWindowedPoStParams{},

		tools.V19.String(): &miner11.SubmitWindowedPoStParams{},
		tools.V20.String(): &miner11.SubmitWindowedPoStParams{},

		tools.V21.String(): &miner12.SubmitWindowedPoStParams{},
		tools.V22.String(): &miner13.SubmitWindowedPoStParams{},
		tools.V23.String(): &miner14.SubmitWindowedPoStParams{},
		tools.V24.String(): &miner15.SubmitWindowedPoStParams{},
		tools.V25.String(): &miner16.SubmitWindowedPoStParams{},
	}
}

func confirmSectorProofsParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V16.String(): &miner8.ConfirmSectorProofsParams{},
		tools.V17.String(): &miner9.ConfirmSectorProofsParams{},
		tools.V18.String(): &miner10.ConfirmSectorProofsParams{},
		tools.V19.String(): &miner11.ConfirmSectorProofsParams{},
		tools.V20.String(): &miner11.ConfirmSectorProofsParams{},
		tools.V21.String(): &miner12.ConfirmSectorProofsParams{},
		tools.V22.String(): &miner13.ConfirmSectorProofsParams{},
	}
}

func checkSectorProvenParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.CheckSectorProvenParams{},

		tools.V8.String(): &legacyv2.CheckSectorProvenParams{},
		tools.V9.String(): &legacyv2.CheckSectorProvenParams{},

		tools.V10.String(): &legacyv3.CheckSectorProvenParams{},
		tools.V11.String(): &legacyv3.CheckSectorProvenParams{},

		tools.V12.String(): &legacyv4.CheckSectorProvenParams{},
		tools.V13.String(): &legacyv5.CheckSectorProvenParams{},
		tools.V14.String(): &legacyv6.CheckSectorProvenParams{},
		tools.V15.String(): &legacyv7.CheckSectorProvenParams{},
		tools.V16.String(): &miner8.CheckSectorProvenParams{},
		tools.V17.String(): &miner9.CheckSectorProvenParams{},
		tools.V18.String(): &miner10.CheckSectorProvenParams{},

		tools.V19.String(): &miner11.CheckSectorProvenParams{},
		tools.V20.String(): &miner11.CheckSectorProvenParams{},

		tools.V21.String(): &miner12.CheckSectorProvenParams{},
		tools.V22.String(): &miner13.CheckSectorProvenParams{},
		tools.V23.String(): &miner14.CheckSectorProvenParams{},
		tools.V24.String(): &miner15.CheckSectorProvenParams{},
		tools.V25.String(): &miner16.CheckSectorProvenParams{},
	}
}

func extendSectorExpirationParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.ExtendSectorExpirationParams{},

		tools.V8.String(): &legacyv2.ExtendSectorExpirationParams{},
		tools.V9.String(): &legacyv2.ExtendSectorExpirationParams{},

		tools.V10.String(): &legacyv3.ExtendSectorExpirationParams{},
		tools.V11.String(): &legacyv3.ExtendSectorExpirationParams{},

		tools.V12.String(): &legacyv4.ExtendSectorExpirationParams{},
		tools.V13.String(): &legacyv5.ExtendSectorExpirationParams{},
		tools.V14.String(): &legacyv6.ExtendSectorExpirationParams{},
		tools.V15.String(): &legacyv7.ExtendSectorExpirationParams{},
		tools.V16.String(): &miner8.ExtendSectorExpirationParams{},
		tools.V17.String(): &miner9.ExtendSectorExpirationParams{},
		tools.V18.String(): &miner10.ExtendSectorExpirationParams{},

		tools.V19.String(): &miner11.ExtendSectorExpirationParams{},
		tools.V20.String(): &miner11.ExtendSectorExpirationParams{},

		tools.V21.String(): &miner12.ExtendSectorExpirationParams{},
		tools.V22.String(): &miner13.ExtendSectorExpirationParams{},
		tools.V23.String(): &miner14.ExtendSectorExpirationParams{},
		tools.V24.String(): &miner15.ExtendSectorExpirationParams{},
		tools.V25.String(): &miner16.ExtendSectorExpirationParams{},
	}
}

func compactSectorNumbersParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.CompactSectorNumbersParams{},

		tools.V8.String(): &legacyv2.CompactSectorNumbersParams{},
		tools.V9.String(): &legacyv2.CompactSectorNumbersParams{},

		tools.V10.String(): &legacyv3.CompactSectorNumbersParams{},
		tools.V11.String(): &legacyv3.CompactSectorNumbersParams{},

		tools.V12.String(): &legacyv4.CompactSectorNumbersParams{},
		tools.V13.String(): &legacyv5.CompactSectorNumbersParams{},
		tools.V14.String(): &legacyv6.CompactSectorNumbersParams{},
		tools.V15.String(): &legacyv7.CompactSectorNumbersParams{},
		tools.V16.String(): &miner8.CompactSectorNumbersParams{},
		tools.V17.String(): &miner9.CompactSectorNumbersParams{},
		tools.V18.String(): &miner10.CompactSectorNumbersParams{},

		tools.V19.String(): &miner11.CompactSectorNumbersParams{},
		tools.V20.String(): &miner11.CompactSectorNumbersParams{},

		tools.V21.String(): &miner12.CompactSectorNumbersParams{},
		tools.V22.String(): &miner13.CompactSectorNumbersParams{},
		tools.V23.String(): &miner14.CompactSectorNumbersParams{},
		tools.V24.String(): &miner15.CompactSectorNumbersParams{},
		tools.V25.String(): &miner16.CompactSectorNumbersParams{},
	}
}

func compactPartitionsParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.CompactPartitionsParams{},

		tools.V8.String(): &legacyv2.CompactPartitionsParams{},
		tools.V9.String(): &legacyv2.CompactPartitionsParams{},

		tools.V10.String(): &legacyv3.CompactPartitionsParams{},
		tools.V11.String(): &legacyv3.CompactPartitionsParams{},

		tools.V12.String(): &legacyv4.CompactPartitionsParams{},
		tools.V13.String(): &legacyv5.CompactPartitionsParams{},
		tools.V14.String(): &legacyv6.CompactPartitionsParams{},
		tools.V15.String(): &legacyv7.CompactPartitionsParams{},
		tools.V16.String(): &miner8.CompactPartitionsParams{},
		tools.V17.String(): &miner9.CompactPartitionsParams{},
		tools.V18.String(): &miner10.CompactPartitionsParams{},

		tools.V19.String(): &miner11.CompactPartitionsParams{},
		tools.V20.String(): &miner11.CompactPartitionsParams{},

		tools.V21.String(): &miner12.CompactPartitionsParams{},
		tools.V22.String(): &miner13.CompactPartitionsParams{},
		tools.V23.String(): &miner14.CompactPartitionsParams{},
		tools.V24.String(): &miner15.CompactPartitionsParams{},
		tools.V25.String(): &miner16.CompactPartitionsParams{},
	}
}

func preCommitSectorBatchParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V13.String(): &legacyv5.PreCommitSectorBatchParams{},
		tools.V14.String(): &legacyv6.PreCommitSectorBatchParams{},
		tools.V15.String(): &legacyv7.PreCommitSectorBatchParams{},
		tools.V16.String(): &miner8.PreCommitSectorBatchParams{},
		tools.V17.String(): &miner9.PreCommitSectorBatchParams{},
		tools.V18.String(): &miner10.PreCommitSectorBatchParams{},

		tools.V19.String(): &miner11.PreCommitSectorBatchParams{},
		tools.V20.String(): &miner11.PreCommitSectorBatchParams{},

		tools.V21.String(): &miner12.PreCommitSectorBatchParams{},
		tools.V22.String(): &miner13.PreCommitSectorBatchParams{},
		tools.V23.String(): &miner14.PreCommitSectorBatchParams{},
		tools.V24.String(): &miner15.PreCommitSectorBatchParams{},
		tools.V25.String(): &miner16.PreCommitSectorBatchParams{},
	}
}

func proveCommitSectorsNIParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V23.String(): &miner14.ProveCommitSectorsNIParams{},
		tools.V24.String(): &miner15.ProveCommitSectorsNIParams{},
		tools.V25.String(): &miner16.ProveCommitSectorsNIParams{},
	}
}

func proveCommitSectorsNIReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V23.String(): &miner14.ProveCommitSectorsNIReturn{},
		tools.V24.String(): &miner15.ProveCommitSectorsNIReturn{},
		tools.V25.String(): &miner16.ProveCommitSectorsNIReturn{},
	}
}
