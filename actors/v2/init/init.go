package init

import (
	"context"
	"fmt"

	"github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/golem/pkg/logger"

	"github.com/filecoin-project/go-state-types/abi"
	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/manifest"
	legacyBuiltin "github.com/filecoin-project/specs-actors/actors/builtin"

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

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
)

type Init struct {
	helper *helper.Helper
	logger *logger.Logger
}

func New(helper *helper.Helper, logger *logger.Logger) *Init {
	return &Init{
		helper: helper,
		logger: logger,
	}
}

func (i *Init) Name() string {
	return manifest.InitKey
}

func (*Init) StartNetworkHeight() int64 {
	return tools.V1.Height()
}

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

func (i *Init) Methods(_ context.Context, network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error) {
	switch {
	// all legacy version
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V15)...):
		return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
			legacyBuiltin.MethodsInit.Constructor: {
				Name:   parser.MethodConstructor,
				Method: actors.ParseConstructor,
			},
			legacyBuiltin.MethodsInit.Exec: {
				Name:   parser.MethodExec,
				Method: i.Exec,
			},
		}, nil
	case tools.V16.IsSupported(network, height):
		return builtinInitv8.Methods, nil
	case tools.V17.IsSupported(network, height):
		return builtinInitv9.Methods, nil
	case tools.V18.IsSupported(network, height):
		return builtinInitv10.Methods, nil
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return builtinInitv11.Methods, nil
	case tools.V21.IsSupported(network, height):
		return builtinInitv12.Methods, nil
	case tools.V22.IsSupported(network, height):
		return builtinInitv13.Methods, nil
	case tools.V23.IsSupported(network, height):
		return builtinInitv14.Methods, nil
	case tools.V24.IsSupported(network, height):
		return builtinInitv15.Methods, nil
	case tools.V25.IsSupported(network, height):
		return builtinInitv16.Methods, nil
	default:
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
}

func (*Init) Constructor(network string, height int64, raw []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := constructorParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return initConstructor(raw, params)
}

func (i *Init) Exec(network string, height int64, msg *parser.LotusMessage, raw []byte) (map[string]interface{}, *types.AddressInfo, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := execParams()[version.String()]
	if !ok {
		return nil, nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := execReturn()[version.String()]
	if !ok {
		return nil, nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseExec(msg, raw, params, returnValue, i.helper)
}

func (i *Init) Exec4(network string, height int64, msg *parser.LotusMessage, raw []byte) (map[string]interface{}, *types.AddressInfo, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := exec4Params()[version.String()]
	if !ok {
		return nil, nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := exec4Return()[version.String()]
	if !ok {
		return nil, nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseExec(msg, raw, params, returnValue, i.helper)
}
