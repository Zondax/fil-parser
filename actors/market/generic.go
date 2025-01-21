package market

import (
	"bytes"

	"github.com/zondax/fil-parser/parser"
)

func parseGeneric[T marketParam, R marketReturn](rawParams, rawReturn []byte, returnCustomParam bool) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params T
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}

	metadata[parser.ParamsKey] = params
	if !returnCustomParam {
		metadata[parser.ParamsKey] = getAddressAsString(params)
		return metadata, nil
	}

	reader = bytes.NewReader(rawReturn)
	var publishReturn R
	err = publishReturn.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = publishReturn
	return metadata, nil
}
