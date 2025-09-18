package account

import (
	"github.com/filecoin-project/go-state-types/abi"
	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"
	accountv10 "github.com/filecoin-project/go-state-types/builtin/v10/account"
	accountv11 "github.com/filecoin-project/go-state-types/builtin/v11/account"
	accountv12 "github.com/filecoin-project/go-state-types/builtin/v12/account"
	accountv13 "github.com/filecoin-project/go-state-types/builtin/v13/account"
	accountv14 "github.com/filecoin-project/go-state-types/builtin/v14/account"
	accountv15 "github.com/filecoin-project/go-state-types/builtin/v15/account"
	accountv16 "github.com/filecoin-project/go-state-types/builtin/v16/account"
	accountv17 "github.com/filecoin-project/go-state-types/builtin/v17/account"
	accountv9 "github.com/filecoin-project/go-state-types/builtin/v9/account"
	legacyBuiltin "github.com/filecoin-project/specs-actors/actors/builtin"
	typegen "github.com/whyrusleeping/cbor-gen"
	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

// All methods can be found in the Actor.Exports method in
// the correct version package for "github.com/filecoin-project/specs-actors/actors/builtin/account"

func v1Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
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
func v2Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return v1Methods()
}
func v3Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return v1Methods()
}
func v4Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return v1Methods()
}
func v5Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return v1Methods()
}
func v6Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return v1Methods()
}
func v7Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return v1Methods()
}

var authenticateMessageParams = map[string]func() typegen.CBORUnmarshaler{
	tools.V17.String(): func() typegen.CBORUnmarshaler { return new(accountv9.AuthenticateMessageParams) },
	tools.V18.String(): func() typegen.CBORUnmarshaler { return new(accountv10.AuthenticateMessageParams) },
	tools.V19.String(): func() typegen.CBORUnmarshaler { return new(accountv11.AuthenticateMessageParams) },
	tools.V20.String(): func() typegen.CBORUnmarshaler { return new(accountv11.AuthenticateMessageParams) },
	tools.V21.String(): func() typegen.CBORUnmarshaler { return new(accountv12.AuthenticateMessageParams) },
	tools.V22.String(): func() typegen.CBORUnmarshaler { return new(accountv13.AuthenticateMessageParams) },
	tools.V23.String(): func() typegen.CBORUnmarshaler { return new(accountv14.AuthenticateMessageParams) },
	tools.V24.String(): func() typegen.CBORUnmarshaler { return new(accountv15.AuthenticateMessageParams) },
	tools.V25.String(): func() typegen.CBORUnmarshaler { return new(accountv16.AuthenticateMessageParams) },
	tools.V26.String(): func() typegen.CBORUnmarshaler { return new(accountv16.AuthenticateMessageParams) },
	tools.V27.String(): func() typegen.CBORUnmarshaler { return new(accountv17.AuthenticateMessageParams) },
}
