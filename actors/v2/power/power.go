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
	powerv8 "github.com/filecoin-project/go-state-types/builtin/v8/power"
	powerv9 "github.com/filecoin-project/go-state-types/builtin/v9/power"

	legacyv1 "github.com/filecoin-project/specs-actors/actors/builtin/power"
	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/power"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/power"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/power"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/power"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/power"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/power"

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

func (*Power) Methods(_ context.Context, network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error) {
	switch {
	// all legacy version
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V15)...):
		return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
			legacyBuiltin.MethodsPower.Constructor: {
				Name: parser.MethodConstructor,
			},
			legacyBuiltin.MethodsPower.CreateMiner: {
				Name: parser.MethodCreateMiner,
			},
			legacyBuiltin.MethodsPower.UpdateClaimedPower: {
				Name: parser.MethodUpdateClaimedPower,
			},
			legacyBuiltin.MethodsPower.EnrollCronEvent: {
				Name: parser.MethodEnrollCronEvent,
			},
			legacyBuiltin.MethodsPower.OnEpochTickEnd: {
				Name: parser.MethodOnEpochTickEnd,
			},
			legacyBuiltin.MethodsPower.UpdatePledgeTotal: {
				Name: parser.MethodUpdatePledgeTotal,
			},
			legacyBuiltin.MethodsPower.OnConsensusFault: {
				Name: parser.MethodOnConsensusFault,
			},
			legacyBuiltin.MethodsPower.SubmitPoRepForBulkVerify: {
				Name: parser.MethodSubmitPoRepForBulkVerify,
			},
			legacyBuiltin.MethodsPower.CurrentTotalPower: {
				Name: parser.MethodCurrentTotalPower,
			},
		}, nil
	case tools.V16.IsSupported(network, height):
		return powerv8.Methods, nil
	case tools.V17.IsSupported(network, height):
		return powerv9.Methods, nil
	case tools.V18.IsSupported(network, height):
		return powerv10.Methods, nil
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return powerv11.Methods, nil
	case tools.V21.IsSupported(network, height):
		return powerv12.Methods, nil
	case tools.V22.IsSupported(network, height):
		return powerv13.Methods, nil
	case tools.V23.IsSupported(network, height):
		return powerv14.Methods, nil
	case tools.V24.IsSupported(network, height):
		return powerv15.Methods, nil
	default:
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
}

