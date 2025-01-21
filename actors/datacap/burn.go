package datacap

import (
	"bytes"
	"fmt"
	"io"

	datacapv10 "github.com/filecoin-project/go-state-types/builtin/v10/datacap"
	datacapv11 "github.com/filecoin-project/go-state-types/builtin/v11/datacap"
	datacapv12 "github.com/filecoin-project/go-state-types/builtin/v12/datacap"
	datacapv13 "github.com/filecoin-project/go-state-types/builtin/v13/datacap"
	datacapv14 "github.com/filecoin-project/go-state-types/builtin/v14/datacap"
	datacapv15 "github.com/filecoin-project/go-state-types/builtin/v15/datacap"
	datacapv9 "github.com/filecoin-project/go-state-types/builtin/v9/datacap"
	"github.com/zondax/fil-parser/parser"
)

type burnParams interface {
	UnmarshalCBOR(r io.Reader) error
}

type burnReturn interface {
	UnmarshalCBOR(r io.Reader) error
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

func BurnExported(height uint64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 9:
		return burnGeneric[*datacapv9.BurnParams, *datacapv9.BurnReturn](raw, rawReturn, &datacapv9.BurnParams{}, &datacapv9.BurnReturn{})
	case 10:
		return burnGeneric[*datacapv10.BurnParams, *datacapv10.BurnReturn](raw, rawReturn, &datacapv10.BurnParams{}, &datacapv10.BurnReturn{})
	case 11:
		return burnGeneric[*datacapv11.BurnParams, *datacapv11.BurnReturn](raw, rawReturn, &datacapv11.BurnParams{}, &datacapv11.BurnReturn{})
	case 14:
		return burnGeneric[*datacapv14.BurnParams, *datacapv14.BurnReturn](raw, rawReturn, &datacapv14.BurnParams{}, &datacapv14.BurnReturn{})
	case 15:
		return burnGeneric[*datacapv15.BurnParams, *datacapv15.BurnReturn](raw, rawReturn, &datacapv15.BurnParams{}, &datacapv15.BurnReturn{})
	}
	return nil, fmt.Errorf("not supported")
}

func BurnFromExported(height uint64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 9:
		return burnGeneric[*datacapv9.BurnFromParams, *datacapv9.BurnFromReturn](raw, rawReturn, &datacapv9.BurnFromParams{}, &datacapv9.BurnFromReturn{})
	case 10:
		return burnGeneric[*datacapv10.BurnFromParams, *datacapv10.BurnFromReturn](raw, rawReturn, &datacapv10.BurnFromParams{}, &datacapv10.BurnFromReturn{})
	case 11:
		return burnGeneric[*datacapv11.BurnFromParams, *datacapv11.BurnFromReturn](raw, rawReturn, &datacapv11.BurnFromParams{}, &datacapv11.BurnFromReturn{})
	case 12:
		return burnGeneric[*datacapv12.BurnFromParams, *datacapv12.BurnFromReturn](raw, rawReturn, &datacapv12.BurnFromParams{}, &datacapv12.BurnFromReturn{})
	case 13:
		return burnGeneric[*datacapv13.BurnFromParams, *datacapv13.BurnFromReturn](raw, rawReturn, &datacapv13.BurnFromParams{}, &datacapv13.BurnFromReturn{})
	case 14:
		return burnGeneric[*datacapv14.BurnFromParams, *datacapv14.BurnFromReturn](raw, rawReturn, &datacapv14.BurnFromParams{}, &datacapv14.BurnFromReturn{})
	case 15:
		return burnGeneric[*datacapv15.BurnFromParams, *datacapv15.BurnFromReturn](raw, rawReturn, &datacapv15.BurnFromParams{}, &datacapv15.BurnFromReturn{})
	}
	return nil, fmt.Errorf("not supported")
}

func DestroyExported(height uint64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 9:
		return burnGeneric[*datacapv9.DestroyParams, *datacapv9.BurnReturn](raw, rawReturn, &datacapv9.DestroyParams{}, &datacapv9.BurnReturn{})
	case 10:
		return burnGeneric[*datacapv10.DestroyParams, *datacapv10.BurnReturn](raw, rawReturn, &datacapv10.DestroyParams{}, &datacapv10.BurnReturn{})
	case 11:
		return burnGeneric[*datacapv11.DestroyParams, *datacapv11.BurnReturn](raw, rawReturn, &datacapv11.DestroyParams{}, &datacapv11.BurnReturn{})
	case 12:
		return burnGeneric[*datacapv12.DestroyParams, *datacapv12.BurnReturn](raw, rawReturn, &datacapv12.DestroyParams{}, &datacapv12.BurnReturn{})
	case 13:
		return burnGeneric[*datacapv13.DestroyParams, *datacapv13.BurnReturn](raw, rawReturn, &datacapv13.DestroyParams{}, &datacapv13.BurnReturn{})
	case 14:
		return burnGeneric[*datacapv14.DestroyParams, *datacapv14.BurnReturn](raw, rawReturn, &datacapv14.DestroyParams{}, &datacapv14.BurnReturn{})
	case 15:
		return burnGeneric[*datacapv15.DestroyParams, *datacapv15.BurnReturn](raw, rawReturn, &datacapv15.DestroyParams{}, &datacapv15.BurnReturn{})
	}
	return nil, fmt.Errorf("not supported")
}
