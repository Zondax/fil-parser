package init

import (
	"context"
	"fmt"

	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/golem/pkg/logger"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/filecoin-project/go-state-types/network"
	filTypes "github.com/filecoin-project/lotus/chain/types"

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

var methods = map[string]map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
	tools.V1.String(): v1Methods(),
	tools.V2.String(): v1Methods(),
	tools.V3.String(): v1Methods(),

	tools.V4.String(): v2Methods(),
	tools.V5.String(): v2Methods(),
	tools.V6.String(): v2Methods(),
	tools.V7.String(): v2Methods(),
	tools.V8.String(): v2Methods(),
	tools.V9.String(): v2Methods(),

	tools.V10.String(): v3Methods(),
	tools.V11.String(): v3Methods(),

	tools.V12.String(): v4Methods(),
	tools.V13.String(): v5Methods(),
	tools.V14.String(): v6Methods(),
	tools.V15.String(): v7Methods(),
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
		createdActorCid, createdActorName, err := i.getActorDetailsFromAddress(height, version.FilNetworkVersion(), addressInfo)
		if err == nil {
			addressInfo.ActorCid = createdActorCid.String()
			addressInfo.ActorType = parseExecActor(createdActorName)
		}
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
		createdActorCid, createdActorName, err := i.getActorDetailsFromAddress(height, version.FilNetworkVersion(), addressInfo)
		if err == nil {
			addressInfo.ActorCid = createdActorCid.String()
			addressInfo.ActorType = parseExecActor(createdActorName)
		}
	}

	return metadata, addressInfo, err
}

func (i *Init) getActorDetailsFromAddress(height int64, version network.Version, addressInfo *types.AddressInfo) (actorCid cid.Cid, actorName string, err error) {
	parsedActorCid, err := cid.Parse(addressInfo.ActorCid)
	if err != nil {
		return cid.Undef, "", err
	}
	addrStr := addressInfo.Robust
	if addrStr == "" {
		addrStr = addressInfo.Short
	}
	addr, err := address.NewFromString(addrStr)
	if err != nil {
		return cid.Undef, "", err
	}

	parsedActorName, err := i.helper.GetFilecoinLib().BuiltinActors.GetActorNameFromCidByVersion(parsedActorCid, version)
	if err != nil {
		i.logger.Warnf("initActor: error getting actor details from rosetta: %s", err)
		gotActorCid, gotActorName, err := i.helper.GetActorNameFromAddress(addr, height, filTypes.EmptyTSK)
		if err != nil {
			i.logger.Errorf("initActor: error getting actor details from node: %s", err)
			return cid.Undef, parsedActorName, err
		}
		return gotActorCid, gotActorName, nil
	}

	return parsedActorCid, parsedActorName, nil
}
