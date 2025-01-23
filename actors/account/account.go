package account

import (
	"bytes"
	"encoding/base64"
	"fmt"

	"github.com/filecoin-project/go-address"
	accountv10 "github.com/filecoin-project/go-state-types/builtin/v10/account"
	accountv11 "github.com/filecoin-project/go-state-types/builtin/v11/account"
	accountv9 "github.com/filecoin-project/go-state-types/builtin/v11/account"
	accountv12 "github.com/filecoin-project/go-state-types/builtin/v12/account"
	accountv13 "github.com/filecoin-project/go-state-types/builtin/v13/account"
	accountv14 "github.com/filecoin-project/go-state-types/builtin/v14/account"
	typegen "github.com/whyrusleeping/cbor-gen"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

type Account struct{}

func (a *Account) PubkeyAddress(network string, raw, rawReturn []byte) (map[string]interface{}, error) {
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

func (a *Account) AuthenticateMessage(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	// all versions before V17
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V16)...):
		return map[string]interface{}{}, nil // method did not exist
	case tools.V17.IsSupported(network, height):
		return authenticateMessageGeneric[*accountv9.AuthenticateMessageParams, *typegen.CborBool](raw, rawReturn, &accountv9.AuthenticateMessageParams{})
	case tools.V18.IsSupported(network, height):
		return authenticateMessageGeneric[*accountv10.AuthenticateMessageParams, *typegen.CborBool](raw, rawReturn, &accountv10.AuthenticateMessageParams{})
	case tools.V19.IsSupported(network, height) || tools.V20.IsSupported(network, height):
		return authenticateMessageGeneric[*accountv11.AuthenticateMessageParams, *typegen.CborBool](raw, rawReturn, &accountv11.AuthenticateMessageParams{})
	case tools.V21.IsSupported(network, height):
		return authenticateMessageGeneric[*accountv12.AuthenticateMessageParams, *typegen.CborBool](raw, rawReturn, &accountv12.AuthenticateMessageParams{})
	case tools.V22.IsSupported(network, height):
		return authenticateMessageGeneric[*accountv13.AuthenticateMessageParams, *typegen.CborBool](raw, rawReturn, &accountv13.AuthenticateMessageParams{})
	case tools.V23.IsSupported(network, height):
		return authenticateMessageGeneric[*accountv14.AuthenticateMessageParams, *typegen.CborBool](raw, rawReturn, &accountv14.AuthenticateMessageParams{})
	default:
		return nil, fmt.Errorf("not supported")
	}
}
