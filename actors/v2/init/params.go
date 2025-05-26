package init

import (
	"github.com/filecoin-project/go-state-types/abi"
	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"
	builtinInitv10 "github.com/filecoin-project/go-state-types/builtin/v10/init"
	builtinInitv11 "github.com/filecoin-project/go-state-types/builtin/v11/init"
	builtinInitv12 "github.com/filecoin-project/go-state-types/builtin/v12/init"
	builtinInitv13 "github.com/filecoin-project/go-state-types/builtin/v13/init"
	builtinInitv14 "github.com/filecoin-project/go-state-types/builtin/v14/init"
	builtinInitv15 "github.com/filecoin-project/go-state-types/builtin/v15/init"
	builtinInitv16 "github.com/filecoin-project/go-state-types/builtin/v16/init"
	builtinInitv8 "github.com/filecoin-project/go-state-types/builtin/v8/init"
	builtinInitv9 "github.com/filecoin-project/go-state-types/builtin/v9/init"
	legacyBuiltin "github.com/filecoin-project/specs-actors/actors/builtin"
	legacyv1 "github.com/filecoin-project/specs-actors/actors/builtin/init"
	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/init"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/init"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/init"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/init"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/init"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/init"
	typegen "github.com/whyrusleeping/cbor-gen"
	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

func v1Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	i := &Init{}
	return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
		legacyBuiltin.MethodsInit.Constructor: {
			Name:   parser.MethodConstructor,
			Method: actors.ParseConstructor,
		},
		legacyBuiltin.MethodsInit.Exec: {
			Name:   parser.MethodExec,
			Method: i.Exec,
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

var constructorParams = map[string]func() typegen.CBORUnmarshaler{
	tools.V1.String(): func() typegen.CBORUnmarshaler { return new(legacyv1.ConstructorParams) },
	tools.V2.String(): func() typegen.CBORUnmarshaler { return new(legacyv1.ConstructorParams) },
	tools.V3.String(): func() typegen.CBORUnmarshaler { return new(legacyv1.ConstructorParams) },

	tools.V4.String(): func() typegen.CBORUnmarshaler { return new(legacyv2.ConstructorParams) },
	tools.V5.String(): func() typegen.CBORUnmarshaler { return new(legacyv2.ConstructorParams) },
	tools.V6.String(): func() typegen.CBORUnmarshaler { return new(legacyv2.ConstructorParams) },
	tools.V7.String(): func() typegen.CBORUnmarshaler { return new(legacyv2.ConstructorParams) },
	tools.V8.String(): func() typegen.CBORUnmarshaler { return new(legacyv2.ConstructorParams) },
	tools.V9.String(): func() typegen.CBORUnmarshaler { return new(legacyv2.ConstructorParams) },

	tools.V10.String(): func() typegen.CBORUnmarshaler { return new(legacyv3.ConstructorParams) },
	tools.V11.String(): func() typegen.CBORUnmarshaler { return new(legacyv3.ConstructorParams) },

	tools.V12.String(): func() typegen.CBORUnmarshaler { return new(legacyv4.ConstructorParams) },
	tools.V13.String(): func() typegen.CBORUnmarshaler { return new(legacyv5.ConstructorParams) },
	tools.V14.String(): func() typegen.CBORUnmarshaler { return new(legacyv6.ConstructorParams) },
	tools.V15.String(): func() typegen.CBORUnmarshaler { return new(legacyv7.ConstructorParams) },
	tools.V16.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv8.ConstructorParams) },
	tools.V17.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv9.ConstructorParams) },
	tools.V18.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv10.ConstructorParams) },

	tools.V19.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv11.ConstructorParams) },
	tools.V20.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv11.ConstructorParams) },

	tools.V21.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv12.ConstructorParams) },
	tools.V22.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv13.ConstructorParams) },
	tools.V23.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv14.ConstructorParams) },
	tools.V24.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv15.ConstructorParams) },
	tools.V25.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv16.ConstructorParams) },
}

