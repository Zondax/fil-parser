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
	multisig8 "github.com/filecoin-project/go-state-types/builtin/v8/multisig"
	multisig9 "github.com/filecoin-project/go-state-types/builtin/v9/multisig"
	"github.com/filecoin-project/go-state-types/exitcode"
	"github.com/zondax/fil-parser/tools"
)

func toBytes(raw any) ([]byte, error) {
	switch v := raw.(type) {
	case []byte:
		return v, nil
	case string:
		return []byte(v), nil
	}
	return nil, errors.New("invalid type")
}

func mapToStruct(m map[string]interface{}, v interface{}) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
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

	switch {
	case tools.V8.IsSupported(height):
		params.Return = multisig8.ApproveReturn{
			Applied: applied,
			Code:    exitcode.ExitCode(code),
			Ret:     []byte(ret),
		}
	case tools.V9.IsSupported(height):
		params.Return = multisig9.ApproveReturn{
			Applied: applied,
			Code:    exitcode.ExitCode(code),
			Ret:     []byte(ret),
		}
	case tools.V10.IsSupported(height):
		params.Return = multisig10.ApproveReturn{
			Applied: applied,
			Code:    exitcode.ExitCode(code),
			Ret:     []byte(ret),
		}
	case tools.V11.IsSupported(height):
		params.Return = multisig11.ApproveReturn{
			Applied: applied,
			Code:    exitcode.ExitCode(code),
			Ret:     []byte(ret),
		}
	case tools.V12.IsSupported(height):
		params.Return = multisig12.ApproveReturn{
			Applied: applied,
			Code:    exitcode.ExitCode(code),
			Ret:     []byte(ret),
		}
	case tools.V13.IsSupported(height):
		params.Return = multisig13.ApproveReturn{
			Applied: applied,
			Code:    exitcode.ExitCode(code),
			Ret:     []byte(ret),
		}
	case tools.V14.IsSupported(height):
		params.Return = multisig14.ApproveReturn{
			Applied: applied,
			Code:    exitcode.ExitCode(code),
			Ret:     []byte(ret),
		}
	case tools.V15.IsSupported(height):
		params.Return = multisig15.ApproveReturn{
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
	switch {
	case tools.V8.IsSupported(height):
		v = multisig8.ChangeNumApprovalsThresholdParams{
			NewThreshold: newValue,
		}
	case tools.V9.IsSupported(height):
		v = multisig9.ChangeNumApprovalsThresholdParams{
			NewThreshold: newValue,
		}
	case tools.V10.IsSupported(height):
		v = multisig10.ChangeNumApprovalsThresholdParams{
			NewThreshold: newValue,
		}
	case tools.V11.IsSupported(height):
		v = multisig11.ChangeNumApprovalsThresholdParams{
			NewThreshold: newValue,
		}
	case tools.V12.IsSupported(height):
		v = multisig12.ChangeNumApprovalsThresholdParams{
			NewThreshold: newValue,
		}
	case tools.V13.IsSupported(height):
		v = multisig13.ChangeNumApprovalsThresholdParams{
			NewThreshold: newValue,
		}
	case tools.V14.IsSupported(height):
		v = multisig14.ChangeNumApprovalsThresholdParams{
			NewThreshold: newValue,
		}
	case tools.V15.IsSupported(height):
		v = multisig15.ChangeNumApprovalsThresholdParams{
			NewThreshold: newValue,
		}
	}

	return v, nil
}
