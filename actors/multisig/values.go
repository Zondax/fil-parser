package multisig

import (
	"encoding/json"

	multisig10 "github.com/filecoin-project/go-state-types/builtin/v10/multisig"
	miner11 "github.com/filecoin-project/go-state-types/builtin/v11/miner"
	multisig11 "github.com/filecoin-project/go-state-types/builtin/v11/multisig"
	multisig12 "github.com/filecoin-project/go-state-types/builtin/v12/multisig"
	multisig13 "github.com/filecoin-project/go-state-types/builtin/v13/multisig"
	multisig14 "github.com/filecoin-project/go-state-types/builtin/v14/multisig"
	multisig15 "github.com/filecoin-project/go-state-types/builtin/v15/multisig"
	multisig8 "github.com/filecoin-project/go-state-types/builtin/v8/multisig"
	multisig9 "github.com/filecoin-project/go-state-types/builtin/v9/multisig"
	"github.com/zondax/fil-parser/tools"
)

func ChangeOwnerAddressValue(network string, height int64, txMetadata string) (interface{}, error) {
	switch {
	case tools.V16.IsSupported(network, height):
		return parse[*ChangeOwnerAddressParams, string](txMetadata, jsonUnmarshaller[*ChangeOwnerAddressParams])
	case tools.V17.IsSupported(network, height):
		return parse[*ChangeOwnerAddressParams, string](txMetadata, jsonUnmarshaller[*ChangeOwnerAddressParams])
	case tools.V18.IsSupported(network, height):
		return parse[*ChangeOwnerAddressParams, string](txMetadata, jsonUnmarshaller[*ChangeOwnerAddressParams])
	case tools.V19.IsSupported(network, height):
		return parse[*ChangeOwnerAddressParams, string](txMetadata, jsonUnmarshaller[*ChangeOwnerAddressParams])
	case tools.V21.IsSupported(network, height):
		return parse[*ChangeOwnerAddressParams, string](txMetadata, jsonUnmarshaller[*ChangeOwnerAddressParams])
	case tools.V22.IsSupported(network, height):
		return parse[*ChangeOwnerAddressParams, string](txMetadata, jsonUnmarshaller[*ChangeOwnerAddressParams])
	case tools.V23.IsSupported(network, height):
		return parse[*ChangeOwnerAddressParams, string](txMetadata, jsonUnmarshaller[*ChangeOwnerAddressParams])
	case tools.V24.IsSupported(network, height):
		return parse[*ChangeOwnerAddressParams, string](txMetadata, jsonUnmarshaller[*ChangeOwnerAddressParams])
	}
	return nil, nil
}

func ParseWithdrawBalanceValue(network string, height int64, txMetadata string) (interface{}, error) {
	switch {
	case tools.V16.IsSupported(network, height):
		return parse[*miner11.WithdrawBalanceParams, string](txMetadata, jsonUnmarshaller[*miner11.WithdrawBalanceParams])
	case tools.V17.IsSupported(network, height):
		return parse[*miner11.WithdrawBalanceParams, string](txMetadata, jsonUnmarshaller[*miner11.WithdrawBalanceParams])
	case tools.V18.IsSupported(network, height):
		return parse[*miner11.WithdrawBalanceParams, string](txMetadata, jsonUnmarshaller[*miner11.WithdrawBalanceParams])
	case tools.V19.IsSupported(network, height):
		return parse[*miner11.WithdrawBalanceParams, string](txMetadata, jsonUnmarshaller[*miner11.WithdrawBalanceParams])
	case tools.V21.IsSupported(network, height):
		return parse[*miner11.WithdrawBalanceParams, string](txMetadata, jsonUnmarshaller[*miner11.WithdrawBalanceParams])
	case tools.V22.IsSupported(network, height):
		return parse[*miner11.WithdrawBalanceParams, string](txMetadata, jsonUnmarshaller[*miner11.WithdrawBalanceParams])
	case tools.V23.IsSupported(network, height):
		return parse[*miner11.WithdrawBalanceParams, string](txMetadata, jsonUnmarshaller[*miner11.WithdrawBalanceParams])
	case tools.V24.IsSupported(network, height):
		return parse[*miner11.WithdrawBalanceParams, string](txMetadata, jsonUnmarshaller[*miner11.WithdrawBalanceParams])
	}
	return nil, nil
}