func (*Power) CurrentTotalPower(network string, msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	var data map[string]interface{}
	var err error

	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		data, err = parse(rawReturn, nil, false, &legacyv1.CurrentTotalPowerReturn{}, &legacyv1.CurrentTotalPowerReturn{}, parser.ReturnKey)
	case tools.AnyIsSupported(network, height, tools.V8, tools.V9):
		data, err = parse(rawReturn, nil, false, &legacyv2.CurrentTotalPowerReturn{}, &legacyv2.CurrentTotalPowerReturn{}, parser.ReturnKey)
	case tools.AnyIsSupported(network, height, tools.V10, tools.V11):
		data, err = parse(rawReturn, nil, false, &legacyv3.CurrentTotalPowerReturn{}, &legacyv3.CurrentTotalPowerReturn{}, parser.ReturnKey)
	case tools.V12.IsSupported(network, height):
		data, err = parse(rawReturn, nil, false, &legacyv4.CurrentTotalPowerReturn{}, &legacyv4.CurrentTotalPowerReturn{}, parser.ReturnKey)
	case tools.V13.IsSupported(network, height):
		data, err = parse(rawReturn, nil, false, &legacyv5.CurrentTotalPowerReturn{}, &legacyv5.CurrentTotalPowerReturn{}, parser.ReturnKey)
	case tools.V14.IsSupported(network, height):
		data, err = parse(rawReturn, nil, false, &legacyv6.CurrentTotalPowerReturn{}, &legacyv6.CurrentTotalPowerReturn{}, parser.ReturnKey)
	case tools.V15.IsSupported(network, height):
		data, err = parse(rawReturn, nil, false, &legacyv7.CurrentTotalPowerReturn{}, &legacyv7.CurrentTotalPowerReturn{}, parser.ReturnKey)
	case tools.V16.IsSupported(network, height):
		data, err = parse(rawReturn, nil, false, &powerv8.CurrentTotalPowerReturn{}, &powerv8.CurrentTotalPowerReturn{}, parser.ReturnKey)
	case tools.V17.IsSupported(network, height):
		data, err = parse(rawReturn, nil, false, &powerv9.CurrentTotalPowerReturn{}, &powerv9.CurrentTotalPowerReturn{}, parser.ReturnKey)
	case tools.V18.IsSupported(network, height):
		data, err = parse(rawReturn, nil, false, &powerv10.CurrentTotalPowerReturn{}, &powerv10.CurrentTotalPowerReturn{}, parser.ReturnKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		data, err = parse(rawReturn, nil, false, &powerv11.CurrentTotalPowerReturn{}, &powerv11.CurrentTotalPowerReturn{}, parser.ReturnKey)
	case tools.V21.IsSupported(network, height):
		data, err = parse(rawReturn, nil, false, &powerv12.CurrentTotalPowerReturn{}, &powerv12.CurrentTotalPowerReturn{}, parser.ReturnKey)
	case tools.V22.IsSupported(network, height):
		data, err = parse(rawReturn, nil, false, &powerv13.CurrentTotalPowerReturn{}, &powerv13.CurrentTotalPowerReturn{}, parser.ReturnKey)
	case tools.V23.IsSupported(network, height):
		data, err = parse(rawReturn, nil, false, &powerv14.CurrentTotalPowerReturn{}, &powerv14.CurrentTotalPowerReturn{}, parser.ReturnKey)
		if err != nil {
			// try to parse with V24 ( 2 extra fields added)
			data, err = parse(rawReturn, nil, false, &powerv15.CurrentTotalPowerReturn{}, &powerv15.CurrentTotalPowerReturn{}, parser.ReturnKey)
		}
	case tools.V24.IsSupported(network, height):
		data, err = parse(rawReturn, nil, false, &powerv15.CurrentTotalPowerReturn{}, &powerv15.CurrentTotalPowerReturn{}, parser.ReturnKey)
	default:
		err = fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return data, err
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
	var data map[string]interface{}
	var err error
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		data, err = parse(raw, nil, false, &legacyv1.MinerConstructorParams{}, &legacyv1.MinerConstructorParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V8, tools.V9):
		data, err = parse(raw, nil, false, &legacyv2.MinerConstructorParams{}, &legacyv2.MinerConstructorParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V10, tools.V11):
		data, err = parse(raw, nil, false, &legacyv3.MinerConstructorParams{}, &legacyv3.MinerConstructorParams{}, parser.ParamsKey)
	case tools.V12.IsSupported(network, height):
		data, err = parse(raw, nil, false, &legacyv4.MinerConstructorParams{}, &legacyv4.MinerConstructorParams{}, parser.ParamsKey)
	case tools.V13.IsSupported(network, height):
		data, err = parse(raw, nil, false, &legacyv5.MinerConstructorParams{}, &legacyv5.MinerConstructorParams{}, parser.ParamsKey)
	case tools.V14.IsSupported(network, height):
		data, err = parse(raw, nil, false, &legacyv6.MinerConstructorParams{}, &legacyv6.MinerConstructorParams{}, parser.ParamsKey)
	case tools.V15.IsSupported(network, height):
		data, err = parse(raw, nil, false, &legacyv7.MinerConstructorParams{}, &legacyv7.MinerConstructorParams{}, parser.ParamsKey)
	case tools.V16.IsSupported(network, height):
		data, err = parse(raw, nil, false, &powerv8.MinerConstructorParams{}, &powerv8.MinerConstructorParams{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		data, err = parse(raw, nil, false, &powerv9.MinerConstructorParams{}, &powerv9.MinerConstructorParams{}, parser.ParamsKey)
	case tools.V18.IsSupported(network, height):
		data, err = parse(raw, nil, false, &powerv10.MinerConstructorParams{}, &powerv10.MinerConstructorParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		data, err = parse(raw, nil, false, &powerv11.MinerConstructorParams{}, &powerv11.MinerConstructorParams{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		data, err = parse(raw, nil, false, &powerv12.MinerConstructorParams{}, &powerv12.MinerConstructorParams{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		data, err = parse(raw, nil, false, &powerv13.MinerConstructorParams{}, &powerv13.MinerConstructorParams{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		data, err = parse(raw, nil, false, &powerv14.MinerConstructorParams{}, &powerv14.MinerConstructorParams{}, parser.ParamsKey)
	case tools.V24.IsSupported(network, height):
		data, err = parse(raw, nil, false, &powerv15.MinerConstructorParams{}, &powerv15.MinerConstructorParams{}, parser.ParamsKey)
	default:
		err = fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return data, err
}

func (p *Power) CreateMinerExported(network string, msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, *types.AddressInfo, error) {
	var data map[string]interface{}
	var addressInfo *types.AddressInfo
	var err error
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		data, addressInfo, err = parseCreateMiner(msg, raw, rawReturn, &legacyv1.CreateMinerParams{}, &legacyv1.CreateMinerReturn{})
	case tools.AnyIsSupported(network, height, tools.V8, tools.V9):
		data, addressInfo, err = parseCreateMiner(msg, raw, rawReturn, &legacyv2.CreateMinerParams{}, &legacyv2.CreateMinerReturn{})
	case tools.AnyIsSupported(network, height, tools.V10, tools.V11):
		data, addressInfo, err = parseCreateMiner(msg, raw, rawReturn, &legacyv3.CreateMinerParams{}, &legacyv3.CreateMinerReturn{})
	case tools.V12.IsSupported(network, height):
		data, addressInfo, err = parseCreateMiner(msg, raw, rawReturn, &legacyv4.CreateMinerParams{}, &legacyv4.CreateMinerReturn{})
	case tools.V13.IsSupported(network, height):
		data, addressInfo, err = parseCreateMiner(msg, raw, rawReturn, &legacyv5.CreateMinerParams{}, &legacyv5.CreateMinerReturn{})
	case tools.V14.IsSupported(network, height):
		data, addressInfo, err = parseCreateMiner(msg, raw, rawReturn, &legacyv6.CreateMinerParams{}, &legacyv6.CreateMinerReturn{})
	case tools.V15.IsSupported(network, height):
		data, addressInfo, err = parseCreateMiner(msg, raw, rawReturn, &legacyv7.CreateMinerParams{}, &legacyv7.CreateMinerReturn{})

	case tools.V16.IsSupported(network, height):
		data, addressInfo, err = parseCreateMiner(msg, raw, rawReturn, &powerv8.CreateMinerParams{}, &powerv8.CreateMinerReturn{})
	case tools.V17.IsSupported(network, height):
		data, addressInfo, err = parseCreateMiner(msg, raw, rawReturn, &powerv9.CreateMinerParams{}, &powerv9.CreateMinerReturn{})
	case tools.V18.IsSupported(network, height):
		data, addressInfo, err = parseCreateMiner(msg, raw, rawReturn, &powerv10.CreateMinerParams{}, &powerv10.CreateMinerReturn{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		data, addressInfo, err = parseCreateMiner(msg, raw, rawReturn, &powerv11.CreateMinerParams{}, &powerv11.CreateMinerReturn{})
	case tools.V21.IsSupported(network, height):
		data, addressInfo, err = parseCreateMiner(msg, raw, rawReturn, &powerv12.CreateMinerParams{}, &powerv12.CreateMinerReturn{})
	case tools.V22.IsSupported(network, height):
		data, addressInfo, err = parseCreateMiner(msg, raw, rawReturn, &powerv13.CreateMinerParams{}, &powerv13.CreateMinerReturn{})
	case tools.V23.IsSupported(network, height):
		data, addressInfo, err = parseCreateMiner(msg, raw, rawReturn, &powerv14.CreateMinerParams{}, &powerv14.CreateMinerReturn{})
	case tools.V24.IsSupported(network, height):
		data, addressInfo, err = parseCreateMiner(msg, raw, rawReturn, &powerv15.CreateMinerParams{}, &powerv15.CreateMinerReturn{})

	default:
		err = fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	if addressInfo != nil {
		p.helper.GetActorsCache().StoreAddressInfoAddress(*addressInfo)
	}

	return data, addressInfo, err
}

func (*Power) EnrollCronEvent(network string, msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	var data map[string]interface{}
	var err error
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		data, err = parse(raw, rawReturn, false, &legacyv1.EnrollCronEventParams{}, &legacyv1.EnrollCronEventParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V8, tools.V9):
		data, err = parse(raw, rawReturn, false, &legacyv2.EnrollCronEventParams{}, &legacyv2.EnrollCronEventParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V10, tools.V11):
		data, err = parse(raw, rawReturn, false, &legacyv3.EnrollCronEventParams{}, &legacyv3.EnrollCronEventParams{}, parser.ParamsKey)
	case tools.V12.IsSupported(network, height):
		data, err = parse(raw, rawReturn, false, &legacyv4.EnrollCronEventParams{}, &legacyv4.EnrollCronEventParams{}, parser.ParamsKey)
	case tools.V13.IsSupported(network, height):
		data, err = parse(raw, rawReturn, false, &legacyv5.EnrollCronEventParams{}, &legacyv5.EnrollCronEventParams{}, parser.ParamsKey)
	case tools.V14.IsSupported(network, height):
		data, err = parse(raw, rawReturn, false, &legacyv6.EnrollCronEventParams{}, &legacyv6.EnrollCronEventParams{}, parser.ParamsKey)
	case tools.V15.IsSupported(network, height):
		data, err = parse(raw, rawReturn, false, &legacyv7.EnrollCronEventParams{}, &legacyv7.EnrollCronEventParams{}, parser.ParamsKey)

	case tools.V16.IsSupported(network, height):
		data, err = parse(raw, rawReturn, false, &powerv8.EnrollCronEventParams{}, &powerv8.EnrollCronEventParams{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		data, err = parse(raw, rawReturn, false, &powerv9.EnrollCronEventParams{}, &powerv9.EnrollCronEventParams{}, parser.ParamsKey)
	case tools.V18.IsSupported(network, height):
		data, err = parse(raw, rawReturn, false, &powerv10.EnrollCronEventParams{}, &powerv10.EnrollCronEventParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		data, err = parse(raw, rawReturn, false, &powerv11.EnrollCronEventParams{}, &powerv11.EnrollCronEventParams{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		data, err = parse(raw, rawReturn, false, &powerv12.EnrollCronEventParams{}, &powerv12.EnrollCronEventParams{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		data, err = parse(raw, rawReturn, false, &powerv13.EnrollCronEventParams{}, &powerv13.EnrollCronEventParams{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		data, err = parse(raw, rawReturn, false, &powerv14.EnrollCronEventParams{}, &powerv14.EnrollCronEventParams{}, parser.ParamsKey)
	case tools.V24.IsSupported(network, height):
		data, err = parse(raw, rawReturn, false, &powerv15.EnrollCronEventParams{}, &powerv15.EnrollCronEventParams{}, parser.ParamsKey)

	default:
		err = fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return data, err
}

func (*Power) UpdateClaimedPower(network string, msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	var data map[string]interface{}
	var err error
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		data, err = parse(raw, rawReturn, false, &legacyv1.UpdateClaimedPowerParams{}, &legacyv1.UpdateClaimedPowerParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V8, tools.V9):
		data, err = parse(raw, rawReturn, false, &legacyv2.UpdateClaimedPowerParams{}, &legacyv2.UpdateClaimedPowerParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V10, tools.V11):
		data, err = parse(raw, rawReturn, false, &legacyv3.UpdateClaimedPowerParams{}, &legacyv3.UpdateClaimedPowerParams{}, parser.ParamsKey)
	case tools.V12.IsSupported(network, height):
		data, err = parse(raw, rawReturn, false, &legacyv4.UpdateClaimedPowerParams{}, &legacyv4.UpdateClaimedPowerParams{}, parser.ParamsKey)
	case tools.V13.IsSupported(network, height):
		data, err = parse(raw, rawReturn, false, &legacyv5.UpdateClaimedPowerParams{}, &legacyv5.UpdateClaimedPowerParams{}, parser.ParamsKey)
	case tools.V14.IsSupported(network, height):
		data, err = parse(raw, rawReturn, false, &legacyv6.UpdateClaimedPowerParams{}, &legacyv6.UpdateClaimedPowerParams{}, parser.ParamsKey)
	case tools.V15.IsSupported(network, height):
		data, err = parse(raw, rawReturn, false, &legacyv7.UpdateClaimedPowerParams{}, &legacyv7.UpdateClaimedPowerParams{}, parser.ParamsKey)
	case tools.V16.IsSupported(network, height):
		data, err = parse(raw, rawReturn, false, &powerv8.UpdateClaimedPowerParams{}, &powerv8.UpdateClaimedPowerParams{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		data, err = parse(raw, rawReturn, false, &powerv9.UpdateClaimedPowerParams{}, &powerv9.UpdateClaimedPowerParams{}, parser.ParamsKey)
	case tools.V18.IsSupported(network, height):
		data, err = parse(raw, rawReturn, false, &powerv10.UpdateClaimedPowerParams{}, &powerv10.UpdateClaimedPowerParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		data, err = parse(raw, rawReturn, false, &powerv11.UpdateClaimedPowerParams{}, &powerv11.UpdateClaimedPowerParams{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		data, err = parse(raw, rawReturn, false, &powerv12.UpdateClaimedPowerParams{}, &powerv12.UpdateClaimedPowerParams{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		data, err = parse(raw, rawReturn, false, &powerv13.UpdateClaimedPowerParams{}, &powerv13.UpdateClaimedPowerParams{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		data, err = parse(raw, rawReturn, false, &powerv14.UpdateClaimedPowerParams{}, &powerv14.UpdateClaimedPowerParams{}, parser.ParamsKey)
	case tools.V24.IsSupported(network, height):
		data, err = parse(raw, rawReturn, false, &powerv15.UpdateClaimedPowerParams{}, &powerv15.UpdateClaimedPowerParams{}, parser.ParamsKey)
	default:
		err = fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return data, err
}

func (*Power) UpdatePledgeTotal(network string, msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	data, err := parse(raw, rawReturn, false, &abi.TokenAmount{}, &abi.TokenAmount{}, parser.ParamsKey)
	return data, err
}

func (*Power) NetworkRawPowerExported(network string, msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	var data map[string]interface{}
	var err error

	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V17)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	case tools.V18.IsSupported(network, height):
		data, err = parse(rawReturn, nil, false, &powerv10.NetworkRawPowerReturn{}, &powerv10.NetworkRawPowerReturn{}, parser.ReturnKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		data, err = parse(rawReturn, nil, false, &powerv11.NetworkRawPowerReturn{}, &powerv11.NetworkRawPowerReturn{}, parser.ReturnKey)
	case tools.V21.IsSupported(network, height):
		data, err = parse(rawReturn, nil, false, &powerv12.NetworkRawPowerReturn{}, &powerv12.NetworkRawPowerReturn{}, parser.ReturnKey)
	case tools.V22.IsSupported(network, height):
		data, err = parse(rawReturn, nil, false, &powerv13.NetworkRawPowerReturn{}, &powerv13.NetworkRawPowerReturn{}, parser.ReturnKey)
	case tools.V23.IsSupported(network, height):
		data, err = parse(rawReturn, nil, false, &powerv14.NetworkRawPowerReturn{}, &powerv14.NetworkRawPowerReturn{}, parser.ReturnKey)
	case tools.V24.IsSupported(network, height):
		data, err = parse(rawReturn, nil, false, &powerv15.NetworkRawPowerReturn{}, &powerv15.NetworkRawPowerReturn{}, parser.ReturnKey)
	default:
		err = fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return data, err
}

func (*Power) MinerRawPowerExported(network string, msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	var data map[string]interface{}
	var err error
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V17)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	case tools.V18.IsSupported(network, height):
		var params powerv10.MinerRawPowerParams
		data, err = parse(raw, rawReturn, true, &params, &powerv10.MinerRawPowerReturn{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		var params powerv11.MinerRawPowerParams
		data, err = parse(raw, rawReturn, true, &params, &powerv11.MinerRawPowerReturn{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		var params powerv12.MinerRawPowerParams
		data, err = parse(raw, rawReturn, true, &params, &powerv12.MinerRawPowerReturn{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		var params powerv13.MinerRawPowerParams
		data, err = parse(raw, rawReturn, true, &params, &powerv13.MinerRawPowerReturn{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		var params powerv14.MinerRawPowerParams
		data, err = parse(raw, rawReturn, true, &params, &powerv14.MinerRawPowerReturn{}, parser.ParamsKey)
	case tools.V24.IsSupported(network, height):
		var params powerv15.MinerRawPowerParams
		data, err = parse(raw, rawReturn, true, &params, &powerv15.MinerRawPowerReturn{}, parser.ParamsKey)
	default:
		err = fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return data, err
}

func (*Power) MinerCountExported(network string, msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	var data map[string]interface{}
	var err error
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V17)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	case tools.V18.IsSupported(network, height):
		var params powerv10.MinerCountReturn
		data, err = parse(rawReturn, nil, false, &params, &params, parser.ReturnKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		var params powerv11.MinerCountReturn
		data, err = parse(rawReturn, nil, false, &params, &params, parser.ReturnKey)
	case tools.V21.IsSupported(network, height):
		var params powerv12.MinerCountReturn
		data, err = parse(rawReturn, nil, false, &params, &params, parser.ReturnKey)
	case tools.V22.IsSupported(network, height):
		var params powerv13.MinerCountReturn
		data, err = parse(rawReturn, nil, false, &params, &params, parser.ReturnKey)
	case tools.V23.IsSupported(network, height):
		var params powerv14.MinerCountReturn
		data, err = parse(rawReturn, nil, false, &params, &params, parser.ReturnKey)
	case tools.V24.IsSupported(network, height):
		var params powerv15.MinerCountReturn
		data, err = parse(rawReturn, nil, false, &params, &params, parser.ReturnKey)
	default:
		err = fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return data, err
}

func (*Power) MinerConsensusCountExported(network string, msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	var data map[string]interface{}
	var err error
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V17)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	case tools.V18.IsSupported(network, height):
		var params powerv10.MinerConsensusCountReturn
		data, err = parse(rawReturn, nil, false, &params, &params, parser.ReturnKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		var params powerv11.MinerConsensusCountReturn
		data, err = parse(rawReturn, nil, false, &params, &params, parser.ReturnKey)
	case tools.V21.IsSupported(network, height):
		var params powerv12.MinerConsensusCountReturn
		data, err = parse(rawReturn, nil, false, &params, &params, parser.ReturnKey)
	case tools.V22.IsSupported(network, height):
		var params powerv13.MinerConsensusCountReturn
		data, err = parse(rawReturn, nil, false, &params, &params, parser.ReturnKey)
	case tools.V23.IsSupported(network, height):
		var params powerv14.MinerConsensusCountReturn
		data, err = parse(rawReturn, nil, false, &params, &params, parser.ReturnKey)
	case tools.V24.IsSupported(network, height):
		var params powerv15.MinerConsensusCountReturn
		data, err = parse(rawReturn, nil, false, &params, &params, parser.ReturnKey)
	default:
		err = fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return data, err
}
