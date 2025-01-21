package datacap

import (
	"bytes"
	"fmt"

	"github.com/filecoin-project/go-state-types/abi"
	datacapv10 "github.com/filecoin-project/go-state-types/builtin/v10/datacap"
	datacapv11 "github.com/filecoin-project/go-state-types/builtin/v11/datacap"
	datacapv12 "github.com/filecoin-project/go-state-types/builtin/v12/datacap"
	datacapv13 "github.com/filecoin-project/go-state-types/builtin/v13/datacap"
	datacapv14 "github.com/filecoin-project/go-state-types/builtin/v14/datacap"
	datacapv15 "github.com/filecoin-project/go-state-types/builtin/v15/datacap"
	datacapv9 "github.com/filecoin-project/go-state-types/builtin/v9/datacap"
	"github.com/zondax/fil-parser/parser"
)

type (
	increaseAllowanceParams = unmarshalCBOR
	decreaseAllowanceParams = unmarshalCBOR
	revokeAllowanceParams   = unmarshalCBOR
	allowanceParams         = unmarshalCBOR
)

type allowanceParamsInterface interface {
	unmarshalCBOR
}

func IncreaseAllowance(height uint64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		return nil, fmt.Errorf("not supported")
	case 9:
		return increaseAllowanceExportedv9(raw, rawReturn)
	case 10:
		return increaseAllowanceExportedv10(raw, rawReturn)
	case 11:
		return increaseAllowanceExportedv11(raw, rawReturn)
	case 14:
		return increaseAllowanceExportedv14(raw, rawReturn)
	case 15:
		return increaseAllowanceExportedv15(raw, rawReturn)
	}
	return nil, fmt.Errorf("not supported")
}

func DecreaseAllowance(height uint64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		return nil, fmt.Errorf("not supported")
	case 9:
		return decreaseAllowanceExportedv9(raw, rawReturn)
	case 10:
		return decreaseAllowanceExportedv10(raw, rawReturn)
	case 11:
		return decreaseAllowanceExportedv11(raw, rawReturn)
	case 14:
		return decreaseAllowanceExportedv14(raw, rawReturn)
	case 15:
		return decreaseAllowanceExportedv15(raw, rawReturn)
	}
	return nil, fmt.Errorf("not supported")
}

func RevokeAllowance(height uint64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		return nil, fmt.Errorf("not supported")
	case 9:
		return revokeAllowanceExportedv9(raw, rawReturn)
	case 10:
		return revokeAllowanceExportedv10(raw, rawReturn)
	case 11:
		return revokeAllowanceExportedv11(raw, rawReturn)
	case 14:
		return revokeAllowanceExportedv14(raw, rawReturn)
	case 15:
		return revokeAllowanceExportedv15(raw, rawReturn)
	}
	return nil, fmt.Errorf("not supported")
}

func GetAllowance(height uint64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		return nil, fmt.Errorf("not supported")
	case 9:
		return allowanceExportedv9(raw, rawReturn)
	case 10:
		return allowanceExportedv10(raw, rawReturn)
	case 11:
		return allowanceExportedv11(raw, rawReturn)
	case 14:
		return allowanceExportedv14(raw, rawReturn)
	case 15:
		return allowanceExportedv15(raw, rawReturn)
	}
	return nil, fmt.Errorf("not supported")
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

// Increase Allowance

func increaseAllowanceExportedv9(raw, rawReturn []byte) (map[string]interface{}, error) {
	return allowanceGeneric(raw, rawReturn, &datacapv9.IncreaseAllowanceParams{})
}

func increaseAllowanceExportedv10(raw, rawReturn []byte) (map[string]interface{}, error) {
	return allowanceGeneric(raw, rawReturn, &datacapv10.IncreaseAllowanceParams{})
}

func increaseAllowanceExportedv11(raw, rawReturn []byte) (map[string]interface{}, error) {
	return allowanceGeneric(raw, rawReturn, &datacapv11.IncreaseAllowanceParams{})
}

func increaseAllowanceExportedv12(raw, rawReturn []byte) (map[string]interface{}, error) {
	return allowanceGeneric(raw, rawReturn, &datacapv12.IncreaseAllowanceParams{})
}

func increaseAllowanceExportedv13(raw, rawReturn []byte) (map[string]interface{}, error) {
	return allowanceGeneric(raw, rawReturn, &datacapv13.IncreaseAllowanceParams{})
}

func increaseAllowanceExportedv14(raw, rawReturn []byte) (map[string]interface{}, error) {
	return allowanceGeneric(raw, rawReturn, &datacapv14.IncreaseAllowanceParams{})
}