func ParseInvokeContractValue(network string, height int64, txMetadata string) (interface{}, error) {
	switch {
	case tools.V16.IsSupported(network, height):
		return parse[*InvokeContractParams, string](txMetadata, jsonUnmarshaller[*InvokeContractParams])
	case tools.V17.IsSupported(network, height):
		return parse[*InvokeContractParams, string](txMetadata, jsonUnmarshaller[*InvokeContractParams])
	case tools.V18.IsSupported(network, height):
		return parse[*InvokeContractParams, string](txMetadata, jsonUnmarshaller[*InvokeContractParams])
	case tools.V19.IsSupported(network, height):
		return parse[*InvokeContractParams, string](txMetadata, jsonUnmarshaller[*InvokeContractParams])
	case tools.V21.IsSupported(network, height):
		return parse[*InvokeContractParams, string](txMetadata, jsonUnmarshaller[*InvokeContractParams])
	case tools.V22.IsSupported(network, height):
		return parse[*InvokeContractParams, string](txMetadata, jsonUnmarshaller[*InvokeContractParams])
	case tools.V23.IsSupported(network, height):
		return parse[*InvokeContractParams, string](txMetadata, jsonUnmarshaller[*InvokeContractParams])
	case tools.V24.IsSupported(network, height):
		return parse[*InvokeContractParams, string](txMetadata, jsonUnmarshaller[*InvokeContractParams])
	}
	return nil, nil
}

func ParseAddSignerValue(network string, height int64, txMetadata string) (interface{}, error) {
	switch {
	case tools.V16.IsSupported(network, height):
		return parse[*multisig11.AddSignerParams, string](txMetadata, jsonUnmarshaller[*multisig11.AddSignerParams])
	case tools.V17.IsSupported(network, height):
		return parse[*multisig9.AddSignerParams, string](txMetadata, jsonUnmarshaller[*multisig9.AddSignerParams])
	case tools.V18.IsSupported(network, height):
		return parse[*multisig10.AddSignerParams, string](txMetadata, jsonUnmarshaller[*multisig10.AddSignerParams])
	case tools.V19.IsSupported(network, height):
		return parse[*multisig11.AddSignerParams, string](txMetadata, jsonUnmarshaller[*multisig11.AddSignerParams])
	case tools.V21.IsSupported(network, height):
		return parse[*multisig12.AddSignerParams, string](txMetadata, jsonUnmarshaller[*multisig12.AddSignerParams])
	case tools.V22.IsSupported(network, height):
		return parse[*multisig13.AddSignerParams, string](txMetadata, jsonUnmarshaller[*multisig13.AddSignerParams])
	case tools.V23.IsSupported(network, height):
		return parse[*multisig14.AddSignerParams, string](txMetadata, jsonUnmarshaller[*multisig14.AddSignerParams])
	case tools.V24.IsSupported(network, height):
		return parse[*multisig15.AddSignerParams, string](txMetadata, jsonUnmarshaller[*multisig15.AddSignerParams])
	}
	return nil, nil
}

