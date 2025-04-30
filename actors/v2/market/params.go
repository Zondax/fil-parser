package market

import (
	"github.com/filecoin-project/go-state-types/abi"
	v10Market "github.com/filecoin-project/go-state-types/builtin/v10/market"
	v11Market "github.com/filecoin-project/go-state-types/builtin/v11/market"
	v12Market "github.com/filecoin-project/go-state-types/builtin/v12/market"
	v13Market "github.com/filecoin-project/go-state-types/builtin/v13/market"
	miner13 "github.com/filecoin-project/go-state-types/builtin/v13/miner"
	v14Market "github.com/filecoin-project/go-state-types/builtin/v14/market"
	miner14 "github.com/filecoin-project/go-state-types/builtin/v14/miner"
	v15Market "github.com/filecoin-project/go-state-types/builtin/v15/market"
	miner15 "github.com/filecoin-project/go-state-types/builtin/v15/miner"
	v16Market "github.com/filecoin-project/go-state-types/builtin/v16/market"
	miner16 "github.com/filecoin-project/go-state-types/builtin/v16/miner"
	v8Market "github.com/filecoin-project/go-state-types/builtin/v8/market"
	v9Market "github.com/filecoin-project/go-state-types/builtin/v9/market"
	legacyv1 "github.com/filecoin-project/specs-actors/actors/builtin/market"
	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/market"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/market"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/market"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/market"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/market"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/market"
	cbg "github.com/whyrusleeping/cbor-gen"
	"github.com/zondax/fil-parser/actors/v2/market/types"
	"github.com/zondax/fil-parser/tools"
)

