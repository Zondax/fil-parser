package verifiedregistry

import (
	"github.com/zondax/fil-parser/parser"
)

func (p *VerifiedRegistry) Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt) (map[string]interface{}, error) {
	switch txType {
	case parser.MethodSend:
		// return p.parseSend(msg), nil
	case parser.MethodConstructor:
		// return p.Constructor(network, height, msg.Params)
	case parser.MethodAddVerifier:
		return p.AddVerifier(network, height, msg.Params)
	case parser.MethodRemoveVerifier: // TODO: not tested
		return p.RemoveVerifier(network, height, msg.Params)
	case parser.MethodAddVerifiedClient, parser.MethodAddVerifiedClientExported:
		return p.AddVerifiedClientExported(network, height, msg.Params)
	case parser.MethodUseBytes:
		return p.UseBytes(network, height, msg.Params)
	case parser.MethodRestoreBytes:
		return p.RestoreBytes(network, height, msg.Params)
	case parser.MethodRemoveVerifiedClientDataCap: // TODO: not tested
		return p.RemoveVerifiedClientDataCap(network, height, msg.Params)
	case parser.MethodRemoveExpiredAllocations, parser.MethodRemoveExpiredAllocationsExported:
		return p.RemoveExpiredAllocationsExported(network, height, msg.Params, msgRct.Return)
	case parser.MethodVerifiedDeprecated1: // UseBytes
		return p.Deprecated1(network, height, msg.Params)
	case parser.MethodVerifiedDeprecated2: // RestoreBytes
		return p.Deprecated2(network, height, msg.Params)
	case parser.MethodClaimAllocations:
		return p.ClaimAllocations(network, height, msg.Params, msgRct.Return)
	case parser.MethodGetClaims, parser.MethodGetClaimsExported: // TODO: not tested
		return p.GetClaimsExported(network, height, msg.Params, msgRct.Return)
	case parser.MethodExtendClaimTerms, parser.MethodExtendClaimTermsExported: // TODO: not tested
		return p.ExtendClaimTermsExported(network, height, msg.Params, msgRct.Return)
	case parser.MethodRemoveExpiredClaims, parser.MethodRemoveExpiredClaimsExported:
		return p.RemoveExpiredClaimsExported(network, height, msg.Params, msgRct.Return)
	case parser.MethodUniversalReceiverHook:
		return p.UniversalReceiverHook(network, height, msg.Params, msgRct.Return)
	}
	return map[string]interface{}{}, parser.ErrUnknownMethod
}
