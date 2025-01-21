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
	mintParams = unmarshalCBOR
	mintReturn = unmarshalCBOR
)

func MintExported(height uint64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 9:
		return mintExportedv9(raw, rawReturn)
	case 10:
		return mintExportedv10(raw, rawReturn)
	case 11:
		return mintExportedv11(raw, rawReturn)
	case 12:
		return mintExportedv12(raw, rawReturn)
	case 13:
		return mintExportedv13(raw, rawReturn)
	case 14:
		return mintExportedv14(raw, rawReturn)
	case 15:
		return mintExportedv15(raw, rawReturn)
	}
	return nil, fmt.Errorf("not supported")
}

func mintGeneric[P mintParams, R mintReturn](raw, rawReturn []byte, params P, r R) (map[string]interface{}, error) {
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

func mintExportedv9(raw, rawReturn []byte) (map[string]interface{}, error) {
	return mintGeneric[*datacapv9.MintParams, *datacapv9.MintReturn](raw, rawReturn, &datacapv9.MintParams{}, &datacapv9.MintReturn{})
}

func mintExportedv10(raw, rawReturn []byte) (map[string]interface{}, error) {
	return mintGeneric[*datacapv10.MintParams, *datacapv10.MintReturn](raw, rawReturn, &datacapv10.MintParams{}, &datacapv10.MintReturn{})
}

func mintExportedv11(raw, rawReturn []byte) (map[string]interface{}, error) {
	return mintGeneric[*datacapv11.MintParams, *datacapv11.MintReturn](raw, rawReturn, &datacapv11.MintParams{}, &datacapv11.MintReturn{})
}

func mintExportedv12(raw, rawReturn []byte) (map[string]interface{}, error) {
	return mintGeneric[*datacapv12.MintParams, *datacapv12.MintReturn](raw, rawReturn, &datacapv12.MintParams{}, &datacapv12.MintReturn{})
}

func mintExportedv13(raw, rawReturn []byte) (map[string]interface{}, error) {
	return mintGeneric[*datacapv13.MintParams, *datacapv13.MintReturn](raw, rawReturn, &datacapv13.MintParams{}, &datacapv13.MintReturn{})
}

func mintExportedv14(raw, rawReturn []byte) (map[string]interface{}, error) {
	return mintGeneric[*datacapv14.MintParams, *datacapv14.MintReturn](raw, rawReturn, &datacapv14.MintParams{}, &datacapv14.MintReturn{})
}

func mintExportedv15(raw, rawReturn []byte) (map[string]interface{}, error) {
	return mintGeneric[*datacapv15.MintParams, *datacapv15.MintReturn](raw, rawReturn, &datacapv15.MintParams{}, &datacapv15.MintReturn{})
}