func ParseApproveValue(network string, height int64, txMetadata string) (interface{}, error) {
	switch {
	case tools.V16.IsSupported(network, height):
		if data, err := parse[metadataWithCbor, string](txMetadata, jsonUnmarshaller[metadataWithCbor]); err != nil {
			return nil, err
		} else {
			return getApproveReturn(network, height, data)
		}
	case tools.V17.IsSupported(network, height):
		if data, err := parse[metadataWithCbor, string](txMetadata, jsonUnmarshaller[metadataWithCbor]); err != nil {
			return nil, err
		} else {
			return getApproveReturn(network, height, data)
		}
	case tools.V18.IsSupported(network, height):
		if data, err := parse[metadataWithCbor, string](txMetadata, jsonUnmarshaller[metadataWithCbor]); err != nil {
			return nil, err
		} else {
			return getApproveReturn(network, height, data)
		}
	case tools.V19.IsSupported(network, height):
		if data, err := parse[metadataWithCbor, string](txMetadata, jsonUnmarshaller[metadataWithCbor]); err != nil {
			return nil, err
		} else {
			return getApproveReturn(network, height, data)
		}
	case tools.V21.IsSupported(network, height):
		if data, err := parse[metadataWithCbor, string](txMetadata, jsonUnmarshaller[metadataWithCbor]); err != nil {
			return nil, err
		} else {
			return getApproveReturn(network, height, data)
		}
	case tools.V22.IsSupported(network, height):
		if data, err := parse[metadataWithCbor, string](txMetadata, jsonUnmarshaller[metadataWithCbor]); err != nil {
			return nil, err
		} else {
			return getApproveReturn(network, height, data)
		}
	case tools.V23.IsSupported(network, height):
		if data, err := parse[metadataWithCbor, string](txMetadata, jsonUnmarshaller[metadataWithCbor]); err != nil {
			return nil, err
		} else {
			return getApproveReturn(network, height, data)
		}
	case tools.V24.IsSupported(network, height):
		if data, err := parse[metadataWithCbor, string](txMetadata, jsonUnmarshaller[metadataWithCbor]); err != nil {
			return nil, err
		} else {
			return getApproveReturn(network, height, data)
		}
	}
	return nil, nil
}

func ParseCancelValue(network string, height int64, txMetadata string) (interface{}, error) {
	switch {
	case tools.V16.IsSupported(network, height):
		if data, err := parse[metadataWithCbor, string](txMetadata, jsonUnmarshaller[metadataWithCbor]); err != nil {
			return nil, err
		} else {
			return getCancelReturn(data)
		}
	case tools.V17.IsSupported(network, height):
		if data, err := parse[metadataWithCbor, string](txMetadata, jsonUnmarshaller[metadataWithCbor]); err != nil {
			return nil, err
		} else {
			return getCancelReturn(data)
		}
	case tools.V18.IsSupported(network, height):
		if data, err := parse[metadataWithCbor, string](txMetadata, jsonUnmarshaller[metadataWithCbor]); err != nil {
			return nil, err
		} else {
			return getCancelReturn(data)
		}
	case tools.V19.IsSupported(network, height):
		if data, err := parse[metadataWithCbor, string](txMetadata, jsonUnmarshaller[metadataWithCbor]); err != nil {
			return nil, err
		} else {
			return getCancelReturn(data)
		}
	case tools.V21.IsSupported(network, height):
		if data, err := parse[metadataWithCbor, string](txMetadata, jsonUnmarshaller[metadataWithCbor]); err != nil {
			return nil, err
		} else {
			return getCancelReturn(data)
		}
	case tools.V22.IsSupported(network, height):
		if data, err := parse[metadataWithCbor, string](txMetadata, jsonUnmarshaller[metadataWithCbor]); err != nil {
			return nil, err
		} else {
			return getCancelReturn(data)
		}
	case tools.V23.IsSupported(network, height):
		if data, err := parse[metadataWithCbor, string](txMetadata, jsonUnmarshaller[metadataWithCbor]); err != nil {
			return nil, err
		} else {
			return getCancelReturn(data)
		}
	case tools.V24.IsSupported(network, height):
		if data, err := parse[metadataWithCbor, string](txMetadata, jsonUnmarshaller[metadataWithCbor]); err != nil {
			return nil, err
		} else {
			return getCancelReturn(data)
		}
	}
	return nil, nil
}

