package parser

import (
	"bytes"
	"github.com/filecoin-project/go-state-types/builtin/v10/verifreg"
	filTypes "github.com/filecoin-project/lotus/chain/types"
)

func (p *Parser) parseVerifiedRegistry(txType string, msg *filTypes.Message, msgRct *filTypes.MessageReceipt) (map[string]interface{}, error) {
	switch txType {
	case MethodSend:
		return p.parseSend(msg), nil
	case MethodConstructor:
	case MethodAddVerifier:
		return p.addVerifier(msg.Params)
	case MethodRemoveVerifier:
	case MethodAddVerifiedClient:
		return p.addVerifiedClient(msg.Params)
	case MethodUseBytes:
		return p.useBytes(msg.Params)
	case MethodRestoreBytes:
		return p.restoreBytes(msg.Params)
	case MethodRemoveVerifiedClientDataCap:
		// TODO: untested
		return p.removeVerifiedClientDataCap(msg.Params)
	case MethodRemoveExpiredAllocations:
		return p.removeExpiredAllocations(msg.Params, msgRct.Return)
	case MethodVerifiedDeprecated1: // UseBytes
		return p.deprecated1(msg.Params)
	case MethodVerifiedDeprecated2: // RestoreBytes
		return p.deprecated2(msg.Params)
	case UnknownStr:
		return p.unknownMetadata(msg.Params, msgRct.Return)
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
