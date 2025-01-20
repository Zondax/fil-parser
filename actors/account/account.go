package account

import (
	"bytes"
	"encoding/base64"

	"github.com/filecoin-project/go-address"
	"github.com/zondax/fil-parser/parser"
)

func PubkeyAddress(raw, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	metadata[parser.ParamsKey] = base64.StdEncoding.EncodeToString(raw)
	reader := bytes.NewReader(rawReturn)
	var r address.Address
	err := r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r.String()
	return metadata, nil
}

func AuthenticateMessage(height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 11:
		return authenticateMessagev11(raw, rawReturn)
	case 14:
		return authenticateMessagev14(raw, rawReturn)
	}
	return nil, nil
}
