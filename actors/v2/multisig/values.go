package multisig

import (
	"encoding/json"
	"fmt"

	multisig10 "github.com/filecoin-project/go-state-types/builtin/v10/multisig"

	miner11 "github.com/filecoin-project/go-state-types/builtin/v11/miner"

	miner10 "github.com/filecoin-project/go-state-types/builtin/v10/miner"
	miner12 "github.com/filecoin-project/go-state-types/builtin/v12/miner"
	miner13 "github.com/filecoin-project/go-state-types/builtin/v13/miner"
	miner14 "github.com/filecoin-project/go-state-types/builtin/v14/miner"
	miner15 "github.com/filecoin-project/go-state-types/builtin/v15/miner"
	miner8 "github.com/filecoin-project/go-state-types/builtin/v8/miner"
	miner9 "github.com/filecoin-project/go-state-types/builtin/v9/miner"
	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/miner"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/miner"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/miner"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/miner"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/miner"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/miner"

	multisig11 "github.com/filecoin-project/go-state-types/builtin/v11/multisig"
	multisig12 "github.com/filecoin-project/go-state-types/builtin/v12/multisig"
	multisig13 "github.com/filecoin-project/go-state-types/builtin/v13/multisig"
	multisig14 "github.com/filecoin-project/go-state-types/builtin/v14/multisig"
	multisig15 "github.com/filecoin-project/go-state-types/builtin/v15/multisig"
	multisig8 "github.com/filecoin-project/go-state-types/builtin/v8/multisig"
	multisig9 "github.com/filecoin-project/go-state-types/builtin/v9/multisig"

	verifreg10 "github.com/filecoin-project/go-state-types/builtin/v10/verifreg"
	verifreg11 "github.com/filecoin-project/go-state-types/builtin/v11/verifreg"
	verifreg12 "github.com/filecoin-project/go-state-types/builtin/v12/verifreg"
	verifreg13 "github.com/filecoin-project/go-state-types/builtin/v13/verifreg"
	verifreg14 "github.com/filecoin-project/go-state-types/builtin/v14/verifreg"
	verifreg15 "github.com/filecoin-project/go-state-types/builtin/v15/verifreg"
	verifreg8 "github.com/filecoin-project/go-state-types/builtin/v8/verifreg"
	verifreg9 "github.com/filecoin-project/go-state-types/builtin/v9/verifreg"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

func (*Msig) ParseChangeOwnerAddressValue(network string, height int64, txMetadata string) (interface{}, error) {
	return parseValue(txMetadata, &ChangeOwnerAddressParams{})
}

func (*Msig) ParseWithdrawBalanceValue(network string, height int64, txMetadata string) (interface{}, error) {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V9)...):
		return parseValue(txMetadata, &legacyv2.WithdrawBalanceParams{})
	case tools.AnyIsSupported(network, height, tools.V10, tools.V11):
		return parseValue(txMetadata, &legacyv3.WithdrawBalanceParams{})
	case tools.V12.IsSupported(network, height):
		return parseValue(txMetadata, &legacyv4.WithdrawBalanceParams{})
	case tools.V13.IsSupported(network, height):
		return parseValue(txMetadata, &legacyv5.WithdrawBalanceParams{})
	case tools.V14.IsSupported(network, height):
		return parseValue(txMetadata, &legacyv6.WithdrawBalanceParams{})
	case tools.V15.IsSupported(network, height):
		return parseValue(txMetadata, &legacyv7.WithdrawBalanceParams{})

	case tools.V16.IsSupported(network, height):
		return parseValue(txMetadata, &miner8.WithdrawBalanceParams{})
	case tools.V17.IsSupported(network, height):
		return parseValue(txMetadata, &miner9.WithdrawBalanceParams{})
	case tools.V18.IsSupported(network, height):
		return parseValue(txMetadata, &miner10.WithdrawBalanceParams{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseValue(txMetadata, &miner11.WithdrawBalanceParams{})
	case tools.V21.IsSupported(network, height):
		return parseValue(txMetadata, &miner12.WithdrawBalanceParams{})
	case tools.V22.IsSupported(network, height):
		return parseValue(txMetadata, &miner13.WithdrawBalanceParams{})
	case tools.V23.IsSupported(network, height):
		return parseValue(txMetadata, &miner14.WithdrawBalanceParams{})
	case tools.V24.IsSupported(network, height):
		return parseValue(txMetadata, &miner15.WithdrawBalanceParams{})
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Msig) ParseInvokeContractValue(network string, height int64, txMetadata string) (interface{}, error) {
	return parseValue(txMetadata, &InvokeContractParams{})

}

func (*Msig) ParseAddSignerValue(network string, height int64, txMetadata string) (interface{}, error) {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V15)...):
		return nil, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	case tools.V16.IsSupported(network, height):
		return parseValue(txMetadata, &multisig8.AddSignerParams{})
	case tools.V17.IsSupported(network, height):
		return parseValue(txMetadata, &multisig9.AddSignerParams{})
	case tools.V18.IsSupported(network, height):
		return parseValue(txMetadata, &multisig10.AddSignerParams{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseValue(txMetadata, &multisig11.AddSignerParams{})
	case tools.V21.IsSupported(network, height):
		return parseValue(txMetadata, &multisig12.AddSignerParams{})
	case tools.V22.IsSupported(network, height):
		return parseValue(txMetadata, &multisig13.AddSignerParams{})
	case tools.V23.IsSupported(network, height):
		return parseValue(txMetadata, &multisig14.AddSignerParams{})
	case tools.V24.IsSupported(network, height):
		return parseValue(txMetadata, &multisig15.AddSignerParams{})
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Msig) ParseAddVerifierValue(network string, height int64, txMetadata string) (interface{}, error) {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V15)...):
		return nil, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	case tools.V16.IsSupported(network, height):
		return parseValue(txMetadata, &verifreg8.AddVerifierParams{})
	case tools.V17.IsSupported(network, height):
		return parseValue(txMetadata, &verifreg9.AddVerifierParams{})
	case tools.V18.IsSupported(network, height):
		return parseValue(txMetadata, &verifreg10.AddVerifierParams{})
	case tools.AnyIsSupported(network, height, tools.V20, tools.V19):
		return parseValue(txMetadata, &verifreg11.AddVerifierParams{})
	case tools.V21.IsSupported(network, height):
		return parseValue(txMetadata, &verifreg12.AddVerifierParams{})
	case tools.V22.IsSupported(network, height):
		return parseValue(txMetadata, &verifreg13.AddVerifierParams{})
	case tools.V23.IsSupported(network, height):
		return parseValue(txMetadata, &verifreg14.AddVerifierParams{})
	case tools.V24.IsSupported(network, height):
		return parseValue(txMetadata, &verifreg15.AddVerifierParams{})
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Msig) ParseApproveValue(network string, height int64, txMetadata string) (interface{}, error) {
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
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V15)...):
		return nil, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
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
	default:
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return ret, nil
}

