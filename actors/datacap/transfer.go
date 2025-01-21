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
	datacapv9 "github.com/filecoin-project/go-state-types/builtin/v9/datacap"
	"github.com/zondax/fil-parser/parser"
)

type (
	transferParams     = unmarshalCBOR
	transferReturn     = unmarshalCBOR
	transferFromParams = unmarshalCBOR
	transferFromReturn = unmarshalCBOR
)

type transferParamsInterface interface {
	unmarshalCBOR
}
type transferReturnInterface interface {
	unmarshalCBOR
}

func TransferExported(height uint64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 9:
		return transferExportedv9(raw, rawReturn)
	case 10:
		return transferExportedv10(raw, rawReturn)
	case 11:
		return transferExportedv11(raw, rawReturn)
	case 12:
		return transferExportedv12(raw, rawReturn)
	case 13:
		return transferExportedv13(raw, rawReturn)
	case 14:
		return transferExportedv14(raw, rawReturn)
	case 15:
		return transferExportedv15(raw, rawReturn)
	}
	return nil, fmt.Errorf("not supported")
}

func transferGeneric[P transferParamsInterface, R transferReturnInterface](raw, rawReturn []byte, params P, r R) (map[string]interface{}, error) {
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

func transferExportedv9(raw, rawReturn []byte) (map[string]interface{}, error) {
	return transferGeneric[*datacapv9.TransferParams, *datacapv9.TransferReturn](raw, rawReturn, &datacapv9.TransferParams{}, &datacapv9.TransferReturn{})
}

func transferExportedv10(raw, rawReturn []byte) (map[string]interface{}, error) {
	return transferGeneric[*datacapv10.TransferParams, *datacapv10.TransferReturn](raw, rawReturn, &datacapv10.TransferParams{}, &datacapv10.TransferReturn{})
}

func transferExportedv11(raw, rawReturn []byte) (map[string]interface{}, error) {
	return transferGeneric[*datacapv11.TransferParams, *datacapv11.TransferReturn](raw, rawReturn, &datacapv11.TransferParams{}, &datacapv11.TransferReturn{})
}

func transferExportedv12(raw, rawReturn []byte) (map[string]interface{}, error) {
	return transferGeneric[*datacapv12.TransferParams, *datacapv12.TransferReturn](raw, rawReturn, &datacapv12.TransferParams{}, &datacapv12.TransferReturn{})
}

func transferExportedv13(raw, rawReturn []byte) (map[string]interface{}, error) {
	return transferGeneric[*datacapv13.TransferParams, *datacapv13.TransferReturn](raw, rawReturn, &datacapv13.TransferParams{}, &datacapv13.TransferReturn{})
}

func transferExportedv14(raw, rawReturn []byte) (map[string]interface{}, error) {
	return transferGeneric[*datacapv14.TransferParams, *datacapv14.TransferReturn](raw, rawReturn, &datacapv14.TransferParams{}, &datacapv14.TransferReturn{})
}

func transferExportedv15(raw, rawReturn []byte) (map[string]interface{}, error) {
	return transferGeneric[*datacapv15.TransferParams, *datacapv15.TransferReturn](raw, rawReturn, &datacapv15.TransferParams{}, &datacapv15.TransferReturn{})
}
