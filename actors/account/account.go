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
	"github.com/zondax/fil-parser/tools"
)

func PubkeyAddress(network string, raw, rawReturn []byte) (map[string]interface{}, error) {
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

func AuthenticateMessage(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(network, height):
		return nil, fmt.Errorf("not supported")
	case tools.V9.IsSupported(network, height):
		return authenticateMessageGeneric[*accountv9.AuthenticateMessageParams, *accountv9.AuthenticateMessageParams](raw, rawReturn, &accountv9.AuthenticateMessageParams{})
	case tools.V10.IsSupported(network, height):
		return authenticateMessageGeneric[*accountv10.AuthenticateMessageParams, *accountv10.AuthenticateMessageParams](raw, rawReturn, &accountv10.AuthenticateMessageParams{})
	case tools.V11.IsSupported(network, height):
		return authenticateMessageGeneric[*accountv11.AuthenticateMessageParams, *accountv11.AuthenticateMessageParams](raw, rawReturn, &accountv11.AuthenticateMessageParams{})
	case tools.V14.IsSupported(network, height):
		return authenticateMessageGeneric[*accountv14.AuthenticateMessageParams, *accountv14.AuthenticateMessageParams](raw, rawReturn, &accountv14.AuthenticateMessageParams{})
	default:
		return authenticateMessageGeneric[*accountv15.AuthenticateMessageParams, *accountv15.AuthenticateMessageParams](raw, rawReturn, &accountv15.AuthenticateMessageParams{})
	}
}
