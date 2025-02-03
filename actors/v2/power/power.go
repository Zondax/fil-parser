package power

import (
	"fmt"

	"github.com/filecoin-project/go-state-types/abi"
	powerv10 "github.com/filecoin-project/go-state-types/builtin/v10/power"
	powerv11 "github.com/filecoin-project/go-state-types/builtin/v11/power"
	powerv12 "github.com/filecoin-project/go-state-types/builtin/v12/power"
	powerv13 "github.com/filecoin-project/go-state-types/builtin/v13/power"
	powerv14 "github.com/filecoin-project/go-state-types/builtin/v14/power"
	powerv15 "github.com/filecoin-project/go-state-types/builtin/v15/power"
	powerv8 "github.com/filecoin-project/go-state-types/builtin/v8/power"
	powerv9 "github.com/filecoin-project/go-state-types/builtin/v9/power"
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/filecoin-project/go-state-types/proof"
	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/power"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/power"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/power"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/power"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/power"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/power"
	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

type Power struct{}

func (p *Power) Name() string {
	return manifest.PowerKey
}

func (*Power) CurrentTotalPower(network string, msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	var data map[string]interface{}
	var err error
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V9)...):
		data, _, err = parse(msg, raw, rawReturn, false, &legacyv2.CurrentTotalPowerReturn{}, &legacyv2.CurrentTotalPowerReturn{})
	case tools.AnyIsSupported(network, height, tools.V10, tools.V11):
		data, _, err = parse(msg, raw, rawReturn, false, &legacyv3.CurrentTotalPowerReturn{}, &legacyv3.CurrentTotalPowerReturn{})
	case tools.V12.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, false, &legacyv4.CurrentTotalPowerReturn{}, &legacyv4.CurrentTotalPowerReturn{})
	case tools.V13.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, false, &legacyv5.CurrentTotalPowerReturn{}, &legacyv5.CurrentTotalPowerReturn{})
	case tools.V14.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, false, &legacyv6.CurrentTotalPowerReturn{}, &legacyv6.CurrentTotalPowerReturn{})
	case tools.V15.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, false, &legacyv7.CurrentTotalPowerReturn{}, &legacyv7.CurrentTotalPowerReturn{})
	case tools.V16.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, false, &powerv8.CurrentTotalPowerReturn{}, &powerv8.CurrentTotalPowerReturn{})
	case tools.V17.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, false, &powerv9.CurrentTotalPowerReturn{}, &powerv9.CurrentTotalPowerReturn{})
	case tools.V18.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, false, &powerv10.CurrentTotalPowerReturn{}, &powerv10.CurrentTotalPowerReturn{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		data, _, err = parse(msg, raw, rawReturn, false, &powerv11.CurrentTotalPowerReturn{}, &powerv11.CurrentTotalPowerReturn{})
	case tools.V21.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, false, &powerv12.CurrentTotalPowerReturn{}, &powerv12.CurrentTotalPowerReturn{})
	case tools.V22.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, false, &powerv13.CurrentTotalPowerReturn{}, &powerv13.CurrentTotalPowerReturn{})
	case tools.V23.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, false, &powerv14.CurrentTotalPowerReturn{}, &powerv14.CurrentTotalPowerReturn{})
	case tools.V24.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, false, &powerv15.CurrentTotalPowerReturn{}, &powerv15.CurrentTotalPowerReturn{})
	default:
		err = fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return data, err
}

func (*Power) SubmitPoRepForBulkVerify(network string, msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	data, _, err := parse(msg, raw, rawReturn, false, &proof.SealVerifyInfo{}, &proof.SealVerifyInfo{})
	return data, err
}

