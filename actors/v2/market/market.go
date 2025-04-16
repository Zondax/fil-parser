package market

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/manifest"
	legacyBuiltin "github.com/filecoin-project/specs-actors/actors/builtin"
	cbg "github.com/whyrusleeping/cbor-gen"
	"github.com/zondax/golem/pkg/logger"

	v10Market "github.com/filecoin-project/go-state-types/builtin/v10/market"
	v11Market "github.com/filecoin-project/go-state-types/builtin/v11/market"
	v12Market "github.com/filecoin-project/go-state-types/builtin/v12/market"
	v13Market "github.com/filecoin-project/go-state-types/builtin/v13/market"
	v14Market "github.com/filecoin-project/go-state-types/builtin/v14/market"
	v15Market "github.com/filecoin-project/go-state-types/builtin/v15/market"
	v16Market "github.com/filecoin-project/go-state-types/builtin/v16/market"
	v8Market "github.com/filecoin-project/go-state-types/builtin/v8/market"
	v9Market "github.com/filecoin-project/go-state-types/builtin/v9/market"

	legacyv1 "github.com/filecoin-project/specs-actors/actors/builtin/market"
	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/market"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/market"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/market"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/market"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/market"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/market"

	miner13 "github.com/filecoin-project/go-state-types/builtin/v13/miner"
	miner14 "github.com/filecoin-project/go-state-types/builtin/v14/miner"
	miner15 "github.com/filecoin-project/go-state-types/builtin/v15/miner"
	miner16 "github.com/filecoin-project/go-state-types/builtin/v16/miner"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

type Market struct {
	logger *logger.Logger
}

func New(logger *logger.Logger) *Market {
	return &Market{
		logger: logger,
	}
}
func (m *Market) Name() string {
	return manifest.MarketKey
}

func (*Market) StartNetworkHeight() int64 {
	return tools.V1.Height()
}

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

func (m *Market) Methods(_ context.Context, network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error) {
	switch {
	// all legacy version
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V15)...):
		return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
			legacyBuiltin.MethodsMarket.Constructor: {
				Name:   parser.MethodConstructor,
				Method: actors.ParseConstructor,
			},
			legacyBuiltin.MethodsMarket.AddBalance: {
				Name:   parser.MethodAddBalance,
				Method: m.AddBalance,
			},
			legacyBuiltin.MethodsMarket.WithdrawBalance: {
				Name:   parser.MethodWithdrawBalance,
				Method: m.WithdrawBalance,
			},
			legacyBuiltin.MethodsMarket.PublishStorageDeals: {
				Name:   parser.MethodPublishStorageDeals,
				Method: m.PublishStorageDealsExported,
			},
			legacyBuiltin.MethodsMarket.VerifyDealsForActivation: {
				Name:   parser.MethodVerifyDealsForActivation,
				Method: m.VerifyDealsForActivationExported,
			},
			legacyBuiltin.MethodsMarket.ActivateDeals: {
				Name:   parser.MethodActivateDeals,
				Method: m.ActivateDealsExported,
			},
			legacyBuiltin.MethodsMarket.OnMinerSectorsTerminate: {
				Name:   parser.MethodOnMinerSectorsTerminate,
				Method: m.OnMinerSectorsTerminateExported,
			},
			legacyBuiltin.MethodsMarket.ComputeDataCommitment: {
				Name:   parser.MethodComputeDataCommitment,
				Method: m.ComputeDataCommitmentExported,
			},
			legacyBuiltin.MethodsMarket.CronTick: {
				Name:   parser.MethodCronTick,
				Method: actors.ParseEmptyParamsAndReturn,
			},
		}, nil
	case tools.V16.IsSupported(network, height):
		return v8Market.Methods, nil
	case tools.V17.IsSupported(network, height):
		return v9Market.Methods, nil
	case tools.V18.IsSupported(network, height):
		return v10Market.Methods, nil
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return v11Market.Methods, nil
	case tools.V21.IsSupported(network, height):
		return v12Market.Methods, nil
	case tools.V22.IsSupported(network, height):
		return v13Market.Methods, nil
	case tools.V23.IsSupported(network, height):
		return v14Market.Methods, nil
	case tools.V24.IsSupported(network, height):
		return v15Market.Methods, nil
	case tools.V25.IsSupported(network, height):
		return v16Market.Methods, nil
	default:
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
}

func (*Market) AddBalance(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	return parseGeneric(rawParams, nil, false, &address.Address{}, &address.Address{})
}

func (*Market) WithdrawBalance(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := withdrawBalanceParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	resp, err := parseGeneric(rawParams, nil, false, params, &abi.EmptyValue{})
	if err != nil {
		return nil, err
	}
	if rawReturn != nil {
		resp[parser.ReturnKey] = base64.StdEncoding.EncodeToString(rawReturn)
	}
	return resp, nil
}

