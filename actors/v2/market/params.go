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
	"github.com/zondax/fil-parser/tools"
)

func withdrawBalanceParams() map[string]cbg.CBORUnmarshaler {
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
		tools.V16.String(): &v8Market.WithdrawBalanceParams{},
		tools.V17.String(): &v9Market.WithdrawBalanceParams{},
		tools.V18.String(): &v10Market.WithdrawBalanceParams{},

		tools.V19.String(): &v11Market.WithdrawBalanceParams{},
		tools.V20.String(): &v11Market.WithdrawBalanceParams{},

		tools.V21.String(): &v12Market.WithdrawBalanceParams{},
		tools.V22.String(): &v13Market.WithdrawBalanceParams{},
		tools.V23.String(): &v14Market.WithdrawBalanceParams{},
		tools.V24.String(): &v15Market.WithdrawBalanceParams{},
		tools.V25.String(): &v16Market.WithdrawBalanceParams{},
	}
}

func publishStorageDealsParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.PublishStorageDealsParams{},

		tools.V8.String(): &legacyv2.PublishStorageDealsParams{},
		tools.V9.String(): &legacyv2.PublishStorageDealsParams{},

		tools.V10.String(): &legacyv3.PublishStorageDealsParams{},
		tools.V11.String(): &legacyv3.PublishStorageDealsParams{},

		tools.V12.String(): &legacyv4.PublishStorageDealsParams{},
		tools.V13.String(): &legacyv5.PublishStorageDealsParams{},
		tools.V14.String(): &legacyv6.PublishStorageDealsParams{},
		tools.V15.String(): &legacyv7.PublishStorageDealsParams{},
		tools.V16.String(): &v8Market.PublishStorageDealsParams{},
		tools.V17.String(): &v9Market.PublishStorageDealsParams{},
		tools.V18.String(): &v10Market.PublishStorageDealsParams{},

		tools.V19.String(): &v11Market.PublishStorageDealsParams{},
		tools.V20.String(): &v11Market.PublishStorageDealsParams{},

		tools.V21.String(): &v12Market.PublishStorageDealsParams{},
		tools.V22.String(): &v13Market.PublishStorageDealsParams{},
		tools.V23.String(): &v14Market.PublishStorageDealsParams{},
		tools.V24.String(): &v15Market.PublishStorageDealsParams{},
		tools.V25.String(): &v16Market.PublishStorageDealsParams{},
	}
}

func publishStorageDealsReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.PublishStorageDealsReturn{},

		tools.V8.String(): &legacyv2.PublishStorageDealsReturn{},
		tools.V9.String(): &legacyv2.PublishStorageDealsReturn{},

		tools.V10.String(): &legacyv3.PublishStorageDealsReturn{},
		tools.V11.String(): &legacyv3.PublishStorageDealsReturn{},

		tools.V12.String(): &legacyv4.PublishStorageDealsReturn{},
		tools.V13.String(): &legacyv5.PublishStorageDealsReturn{},
		tools.V14.String(): &legacyv6.PublishStorageDealsReturn{},
		tools.V15.String(): &legacyv7.PublishStorageDealsReturn{},
		tools.V16.String(): &v8Market.PublishStorageDealsReturn{},
		tools.V17.String(): &v9Market.PublishStorageDealsReturn{},
		tools.V18.String(): &v10Market.PublishStorageDealsReturn{},

		tools.V19.String(): &v11Market.PublishStorageDealsReturn{},
		tools.V20.String(): &v11Market.PublishStorageDealsReturn{},

		tools.V21.String(): &v12Market.PublishStorageDealsReturn{},
		tools.V22.String(): &v13Market.PublishStorageDealsReturn{},
		tools.V23.String(): &v14Market.PublishStorageDealsReturn{},
		tools.V24.String(): &v15Market.PublishStorageDealsReturn{},
		tools.V25.String(): &v16Market.PublishStorageDealsReturn{},
	}
}

func verifyDealsForActivationParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.VerifyDealsForActivationParams{},

		tools.V8.String(): &legacyv2.VerifyDealsForActivationParams{},
		tools.V9.String(): &legacyv2.VerifyDealsForActivationParams{},

		tools.V10.String(): &legacyv3.VerifyDealsForActivationParams{},
		tools.V11.String(): &legacyv3.VerifyDealsForActivationParams{},

		tools.V12.String(): &legacyv4.VerifyDealsForActivationParams{},
		tools.V13.String(): &legacyv5.VerifyDealsForActivationParams{},
		tools.V14.String(): &legacyv6.VerifyDealsForActivationParams{},
		tools.V15.String(): &legacyv7.VerifyDealsForActivationParams{},
		tools.V16.String(): &v8Market.VerifyDealsForActivationParams{},
		tools.V17.String(): &v9Market.VerifyDealsForActivationParams{},
		tools.V18.String(): &v10Market.VerifyDealsForActivationParams{},

		tools.V19.String(): &v11Market.VerifyDealsForActivationParams{},
		tools.V20.String(): &v11Market.VerifyDealsForActivationParams{},

		tools.V21.String(): &v12Market.VerifyDealsForActivationParams{},
		tools.V22.String(): &v13Market.VerifyDealsForActivationParams{},
		tools.V23.String(): &v14Market.VerifyDealsForActivationParams{},
		tools.V24.String(): &v15Market.VerifyDealsForActivationParams{},
		tools.V25.String(): &v16Market.VerifyDealsForActivationParams{},
	}
}

func verifyDealsForActivationReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.VerifyDealsForActivationReturn{},

		tools.V8.String(): &legacyv2.VerifyDealsForActivationReturn{},
		tools.V9.String(): &legacyv2.VerifyDealsForActivationReturn{},

		tools.V10.String(): &legacyv3.VerifyDealsForActivationReturn{},
		tools.V11.String(): &legacyv3.VerifyDealsForActivationReturn{},

		tools.V12.String(): &legacyv4.VerifyDealsForActivationReturn{},
		tools.V13.String(): &legacyv5.VerifyDealsForActivationReturn{},
		tools.V14.String(): &legacyv6.VerifyDealsForActivationReturn{},
		tools.V15.String(): &legacyv7.VerifyDealsForActivationReturn{},
		tools.V16.String(): &v8Market.VerifyDealsForActivationReturn{},
		tools.V17.String(): &v9Market.VerifyDealsForActivationReturn{},
		tools.V18.String(): &v10Market.VerifyDealsForActivationReturn{},

		tools.V19.String(): &v11Market.VerifyDealsForActivationReturn{},
		tools.V20.String(): &v11Market.VerifyDealsForActivationReturn{},

		tools.V21.String(): &v12Market.VerifyDealsForActivationReturn{},
		tools.V22.String(): &v13Market.VerifyDealsForActivationReturn{},
		tools.V23.String(): &v14Market.VerifyDealsForActivationReturn{},
		tools.V24.String(): &v15Market.VerifyDealsForActivationReturn{},
		tools.V25.String(): &v16Market.VerifyDealsForActivationReturn{},
	}
}

func activateDealsParams() map[string]cbg.CBORUnmarshaler {
	vmap := map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.ActivateDealsParams{},

		tools.V8.String(): &legacyv2.ActivateDealsParams{},
		tools.V9.String(): &legacyv2.ActivateDealsParams{},

		tools.V10.String(): &legacyv3.ActivateDealsParams{},
		tools.V11.String(): &legacyv3.ActivateDealsParams{},

		tools.V12.String(): &legacyv4.ActivateDealsParams{},
		tools.V13.String(): &legacyv5.ActivateDealsParams{},
		tools.V14.String(): &legacyv6.ActivateDealsParams{},
		tools.V15.String(): &legacyv7.ActivateDealsParams{},
		tools.V16.String(): &v8Market.ActivateDealsParams{},
		tools.V17.String(): &v9Market.ActivateDealsParams{},
		tools.V18.String(): &v10Market.ActivateDealsParams{},

		tools.V19.String(): &v11Market.ActivateDealsParams{},
		tools.V20.String(): &v11Market.ActivateDealsParams{},

		tools.V21.String(): &v12Market.ActivateDealsParams{},
		tools.V22.String(): &v13Market.ActivateDealsParams{},
		tools.V23.String(): &v14Market.ActivateDealsParams{},
		tools.V24.String(): &v15Market.ActivateDealsParams{},
		tools.V25.String(): &v16Market.ActivateDealsParams{},
	}
	// set all versions to the same value as V7
	versions := tools.VersionsBefore(tools.V6)
	for _, version := range versions {
		vmap[version.String()] = &legacyv1.ActivateDealsParams{}
	}
	return vmap
}