func (*Power) Constructor(network string, height int64, msg *parser.LotusMessage, raw []byte) (map[string]interface{}, error) {
	var data map[string]interface{}
	var err error
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V9)...):
		data, _, err = parse(msg, raw, nil, false, &legacyv2.MinerConstructorParams{}, &legacyv2.MinerConstructorParams{})
	case tools.AnyIsSupported(network, height, tools.V10, tools.V11):
		data, _, err = parse(msg, raw, nil, false, &legacyv3.MinerConstructorParams{}, &legacyv3.MinerConstructorParams{})
	case tools.V12.IsSupported(network, height):
		data, _, err = parse(msg, raw, nil, false, &legacyv4.MinerConstructorParams{}, &legacyv4.MinerConstructorParams{})
	case tools.V13.IsSupported(network, height):
		data, _, err = parse(msg, raw, nil, false, &legacyv5.MinerConstructorParams{}, &legacyv5.MinerConstructorParams{})
	case tools.V14.IsSupported(network, height):
		data, _, err = parse(msg, raw, nil, false, &legacyv6.MinerConstructorParams{}, &legacyv6.MinerConstructorParams{})
	case tools.V15.IsSupported(network, height):
		data, _, err = parse(msg, raw, nil, false, &legacyv7.MinerConstructorParams{}, &legacyv7.MinerConstructorParams{})
	case tools.V16.IsSupported(network, height):
		data, _, err = parse(msg, raw, nil, false, &powerv8.MinerConstructorParams{}, &powerv8.MinerConstructorParams{})
	case tools.V17.IsSupported(network, height):
		data, _, err = parse(msg, raw, nil, false, &powerv9.MinerConstructorParams{}, &powerv9.MinerConstructorParams{})
	case tools.V18.IsSupported(network, height):
		data, _, err = parse(msg, raw, nil, false, &powerv10.MinerConstructorParams{}, &powerv10.MinerConstructorParams{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		data, _, err = parse(msg, raw, nil, false, &powerv11.MinerConstructorParams{}, &powerv11.MinerConstructorParams{})
	case tools.V21.IsSupported(network, height):
		data, _, err = parse(msg, raw, nil, false, &powerv12.MinerConstructorParams{}, &powerv12.MinerConstructorParams{})
	case tools.V22.IsSupported(network, height):
		data, _, err = parse(msg, raw, nil, false, &powerv13.MinerConstructorParams{}, &powerv13.MinerConstructorParams{})
	case tools.V23.IsSupported(network, height):
		data, _, err = parse(msg, raw, nil, false, &powerv14.MinerConstructorParams{}, &powerv14.MinerConstructorParams{})
	case tools.V24.IsSupported(network, height):
		data, _, err = parse(msg, raw, nil, false, &powerv15.MinerConstructorParams{}, &powerv15.MinerConstructorParams{})
	default:
		err = fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return data, err
}

func (*Power) CreateMinerExported(network string, msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	var data map[string]interface{}
	var err error
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V9)...):
		data, _, err = parse(msg, raw, rawReturn, true, &legacyv2.CreateMinerParams{}, &legacyv2.CreateMinerReturn{})
	case tools.AnyIsSupported(network, height, tools.V10, tools.V11):
		data, _, err = parse(msg, raw, rawReturn, true, &legacyv3.CreateMinerParams{}, &legacyv3.CreateMinerReturn{})
	case tools.V12.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, true, &legacyv4.CreateMinerParams{}, &legacyv4.CreateMinerReturn{})
	case tools.V13.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, true, &legacyv5.CreateMinerParams{}, &legacyv5.CreateMinerReturn{})
	case tools.V14.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, true, &legacyv6.CreateMinerParams{}, &legacyv6.CreateMinerReturn{})
	case tools.V15.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, true, &legacyv7.CreateMinerParams{}, &legacyv7.CreateMinerReturn{})

	case tools.V16.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, true, &powerv8.CreateMinerParams{}, &powerv8.CreateMinerReturn{})
	case tools.V17.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, true, &powerv9.CreateMinerParams{}, &powerv9.CreateMinerReturn{})
	case tools.V18.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, true, &powerv10.CreateMinerParams{}, &powerv10.CreateMinerReturn{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		data, _, err = parse(msg, raw, rawReturn, true, &powerv11.CreateMinerParams{}, &powerv11.CreateMinerReturn{})
	case tools.V21.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, true, &powerv12.CreateMinerParams{}, &powerv12.CreateMinerReturn{})
	case tools.V22.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, true, &powerv13.CreateMinerParams{}, &powerv13.CreateMinerReturn{})
	case tools.V23.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, true, &powerv14.CreateMinerParams{}, &powerv14.CreateMinerReturn{})
	case tools.V24.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, true, &powerv15.CreateMinerParams{}, &powerv15.CreateMinerReturn{})

	default:
		err = fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return data, err
}

