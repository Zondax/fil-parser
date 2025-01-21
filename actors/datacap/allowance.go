package datacap

import (
	"bytes"
	"fmt"

	"github.com/filecoin-project/go-state-types/abi"
	datacapv10 "github.com/filecoin-project/go-state-types/builtin/v10/datacap"
	datacapv11 "github.com/filecoin-project/go-state-types/builtin/v11/datacap"
	datacapv14 "github.com/filecoin-project/go-state-types/builtin/v14/datacap"
	datacapv15 "github.com/filecoin-project/go-state-types/builtin/v15/datacap"
	datacapv9 "github.com/filecoin-project/go-state-types/builtin/v9/datacap"
	"github.com/zondax/fil-parser/parser"
)

type allowanceParamsInterface interface {
	unmarshalCBOR
}

func allowanceGeneric[T allowanceParamsInterface](raw []byte, rawReturn []byte, params T) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params

	reader = bytes.NewReader(rawReturn)
	var r abi.TokenAmount
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r
	return metadata, nil
}

func IncreaseAllowance(height uint64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		return nil, fmt.Errorf("not supported")
	case 9:
		return allowanceGeneric[*datacapv9.IncreaseAllowanceParams](raw, rawReturn, &datacapv9.IncreaseAllowanceParams{})
	case 10:
		return allowanceGeneric[*datacapv10.IncreaseAllowanceParams](raw, rawReturn, &datacapv10.IncreaseAllowanceParams{})
	case 11:
		return allowanceGeneric[*datacapv11.IncreaseAllowanceParams](raw, rawReturn, &datacapv11.IncreaseAllowanceParams{})
	case 14:
		return allowanceGeneric[*datacapv14.IncreaseAllowanceParams](raw, rawReturn, &datacapv14.IncreaseAllowanceParams{})
	case 15:
		return allowanceGeneric[*datacapv15.IncreaseAllowanceParams](raw, rawReturn, &datacapv15.IncreaseAllowanceParams{})
	}
	return nil, fmt.Errorf("not supported")
}

func DecreaseAllowance(height uint64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		return nil, fmt.Errorf("not supported")
	case 9:
		return allowanceGeneric[*datacapv9.DecreaseAllowanceParams](raw, rawReturn, &datacapv9.DecreaseAllowanceParams{})
	case 10:
		return allowanceGeneric[*datacapv10.DecreaseAllowanceParams](raw, rawReturn, &datacapv10.DecreaseAllowanceParams{})
	case 11:
		return allowanceGeneric[*datacapv11.DecreaseAllowanceParams](raw, rawReturn, &datacapv11.DecreaseAllowanceParams{})
	case 14:
		return allowanceGeneric[*datacapv14.DecreaseAllowanceParams](raw, rawReturn, &datacapv14.DecreaseAllowanceParams{})
	case 15:
		return allowanceGeneric[*datacapv15.DecreaseAllowanceParams](raw, rawReturn, &datacapv15.DecreaseAllowanceParams{})
	}
	return nil, fmt.Errorf("not supported")
}

func RevokeAllowance(height uint64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		return nil, fmt.Errorf("not supported")
	case 9:
		return allowanceGeneric[*datacapv9.RevokeAllowanceParams](raw, rawReturn, &datacapv9.RevokeAllowanceParams{})
	case 10:
		return allowanceGeneric[*datacapv10.RevokeAllowanceParams](raw, rawReturn, &datacapv10.RevokeAllowanceParams{})
	case 11:
		return allowanceGeneric[*datacapv11.RevokeAllowanceParams](raw, rawReturn, &datacapv11.RevokeAllowanceParams{})
	case 14:
		return allowanceGeneric[*datacapv14.RevokeAllowanceParams](raw, rawReturn, &datacapv14.RevokeAllowanceParams{})
	case 15:
		return allowanceGeneric[*datacapv15.RevokeAllowanceParams](raw, rawReturn, &datacapv15.RevokeAllowanceParams{})
	}
	return nil, fmt.Errorf("not supported")
}

func GetAllowance(height uint64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		return nil, fmt.Errorf("not supported")
	case 9:
		return allowanceGeneric[*datacapv9.GetAllowanceParams](raw, rawReturn, &datacapv9.GetAllowanceParams{})
	case 10:
		return allowanceGeneric[*datacapv10.GetAllowanceParams](raw, rawReturn, &datacapv10.GetAllowanceParams{})
	case 11:
		return allowanceGeneric[*datacapv11.GetAllowanceParams](raw, rawReturn, &datacapv11.GetAllowanceParams{})
	case 14:
		return allowanceGeneric[*datacapv14.GetAllowanceParams](raw, rawReturn, &datacapv14.GetAllowanceParams{})
	case 15:
		return allowanceGeneric[*datacapv15.GetAllowanceParams](raw, rawReturn, &datacapv15.GetAllowanceParams{})
	}
	return nil, fmt.Errorf("not supported")
}
