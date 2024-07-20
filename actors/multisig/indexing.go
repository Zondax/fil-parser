package multisig

import (
	"encoding/json"
	"fmt"

	multisig2 "github.com/filecoin-project/go-state-types/builtin/v14/multisig"
	"github.com/filecoin-project/go-state-types/builtin/v8/miner"
	"github.com/filecoin-project/go-state-types/builtin/v8/verifreg"
	"github.com/filecoin-project/go-state-types/exitcode"
	"github.com/zondax/fil-parser/parser"
)

func ParseMultisigMetadata(txType string, txMetadata string) (interface{}, error) {
	deserializationFuncs := map[string]func(string) (interface{}, error){
		parser.MethodAddSigner:                           parseAddSignerValue,
		parser.MethodApprove:                             parseApproveValue,
		parser.MethodCancel:                              parseCancelValue,
		parser.MethodChangeNumApprovalsThreshold:         parseChangeNumApprovalsThresholdValue,
		parser.MethodConstructor:                         parseConstructorValue,
		parser.MethodLockBalance:                         parseLockBalanceValue,
		parser.MethodRemoveSigner:                        parseRemoveSignerValue,
		parser.MethodSend:                                parseSendValue,
		parser.MethodSwapSigner:                          parseSwapSignerValue,
		parser.MethodAddVerifier:                         parseAddVerifierValue,
		parser.MethodChangeOwnerAddress:                  parseChangeOwnerAddressValue,
		parser.MethodWithdrawBalance:                     parseWithdrawBalanceValue,
		parser.MethodInvokeContract:                      parseInvokeContractValue,
		parser.MethodApproveExported:                     parseApproveValue,
		parser.MethodCancelExported:                      parseCancelValue,
		parser.MethodAddSignerExported:                   parseAddSignerValue,
		parser.MethodSwapSignerExported:                  parseSwapSignerValue,
		parser.MethodRemoveSignerExported:                parseRemoveSignerValue,
		parser.MethodChangeNumApprovalsThresholdExported: parseChangeNumApprovalsThresholdValue,
		parser.MethodLockBalanceExported:                 parseLockBalanceValue,
		parser.MethodMsigUniversalReceiverHook:           parseUniversalReceiverHookValue,
		parser.MethodChangeOwnerAddressExported:          parseChangeOwnerAddressValue,
		parser.MethodWithdrawBalanceExported:             parseWithdrawBalanceValue,
	}

	if parseFunc, found := deserializationFuncs[txType]; found {
		return parseFunc(txMetadata)
	}

	return nil, fmt.Errorf("unknown tx type: %s", txType)
}

func parseAddVerifierValue(txMetadata string) (interface{}, error) {
	var v verifreg.AddVerifierParams
	err := json.Unmarshal([]byte(txMetadata), &v)
	return v, err
}

func parseChangeOwnerAddressValue(txMetadata string) (interface{}, error) {
	var v ChangeOwnerAddressParams
	err := json.Unmarshal([]byte(txMetadata), &v)
	return v, err
}

func parseWithdrawBalanceValue(txMetadata string) (interface{}, error) {
	var v miner.WithdrawBalanceParams
	err := json.Unmarshal([]byte(txMetadata), &v)
	return v, err
}

func parseInvokeContractValue(txMetadata string) (interface{}, error) {
	var v InvokeContractParams
	err := json.Unmarshal([]byte(txMetadata), &v)
	return v, err
}

func parseAddSignerValue(txMetadata string) (interface{}, error) {
	var v multisig2.AddSignerParams
	err := json.Unmarshal([]byte(txMetadata), &v)
	return v, err
}

