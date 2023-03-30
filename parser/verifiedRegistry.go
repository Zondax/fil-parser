package parser

import (
	"bytes"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/builtin/v11/verifreg"
	filTypes "github.com/filecoin-project/lotus/chain/types"
)

func (p *Parser) parseVerifiedRegistry(txType string, msg *filTypes.Message, msgRct *filTypes.MessageReceipt) (map[string]interface{}, error) {
	switch txType {
	case MethodSend:
		return p.parseSend(msg), nil
	case MethodConstructor:
		return p.parseConstructor(msg.Params)
	case MethodAddVerifier:
		return p.addVerifier(msg.Params)
	case MethodRemoveVerifier: // TODO: not tested
		return p.removeVerifier(msg.Params)
	case MethodAddVerifiedClient, MethodAddVerifiedClientExported:
		return p.addVerifiedClient(msg.Params)
	case MethodUseBytes:
		return p.useBytes(msg.Params)
	case MethodRestoreBytes:
		return p.restoreBytes(msg.Params)
	case MethodRemoveVerifiedClientDataCap: // TODO: not tested
		return p.removeVerifiedClientDataCap(msg.Params)
	case MethodRemoveExpiredAllocations, MethodRemoveExpiredAllocationsExported:
		return p.removeExpiredAllocations(msg.Params, msgRct.Return)
	case MethodVerifiedDeprecated1: // UseBytes
		return p.deprecated1(msg.Params)
	case MethodVerifiedDeprecated2: // RestoreBytes
		return p.deprecated2(msg.Params)
	case MethodClaimAllocations:
		return p.claimAllocations(msg.Params, msgRct.Return)
	case MethodGetClaims, MethodGetClaimsExported: // TODO: not tested
		return p.getClaims(msg.Params, msgRct.Return)
	case MethodExtendClaimTerms, MethodExtendClaimTermsExported: // TODO: not tested
		return p.extendClaimTerms(msg.Params, msgRct.Return)
	case MethodRemoveExpiredClaims, MethodRemoveExpiredClaimsExported:
		return p.removeExpiredClaims(msg.Params, msgRct.Return)
	case MethodUniversalReceiverHook:
		return p.verifregUniversalReceiverHook(msg.Params, msgRct.Return)
	}
	return map[string]interface{}{}, errUnknownMethod
}

func (p *Parser) addVerifier(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params verifreg.AddVerifierParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) removeVerifier(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params address.Address
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params.String()
	return metadata, nil
}

func (p *Parser) addVerifiedClient(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params verifreg.AddVerifiedClientParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) useBytes(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params verifreg.UseBytesParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) restoreBytes(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params verifreg.RestoreBytesParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

// TODO: untested
func (p *Parser) removeVerifiedClientDataCap(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var datacap verifreg.DataCap
	err := datacap.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = datacap
	return metadata, nil
}

func (p *Parser) removeExpiredAllocations(raw, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params verifreg.RemoveExpiredAllocationsParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var expiredReturn verifreg.RemoveExpiredAllocationsReturn
	err = expiredReturn.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = expiredReturn
	return metadata, nil
}

func (p *Parser) deprecated1(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params verifreg.RestoreBytesParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) deprecated2(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params verifreg.UseBytesParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) claimAllocations(raw, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params verifreg.ClaimAllocationsParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params

	reader = bytes.NewReader(rawReturn)
	var expiredReturn verifreg.ClaimAllocationsReturn
	err = expiredReturn.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = expiredReturn
	return metadata, nil
}

func (p *Parser) getClaims(raw, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params verifreg.GetClaimsParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var expiredReturn verifreg.GetClaimsReturn
	err = expiredReturn.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = expiredReturn
	return metadata, nil
}

func (p *Parser) extendClaimTerms(raw, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params verifreg.ExtendClaimTermsParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var expiredReturn verifreg.ExtendClaimTermsReturn
	err = expiredReturn.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = expiredReturn
	return metadata, nil
}

func (p *Parser) removeExpiredClaims(raw, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params verifreg.RemoveExpiredClaimsParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var expiredReturn verifreg.RemoveExpiredClaimsReturn
	err = expiredReturn.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = expiredReturn
	return metadata, nil
}

func (p *Parser) verifregUniversalReceiverHook(raw, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	var params verifreg.UniversalReceiverParams
	reader := bytes.NewReader(raw)
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params

	reader = bytes.NewReader(rawReturn)
	var r verifreg.AllocationsResponse
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = r
	return metadata, nil
}
