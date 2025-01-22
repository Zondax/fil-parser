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

	"github.com/zondax/fil-parser/tools"
)

func AddVerifier(network string, height int64, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(network, height):
		return parse[*verifregv8.AddVerifierParams, *verifregv8.AddVerifierParams](raw, nil, false)
	case tools.V9.IsSupported(network, height):
		return parse[*verifregv9.AddVerifierParams, *verifregv9.AddVerifierParams](raw, nil, false)
	case tools.V10.IsSupported(network, height):
		return parse[*verifregv10.AddVerifierParams, *verifregv10.AddVerifierParams](raw, nil, false)
	case tools.V11.IsSupported(network, height):
		return parse[*verifregv11.AddVerifierParams, *verifregv11.AddVerifierParams](raw, nil, false)
	case tools.V12.IsSupported(network, height):
		return parse[*verifregv12.AddVerifierParams, *verifregv12.AddVerifierParams](raw, nil, false)
	case tools.V13.IsSupported(network, height):
		return parse[*verifregv13.AddVerifierParams, *verifregv13.AddVerifierParams](raw, nil, false)
	case tools.V14.IsSupported(network, height):
		return parse[*verifregv14.AddVerifierParams, *verifregv14.AddVerifierParams](raw, nil, false)
	case tools.V15.IsSupported(network, height):
		return parse[*verifregv15.AddVerifierParams, *verifregv15.AddVerifierParams](raw, nil, false)
	}
	return nil, nil
}

func RemoveVerifier(network string, height int64, raw []byte) (map[string]interface{}, error) {
	return parse[*address.Address, *address.Address](raw, nil, false)
}

func AddVerifiedClient(network string, height int64, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(network, height):
		return parse[*verifregv8.AddVerifiedClientParams, *verifregv8.AddVerifiedClientParams](raw, nil, false)
	case tools.V9.IsSupported(network, height):
		return parse[*verifregv9.AddVerifiedClientParams, *verifregv9.AddVerifiedClientParams](raw, nil, false)
	case tools.V10.IsSupported(network, height):
		return parse[*verifregv10.AddVerifiedClientParams, *verifregv10.AddVerifiedClientParams](raw, nil, false)
	case tools.V11.IsSupported(network, height):
		return parse[*verifregv11.AddVerifiedClientParams, *verifregv11.AddVerifiedClientParams](raw, nil, false)
	case tools.V12.IsSupported(network, height):
		return parse[*verifregv12.AddVerifiedClientParams, *verifregv12.AddVerifiedClientParams](raw, nil, false)
	case tools.V13.IsSupported(network, height):
		return parse[*verifregv13.AddVerifiedClientParams, *verifregv13.AddVerifiedClientParams](raw, nil, false)
	case tools.V14.IsSupported(network, height):
		return parse[*verifregv14.AddVerifiedClientParams, *verifregv14.AddVerifiedClientParams](raw, nil, false)
	case tools.V15.IsSupported(network, height):
		return parse[*verifregv15.AddVerifiedClientParams, *verifregv15.AddVerifiedClientParams](raw, nil, false)
	}
	return nil, nil
}

func UseBytes(network string, height int64, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(network, height):
		return parse[*verifregv8.UseBytesParams, *verifregv8.UseBytesParams](raw, nil, false)
	case tools.V9.IsSupported(network, height):
		return parse[*verifregv9.UseBytesParams, *verifregv9.UseBytesParams](raw, nil, false)
	case tools.V10.IsSupported(network, height):
		return parse[*verifregv10.UseBytesParams, *verifregv10.UseBytesParams](raw, nil, false)
	case tools.V11.IsSupported(network, height):
		return parse[*verifregv11.UseBytesParams, *verifregv11.UseBytesParams](raw, nil, false)
	case tools.V12.IsSupported(network, height):
		return parse[*verifregv12.UseBytesParams, *verifregv12.UseBytesParams](raw, nil, false)
	case tools.V13.IsSupported(network, height):
		return parse[*verifregv13.UseBytesParams, *verifregv13.UseBytesParams](raw, nil, false)
	case tools.V14.IsSupported(network, height):
		return parse[*verifregv14.UseBytesParams, *verifregv14.UseBytesParams](raw, nil, false)
	case tools.V15.IsSupported(network, height):
		return parse[*verifregv15.UseBytesParams, *verifregv15.UseBytesParams](raw, nil, false)
	}
	return nil, nil
}