func parseApproveValue(txMetadata string) (interface{}, error) {
	var raw map[string]interface{}
	err := json.Unmarshal([]byte(txMetadata), &raw)
	if err != nil {
		return nil, err
	}

	paramsRaw, ok := raw["Params"].(string)
	if !ok {
		return nil, fmt.Errorf("Params not found or not a string")
	}

	var params ApproveValue
	err = json.Unmarshal([]byte(paramsRaw), &params)
	if err != nil {
		return nil, err
	}

	returnRaw, ok := raw["Return"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("Return not found or not a map")
	}

	applied, ok := returnRaw["Applied"].(bool)
	if !ok {
		return nil, fmt.Errorf("Applied not found or not a bool")
	}

	code, ok := returnRaw["Code"].(float64)
	if !ok {
		return nil, fmt.Errorf("Code not found or not a float64")
	}

	ret, ok := returnRaw["Ret"].(string)
	if !ok {
		return nil, fmt.Errorf("Ret not found or not a string")
	}

	params.Return = multisig2.ApproveReturn{
		Applied: applied,
		Code:    exitcode.ExitCode(code),
		Ret:     []byte(ret),
	}

	return params, nil
}

func parseCancelValue(txMetadata string) (interface{}, error) {
	var raw map[string]interface{}
	err := json.Unmarshal([]byte(txMetadata), &raw)
	if err != nil {
		return nil, err
	}

	paramsStr, ok := raw["Params"].(string)
	if !ok {
		return nil, fmt.Errorf("Params not found or not a string")
	}

	var paramsRaw map[string]interface{}
	err = json.Unmarshal([]byte(paramsStr), &paramsRaw)
	if err != nil {
		return nil, err
	}

	var v CancelValue
	err = mapToStruct(paramsRaw, &v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func parseChangeNumApprovalsThresholdValue(txMetadata string) (interface{}, error) {
	var raw map[string]interface{}
	err := json.Unmarshal([]byte(txMetadata), &raw)
	if err != nil {
		return nil, err
	}

	paramsRaw, ok := raw["Params"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("params not found or not a map")
	}

	var v multisig2.ChangeNumApprovalsThresholdParams
	if newThreshold, ok := paramsRaw["NewThreshold"].(float64); ok {
		v.NewThreshold = uint64(newThreshold)
	} else {
		return nil, fmt.Errorf("NewThreshold not found or not a number")
	}

	return v, nil
}

func parseConstructorValue(txMetadata string) (interface{}, error) {
	var raw map[string]interface{}
	err := json.Unmarshal([]byte(txMetadata), &raw)
	if err != nil {
		return nil, err
	}

	paramsRaw, ok := raw["Params"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("params not found or not a map")
	}

	var v multisig2.ConstructorParams
	err = mapToStruct(paramsRaw, &v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func parseLockBalanceValue(txMetadata string) (interface{}, error) {
	var raw map[string]interface{}
	err := json.Unmarshal([]byte(txMetadata), &raw)
	if err != nil {
		return nil, err
	}

	paramsRaw, ok := raw["Params"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("params not found or not a map")
	}

	var v multisig2.LockBalanceParams
	err = mapToStruct(paramsRaw, &v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func parseRemoveSignerValue(txMetadata string) (interface{}, error) {
	var raw map[string]interface{}
	err := json.Unmarshal([]byte(txMetadata), &raw)
	if err != nil {
		return nil, err
	}
	paramsRaw, ok := raw["Params"].(string)
	if !ok {
		return nil, fmt.Errorf("Params not found or not a string")
	}
	var params multisig2.RemoveSignerParams
	err = json.Unmarshal([]byte(paramsRaw), &params)
	if err != nil {
		return nil, err
	}
	return params, nil
}

func parseSendValue(txMetadata string) (interface{}, error) {
	var v SendValue
	err := json.Unmarshal([]byte(txMetadata), &v)
	return v, err
}

func parseSwapSignerValue(txMetadata string) (interface{}, error) {
	var v multisig2.SwapSignerParams
	err := json.Unmarshal([]byte(txMetadata), &v)
	return v, err
}

func parseUniversalReceiverHookValue(txMetadata string) (interface{}, error) {
	var tx TransactionUniversalReceiverHookMetadata
	err := json.Unmarshal([]byte(txMetadata), &tx)
	if err != nil {
		return nil, err
	}

	var params UniversalReceiverHookParams
	err = json.Unmarshal([]byte(tx.Params), &params)
	if err != nil {
		return nil, err
	}

	result := UniversalReceiverHookValue{
		Type:    uint64(params.Type_),
		Payload: params.Payload,
		Return:  tx.Return,
	}

	return result, nil
}

func mapToStruct(m map[string]interface{}, v interface{}) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}