func (*Msig) ParseConstructorValue(network string, height int64, txMetadata string) (interface{}, error) {
	raw := map[string]interface{}{}
	if err := json.Unmarshal([]byte(txMetadata), &raw); err != nil {
		return nil, err
	}
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V15)...):
		return nil, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	case tools.V16.IsSupported(network, height):
		return getValue(height, raw, &multisig8.ConstructorParams{})
	case tools.V17.IsSupported(network, height):
		return getValue(height, raw, &multisig9.ConstructorParams{})
	case tools.V18.IsSupported(network, height):
		return getValue(height, raw, &multisig10.ConstructorParams{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return getValue(height, raw, &multisig11.ConstructorParams{})
	case tools.V21.IsSupported(network, height):
		return getValue(height, raw, &multisig12.ConstructorParams{})
	case tools.V22.IsSupported(network, height):
		return getValue(height, raw, &multisig13.ConstructorParams{})
	case tools.V23.IsSupported(network, height):
		return getValue(height, raw, &multisig14.ConstructorParams{})
	case tools.V24.IsSupported(network, height):
		return getValue(height, raw, &multisig15.ConstructorParams{})
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Msig) ParseLockBalanceValue(network string, height int64, txMetadata string) (interface{}, error) {
	raw := map[string]interface{}{}
	if err := json.Unmarshal([]byte(txMetadata), &raw); err != nil {
		return nil, err
	}
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V15)...):
		return nil, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	case tools.V16.IsSupported(network, height):
		return getValue(height, raw, &multisig8.LockBalanceParams{})
	case tools.V17.IsSupported(network, height):
		return getValue(height, raw, &multisig9.LockBalanceParams{})
	case tools.V18.IsSupported(network, height):
		return getValue(height, raw, &multisig10.LockBalanceParams{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return getValue(height, raw, &multisig11.LockBalanceParams{})
	case tools.V21.IsSupported(network, height):
		return getValue(height, raw, &multisig12.LockBalanceParams{})
	case tools.V22.IsSupported(network, height):
		return getValue(height, raw, &multisig13.LockBalanceParams{})
	case tools.V23.IsSupported(network, height):
		return getValue(height, raw, &multisig14.LockBalanceParams{})
	case tools.V24.IsSupported(network, height):
		return getValue(height, raw, &multisig15.LockBalanceParams{})
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Msig) ParseRemoveSignerValue(network string, height int64, txMetadata string) (interface{}, error) {
	var raw map[string]interface{}
	if err := json.Unmarshal([]byte(txMetadata), &raw); err != nil {
		return nil, err
	}
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V15)...):
		return nil, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	case tools.V16.IsSupported(network, height):
		return getValue(height, raw, &multisig8.RemoveSignerParams{})
	case tools.V17.IsSupported(network, height):
		return getValue(height, raw, &multisig9.RemoveSignerParams{})
	case tools.V18.IsSupported(network, height):
		return getValue(height, raw, &multisig10.RemoveSignerParams{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return getValue(height, raw, &multisig11.RemoveSignerParams{})
	case tools.V21.IsSupported(network, height):
		return getValue(height, raw, &multisig12.RemoveSignerParams{})
	case tools.V22.IsSupported(network, height):
		return getValue(height, raw, &multisig13.RemoveSignerParams{})
	case tools.V23.IsSupported(network, height):
		return getValue(height, raw, &multisig14.RemoveSignerParams{})
	case tools.V24.IsSupported(network, height):
		return getValue(height, raw, &multisig15.RemoveSignerParams{})
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Msig) ParseSendValue(network string, height int64, txMetadata string) (interface{}, error) {
	var v SendValue
	err := json.Unmarshal([]byte(txMetadata), &v)
	return v, err
}

func (*Msig) ParseSwapSignerValue(network string, height int64, txMetadata string) (interface{}, error) {
	var v multisigParams
	var err error
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V15)...):
		return nil, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	case tools.V16.IsSupported(network, height):
		v = &multisig8.SwapSignerParams{}
	case tools.V17.IsSupported(network, height):
		v = &multisig9.SwapSignerParams{}
	case tools.V18.IsSupported(network, height):
		v = &multisig10.SwapSignerParams{}
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		v = &multisig11.SwapSignerParams{}
	case tools.V21.IsSupported(network, height):
		v = &multisig12.SwapSignerParams{}
	case tools.V22.IsSupported(network, height):
		v = &multisig13.SwapSignerParams{}
	case tools.V23.IsSupported(network, height):
		v = &multisig14.SwapSignerParams{}
	case tools.V24.IsSupported(network, height):
		v = &multisig15.SwapSignerParams{}
	default:
		err = fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	if err == nil {
		err = json.Unmarshal([]byte(txMetadata), &v)
		return v, err
	}
	return nil, err
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
		Type:    uint64(params.Type_),
		Payload: params.Payload,
		Return:  tx.Return,
	}

	return result, nil
}