func (*Power) EnrollCronEvent(network string, msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	var data map[string]interface{}
	var err error
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V9)...):
		data, _, err = parse(msg, raw, rawReturn, true, &legacyv2.EnrollCronEventParams{}, &legacyv2.EnrollCronEventParams{})
	case tools.AnyIsSupported(network, height, tools.V10, tools.V11):
		data, _, err = parse(msg, raw, rawReturn, true, &legacyv3.EnrollCronEventParams{}, &legacyv3.EnrollCronEventParams{})
	case tools.V12.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, true, &legacyv4.EnrollCronEventParams{}, &legacyv4.EnrollCronEventParams{})
	case tools.V13.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, true, &legacyv5.EnrollCronEventParams{}, &legacyv5.EnrollCronEventParams{})
	case tools.V14.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, true, &legacyv6.EnrollCronEventParams{}, &legacyv6.EnrollCronEventParams{})
	case tools.V15.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, true, &legacyv7.EnrollCronEventParams{}, &legacyv7.EnrollCronEventParams{})

	case tools.V16.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, true, &powerv8.EnrollCronEventParams{}, &powerv8.EnrollCronEventParams{})
	case tools.V17.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, true, &powerv9.EnrollCronEventParams{}, &powerv9.EnrollCronEventParams{})
	case tools.V18.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, true, &powerv10.EnrollCronEventParams{}, &powerv10.EnrollCronEventParams{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		data, _, err = parse(msg, raw, rawReturn, true, &powerv11.EnrollCronEventParams{}, &powerv11.EnrollCronEventParams{})
	case tools.V21.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, true, &powerv12.EnrollCronEventParams{}, &powerv12.EnrollCronEventParams{})
	case tools.V22.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, true, &powerv13.EnrollCronEventParams{}, &powerv13.EnrollCronEventParams{})
	case tools.V23.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, true, &powerv14.EnrollCronEventParams{}, &powerv14.EnrollCronEventParams{})
	case tools.V24.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, true, &powerv15.EnrollCronEventParams{}, &powerv15.EnrollCronEventParams{})

	default:
		err = fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return data, err
}

func (*Power) UpdateClaimedPower(network string, msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	var data map[string]interface{}
	var err error
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V9)...):
		data, _, err = parse(msg, raw, rawReturn, true, &legacyv2.UpdateClaimedPowerParams{}, &legacyv2.UpdateClaimedPowerParams{})
	case tools.AnyIsSupported(network, height, tools.V10, tools.V11):
		data, _, err = parse(msg, raw, rawReturn, true, &legacyv3.UpdateClaimedPowerParams{}, &legacyv3.UpdateClaimedPowerParams{})
	case tools.V12.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, true, &legacyv4.UpdateClaimedPowerParams{}, &legacyv4.UpdateClaimedPowerParams{})
	case tools.V13.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, true, &legacyv5.UpdateClaimedPowerParams{}, &legacyv5.UpdateClaimedPowerParams{})
	case tools.V14.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, true, &legacyv6.UpdateClaimedPowerParams{}, &legacyv6.UpdateClaimedPowerParams{})
	case tools.V15.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, true, &legacyv7.UpdateClaimedPowerParams{}, &legacyv7.UpdateClaimedPowerParams{})
	case tools.V16.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, true, &powerv8.UpdateClaimedPowerParams{}, &powerv8.UpdateClaimedPowerParams{})
	case tools.V17.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, true, &powerv9.UpdateClaimedPowerParams{}, &powerv9.UpdateClaimedPowerParams{})
	case tools.V18.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, true, &powerv10.UpdateClaimedPowerParams{}, &powerv10.UpdateClaimedPowerParams{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		data, _, err = parse(msg, raw, rawReturn, true, &powerv11.UpdateClaimedPowerParams{}, &powerv11.UpdateClaimedPowerParams{})
	case tools.V21.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, true, &powerv12.UpdateClaimedPowerParams{}, &powerv12.UpdateClaimedPowerParams{})
	case tools.V22.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, true, &powerv13.UpdateClaimedPowerParams{}, &powerv13.UpdateClaimedPowerParams{})
	case tools.V23.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, true, &powerv14.UpdateClaimedPowerParams{}, &powerv14.UpdateClaimedPowerParams{})
	case tools.V24.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, true, &powerv15.UpdateClaimedPowerParams{}, &powerv15.UpdateClaimedPowerParams{})
	default:
		err = fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return data, err
}