func activateDealsReturn() map[string]cbg.CBORUnmarshaler {
	vmap := map[string]cbg.CBORUnmarshaler{
		tools.V16.String(): &abi.EmptyValue{},

		tools.V17.String(): &v9Market.ActivateDealsResult{},
		tools.V18.String(): &v10Market.ActivateDealsResult{},

		tools.V19.String(): &v11Market.ActivateDealsResult{},
		tools.V20.String(): &v11Market.ActivateDealsResult{},

		tools.V21.String(): &v12Market.ActivateDealsResult{},
		tools.V22.String(): &v13Market.ActivateDealsResult{},
		tools.V23.String(): &v14Market.ActivateDealsResult{},
		tools.V24.String(): &v15Market.ActivateDealsResult{},
		tools.V25.String(): &v16Market.ActivateDealsResult{},
	}
	// set all versions to the same value as V16
	versions := tools.VersionsBefore(tools.V15)
	for _, version := range versions {
		vmap[version.String()] = &abi.EmptyValue{}
	}
	return vmap
}

func onMinerSectorsTerminateParams() map[string]cbg.CBORUnmarshaler {
	vmap := map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.OnMinerSectorsTerminateParams{},

		tools.V8.String(): &legacyv2.OnMinerSectorsTerminateParams{},
		tools.V9.String(): &legacyv2.OnMinerSectorsTerminateParams{},

		tools.V10.String(): &legacyv3.OnMinerSectorsTerminateParams{},
		tools.V11.String(): &legacyv3.OnMinerSectorsTerminateParams{},

		tools.V12.String(): &legacyv4.OnMinerSectorsTerminateParams{},
		tools.V13.String(): &legacyv5.OnMinerSectorsTerminateParams{},
		tools.V14.String(): &legacyv6.OnMinerSectorsTerminateParams{},
		tools.V15.String(): &legacyv7.OnMinerSectorsTerminateParams{},
		tools.V16.String(): &v8Market.OnMinerSectorsTerminateParams{},
		tools.V17.String(): &v9Market.OnMinerSectorsTerminateParams{},
		tools.V18.String(): &v10Market.OnMinerSectorsTerminateParams{},

		tools.V19.String(): &v11Market.OnMinerSectorsTerminateParams{},
		tools.V20.String(): &v11Market.OnMinerSectorsTerminateParams{},

		tools.V21.String(): &v12Market.OnMinerSectorsTerminateParams{},
		tools.V22.String(): &v13Market.OnMinerSectorsTerminateParams{},
		tools.V23.String(): &v14Market.OnMinerSectorsTerminateParams{},
		tools.V24.String(): &v15Market.OnMinerSectorsTerminateParams{},
		tools.V25.String(): &v16Market.OnMinerSectorsTerminateParams{},
	}
	// set all versions to the same value as V7
	versions := tools.VersionsBefore(tools.V6)
	for _, version := range versions {
		vmap[version.String()] = &legacyv1.OnMinerSectorsTerminateParams{}
	}
	return vmap
}

