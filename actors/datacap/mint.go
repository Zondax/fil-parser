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

func MintExported(height uint64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 9:
		var r datacapv9.MintReturn
		return mintGeneric[*datacapv9.MintParams, *datacapv9.MintReturn](raw, rawReturn, &datacapv9.MintParams{}, &r)
	case 10:
		var r datacapv10.MintReturn
		return mintGeneric[*datacapv10.MintParams, *datacapv10.MintReturn](raw, rawReturn, &datacapv10.MintParams{}, &r)
	case 11:
		var r datacapv11.MintReturn
		return mintGeneric[*datacapv11.MintParams, *datacapv11.MintReturn](raw, rawReturn, &datacapv11.MintParams{}, &r)
	case 12:
		var r datacapv12.MintReturn
		return mintGeneric[*datacapv12.MintParams, *datacapv12.MintReturn](raw, rawReturn, &datacapv12.MintParams{}, &r)
	case 13:
		var r datacapv13.MintReturn
		return mintGeneric[*datacapv13.MintParams, *datacapv13.MintReturn](raw, rawReturn, &datacapv13.MintParams{}, &r)
	case 14:
		var r datacapv14.MintReturn
		return mintGeneric[*datacapv14.MintParams, *datacapv14.MintReturn](raw, rawReturn, &datacapv14.MintParams{}, &r)
	case 15:
		var r datacapv15.MintReturn
		return mintGeneric[*datacapv15.MintParams, *datacapv15.MintReturn](raw, rawReturn, &datacapv15.MintParams{}, &r)
	}
	return nil, fmt.Errorf("not supported")
}
