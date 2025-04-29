package power

import (
	"context"
	"fmt"

	"github.com/zondax/golem/pkg/logger"

	"github.com/filecoin-project/go-state-types/abi"
	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/filecoin-project/go-state-types/proof"
	legacyBuiltin "github.com/filecoin-project/specs-actors/actors/builtin"

	powerv10 "github.com/filecoin-project/go-state-types/builtin/v10/power"
	powerv11 "github.com/filecoin-project/go-state-types/builtin/v11/power"
	powerv12 "github.com/filecoin-project/go-state-types/builtin/v12/power"
	powerv13 "github.com/filecoin-project/go-state-types/builtin/v13/power"
	powerv14 "github.com/filecoin-project/go-state-types/builtin/v14/power"
	powerv15 "github.com/filecoin-project/go-state-types/builtin/v15/power"
	powerv16 "github.com/filecoin-project/go-state-types/builtin/v16/power"
	powerv8 "github.com/filecoin-project/go-state-types/builtin/v8/power"
	powerv9 "github.com/filecoin-project/go-state-types/builtin/v9/power"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
)

type Power struct {
	helper *helper.Helper
	logger *logger.Logger
}

func New(helper *helper.Helper, logger *logger.Logger) *Power {
	return &Power{
		helper: helper,
		logger: logger,
	}
}
func (p *Power) Name() string {
	return manifest.PowerKey
}

func (*Power) StartNetworkHeight() int64 {
	return tools.V1.Height()
}

func legacyMethods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	p := &Power{}
	return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
		legacyBuiltin.MethodsPower.Constructor: {
			Name:   parser.MethodConstructor,
			Method: actors.ParseConstructor,
		},
		legacyBuiltin.MethodsPower.CreateMiner: {
			Name:   parser.MethodCreateMiner,
			Method: p.CreateMinerExported,
		},
		legacyBuiltin.MethodsPower.UpdateClaimedPower: {
			Name:   parser.MethodUpdateClaimedPower,
			Method: p.UpdateClaimedPower,
		},
		legacyBuiltin.MethodsPower.EnrollCronEvent: {
			Name:   parser.MethodEnrollCronEvent,
			Method: p.EnrollCronEvent,
		},
		legacyBuiltin.MethodsPower.OnEpochTickEnd: {
			Name:   parser.MethodOnEpochTickEnd,
			Method: actors.ParseEmptyParamsAndReturn,
		},
		legacyBuiltin.MethodsPower.UpdatePledgeTotal: {
			Name:   parser.MethodUpdatePledgeTotal,
			Method: p.UpdatePledgeTotal,
		},
		legacyBuiltin.MethodsPower.OnConsensusFault: {
			Name:   parser.MethodOnConsensusFault,
			Method: p.OnConsensusFault,
		},
		legacyBuiltin.MethodsPower.SubmitPoRepForBulkVerify: {
			Name:   parser.MethodSubmitPoRepForBulkVerify,
			Method: p.SubmitPoRepForBulkVerify,
		},
		legacyBuiltin.MethodsPower.CurrentTotalPower: {
			Name:   parser.MethodCurrentTotalPower,
			Method: p.CurrentTotalPower,
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
	tools.V16.String(): actors.CopyMethods(powerv8.Methods),
	tools.V17.String(): actors.CopyMethods(powerv9.Methods),
	tools.V18.String(): actors.CopyMethods(powerv10.Methods),
	tools.V19.String(): actors.CopyMethods(powerv11.Methods),
	tools.V20.String(): actors.CopyMethods(powerv11.Methods),
	tools.V21.String(): actors.CopyMethods(powerv12.Methods),
	tools.V22.String(): actors.CopyMethods(powerv13.Methods),
	tools.V23.String(): actors.CopyMethods(powerv14.Methods),
	tools.V24.String(): actors.CopyMethods(powerv15.Methods),
	tools.V25.String(): actors.CopyMethods(powerv16.Methods),
}

func (p *Power) Methods(_ context.Context, network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error) {
	version := tools.VersionFromHeight(network, height)
	methods, ok := methods[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return methods, nil
}

func (*Power) CurrentTotalPower(network string, msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	returnValue, ok := currentTotalPowerReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parse(rawReturn, nil, false, returnValue(), &abi.EmptyValue{}, parser.ReturnKey)
}

func (*Power) SubmitPoRepForBulkVerify(network string, msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	data, err := parse(raw, rawReturn, false, &proof.SealVerifyInfo{}, &proof.SealVerifyInfo{}, parser.ParamsKey)
	return data, err
}

func (*Power) OnConsensusFault(network string, height int64, msg *parser.LotusMessage, raw []byte) (map[string]interface{}, error) {
	var data map[string]interface{}
	var err error
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V15)...):
		return nil, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	case tools.AnyIsSupported(network, height, tools.VersionsAfter(tools.V16)...):
		return map[string]interface{}{}, nil
	default:
		err = fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return data, err
}

func (*Power) Constructor(network string, height int64, msg *parser.LotusMessage, raw []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := constructorParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parse(raw, nil, false, params(), &abi.EmptyValue{}, parser.ParamsKey)
}

func (p *Power) CreateMinerExported(network string, msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, *types.AddressInfo, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := createMinerParams[version.String()]
	if !ok {
		return nil, nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	returnValue, ok := createMinerReturn[version.String()]
	if !ok {
		return nil, nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	data, addressInfo, err := parseCreateMiner(msg, raw, rawReturn, params(), returnValue())
	if err != nil {
		return nil, nil, err
	}

	if addressInfo != nil {
		p.helper.GetActorsCache().StoreAddressInfoAddress(*addressInfo)
	}

	return data, addressInfo, err
}

func (*Power) EnrollCronEvent(network string, msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := enrollCronEventParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parse(raw, rawReturn, false, params(), &abi.EmptyValue{}, parser.ParamsKey)
}

func (*Power) UpdateClaimedPower(network string, msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := updateClaimedPowerParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parse(raw, rawReturn, false, params(), &abi.EmptyValue{}, parser.ParamsKey)
}

func (*Power) UpdatePledgeTotal(network string, msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	data, err := parse(raw, rawReturn, false, &abi.TokenAmount{}, &abi.TokenAmount{}, parser.ParamsKey)
	return data, err
}

func (*Power) NetworkRawPowerExported(network string, msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	returnValue, ok := networkRawPowerReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parse(rawReturn, nil, false, returnValue(), &abi.EmptyValue{}, parser.ReturnKey)
}

func (*Power) MinerRawPowerExported(network string, msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := minerRawPowerParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	returnValue, ok := minerRawPowerReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parse(raw, rawReturn, true, params(), returnValue(), parser.ParamsKey)
}

func (*Power) MinerCountExported(network string, msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	returnValue, ok := minerCountReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parse(rawReturn, nil, false, returnValue(), &abi.EmptyValue{}, parser.ReturnKey)
}

func (*Power) MinerConsensusCountExported(network string, msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	returnValue, ok := minerConsensusCountReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parse(rawReturn, nil, false, returnValue(), &abi.EmptyValue{}, parser.ReturnKey)
}
