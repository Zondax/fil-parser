package market

import (
	"fmt"

	"github.com/filecoin-project/go-address"
	v10Market "github.com/filecoin-project/go-state-types/builtin/v10/market"
	v11Market "github.com/filecoin-project/go-state-types/builtin/v11/market"
	v12Market "github.com/filecoin-project/go-state-types/builtin/v12/market"
	v13Market "github.com/filecoin-project/go-state-types/builtin/v13/market"
	v14Market "github.com/filecoin-project/go-state-types/builtin/v14/market"
	v15Market "github.com/filecoin-project/go-state-types/builtin/v15/market"
	v8Market "github.com/filecoin-project/go-state-types/builtin/v8/market"
	v9Market "github.com/filecoin-project/go-state-types/builtin/v9/market"
	"github.com/filecoin-project/go-state-types/manifest"
	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/market"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/market"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/market"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/market"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/market"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/market"
	cbg "github.com/whyrusleeping/cbor-gen"
	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/tools"
)

type Market struct{}

func (m *Market) Name() string {
	return manifest.MarketKey
}

func (*Market) AddBalance(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	return parseGeneric(rawParams, nil, false, &address.Address{}, &address.Address{})
}

func (*Market) WithdrawBalance(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &v15Market.WithdrawBalanceParams{}, &v15Market.WithdrawBalanceParams{})
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &v14Market.WithdrawBalanceParams{}, &v14Market.WithdrawBalanceParams{})
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &v13Market.WithdrawBalanceParams{}, &v13Market.WithdrawBalanceParams{})
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &v12Market.WithdrawBalanceParams{}, &v12Market.WithdrawBalanceParams{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &v11Market.WithdrawBalanceParams{}, &v11Market.WithdrawBalanceParams{})
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &v10Market.WithdrawBalanceParams{}, &v10Market.WithdrawBalanceParams{})
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &v9Market.WithdrawBalanceParams{}, &v9Market.WithdrawBalanceParams{})
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &v8Market.WithdrawBalanceParams{}, &v8Market.WithdrawBalanceParams{})
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv7.WithdrawBalanceParams{}, &legacyv7.WithdrawBalanceParams{})
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv6.WithdrawBalanceParams{}, &legacyv6.WithdrawBalanceParams{})
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv5.WithdrawBalanceParams{}, &legacyv5.WithdrawBalanceParams{})
	case tools.V12.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv4.WithdrawBalanceParams{}, &legacyv4.WithdrawBalanceParams{})
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parseGeneric(rawParams, nil, false, &legacyv3.WithdrawBalanceParams{}, &legacyv3.WithdrawBalanceParams{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V9)...):
		return parseGeneric(rawParams, nil, false, &legacyv2.WithdrawBalanceParams{}, &legacyv2.WithdrawBalanceParams{})
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Market) PublishStorageDealsExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &v15Market.PublishStorageDealsParams{}, &v15Market.PublishStorageDealsReturn{})
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &v14Market.PublishStorageDealsParams{}, &v14Market.PublishStorageDealsReturn{})
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &v13Market.PublishStorageDealsParams{}, &v13Market.PublishStorageDealsReturn{})
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &v12Market.PublishStorageDealsParams{}, &v12Market.PublishStorageDealsReturn{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, rawReturn, true, &v11Market.PublishStorageDealsParams{}, &v11Market.PublishStorageDealsReturn{})
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &v10Market.PublishStorageDealsParams{}, &v10Market.PublishStorageDealsReturn{})
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &v9Market.PublishStorageDealsParams{}, &v9Market.PublishStorageDealsReturn{})
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &v8Market.PublishStorageDealsParams{}, &v8Market.PublishStorageDealsReturn{})
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &legacyv7.PublishStorageDealsParams{}, &legacyv7.PublishStorageDealsReturn{})
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &legacyv6.PublishStorageDealsParams{}, &legacyv6.PublishStorageDealsReturn{})
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &legacyv5.PublishStorageDealsParams{}, &legacyv5.PublishStorageDealsReturn{})
	case tools.V12.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &legacyv4.PublishStorageDealsParams{}, &legacyv4.PublishStorageDealsReturn{})
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parseGeneric(rawParams, rawReturn, true, &legacyv3.PublishStorageDealsParams{}, &legacyv3.PublishStorageDealsReturn{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V9)...):
		return parseGeneric(rawParams, rawReturn, true, &legacyv2.PublishStorageDealsParams{}, &legacyv2.PublishStorageDealsReturn{})
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Market) VerifyDealsForActivationExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &v15Market.VerifyDealsForActivationParams{}, &v15Market.VerifyDealsForActivationReturn{})
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &v14Market.VerifyDealsForActivationParams{}, &v14Market.VerifyDealsForActivationReturn{})
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &v13Market.VerifyDealsForActivationParams{}, &v13Market.VerifyDealsForActivationReturn{})
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &v12Market.VerifyDealsForActivationParams{}, &v12Market.VerifyDealsForActivationReturn{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, rawReturn, true, &v11Market.VerifyDealsForActivationParams{}, &v11Market.VerifyDealsForActivationReturn{})
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &v10Market.VerifyDealsForActivationParams{}, &v10Market.VerifyDealsForActivationReturn{})
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &v9Market.VerifyDealsForActivationParams{}, &v9Market.VerifyDealsForActivationReturn{})
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &v8Market.VerifyDealsForActivationParams{}, &v8Market.VerifyDealsForActivationReturn{})
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &legacyv7.VerifyDealsForActivationParams{}, &legacyv7.VerifyDealsForActivationReturn{})
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &legacyv6.VerifyDealsForActivationParams{}, &legacyv6.VerifyDealsForActivationReturn{})
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &legacyv5.VerifyDealsForActivationParams{}, &legacyv5.VerifyDealsForActivationReturn{})
	case tools.V12.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &legacyv4.VerifyDealsForActivationParams{}, &legacyv4.VerifyDealsForActivationReturn{})
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parseGeneric(rawParams, rawReturn, true, &legacyv3.VerifyDealsForActivationParams{}, &legacyv3.VerifyDealsForActivationReturn{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V9)...):
		return parseGeneric(rawParams, rawReturn, true, &legacyv2.VerifyDealsForActivationParams{}, &legacyv2.VerifyDealsForActivationReturn{})
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Market) ActivateDealsExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &v15Market.ActivateDealsParams{}, &v15Market.ActivateDealsResult{})
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &v14Market.ActivateDealsParams{}, &v14Market.ActivateDealsResult{})
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &v13Market.ActivateDealsParams{}, &v13Market.ActivateDealsResult{})
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &v12Market.ActivateDealsParams{}, &v12Market.ActivateDealsResult{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, rawReturn, true, &v11Market.ActivateDealsParams{}, &v11Market.ActivateDealsResult{})
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &v10Market.ActivateDealsParams{}, &v10Market.ActivateDealsResult{})
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &v9Market.ActivateDealsParams{}, &v9Market.ActivateDealsResult{})

	// the method used to return an empty value before
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &v8Market.ActivateDealsParams{}, &v8Market.ActivateDealsParams{})
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &legacyv7.ActivateDealsParams{}, &legacyv7.ActivateDealsParams{})
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &legacyv6.ActivateDealsParams{}, &legacyv6.ActivateDealsParams{})
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &legacyv5.ActivateDealsParams{}, &legacyv5.ActivateDealsParams{})
	case tools.V12.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &legacyv4.ActivateDealsParams{}, &legacyv4.ActivateDealsParams{})
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parseGeneric(rawParams, rawReturn, true, &legacyv3.ActivateDealsParams{}, &legacyv3.ActivateDealsParams{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V9)...):
		return parseGeneric(rawParams, rawReturn, true, &legacyv2.ActivateDealsParams{}, &legacyv2.ActivateDealsParams{})
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Market) OnMinerSectorsTerminateExported(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &v15Market.OnMinerSectorsTerminateParams{}, &v15Market.OnMinerSectorsTerminateParams{})
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &v14Market.OnMinerSectorsTerminateParams{}, &v14Market.OnMinerSectorsTerminateParams{})
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &v13Market.OnMinerSectorsTerminateParams{}, &v13Market.OnMinerSectorsTerminateParams{})
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &v12Market.OnMinerSectorsTerminateParams{}, &v12Market.OnMinerSectorsTerminateParams{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &v11Market.OnMinerSectorsTerminateParams{}, &v11Market.OnMinerSectorsTerminateParams{})
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &v10Market.OnMinerSectorsTerminateParams{}, &v10Market.OnMinerSectorsTerminateParams{})
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &v9Market.OnMinerSectorsTerminateParams{}, &v9Market.OnMinerSectorsTerminateParams{})
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &v8Market.OnMinerSectorsTerminateParams{}, &v8Market.OnMinerSectorsTerminateParams{})
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv7.OnMinerSectorsTerminateParams{}, &legacyv7.OnMinerSectorsTerminateParams{})
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv6.OnMinerSectorsTerminateParams{}, &legacyv6.OnMinerSectorsTerminateParams{})
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv5.OnMinerSectorsTerminateParams{}, &legacyv5.OnMinerSectorsTerminateParams{})
	case tools.V12.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv4.OnMinerSectorsTerminateParams{}, &legacyv4.OnMinerSectorsTerminateParams{})
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parseGeneric(rawParams, nil, false, &legacyv3.OnMinerSectorsTerminateParams{}, &legacyv3.OnMinerSectorsTerminateParams{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V9)...):
		return parseGeneric(rawParams, nil, false, &legacyv2.OnMinerSectorsTerminateParams{}, &legacyv2.OnMinerSectorsTerminateParams{})
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Market) ComputeDataCommitmentExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsAfter(tools.V20)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, rawReturn, true, &v11Market.ComputeDataCommitmentParams{}, &v11Market.ComputeDataCommitmentReturn{})
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &v10Market.ComputeDataCommitmentParams{}, &v10Market.ComputeDataCommitmentReturn{})
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &v9Market.ComputeDataCommitmentParams{}, &v9Market.ComputeDataCommitmentReturn{})
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &v8Market.ComputeDataCommitmentParams{}, &v8Market.ComputeDataCommitmentReturn{})
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &legacyv7.ComputeDataCommitmentParams{}, &legacyv7.ComputeDataCommitmentReturn{})
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &legacyv6.ComputeDataCommitmentParams{}, &legacyv6.ComputeDataCommitmentReturn{})
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &legacyv5.ComputeDataCommitmentParams{}, &legacyv5.ComputeDataCommitmentReturn{})
	case tools.V12.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &legacyv4.ComputeDataCommitmentParams{}, &cbg.CborCid{})
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parseGeneric(rawParams, rawReturn, true, &legacyv3.ComputeDataCommitmentParams{}, &cbg.CborCid{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V9)...):
		return parseGeneric(rawParams, rawReturn, true, &legacyv2.ComputeDataCommitmentParams{}, &cbg.CborCid{})
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Market) GetBalanceExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &address.Address{}, &v15Market.GetBalanceReturn{})
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &address.Address{}, &v14Market.GetBalanceReturn{})
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &address.Address{}, &v13Market.GetBalanceReturn{})
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &address.Address{}, &v12Market.GetBalanceReturn{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, rawReturn, true, &address.Address{}, &v11Market.GetBalanceReturn{})
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &address.Address{}, &v10Market.GetBalanceReturn{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V17)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Market) GetDealDataCommitmentExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		var params v15Market.GetDealDataCommitmentParams
		return parseGeneric(rawParams, rawReturn, true, &params, &v15Market.GetDealDataCommitmentReturn{})
	case tools.V23.IsSupported(network, height):
		var params v14Market.GetDealDataCommitmentParams
		return parseGeneric(rawParams, rawReturn, true, &params, &v14Market.GetDealDataCommitmentReturn{})
	case tools.V22.IsSupported(network, height):
		var params v13Market.GetDealDataCommitmentParams
		return parseGeneric(rawParams, rawReturn, true, &params, &v13Market.GetDealDataCommitmentReturn{})
	case tools.V21.IsSupported(network, height):
		var params v12Market.GetDealDataCommitmentParams
		return parseGeneric(rawParams, rawReturn, true, &params, &v12Market.GetDealDataCommitmentReturn{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		var params v11Market.GetDealDataCommitmentParams
		return parseGeneric(rawParams, rawReturn, true, &params, &v11Market.GetDealDataCommitmentReturn{})
	case tools.V18.IsSupported(network, height):
		var params v10Market.GetDealDataCommitmentParams
		return parseGeneric(rawParams, rawReturn, true, &params, &v10Market.GetDealDataCommitmentReturn{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V17)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Market) GetDealClientExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		var params v15Market.GetDealClientParams
		var r v15Market.GetDealClientReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.V23.IsSupported(network, height):
		var params v14Market.GetDealClientParams
		var r v14Market.GetDealClientReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.V22.IsSupported(network, height):
		var params v13Market.GetDealClientParams
		var r v13Market.GetDealClientReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.V21.IsSupported(network, height):
		var params v12Market.GetDealClientParams
		var r v12Market.GetDealClientReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		var params v11Market.GetDealClientParams
		var r v11Market.GetDealClientReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.V18.IsSupported(network, height):
		var params v10Market.GetDealClientParams
		var r v10Market.GetDealClientReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V17)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Market) GetDealProviderExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		var params v15Market.GetDealProviderParams
		var r v15Market.GetDealProviderReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.V23.IsSupported(network, height):
		var params v14Market.GetDealProviderParams
		var r v14Market.GetDealProviderReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.V22.IsSupported(network, height):
		var params v13Market.GetDealProviderParams
		var r v13Market.GetDealProviderReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.V21.IsSupported(network, height):
		var params v12Market.GetDealProviderParams
		var r v12Market.GetDealProviderReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		var params v11Market.GetDealProviderParams
		var r v11Market.GetDealProviderReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.V18.IsSupported(network, height):
		var params v10Market.GetDealProviderParams
		var r v10Market.GetDealProviderReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V17)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Market) GetDealLabelExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		var params v15Market.GetDealLabelParams
		var r v15Market.GetDealLabelReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.V23.IsSupported(network, height):
		var params v14Market.GetDealLabelParams
		var r v14Market.GetDealLabelReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.V22.IsSupported(network, height):
		var params v13Market.GetDealLabelParams
		var r v13Market.GetDealLabelReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.V21.IsSupported(network, height):
		var params v12Market.GetDealLabelParams
		var r v12Market.GetDealLabelReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		var params v11Market.GetDealLabelParams
		var r v11Market.GetDealLabelReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.V18.IsSupported(network, height):
		var params v10Market.GetDealLabelParams
		var r v10Market.GetDealLabelReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V17)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Market) GetDealTermExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		var params v15Market.GetDealTermParams
		var r v15Market.GetDealTermReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.V23.IsSupported(network, height):
		var params v14Market.GetDealTermParams
		var r v14Market.GetDealTermReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.V22.IsSupported(network, height):
		var params v13Market.GetDealTermParams
		var r v13Market.GetDealTermReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.V21.IsSupported(network, height):
		var params v12Market.GetDealTermParams
		var r v12Market.GetDealTermReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		var params v11Market.GetDealTermParams
		var r v11Market.GetDealTermReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.V18.IsSupported(network, height):
		var params v10Market.GetDealTermParams
		var r v10Market.GetDealTermReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V17)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Market) GetDealTotalPriceExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		var params v15Market.GetDealTotalPriceParams
		var r v15Market.GetDealTotalPriceReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.V23.IsSupported(network, height):
		var params v14Market.GetDealTotalPriceParams
		var r v14Market.GetDealTotalPriceReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.V22.IsSupported(network, height):
		var params v13Market.GetDealTotalPriceParams
		var r v13Market.GetDealTotalPriceReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.V21.IsSupported(network, height):
		var params v12Market.GetDealTotalPriceParams
		var r v12Market.GetDealTotalPriceReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		var params v11Market.GetDealTotalPriceParams
		var r v11Market.GetDealTotalPriceReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.V18.IsSupported(network, height):
		var params v10Market.GetDealTotalPriceParams
		var r v10Market.GetDealTotalPriceReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V17)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func (*Market) GetDealClientCollateralExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		var params v15Market.GetDealClientCollateralParams
		var r v15Market.GetDealClientCollateralReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.V23.IsSupported(network, height):
		var params v14Market.GetDealClientCollateralParams
		var r v14Market.GetDealClientCollateralReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.V22.IsSupported(network, height):
		var params v13Market.GetDealClientCollateralParams
		var r v13Market.GetDealClientCollateralReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.V21.IsSupported(network, height):
		var params v12Market.GetDealClientCollateralParams
		var r v12Market.GetDealClientCollateralReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		var params v11Market.GetDealClientCollateralParams
		var r v11Market.GetDealClientCollateralReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.V18.IsSupported(network, height):
		var params v10Market.GetDealClientCollateralParams
		var r v10Market.GetDealClientCollateralReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V17)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func (*Market) GetDealProviderCollateralExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		var params v15Market.GetDealProviderCollateralParams
		var r v15Market.GetDealProviderCollateralReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.V23.IsSupported(network, height):
		var params v14Market.GetDealProviderCollateralParams
		var r v14Market.GetDealProviderCollateralReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.V22.IsSupported(network, height):
		var params v13Market.GetDealProviderCollateralParams
		var r v13Market.GetDealProviderCollateralReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.V21.IsSupported(network, height):
		var params v12Market.GetDealProviderCollateralParams
		var r v12Market.GetDealProviderCollateralReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.AnyIsSupported(network, height, tools.V20, tools.V19):
		var params v11Market.GetDealProviderCollateralParams
		var r v11Market.GetDealProviderCollateralReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.V18.IsSupported(network, height):
		var params v10Market.GetDealProviderCollateralParams
		var r v10Market.GetDealProviderCollateralReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V17)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Market) GetDealVerifiedExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		var params v15Market.GetDealVerifiedParams
		var r v15Market.GetDealVerifiedReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.V23.IsSupported(network, height):
		var params v14Market.GetDealVerifiedParams
		var r v14Market.GetDealVerifiedReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.V22.IsSupported(network, height):
		var params v13Market.GetDealVerifiedParams
		var r v13Market.GetDealVerifiedReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.V21.IsSupported(network, height):
		var params v12Market.GetDealVerifiedParams
		var r v12Market.GetDealVerifiedReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		var params v11Market.GetDealVerifiedParams
		var r v11Market.GetDealVerifiedReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.V18.IsSupported(network, height):
		var params v10Market.GetDealVerifiedParams
		var r v10Market.GetDealVerifiedReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V17)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Market) GetDealActivationExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		var params v15Market.GetDealActivationParams
		var r v15Market.GetDealActivationReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.V23.IsSupported(network, height):
		var params v14Market.GetDealActivationParams
		var r v14Market.GetDealActivationReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.V22.IsSupported(network, height):
		var params v13Market.GetDealActivationParams
		var r v13Market.GetDealActivationReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.V21.IsSupported(network, height):
		var params v12Market.GetDealActivationParams
		var r v12Market.GetDealActivationReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		var params v11Market.GetDealActivationParams
		var r v11Market.GetDealActivationReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.V18.IsSupported(network, height):
		var params v10Market.GetDealActivationParams
		var r v10Market.GetDealActivationReturn
		return parseGeneric(rawParams, rawReturn, true, &params, &r)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V17)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}
