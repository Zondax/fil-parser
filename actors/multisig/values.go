package multisig

import (
	"encoding/json"
	"fmt"

	miner11 "github.com/filecoin-project/go-state-types/builtin/v11/miner"
	multisig11 "github.com/filecoin-project/go-state-types/builtin/v11/multisig"
	"github.com/filecoin-project/go-state-types/exitcode"
)

func ChangeOwnerAddressValue(height int64, txMetadata string) (interface{}, error) {
	switch height {
	case 11:
		return parse[*ChangeOwnerAddressParams, string](txMetadata, jsonUnmarshaller[*ChangeOwnerAddressParams])
	}
	return nil, nil
}

func ParseWithdrawBalanceValue(height int64, txMetadata string) (interface{}, error) {
	switch height {
	case 11:
		return parse[*miner11.WithdrawBalanceParams, string](txMetadata, jsonUnmarshaller[*miner11.WithdrawBalanceParams])
	}
	return nil, nil
}

func ParseInvokeContractValue(height int64, txMetadata string) (interface{}, error) {
	switch height {
	case 11:
		return parse[*InvokeContractParams, string](txMetadata, jsonUnmarshaller[*InvokeContractParams])
	}
	return nil, nil
}

func ParseAddSignerValue(height int64, txMetadata string) (interface{}, error) {
	switch height {
	case 11:
		return parse[*multisig11.AddSignerParams, string](txMetadata, jsonUnmarshaller[*multisig11.AddSignerParams])
	}
	return nil, nil
}

func ParseApproveValue(height int64, txMetadata string) (interface{}, error) {
	switch height {
	case 11:
		if data, err := parse[metadataWithCbor, string](txMetadata, jsonUnmarshaller[metadataWithCbor]); err != nil {
			return nil, err
		} else {
			return getApproveReturn(height, data)
		}
	}
	return nil, nil
}

func ParseCancelValue(height int64, txMetadata string) (interface{}, error) {
	switch height {
	case 11:
		if data, err := parse[metadataWithCbor, string](txMetadata, jsonUnmarshaller[metadataWithCbor]); err != nil {
			return nil, err
		} else {
			return getCancelReturn(height, data)
		}
	}
	return nil, nil
}

func ChangeNumApprovalsThresholdValue(height int64, txMetadata string) (interface{}, error) {
	switch height {
	case 11:
		return parse[*multisig11.ChangeNumApprovalsThresholdParams, string](txMetadata, jsonUnmarshaller[*multisig11.ChangeNumApprovalsThresholdParams])
	}
	return nil, nil
}

func ParseConstructorValue(height int64, txMetadata string) (interface{}, error) {
	switch height {
	case 11:
		if data, err := parse[*multisig11.ConstructorParams, string](txMetadata, jsonUnmarshaller[*multisig11.ConstructorParams]); err != nil {
			return nil, err
		} else {
			return getValue[*multisig11.ConstructorParams](height, data)
		}
	}
	return nil, nil
}

func ParseLockBalanceValue(height int64, txMetadata string) (interface{}, error) {
	switch height {
	case 11:
		if data, err := parse[*multisig11.LockBalanceParams, string](txMetadata, jsonUnmarshaller[*multisig11.LockBalanceParams]); err != nil {
			return nil, err
		} else {
			return getValue[*multisig11.LockBalanceParams](height, data)
		}
	}
	return nil, nil
}

func ParseRemoveSignerValue(height int64, txMetadata string) (interface{}, error) {
	switch height {
	case 11:
		if data, err := parse[*multisig11.RemoveSignerParams, string](txMetadata, jsonUnmarshaller[*multisig11.RemoveSignerParams]); err != nil {
			return nil, err
		} else {
			return getValue[*multisig11.RemoveSignerParams](height, data)
		}
	}
	return nil, nil
}

func ParseSendValue(height int64, txMetadata string) (interface{}, error) {
	switch height {
	case 11:
		return parse[*SendValue, string](txMetadata, jsonUnmarshaller[*SendValue])
	}
	return nil, nil
}

func ParseSwapSignerValue(height int64, txMetadata string) (interface{}, error) {
	switch height {
	case 11:
		return parse[*multisig11.SwapSignerParams, string](txMetadata, jsonUnmarshaller[*multisig11.SwapSignerParams])
	}
	return nil, nil
}

func ParseUniversalReceiverHookValue(height int64, txMetadata string) (interface{}, error) {
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

func getApproveReturn(height int64, raw map[string]interface{}) (interface{}, error) {
	var params ApproveValue

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

	switch height {
	case 11:
		params.Return = multisig11.ApproveReturn{
			Applied: applied,
			Code:    exitcode.ExitCode(code),
			Ret:     []byte(ret),
		}
	}

	return params, nil

}

func getCancelReturn(height int64, raw map[string]interface{}) (interface{}, error) {
	paramsStr, ok := raw["Params"].(string)
	if !ok {
		return nil, fmt.Errorf("Params not found or not a string")
	}

	var paramsRaw map[string]interface{}
	err := json.Unmarshal([]byte(paramsStr), &paramsRaw)
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

func getChangeNumApprovalsThresholdValue(height int64, raw map[string]interface{}) (interface{}, error) {
	paramsStr, ok := raw["Params"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("Params not found or not a map")
	}

	var newValue uint64
	if newThreshold, ok := paramsStr["NewThreshold"].(float64); ok {
		newValue = uint64(newThreshold)
	} else {
		return nil, fmt.Errorf("NewThreshold not found or not a number")
	}
	var v any
	switch height {
	case 11:
		v = multisig11.ChangeNumApprovalsThresholdParams{
			NewThreshold: newValue,
		}
	}

	return v, nil
}