func ChangeNumApprovalsThresholdValue(network string, height int64, txMetadata string) (interface{}, error) {
	switch {
	case tools.V16.IsSupported(network, height):
		return parse[*multisig8.ChangeNumApprovalsThresholdParams, string](txMetadata, jsonUnmarshaller[*multisig8.ChangeNumApprovalsThresholdParams])
	case tools.V17.IsSupported(network, height):
		return parse[*multisig9.ChangeNumApprovalsThresholdParams, string](txMetadata, jsonUnmarshaller[*multisig9.ChangeNumApprovalsThresholdParams])
	case tools.V18.IsSupported(network, height):
		return parse[*multisig10.ChangeNumApprovalsThresholdParams, string](txMetadata, jsonUnmarshaller[*multisig10.ChangeNumApprovalsThresholdParams])
	case tools.V19.IsSupported(network, height):
		return parse[*multisig11.ChangeNumApprovalsThresholdParams, string](txMetadata, jsonUnmarshaller[*multisig11.ChangeNumApprovalsThresholdParams])
	case tools.V21.IsSupported(network, height):
		return parse[*multisig12.ChangeNumApprovalsThresholdParams, string](txMetadata, jsonUnmarshaller[*multisig12.ChangeNumApprovalsThresholdParams])
	case tools.V22.IsSupported(network, height):
		return parse[*multisig13.ChangeNumApprovalsThresholdParams, string](txMetadata, jsonUnmarshaller[*multisig13.ChangeNumApprovalsThresholdParams])
	case tools.V23.IsSupported(network, height):
		return parse[*multisig14.ChangeNumApprovalsThresholdParams, string](txMetadata, jsonUnmarshaller[*multisig14.ChangeNumApprovalsThresholdParams])
	case tools.V24.IsSupported(network, height):
		return parse[*multisig15.ChangeNumApprovalsThresholdParams, string](txMetadata, jsonUnmarshaller[*multisig15.ChangeNumApprovalsThresholdParams])
	}
	return nil, nil
}

func ParseConstructorValue(network string, height int64, txMetadata string) (interface{}, error) {
	switch {
	case tools.V16.IsSupported(network, height):
		if data, err := parse[*multisig8.ConstructorParams, string](txMetadata, jsonUnmarshaller[*multisig8.ConstructorParams]); err != nil {
			return nil, err
		} else {
			return getValue[*multisig8.ConstructorParams](height, data)
		}
	case tools.V17.IsSupported(network, height):
		if data, err := parse[*multisig9.ConstructorParams, string](txMetadata, jsonUnmarshaller[*multisig9.ConstructorParams]); err != nil {
			return nil, err
		} else {
			return getValue[*multisig9.ConstructorParams](height, data)
		}
	case tools.V18.IsSupported(network, height):
		if data, err := parse[*multisig10.ConstructorParams, string](txMetadata, jsonUnmarshaller[*multisig10.ConstructorParams]); err != nil {
			return nil, err
		} else {
			return getValue[*multisig10.ConstructorParams](height, data)
		}
	case tools.V19.IsSupported(network, height):
		if data, err := parse[*multisig11.ConstructorParams, string](txMetadata, jsonUnmarshaller[*multisig11.ConstructorParams]); err != nil {
			return nil, err
		} else {
			return getValue[*multisig11.ConstructorParams](height, data)
		}
	case tools.V21.IsSupported(network, height):
		if data, err := parse[*multisig12.ConstructorParams, string](txMetadata, jsonUnmarshaller[*multisig12.ConstructorParams]); err != nil {
			return nil, err
		} else {
			return getValue[*multisig12.ConstructorParams](height, data)
		}
	case tools.V22.IsSupported(network, height):
		if data, err := parse[*multisig13.ConstructorParams, string](txMetadata, jsonUnmarshaller[*multisig13.ConstructorParams]); err != nil {
			return nil, err
		} else {
			return getValue[*multisig13.ConstructorParams](height, data)
		}
	case tools.V23.IsSupported(network, height):
		if data, err := parse[*multisig14.ConstructorParams, string](txMetadata, jsonUnmarshaller[*multisig14.ConstructorParams]); err != nil {
			return nil, err
		} else {
			return getValue[*multisig14.ConstructorParams](height, data)
		}
	case tools.V24.IsSupported(network, height):
		if data, err := parse[*multisig15.ConstructorParams, string](txMetadata, jsonUnmarshaller[*multisig15.ConstructorParams]); err != nil {
			return nil, err
		} else {
			return getValue[*multisig15.ConstructorParams](height, data)
		}
	}
	return nil, nil
}

