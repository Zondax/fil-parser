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

func (*Market) Methods(_ context.Context, network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error) {
	switch {
	// all legacy version
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V15)...):
		return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
			legacyBuiltin.MethodsMarket.Constructor: {
				Name: parser.MethodConstructor,
			},
			legacyBuiltin.MethodsMarket.AddBalance: {
				Name: parser.MethodAddBalance,
			},
			legacyBuiltin.MethodsMarket.WithdrawBalance: {
				Name: parser.MethodWithdrawBalance,
			},
			legacyBuiltin.MethodsMarket.PublishStorageDeals: {
				Name: parser.MethodPublishStorageDeals,
			},
			legacyBuiltin.MethodsMarket.VerifyDealsForActivation: {
				Name: parser.MethodVerifyDealsForActivation,
			},
			legacyBuiltin.MethodsMarket.ActivateDeals: {
				Name: parser.MethodActivateDeals,
			},
			legacyBuiltin.MethodsMarket.OnMinerSectorsTerminate: {
				Name: parser.MethodOnMinerSectorsTerminate,
			},
			legacyBuiltin.MethodsMarket.ComputeDataCommitment: {
				Name: parser.MethodComputeDataCommitment,
			},
			legacyBuiltin.MethodsMarket.CronTick: {
				Name: parser.MethodCronTick,
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
	default:
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
}

func (*Market) AddBalance(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	return parseGeneric(rawParams, nil, false, &address.Address{}, &address.Address{})
}