func RestoreBytes(network string, height int64, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(network, height):
		return parse[*verifregv8.RestoreBytesParams, *verifregv8.RestoreBytesParams](raw, nil, false)
	case tools.V9.IsSupported(network, height):
		return parse[*verifregv9.RestoreBytesParams, *verifregv9.RestoreBytesParams](raw, nil, false)
	case tools.V10.IsSupported(network, height):
		return parse[*verifregv10.RestoreBytesParams, *verifregv10.RestoreBytesParams](raw, nil, false)
	case tools.V11.IsSupported(network, height):
		return parse[*verifregv11.RestoreBytesParams, *verifregv11.RestoreBytesParams](raw, nil, false)
	case tools.V12.IsSupported(network, height):
		return parse[*verifregv12.RestoreBytesParams, *verifregv12.RestoreBytesParams](raw, nil, false)
	case tools.V13.IsSupported(network, height):
		return parse[*verifregv13.RestoreBytesParams, *verifregv13.RestoreBytesParams](raw, nil, false)
	case tools.V14.IsSupported(network, height):
		return parse[*verifregv14.RestoreBytesParams, *verifregv14.RestoreBytesParams](raw, nil, false)
	case tools.V15.IsSupported(network, height):
		return parse[*verifregv15.RestoreBytesParams, *verifregv15.RestoreBytesParams](raw, nil, false)
	}
	return nil, nil
}

func RemoveVerifiedClientDataCap(network string, height int64, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(network, height):
		return parse[*verifregv8.DataCap, *verifregv8.DataCap](raw, nil, false)
	case tools.V9.IsSupported(network, height):
		return parse[*verifregv9.DataCap, *verifregv9.DataCap](raw, nil, false)
	case tools.V10.IsSupported(network, height):
		return parse[*verifregv10.DataCap, *verifregv10.DataCap](raw, nil, false)
	case tools.V11.IsSupported(network, height):
		return parse[*verifregv11.DataCap, *verifregv11.DataCap](raw, nil, false)
	case tools.V12.IsSupported(network, height):
		return parse[*verifregv12.DataCap, *verifregv12.DataCap](raw, nil, false)
	case tools.V13.IsSupported(network, height):
		return parse[*verifregv13.DataCap, *verifregv13.DataCap](raw, nil, false)
	case tools.V14.IsSupported(network, height):
		return parse[*verifregv14.DataCap, *verifregv14.DataCap](raw, nil, false)
	case tools.V15.IsSupported(network, height):
		return parse[*verifregv15.DataCap, *verifregv15.DataCap](raw, nil, false)
	}
	return nil, nil
}

func RemoveExpiredAllocations(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(network, height):
		return nil, fmt.Errorf("not supported")
	case tools.V9.IsSupported(network, height):
		return parse[*verifregv9.RemoveExpiredAllocationsParams, *verifregv9.RemoveExpiredAllocationsReturn](raw, rawReturn, true)
	case tools.V10.IsSupported(network, height):
		return parse[*verifregv10.RemoveExpiredAllocationsParams, *verifregv10.RemoveExpiredAllocationsReturn](raw, rawReturn, true)
	case tools.V11.IsSupported(network, height):
		return parse[*verifregv11.RemoveExpiredAllocationsParams, *verifregv11.RemoveExpiredAllocationsReturn](raw, rawReturn, true)
	case tools.V12.IsSupported(network, height):
		return parse[*verifregv12.RemoveExpiredAllocationsParams, *verifregv12.RemoveExpiredAllocationsReturn](raw, rawReturn, true)
	case tools.V13.IsSupported(network, height):
		return parse[*verifregv13.RemoveExpiredAllocationsParams, *verifregv13.RemoveExpiredAllocationsReturn](raw, rawReturn, true)
	case tools.V14.IsSupported(network, height):
		return parse[*verifregv14.RemoveExpiredAllocationsParams, *verifregv14.RemoveExpiredAllocationsReturn](raw, rawReturn, true)
	case tools.V15.IsSupported(network, height):
		return parse[*verifregv15.RemoveExpiredAllocationsParams, *verifregv15.RemoveExpiredAllocationsReturn](raw, rawReturn, true)
	}
	return nil, nil
}

