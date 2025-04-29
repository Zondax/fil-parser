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
	params, ok := withdrawBalanceParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	resp, err := parseGeneric(rawParams, nil, false, params(), &abi.EmptyValue{})
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
	params, ok := publishStorageDealsParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := publishStorageDealsReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parseGeneric(rawParams, rawReturn, true, params(), returnValue())
}

func (*Market) VerifyDealsForActivationExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := verifyDealsForActivationParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := verifyDealsForActivationReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parseGeneric(rawParams, rawReturn, true, params(), returnValue())
}

func (*Market) ActivateDealsExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := activateDealsParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := activateDealsReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, rawReturn, true, params(), returnValue())

}

func (*Market) OnMinerSectorsTerminateExported(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := onMinerSectorsTerminateParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, nil, false, params(), &abi.EmptyValue{})
}

func (*Market) ComputeDataCommitmentExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := computeDataCommitmentParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := computeDataCommitmentReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, rawReturn, true, params(), returnValue())
}

func (*Market) GetBalanceExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	returnValue, ok := getBalanceReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, rawReturn, true, &address.Address{}, returnValue())
}

func (*Market) GetDealDataCommitmentExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := getDealDataCommitmentParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := getDealDataCommitmentReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, rawReturn, true, params(), returnValue())
}

func (*Market) GetDealClientExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := getDealClientParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := getDealClientReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, rawReturn, true, params(), returnValue())
}

func (*Market) GetDealProviderExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := getDealProviderParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := getDealProviderReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, rawReturn, true, params(), returnValue())
}

func (*Market) GetDealLabelExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := getDealLabelParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := getDealLabelReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, rawReturn, true, params(), returnValue())
}

func (*Market) GetDealTermExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := getDealTermParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := getDealTermReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, rawReturn, true, params(), returnValue())
}

func (*Market) GetDealTotalPriceExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := getDealTotalPriceParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := getDealTotalPriceReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, rawReturn, true, params(), returnValue())
}

func (*Market) GetDealClientCollateralExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := getDealClientCollateralParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := getDealClientCollateralReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, rawReturn, true, params(), returnValue())
}

func (*Market) GetDealProviderCollateralExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := getDealProviderCollateralParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := getDealProviderCollateralReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, rawReturn, true, params(), returnValue())
}

func (*Market) GetDealVerifiedExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := getDealVerifiedParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := getDealVerifiedReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, rawReturn, true, params(), returnValue())
}

func (*Market) GetDealActivationExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := getDealActivationParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := getDealActivationReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, rawReturn, true, params(), returnValue())
}

func (*Market) SettleDealPaymentsExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := settleDealPaymentsParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := settleDealPaymentsReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, rawReturn, true, params(), returnValue())
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
	params, ok := sectorChanges[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return params(), nil
}

func sectorContentChangedReturn(network string, height int64) (cbg.CBORUnmarshaler, error) {
	version := tools.VersionFromHeight(network, height)
	returnValue, ok := pieceChange[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return returnValue(), nil
}

func (*Market) GetDealSectorExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := map[string]interface{}{}
	var extractedReturn abi.SectorNumber

	version := tools.VersionFromHeight(network, height)
	extractedParams, ok := getDealSectorParams[version.String()]
	if !ok {
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	err := extractedParams().UnmarshalCBOR(bytes.NewReader(rawParams))
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