var execParams = map[string]func() typegen.CBORUnmarshaler{
	tools.V1.String(): func() typegen.CBORUnmarshaler { return new(legacyv1.ExecParams) },
	tools.V2.String(): func() typegen.CBORUnmarshaler { return new(legacyv1.ExecParams) },
	tools.V3.String(): func() typegen.CBORUnmarshaler { return new(legacyv1.ExecParams) },

	tools.V4.String(): func() typegen.CBORUnmarshaler { return new(legacyv2.ExecParams) },
	tools.V5.String(): func() typegen.CBORUnmarshaler { return new(legacyv2.ExecParams) },
	tools.V6.String(): func() typegen.CBORUnmarshaler { return new(legacyv2.ExecParams) },
	tools.V7.String(): func() typegen.CBORUnmarshaler { return new(legacyv2.ExecParams) },
	tools.V8.String(): func() typegen.CBORUnmarshaler { return new(legacyv2.ExecParams) },
	tools.V9.String(): func() typegen.CBORUnmarshaler { return new(legacyv2.ExecParams) },

	tools.V10.String(): func() typegen.CBORUnmarshaler { return new(legacyv3.ExecParams) },
	tools.V11.String(): func() typegen.CBORUnmarshaler { return new(legacyv3.ExecParams) },

	tools.V12.String(): func() typegen.CBORUnmarshaler { return new(legacyv4.ExecParams) },
	tools.V13.String(): func() typegen.CBORUnmarshaler { return new(legacyv5.ExecParams) },
	tools.V14.String(): func() typegen.CBORUnmarshaler { return new(legacyv6.ExecParams) },
	tools.V15.String(): func() typegen.CBORUnmarshaler { return new(legacyv7.ExecParams) },
	tools.V16.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv8.ExecParams) },
	tools.V17.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv9.ExecParams) },
	tools.V18.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv10.ExecParams) },

	tools.V19.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv11.ExecParams) },
	tools.V20.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv11.ExecParams) },

	tools.V21.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv12.ExecParams) },
	tools.V22.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv13.ExecParams) },
	tools.V23.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv14.ExecParams) },
	tools.V24.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv15.ExecParams) },
	tools.V25.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv16.ExecParams) },
}

var execReturn = map[string]func() typegen.CBORUnmarshaler{
	tools.V1.String(): func() typegen.CBORUnmarshaler { return new(legacyv1.ExecReturn) },
	tools.V2.String(): func() typegen.CBORUnmarshaler { return new(legacyv1.ExecReturn) },
	tools.V3.String(): func() typegen.CBORUnmarshaler { return new(legacyv1.ExecReturn) },

	tools.V4.String(): func() typegen.CBORUnmarshaler { return new(legacyv2.ExecReturn) },
	tools.V5.String(): func() typegen.CBORUnmarshaler { return new(legacyv2.ExecReturn) },
	tools.V6.String(): func() typegen.CBORUnmarshaler { return new(legacyv2.ExecReturn) },
	tools.V7.String(): func() typegen.CBORUnmarshaler { return new(legacyv2.ExecReturn) },
	tools.V8.String(): func() typegen.CBORUnmarshaler { return new(legacyv2.ExecReturn) },
	tools.V9.String(): func() typegen.CBORUnmarshaler { return new(legacyv2.ExecReturn) },

	tools.V10.String(): func() typegen.CBORUnmarshaler { return new(legacyv3.ExecReturn) },
	tools.V11.String(): func() typegen.CBORUnmarshaler { return new(legacyv3.ExecReturn) },

	tools.V12.String(): func() typegen.CBORUnmarshaler { return new(legacyv4.ExecReturn) },
	tools.V13.String(): func() typegen.CBORUnmarshaler { return new(legacyv5.ExecReturn) },
	tools.V14.String(): func() typegen.CBORUnmarshaler { return new(legacyv6.ExecReturn) },
	tools.V15.String(): func() typegen.CBORUnmarshaler { return new(legacyv7.ExecReturn) },
	tools.V16.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv8.ExecReturn) },
	tools.V17.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv9.ExecReturn) },
	tools.V18.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv10.ExecReturn) },

	tools.V19.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv11.ExecReturn) },
	tools.V20.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv11.ExecReturn) },

	tools.V21.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv12.ExecReturn) },
	tools.V22.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv13.ExecReturn) },
	tools.V23.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv14.ExecReturn) },
	tools.V24.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv15.ExecReturn) },
	tools.V25.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv16.ExecReturn) },
}

var exec4Params = map[string]func() typegen.CBORUnmarshaler{
	tools.V18.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv10.Exec4Params) },

	tools.V19.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv11.Exec4Params) },
	tools.V20.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv11.Exec4Params) },

	tools.V21.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv12.Exec4Params) },
	tools.V22.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv13.Exec4Params) },
	tools.V23.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv14.Exec4Params) },
	tools.V24.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv15.Exec4Params) },
	tools.V25.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv16.Exec4Params) },
}

var exec4Return = map[string]func() typegen.CBORUnmarshaler{
	tools.V18.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv10.Exec4Return) },

	tools.V19.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv11.Exec4Return) },
	tools.V20.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv11.Exec4Return) },

	tools.V21.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv12.Exec4Return) },
	tools.V22.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv13.Exec4Return) },
	tools.V23.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv14.Exec4Return) },
	tools.V24.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv15.Exec4Return) },
	tools.V25.String(): func() typegen.CBORUnmarshaler { return new(builtinInitv16.Exec4Return) },
}
