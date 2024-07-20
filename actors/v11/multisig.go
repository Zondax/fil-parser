package v11

import (
	multisigV11 "github.com/filecoin-project/go-state-types/builtin/v11/multisig"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/zondax/fil-parser/actors/multisig"
	"github.com/zondax/fil-parser/parser"
)

type V11MultisigParser struct {
	service multisig.MultisigMethods
}

func NewV11MultisigParser(service multisig.MultisigMethods) multisig.MultisigParser {
	return &V11MultisigParser{service: service}
}

func (p *V11MultisigParser) ParseMultisig(txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, height int64, key filTypes.TipSetKey) (map[string]interface{}, error) {
	switch txType {
	case parser.MethodConstructor:
		return p.service.ParseMsigConstructor(msg.Params, new(multisigV11.ConstructorParams))
	case parser.MethodChangeNumApprovalsThreshold:
		return p.service.ParseChangeNumApprovalsThreshold(msg.Params, new(multisigV11.ChangeNumApprovalsThresholdParams))
	case parser.MethodLockBalance:
		return p.service.ParseLockBalance(msg.Params, new(multisigV11.LockBalanceParams))
	case parser.MethodRemoveSigner:
		return p.service.ParseRemoveSigner(msg, height, key, new(multisigV11.RemoveSignerParams))
	case parser.MethodApprove:
		return p.service.ParseApprove(msg, msgRct.Return, height, key, new(multisigV11.TxnIDParams), new(multisigV11.ApproveReturn))
	case parser.MethodCancel:
		return p.service.ParseCancel(msg, height, key, new(multisigV11.TxnIDParams))
	case parser.MethodPropose:
		return p.service.ParsePropose(msg.Params, msgRct.Return, new(multisigV11.ProposeParams), new(multisigV11.ProposeReturn))

		// TODO: the rest of the methods
		//case parser.MethodAddSigner:
		//	return p.service.ParseAddSigner(msg.Params, new(multisigV11.AddSignerParams))
		//case parser.MethodSwapSigner:
		//	return p.service.ParseSwapSigner(msg.Params, new(multisigV11.SwapSignerParams))
		// case parser.MethodAddVerifier:
		// 	return p.service.ParseAddVerifier(msg.Params, new(multisigV11.AddVerifierParams))
		// case parser.MethodRemoveVerifier:
		// 	return p.service.ParseRemoveVerifier(msg.Params, new(multisigV11.RemoveVerifierParams))
		// case parser.MethodUniversalReceiverHook:
		// 	return p.service.ParseUniversalReceiverHook(msg.Params, new(multisigV11.UniversalReceiverHookParams))
	}
	return map[string]interface{}{}, parser.ErrUnknownMethod
}
