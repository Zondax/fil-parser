package verifiedregistry

import (
	"github.com/filecoin-project/go-address"
	verifregv11 "github.com/filecoin-project/go-state-types/builtin/v11/verifreg"
	verifregv8 "github.com/filecoin-project/go-state-types/builtin/v8/verifreg"
)

func AddVerifier(height int64, raw []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		return parse[*verifregv8.AddVerifierParams, *verifregv8.AddVerifierParams](raw, nil, false)
	case 11:
		return parse[*verifregv11.AddVerifierParams, *verifregv11.AddVerifierParams](raw, nil, false)
	}
	return nil, nil
}

func RemoveVerifier(height int64, raw []byte) (map[string]interface{}, error) {
	return parse[*address.Address, *address.Address](raw, nil, false)
}

func AddVerifiedClient(height int64, raw []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		return parse[*verifregv8.AddVerifiedClientParams, *verifregv8.AddVerifiedClientParams](raw, nil, false)
	case 11:
		return parse[*verifregv11.AddVerifiedClientParams, *verifregv11.AddVerifiedClientParams](raw, nil, false)
	}
	return nil, nil
}

func UseBytes(height int64, raw []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		return parse[*verifregv8.UseBytesParams, *verifregv8.UseBytesParams](raw, nil, false)
	case 11:
		return parse[*verifregv11.UseBytesParams, *verifregv11.UseBytesParams](raw, nil, false)
	}
	return nil, nil
}

func RestoreBytes(height int64, raw []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		return parse[*verifregv8.RestoreBytesParams, *verifregv8.RestoreBytesParams](raw, nil, false)
	case 11:
		return parse[*verifregv11.RestoreBytesParams, *verifregv11.RestoreBytesParams](raw, nil, false)
	}
	return nil, nil
}

func RemoveVerifiedClientDataCap(height int64, raw []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		return parse[*verifregv8.DataCap, *verifregv8.DataCap](raw, nil, false)
	case 11:
		return parse[*verifregv11.DataCap, *verifregv11.DataCap](raw, nil, false)
	}
	return nil, nil
}

func RemoveExpiredAllocations(height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		return parse[*verifregv8.RemoveExpiredAllocationsParams, *verifregv8.RemoveExpiredAllocationsReturn](raw, rawReturn, true)
	case 11:
		return parse[*verifregv11.RemoveExpiredAllocationsParams, *verifregv11.RemoveExpiredAllocationsReturn](raw, rawReturn, true)
	}
	return nil, nil
}

func Deprecated1(height int64, raw []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		return parse[*verifregv8.RestoreBytesParams, *verifregv8.RestoreBytesParams](raw, nil, false)
	case 11:
		return parse[*verifregv11.RestoreBytesParams, *verifregv11.RestoreBytesParams](raw, nil, false)
	}
	return nil, nil
}

func Deprecated2(height int64, raw []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		return parse[*verifregv8.UseBytesParams, *verifregv8.UseBytesParams](raw, nil, false)
	case 11:
		return parse[*verifregv11.UseBytesParams, *verifregv11.UseBytesParams](raw, nil, false)
	}
	return nil, nil
}

func ClaimAllocations(height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		return parse[*verifregv8.ClaimAllocationsParams, *verifregv8.ClaimAllocationsReturn](raw, rawReturn, true)
	case 11:
		return parse[*verifregv11.ClaimAllocationsParams, *verifregv11.ClaimAllocationsReturn](raw, rawReturn, true)
	}
	return nil, nil
}

func GetClaims(height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		return parse[*verifregv8.GetClaimsParams, *verifregv8.GetClaimsReturn](raw, rawReturn, true)
	case 11:
		return parse[*verifregv11.GetClaimsParams, *verifregv11.GetClaimsReturn](raw, rawReturn, true)
	}
	return nil, nil
}

func ExtendClaimTerms(height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		return parse[*verifregv8.ExtendClaimTermsParams, *verifregv8.ExtendClaimTermsReturn](raw, rawReturn, true)
	case 11:
		return parse[*verifregv11.ExtendClaimTermsParams, *verifregv11.ExtendClaimTermsReturn](raw, rawReturn, true)
	}
	return nil, nil
}

func RemoveExpiredClaims(height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		return parse[*verifregv8.RemoveExpiredClaimsParams, *verifregv8.RemoveExpiredClaimsReturn](raw, rawReturn, true)
	case 11:
		return parse[*verifregv11.RemoveExpiredClaimsParams, *verifregv11.RemoveExpiredClaimsReturn](raw, rawReturn, true)
	}
	return nil, nil
}

func VerifregUniversalReceiverHook(height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		return parse[*verifregv8.UniversalReceiverParams, *verifregv8.AllocationsResponse](raw, rawReturn, true)
	case 11:
		return parse[*verifregv11.UniversalReceiverParams, *verifregv11.AllocationsResponse](raw, rawReturn, true)
	}
	return nil, nil
}