func Deprecated1(network string, height int64, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(network, height):
		return parse[*verifregv8.RestoreBytesParams, *verifregv8.RestoreBytesParams](raw, nil, false)
	case tools.V9.IsSupported(network, height):
		return parse[*verifregv9.RestoreBytesParams, *verifregv9.RestoreBytesParams](raw, nil, false)
	case tools.V10.IsSupported(network, height):
		return parse[*verifregv10.RestoreBytesParams, *verifregv10.RestoreBytesParams](raw, nil, false)
	case tools.V11.IsSupported(network, height):
		return parse[*verifregv11.RestoreBytesParams, *verifregv11.RestoreBytesParams](raw, nil, false)
	case tools.V12.IsSupported(network, height):
		return parse[*verifregv12.RestoreBytesParams, *verifregv12.RestoreBytesParams](raw, nil, false)
	case tools.V13.IsSupported(network, height):
		return parse[*verifregv13.RestoreBytesParams, *verifregv13.RestoreBytesParams](raw, nil, false)
	case tools.V14.IsSupported(network, height):
		return parse[*verifregv14.RestoreBytesParams, *verifregv14.RestoreBytesParams](raw, nil, false)
	case tools.V15.IsSupported(network, height):
		return parse[*verifregv15.RestoreBytesParams, *verifregv15.RestoreBytesParams](raw, nil, false)
	}
	return nil, nil
}

func Deprecated2(network string, height int64, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(network, height):
		return parse[*verifregv8.UseBytesParams, *verifregv8.UseBytesParams](raw, nil, false)
	case tools.V9.IsSupported(network, height):
		return parse[*verifregv9.UseBytesParams, *verifregv9.UseBytesParams](raw, nil, false)
	case tools.V10.IsSupported(network, height):
		return parse[*verifregv10.UseBytesParams, *verifregv10.UseBytesParams](raw, nil, false)
	case tools.V11.IsSupported(network, height):
		return parse[*verifregv11.UseBytesParams, *verifregv11.UseBytesParams](raw, nil, false)
	case tools.V12.IsSupported(network, height):
		return parse[*verifregv12.UseBytesParams, *verifregv12.UseBytesParams](raw, nil, false)
	case tools.V13.IsSupported(network, height):
		return parse[*verifregv13.UseBytesParams, *verifregv13.UseBytesParams](raw, nil, false)
	case tools.V14.IsSupported(network, height):
		return parse[*verifregv14.UseBytesParams, *verifregv14.UseBytesParams](raw, nil, false)
	case tools.V15.IsSupported(network, height):
		return parse[*verifregv15.UseBytesParams, *verifregv15.UseBytesParams](raw, nil, false)
	}
	return nil, nil
}

func ClaimAllocations(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(network, height):
		return nil, fmt.Errorf("not supported")
	case tools.V9.IsSupported(network, height):
		return parse[*verifregv9.ClaimAllocationsParams, *verifregv9.ClaimAllocationsReturn](raw, rawReturn, true)
	case tools.V10.IsSupported(network, height):
		return parse[*verifregv10.ClaimAllocationsParams, *verifregv10.ClaimAllocationsReturn](raw, rawReturn, true)
	case tools.V11.IsSupported(network, height):
		return parse[*verifregv11.ClaimAllocationsParams, *verifregv11.ClaimAllocationsReturn](raw, rawReturn, true)
	case tools.V12.IsSupported(network, height):
		return parse[*verifregv12.ClaimAllocationsParams, *verifregv12.ClaimAllocationsReturn](raw, rawReturn, true)
	case tools.V13.IsSupported(network, height):
		return parse[*verifregv13.ClaimAllocationsParams, *verifregv13.ClaimAllocationsReturn](raw, rawReturn, true)
	case tools.V14.IsSupported(network, height):
		return parse[*verifregv14.ClaimAllocationsParams, *verifregv14.ClaimAllocationsReturn](raw, rawReturn, true)
	case tools.V15.IsSupported(network, height):
		return parse[*verifregv15.ClaimAllocationsParams, *verifregv15.ClaimAllocationsReturn](raw, rawReturn, true)
	}
	return nil, nil
}

func GetClaims(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(network, height):
		return nil, fmt.Errorf("not supported")
	case tools.V9.IsSupported(network, height):
		return parse[*verifregv9.GetClaimsParams, *verifregv9.GetClaimsReturn](raw, rawReturn, true)
	case tools.V10.IsSupported(network, height):
		return parse[*verifregv10.GetClaimsParams, *verifregv10.GetClaimsReturn](raw, rawReturn, true)
	case tools.V11.IsSupported(network, height):
		return parse[*verifregv11.GetClaimsParams, *verifregv11.GetClaimsReturn](raw, rawReturn, true)
	case tools.V12.IsSupported(network, height):
		return parse[*verifregv12.GetClaimsParams, *verifregv12.GetClaimsReturn](raw, rawReturn, true)
	case tools.V13.IsSupported(network, height):
		return parse[*verifregv13.GetClaimsParams, *verifregv13.GetClaimsReturn](raw, rawReturn, true)
	case tools.V14.IsSupported(network, height):
		return parse[*verifregv14.GetClaimsParams, *verifregv14.GetClaimsReturn](raw, rawReturn, true)
	case tools.V15.IsSupported(network, height):
		return parse[*verifregv15.GetClaimsParams, *verifregv15.GetClaimsReturn](raw, rawReturn, true)
	}
	return nil, nil
}

