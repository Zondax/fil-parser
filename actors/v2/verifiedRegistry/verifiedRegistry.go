package verifiedRegistry

import (
	"context"
	"fmt"

	"github.com/zondax/golem/pkg/logger"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/manifest"

	verifregv10 "github.com/filecoin-project/go-state-types/builtin/v10/verifreg"
	verifregv11 "github.com/filecoin-project/go-state-types/builtin/v11/verifreg"
	verifregv12 "github.com/filecoin-project/go-state-types/builtin/v12/verifreg"
	verifregv13 "github.com/filecoin-project/go-state-types/builtin/v13/verifreg"
	verifregv14 "github.com/filecoin-project/go-state-types/builtin/v14/verifreg"
	verifregv15 "github.com/filecoin-project/go-state-types/builtin/v15/verifreg"
	verifregv16 "github.com/filecoin-project/go-state-types/builtin/v16/verifreg"
	verifregv8 "github.com/filecoin-project/go-state-types/builtin/v8/verifreg"
	verifregv9 "github.com/filecoin-project/go-state-types/builtin/v9/verifreg"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/tools"
)

type VerifiedRegistry struct {
	logger *logger.Logger
}

func New(logger *logger.Logger) *VerifiedRegistry {
	return &VerifiedRegistry{
		logger: logger,
	}
}

func (v *VerifiedRegistry) Name() string {
	return manifest.VerifregKey
}

func (*VerifiedRegistry) StartNetworkHeight() int64 {
	return tools.V1.Height()
}

var methods = map[string]map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
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
	tools.V16.String(): actors.CopyMethods(verifregv8.Methods),
	tools.V17.String(): actors.CopyMethods(verifregv9.Methods),
	tools.V18.String(): actors.CopyMethods(verifregv10.Methods),
	tools.V19.String(): actors.CopyMethods(verifregv11.Methods),
	tools.V20.String(): actors.CopyMethods(verifregv11.Methods),
	tools.V21.String(): actors.CopyMethods(verifregv12.Methods),
	tools.V22.String(): actors.CopyMethods(verifregv13.Methods),
	tools.V23.String(): actors.CopyMethods(verifregv14.Methods),
	tools.V24.String(): actors.CopyMethods(verifregv15.Methods),
	tools.V25.String(): actors.CopyMethods(verifregv16.Methods),
}

func (v *VerifiedRegistry) Methods(_ context.Context, network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error) {
	version := tools.VersionFromHeight(network, height)
	methods, ok := methods[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return methods, nil
}

func (*VerifiedRegistry) AddVerifier(network string, height int64, raw []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := addVerifierParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parse(raw, nil, false, params(), &abi.EmptyValue{})
}

func (*VerifiedRegistry) RemoveVerifier(network string, height int64, raw []byte) (map[string]interface{}, error) {

	return parse(raw, nil, false, &address.Address{}, &address.Address{})
}

func (*VerifiedRegistry) AddVerifiedClientExported(network string, height int64, raw []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := addVerifiedClientParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parse(raw, nil, false, params(), &abi.EmptyValue{})
}

func (*VerifiedRegistry) UseBytes(network string, height int64, raw []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := useBytesParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parse(raw, nil, false, params(), &abi.EmptyValue{})
}

func (*VerifiedRegistry) RestoreBytes(network string, height int64, raw []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := restoreBytesParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parse(raw, nil, false, params(), &abi.EmptyValue{})
}

func (*VerifiedRegistry) RemoveVerifiedClientDataCap(network string, height int64, raw []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := dataCap[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parse(raw, nil, false, params(), &abi.EmptyValue{})
}

func (*VerifiedRegistry) RemoveExpiredAllocationsExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := removeExpiredAllocationsParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := removeExpiredAllocationsReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parse(raw, rawReturn, true, params(), returnValue())
}

func (*VerifiedRegistry) Deprecated1(network string, height int64, raw []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := restoreBytesParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parse(raw, nil, false, params(), &abi.EmptyValue{})
}

func (*VerifiedRegistry) Deprecated2(network string, height int64, raw []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := useBytesParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parse(raw, nil, false, params(), &abi.EmptyValue{})
}

func (*VerifiedRegistry) ClaimAllocations(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := claimAllocationsParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := claimAllocationsReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	metadata, err := parse(raw, rawReturn, true, params(), returnValue())
	if err != nil {
		// fallback
		// There is mixed use of SectorAllocationClaims with flat parameters and with info inside structs
		versions := tools.GetSupportedVersions(network)
		for _, v := range versions {
			params, paramsOk := claimAllocationsParams[v.String()]
			returnValue, returnValueOk := claimAllocationsReturn[v.String()]
			if !paramsOk || !returnValueOk {
				continue
			}
			metadata, err = parse(raw, rawReturn, true, params(), returnValue())
			if err != nil {
				continue
			}
			break
		}
	}
	return metadata, err
}

func (*VerifiedRegistry) GetClaimsExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := getClaimsParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := getClaimsReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parse(raw, rawReturn, true, params(), returnValue())
}

func (*VerifiedRegistry) ExtendClaimTermsExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := extendClaimTermsParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := extendClaimTermsReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parse(raw, rawReturn, true, params(), returnValue())
}

func (*VerifiedRegistry) RemoveExpiredClaimsExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := removeExpiredClaimsParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := removeExpiredClaimsReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parse(raw, rawReturn, true, params(), returnValue())
}

func (*VerifiedRegistry) UniversalReceiverHook(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := universalReceiverParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := allocationsResponse[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parse(raw, rawReturn, true, params(), returnValue())
}
