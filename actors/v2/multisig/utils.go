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
	"github.com/zondax/fil-parser/actors"
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

func getApproveReturn(network string, height int64, raw map[string]interface{}) (interface{}, error) {

	returnRaw, ok := raw["Return"].(map[string]interface{})
	if !ok {
		return nil, errors.New("getApproveReturn: Return not found or not a map")
	}

	applied, ok := returnRaw["Applied"].(bool)
	if !ok {
		return nil, errors.New("getApproveReturn: Applied not found or not a bool")
	}

	code, ok := returnRaw["Code"].(float64)
	if !ok {
		return nil, errors.New("getApproveReturn: Code not found or not a float64")
	}

	ret, ok := returnRaw["Ret"].(string)
	if !ok {
		return nil, errors.New("getApproveReturn: Ret not found or not a string")
	}

	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V15)...):
		return nil, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	case tools.V16.IsSupported(network, height):
		return multisig8.ApproveReturn{
			Applied: applied,
			Code:    exitcode.ExitCode(code),
			Ret:     []byte(ret),
		}, nil
	case tools.V17.IsSupported(network, height):
		return multisig9.ApproveReturn{
			Applied: applied,
			Code:    exitcode.ExitCode(code),
			Ret:     []byte(ret),
		}, nil
	case tools.V18.IsSupported(network, height):
		return multisig10.ApproveReturn{
			Applied: applied,
			Code:    exitcode.ExitCode(code),
			Ret:     []byte(ret),
		}, nil
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return multisig11.ApproveReturn{
			Applied: applied,
			Code:    exitcode.ExitCode(code),
			Ret:     []byte(ret),
		}, nil
	case tools.V21.IsSupported(network, height):
		return multisig12.ApproveReturn{
			Applied: applied,
			Code:    exitcode.ExitCode(code),
			Ret:     []byte(ret),
		}, nil
	case tools.V22.IsSupported(network, height):
		return multisig13.ApproveReturn{
			Applied: applied,
			Code:    exitcode.ExitCode(code),
			Ret:     []byte(ret),
		}, nil
	case tools.V23.IsSupported(network, height):
		return multisig14.ApproveReturn{
			Applied: applied,
			Code:    exitcode.ExitCode(code),
			Ret:     []byte(ret),
		}, nil
	case tools.V24.IsSupported(network, height):
		return multisig15.ApproveReturn{
			Applied: applied,
			Code:    exitcode.ExitCode(code),
			Ret:     []byte(ret),
		}, nil
	}

	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}