func increaseAllowanceExportedv15(raw, rawReturn []byte) (map[string]interface{}, error) {
	return allowanceGeneric(raw, rawReturn, &datacapv15.IncreaseAllowanceParams{})
}

// Decrease Allowance

func decreaseAllowanceExportedv9(raw, rawReturn []byte) (map[string]interface{}, error) {
	return allowanceGeneric(raw, rawReturn, &datacapv9.DecreaseAllowanceParams{})
}

func decreaseAllowanceExportedv10(raw, rawReturn []byte) (map[string]interface{}, error) {
	return allowanceGeneric(raw, rawReturn, &datacapv10.DecreaseAllowanceParams{})
}

func decreaseAllowanceExportedv11(raw, rawReturn []byte) (map[string]interface{}, error) {
	return allowanceGeneric(raw, rawReturn, &datacapv11.DecreaseAllowanceParams{})
}

func decreaseAllowanceExportedv12(raw, rawReturn []byte) (map[string]interface{}, error) {
	return allowanceGeneric(raw, rawReturn, &datacapv12.DecreaseAllowanceParams{})
}

func decreaseAllowanceExportedv13(raw, rawReturn []byte) (map[string]interface{}, error) {
	return allowanceGeneric(raw, rawReturn, &datacapv13.DecreaseAllowanceParams{})
}

func decreaseAllowanceExportedv14(raw, rawReturn []byte) (map[string]interface{}, error) {
	return allowanceGeneric(raw, rawReturn, &datacapv14.DecreaseAllowanceParams{})
}

func decreaseAllowanceExportedv15(raw, rawReturn []byte) (map[string]interface{}, error) {
	return allowanceGeneric(raw, rawReturn, &datacapv15.DecreaseAllowanceParams{})
}

// Revoke Allowance

func revokeAllowanceExportedv9(raw, rawReturn []byte) (map[string]interface{}, error) {
	return allowanceGeneric(raw, rawReturn, &datacapv9.RevokeAllowanceParams{})
}

func revokeAllowanceExportedv10(raw, rawReturn []byte) (map[string]interface{}, error) {
	return allowanceGeneric(raw, rawReturn, &datacapv10.RevokeAllowanceParams{})
}

func revokeAllowanceExportedv11(raw, rawReturn []byte) (map[string]interface{}, error) {
	return allowanceGeneric(raw, rawReturn, &datacapv11.RevokeAllowanceParams{})
}

func revokeAllowanceExportedv12(raw, rawReturn []byte) (map[string]interface{}, error) {
	return allowanceGeneric(raw, rawReturn, &datacapv12.RevokeAllowanceParams{})
}

func revokeAllowanceExportedv13(raw, rawReturn []byte) (map[string]interface{}, error) {
	return allowanceGeneric(raw, rawReturn, &datacapv13.RevokeAllowanceParams{})
}

func revokeAllowanceExportedv14(raw, rawReturn []byte) (map[string]interface{}, error) {
	return allowanceGeneric(raw, rawReturn, &datacapv14.RevokeAllowanceParams{})
}

func revokeAllowanceExportedv15(raw, rawReturn []byte) (map[string]interface{}, error) {
	return allowanceGeneric(raw, rawReturn, &datacapv15.RevokeAllowanceParams{})
}

// Get Allowance

func allowanceExportedv9(raw, rawReturn []byte) (map[string]interface{}, error) {
	return allowanceGeneric(raw, rawReturn, &datacapv9.GetAllowanceParams{})
}

func allowanceExportedv10(raw, rawReturn []byte) (map[string]interface{}, error) {
	return allowanceGeneric(raw, rawReturn, &datacapv10.GetAllowanceParams{})
}

func allowanceExportedv11(raw, rawReturn []byte) (map[string]interface{}, error) {
	return allowanceGeneric(raw, rawReturn, &datacapv11.GetAllowanceParams{})
}

func allowanceExportedv12(raw, rawReturn []byte) (map[string]interface{}, error) {
	return allowanceGeneric(raw, rawReturn, &datacapv12.GetAllowanceParams{})
}

func allowanceExportedv13(raw, rawReturn []byte) (map[string]interface{}, error) {
	return allowanceGeneric(raw, rawReturn, &datacapv13.GetAllowanceParams{})
}

func allowanceExportedv14(raw, rawReturn []byte) (map[string]interface{}, error) {
	return allowanceGeneric(raw, rawReturn, &datacapv14.GetAllowanceParams{})
}

func allowanceExportedv15(raw, rawReturn []byte) (map[string]interface{}, error) {
	return allowanceGeneric(raw, rawReturn, &datacapv15.GetAllowanceParams{})
}
