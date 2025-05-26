package actors

import (
	"bytes"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/builtin/v11/verifreg"
	"github.com/zondax/fil-parser/parser"
)

func (p *ActorParser) ParseVerifiedRegistry(txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt) (map[string]interface{}, error) {
	switch txType {
	case parser.MethodSend:
		return p.parseSend(msg), nil
	case parser.MethodConstructor:
		return p.parseConstructor(msg.Params)
	case parser.MethodAddVerifier:
		return p.addVerifier(msg.Params)
	case parser.MethodRemoveVerifier: // TODO: not tested
		return p.removeVerifier(msg.Params)
	case parser.MethodAddVerifiedClient, parser.MethodAddVerifiedClientExported:
		return p.addVerifiedClient(msg.Params)
	case parser.MethodUseBytes:
		return p.useBytes(msg.Params)
	case parser.MethodRestoreBytes:
		return p.restoreBytes(msg.Params)
	case parser.MethodRemoveVerifiedClientDataCap: // TODO: not tested
		return p.removeVerifiedClientDataCap(msg.Params)
	case parser.MethodRemoveExpiredAllocations, parser.MethodRemoveExpiredAllocationsExported:
		return p.removeExpiredAllocations(msg.Params, msgRct.Return)
	case parser.MethodVerifiedDeprecated1: // UseBytes
		return p.deprecated1(msg.Params)
	case parser.MethodVerifiedDeprecated2: // RestoreBytes
		return p.deprecated2(msg.Params)
	case parser.MethodClaimAllocations:
		return p.claimAllocations(msg.Params, msgRct.Return)
	case parser.MethodGetClaims, parser.MethodGetClaimsExported: // TODO: not tested
		return p.getClaims(msg.Params, msgRct.Return)
	case parser.MethodExtendClaimTerms, parser.MethodExtendClaimTermsExported: // TODO: not tested
		return p.extendClaimTerms(msg.Params, msgRct.Return)
	case parser.MethodRemoveExpiredClaims, parser.MethodRemoveExpiredClaimsExported:
		return p.removeExpiredClaims(msg.Params, msgRct.Return)
	case parser.MethodUniversalReceiverHook:
		return p.verifregUniversalReceiverHook(msg.Params, msgRct.Return)
	}
	return map[string]interface{}{}, parser.ErrUnknownMethod
}

func (p *ActorParser) addVerifier(raw []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) removeVerifier(raw []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) addVerifiedClient(raw []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) useBytes(raw []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) restoreBytes(raw []byte) (map[string]interface{}, error) {
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
func (p *ActorParser) removeVerifiedClientDataCap(raw []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) removeExpiredAllocations(raw, rawReturn []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) deprecated1(raw []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) deprecated2(raw []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) claimAllocations(raw, rawReturn []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) getClaims(raw, rawReturn []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) extendClaimTerms(raw, rawReturn []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) removeExpiredClaims(raw, rawReturn []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) verifregUniversalReceiverHook(raw, rawReturn []byte) (map[string]interface{}, error) {
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