func computeDataCommitmentParams() map[string]cbg.CBORUnmarshaler {
	vmap := map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.ComputeDataCommitmentParams{},

		tools.V8.String(): &legacyv2.ComputeDataCommitmentParams{},
		tools.V9.String(): &legacyv2.ComputeDataCommitmentParams{},

		tools.V10.String(): &legacyv3.ComputeDataCommitmentParams{},
		tools.V11.String(): &legacyv3.ComputeDataCommitmentParams{},

		tools.V12.String(): &legacyv4.ComputeDataCommitmentParams{},
		tools.V13.String(): &legacyv5.ComputeDataCommitmentParams{},
		tools.V14.String(): &legacyv6.ComputeDataCommitmentParams{},
		tools.V15.String(): &legacyv7.ComputeDataCommitmentParams{},
		tools.V16.String(): &v8Market.ComputeDataCommitmentParams{},
		tools.V17.String(): &v9Market.ComputeDataCommitmentParams{},
		tools.V18.String(): &v10Market.ComputeDataCommitmentParams{},

		tools.V19.String(): &v11Market.ComputeDataCommitmentParams{},
		tools.V20.String(): &v11Market.ComputeDataCommitmentParams{},
	}
	// set all versions to the same value as V7
	versions := tools.VersionsBefore(tools.V6)
	for _, version := range versions {
		vmap[version.String()] = &legacyv1.ComputeDataCommitmentParams{}
	}
	return vmap
}

func computeDataCommitmentReturn() map[string]cbg.CBORUnmarshaler {
	vmap := map[string]cbg.CBORUnmarshaler{
		tools.V12.String(): &cbg.CborCid{},
		tools.V13.String(): &legacyv5.ComputeDataCommitmentReturn{},
		tools.V14.String(): &legacyv6.ComputeDataCommitmentReturn{},
		tools.V15.String(): &legacyv7.ComputeDataCommitmentReturn{},
		tools.V16.String(): &v8Market.ComputeDataCommitmentReturn{},
		tools.V17.String(): &v9Market.ComputeDataCommitmentReturn{},
		tools.V18.String(): &v10Market.ComputeDataCommitmentReturn{},

		tools.V19.String(): &v11Market.ComputeDataCommitmentReturn{},
		tools.V20.String(): &v11Market.ComputeDataCommitmentReturn{},
	}
	// set all versions to the same value as V12
	versions := tools.VersionsBefore(tools.V11)
	for _, version := range versions {
		vmap[version.String()] = &cbg.CborCid{}
	}
	return vmap
}

func getBalanceReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V18.String(): &v10Market.GetBalanceReturn{},

		tools.V19.String(): &v11Market.GetBalanceReturn{},
		tools.V20.String(): &v11Market.GetBalanceReturn{},

		tools.V21.String(): &v12Market.GetBalanceReturn{},
		tools.V22.String(): &v13Market.GetBalanceReturn{},
		tools.V23.String(): &v14Market.GetBalanceReturn{},
		tools.V24.String(): &v15Market.GetBalanceReturn{},
		tools.V25.String(): &v16Market.GetBalanceReturn{},
	}
}

func getDealDataCommitmentParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V18.String(): new(v10Market.GetDealDataCommitmentParams),

		tools.V19.String(): new(v11Market.GetDealDataCommitmentParams),
		tools.V20.String(): new(v11Market.GetDealDataCommitmentParams),

		tools.V21.String(): new(v12Market.GetDealDataCommitmentParams),
		tools.V22.String(): new(v13Market.GetDealDataCommitmentParams),
		tools.V23.String(): new(v14Market.GetDealDataCommitmentParams),
		tools.V24.String(): new(v15Market.GetDealDataCommitmentParams),
		tools.V25.String(): new(v16Market.GetDealDataCommitmentParams),
	}
}

func getDealDataCommitmentReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V18.String(): &v10Market.GetDealDataCommitmentReturn{},

		tools.V19.String(): &v11Market.GetDealDataCommitmentReturn{},
		tools.V20.String(): &v11Market.GetDealDataCommitmentReturn{},

		tools.V21.String(): &v12Market.GetDealDataCommitmentReturn{},
		tools.V22.String(): &v13Market.GetDealDataCommitmentReturn{},
		tools.V23.String(): &v14Market.GetDealDataCommitmentReturn{},
		tools.V24.String(): &v15Market.GetDealDataCommitmentReturn{},
		tools.V25.String(): &v16Market.GetDealDataCommitmentReturn{},
	}
}

func getDealClientParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V18.String(): new(v10Market.GetDealClientParams),

		tools.V19.String(): new(v11Market.GetDealClientParams),
		tools.V20.String(): new(v11Market.GetDealClientParams),

		tools.V21.String(): new(v12Market.GetDealClientParams),
		tools.V22.String(): new(v13Market.GetDealClientParams),
		tools.V23.String(): new(v14Market.GetDealClientParams),
		tools.V24.String(): new(v15Market.GetDealClientParams),
		tools.V25.String(): new(v16Market.GetDealClientParams),
	}
}

func getDealClientReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V18.String(): new(v10Market.GetDealClientReturn),

		tools.V19.String(): new(v11Market.GetDealClientReturn),
		tools.V20.String(): new(v11Market.GetDealClientReturn),

		tools.V21.String(): new(v12Market.GetDealClientReturn),
		tools.V22.String(): new(v13Market.GetDealClientReturn),
		tools.V23.String(): new(v14Market.GetDealClientReturn),
		tools.V24.String(): new(v15Market.GetDealClientReturn),
		tools.V25.String(): new(v16Market.GetDealClientReturn),
	}
}

func getDealProviderParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V18.String(): new(v10Market.GetDealProviderParams),

		tools.V19.String(): new(v11Market.GetDealProviderParams),
		tools.V20.String(): new(v11Market.GetDealProviderParams),

		tools.V21.String(): new(v12Market.GetDealProviderParams),
		tools.V22.String(): new(v13Market.GetDealProviderParams),
		tools.V23.String(): new(v14Market.GetDealProviderParams),
		tools.V24.String(): new(v15Market.GetDealProviderParams),
		tools.V25.String(): new(v16Market.GetDealProviderParams),
	}
}

func getDealProviderReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V18.String(): new(v10Market.GetDealProviderReturn),

		tools.V19.String(): new(v11Market.GetDealProviderReturn),
		tools.V20.String(): new(v11Market.GetDealProviderReturn),

		tools.V21.String(): new(v12Market.GetDealProviderReturn),
		tools.V22.String(): new(v13Market.GetDealProviderReturn),
		tools.V23.String(): new(v14Market.GetDealProviderReturn),
		tools.V24.String(): new(v15Market.GetDealProviderReturn),
		tools.V25.String(): new(v16Market.GetDealProviderReturn),
	}
}

func getDealLabelParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V18.String(): new(v10Market.GetDealLabelParams),

		tools.V19.String(): new(v11Market.GetDealLabelParams),
		tools.V20.String(): new(v11Market.GetDealLabelParams),

		tools.V21.String(): new(v12Market.GetDealLabelParams),
		tools.V22.String(): new(v13Market.GetDealLabelParams),
		tools.V23.String(): new(v14Market.GetDealLabelParams),
		tools.V24.String(): new(v15Market.GetDealLabelParams),
		tools.V25.String(): new(v16Market.GetDealLabelParams),
	}
}

func getDealLabelReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V18.String(): new(v10Market.GetDealLabelReturn),

		tools.V19.String(): new(v11Market.GetDealLabelReturn),
		tools.V20.String(): new(v11Market.GetDealLabelReturn),

		tools.V21.String(): new(v12Market.GetDealLabelReturn),
		tools.V22.String(): new(v13Market.GetDealLabelReturn),
		tools.V23.String(): new(v14Market.GetDealLabelReturn),
		tools.V24.String(): new(v15Market.GetDealLabelReturn),
		tools.V25.String(): new(v16Market.GetDealLabelReturn),
	}
}

func getDealTermParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V18.String(): new(v10Market.GetDealTermParams),

		tools.V19.String(): new(v11Market.GetDealTermParams),
		tools.V20.String(): new(v11Market.GetDealTermParams),

		tools.V21.String(): new(v12Market.GetDealTermParams),
		tools.V22.String(): new(v13Market.GetDealTermParams),
		tools.V23.String(): new(v14Market.GetDealTermParams),
		tools.V24.String(): new(v15Market.GetDealTermParams),
		tools.V25.String(): new(v16Market.GetDealTermParams),
	}
}

func getDealTermReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V18.String(): new(v10Market.GetDealTermReturn),

		tools.V19.String(): new(v11Market.GetDealTermReturn),
		tools.V20.String(): new(v11Market.GetDealTermReturn),

		tools.V21.String(): new(v12Market.GetDealTermReturn),
		tools.V22.String(): new(v13Market.GetDealTermReturn),
		tools.V23.String(): new(v14Market.GetDealTermReturn),
		tools.V24.String(): new(v15Market.GetDealTermReturn),
		tools.V25.String(): new(v16Market.GetDealTermReturn),
	}
}

func getDealTotalPriceParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V18.String(): new(v10Market.GetDealTotalPriceParams),

		tools.V19.String(): new(v11Market.GetDealTotalPriceParams),
		tools.V20.String(): new(v11Market.GetDealTotalPriceParams),

		tools.V21.String(): new(v12Market.GetDealTotalPriceParams),
		tools.V22.String(): new(v13Market.GetDealTotalPriceParams),
		tools.V23.String(): new(v14Market.GetDealTotalPriceParams),
		tools.V24.String(): new(v15Market.GetDealTotalPriceParams),
		tools.V25.String(): new(v16Market.GetDealTotalPriceParams),
	}
}

func getDealTotalPriceReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V18.String(): new(v10Market.GetDealTotalPriceReturn),

		tools.V19.String(): new(v11Market.GetDealTotalPriceReturn),
		tools.V20.String(): new(v11Market.GetDealTotalPriceReturn),

		tools.V21.String(): new(v12Market.GetDealTotalPriceReturn),
		tools.V22.String(): new(v13Market.GetDealTotalPriceReturn),
		tools.V23.String(): new(v14Market.GetDealTotalPriceReturn),
		tools.V24.String(): new(v15Market.GetDealTotalPriceReturn),
		tools.V25.String(): new(v16Market.GetDealTotalPriceReturn),
	}
}

func getDealClientCollateralParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V18.String(): new(v10Market.GetDealClientCollateralParams),

		tools.V19.String(): new(v11Market.GetDealClientCollateralParams),
		tools.V20.String(): new(v11Market.GetDealClientCollateralParams),

		tools.V21.String(): new(v12Market.GetDealClientCollateralParams),
		tools.V22.String(): new(v13Market.GetDealClientCollateralParams),
		tools.V23.String(): new(v14Market.GetDealClientCollateralParams),
		tools.V24.String(): new(v15Market.GetDealClientCollateralParams),
		tools.V25.String(): new(v16Market.GetDealClientCollateralParams),
	}
}

func getDealClientCollateralReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V18.String(): new(v10Market.GetDealClientCollateralReturn),

		tools.V19.String(): new(v11Market.GetDealClientCollateralReturn),
		tools.V20.String(): new(v11Market.GetDealClientCollateralReturn),

		tools.V21.String(): new(v12Market.GetDealClientCollateralReturn),
		tools.V22.String(): new(v13Market.GetDealClientCollateralReturn),
		tools.V23.String(): new(v14Market.GetDealClientCollateralReturn),
		tools.V24.String(): new(v15Market.GetDealClientCollateralReturn),
		tools.V25.String(): new(v16Market.GetDealClientCollateralReturn),
	}
}

func getDealProviderCollateralParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V18.String(): new(v10Market.GetDealProviderCollateralParams),

		tools.V19.String(): new(v11Market.GetDealProviderCollateralParams),
		tools.V20.String(): new(v11Market.GetDealProviderCollateralParams),

		tools.V21.String(): new(v12Market.GetDealProviderCollateralParams),
		tools.V22.String(): new(v13Market.GetDealProviderCollateralParams),
		tools.V23.String(): new(v14Market.GetDealProviderCollateralParams),
		tools.V24.String(): new(v15Market.GetDealProviderCollateralParams),
		tools.V25.String(): new(v16Market.GetDealProviderCollateralParams),
	}
}

func getDealProviderCollateralReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V18.String(): new(v10Market.GetDealProviderCollateralReturn),

		tools.V19.String(): new(v11Market.GetDealProviderCollateralReturn),
		tools.V20.String(): new(v11Market.GetDealProviderCollateralReturn),

		tools.V21.String(): new(v12Market.GetDealProviderCollateralReturn),
		tools.V22.String(): new(v13Market.GetDealProviderCollateralReturn),
		tools.V23.String(): new(v14Market.GetDealProviderCollateralReturn),
		tools.V24.String(): new(v15Market.GetDealProviderCollateralReturn),
		tools.V25.String(): new(v16Market.GetDealProviderCollateralReturn),
	}
}

func getDealVerifiedParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V18.String(): new(v10Market.GetDealVerifiedParams),

		tools.V19.String(): new(v11Market.GetDealVerifiedParams),
		tools.V20.String(): new(v11Market.GetDealVerifiedParams),

		tools.V21.String(): new(v12Market.GetDealVerifiedParams),
		tools.V22.String(): new(v13Market.GetDealVerifiedParams),
		tools.V23.String(): new(v14Market.GetDealVerifiedParams),
		tools.V24.String(): new(v15Market.GetDealVerifiedParams),
		tools.V25.String(): new(v16Market.GetDealVerifiedParams),
	}
}

func getDealVerifiedReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V18.String(): new(v10Market.GetDealVerifiedReturn),

		tools.V19.String(): new(v11Market.GetDealVerifiedReturn),
		tools.V20.String(): new(v11Market.GetDealVerifiedReturn),

		tools.V21.String(): new(v12Market.GetDealVerifiedReturn),
		tools.V22.String(): new(v13Market.GetDealVerifiedReturn),
		tools.V23.String(): new(v14Market.GetDealVerifiedReturn),
		tools.V24.String(): new(v15Market.GetDealVerifiedReturn),
		tools.V25.String(): new(v16Market.GetDealVerifiedReturn),
	}
}

func getDealActivationParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V18.String(): new(v10Market.GetDealActivationParams),

		tools.V19.String(): new(v11Market.GetDealActivationParams),
		tools.V20.String(): new(v11Market.GetDealActivationParams),

		tools.V21.String(): new(v12Market.GetDealActivationParams),
		tools.V22.String(): new(v13Market.GetDealActivationParams),
		tools.V23.String(): new(v14Market.GetDealActivationParams),
		tools.V24.String(): new(v15Market.GetDealActivationParams),
		tools.V25.String(): new(v16Market.GetDealActivationParams),
	}
}

func getDealActivationReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V18.String(): new(v10Market.GetDealActivationReturn),

		tools.V19.String(): new(v11Market.GetDealActivationReturn),
		tools.V20.String(): new(v11Market.GetDealActivationReturn),

		tools.V21.String(): new(v12Market.GetDealActivationReturn),
		tools.V22.String(): new(v13Market.GetDealActivationReturn),
		tools.V23.String(): new(v14Market.GetDealActivationReturn),
		tools.V24.String(): new(v15Market.GetDealActivationReturn),
		tools.V25.String(): new(v16Market.GetDealActivationReturn),
	}
}

func settleDealPaymentsParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V22.String(): &v13Market.SettleDealPaymentsParams{},
		tools.V23.String(): &v14Market.SettleDealPaymentsParams{},
		tools.V24.String(): &v15Market.SettleDealPaymentsParams{},
		tools.V25.String(): &v16Market.SettleDealPaymentsParams{},
	}
}

func settleDealPaymentsReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V22.String(): &v13Market.SettleDealPaymentsReturn{},
		tools.V23.String(): &v14Market.SettleDealPaymentsReturn{},
		tools.V24.String(): &v15Market.SettleDealPaymentsReturn{},
		tools.V25.String(): &v16Market.SettleDealPaymentsReturn{},
	}
}

func sectorChanges() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V22.String(): &miner13.SectorChanges{},
		tools.V23.String(): &miner14.SectorChanges{},
		tools.V24.String(): &miner15.SectorChanges{},
		tools.V25.String(): &miner16.SectorChanges{},
	}
}

func pieceChange() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V22.String(): &miner13.PieceChange{},
		tools.V23.String(): &miner14.PieceChange{},
		tools.V24.String(): &miner15.PieceChange{},
		tools.V25.String(): &miner16.PieceChange{},
	}
}

func getDealSectorParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V22.String(): new(v13Market.GetDealSectorParams),
		tools.V23.String(): new(v14Market.GetDealSectorParams),
		tools.V24.String(): new(v15Market.GetDealSectorParams),
		tools.V25.String(): new(v16Market.GetDealSectorParams),
	}
}
