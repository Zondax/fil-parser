package verifiedRegistry

import (
	"context"
	"fmt"

	"github.com/zondax/golem/pkg/logger"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/manifest"
	legacyBuiltin "github.com/filecoin-project/specs-actors/actors/builtin"

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
	"github.com/zondax/fil-parser/parser"
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

func (v *VerifiedRegistry) Methods(_ context.Context, network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error) {
	switch {
	// all legacy version
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V15)...):
		return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
			legacyBuiltin.MethodsVerifiedRegistry.Constructor: {
				Name:   parser.MethodConstructor,
				Method: actors.ParseConstructor,
			},
			legacyBuiltin.MethodsVerifiedRegistry.AddVerifier: {
				Name:   parser.MethodAddVerifier,
				Method: v.AddVerifier,
			},
			legacyBuiltin.MethodsVerifiedRegistry.RemoveVerifier: {
				Name:   parser.MethodRemoveVerifier,
				Method: v.RemoveVerifier,
			},
			legacyBuiltin.MethodsVerifiedRegistry.AddVerifiedClient: {
				Name:   parser.MethodAddVerifiedClient,
				Method: v.AddVerifiedClientExported,
			},
			legacyBuiltin.MethodsVerifiedRegistry.UseBytes: {
				Name:   parser.MethodUseBytes,
				Method: v.UseBytes,
			},
			legacyBuiltin.MethodsVerifiedRegistry.RestoreBytes: {
				Name:   parser.MethodRestoreBytes,
				Method: v.RestoreBytes,
			},
		}, nil
	case tools.V16.IsSupported(network, height):
		return verifregv8.Methods, nil
	case tools.V17.IsSupported(network, height):
		return verifregv9.Methods, nil
	case tools.V18.IsSupported(network, height):
		return verifregv10.Methods, nil
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return verifregv11.Methods, nil
	case tools.V21.IsSupported(network, height):
		return verifregv12.Methods, nil
	case tools.V22.IsSupported(network, height):
		return verifregv13.Methods, nil
	case tools.V23.IsSupported(network, height):
		return verifregv14.Methods, nil
	case tools.V24.IsSupported(network, height):
		return verifregv15.Methods, nil
	case tools.V25.IsSupported(network, height):
		return verifregv16.Methods, nil
	default:
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
}

func (*VerifiedRegistry) AddVerifier(network string, height int64, raw []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := addVerifierParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parse(raw, nil, false, params, &abi.EmptyValue{})
}

func (*VerifiedRegistry) RemoveVerifier(network string, height int64, raw []byte) (map[string]interface{}, error) {

	return parse(raw, nil, false, &address.Address{}, &address.Address{})
}

func (*VerifiedRegistry) AddVerifiedClientExported(network string, height int64, raw []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := addVerifiedClientParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parse(raw, nil, false, params, &abi.EmptyValue{})
}

func (*VerifiedRegistry) UseBytes(network string, height int64, raw []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := useBytesParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parse(raw, nil, false, params, &abi.EmptyValue{})
}

func (*VerifiedRegistry) RestoreBytes(network string, height int64, raw []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := restoreBytesParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parse(raw, nil, false, params, &abi.EmptyValue{})
}

func (*VerifiedRegistry) RemoveVerifiedClientDataCap(network string, height int64, raw []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := dataCap()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parse(raw, nil, false, params, &abi.EmptyValue{})
}

func (*VerifiedRegistry) RemoveExpiredAllocationsExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := removeExpiredAllocationsParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := removeExpiredAllocationsReturn()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parse(raw, rawReturn, true, params, returnValue)
}

func (*VerifiedRegistry) Deprecated1(network string, height int64, raw []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := restoreBytesParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parse(raw, nil, false, params, &abi.EmptyValue{})
}

func (*VerifiedRegistry) Deprecated2(network string, height int64, raw []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := useBytesParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parse(raw, nil, false, params, &abi.EmptyValue{})
}

func (*VerifiedRegistry) ClaimAllocations(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := claimAllocationsParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := claimAllocationsReturn()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parse(raw, rawReturn, true, params, returnValue)
}

func (*VerifiedRegistry) GetClaimsExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := getClaimsParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := getClaimsReturn()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parse(raw, rawReturn, true, params, returnValue)
}

func (*VerifiedRegistry) ExtendClaimTermsExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := extendClaimTermsParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := extendClaimTermsReturn()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parse(raw, rawReturn, true, params, returnValue)
}

func (*VerifiedRegistry) RemoveExpiredClaimsExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := removeExpiredClaimsParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := removeExpiredClaimsReturn()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parse(raw, rawReturn, true, params, returnValue)
}

func (*VerifiedRegistry) UniversalReceiverHook(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := universalReceiverParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := allocationsResponse()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parse(raw, rawReturn, true, params, returnValue)
}
