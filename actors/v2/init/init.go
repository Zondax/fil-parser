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
	params, ok := constructorParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return initConstructor(raw, params())
}

func (i *Init) Exec(network string, height int64, msg *parser.LotusMessage, raw []byte) (map[string]interface{}, *types.AddressInfo, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := execParams[version.String()]
	if !ok {
		return nil, nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := execReturn[version.String()]
	if !ok {
		return nil, nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseExec(msg, raw, params(), returnValue(), i.helper)
}

func (i *Init) Exec4(network string, height int64, msg *parser.LotusMessage, raw []byte) (map[string]interface{}, *types.AddressInfo, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := exec4Params[version.String()]
	if !ok {
		return nil, nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := exec4Return[version.String()]
	if !ok {
		return nil, nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseExec(msg, raw, params(), returnValue(), i.helper)
}
