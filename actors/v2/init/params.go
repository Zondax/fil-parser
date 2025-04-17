package init

import (
	builtinInitv10 "github.com/filecoin-project/go-state-types/builtin/v10/init"
	builtinInitv11 "github.com/filecoin-project/go-state-types/builtin/v11/init"
	builtinInitv12 "github.com/filecoin-project/go-state-types/builtin/v12/init"
	builtinInitv13 "github.com/filecoin-project/go-state-types/builtin/v13/init"
	builtinInitv14 "github.com/filecoin-project/go-state-types/builtin/v14/init"
	builtinInitv15 "github.com/filecoin-project/go-state-types/builtin/v15/init"
	builtinInitv16 "github.com/filecoin-project/go-state-types/builtin/v16/init"
	builtinInitv8 "github.com/filecoin-project/go-state-types/builtin/v8/init"
	builtinInitv9 "github.com/filecoin-project/go-state-types/builtin/v9/init"
	legacyv1 "github.com/filecoin-project/specs-actors/actors/builtin/init"
	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/init"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/init"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/init"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/init"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/init"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/init"
	typegen "github.com/whyrusleeping/cbor-gen"
	"github.com/zondax/fil-parser/tools"
)

func constructorParams() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.ConstructorParams{},

		tools.V8.String(): &legacyv2.ConstructorParams{},
		tools.V9.String(): &legacyv2.ConstructorParams{},

		tools.V10.String(): &legacyv3.ConstructorParams{},
		tools.V11.String(): &legacyv3.ConstructorParams{},

		tools.V12.String(): &legacyv4.ConstructorParams{},
		tools.V13.String(): &legacyv5.ConstructorParams{},
		tools.V14.String(): &legacyv6.ConstructorParams{},
		tools.V15.String(): &legacyv7.ConstructorParams{},
		tools.V16.String(): &builtinInitv8.ConstructorParams{},
		tools.V17.String(): &builtinInitv9.ConstructorParams{},
		tools.V18.String(): &builtinInitv10.ConstructorParams{},

		tools.V19.String(): &builtinInitv11.ConstructorParams{},
		tools.V20.String(): &builtinInitv11.ConstructorParams{},

		tools.V21.String(): &builtinInitv12.ConstructorParams{},
		tools.V22.String(): &builtinInitv13.ConstructorParams{},
		tools.V23.String(): &builtinInitv14.ConstructorParams{},
		tools.V24.String(): &builtinInitv15.ConstructorParams{},
		tools.V25.String(): &builtinInitv16.ConstructorParams{},
	}
}

func execParams() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.ExecParams{},
		tools.V8.String(): &legacyv2.ExecParams{},
		tools.V9.String(): &legacyv2.ExecParams{},

		tools.V10.String(): &legacyv3.ExecParams{},
		tools.V11.String(): &legacyv3.ExecParams{},

		tools.V12.String(): &legacyv4.ExecParams{},
		tools.V13.String(): &legacyv5.ExecParams{},
		tools.V14.String(): &legacyv6.ExecParams{},
		tools.V15.String(): &legacyv7.ExecParams{},
		tools.V16.String(): &builtinInitv8.ExecParams{},
		tools.V17.String(): &builtinInitv9.ExecParams{},
		tools.V18.String(): &builtinInitv10.ExecParams{},

		tools.V19.String(): &builtinInitv11.ExecParams{},
		tools.V20.String(): &builtinInitv11.ExecParams{},

		tools.V21.String(): &builtinInitv12.ExecParams{},
		tools.V22.String(): &builtinInitv13.ExecParams{},
		tools.V23.String(): &builtinInitv14.ExecParams{},
		tools.V24.String(): &builtinInitv15.ExecParams{},
		tools.V25.String(): &builtinInitv16.ExecParams{},
	}
}

func execReturn() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.ExecReturn{},
		tools.V8.String(): &legacyv2.ExecReturn{},
		tools.V9.String(): &legacyv2.ExecReturn{},

		tools.V10.String(): &legacyv3.ExecReturn{},
		tools.V11.String(): &legacyv3.ExecReturn{},

		tools.V12.String(): &legacyv4.ExecReturn{},
		tools.V13.String(): &legacyv5.ExecReturn{},
		tools.V14.String(): &legacyv6.ExecReturn{},
		tools.V15.String(): &legacyv7.ExecReturn{},
		tools.V16.String(): &builtinInitv8.ExecReturn{},
		tools.V17.String(): &builtinInitv9.ExecReturn{},
		tools.V18.String(): &builtinInitv10.ExecReturn{},

		tools.V19.String(): &builtinInitv11.ExecReturn{},
		tools.V20.String(): &builtinInitv11.ExecReturn{},

		tools.V21.String(): &builtinInitv12.ExecReturn{},
		tools.V22.String(): &builtinInitv13.ExecReturn{},
		tools.V23.String(): &builtinInitv14.ExecReturn{},
		tools.V24.String(): &builtinInitv15.ExecReturn{},
		tools.V25.String(): &builtinInitv16.ExecReturn{},
	}
}

func exec4Params() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V18.String(): &builtinInitv10.Exec4Params{},

		tools.V19.String(): &builtinInitv11.Exec4Params{},
		tools.V20.String(): &builtinInitv11.Exec4Params{},

		tools.V21.String(): &builtinInitv12.Exec4Params{},
		tools.V22.String(): &builtinInitv13.Exec4Params{},
		tools.V23.String(): &builtinInitv14.Exec4Params{},
		tools.V24.String(): &builtinInitv15.Exec4Params{},
		tools.V25.String(): &builtinInitv16.Exec4Params{},
	}
}

func exec4Return() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V18.String(): &builtinInitv10.Exec4Return{},

		tools.V19.String(): &builtinInitv11.Exec4Return{},
		tools.V20.String(): &builtinInitv11.Exec4Return{},

		tools.V21.String(): &builtinInitv12.Exec4Return{},
		tools.V22.String(): &builtinInitv13.Exec4Return{},
		tools.V23.String(): &builtinInitv14.Exec4Return{},
		tools.V24.String(): &builtinInitv15.Exec4Return{},
		tools.V25.String(): &builtinInitv16.Exec4Return{},
	}
}
