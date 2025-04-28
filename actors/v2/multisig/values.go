package multisig

import (
	"encoding/json"
	"errors"
	"fmt"

	multisig10 "github.com/filecoin-project/go-state-types/builtin/v10/multisig"
	multisig11 "github.com/filecoin-project/go-state-types/builtin/v11/multisig"
	multisig12 "github.com/filecoin-project/go-state-types/builtin/v12/multisig"
	multisig13 "github.com/filecoin-project/go-state-types/builtin/v13/multisig"
	multisig14 "github.com/filecoin-project/go-state-types/builtin/v14/multisig"
	multisig15 "github.com/filecoin-project/go-state-types/builtin/v15/multisig"
	multisig16 "github.com/filecoin-project/go-state-types/builtin/v16/multisig"
	multisig8 "github.com/filecoin-project/go-state-types/builtin/v8/multisig"
	multisig9 "github.com/filecoin-project/go-state-types/builtin/v9/multisig"

	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/multisig"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/multisig"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/multisig"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/multisig"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/multisig"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/multisig"

	actor_tools "github.com/zondax/fil-parser/actors/v2/tools"

	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

func (*Msig) ParseChangeOwnerAddressValue(network string, height int64, txMetadata string) (interface{}, error) {
	return parseValue(txMetadata, &ChangeOwnerAddressParams{})
}

func (*Msig) ParseWithdrawBalanceValue(network string, height int64, txMetadata string) (interface{}, error) {
	withdrawBalanceParams, err := withdrawBalanceParams(network, height)
	if err != nil {
		return nil, err
	}
	return parseValue(txMetadata, withdrawBalanceParams)

}

func (*Msig) ParseInvokeContractValue(network string, height int64, txMetadata string) (interface{}, error) {
	return parseValue(txMetadata, &InvokeContractParams{})

}

func (*Msig) ParseAddSignerValue(network string, height int64, txMetadata string) (interface{}, error) {
	addSignerParams, err := addSignerParams(network, height)
	if err != nil {
		return nil, err
	}
	return parseValue(txMetadata, addSignerParams)
}

func (*Msig) ParseAddVerifierValue(network string, height int64, txMetadata string) (interface{}, error) {
	addVerifierParams, err := verifierParams(network, height)
	if err != nil {
		return nil, err
	}
	return parseValue(txMetadata, addVerifierParams)
}

func (*Msig) ParseApproveValue(network string, height int64, txMetadata string) (interface{}, error) {
	var raw map[string]interface{}
	err := json.Unmarshal([]byte(txMetadata), &raw)
	if err != nil {
		return nil, err
	}

	paramsRaw, ok := raw["Params"].(string)
	if !ok {
		return nil, errors.New("parseApproveValue: Params not found or not a string")
	}

	var params ApproveValue
	err = json.Unmarshal([]byte(paramsRaw), &params)
	if err != nil {
		return nil, err
	}
	ret, err := getApproveReturn(network, height, raw)
	if err != nil {
		return nil, err
	}
	params.Return = ret
	return params, nil
}

func (*Msig) ParseCancelValue(network string, height int64, txMetadata string) (interface{}, error) {
	var raw map[string]interface{}
	err := json.Unmarshal([]byte(txMetadata), &raw)
	if err != nil {
		return nil, err
	}

	paramsStr, ok := raw[parser.ParamsKey].(string)
	if !ok {
		return nil, errors.New("parseCancelValue: Params not found or not a string")
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

func (*Msig) ParseChangeNumApprovalsThresholdValue(network string, height int64, txMetadata string) (interface{}, error) {
	var raw map[string]interface{}
	err := json.Unmarshal([]byte(txMetadata), &raw)
	if err != nil {
		return nil, err
	}

	paramsRaw, ok := raw["Params"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("params not found or not a map")
	}

	threshold := uint64(0)
	if newThreshold, ok := paramsRaw["NewThreshold"].(float64); ok {
		threshold = uint64(newThreshold)
	} else {
		return nil, fmt.Errorf("NewThreshold not found or not a number")
	}

	var ret multisigParams
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V9)...):
		return parseValue(txMetadata, &legacyv2.ChangeNumApprovalsThresholdParams{
			NewThreshold: threshold,
		})
	case tools.AnyIsSupported(network, height, tools.V10, tools.V11):
		return parseValue(txMetadata, &legacyv3.ChangeNumApprovalsThresholdParams{
			NewThreshold: threshold,
		})
	case tools.V12.IsSupported(network, height):
		return parseValue(txMetadata, &legacyv4.ChangeNumApprovalsThresholdParams{
			NewThreshold: threshold,
		})
	case tools.V13.IsSupported(network, height):
		return parseValue(txMetadata, &legacyv5.ChangeNumApprovalsThresholdParams{
			NewThreshold: threshold,
		})
	case tools.V14.IsSupported(network, height):
		return parseValue(txMetadata, &legacyv6.ChangeNumApprovalsThresholdParams{
			NewThreshold: threshold,
		})
	case tools.V15.IsSupported(network, height):
		return parseValue(txMetadata, &legacyv7.ChangeNumApprovalsThresholdParams{
			NewThreshold: threshold,
		})

	case tools.V16.IsSupported(network, height):
		ret = &multisig8.ChangeNumApprovalsThresholdParams{
			NewThreshold: threshold,
		}
	case tools.V17.IsSupported(network, height):
		ret = &multisig9.ChangeNumApprovalsThresholdParams{
			NewThreshold: threshold,
		}
	case tools.V18.IsSupported(network, height):
		ret = &multisig10.ChangeNumApprovalsThresholdParams{
			NewThreshold: threshold,
		}
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		ret = &multisig11.ChangeNumApprovalsThresholdParams{
			NewThreshold: threshold,
		}
	case tools.V21.IsSupported(network, height):
		ret = &multisig12.ChangeNumApprovalsThresholdParams{
			NewThreshold: threshold,
		}
	case tools.V22.IsSupported(network, height):
		ret = &multisig13.ChangeNumApprovalsThresholdParams{
			NewThreshold: threshold,
		}
	case tools.V24.IsSupported(network, height):
		ret = &multisig15.ChangeNumApprovalsThresholdParams{
			NewThreshold: threshold,
		}
	case tools.V25.IsSupported(network, height):
		ret = &multisig16.ChangeNumApprovalsThresholdParams{
			NewThreshold: threshold,
		}
	default:
		return nil, fmt.Errorf("%w: %d", actor_tools.ErrUnsupportedHeight, height)
	}
	return ret, nil
}