func (*Market) PublishStorageDealsExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := publishStorageDealsParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := publishStorageDealsReturn()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parseGeneric(rawParams, rawReturn, true, params, returnValue)
}

func (*Market) VerifyDealsForActivationExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := verifyDealsForActivationParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := verifyDealsForActivationReturn()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parseGeneric(rawParams, rawReturn, true, params, returnValue)
}

func (*Market) ActivateDealsExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := activateDealsParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := activateDealsReturn()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, rawReturn, true, params, returnValue)

}

func (*Market) OnMinerSectorsTerminateExported(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := onMinerSectorsTerminateParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, nil, false, params, &abi.EmptyValue{})
}

func (*Market) ComputeDataCommitmentExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := computeDataCommitmentParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := computeDataCommitmentReturn()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, rawReturn, true, params, returnValue)
}

func (*Market) GetBalanceExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	returnValue, ok := getBalanceReturn()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, rawReturn, true, &address.Address{}, returnValue)
}

func (*Market) GetDealDataCommitmentExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := getDealDataCommitmentParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := getDealDataCommitmentReturn()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, rawReturn, true, params, returnValue)
}

func (*Market) GetDealClientExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := getDealClientParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := getDealClientReturn()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, rawReturn, true, params, returnValue)
}

func (*Market) GetDealProviderExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := getDealProviderParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := getDealProviderReturn()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, rawReturn, true, params, returnValue)
}

func (*Market) GetDealLabelExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := getDealLabelParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := getDealLabelReturn()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, rawReturn, true, params, returnValue)
}

func (*Market) GetDealTermExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := getDealTermParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := getDealTermReturn()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, rawReturn, true, params, returnValue)
}

func (*Market) GetDealTotalPriceExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := getDealTotalPriceParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := getDealTotalPriceReturn()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, rawReturn, true, params, returnValue)
}

func (*Market) GetDealClientCollateralExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := getDealClientCollateralParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := getDealClientCollateralReturn()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, rawReturn, true, params, returnValue)
}

func (*Market) GetDealProviderCollateralExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := getDealProviderCollateralParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := getDealProviderCollateralReturn()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, rawReturn, true, params, returnValue)
}

func (*Market) GetDealVerifiedExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := getDealVerifiedParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := getDealVerifiedReturn()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, rawReturn, true, params, returnValue)
}

func (*Market) GetDealActivationExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := getDealActivationParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := getDealActivationReturn()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, rawReturn, true, params, returnValue)
}

func (*Market) SettleDealPaymentsExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := settleDealPaymentsParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := settleDealPaymentsReturn()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, rawReturn, true, params, returnValue)
}

func (*Market) SectorContentChanged(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := map[string]interface{}{}

	// Parse params which is an array of SectorChanges
	// SectorChanges does not implement UnmarshalCBOR
	// So we need to parse the individual elements manually
	params, err := parseCBORArray(network, height, rawParams, sectorContentChangedParams)
	if err != nil {
		return nil, fmt.Errorf("error parsing CBOR array: %w", err)
	}

	metadata[parser.ParamsKey] = params
	r, err := parseCBORArray(network, height, rawReturn, sectorContentChangedReturn)
	if err != nil {
		return nil, fmt.Errorf("error parsing CBOR array: %w", err)
	}
	metadata[parser.ReturnKey] = r

	return metadata, nil

}

func sectorContentChangedParams(network string, height int64) (cbg.CBORUnmarshaler, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := sectorChanges()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return params, nil
}

func sectorContentChangedReturn(network string, height int64) (cbg.CBORUnmarshaler, error) {
	version := tools.VersionFromHeight(network, height)
	returnValue, ok := pieceChange()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return returnValue, nil
}

func (*Market) GetDealSectorExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := map[string]interface{}{}
	var extractedReturn abi.SectorNumber

	version := tools.VersionFromHeight(network, height)
	extractedParams, ok := getDealSectorParams()[version.String()]
	if !ok {
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	err := extractedParams.UnmarshalCBOR(bytes.NewReader(rawParams))
	if err != nil {
		return metadata, err
	}

	metadata[parser.ParamsKey] = extractedParams

	sectorNumber, err := abi.ParseUIntKey(string(rawReturn))
	if err != nil {
		return metadata, fmt.Errorf("error parsing return: %w", err)
	}
	extractedReturn = abi.SectorNumber(sectorNumber)

	metadata[parser.ReturnKey] = extractedReturn.String()

	return metadata, nil

}