var withdrawBalanceParams = map[string]func() cbg.CBORUnmarshaler{
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
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(v8Market.WithdrawBalanceParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(v9Market.WithdrawBalanceParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(v10Market.WithdrawBalanceParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(v11Market.WithdrawBalanceParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(v11Market.WithdrawBalanceParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(v12Market.WithdrawBalanceParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(v13Market.WithdrawBalanceParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(v14Market.WithdrawBalanceParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(v15Market.WithdrawBalanceParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(v16Market.WithdrawBalanceParams) },
}

var publishStorageDealsParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.PublishStorageDealsParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.PublishStorageDealsParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.PublishStorageDealsParams) },
	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.PublishStorageDealsParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.PublishStorageDealsParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.PublishStorageDealsParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.PublishStorageDealsParams) },

	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.PublishStorageDealsParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.PublishStorageDealsParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.PublishStorageDealsParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.PublishStorageDealsParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.PublishStorageDealsParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.PublishStorageDealsParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.PublishStorageDealsParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.PublishStorageDealsParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(v8Market.PublishStorageDealsParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(v9Market.PublishStorageDealsParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(v10Market.PublishStorageDealsParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(v11Market.PublishStorageDealsParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(v11Market.PublishStorageDealsParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(v12Market.PublishStorageDealsParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(v13Market.PublishStorageDealsParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(v14Market.PublishStorageDealsParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(v15Market.PublishStorageDealsParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(v16Market.PublishStorageDealsParams) },
}

var publishStorageDealsReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.PublishStorageDealsReturn) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.PublishStorageDealsReturn) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.PublishStorageDealsReturn) },
	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.PublishStorageDealsReturn) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.PublishStorageDealsReturn) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.PublishStorageDealsReturn) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.PublishStorageDealsReturn) },

	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.PublishStorageDealsReturn) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.PublishStorageDealsReturn) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.PublishStorageDealsReturn) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.PublishStorageDealsReturn) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.PublishStorageDealsReturn) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.PublishStorageDealsReturn) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.PublishStorageDealsReturn) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.PublishStorageDealsReturn) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(v8Market.PublishStorageDealsReturn) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(v9Market.PublishStorageDealsReturn) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(v10Market.PublishStorageDealsReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(v11Market.PublishStorageDealsReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(v11Market.PublishStorageDealsReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(v12Market.PublishStorageDealsReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(v13Market.PublishStorageDealsReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(v14Market.PublishStorageDealsReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(v15Market.PublishStorageDealsReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(v16Market.PublishStorageDealsReturn) },
}

var verifyDealsForActivationParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.VerifyDealsForActivationParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.VerifyDealsForActivationParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.VerifyDealsForActivationParams) },
	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.VerifyDealsForActivationParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.VerifyDealsForActivationParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.VerifyDealsForActivationParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.VerifyDealsForActivationParams) },

	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.VerifyDealsForActivationParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.VerifyDealsForActivationParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.VerifyDealsForActivationParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.VerifyDealsForActivationParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.VerifyDealsForActivationParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.VerifyDealsForActivationParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.VerifyDealsForActivationParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.VerifyDealsForActivationParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(v8Market.VerifyDealsForActivationParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(v9Market.VerifyDealsForActivationParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(v10Market.VerifyDealsForActivationParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(v11Market.VerifyDealsForActivationParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(v11Market.VerifyDealsForActivationParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(v12Market.VerifyDealsForActivationParams) },

	// go-state-types impl. of ActivateDealsParams not upto date with builtin-actors
	tools.V22.String(): func() cbg.CBORUnmarshaler { return types.NewVerifyDealsForActivationParams(tools.V22.String()) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return types.NewVerifyDealsForActivationParams(tools.V23.String()) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return types.NewVerifyDealsForActivationParams(tools.V24.String()) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return types.NewVerifyDealsForActivationParams(tools.V25.String()) },
}

var verifyDealsForActivationReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.VerifyDealsForActivationReturn) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.VerifyDealsForActivationReturn) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.VerifyDealsForActivationReturn) },
	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.VerifyDealsForActivationReturn) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.VerifyDealsForActivationReturn) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.VerifyDealsForActivationReturn) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.VerifyDealsForActivationReturn) },

	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.VerifyDealsForActivationReturn) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.VerifyDealsForActivationReturn) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.VerifyDealsForActivationReturn) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.VerifyDealsForActivationReturn) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.VerifyDealsForActivationReturn) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.VerifyDealsForActivationReturn) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.VerifyDealsForActivationReturn) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.VerifyDealsForActivationReturn) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(v8Market.VerifyDealsForActivationReturn) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(v9Market.VerifyDealsForActivationReturn) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(v10Market.VerifyDealsForActivationReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(v11Market.VerifyDealsForActivationReturn) },

	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(v11Market.VerifyDealsForActivationReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(v12Market.VerifyDealsForActivationReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(v13Market.VerifyDealsForActivationReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(v14Market.VerifyDealsForActivationReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(v15Market.VerifyDealsForActivationReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(v16Market.VerifyDealsForActivationReturn) },
}

var activateDealsParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ActivateDealsParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ActivateDealsParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ActivateDealsParams) },
	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ActivateDealsParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ActivateDealsParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ActivateDealsParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ActivateDealsParams) },

	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ActivateDealsParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ActivateDealsParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.ActivateDealsParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.ActivateDealsParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.ActivateDealsParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.ActivateDealsParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.ActivateDealsParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.ActivateDealsParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(v8Market.ActivateDealsParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(v9Market.ActivateDealsParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(v10Market.ActivateDealsParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(v11Market.ActivateDealsParams) },

	// go-state-types impl. of ActivateDealsParams not upto date with builtin-actors
	tools.V20.String(): func() cbg.CBORUnmarshaler { return types.NewBatchActivateDealsParams(tools.V20.String()) },
	tools.V21.String(): func() cbg.CBORUnmarshaler { return types.NewBatchActivateDealsParams(tools.V21.String()) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return types.NewBatchActivateDealsParams(tools.V22.String()) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return types.NewBatchActivateDealsParams(tools.V23.String()) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return types.NewBatchActivateDealsParams(tools.V24.String()) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return types.NewBatchActivateDealsParams(tools.V25.String()) },
}

var activateDealsReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V1.String():  func() cbg.CBORUnmarshaler { return new(abi.EmptyValue) },
	tools.V2.String():  func() cbg.CBORUnmarshaler { return new(abi.EmptyValue) },
	tools.V3.String():  func() cbg.CBORUnmarshaler { return new(abi.EmptyValue) },
	tools.V4.String():  func() cbg.CBORUnmarshaler { return new(abi.EmptyValue) },
	tools.V5.String():  func() cbg.CBORUnmarshaler { return new(abi.EmptyValue) },
	tools.V6.String():  func() cbg.CBORUnmarshaler { return new(abi.EmptyValue) },
	tools.V7.String():  func() cbg.CBORUnmarshaler { return new(abi.EmptyValue) },
	tools.V8.String():  func() cbg.CBORUnmarshaler { return new(abi.EmptyValue) },
	tools.V9.String():  func() cbg.CBORUnmarshaler { return new(abi.EmptyValue) },
	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(abi.EmptyValue) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(abi.EmptyValue) },
	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(abi.EmptyValue) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(abi.EmptyValue) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(abi.EmptyValue) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(abi.EmptyValue) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(abi.EmptyValue) },

	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(v9Market.ActivateDealsResult) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(v10Market.ActivateDealsResult) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(v11Market.ActivateDealsResult) },

	// go-state-types impl. of ActivateDealsParams not upto date with builtin-actors
	tools.V20.String(): func() cbg.CBORUnmarshaler { return types.NewBatchActivateDealsResult(tools.V20.String()) },
	tools.V21.String(): func() cbg.CBORUnmarshaler { return types.NewBatchActivateDealsResult(tools.V21.String()) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return types.NewBatchActivateDealsResult(tools.V22.String()) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return types.NewBatchActivateDealsResult(tools.V23.String()) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return types.NewBatchActivateDealsResult(tools.V24.String()) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return types.NewBatchActivateDealsResult(tools.V25.String()) },
}

var onMinerSectorsTerminateParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.OnMinerSectorsTerminateParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.OnMinerSectorsTerminateParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.OnMinerSectorsTerminateParams) },
	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.OnMinerSectorsTerminateParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.OnMinerSectorsTerminateParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.OnMinerSectorsTerminateParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.OnMinerSectorsTerminateParams) },

	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.OnMinerSectorsTerminateParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.OnMinerSectorsTerminateParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.OnMinerSectorsTerminateParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.OnMinerSectorsTerminateParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.OnMinerSectorsTerminateParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.OnMinerSectorsTerminateParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.OnMinerSectorsTerminateParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.OnMinerSectorsTerminateParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(v8Market.OnMinerSectorsTerminateParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(v9Market.OnMinerSectorsTerminateParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(v10Market.OnMinerSectorsTerminateParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(v11Market.OnMinerSectorsTerminateParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(v11Market.OnMinerSectorsTerminateParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(v12Market.OnMinerSectorsTerminateParams) },

	// go-state-types impl. of OnMinerSectorsTerminateParams not upto date with builtin-actors
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(types.OnMinerSectorsTerminateParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(types.OnMinerSectorsTerminateParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(types.OnMinerSectorsTerminateParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(types.OnMinerSectorsTerminateParams) },
}

var computeDataCommitmentParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ComputeDataCommitmentParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ComputeDataCommitmentParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ComputeDataCommitmentParams) },
	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ComputeDataCommitmentParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ComputeDataCommitmentParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ComputeDataCommitmentParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ComputeDataCommitmentParams) },

	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ComputeDataCommitmentParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ComputeDataCommitmentParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.ComputeDataCommitmentParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.ComputeDataCommitmentParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.ComputeDataCommitmentParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.ComputeDataCommitmentParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.ComputeDataCommitmentParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.ComputeDataCommitmentParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(v8Market.ComputeDataCommitmentParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(v9Market.ComputeDataCommitmentParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(v10Market.ComputeDataCommitmentParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(v11Market.ComputeDataCommitmentParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(v11Market.ComputeDataCommitmentParams) },
}

var computeDataCommitmentReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(cbg.CborCid) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(cbg.CborCid) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(cbg.CborCid) },
	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(cbg.CborCid) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(cbg.CborCid) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(cbg.CborCid) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(cbg.CborCid) },

	tools.V8.String():  func() cbg.CBORUnmarshaler { return new(cbg.CborCid) },
	tools.V9.String():  func() cbg.CBORUnmarshaler { return new(cbg.CborCid) },
	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(cbg.CborCid) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(cbg.CborCid) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(cbg.CborCid) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.ComputeDataCommitmentReturn) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.ComputeDataCommitmentReturn) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.ComputeDataCommitmentReturn) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(v8Market.ComputeDataCommitmentReturn) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(v9Market.ComputeDataCommitmentReturn) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(v10Market.ComputeDataCommitmentReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(v11Market.ComputeDataCommitmentReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(v11Market.ComputeDataCommitmentReturn) },
}

var getBalanceReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(v10Market.GetBalanceReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetBalanceReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetBalanceReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(v12Market.GetBalanceReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(v13Market.GetBalanceReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(v14Market.GetBalanceReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(v15Market.GetBalanceReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(v16Market.GetBalanceReturn) },
}

var getDealDataCommitmentParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(v10Market.GetDealDataCommitmentParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetDealDataCommitmentParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetDealDataCommitmentParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(v12Market.GetDealDataCommitmentParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(v13Market.GetDealDataCommitmentParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(v14Market.GetDealDataCommitmentParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(v15Market.GetDealDataCommitmentParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(v16Market.GetDealDataCommitmentParams) },
}

var getDealDataCommitmentReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(v10Market.GetDealDataCommitmentReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetDealDataCommitmentReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetDealDataCommitmentReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(v12Market.GetDealDataCommitmentReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(v13Market.GetDealDataCommitmentReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(v14Market.GetDealDataCommitmentReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(v15Market.GetDealDataCommitmentReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(v16Market.GetDealDataCommitmentReturn) },
}

var getDealClientParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(v10Market.GetDealClientParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetDealClientParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetDealClientParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(v12Market.GetDealClientParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(v13Market.GetDealClientParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(v14Market.GetDealClientParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(v15Market.GetDealClientParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(v16Market.GetDealClientParams) },
}

var getDealClientReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(v10Market.GetDealClientReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetDealClientReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetDealClientReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(v12Market.GetDealClientReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(v13Market.GetDealClientReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(v14Market.GetDealClientReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(v15Market.GetDealClientReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(v16Market.GetDealClientReturn) },
}

var getDealProviderParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(v10Market.GetDealProviderParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetDealProviderParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetDealProviderParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(v12Market.GetDealProviderParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(v13Market.GetDealProviderParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(v14Market.GetDealProviderParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(v15Market.GetDealProviderParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(v16Market.GetDealProviderParams) },
}

var getDealProviderReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(v10Market.GetDealProviderReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetDealProviderReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetDealProviderReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(v12Market.GetDealProviderReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(v13Market.GetDealProviderReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(v14Market.GetDealProviderReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(v15Market.GetDealProviderReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(v16Market.GetDealProviderReturn) },
}

var getDealLabelParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(v10Market.GetDealLabelParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetDealLabelParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetDealLabelParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(v12Market.GetDealLabelParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(v13Market.GetDealLabelParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(v14Market.GetDealLabelParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(v15Market.GetDealLabelParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(v16Market.GetDealLabelParams) },
}

var getDealLabelReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(v10Market.GetDealLabelReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetDealLabelReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetDealLabelReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(v12Market.GetDealLabelReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(v13Market.GetDealLabelReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(v14Market.GetDealLabelReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(v15Market.GetDealLabelReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(v16Market.GetDealLabelReturn) },
}

var getDealTermParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(v10Market.GetDealTermParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetDealTermParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetDealTermParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(v12Market.GetDealTermParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(v13Market.GetDealTermParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(v14Market.GetDealTermParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(v15Market.GetDealTermParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(v16Market.GetDealTermParams) },
}

var getDealTermReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(v10Market.GetDealTermReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetDealTermReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetDealTermReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(v12Market.GetDealTermReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(v13Market.GetDealTermReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(v14Market.GetDealTermReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(v15Market.GetDealTermReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(v16Market.GetDealTermReturn) },
}

var getDealTotalPriceParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(v10Market.GetDealTotalPriceParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetDealTotalPriceParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetDealTotalPriceParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(v12Market.GetDealTotalPriceParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(v13Market.GetDealTotalPriceParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(v14Market.GetDealTotalPriceParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(v15Market.GetDealTotalPriceParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(v16Market.GetDealTotalPriceParams) },
}

var getDealTotalPriceReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(v10Market.GetDealTotalPriceReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetDealTotalPriceReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetDealTotalPriceReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(v12Market.GetDealTotalPriceReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(v13Market.GetDealTotalPriceReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(v14Market.GetDealTotalPriceReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(v15Market.GetDealTotalPriceReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(v16Market.GetDealTotalPriceReturn) },
}

var getDealClientCollateralParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(v10Market.GetDealClientCollateralParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetDealClientCollateralParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetDealClientCollateralParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(v12Market.GetDealClientCollateralParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(v13Market.GetDealClientCollateralParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(v14Market.GetDealClientCollateralParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(v15Market.GetDealClientCollateralParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(v16Market.GetDealClientCollateralParams) },
}

var getDealClientCollateralReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(v10Market.GetDealClientCollateralReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetDealClientCollateralReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetDealClientCollateralReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(v12Market.GetDealClientCollateralReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(v13Market.GetDealClientCollateralReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(v14Market.GetDealClientCollateralReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(v15Market.GetDealClientCollateralReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(v16Market.GetDealClientCollateralReturn) },
}

var getDealProviderCollateralParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(v10Market.GetDealProviderCollateralParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetDealProviderCollateralParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetDealProviderCollateralParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(v12Market.GetDealProviderCollateralParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(v13Market.GetDealProviderCollateralParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(v14Market.GetDealProviderCollateralParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(v15Market.GetDealProviderCollateralParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(v16Market.GetDealProviderCollateralParams) },
}

var getDealProviderCollateralReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(v10Market.GetDealProviderCollateralReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetDealProviderCollateralReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetDealProviderCollateralReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(v12Market.GetDealProviderCollateralReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(v13Market.GetDealProviderCollateralReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(v14Market.GetDealProviderCollateralReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(v15Market.GetDealProviderCollateralReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(v16Market.GetDealProviderCollateralReturn) },
}

var getDealVerifiedParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(v10Market.GetDealVerifiedParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetDealVerifiedParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetDealVerifiedParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(v12Market.GetDealVerifiedParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(v13Market.GetDealVerifiedParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(v14Market.GetDealVerifiedParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(v15Market.GetDealVerifiedParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(v16Market.GetDealVerifiedParams) },
}

var getDealVerifiedReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(v10Market.GetDealVerifiedReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetDealVerifiedReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetDealVerifiedReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(v12Market.GetDealVerifiedReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(v13Market.GetDealVerifiedReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(v14Market.GetDealVerifiedReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(v15Market.GetDealVerifiedReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(v16Market.GetDealVerifiedReturn) },
}

var getDealActivationParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(v10Market.GetDealActivationParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetDealActivationParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetDealActivationParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(v12Market.GetDealActivationParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(v13Market.GetDealActivationParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(v14Market.GetDealActivationParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(v15Market.GetDealActivationParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(v16Market.GetDealActivationParams) },
}

var getDealActivationReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(v10Market.GetDealActivationReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetDealActivationReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(v11Market.GetDealActivationReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(v12Market.GetDealActivationReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(v13Market.GetDealActivationReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(v14Market.GetDealActivationReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(v15Market.GetDealActivationReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(v16Market.GetDealActivationReturn) },
}

var settleDealPaymentsParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(v13Market.SettleDealPaymentsParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(v14Market.SettleDealPaymentsParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(v15Market.SettleDealPaymentsParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(v16Market.SettleDealPaymentsParams) },
}

var settleDealPaymentsReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(v13Market.SettleDealPaymentsReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(v14Market.SettleDealPaymentsReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(v15Market.SettleDealPaymentsReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(v16Market.SettleDealPaymentsReturn) },
}

var sectorChanges = map[string]func() cbg.CBORUnmarshaler{
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.SectorChanges) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.SectorChanges) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.SectorChanges) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.SectorChanges) },
}

var pieceChange = map[string]func() cbg.CBORUnmarshaler{
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.PieceChange) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.PieceChange) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.PieceChange) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.PieceChange) },
}

var getDealSectorParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(v13Market.GetDealSectorParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(v14Market.GetDealSectorParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(v15Market.GetDealSectorParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(v16Market.GetDealSectorParams) },
}
