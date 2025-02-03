package verifiedregistry

import (
	"fmt"

	"github.com/filecoin-project/go-address"
	verifregv10 "github.com/filecoin-project/go-state-types/builtin/v10/verifreg"
	verifregv11 "github.com/filecoin-project/go-state-types/builtin/v11/verifreg"
	verifregv12 "github.com/filecoin-project/go-state-types/builtin/v12/verifreg"
	verifregv13 "github.com/filecoin-project/go-state-types/builtin/v13/verifreg"
	verifregv14 "github.com/filecoin-project/go-state-types/builtin/v14/verifreg"
	verifregv15 "github.com/filecoin-project/go-state-types/builtin/v15/verifreg"
	verifregv8 "github.com/filecoin-project/go-state-types/builtin/v8/verifreg"
	verifregv9 "github.com/filecoin-project/go-state-types/builtin/v9/verifreg"
	"github.com/filecoin-project/go-state-types/manifest"
	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/verifreg"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/verifreg"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/verifreg"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/verifreg"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/verifreg"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/verifreg"
	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/tools"
)

type VerifiedRegistry struct{}

func (v *VerifiedRegistry) Name() string {
	return manifest.VerifregKey
}

func (*VerifiedRegistry) AddVerifier(network string, height int64, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv15.AddVerifierParams{}, &verifregv15.AddVerifierParams{})
	case tools.V23.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv14.AddVerifierParams{}, &verifregv14.AddVerifierParams{})
	case tools.V22.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv13.AddVerifierParams{}, &verifregv13.AddVerifierParams{})
	case tools.V21.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv12.AddVerifierParams{}, &verifregv12.AddVerifierParams{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parse(raw, nil, false, &verifregv11.AddVerifierParams{}, &verifregv11.AddVerifierParams{})
	case tools.V18.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv10.AddVerifierParams{}, &verifregv10.AddVerifierParams{})
	case tools.V17.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv9.AddVerifierParams{}, &verifregv9.AddVerifierParams{})
	case tools.V16.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv8.AddVerifierParams{}, &verifregv8.AddVerifierParams{})
	case tools.V15.IsSupported(network, height):
		return parse(raw, nil, false, &legacyv7.AddVerifierParams{}, &legacyv7.AddVerifierParams{})
	case tools.V14.IsSupported(network, height):
		return parse(raw, nil, false, &legacyv6.AddVerifierParams{}, &legacyv6.AddVerifierParams{})
	case tools.V13.IsSupported(network, height):
		return parse(raw, nil, false, &legacyv5.AddVerifierParams{}, &legacyv5.AddVerifierParams{})
	case tools.V12.IsSupported(network, height):
		return parse(raw, nil, false, &legacyv4.AddVerifierParams{}, &legacyv4.AddVerifierParams{})
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parse(raw, nil, false, &legacyv3.AddVerifierParams{}, &legacyv3.AddVerifierParams{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V9)...):
		return parse(raw, nil, false, &legacyv2.AddVerifierParams{}, &legacyv2.AddVerifierParams{})
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*VerifiedRegistry) RemoveVerifier(network string, height int64, raw []byte) (map[string]interface{}, error) {

	return parse(raw, nil, false, &address.Address{}, &address.Address{})
}