func ParseLockBalanceValue(network string, height int64, txMetadata string) (interface{}, error) {
	switch {
	case tools.V16.IsSupported(network, height):
		if data, err := parse[*multisig8.LockBalanceParams, string](txMetadata, jsonUnmarshaller[*multisig8.LockBalanceParams]); err != nil {
			return nil, err
		} else {
			return getValue[*multisig8.LockBalanceParams](height, data)
		}
	case tools.V17.IsSupported(network, height):
		if data, err := parse[*multisig9.LockBalanceParams, string](txMetadata, jsonUnmarshaller[*multisig9.LockBalanceParams]); err != nil {
			return nil, err
		} else {
			return getValue[*multisig9.LockBalanceParams](height, data)
		}
	case tools.V18.IsSupported(network, height):
		if data, err := parse[*multisig10.LockBalanceParams, string](txMetadata, jsonUnmarshaller[*multisig10.LockBalanceParams]); err != nil {
			return nil, err
		} else {
			return getValue[*multisig10.LockBalanceParams](height, data)
		}
	case tools.V19.IsSupported(network, height):
		if data, err := parse[*multisig11.LockBalanceParams, string](txMetadata, jsonUnmarshaller[*multisig11.LockBalanceParams]); err != nil {
			return nil, err
		} else {
			return getValue[*multisig11.LockBalanceParams](height, data)
		}
	case tools.V21.IsSupported(network, height):
		if data, err := parse[*multisig12.LockBalanceParams, string](txMetadata, jsonUnmarshaller[*multisig12.LockBalanceParams]); err != nil {
			return nil, err
		} else {
			return getValue[*multisig12.LockBalanceParams](height, data)
		}
	case tools.V22.IsSupported(network, height):
		if data, err := parse[*multisig13.LockBalanceParams, string](txMetadata, jsonUnmarshaller[*multisig13.LockBalanceParams]); err != nil {
			return nil, err
		} else {
			return getValue[*multisig13.LockBalanceParams](height, data)
		}
	case tools.V23.IsSupported(network, height):
		if data, err := parse[*multisig14.LockBalanceParams, string](txMetadata, jsonUnmarshaller[*multisig14.LockBalanceParams]); err != nil {
			return nil, err
		} else {
			return getValue[*multisig14.LockBalanceParams](height, data)
		}
	case tools.V24.IsSupported(network, height):
		if data, err := parse[*multisig15.LockBalanceParams, string](txMetadata, jsonUnmarshaller[*multisig15.LockBalanceParams]); err != nil {
			return nil, err
		} else {
			return getValue[*multisig15.LockBalanceParams](height, data)
		}
	}
	return nil, nil
}

func ParseRemoveSignerValue(network string, height int64, txMetadata string) (interface{}, error) {
	switch {
	case tools.V16.IsSupported(network, height):
		if data, err := parse[*multisig8.RemoveSignerParams, string](txMetadata, jsonUnmarshaller[*multisig8.RemoveSignerParams]); err != nil {
			return nil, err
		} else {
			return getValue[*multisig8.RemoveSignerParams](height, data)
		}
	case tools.V17.IsSupported(network, height):
		if data, err := parse[*multisig9.RemoveSignerParams, string](txMetadata, jsonUnmarshaller[*multisig9.RemoveSignerParams]); err != nil {
			return nil, err
		} else {
			return getValue[*multisig9.RemoveSignerParams](height, data)
		}
	case tools.V18.IsSupported(network, height):
		if data, err := parse[*multisig10.RemoveSignerParams, string](txMetadata, jsonUnmarshaller[*multisig10.RemoveSignerParams]); err != nil {
			return nil, err
		} else {
			return getValue[*multisig10.RemoveSignerParams](height, data)
		}
	case tools.V19.IsSupported(network, height):
		if data, err := parse[*multisig11.RemoveSignerParams, string](txMetadata, jsonUnmarshaller[*multisig11.RemoveSignerParams]); err != nil {
			return nil, err
		} else {
			return getValue[*multisig11.RemoveSignerParams](height, data)
		}
	case tools.V21.IsSupported(network, height):
		if data, err := parse[*multisig12.RemoveSignerParams, string](txMetadata, jsonUnmarshaller[*multisig12.RemoveSignerParams]); err != nil {
			return nil, err
		} else {
			return getValue[*multisig12.RemoveSignerParams](height, data)
		}
	case tools.V22.IsSupported(network, height):
		if data, err := parse[*multisig13.RemoveSignerParams, string](txMetadata, jsonUnmarshaller[*multisig13.RemoveSignerParams]); err != nil {
			return nil, err
		} else {
			return getValue[*multisig13.RemoveSignerParams](height, data)
		}
	case tools.V23.IsSupported(network, height):
		if data, err := parse[*multisig14.RemoveSignerParams, string](txMetadata, jsonUnmarshaller[*multisig14.RemoveSignerParams]); err != nil {
			return nil, err
		} else {
			return getValue[*multisig14.RemoveSignerParams](height, data)
		}
	case tools.V24.IsSupported(network, height):
		if data, err := parse[*multisig15.RemoveSignerParams, string](txMetadata, jsonUnmarshaller[*multisig15.RemoveSignerParams]); err != nil {
			return nil, err
		} else {
			return getValue[*multisig15.RemoveSignerParams](height, data)
		}
	}
	return nil, nil
}

