package account

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/hex"
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
	accountv16 "github.com/filecoin-project/go-state-types/builtin/v16/account"
	accountv8 "github.com/filecoin-project/go-state-types/builtin/v8/account"
	accountv9 "github.com/filecoin-project/go-state-types/builtin/v9/account"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

func legacyMethods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	a := &Account{}
	return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
		legacyBuiltin.MethodsAccount.Constructor: {
			Name:   parser.MethodConstructor,
			Method: actors.ParseConstructor,
		},
		legacyBuiltin.MethodsAccount.PubkeyAddress: {
			Name:   parser.MethodPubkeyAddress,
			Method: a.PubkeyAddress,
		},
	}
}

var methods = map[string]map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
	tools.V1.String():  legacyMethods(),
	tools.V2.String():  legacyMethods(),
	tools.V3.String():  legacyMethods(),
	tools.V4.String():  legacyMethods(),
	tools.V5.String():  legacyMethods(),
	tools.V6.String():  legacyMethods(),
	tools.V7.String():  legacyMethods(),
	tools.V8.String():  legacyMethods(),
	tools.V9.String():  legacyMethods(),
	tools.V10.String(): legacyMethods(),
	tools.V11.String(): legacyMethods(),
	tools.V12.String(): legacyMethods(),
	tools.V13.String(): legacyMethods(),
	tools.V14.String(): legacyMethods(),
	tools.V15.String(): legacyMethods(),
	tools.V16.String(): actors.CopyMethods(accountv8.Methods),
	tools.V17.String(): actors.CopyMethods(accountv9.Methods),
	tools.V18.String(): actors.CopyMethods(accountv10.Methods),
	tools.V19.String(): actors.CopyMethods(accountv11.Methods),
	tools.V20.String(): actors.CopyMethods(accountv11.Methods),
	tools.V21.String(): actors.CopyMethods(accountv12.Methods),
	tools.V22.String(): actors.CopyMethods(accountv13.Methods),
	tools.V23.String(): actors.CopyMethods(accountv14.Methods),
	tools.V24.String(): actors.CopyMethods(accountv15.Methods),
	tools.V25.String(): actors.CopyMethods(accountv16.Methods),
}

func (a *Account) Methods(_ context.Context, network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error) {
	version := tools.VersionFromHeight(network, height)
	methods, ok := methods[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return methods, nil
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
	version := tools.VersionFromHeight(network, height)
	params, ok := authenticateMessageParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	var r typegen.CborBool
	return authenticateMessageGeneric(raw, rawReturn, params(), &r)
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

func (a *Account) Fallback(network string, height int64, raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	metadata[parser.ParamsRawKey] = hex.EncodeToString(raw)
	return metadata, nil
}