func (*VerifiedRegistry) AddVerifiedClientExported(network string, height int64, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv15.AddVerifiedClientParams{}, &verifregv15.AddVerifiedClientParams{})
	case tools.V23.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv14.AddVerifiedClientParams{}, &verifregv14.AddVerifiedClientParams{})
	case tools.V22.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv13.AddVerifiedClientParams{}, &verifregv13.AddVerifiedClientParams{})
	case tools.V21.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv12.AddVerifiedClientParams{}, &verifregv12.AddVerifiedClientParams{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parse(raw, nil, false, &verifregv11.AddVerifiedClientParams{}, &verifregv11.AddVerifiedClientParams{})
	case tools.V18.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv10.AddVerifiedClientParams{}, &verifregv10.AddVerifiedClientParams{})
	case tools.V17.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv9.AddVerifiedClientParams{}, &verifregv9.AddVerifiedClientParams{})
	case tools.V16.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv8.AddVerifiedClientParams{}, &verifregv8.AddVerifiedClientParams{})
	case tools.V15.IsSupported(network, height):
		return parse(raw, nil, false, &legacyv7.AddVerifiedClientParams{}, &legacyv7.AddVerifiedClientParams{})
	case tools.V14.IsSupported(network, height):
		return parse(raw, nil, false, &legacyv6.AddVerifiedClientParams{}, &legacyv6.AddVerifiedClientParams{})
	case tools.V13.IsSupported(network, height):
		return parse(raw, nil, false, &legacyv5.AddVerifiedClientParams{}, &legacyv5.AddVerifiedClientParams{})
	case tools.V12.IsSupported(network, height):
		return parse(raw, nil, false, &legacyv4.AddVerifiedClientParams{}, &legacyv4.AddVerifiedClientParams{})
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parse(raw, nil, false, &legacyv3.AddVerifiedClientParams{}, &legacyv3.AddVerifiedClientParams{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V9)...):
		return parse(raw, nil, false, &legacyv2.AddVerifiedClientParams{}, &legacyv2.AddVerifiedClientParams{})
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*VerifiedRegistry) UseBytes(network string, height int64, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv15.UseBytesParams{}, &verifregv15.UseBytesParams{})
	case tools.V23.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv14.UseBytesParams{}, &verifregv14.UseBytesParams{})
	case tools.V22.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv13.UseBytesParams{}, &verifregv13.UseBytesParams{})
	case tools.V21.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv12.UseBytesParams{}, &verifregv12.UseBytesParams{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parse(raw, nil, false, &verifregv11.UseBytesParams{}, &verifregv11.UseBytesParams{})
	case tools.V18.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv10.UseBytesParams{}, &verifregv10.UseBytesParams{})
	case tools.V17.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv9.UseBytesParams{}, &verifregv9.UseBytesParams{})
	case tools.V16.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv8.UseBytesParams{}, &verifregv8.UseBytesParams{})
	case tools.V15.IsSupported(network, height):
		return parse(raw, nil, false, &legacyv7.UseBytesParams{}, &legacyv7.UseBytesParams{})
	case tools.V14.IsSupported(network, height):
		return parse(raw, nil, false, &legacyv6.UseBytesParams{}, &legacyv6.UseBytesParams{})
	case tools.V13.IsSupported(network, height):
		return parse(raw, nil, false, &legacyv5.UseBytesParams{}, &legacyv5.UseBytesParams{})
	case tools.V12.IsSupported(network, height):
		return parse(raw, nil, false, &legacyv4.UseBytesParams{}, &legacyv4.UseBytesParams{})
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parse(raw, nil, false, &legacyv3.UseBytesParams{}, &legacyv3.UseBytesParams{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V9)...):
		return parse(raw, nil, false, &legacyv2.UseBytesParams{}, &legacyv2.UseBytesParams{})
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*VerifiedRegistry) RestoreBytes(network string, height int64, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv15.RestoreBytesParams{}, &verifregv15.RestoreBytesParams{})
	case tools.V23.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv14.RestoreBytesParams{}, &verifregv14.RestoreBytesParams{})
	case tools.V22.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv13.RestoreBytesParams{}, &verifregv13.RestoreBytesParams{})
	case tools.V21.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv12.RestoreBytesParams{}, &verifregv12.RestoreBytesParams{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parse(raw, nil, false, &verifregv11.RestoreBytesParams{}, &verifregv11.RestoreBytesParams{})
	case tools.V18.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv10.RestoreBytesParams{}, &verifregv10.RestoreBytesParams{})
	case tools.V17.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv9.RestoreBytesParams{}, &verifregv9.RestoreBytesParams{})
	case tools.V16.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv8.RestoreBytesParams{}, &verifregv8.RestoreBytesParams{})
	case tools.V15.IsSupported(network, height):
		return parse(raw, nil, false, &legacyv7.RestoreBytesParams{}, &legacyv7.RestoreBytesParams{})
	case tools.V14.IsSupported(network, height):
		return parse(raw, nil, false, &legacyv6.RestoreBytesParams{}, &legacyv6.RestoreBytesParams{})
	case tools.V13.IsSupported(network, height):
		return parse(raw, nil, false, &legacyv5.RestoreBytesParams{}, &legacyv5.RestoreBytesParams{})
	case tools.V12.IsSupported(network, height):
		return parse(raw, nil, false, &legacyv4.RestoreBytesParams{}, &legacyv4.RestoreBytesParams{})
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parse(raw, nil, false, &legacyv3.RestoreBytesParams{}, &legacyv3.RestoreBytesParams{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V9)...):
		return parse(raw, nil, false, &legacyv2.RestoreBytesParams{}, &legacyv2.RestoreBytesParams{})
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*VerifiedRegistry) RemoveVerifiedClientDataCap(network string, height int64, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv15.DataCap{}, &verifregv15.DataCap{})
	case tools.V23.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv14.DataCap{}, &verifregv14.DataCap{})
	case tools.V22.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv13.DataCap{}, &verifregv13.DataCap{})
	case tools.V21.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv12.DataCap{}, &verifregv12.DataCap{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parse(raw, nil, false, &verifregv11.DataCap{}, &verifregv11.DataCap{})
	case tools.V18.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv10.DataCap{}, &verifregv10.DataCap{})
	case tools.V17.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv9.DataCap{}, &verifregv9.DataCap{})
	case tools.V16.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv8.DataCap{}, &verifregv8.DataCap{})
	case tools.V15.IsSupported(network, height):
		return parse(raw, nil, false, &legacyv7.DataCap{}, &legacyv7.DataCap{})
	case tools.V14.IsSupported(network, height):
		return parse(raw, nil, false, &legacyv6.DataCap{}, &legacyv6.DataCap{})
	case tools.V13.IsSupported(network, height):
		return parse(raw, nil, false, &legacyv5.DataCap{}, &legacyv5.DataCap{})
	case tools.V12.IsSupported(network, height):
		return parse(raw, nil, false, &legacyv4.DataCap{}, &legacyv4.DataCap{})
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parse(raw, nil, false, &legacyv3.DataCap{}, &legacyv3.DataCap{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V9)...):
		return parse(raw, nil, false, &legacyv2.DataCap{}, &legacyv2.DataCap{})
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*VerifiedRegistry) RemoveExpiredAllocationsExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parse(raw, rawReturn, true, &verifregv15.RemoveExpiredAllocationsParams{}, &verifregv15.RemoveExpiredAllocationsReturn{})
	case tools.V23.IsSupported(network, height):
		return parse(raw, rawReturn, true, &verifregv14.RemoveExpiredAllocationsParams{}, &verifregv14.RemoveExpiredAllocationsReturn{})
	case tools.V22.IsSupported(network, height):
		return parse(raw, rawReturn, true, &verifregv13.RemoveExpiredAllocationsParams{}, &verifregv13.RemoveExpiredAllocationsReturn{})
	case tools.V21.IsSupported(network, height):
		return parse(raw, rawReturn, true, &verifregv12.RemoveExpiredAllocationsParams{}, &verifregv12.RemoveExpiredAllocationsReturn{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parse(raw, rawReturn, true, &verifregv11.RemoveExpiredAllocationsParams{}, &verifregv11.RemoveExpiredAllocationsReturn{})
	case tools.V18.IsSupported(network, height):
		return parse(raw, rawReturn, true, &verifregv10.RemoveExpiredAllocationsParams{}, &verifregv10.RemoveExpiredAllocationsReturn{})
	case tools.V17.IsSupported(network, height):
		return parse(raw, rawReturn, true, &verifregv9.RemoveExpiredAllocationsParams{}, &verifregv9.RemoveExpiredAllocationsReturn{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V16)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*VerifiedRegistry) Deprecated1(network string, height int64, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv15.RestoreBytesParams{}, &verifregv15.RestoreBytesParams{})
	case tools.V23.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv14.RestoreBytesParams{}, &verifregv14.RestoreBytesParams{})
	case tools.V22.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv13.RestoreBytesParams{}, &verifregv13.RestoreBytesParams{})
	case tools.V21.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv12.RestoreBytesParams{}, &verifregv12.RestoreBytesParams{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parse(raw, nil, false, &verifregv11.RestoreBytesParams{}, &verifregv11.RestoreBytesParams{})
	case tools.V18.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv10.RestoreBytesParams{}, &verifregv10.RestoreBytesParams{})
	case tools.V17.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv9.RestoreBytesParams{}, &verifregv9.RestoreBytesParams{})
	case tools.V16.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv8.RestoreBytesParams{}, &verifregv8.RestoreBytesParams{})
	case tools.V15.IsSupported(network, height):
		return parse(raw, nil, false, &legacyv7.RestoreBytesParams{}, &legacyv7.RestoreBytesParams{})
	case tools.V14.IsSupported(network, height):
		return parse(raw, nil, false, &legacyv6.RestoreBytesParams{}, &legacyv6.RestoreBytesParams{})
	case tools.V13.IsSupported(network, height):
		return parse(raw, nil, false, &legacyv5.RestoreBytesParams{}, &legacyv5.RestoreBytesParams{})
	case tools.V12.IsSupported(network, height):
		return parse(raw, nil, false, &legacyv4.RestoreBytesParams{}, &legacyv4.RestoreBytesParams{})
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parse(raw, nil, false, &legacyv3.RestoreBytesParams{}, &legacyv3.RestoreBytesParams{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V9)...):
		return parse(raw, nil, false, &legacyv2.RestoreBytesParams{}, &legacyv2.RestoreBytesParams{})
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*VerifiedRegistry) Deprecated2(network string, height int64, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv15.UseBytesParams{}, &verifregv15.UseBytesParams{})
	case tools.V23.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv14.UseBytesParams{}, &verifregv14.UseBytesParams{})
	case tools.V22.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv13.UseBytesParams{}, &verifregv13.UseBytesParams{})
	case tools.V21.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv12.UseBytesParams{}, &verifregv12.UseBytesParams{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parse(raw, nil, false, &verifregv11.UseBytesParams{}, &verifregv11.UseBytesParams{})
	case tools.V18.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv10.UseBytesParams{}, &verifregv10.UseBytesParams{})
	case tools.V17.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv9.UseBytesParams{}, &verifregv9.UseBytesParams{})
	case tools.V16.IsSupported(network, height):
		return parse(raw, nil, false, &verifregv8.UseBytesParams{}, &verifregv8.UseBytesParams{})
	case tools.V15.IsSupported(network, height):
		return parse(raw, nil, false, &legacyv7.UseBytesParams{}, &legacyv7.UseBytesParams{})
	case tools.V14.IsSupported(network, height):
		return parse(raw, nil, false, &legacyv6.UseBytesParams{}, &legacyv6.UseBytesParams{})
	case tools.V13.IsSupported(network, height):
		return parse(raw, nil, false, &legacyv5.UseBytesParams{}, &legacyv5.UseBytesParams{})
	case tools.V12.IsSupported(network, height):
		return parse(raw, nil, false, &legacyv4.UseBytesParams{}, &legacyv4.UseBytesParams{})
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parse(raw, nil, false, &legacyv3.UseBytesParams{}, &legacyv3.UseBytesParams{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V9)...):
		return parse(raw, nil, false, &legacyv2.UseBytesParams{}, &legacyv2.UseBytesParams{})
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*VerifiedRegistry) ClaimAllocations(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parse(raw, rawReturn, true, &verifregv15.ClaimAllocationsParams{}, &verifregv15.ClaimAllocationsReturn{})
	case tools.V23.IsSupported(network, height):
		return parse(raw, rawReturn, true, &verifregv14.ClaimAllocationsParams{}, &verifregv14.ClaimAllocationsReturn{})
	case tools.V22.IsSupported(network, height):
		return parse(raw, rawReturn, true, &verifregv13.ClaimAllocationsParams{}, &verifregv13.ClaimAllocationsReturn{})
	case tools.V21.IsSupported(network, height):
		return parse(raw, rawReturn, true, &verifregv12.ClaimAllocationsParams{}, &verifregv12.ClaimAllocationsReturn{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parse(raw, rawReturn, true, &verifregv11.ClaimAllocationsParams{}, &verifregv11.ClaimAllocationsReturn{})
	case tools.V18.IsSupported(network, height):
		return parse(raw, rawReturn, true, &verifregv10.ClaimAllocationsParams{}, &verifregv10.ClaimAllocationsReturn{})
	case tools.V17.IsSupported(network, height):
		return parse(raw, rawReturn, true, &verifregv9.ClaimAllocationsParams{}, &verifregv9.ClaimAllocationsReturn{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V16)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*VerifiedRegistry) GetClaimsExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parse(raw, rawReturn, true, &verifregv15.GetClaimsParams{}, &verifregv15.GetClaimsReturn{})
	case tools.V23.IsSupported(network, height):
		return parse(raw, rawReturn, true, &verifregv14.GetClaimsParams{}, &verifregv14.GetClaimsReturn{})
	case tools.V22.IsSupported(network, height):
		return parse(raw, rawReturn, true, &verifregv13.GetClaimsParams{}, &verifregv13.GetClaimsReturn{})
	case tools.V21.IsSupported(network, height):
		return parse(raw, rawReturn, true, &verifregv12.GetClaimsParams{}, &verifregv12.GetClaimsReturn{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parse(raw, rawReturn, true, &verifregv11.GetClaimsParams{}, &verifregv11.GetClaimsReturn{})
	case tools.V18.IsSupported(network, height):
		return parse(raw, rawReturn, true, &verifregv10.GetClaimsParams{}, &verifregv10.GetClaimsReturn{})
	case tools.V17.IsSupported(network, height):
		return parse(raw, rawReturn, true, &verifregv9.GetClaimsParams{}, &verifregv9.GetClaimsReturn{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V16)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*VerifiedRegistry) ExtendClaimTermsExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parse(raw, rawReturn, true, &verifregv15.ExtendClaimTermsParams{}, &verifregv15.ExtendClaimTermsReturn{})
	case tools.V23.IsSupported(network, height):
		return parse(raw, rawReturn, true, &verifregv14.ExtendClaimTermsParams{}, &verifregv14.ExtendClaimTermsReturn{})
	case tools.V22.IsSupported(network, height):
		return parse(raw, rawReturn, true, &verifregv13.ExtendClaimTermsParams{}, &verifregv13.ExtendClaimTermsReturn{})
	case tools.V21.IsSupported(network, height):
		return parse(raw, rawReturn, true, &verifregv12.ExtendClaimTermsParams{}, &verifregv12.ExtendClaimTermsReturn{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parse(raw, rawReturn, true, &verifregv11.ExtendClaimTermsParams{}, &verifregv11.ExtendClaimTermsReturn{})
	case tools.V18.IsSupported(network, height):
		return parse(raw, rawReturn, true, &verifregv10.ExtendClaimTermsParams{}, &verifregv10.ExtendClaimTermsReturn{})
	case tools.V17.IsSupported(network, height):
		return parse(raw, rawReturn, true, &verifregv9.ExtendClaimTermsParams{}, &verifregv9.ExtendClaimTermsReturn{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V16)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*VerifiedRegistry) RemoveExpiredClaimsExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parse(raw, rawReturn, true, &verifregv15.RemoveExpiredClaimsParams{}, &verifregv15.RemoveExpiredClaimsReturn{})
	case tools.V23.IsSupported(network, height):
		return parse(raw, rawReturn, true, &verifregv14.RemoveExpiredClaimsParams{}, &verifregv14.RemoveExpiredClaimsReturn{})
	case tools.V22.IsSupported(network, height):
		return parse(raw, rawReturn, true, &verifregv13.RemoveExpiredClaimsParams{}, &verifregv13.RemoveExpiredClaimsReturn{})
	case tools.V21.IsSupported(network, height):
		return parse(raw, rawReturn, true, &verifregv12.RemoveExpiredClaimsParams{}, &verifregv12.RemoveExpiredClaimsReturn{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parse(raw, rawReturn, true, &verifregv11.RemoveExpiredClaimsParams{}, &verifregv11.RemoveExpiredClaimsReturn{})
	case tools.V18.IsSupported(network, height):
		return parse(raw, rawReturn, true, &verifregv10.RemoveExpiredClaimsParams{}, &verifregv10.RemoveExpiredClaimsReturn{})
	case tools.V17.IsSupported(network, height):
		return parse(raw, rawReturn, true, &verifregv9.RemoveExpiredClaimsParams{}, &verifregv9.RemoveExpiredClaimsReturn{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V16)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*VerifiedRegistry) UniversalReceiverHook(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parse(raw, rawReturn, true, &verifregv15.UniversalReceiverParams{}, &verifregv15.AllocationsResponse{})
	case tools.V23.IsSupported(network, height):
		return parse(raw, rawReturn, true, &verifregv14.UniversalReceiverParams{}, &verifregv14.AllocationsResponse{})
	case tools.V22.IsSupported(network, height):
		return parse(raw, rawReturn, true, &verifregv13.UniversalReceiverParams{}, &verifregv13.AllocationsResponse{})
	case tools.V21.IsSupported(network, height):
		return parse(raw, rawReturn, true, &verifregv12.UniversalReceiverParams{}, &verifregv12.AllocationsResponse{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parse(raw, rawReturn, true, &verifregv11.UniversalReceiverParams{}, &verifregv11.AllocationsResponse{})
	case tools.V18.IsSupported(network, height):
		return parse(raw, rawReturn, true, &verifregv10.UniversalReceiverParams{}, &verifregv10.AllocationsResponse{})
	case tools.V17.IsSupported(network, height):
		return parse(raw, rawReturn, true, &verifregv9.UniversalReceiverParams{}, &verifregv9.AllocationsResponse{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V16)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}
