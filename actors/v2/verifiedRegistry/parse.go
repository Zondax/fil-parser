package verifiedregistry

import (
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
)

func (p *VerifiedRegistry) Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, _ cid.Cid) (map[string]interface{}, *types.AddressInfo, error) {
	switch txType {
	case parser.MethodSend:
		// return p.parseSend(msg), nil
	case parser.MethodConstructor:
		// return p.Constructor(network, height, msg.Params)
	case parser.MethodAddVerifier:
		resp, err := p.AddVerifier(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodRemoveVerifier: // TODO: not tested
		resp, err := p.RemoveVerifier(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodAddVerifiedClient, parser.MethodAddVerifiedClientExported:
		resp, err := p.AddVerifiedClientExported(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodUseBytes:
		resp, err := p.UseBytes(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodRestoreBytes:
		resp, err := p.RestoreBytes(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodRemoveVerifiedClientDataCap: // TODO: not tested
		resp, err := p.RemoveVerifiedClientDataCap(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodRemoveExpiredAllocations, parser.MethodRemoveExpiredAllocationsExported:
		resp, err := p.RemoveExpiredAllocationsExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodVerifiedDeprecated1: // UseBytes
		resp, err := p.Deprecated1(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodVerifiedDeprecated2: // RestoreBytes
		resp, err := p.Deprecated2(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodClaimAllocations:
		resp, err := p.ClaimAllocations(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodGetClaims, parser.MethodGetClaimsExported: // TODO: not tested
		resp, err := p.GetClaimsExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodExtendClaimTerms, parser.MethodExtendClaimTermsExported: // TODO: not tested
		resp, err := p.ExtendClaimTermsExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodRemoveExpiredClaims, parser.MethodRemoveExpiredClaimsExported:
		resp, err := p.RemoveExpiredClaimsExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodUniversalReceiverHook:
		resp, err := p.UniversalReceiverHook(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	}
	return map[string]interface{}{}, nil, parser.ErrUnknownMethod
}

func (p *VerifiedRegistry) TransactionTypes() []string {
	return []string{
		parser.MethodSend,
		parser.MethodConstructor,
		parser.MethodAddVerifier,
		parser.MethodRemoveVerifier,
		parser.MethodAddVerifiedClient,
		parser.MethodAddVerifiedClientExported,
		parser.MethodUseBytes,
		parser.MethodRestoreBytes,
		parser.MethodRemoveVerifiedClientDataCap,
		parser.MethodRemoveExpiredAllocations,
		parser.MethodRemoveExpiredAllocationsExported,
		parser.MethodVerifiedDeprecated1,
		parser.MethodVerifiedDeprecated2,
		parser.MethodClaimAllocations,
		parser.MethodGetClaims,
		parser.MethodGetClaimsExported,
		parser.MethodExtendClaimTerms,
		parser.MethodExtendClaimTermsExported,
		parser.MethodRemoveExpiredClaims,
		parser.MethodRemoveExpiredClaimsExported,
		parser.MethodUniversalReceiverHook,
	}
}
