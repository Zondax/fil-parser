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

var changeMultiaddrsParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ChangeMultiaddrsParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ChangeMultiaddrsParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ChangeMultiaddrsParams) },
	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ChangeMultiaddrsParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ChangeMultiaddrsParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ChangeMultiaddrsParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ChangeMultiaddrsParams) },

	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ChangeMultiaddrsParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ChangeMultiaddrsParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.ChangeMultiaddrsParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.ChangeMultiaddrsParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.ChangeMultiaddrsParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.ChangeMultiaddrsParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.ChangeMultiaddrsParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.ChangeMultiaddrsParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(miner8.ChangeMultiaddrsParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.ChangeMultiaddrsParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.ChangeMultiaddrsParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.ChangeMultiaddrsParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.ChangeMultiaddrsParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.ChangeMultiaddrsParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.ChangeMultiaddrsParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.ChangeMultiaddrsParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.ChangeMultiaddrsParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.ChangeMultiaddrsParams) },
}

var changePeerIDParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ChangePeerIDParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ChangePeerIDParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ChangePeerIDParams) },
	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ChangePeerIDParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ChangePeerIDParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ChangePeerIDParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ChangePeerIDParams) },

	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ChangePeerIDParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ChangePeerIDParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.ChangePeerIDParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.ChangePeerIDParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.ChangePeerIDParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.ChangePeerIDParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.ChangePeerIDParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.ChangePeerIDParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(miner8.ChangePeerIDParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.ChangePeerIDParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.ChangePeerIDParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.ChangePeerIDParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.ChangePeerIDParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.ChangePeerIDParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.ChangePeerIDParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.ChangePeerIDParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.ChangePeerIDParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.ChangePeerIDParams) },
}

var changeWorkerAddressParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ChangeWorkerAddressParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ChangeWorkerAddressParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ChangeWorkerAddressParams) },
	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ChangeWorkerAddressParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ChangeWorkerAddressParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ChangeWorkerAddressParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ChangeWorkerAddressParams) },

	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ChangeWorkerAddressParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ChangeWorkerAddressParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.ChangeWorkerAddressParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.ChangeWorkerAddressParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.ChangeWorkerAddressParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.ChangeWorkerAddressParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.ChangeWorkerAddressParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.ChangeWorkerAddressParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(miner8.ChangeWorkerAddressParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.ChangeWorkerAddressParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.ChangeWorkerAddressParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.ChangeWorkerAddressParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.ChangeWorkerAddressParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.ChangeWorkerAddressParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.ChangeWorkerAddressParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.ChangeWorkerAddressParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.ChangeWorkerAddressParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.ChangeWorkerAddressParams) },
}

var isControllingAddressParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.IsControllingAddressParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.IsControllingAddressParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.IsControllingAddressParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.IsControllingAddressParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.IsControllingAddressParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.IsControllingAddressParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.IsControllingAddressParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.IsControllingAddressParams) },
}

var isControllingAddressReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.IsControllingAddressReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.IsControllingAddressReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.IsControllingAddressReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.IsControllingAddressReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.IsControllingAddressReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.IsControllingAddressReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.IsControllingAddressReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.IsControllingAddressReturn) },
}

var getOwnerReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.GetOwnerReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.GetOwnerReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.GetOwnerReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.GetOwnerReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.GetOwnerReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.GetOwnerReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.GetOwnerReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.GetOwnerReturn) },
}

var getPeerIDReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.GetPeerIDReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.GetPeerIDReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.GetPeerIDReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.GetPeerIDReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.GetPeerIDReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.GetPeerIDReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.GetPeerIDReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.GetPeerIDReturn) },
}

var getMultiAddrsReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.GetMultiAddrsReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.GetMultiAddrsReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.GetMultiAddrsReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.GetMultiAddrsReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.GetMultiAddrsReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.GetMultiAddrsReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.GetMultiAddrsReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.GetMultiAddrsReturn) },
}

var getControlAddressesReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.GetControlAddressesReturn) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.GetControlAddressesReturn) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.GetControlAddressesReturn) },
	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.GetControlAddressesReturn) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.GetControlAddressesReturn) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.GetControlAddressesReturn) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.GetControlAddressesReturn) },

	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.GetControlAddressesReturn) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.GetControlAddressesReturn) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.GetControlAddressesReturn) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.GetControlAddressesReturn) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.GetControlAddressesReturn) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.GetControlAddressesReturn) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.GetControlAddressesReturn) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.GetControlAddressesReturn) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(miner8.GetControlAddressesReturn) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.GetControlAddressesReturn) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.GetControlAddressesReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.GetControlAddressesReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.GetControlAddressesReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.GetControlAddressesReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.GetControlAddressesReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.GetControlAddressesReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.GetControlAddressesReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.GetControlAddressesReturn) },
}

var getAvailableBalanceReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.GetAvailableBalanceReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.GetAvailableBalanceReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.GetAvailableBalanceReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.GetAvailableBalanceReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.GetAvailableBalanceReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.GetAvailableBalanceReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.GetAvailableBalanceReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.GetAvailableBalanceReturn) },
}

var getVestingFundsReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.GetVestingFundsReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.GetVestingFundsReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.GetVestingFundsReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.GetVestingFundsReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.GetVestingFundsReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.GetVestingFundsReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.GetVestingFundsReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.GetVestingFundsReturn) },
}

var getWithdrawBalanceParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.WithdrawBalanceParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.WithdrawBalanceParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.WithdrawBalanceParams) },
	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.WithdrawBalanceParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.WithdrawBalanceParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.WithdrawBalanceParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.WithdrawBalanceParams) },

	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.WithdrawBalanceParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.WithdrawBalanceParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.WithdrawBalanceParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.WithdrawBalanceParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.WithdrawBalanceParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.WithdrawBalanceParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.WithdrawBalanceParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.WithdrawBalanceParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(miner8.WithdrawBalanceParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.WithdrawBalanceParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.WithdrawBalanceParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.WithdrawBalanceParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.WithdrawBalanceParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.WithdrawBalanceParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.WithdrawBalanceParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.WithdrawBalanceParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.WithdrawBalanceParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.WithdrawBalanceParams) },
}

var extendSectorExpiration2Params = map[string]func() cbg.CBORUnmarshaler{
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.ExtendSectorExpiration2Params) },

	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.ExtendSectorExpiration2Params) },
	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.ExtendSectorExpiration2Params) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.ExtendSectorExpiration2Params) },
	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.ExtendSectorExpiration2Params) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.ExtendSectorExpiration2Params) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.ExtendSectorExpiration2Params) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.ExtendSectorExpiration2Params) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.ExtendSectorExpiration2Params) },
}

var preCommitSectorParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.SectorPreCommitInfo) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.SectorPreCommitInfo) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.SectorPreCommitInfo) },
	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.SectorPreCommitInfo) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.SectorPreCommitInfo) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.SectorPreCommitInfo) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.SectorPreCommitInfo) },

	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.SectorPreCommitInfo) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.SectorPreCommitInfo) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.SectorPreCommitInfo) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.SectorPreCommitInfo) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.SectorPreCommitInfo) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.SectorPreCommitInfo) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.SectorPreCommitInfo) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.SectorPreCommitInfo) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(miner8.SectorPreCommitInfo) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.SectorPreCommitInfo) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.SectorPreCommitInfo) },
	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.SectorPreCommitInfo) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.SectorPreCommitInfo) },
	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.SectorPreCommitInfo) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.SectorPreCommitInfo) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.SectorPreCommitInfo) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.SectorPreCommitInfo) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.SectorPreCommitInfo) },
}

var proveCommitSectorParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ProveCommitSectorParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ProveCommitSectorParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ProveCommitSectorParams) },
	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ProveCommitSectorParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ProveCommitSectorParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ProveCommitSectorParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ProveCommitSectorParams) },

	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ProveCommitSectorParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ProveCommitSectorParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.ProveCommitSectorParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.ProveCommitSectorParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.ProveCommitSectorParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.ProveCommitSectorParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.ProveCommitSectorParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.ProveCommitSectorParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(miner8.ProveCommitSectorParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.ProveCommitSectorParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.ProveCommitSectorParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.ProveCommitSectorParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.ProveCommitSectorParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.ProveCommitSectorParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.ProveCommitSectorParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.ProveCommitSectorParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.ProveCommitSectorParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.ProveCommitSectorParams) },
}

var proveCommitSectors3Params = map[string]func() cbg.CBORUnmarshaler{
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.ProveCommitSectors3Params) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.ProveCommitSectors3Params) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.ProveCommitSectors3Params) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.ProveCommitSectors3Params) },
}

var proveCommitSectors3Return = map[string]func() cbg.CBORUnmarshaler{
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.ProveCommitSectors3Return) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.ProveCommitSectors3Return) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.ProveCommitSectors3Return) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.ProveCommitSectors3Return) },
}

var internalSectorSetupForPresealParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.InternalSectorSetupForPresealParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.InternalSectorSetupForPresealParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.InternalSectorSetupForPresealParams) },
}

var submitWindowedPoStParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.SubmitWindowedPoStParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.SubmitWindowedPoStParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.SubmitWindowedPoStParams) },
	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.SubmitWindowedPoStParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.SubmitWindowedPoStParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.SubmitWindowedPoStParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.SubmitWindowedPoStParams) },

	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.SubmitWindowedPoStParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.SubmitWindowedPoStParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.SubmitWindowedPoStParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.SubmitWindowedPoStParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.SubmitWindowedPoStParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.SubmitWindowedPoStParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.SubmitWindowedPoStParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.SubmitWindowedPoStParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(miner8.SubmitWindowedPoStParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.SubmitWindowedPoStParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.SubmitWindowedPoStParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.SubmitWindowedPoStParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.SubmitWindowedPoStParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.SubmitWindowedPoStParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.SubmitWindowedPoStParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.SubmitWindowedPoStParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.SubmitWindowedPoStParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.SubmitWindowedPoStParams) },
}

var confirmSectorProofsParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(miner8.ConfirmSectorProofsParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.ConfirmSectorProofsParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.ConfirmSectorProofsParams) },
	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.ConfirmSectorProofsParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.ConfirmSectorProofsParams) },
	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.ConfirmSectorProofsParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.ConfirmSectorProofsParams) },
}

var checkSectorProvenParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CheckSectorProvenParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CheckSectorProvenParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CheckSectorProvenParams) },
	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CheckSectorProvenParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CheckSectorProvenParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CheckSectorProvenParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CheckSectorProvenParams) },

	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CheckSectorProvenParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CheckSectorProvenParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.CheckSectorProvenParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.CheckSectorProvenParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.CheckSectorProvenParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.CheckSectorProvenParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.CheckSectorProvenParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.CheckSectorProvenParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(miner8.CheckSectorProvenParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.CheckSectorProvenParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.CheckSectorProvenParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.CheckSectorProvenParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.CheckSectorProvenParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.CheckSectorProvenParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.CheckSectorProvenParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.CheckSectorProvenParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.CheckSectorProvenParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.CheckSectorProvenParams) },
}

var extendSectorExpirationParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ExtendSectorExpirationParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ExtendSectorExpirationParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ExtendSectorExpirationParams) },
	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ExtendSectorExpirationParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ExtendSectorExpirationParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ExtendSectorExpirationParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ExtendSectorExpirationParams) },

	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ExtendSectorExpirationParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ExtendSectorExpirationParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.ExtendSectorExpirationParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.ExtendSectorExpirationParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.ExtendSectorExpirationParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.ExtendSectorExpirationParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.ExtendSectorExpirationParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.ExtendSectorExpirationParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(miner8.ExtendSectorExpirationParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.ExtendSectorExpirationParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.ExtendSectorExpirationParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.ExtendSectorExpirationParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.ExtendSectorExpirationParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.ExtendSectorExpirationParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.ExtendSectorExpirationParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.ExtendSectorExpirationParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.ExtendSectorExpirationParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.ExtendSectorExpirationParams) },
}

var compactSectorNumbersParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CompactSectorNumbersParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CompactSectorNumbersParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CompactSectorNumbersParams) },
	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CompactSectorNumbersParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CompactSectorNumbersParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CompactSectorNumbersParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CompactSectorNumbersParams) },

	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CompactSectorNumbersParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CompactSectorNumbersParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.CompactSectorNumbersParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.CompactSectorNumbersParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.CompactSectorNumbersParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.CompactSectorNumbersParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.CompactSectorNumbersParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.CompactSectorNumbersParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(miner8.CompactSectorNumbersParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.CompactSectorNumbersParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.CompactSectorNumbersParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.CompactSectorNumbersParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.CompactSectorNumbersParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.CompactSectorNumbersParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.CompactSectorNumbersParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.CompactSectorNumbersParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.CompactSectorNumbersParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.CompactSectorNumbersParams) },
}

func compactPartitionsParams() map[string]func() cbg.CBORUnmarshaler {
	return map[string]func() cbg.CBORUnmarshaler{
		tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CompactPartitionsParams) },
		tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CompactPartitionsParams) },
		tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CompactPartitionsParams) },
		tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CompactPartitionsParams) },
		tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CompactPartitionsParams) },
		tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CompactPartitionsParams) },
		tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CompactPartitionsParams) },

		tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CompactPartitionsParams) },
		tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CompactPartitionsParams) },

		tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.CompactPartitionsParams) },
		tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.CompactPartitionsParams) },

		tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.CompactPartitionsParams) },
		tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.CompactPartitionsParams) },
		tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.CompactPartitionsParams) },
		tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.CompactPartitionsParams) },
		tools.V16.String(): func() cbg.CBORUnmarshaler { return new(miner8.CompactPartitionsParams) },
		tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.CompactPartitionsParams) },
		tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.CompactPartitionsParams) },

		tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.CompactPartitionsParams) },
		tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.CompactPartitionsParams) },

		tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.CompactPartitionsParams) },
		tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.CompactPartitionsParams) },
		tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.CompactPartitionsParams) },
		tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.CompactPartitionsParams) },
		tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.CompactPartitionsParams) },
	}
}

var preCommitSectorBatchParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.PreCommitSectorBatchParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.PreCommitSectorBatchParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.PreCommitSectorBatchParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(miner8.PreCommitSectorBatchParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.PreCommitSectorBatchParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.PreCommitSectorBatchParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.PreCommitSectorBatchParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.PreCommitSectorBatchParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.PreCommitSectorBatchParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.PreCommitSectorBatchParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.PreCommitSectorBatchParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.PreCommitSectorBatchParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.PreCommitSectorBatchParams) },
}

var proveCommitSectorsNIParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.ProveCommitSectorsNIParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.ProveCommitSectorsNIParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.ProveCommitSectorsNIParams) },
}

var proveCommitSectorsNIReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.ProveCommitSectorsNIReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.ProveCommitSectorsNIReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.ProveCommitSectorsNIReturn) },
}