func (*Power) UpdatePledgeTotal(network string, msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	data, _, err := parse(msg, raw, rawReturn, false, &abi.TokenAmount{}, &abi.TokenAmount{})
	return data, err
}

func (*Power) NetworkRawPowerExported(network string, msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	var data map[string]interface{}
	var err error
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V17)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	case tools.V18.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, false, &powerv10.NetworkRawPowerReturn{}, &powerv10.NetworkRawPowerReturn{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		data, _, err = parse(msg, raw, rawReturn, false, &powerv11.NetworkRawPowerReturn{}, &powerv11.NetworkRawPowerReturn{})
	case tools.V21.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, false, &powerv12.NetworkRawPowerReturn{}, &powerv12.NetworkRawPowerReturn{})
	case tools.V22.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, false, &powerv13.NetworkRawPowerReturn{}, &powerv13.NetworkRawPowerReturn{})
	case tools.V23.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, false, &powerv14.NetworkRawPowerReturn{}, &powerv14.NetworkRawPowerReturn{})
	case tools.V24.IsSupported(network, height):
		data, _, err = parse(msg, raw, rawReturn, false, &powerv15.NetworkRawPowerReturn{}, &powerv15.NetworkRawPowerReturn{})
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
		data, _, err = parse(msg, raw, rawReturn, true, &params, &powerv10.MinerRawPowerReturn{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		var params powerv11.MinerRawPowerParams
		data, _, err = parse(msg, raw, rawReturn, true, &params, &powerv11.MinerRawPowerReturn{})
	case tools.V21.IsSupported(network, height):
		var params powerv12.MinerRawPowerParams
		data, _, err = parse(msg, raw, rawReturn, true, &params, &powerv12.MinerRawPowerReturn{})
	case tools.V22.IsSupported(network, height):
		var params powerv13.MinerRawPowerParams
		data, _, err = parse(msg, raw, rawReturn, true, &params, &powerv13.MinerRawPowerReturn{})
	case tools.V23.IsSupported(network, height):
		var params powerv14.MinerRawPowerParams
		data, _, err = parse(msg, raw, rawReturn, true, &params, &powerv14.MinerRawPowerReturn{})
	case tools.V24.IsSupported(network, height):
		var params powerv15.MinerRawPowerParams
		data, _, err = parse(msg, raw, rawReturn, true, &params, &powerv15.MinerRawPowerReturn{})
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
		data, _, err = parse(msg, raw, rawReturn, false, &params, &params)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		var params powerv11.MinerCountReturn
		data, _, err = parse(msg, raw, rawReturn, false, &params, &params)
	case tools.V21.IsSupported(network, height):
		var params powerv12.MinerCountReturn
		data, _, err = parse(msg, raw, rawReturn, false, &params, &params)
	case tools.V22.IsSupported(network, height):
		var params powerv13.MinerCountReturn
		data, _, err = parse(msg, raw, rawReturn, false, &params, &params)
	case tools.V23.IsSupported(network, height):
		var params powerv14.MinerCountReturn
		data, _, err = parse(msg, raw, rawReturn, false, &params, &params)
	case tools.V24.IsSupported(network, height):
		var params powerv15.MinerCountReturn
		data, _, err = parse(msg, raw, rawReturn, false, &params, &params)
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
		data, _, err = parse(msg, raw, rawReturn, false, &params, &params)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		var params powerv11.MinerConsensusCountReturn
		data, _, err = parse(msg, raw, rawReturn, false, &params, &params)
	case tools.V21.IsSupported(network, height):
		var params powerv12.MinerConsensusCountReturn
		data, _, err = parse(msg, raw, rawReturn, false, &params, &params)
	case tools.V22.IsSupported(network, height):
		var params powerv13.MinerConsensusCountReturn
		data, _, err = parse(msg, raw, rawReturn, false, &params, &params)
	case tools.V23.IsSupported(network, height):
		var params powerv14.MinerConsensusCountReturn
		data, _, err = parse(msg, raw, rawReturn, false, &params, &params)
	case tools.V24.IsSupported(network, height):
		var params powerv15.MinerConsensusCountReturn
		data, _, err = parse(msg, raw, rawReturn, false, &params, &params)
	default:
		err = fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return data, err
}
