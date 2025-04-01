package init

import (
	"context"
	"fmt"
	"github.com/zondax/fil-parser/parser/helper"

	"go.uber.org/zap"

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
	builtinInitv8 "github.com/filecoin-project/go-state-types/builtin/v8/init"
	builtinInitv9 "github.com/filecoin-project/go-state-types/builtin/v9/init"

	legacyv1 "github.com/filecoin-project/specs-actors/actors/builtin/init"
	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/init"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/init"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/init"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/init"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/init"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/init"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
)

type Init struct {
	helper *helper.Helper
	logger *zap.Logger
}

func New(helper *helper.Helper, logger *zap.Logger) *Init {
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

func (i *Init) Methods(_ context.Context, network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error) {
	switch {
	// all legacy version
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V15)...):
		return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
			legacyBuiltin.MethodsInit.Constructor: {
				Name: parser.MethodConstructor,
			},
			legacyBuiltin.MethodsInit.Exec: {
				Name: parser.MethodExec,
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
	default:
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
}

func (*Init) Constructor(network string, height int64, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return initConstructor(raw, &builtinInitv15.ConstructorParams{})
	case tools.V23.IsSupported(network, height):
		return initConstructor(raw, &builtinInitv14.ConstructorParams{})
	case tools.V22.IsSupported(network, height):
		return initConstructor(raw, &builtinInitv13.ConstructorParams{})
	case tools.V21.IsSupported(network, height):
		return initConstructor(raw, &builtinInitv12.ConstructorParams{})
	case tools.AnyIsSupported(network, height, tools.V20, tools.V19):
		return initConstructor(raw, &builtinInitv11.ConstructorParams{})
	case tools.V18.IsSupported(network, height):
		return initConstructor(raw, &builtinInitv10.ConstructorParams{})
	case tools.V17.IsSupported(network, height):
		return initConstructor(raw, &builtinInitv9.ConstructorParams{})
	case tools.V16.IsSupported(network, height):
		return initConstructor(raw, &builtinInitv8.ConstructorParams{})
	case tools.V15.IsSupported(network, height):
		return initConstructor(raw, &legacyv7.ConstructorParams{})
	case tools.V14.IsSupported(network, height):
		return initConstructor(raw, &legacyv6.ConstructorParams{})
	case tools.V13.IsSupported(network, height):
		return initConstructor(raw, &legacyv5.ConstructorParams{})
	case tools.V12.IsSupported(network, height):
		return initConstructor(raw, &legacyv4.ConstructorParams{})
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return initConstructor(raw, &legacyv3.ConstructorParams{})
	case tools.AnyIsSupported(network, height, tools.V8, tools.V9):
		return initConstructor(raw, &legacyv2.ConstructorParams{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return initConstructor(raw, &legacyv1.ConstructorParams{})
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (i *Init) Exec(network string, height int64, msg *parser.LotusMessage, raw []byte) (map[string]interface{}, *types.AddressInfo, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseExec(msg, raw, &builtinInitv15.ExecParams{}, &builtinInitv15.ExecReturn{}, i.helper)
	case tools.V23.IsSupported(network, height):
		return parseExec(msg, raw, &builtinInitv14.ExecParams{}, &builtinInitv14.ExecReturn{}, i.helper)
	case tools.V22.IsSupported(network, height):
		return parseExec(msg, raw, &builtinInitv13.ExecParams{}, &builtinInitv13.ExecReturn{}, i.helper)
	case tools.V21.IsSupported(network, height):
		return parseExec(msg, raw, &builtinInitv12.ExecParams{}, &builtinInitv12.ExecReturn{}, i.helper)
	case tools.AnyIsSupported(network, height, tools.V20, tools.V19):
		return parseExec(msg, raw, &builtinInitv11.ExecParams{}, &builtinInitv11.ExecReturn{}, i.helper)
	case tools.V18.IsSupported(network, height):
		return parseExec(msg, raw, &builtinInitv10.ExecParams{}, &builtinInitv10.ExecReturn{}, i.helper)
	case tools.V17.IsSupported(network, height):
		return parseExec(msg, raw, &builtinInitv9.ExecParams{}, &builtinInitv9.ExecReturn{}, i.helper)
	case tools.V16.IsSupported(network, height):
		return parseExec(msg, raw, &builtinInitv8.ExecParams{}, &builtinInitv8.ExecReturn{}, i.helper)
	case tools.V15.IsSupported(network, height):
		return parseExec(msg, raw, &legacyv7.ExecParams{}, &legacyv7.ExecReturn{}, i.helper)
	case tools.V14.IsSupported(network, height):
		return parseExec(msg, raw, &legacyv6.ExecParams{}, &legacyv6.ExecReturn{}, i.helper)
	case tools.V13.IsSupported(network, height):
		return parseExec(msg, raw, &legacyv5.ExecParams{}, &legacyv5.ExecReturn{}, i.helper)
	case tools.V12.IsSupported(network, height):
		return parseExec(msg, raw, &legacyv4.ExecParams{}, &legacyv4.ExecReturn{}, i.helper)
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parseExec(msg, raw, &legacyv3.ExecParams{}, &legacyv3.ExecReturn{}, i.helper)
	case tools.AnyIsSupported(network, height, tools.V8, tools.V9):
		return parseExec(msg, raw, &legacyv2.ExecParams{}, &legacyv2.ExecReturn{}, i.helper)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return parseExec(msg, raw, &legacyv1.ExecParams{}, &legacyv1.ExecReturn{}, i.helper)
	}
	return nil, nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (i *Init) Exec4(network string, height int64, msg *parser.LotusMessage, raw []byte) (map[string]interface{}, *types.AddressInfo, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseExec(msg, raw, &builtinInitv15.Exec4Params{}, &builtinInitv15.Exec4Return{}, i.helper)
	case tools.V23.IsSupported(network, height):
		return parseExec(msg, raw, &builtinInitv14.Exec4Params{}, &builtinInitv14.Exec4Return{}, i.helper)
	case tools.V22.IsSupported(network, height):
		return parseExec(msg, raw, &builtinInitv13.Exec4Params{}, &builtinInitv13.Exec4Return{}, i.helper)
	case tools.V21.IsSupported(network, height):
		return parseExec(msg, raw, &builtinInitv12.Exec4Params{}, &builtinInitv12.Exec4Return{}, i.helper)
	case tools.AnyIsSupported(network, height, tools.V20, tools.V19):
		return parseExec(msg, raw, &builtinInitv11.Exec4Params{}, &builtinInitv11.Exec4Return{}, i.helper)
	case tools.V18.IsSupported(network, height):
		return parseExec(msg, raw, &builtinInitv10.Exec4Params{}, &builtinInitv10.Exec4Return{}, i.helper)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V17)...):
		return map[string]interface{}{}, nil, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}