func ExtendClaimTerms(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(network, height):
		return nil, fmt.Errorf("not supported")
	case tools.V9.IsSupported(network, height):
		return parse[*verifregv9.ExtendClaimTermsParams, *verifregv9.ExtendClaimTermsReturn](raw, rawReturn, true)
	case tools.V10.IsSupported(network, height):
		return parse[*verifregv10.ExtendClaimTermsParams, *verifregv10.ExtendClaimTermsReturn](raw, rawReturn, true)
	case tools.V11.IsSupported(network, height):
		return parse[*verifregv11.ExtendClaimTermsParams, *verifregv11.ExtendClaimTermsReturn](raw, rawReturn, true)
	case tools.V12.IsSupported(network, height):
		return parse[*verifregv12.ExtendClaimTermsParams, *verifregv12.ExtendClaimTermsReturn](raw, rawReturn, true)
	case tools.V13.IsSupported(network, height):
		return parse[*verifregv13.ExtendClaimTermsParams, *verifregv13.ExtendClaimTermsReturn](raw, rawReturn, true)
	case tools.V14.IsSupported(network, height):
		return parse[*verifregv14.ExtendClaimTermsParams, *verifregv14.ExtendClaimTermsReturn](raw, rawReturn, true)
	case tools.V15.IsSupported(network, height):
		return parse[*verifregv15.ExtendClaimTermsParams, *verifregv15.ExtendClaimTermsReturn](raw, rawReturn, true)
	}
	return nil, nil
}

func RemoveExpiredClaims(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(network, height):
		return nil, fmt.Errorf("not supported")
	case tools.V9.IsSupported(network, height):
		return parse[*verifregv9.RemoveExpiredClaimsParams, *verifregv9.RemoveExpiredClaimsReturn](raw, rawReturn, true)
	case tools.V10.IsSupported(network, height):
		return parse[*verifregv10.RemoveExpiredClaimsParams, *verifregv10.RemoveExpiredClaimsReturn](raw, rawReturn, true)
	case tools.V11.IsSupported(network, height):
		return parse[*verifregv11.RemoveExpiredClaimsParams, *verifregv11.RemoveExpiredClaimsReturn](raw, rawReturn, true)
	case tools.V12.IsSupported(network, height):
		return parse[*verifregv12.RemoveExpiredClaimsParams, *verifregv12.RemoveExpiredClaimsReturn](raw, rawReturn, true)
	case tools.V13.IsSupported(network, height):
		return parse[*verifregv13.RemoveExpiredClaimsParams, *verifregv13.RemoveExpiredClaimsReturn](raw, rawReturn, true)
	case tools.V14.IsSupported(network, height):
		return parse[*verifregv14.RemoveExpiredClaimsParams, *verifregv14.RemoveExpiredClaimsReturn](raw, rawReturn, true)
	case tools.V15.IsSupported(network, height):
		return parse[*verifregv15.RemoveExpiredClaimsParams, *verifregv15.RemoveExpiredClaimsReturn](raw, rawReturn, true)
	}
	return nil, nil
}

func VerifregUniversalReceiverHook(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(network, height):
		return nil, fmt.Errorf("not supported")
	case tools.V9.IsSupported(network, height):
		return parse[*verifregv9.UniversalReceiverParams, *verifregv9.AllocationsResponse](raw, rawReturn, true)
	case tools.V10.IsSupported(network, height):
		return parse[*verifregv10.UniversalReceiverParams, *verifregv10.AllocationsResponse](raw, rawReturn, true)
	case tools.V11.IsSupported(network, height):
		return parse[*verifregv11.UniversalReceiverParams, *verifregv11.AllocationsResponse](raw, rawReturn, true)
	case tools.V12.IsSupported(network, height):
		return parse[*verifregv12.UniversalReceiverParams, *verifregv12.AllocationsResponse](raw, rawReturn, true)
	case tools.V13.IsSupported(network, height):
		return parse[*verifregv13.UniversalReceiverParams, *verifregv13.AllocationsResponse](raw, rawReturn, true)
	case tools.V14.IsSupported(network, height):
		return parse[*verifregv14.UniversalReceiverParams, *verifregv14.AllocationsResponse](raw, rawReturn, true)
	case tools.V15.IsSupported(network, height):
		return parse[*verifregv15.UniversalReceiverParams, *verifregv15.AllocationsResponse](raw, rawReturn, true)
	}
	return nil, nil
}
