package multisig

import (
	"fmt"

	"github.com/zondax/fil-parser/parser"
)

func (m *Msig) ParseMultisigMetadata(network string, height int64, txType string, txMetadata string) (interface{}, error) {
	deserializationFuncs := map[string]func(string, int64, string) (interface{}, error){
		parser.MethodAddSigner:                           m.ParseAddSignerValue,
		parser.MethodApprove:                             m.ParseApproveValue,
		parser.MethodCancel:                              m.ParseCancelValue,
		parser.MethodChangeNumApprovalsThreshold:         m.ParseChangeNumApprovalsThresholdValue,
		parser.MethodConstructor:                         m.ParseConstructorValue,
		parser.MethodLockBalance:                         m.ParseLockBalanceValue,
		parser.MethodRemoveSigner:                        m.ParseRemoveSignerValue,
		parser.MethodSend:                                m.ParseSendValue,
		parser.MethodSwapSigner:                          m.ParseSwapSignerValue,
		parser.MethodAddVerifier:                         m.ParseAddVerifierValue,
		parser.MethodChangeOwnerAddress:                  m.ParseChangeOwnerAddressValue,
		parser.MethodWithdrawBalance:                     m.ParseWithdrawBalanceValue,
		parser.MethodInvokeContract:                      m.ParseInvokeContractValue,
		parser.MethodApproveExported:                     m.ParseApproveValue,
		parser.MethodCancelExported:                      m.ParseCancelValue,
		parser.MethodAddSignerExported:                   m.ParseAddSignerValue,
		parser.MethodSwapSignerExported:                  m.ParseSwapSignerValue,
		parser.MethodRemoveSignerExported:                m.ParseRemoveSignerValue,
		parser.MethodChangeNumApprovalsThresholdExported: m.ParseChangeNumApprovalsThresholdValue,
		parser.MethodLockBalanceExported:                 m.ParseLockBalanceValue,
		parser.MethodMsigUniversalReceiverHook:           m.ParseUniversalReceiverHookValue,
		parser.MethodChangeOwnerAddressExported:          m.ParseChangeOwnerAddressValue,
		parser.MethodWithdrawBalanceExported:             m.ParseWithdrawBalanceValue,
	}

	if parseFunc, found := deserializationFuncs[txType]; found {
		return parseFunc(network, height, txMetadata)
	}

	return nil, fmt.Errorf("unknown tx type: %s", txType)
}