func (*Msig) ParseConstructorValue(network string, height int64, txMetadata string) (interface{}, error) {
	raw := map[string]interface{}{}
	if err := json.Unmarshal([]byte(txMetadata), &raw); err != nil {
		return nil, err
	}
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V9)...):
		return parseValue(txMetadata, &legacyv2.ConstructorParams{})
	case tools.AnyIsSupported(network, height, tools.V10, tools.V11):
		return parseValue(txMetadata, &legacyv3.ConstructorParams{})
	case tools.V12.IsSupported(network, height):
		return parseValue(txMetadata, &legacyv4.ConstructorParams{})
	case tools.V13.IsSupported(network, height):
		return parseValue(txMetadata, &legacyv5.ConstructorParams{})
	case tools.V14.IsSupported(network, height):
		return parseValue(txMetadata, &legacyv6.ConstructorParams{})
	case tools.V15.IsSupported(network, height):
		return parseValue(txMetadata, &legacyv7.ConstructorParams{})
	case tools.V16.IsSupported(network, height):
		return getValue(raw, &multisig8.ConstructorParams{})
	case tools.V17.IsSupported(network, height):
		return getValue(raw, &multisig9.ConstructorParams{})
	case tools.V18.IsSupported(network, height):
		return getValue(raw, &multisig10.ConstructorParams{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return getValue(raw, &multisig11.ConstructorParams{})
	case tools.V21.IsSupported(network, height):
		return getValue(raw, &multisig12.ConstructorParams{})
	case tools.V22.IsSupported(network, height):
		return getValue(raw, &multisig13.ConstructorParams{})
	case tools.V23.IsSupported(network, height):
		return getValue(raw, &multisig14.ConstructorParams{})
	case tools.V24.IsSupported(network, height):
		return getValue(raw, &multisig15.ConstructorParams{})
	case tools.V25.IsSupported(network, height):
		return getValue(raw, &multisig16.ConstructorParams{})
	}
	return nil, fmt.Errorf("%w: %d", actor_tools.ErrUnsupportedHeight, height)
}

func (*Msig) ParseLockBalanceValue(network string, height int64, txMetadata string) (interface{}, error) {
	raw := map[string]interface{}{}
	if err := json.Unmarshal([]byte(txMetadata), &raw); err != nil {
		return nil, err
	}
	lockBalanceParams, err := lockBalanceParams(network, height)
	if err != nil {
		return nil, err
	}
	return getValue(raw, lockBalanceParams)
}

func (*Msig) ParseRemoveSignerValue(network string, height int64, txMetadata string) (interface{}, error) {
	var raw map[string]interface{}
	if err := json.Unmarshal([]byte(txMetadata), &raw); err != nil {
		return nil, err
	}
	removeSignerParams, err := removeSignerParams(network, height)
	if err != nil {
		return nil, err
	}
	return getValue(raw, removeSignerParams)
}

func (*Msig) ParseSendValue(network string, height int64, txMetadata string) (interface{}, error) {
	var v SendValue
	err := json.Unmarshal([]byte(txMetadata), &v)
	return v, err
}

func (*Msig) ParseSwapSignerValue(network string, height int64, txMetadata string) (interface{}, error) {
	swapSignerParams, err := swapSignerParams(network, height)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(txMetadata), &swapSignerParams)
	return swapSignerParams, err
}

func (*Msig) ParseUniversalReceiverHookValue(network string, height int64, txMetadata string) (interface{}, error) {
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
		// #nosec G115
		Type:    uint64(params.Type_),
		Payload: params.Payload,
		Return:  tx.Return,
	}

	return result, nil
}
