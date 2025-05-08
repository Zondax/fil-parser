package init

import (
	"context"
	"fmt"

	"github.com/ipfs/go-cid"
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

func legacyMethods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
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
	tools.V16.String(): actors.CopyMethods(builtinInitv8.Methods),
	tools.V17.String(): actors.CopyMethods(builtinInitv9.Methods),
	tools.V18.String(): actors.CopyMethods(builtinInitv10.Methods),
	tools.V19.String(): actors.CopyMethods(builtinInitv11.Methods),
	tools.V20.String(): actors.CopyMethods(builtinInitv11.Methods),
	tools.V21.String(): actors.CopyMethods(builtinInitv12.Methods),
	tools.V22.String(): actors.CopyMethods(builtinInitv13.Methods),
	tools.V23.String(): actors.CopyMethods(builtinInitv14.Methods),
	tools.V24.String(): actors.CopyMethods(builtinInitv15.Methods),
	tools.V25.String(): actors.CopyMethods(builtinInitv16.Methods),
}

func (i *Init) Methods(_ context.Context, network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error) {
	version := tools.VersionFromHeight(network, height)
	methods, ok := methods[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return methods, nil
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

	metadata, addressInfo, err := parseExec(msg, raw, params(), returnValue(), i.helper)
	if addressInfo != nil {
		c, err := cid.Parse(addressInfo.ActorCid)
		if err == nil {
			createdActorName, err := i.helper.GetFilecoinLib().BuiltinActors.GetActorNameFromCidByVersion(c, version.FilNetworkVersion())
			if err == nil {
				addressInfo.ActorType = parseExecActor(createdActorName)
			}
		}
	}
	if addressInfo != nil {
		i.helper.GetActorsCache().StoreAddressInfoAddress(*addressInfo)
	}
	return metadata, addressInfo, err
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

	metadata, addressInfo, err := parseExec(msg, raw, params(), returnValue(), i.helper)
	if addressInfo != nil {
		c, err := cid.Parse(addressInfo.ActorCid)
		if err == nil {
			createdActorName, err := i.helper.GetFilecoinLib().BuiltinActors.GetActorNameFromCidByVersion(c, version.FilNetworkVersion())
			if err == nil {
				addressInfo.ActorType = parseExecActor(createdActorName)
			}
		}
	}
	if addressInfo != nil {
		i.helper.GetActorsCache().StoreAddressInfoAddress(*addressInfo)
	}
	return metadata, addressInfo, err
}
