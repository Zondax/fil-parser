package account

import (
	"bytes"
	"encoding/base64"
	"fmt"

	"github.com/filecoin-project/go-address"
	accountv10 "github.com/filecoin-project/go-state-types/builtin/v10/account"
	accountv11 "github.com/filecoin-project/go-state-types/builtin/v11/account"
	accountv9 "github.com/filecoin-project/go-state-types/builtin/v11/account"
	accountv14 "github.com/filecoin-project/go-state-types/builtin/v14/account"
	accountv15 "github.com/filecoin-project/go-state-types/builtin/v15/account"
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
	case 8:
		return nil, fmt.Errorf("not supported")
	case 9:
		return authenticateMessageGeneric[*accountv9.AuthenticateMessageParams, *accountv9.AuthenticateMessageParams](raw, rawReturn, &accountv9.AuthenticateMessageParams{})
	case 10:
		return authenticateMessageGeneric[*accountv10.AuthenticateMessageParams, *accountv10.AuthenticateMessageParams](raw, rawReturn, &accountv10.AuthenticateMessageParams{})
	case 11:
		return authenticateMessageGeneric[*accountv11.AuthenticateMessageParams, *accountv11.AuthenticateMessageParams](raw, rawReturn, &accountv11.AuthenticateMessageParams{})
	case 14:
		return authenticateMessageGeneric[*accountv14.AuthenticateMessageParams, *accountv14.AuthenticateMessageParams](raw, rawReturn, &accountv14.AuthenticateMessageParams{})
	case 15:
		return authenticateMessageGeneric[*accountv15.AuthenticateMessageParams, *accountv15.AuthenticateMessageParams](raw, rawReturn, &accountv15.AuthenticateMessageParams{})
	}
	return nil, nil
}
