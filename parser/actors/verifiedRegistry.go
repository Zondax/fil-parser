package actors

import (
	"bytes"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/builtin/v11/verifreg"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/zondax/fil-parser/parser"
)

func ParseVerifiedRegistry(txType string, msg *parser.LotusMessage, msgRct *filTypes.MessageReceipt) (map[string]interface{}, error) {
	switch txType {
	case parser.MethodSend:
		return parseSend(msg), nil
	case parser.MethodConstructor:
		return parseConstructor(msg.Params)
	case parser.MethodAddVerifier:
		return addVerifier(msg.Params)
	case parser.MethodRemoveVerifier: // TODO: not tested
		return removeVerifier(msg.Params)
	case parser.MethodAddVerifiedClient, parser.MethodAddVerifiedClientExported:
		return addVerifiedClient(msg.Params)
	case parser.MethodUseBytes:
		return useBytes(msg.Params)
	case parser.MethodRestoreBytes:
		return restoreBytes(msg.Params)
	case parser.MethodRemoveVerifiedClientDataCap: // TODO: not tested
		return removeVerifiedClientDataCap(msg.Params)
	case parser.MethodRemoveExpiredAllocations, parser.MethodRemoveExpiredAllocationsExported:
		return removeExpiredAllocations(msg.Params, msgRct.Return)
	case parser.MethodVerifiedDeprecated1: // UseBytes
		return deprecated1(msg.Params)
	case parser.MethodVerifiedDeprecated2: // RestoreBytes
		return deprecated2(msg.Params)
	case parser.MethodClaimAllocations:
		return claimAllocations(msg.Params, msgRct.Return)
	case parser.MethodGetClaims, parser.MethodGetClaimsExported: // TODO: not tested
		return getClaims(msg.Params, msgRct.Return)
	case parser.MethodExtendClaimTerms, parser.MethodExtendClaimTermsExported: // TODO: not tested
		return extendClaimTerms(msg.Params, msgRct.Return)
	case parser.MethodRemoveExpiredClaims, parser.MethodRemoveExpiredClaimsExported:
		return removeExpiredClaims(msg.Params, msgRct.Return)
	case parser.MethodUniversalReceiverHook:
		return verifregUniversalReceiverHook(msg.Params, msgRct.Return)
	}
	return map[string]interface{}{}, parser.ErrUnknownMethod
}

func addVerifier(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params verifreg.AddVerifierParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func removeVerifier(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params address.Address
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params.String()
	return metadata, nil
}

func addVerifiedClient(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params verifreg.AddVerifiedClientParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func useBytes(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params verifreg.UseBytesParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func restoreBytes(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params verifreg.RestoreBytesParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

// TODO: untested
func removeVerifiedClientDataCap(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var datacap verifreg.DataCap
	err := datacap.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = datacap
	return metadata, nil
}

func removeExpiredAllocations(raw, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params verifreg.RemoveExpiredAllocationsParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var expiredReturn verifreg.RemoveExpiredAllocationsReturn
	err = expiredReturn.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = expiredReturn
	return metadata, nil
}

func deprecated1(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params verifreg.RestoreBytesParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func deprecated2(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params verifreg.UseBytesParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func claimAllocations(raw, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params verifreg.ClaimAllocationsParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params

	reader = bytes.NewReader(rawReturn)
	var expiredReturn verifreg.ClaimAllocationsReturn
	err = expiredReturn.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = expiredReturn
	return metadata, nil
}

func getClaims(raw, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params verifreg.GetClaimsParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var expiredReturn verifreg.GetClaimsReturn
	err = expiredReturn.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = expiredReturn
	return metadata, nil
}

func extendClaimTerms(raw, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params verifreg.ExtendClaimTermsParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var expiredReturn verifreg.ExtendClaimTermsReturn
	err = expiredReturn.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = expiredReturn
	return metadata, nil
}

func removeExpiredClaims(raw, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params verifreg.RemoveExpiredClaimsParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var expiredReturn verifreg.RemoveExpiredClaimsReturn
	err = expiredReturn.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = expiredReturn
	return metadata, nil
}

func verifregUniversalReceiverHook(raw, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	var params verifreg.UniversalReceiverParams
	reader := bytes.NewReader(raw)
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params

	reader = bytes.NewReader(rawReturn)
	var r verifreg.AllocationsResponse
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r
	return metadata, nil
}