func (*Market) WithdrawBalance(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	var resp map[string]interface{}
	var err error
	switch {
	case tools.V24.IsSupported(network, height):
		resp, err = parseGeneric(rawParams, nil, false, &v15Market.WithdrawBalanceParams{}, &abi.EmptyValue{})
	case tools.V23.IsSupported(network, height):
		resp, err = parseGeneric(rawParams, nil, false, &v14Market.WithdrawBalanceParams{}, &abi.EmptyValue{})
	case tools.V22.IsSupported(network, height):
		resp, err = parseGeneric(rawParams, nil, false, &v13Market.WithdrawBalanceParams{}, &abi.EmptyValue{})
	case tools.V21.IsSupported(network, height):
		resp, err = parseGeneric(rawParams, nil, false, &v12Market.WithdrawBalanceParams{}, &abi.EmptyValue{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		resp, err = parseGeneric(rawParams, nil, false, &v11Market.WithdrawBalanceParams{}, &abi.EmptyValue{})
	case tools.V18.IsSupported(network, height):
		resp, err = parseGeneric(rawParams, nil, false, &v10Market.WithdrawBalanceParams{}, &abi.EmptyValue{})
	case tools.V17.IsSupported(network, height):
		resp, err = parseGeneric(rawParams, nil, false, &v9Market.WithdrawBalanceParams{}, &abi.EmptyValue{})
	case tools.V16.IsSupported(network, height):
		resp, err = parseGeneric(rawParams, nil, false, &v8Market.WithdrawBalanceParams{}, &abi.EmptyValue{})
	case tools.V15.IsSupported(network, height):
		resp, err = parseGeneric(rawParams, nil, false, &legacyv7.WithdrawBalanceParams{}, &abi.EmptyValue{})
	case tools.V14.IsSupported(network, height):
		resp, err = parseGeneric(rawParams, nil, false, &legacyv6.WithdrawBalanceParams{}, &abi.EmptyValue{})
	case tools.V13.IsSupported(network, height):
		resp, err = parseGeneric(rawParams, nil, false, &legacyv5.WithdrawBalanceParams{}, &abi.EmptyValue{})
	case tools.V12.IsSupported(network, height):
		resp, err = parseGeneric(rawParams, nil, false, &legacyv4.WithdrawBalanceParams{}, &abi.EmptyValue{})
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		resp, err = parseGeneric(rawParams, nil, false, &legacyv3.WithdrawBalanceParams{}, &abi.EmptyValue{})
	case tools.AnyIsSupported(network, height, tools.V8, tools.V9):
		resp, err = parseGeneric(rawParams, nil, false, &legacyv2.WithdrawBalanceParams{}, &abi.EmptyValue{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		resp, err = parseGeneric(rawParams, nil, false, &legacyv1.WithdrawBalanceParams{}, &abi.EmptyValue{})
	default:
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	if err != nil {
		return nil, err
	}
	if rawReturn != nil {
		resp[parser.ReturnKey] = base64.StdEncoding.EncodeToString(rawReturn)
	}
	return resp, nil
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
	case tools.AnyIsSupported(network, height, tools.V8, tools.V9):
		return parseGeneric(rawParams, rawReturn, true, &legacyv2.PublishStorageDealsParams{}, &legacyv2.PublishStorageDealsReturn{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return parseGeneric(rawParams, rawReturn, true, &legacyv1.PublishStorageDealsParams{}, &legacyv1.PublishStorageDealsReturn{})
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Market) VerifyDealsForActivationExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	var resp map[string]interface{}
	var err error

	switch {
	case tools.V24.IsSupported(network, height):
		resp, err = parseGeneric(rawParams, rawReturn, true, &v15Market.VerifyDealsForActivationParams{}, &v15Market.VerifyDealsForActivationReturn{})
	case tools.V23.IsSupported(network, height):
		resp, err = parseGeneric(rawParams, rawReturn, true, &v14Market.VerifyDealsForActivationParams{}, &v14Market.VerifyDealsForActivationReturn{})
	case tools.V22.IsSupported(network, height):
		resp, err = parseGeneric(rawParams, rawReturn, true, &v13Market.VerifyDealsForActivationParams{}, &v13Market.VerifyDealsForActivationReturn{})
	case tools.V21.IsSupported(network, height):
		resp, err = parseGeneric(rawParams, rawReturn, true, &v12Market.VerifyDealsForActivationParams{}, &v12Market.VerifyDealsForActivationReturn{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		resp, err = parseGeneric(rawParams, rawReturn, true, &v11Market.VerifyDealsForActivationParams{}, &v11Market.VerifyDealsForActivationReturn{})
	case tools.V18.IsSupported(network, height):
		resp, err = parseGeneric(rawParams, rawReturn, true, &v10Market.VerifyDealsForActivationParams{}, &v10Market.VerifyDealsForActivationReturn{})
	case tools.V17.IsSupported(network, height):
		resp, err = parseGeneric(rawParams, rawReturn, true, &v9Market.VerifyDealsForActivationParams{}, &v9Market.VerifyDealsForActivationReturn{})

	case tools.V16.IsSupported(network, height):
		resp, err = parseGeneric(rawParams, rawReturn, true, &v8Market.VerifyDealsForActivationParams{}, &v8Market.VerifyDealsForActivationReturn{})
	case tools.V15.IsSupported(network, height):
		resp, err = parseGeneric(rawParams, rawReturn, true, &legacyv7.VerifyDealsForActivationParams{}, &legacyv7.VerifyDealsForActivationReturn{})
	case tools.V14.IsSupported(network, height):
		resp, err = parseGeneric(rawParams, rawReturn, true, &legacyv6.VerifyDealsForActivationParams{}, &legacyv6.VerifyDealsForActivationReturn{})
	case tools.V13.IsSupported(network, height):
		resp, err = parseGeneric(rawParams, rawReturn, true, &legacyv5.VerifyDealsForActivationParams{}, &legacyv5.VerifyDealsForActivationReturn{})
	case tools.V12.IsSupported(network, height):
		resp, err = parseGeneric(rawParams, rawReturn, true, &legacyv4.VerifyDealsForActivationParams{}, &legacyv4.VerifyDealsForActivationReturn{})
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		resp, err = parseGeneric(rawParams, rawReturn, true, &legacyv3.VerifyDealsForActivationParams{}, &legacyv3.VerifyDealsForActivationReturn{})
	case tools.AnyIsSupported(network, height, tools.V8, tools.V9):
		resp, err = parseGeneric(rawParams, rawReturn, true, &legacyv2.VerifyDealsForActivationParams{}, &legacyv2.VerifyDealsForActivationReturn{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		resp, err = parseGeneric(rawParams, rawReturn, true, &legacyv1.VerifyDealsForActivationParams{}, &legacyv1.VerifyDealsForActivationReturn{})
	default:
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return resp, err
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
		return parseGeneric(rawParams, rawReturn, true, &v8Market.ActivateDealsParams{}, &abi.EmptyValue{})
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &legacyv7.ActivateDealsParams{}, &abi.EmptyValue{})
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &legacyv6.ActivateDealsParams{}, &abi.EmptyValue{})
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &legacyv5.ActivateDealsParams{}, &abi.EmptyValue{})
	case tools.V12.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &legacyv4.ActivateDealsParams{}, &abi.EmptyValue{})
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parseGeneric(rawParams, rawReturn, true, &legacyv3.ActivateDealsParams{}, &abi.EmptyValue{})
	case tools.AnyIsSupported(network, height, tools.V8, tools.V9):
		return parseGeneric(rawParams, rawReturn, true, &legacyv2.ActivateDealsParams{}, &abi.EmptyValue{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return parseGeneric(rawParams, rawReturn, true, &legacyv1.ActivateDealsParams{}, &abi.EmptyValue{})
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Market) OnMinerSectorsTerminateExported(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &v15Market.OnMinerSectorsTerminateParams{}, &abi.EmptyValue{})
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &v14Market.OnMinerSectorsTerminateParams{}, &abi.EmptyValue{})
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &v13Market.OnMinerSectorsTerminateParams{}, &abi.EmptyValue{})
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &v12Market.OnMinerSectorsTerminateParams{}, &abi.EmptyValue{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &v11Market.OnMinerSectorsTerminateParams{}, &abi.EmptyValue{})
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &v10Market.OnMinerSectorsTerminateParams{}, &abi.EmptyValue{})
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &v9Market.OnMinerSectorsTerminateParams{}, &abi.EmptyValue{})
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &v8Market.OnMinerSectorsTerminateParams{}, &abi.EmptyValue{})
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv7.OnMinerSectorsTerminateParams{}, &abi.EmptyValue{})
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv6.OnMinerSectorsTerminateParams{}, &abi.EmptyValue{})
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv5.OnMinerSectorsTerminateParams{}, &abi.EmptyValue{})
	case tools.V12.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv4.OnMinerSectorsTerminateParams{}, &abi.EmptyValue{})
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parseGeneric(rawParams, nil, false, &legacyv3.OnMinerSectorsTerminateParams{}, &abi.EmptyValue{})
	case tools.AnyIsSupported(network, height, tools.V8, tools.V9):
		return parseGeneric(rawParams, nil, false, &legacyv2.OnMinerSectorsTerminateParams{}, &abi.EmptyValue{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return parseGeneric(rawParams, nil, false, &legacyv1.OnMinerSectorsTerminateParams{}, &abi.EmptyValue{})
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
	case tools.AnyIsSupported(network, height, tools.V8, tools.V9):
		return parseGeneric(rawParams, rawReturn, true, &legacyv2.ComputeDataCommitmentParams{}, &cbg.CborCid{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return parseGeneric(rawParams, rawReturn, true, &legacyv1.ComputeDataCommitmentParams{}, &cbg.CborCid{})
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

func (*Market) SettleDealPaymentsExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &v15Market.SettleDealPaymentsParams{}, &v15Market.SettleDealPaymentsReturn{})
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &v14Market.SettleDealPaymentsParams{}, &v14Market.SettleDealPaymentsReturn{})
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &v13Market.SettleDealPaymentsParams{}, &v13Market.SettleDealPaymentsReturn{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V21)...):
		return nil, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
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
	switch {
	case tools.V24.IsSupported(network, height):
		return &miner15.SectorChanges{}, nil
	case tools.V23.IsSupported(network, height):
		return &miner14.SectorChanges{}, nil
	case tools.V22.IsSupported(network, height):
		return &miner13.SectorChanges{}, nil
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V21)...):
		return nil, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}
func sectorContentChangedReturn(network string, height int64) (cbg.CBORUnmarshaler, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return &miner15.PieceChange{}, nil
	case tools.V23.IsSupported(network, height):
		return &miner14.PieceChange{}, nil
	case tools.V22.IsSupported(network, height):
		return &miner13.PieceChange{}, nil
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V21)...):
		return nil, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Market) GetDealSectorExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := map[string]interface{}{}
	var extractedParams marketParam
	var extractedReturn abi.SectorNumber

	switch {
	case tools.V24.IsSupported(network, height):
		var params v15Market.GetDealSectorParams
		extractedParams = &params
	case tools.V23.IsSupported(network, height):
		var params v14Market.GetDealSectorParams
		extractedParams = &params
	case tools.V22.IsSupported(network, height):
		var params v13Market.GetDealSectorParams
		extractedParams = &params
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V21)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
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
