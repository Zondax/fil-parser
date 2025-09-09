package power

import (
	"context"
	"fmt"

	"github.com/zondax/golem/pkg/logger"

	"github.com/filecoin-project/go-state-types/abi"
	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/exitcode"
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/filecoin-project/go-state-types/proof"

	powerv10 "github.com/filecoin-project/go-state-types/builtin/v10/power"
	powerv11 "github.com/filecoin-project/go-state-types/builtin/v11/power"
	powerv12 "github.com/filecoin-project/go-state-types/builtin/v12/power"
	powerv13 "github.com/filecoin-project/go-state-types/builtin/v13/power"
	powerv14 "github.com/filecoin-project/go-state-types/builtin/v14/power"
	powerv15 "github.com/filecoin-project/go-state-types/builtin/v15/power"
	powerv16 "github.com/filecoin-project/go-state-types/builtin/v16/power"
	powerv17 "github.com/filecoin-project/go-state-types/builtin/v17/power"
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

var methods = map[string]map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
	tools.V0.String(): v1Methods(),
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
	tools.V27.String(): actors.CopyMethods(powerv17.Methods),
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
	metadata, err := parse(rawReturn, nil, false, returnValue(), &abi.EmptyValue{}, parser.ReturnKey)
	if err != nil {
		versions := tools.GetSupportedVersions(network)
		for _, v := range versions {
			returnValue, ok := currentTotalPowerReturn[v.String()]
			if !ok {
				continue
			}
			metadata, err = parse(rawReturn, nil, false, returnValue(), &abi.EmptyValue{}, parser.ReturnKey)
			if err != nil {
				continue
			}
			break
		}
	}
	return metadata, err
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

func (p *Power) CreateMinerExported(network string, msg *parser.LotusMessage, height int64, raw, rawReturn []byte, ec exitcode.ExitCode) (map[string]interface{}, *types.AddressInfo, error) {
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

	if ec.IsSuccess() && addressInfo != nil {
		p.helper.GetActorsCache().StoreAddressInfo(*addressInfo)
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
