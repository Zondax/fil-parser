package datacap

import (
	"bytes"
	"fmt"

	datacapv10 "github.com/filecoin-project/go-state-types/builtin/v10/datacap"
	datacapv11 "github.com/filecoin-project/go-state-types/builtin/v11/datacap"
	datacapv12 "github.com/filecoin-project/go-state-types/builtin/v12/datacap"
	datacapv13 "github.com/filecoin-project/go-state-types/builtin/v13/datacap"
	datacapv14 "github.com/filecoin-project/go-state-types/builtin/v14/datacap"
	datacapv15 "github.com/filecoin-project/go-state-types/builtin/v15/datacap"
	"github.com/filecoin-project/go-state-types/builtin/v9/datacap"
	datacapv9 "github.com/filecoin-project/go-state-types/builtin/v9/datacap"
	"github.com/zondax/fil-parser/parser"
)

type (
	burnParams     = unmarshalCBOR
	burnReturn     = unmarshalCBOR
	destroyParams  = unmarshalCBOR
	burnFromParams = unmarshalCBOR
	burnFromReturn = unmarshalCBOR
)

func BurnExported(height uint64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 9:
		return burnExportedv9(raw, rawReturn)
	case 10:
		return burnExportedv10(raw, rawReturn)
	case 11:
		return burnExportedv11(raw, rawReturn)
	case 14:
		return burnExportedv14(raw, rawReturn)
	case 15:
		return burnExportedv15(raw, rawReturn)
	}
	return nil, fmt.Errorf("not supported")
}

func BurnFromExported(height uint64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 9:
		return burnFromExportedv9(raw, rawReturn)
	case 10:
		return burnFromExportedv10(raw, rawReturn)
	case 11:
		return burnFromExportedv11(raw, rawReturn)
	case 12:
		return burnFromExportedv12(raw, rawReturn)
	case 13:
		return burnFromExportedv13(raw, rawReturn)
	case 14:
		return burnFromExportedv14(raw, rawReturn)
	case 15:
		return burnFromExportedv15(raw, rawReturn)
	}
	return nil, fmt.Errorf("not supported")
}

func DestroyExported(height uint64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 9:
		return destroyExportedv9(raw, rawReturn)
	case 10:
		return destroyExportedv10(raw, rawReturn)
	case 11:
		return destroyExportedv11(raw, rawReturn)
	case 12:
		return destroyExportedv12(raw, rawReturn)
	case 13:
		return destroyExportedv13(raw, rawReturn)
	case 14:
		return destroyExportedv14(raw, rawReturn)
	case 15:
		return destroyExportedv15(raw, rawReturn)
	}
	return nil, fmt.Errorf("not supported")
}

func burnGeneric[P unmarshalCBOR, R unmarshalCBOR](raw []byte, rawReturn []byte, params P, r R) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params

	reader = bytes.NewReader(rawReturn)
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r
	return metadata, nil
}

// Burn Exported

func burnExportedv9(raw, rawReturn []byte) (map[string]interface{}, error) {
	return burnGeneric(raw, rawReturn, &datacapv9.BurnParams{}, &datacapv9.BurnReturn{})
}

func burnExportedv10(raw, rawReturn []byte) (map[string]interface{}, error) {
	return burnGeneric(raw, rawReturn, &datacapv10.BurnParams{}, &datacapv10.BurnReturn{})
}

func burnExportedv11(raw, rawReturn []byte) (map[string]interface{}, error) {
	return burnGeneric(raw, rawReturn, &datacapv11.BurnParams{}, &datacapv11.BurnReturn{})
}

func burnExportedv12(raw, rawReturn []byte) (map[string]interface{}, error) {
	return burnGeneric(raw, rawReturn, &datacapv12.BurnParams{}, &datacapv12.BurnReturn{})
}

func burnExportedv13(raw, rawReturn []byte) (map[string]interface{}, error) {
	return burnGeneric(raw, rawReturn, &datacapv13.BurnParams{}, &datacapv13.BurnReturn{})
}

func burnExportedv14(raw, rawReturn []byte) (map[string]interface{}, error) {
	return burnGeneric(raw, rawReturn, &datacapv14.BurnParams{}, &datacapv14.BurnReturn{})
}

func burnExportedv15(raw, rawReturn []byte) (map[string]interface{}, error) {
	return burnGeneric(raw, rawReturn, &datacapv15.BurnParams{}, &datacapv15.BurnReturn{})
}

// Destroy Exported

func destroyExportedv9(raw, rawReturn []byte) (map[string]interface{}, error) {
	return burnGeneric(raw, rawReturn, &datacapv9.DestroyParams{}, &datacapv9.BurnReturn{})
}

func destroyExportedv10(raw, rawReturn []byte) (map[string]interface{}, error) {
	return burnGeneric(raw, rawReturn, &datacapv10.DestroyParams{}, &datacapv10.BurnReturn{})
}

func destroyExportedv11(raw, rawReturn []byte) (map[string]interface{}, error) {
	return burnGeneric(raw, rawReturn, &datacap.DestroyParams{}, &datacap.BurnReturn{})
}

func destroyExportedv12(raw, rawReturn []byte) (map[string]interface{}, error) {
	return burnGeneric(raw, rawReturn, &datacapv12.DestroyParams{}, &datacapv12.BurnReturn{})
}

func destroyExportedv13(raw, rawReturn []byte) (map[string]interface{}, error) {
	return burnGeneric(raw, rawReturn, &datacapv13.DestroyParams{}, &datacapv13.BurnReturn{})
}

func destroyExportedv14(raw, rawReturn []byte) (map[string]interface{}, error) {
	return burnGeneric(raw, rawReturn, &datacapv14.DestroyParams{}, &datacapv14.BurnReturn{})
}

func destroyExportedv15(raw, rawReturn []byte) (map[string]interface{}, error) {
	return burnGeneric(raw, rawReturn, &datacapv15.DestroyParams{}, &datacapv15.BurnReturn{})
}

// Burn From Exported

func burnFromExportedv9(raw, rawReturn []byte) (map[string]interface{}, error) {
	return burnGeneric(raw, rawReturn, &datacapv9.BurnFromParams{}, &datacapv9.BurnFromReturn{})
}

func burnFromExportedv10(raw, rawReturn []byte) (map[string]interface{}, error) {
	return burnGeneric(raw, rawReturn, &datacapv10.BurnFromParams{}, &datacapv10.BurnFromReturn{})
}

func burnFromExportedv11(raw, rawReturn []byte) (map[string]interface{}, error) {
	return burnGeneric(raw, rawReturn, &datacapv11.BurnFromParams{}, &datacapv11.BurnFromReturn{})
}

func burnFromExportedv12(raw, rawReturn []byte) (map[string]interface{}, error) {
	return burnGeneric(raw, rawReturn, &datacapv12.BurnFromParams{}, &datacapv12.BurnFromReturn{})
}

func burnFromExportedv13(raw, rawReturn []byte) (map[string]interface{}, error) {
	return burnGeneric(raw, rawReturn, &datacapv13.BurnFromParams{}, &datacapv13.BurnFromReturn{})
}

func burnFromExportedv14(raw, rawReturn []byte) (map[string]interface{}, error) {
	return burnGeneric(raw, rawReturn, &datacapv14.BurnFromParams{}, &datacapv14.BurnFromReturn{})
}

func burnFromExportedv15(raw, rawReturn []byte) (map[string]interface{}, error) {
	return burnGeneric(raw, rawReturn, &datacapv15.BurnFromParams{}, &datacapv15.BurnFromReturn{})
}
