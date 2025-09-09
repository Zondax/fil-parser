package market

import (
	"bytes"
	"context"
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/manifest"
	cbg "github.com/whyrusleeping/cbor-gen"
	"github.com/zondax/golem/pkg/logger"

	v10Market "github.com/filecoin-project/go-state-types/builtin/v10/market"
	v11Market "github.com/filecoin-project/go-state-types/builtin/v11/market"
	v12Market "github.com/filecoin-project/go-state-types/builtin/v12/market"
	v13Market "github.com/filecoin-project/go-state-types/builtin/v13/market"
	v14Market "github.com/filecoin-project/go-state-types/builtin/v14/market"
	v15Market "github.com/filecoin-project/go-state-types/builtin/v15/market"
	v16Market "github.com/filecoin-project/go-state-types/builtin/v16/market"
	v17Market "github.com/filecoin-project/go-state-types/builtin/v17/market"
	v8Market "github.com/filecoin-project/go-state-types/builtin/v8/market"
	v9Market "github.com/filecoin-project/go-state-types/builtin/v9/market"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/actors/v2/market/types"
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

var methods = map[string]map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
	tools.V0.String(): v1Methods(),
	tools.V1.String(): v1Methods(),
	tools.V2.String(): v1Methods(),
	tools.V3.String(): v1Methods(),

	tools.V4.String(): v2Methods(),
	tools.V5.String(): v2Methods(),
	tools.V6.String(): v2Methods(),
	tools.V7.String(): v2Methods(),
	tools.V8.String(): v2Methods(),
	tools.V9.String(): v2Methods(),

	tools.V10.String(): v3Methods(),
	tools.V11.String(): v3Methods(),

	tools.V12.String(): v4Methods(),
	tools.V13.String(): v5Methods(),
	tools.V14.String(): v6Methods(),
	tools.V15.String(): v7Methods(),
	tools.V16.String(): actors.CopyMethods(v8Market.Methods),
	tools.V17.String(): actors.CopyMethods(v9Market.Methods),
	tools.V18.String(): actors.CopyMethods(v10Market.Methods),
	tools.V19.String(): actors.CopyMethods(v11Market.Methods),
	tools.V20.String(): actors.CopyMethods(v11Market.Methods),
	tools.V21.String(): actors.CopyMethods(v12Market.Methods),
	tools.V22.String(): actors.CopyMethods(v13Market.Methods),
	tools.V23.String(): actors.CopyMethods(v14Market.Methods),
	tools.V24.String(): actors.CopyMethods(v15Market.Methods),
	tools.V25.String(): actors.CopyMethods(v16Market.Methods),
	tools.V26.String(): actors.CopyMethods(v17Market.Methods),
}

func (m *Market) Methods(_ context.Context, network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error) {
	version := tools.VersionFromHeight(network, height)
	methods, ok := methods[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return methods, nil
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
	resp, err := parseGeneric(rawParams, rawReturn, true, params(), &types.WithdrawBalanceReturn{})
	if err != nil {
		return nil, err
	}
	// if rawReturn != nil {
	// 	resp[parser.ReturnKey] = base64.StdEncoding.EncodeToString(rawReturn)
	// }
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

	metadata, err := parseGeneric(rawParams, rawReturn, true, params(), returnValue())
	if err != nil && metadata[parser.ReturnKey] == nil {
		versions := tools.GetSupportedVersions(network)
		for _, v := range versions {
			returnValue, returnOk := publishStorageDealsReturn[v.String()]
			if !returnOk {
				continue
			}
			metadata, err = parseGeneric(rawParams, rawReturn, true, params(), returnValue())
			if err == nil {
				break
			}
		}
	}
	return metadata, err
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

	metadata, err := parseGeneric(rawParams, rawReturn, true, params(), returnValue())
	if err != nil {
		versions := tools.GetSupportedVersions(network)
		var mErr error
		for _, v := range versions {
			params, paramsOk := verifyDealsForActivationParams[v.String()]
			returnValue, returnOk := verifyDealsForActivationReturn[v.String()]
			if !paramsOk || !returnOk {
				continue
			}
			metadata, mErr = parseGeneric(rawParams, rawReturn, true, params(), returnValue())
			if mErr != nil {
				continue
			}
			break
		}
		return metadata, mErr
	}
	return metadata, err
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

	metadata, err := parseGeneric(rawParams, rawReturn, true, params(), returnValue())
	if err != nil {
		versions := tools.GetSupportedVersions(network)
		for _, v := range versions {
			params, paramsOk := activateDealsParams[v.String()]
			returnValue, returnOk := activateDealsReturn[v.String()]
			if !paramsOk || !returnOk {
				continue
			}
			metadata, err = parseGeneric(rawParams, rawReturn, true, params(), returnValue())
			if err != nil {
				continue
			}
			break
		}
		return metadata, err
	}
	return metadata, nil
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
	metadata, err := parseGeneric(rawParams, rawReturn, true, params(), returnValue())
	if err != nil {
		versions := tools.GetSupportedVersions(network)
		for _, v := range versions {
			params, paramsOk := computeDataCommitmentParams[v.String()]
			returnValue, returnOk := computeDataCommitmentReturn[v.String()]
			if !paramsOk || !returnOk {
				continue
			}
			metadata, err = parseGeneric(rawParams, rawReturn, true, params(), returnValue())
			if err == nil {
				break
			}
		}
	}
	return metadata, err
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

	if len(rawReturn) > 0 {
		r := types.SectorReturn{}
		err = r.UnmarshalCBOR(bytes.NewReader(rawReturn))
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling sector return: %w", err)
		}

		metadata[parser.ReturnKey] = r
	}

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
