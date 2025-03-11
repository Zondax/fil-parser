package account

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"

	typegen "github.com/whyrusleeping/cbor-gen"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"
	legacyBuiltin "github.com/filecoin-project/specs-actors/actors/builtin"

	accountv10 "github.com/filecoin-project/go-state-types/builtin/v10/account"
	accountv11 "github.com/filecoin-project/go-state-types/builtin/v11/account"
	accountv12 "github.com/filecoin-project/go-state-types/builtin/v12/account"
	accountv13 "github.com/filecoin-project/go-state-types/builtin/v13/account"
	accountv14 "github.com/filecoin-project/go-state-types/builtin/v14/account"
	accountv15 "github.com/filecoin-project/go-state-types/builtin/v15/account"
	accountv8 "github.com/filecoin-project/go-state-types/builtin/v8/account"
	accountv9 "github.com/filecoin-project/go-state-types/builtin/v9/account"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

func (a *Account) Methods(_ context.Context, network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error) {
	switch {
	// all legacy version
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V15)...):
		return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
			legacyBuiltin.MethodsAccount.Constructor: {
				Name: parser.MethodConstructor,
			},
			legacyBuiltin.MethodsAccount.PubkeyAddress: {
				Name: parser.MethodPubkeyAddress,
			},
		}, nil
	case tools.V16.IsSupported(network, height):
		return accountv8.Methods, nil
	case tools.V17.IsSupported(network, height):
		return accountv9.Methods, nil
	case tools.V18.IsSupported(network, height):
		return accountv10.Methods, nil
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return accountv11.Methods, nil
	case tools.V21.IsSupported(network, height):
		return accountv12.Methods, nil
	case tools.V22.IsSupported(network, height):
		return accountv13.Methods, nil
	case tools.V23.IsSupported(network, height):
		return accountv14.Methods, nil
	case tools.V24.IsSupported(network, height):
		return accountv15.Methods, nil
	default:
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
}

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
	var r typegen.CborBool
	switch {
	// all versions before V17
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V16)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height) // method did not exist
	case tools.V17.IsSupported(network, height):
		return authenticateMessageGeneric(raw, rawReturn, &accountv9.AuthenticateMessageParams{}, &r)
	case tools.V18.IsSupported(network, height):
		return authenticateMessageGeneric(raw, rawReturn, &accountv10.AuthenticateMessageParams{}, &r)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return authenticateMessageGeneric(raw, rawReturn, &accountv11.AuthenticateMessageParams{}, &r)
	case tools.V21.IsSupported(network, height):
		return authenticateMessageGeneric(raw, rawReturn, &accountv12.AuthenticateMessageParams{}, &r)
	case tools.V22.IsSupported(network, height):
		return authenticateMessageGeneric(raw, rawReturn, &accountv13.AuthenticateMessageParams{}, &r)
	case tools.V23.IsSupported(network, height):
		return authenticateMessageGeneric(raw, rawReturn, &accountv14.AuthenticateMessageParams{}, &r)
	case tools.V24.IsSupported(network, height):
		return authenticateMessageGeneric(raw, rawReturn, &accountv15.AuthenticateMessageParams{}, &r)
	default:
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
}

func (a *Account) UniversalReceiverHook(network string, height int64, raw []byte) (map[string]interface{}, error) {
	var data abi.CborBytesTransparent
	if err := data.UnmarshalCBOR(bytes.NewReader(raw)); err != nil {
		return nil, err
	}

	metadata := make(map[string]interface{})
	metadata[parser.ParamsKey] = base64.StdEncoding.EncodeToString(raw)
	return metadata, nil
}