func ParseSendValue(network string, height int64, txMetadata string) (interface{}, error) {
	switch {
	case tools.V16.IsSupported(network, height):
		return parse[*SendValue, string](txMetadata, jsonUnmarshaller[*SendValue])
	case tools.V17.IsSupported(network, height):
		return parse[*SendValue, string](txMetadata, jsonUnmarshaller[*SendValue])
	case tools.V18.IsSupported(network, height):
		return parse[*SendValue, string](txMetadata, jsonUnmarshaller[*SendValue])
	case tools.V19.IsSupported(network, height):
		return parse[*SendValue, string](txMetadata, jsonUnmarshaller[*SendValue])
	case tools.V21.IsSupported(network, height):
		return parse[*SendValue, string](txMetadata, jsonUnmarshaller[*SendValue])
	case tools.V22.IsSupported(network, height):
		return parse[*SendValue, string](txMetadata, jsonUnmarshaller[*SendValue])
	case tools.V23.IsSupported(network, height):
		return parse[*SendValue, string](txMetadata, jsonUnmarshaller[*SendValue])
	case tools.V24.IsSupported(network, height):
		return parse[*SendValue, string](txMetadata, jsonUnmarshaller[*SendValue])
	}
	return nil, nil
}

func ParseSwapSignerValue(network string, height int64, txMetadata string) (interface{}, error) {
	switch {
	case tools.V16.IsSupported(network, height):
		return parse[*multisig8.SwapSignerParams, string](txMetadata, jsonUnmarshaller[*multisig8.SwapSignerParams])
	case tools.V17.IsSupported(network, height):
		return parse[*multisig9.SwapSignerParams, string](txMetadata, jsonUnmarshaller[*multisig9.SwapSignerParams])
	case tools.V18.IsSupported(network, height):
		return parse[*multisig10.SwapSignerParams, string](txMetadata, jsonUnmarshaller[*multisig10.SwapSignerParams])
	case tools.V19.IsSupported(network, height):
		return parse[*multisig11.SwapSignerParams, string](txMetadata, jsonUnmarshaller[*multisig11.SwapSignerParams])
	case tools.V21.IsSupported(network, height):
		return parse[*multisig12.SwapSignerParams, string](txMetadata, jsonUnmarshaller[*multisig12.SwapSignerParams])
	case tools.V22.IsSupported(network, height):
		return parse[*multisig13.SwapSignerParams, string](txMetadata, jsonUnmarshaller[*multisig13.SwapSignerParams])
	case tools.V23.IsSupported(network, height):
		return parse[*multisig14.SwapSignerParams, string](txMetadata, jsonUnmarshaller[*multisig14.SwapSignerParams])
	case tools.V24.IsSupported(network, height):
		return parse[*multisig15.SwapSignerParams, string](txMetadata, jsonUnmarshaller[*multisig15.SwapSignerParams])
	}
	return nil, nil
}

func ParseUniversalReceiverHookValue(network string, height int64, txMetadata string) (interface{}, error) {
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
