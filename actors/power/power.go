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
	"github.com/filecoin-project/go-state-types/proof"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

func CurrentTotalPower(network string, msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(network, height):
		data, _, err := parse[*powerv8.CurrentTotalPowerReturn, *powerv8.CurrentTotalPowerReturn](msg, raw, rawReturn, false)
		return data, err
	case tools.V9.IsSupported(network, height):
		data, _, err := parse[*powerv9.CurrentTotalPowerReturn, *powerv9.CurrentTotalPowerReturn](msg, raw, rawReturn, false)
		return data, err
	case tools.V10.IsSupported(network, height):
		data, _, err := parse[*powerv10.CurrentTotalPowerReturn, *powerv10.CurrentTotalPowerReturn](msg, raw, rawReturn, false)
		return data, err
	case tools.V11.IsSupported(network, height):
		data, _, err := parse[*powerv11.CurrentTotalPowerReturn, *powerv11.CurrentTotalPowerReturn](msg, raw, rawReturn, false)
		return data, err
	case tools.V12.IsSupported(network, height):
		data, _, err := parse[*powerv12.CurrentTotalPowerReturn, *powerv12.CurrentTotalPowerReturn](msg, raw, rawReturn, false)
		return data, err
	case tools.V13.IsSupported(network, height):
		data, _, err := parse[*powerv13.CurrentTotalPowerReturn, *powerv13.CurrentTotalPowerReturn](msg, raw, rawReturn, false)
		return data, err
	case tools.V14.IsSupported(network, height):
		data, _, err := parse[*powerv14.CurrentTotalPowerReturn, *powerv14.CurrentTotalPowerReturn](msg, raw, rawReturn, false)
		return data, err
	case tools.V15.IsSupported(network, height):
		data, _, err := parse[*powerv15.CurrentTotalPowerReturn, *powerv15.CurrentTotalPowerReturn](msg, raw, rawReturn, false)
		return data, err
	}
	return nil, nil
}

func SubmitPoRepForBulkVerify(network string, msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	data, _, err := parse[*proof.SealVerifyInfo, *proof.SealVerifyInfo](msg, raw, rawReturn, false)
	return data, err
}

func PowerConstructor(network string, height int64, msg *parser.LotusMessage, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(network, height):
		data, _, err := parse[*powerv8.MinerConstructorParams, *powerv8.MinerConstructorParams](msg, raw, nil, false)
		return data, err
	case tools.V9.IsSupported(network, height):
		data, _, err := parse[*powerv9.MinerConstructorParams, *powerv9.MinerConstructorParams](msg, raw, nil, false)
		return data, err
	case tools.V10.IsSupported(network, height):
		data, _, err := parse[*powerv10.MinerConstructorParams, *powerv10.MinerConstructorParams](msg, raw, nil, false)
		return data, err
	case tools.V11.IsSupported(network, height):
		data, _, err := parse[*powerv11.MinerConstructorParams, *powerv11.MinerConstructorParams](msg, raw, nil, false)
		return data, err
	case tools.V12.IsSupported(network, height):
		data, _, err := parse[*powerv12.MinerConstructorParams, *powerv12.MinerConstructorParams](msg, raw, nil, false)
		return data, err
	case tools.V13.IsSupported(network, height):
		data, _, err := parse[*powerv13.MinerConstructorParams, *powerv13.MinerConstructorParams](msg, raw, nil, false)
		return data, err
	case tools.V14.IsSupported(network, height):
		data, _, err := parse[*powerv14.MinerConstructorParams, *powerv14.MinerConstructorParams](msg, raw, nil, false)
		return data, err
	case tools.V15.IsSupported(network, height):
		data, _, err := parse[*powerv15.MinerConstructorParams, *powerv15.MinerConstructorParams](msg, raw, nil, false)
		return data, err
	}
	return nil, nil
}

func CreateMiner(network string, msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(network, height):
		data, _, err := parse[*powerv8.CreateMinerParams, *powerv8.CreateMinerReturn](msg, raw, rawReturn, true)
		return data, err
	case tools.V9.IsSupported(network, height):
		data, _, err := parse[*powerv9.CreateMinerParams, *powerv9.CreateMinerReturn](msg, raw, rawReturn, true)
		return data, err
	case tools.V10.IsSupported(network, height):
		data, _, err := parse[*powerv10.CreateMinerParams, *powerv10.CreateMinerReturn](msg, raw, rawReturn, true)
		return data, err
	case tools.V11.IsSupported(network, height):
		data, _, err := parse[*powerv11.CreateMinerParams, *powerv11.CreateMinerReturn](msg, raw, rawReturn, true)
		return data, err
	case tools.V12.IsSupported(network, height):
		data, _, err := parse[*powerv12.CreateMinerParams, *powerv12.CreateMinerReturn](msg, raw, rawReturn, true)
		return data, err
	case tools.V13.IsSupported(network, height):
		data, _, err := parse[*powerv13.CreateMinerParams, *powerv13.CreateMinerReturn](msg, raw, rawReturn, true)
		return data, err
	case tools.V14.IsSupported(network, height):
		data, _, err := parse[*powerv14.CreateMinerParams, *powerv14.CreateMinerReturn](msg, raw, rawReturn, true)
		return data, err
	case tools.V15.IsSupported(network, height):
		data, _, err := parse[*powerv15.CreateMinerParams, *powerv15.CreateMinerReturn](msg, raw, rawReturn, true)
		return data, err
	}
	return nil, nil
}

func EnrollCronEvent(network string, msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(network, height):
		data, _, err := parse[*powerv8.EnrollCronEventParams, *powerv8.EnrollCronEventParams](msg, raw, rawReturn, true)
		return data, err
	case tools.V9.IsSupported(network, height):
		data, _, err := parse[*powerv9.EnrollCronEventParams, *powerv9.EnrollCronEventParams](msg, raw, rawReturn, true)
		return data, err
	case tools.V10.IsSupported(network, height):
		data, _, err := parse[*powerv10.EnrollCronEventParams, *powerv10.EnrollCronEventParams](msg, raw, rawReturn, true)
		return data, err
	case tools.V11.IsSupported(network, height):
		data, _, err := parse[*powerv11.EnrollCronEventParams, *powerv11.EnrollCronEventParams](msg, raw, rawReturn, true)
		return data, err
	case tools.V12.IsSupported(network, height):
		data, _, err := parse[*powerv12.EnrollCronEventParams, *powerv12.EnrollCronEventParams](msg, raw, rawReturn, true)
		return data, err
	case tools.V13.IsSupported(network, height):
		data, _, err := parse[*powerv13.EnrollCronEventParams, *powerv13.EnrollCronEventParams](msg, raw, rawReturn, true)
		return data, err
	case tools.V14.IsSupported(network, height):
		data, _, err := parse[*powerv14.EnrollCronEventParams, *powerv14.EnrollCronEventParams](msg, raw, rawReturn, true)
		return data, err
	case tools.V15.IsSupported(network, height):
		data, _, err := parse[*powerv15.EnrollCronEventParams, *powerv15.EnrollCronEventParams](msg, raw, rawReturn, true)
		return data, err
	}
	return nil, nil
}

func UpdateClaimedPower(network string, msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(network, height):
		data, _, err := parse[*powerv8.UpdateClaimedPowerParams, *powerv8.UpdateClaimedPowerParams](msg, raw, rawReturn, true)
		return data, err
	case tools.V9.IsSupported(network, height):
		data, _, err := parse[*powerv9.UpdateClaimedPowerParams, *powerv9.UpdateClaimedPowerParams](msg, raw, rawReturn, true)
		return data, err
	case tools.V10.IsSupported(network, height):
		data, _, err := parse[*powerv10.UpdateClaimedPowerParams, *powerv10.UpdateClaimedPowerParams](msg, raw, rawReturn, true)
		return data, err
	case tools.V11.IsSupported(network, height):
		data, _, err := parse[*powerv11.UpdateClaimedPowerParams, *powerv11.UpdateClaimedPowerParams](msg, raw, rawReturn, true)
		return data, err
	case tools.V12.IsSupported(network, height):
		data, _, err := parse[*powerv12.UpdateClaimedPowerParams, *powerv12.UpdateClaimedPowerParams](msg, raw, rawReturn, true)
		return data, err
	case tools.V13.IsSupported(network, height):
		data, _, err := parse[*powerv13.UpdateClaimedPowerParams, *powerv13.UpdateClaimedPowerParams](msg, raw, rawReturn, true)
		return data, err
	case tools.V14.IsSupported(network, height):
		data, _, err := parse[*powerv14.UpdateClaimedPowerParams, *powerv14.UpdateClaimedPowerParams](msg, raw, rawReturn, true)
		return data, err
	case tools.V15.IsSupported(network, height):
		data, _, err := parse[*powerv15.UpdateClaimedPowerParams, *powerv15.UpdateClaimedPowerParams](msg, raw, rawReturn, true)
		return data, err
	}
	return nil, nil
}

func UpdatePledgeTotal(network string, msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(network, height):
		data, _, err := parse[*abi.TokenAmount, *abi.TokenAmount](msg, raw, rawReturn, false)
		return data, err
	case tools.V9.IsSupported(network, height):
		data, _, err := parse[*abi.TokenAmount, *abi.TokenAmount](msg, raw, rawReturn, false)
		return data, err
	case tools.V10.IsSupported(network, height):
		data, _, err := parse[*abi.TokenAmount, *abi.TokenAmount](msg, raw, rawReturn, false)
		return data, err
	case tools.V11.IsSupported(network, height):
		data, _, err := parse[*abi.TokenAmount, *abi.TokenAmount](msg, raw, rawReturn, false)
		return data, err
	case tools.V12.IsSupported(network, height):
		data, _, err := parse[*abi.TokenAmount, *abi.TokenAmount](msg, raw, rawReturn, false)
		return data, err
	case tools.V13.IsSupported(network, height):
		data, _, err := parse[*abi.TokenAmount, *abi.TokenAmount](msg, raw, rawReturn, false)
		return data, err
	case tools.V14.IsSupported(network, height):
		data, _, err := parse[*abi.TokenAmount, *abi.TokenAmount](msg, raw, rawReturn, false)
		return data, err
	case tools.V15.IsSupported(network, height):
		data, _, err := parse[*abi.TokenAmount, *abi.TokenAmount](msg, raw, rawReturn, false)
		return data, err
	}
	return nil, nil
}

func NetworkRawPower(network string, msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(network, height):
		return nil, fmt.Errorf("not supported")
	case tools.V9.IsSupported(network, height):
		return nil, fmt.Errorf("not supported")
	case tools.V10.IsSupported(network, height):
		data, _, err := parse[*powerv10.NetworkRawPowerReturn, *powerv10.NetworkRawPowerReturn](msg, raw, rawReturn, false)
		return data, err
	case tools.V11.IsSupported(network, height):
		data, _, err := parse[*powerv11.NetworkRawPowerReturn, *powerv11.NetworkRawPowerReturn](msg, raw, rawReturn, false)
		return data, err
	case tools.V12.IsSupported(network, height):
		data, _, err := parse[*powerv12.NetworkRawPowerReturn, *powerv12.NetworkRawPowerReturn](msg, raw, rawReturn, false)
		return data, err
	case tools.V13.IsSupported(network, height):
		data, _, err := parse[*powerv13.NetworkRawPowerReturn, *powerv13.NetworkRawPowerReturn](msg, raw, rawReturn, false)
		return data, err
	case tools.V14.IsSupported(network, height):
		data, _, err := parse[*powerv14.NetworkRawPowerReturn, *powerv14.NetworkRawPowerReturn](msg, raw, rawReturn, false)
		return data, err
	case tools.V15.IsSupported(network, height):
		data, _, err := parse[*powerv15.NetworkRawPowerReturn, *powerv15.NetworkRawPowerReturn](msg, raw, rawReturn, false)
		return data, err
	}
	return nil, nil
}

func MinerRawPower(network string, msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(network, height):
		return nil, fmt.Errorf("not supported")
	case tools.V9.IsSupported(network, height):
		return nil, fmt.Errorf("not supported")
	case tools.V10.IsSupported(network, height):
		data, _, err := parse[*powerv10.MinerRawPowerParams, *powerv10.MinerRawPowerReturn](msg, raw, rawReturn, true)
		return data, err
	case tools.V11.IsSupported(network, height):
		data, _, err := parse[*powerv11.MinerRawPowerParams, *powerv11.MinerRawPowerReturn](msg, raw, rawReturn, true)
		return data, err
	case tools.V12.IsSupported(network, height):
		data, _, err := parse[*powerv12.MinerRawPowerParams, *powerv12.MinerRawPowerReturn](msg, raw, rawReturn, true)
		return data, err
	case tools.V13.IsSupported(network, height):
		data, _, err := parse[*powerv13.MinerRawPowerParams, *powerv13.MinerRawPowerReturn](msg, raw, rawReturn, true)
		return data, err
	case tools.V14.IsSupported(network, height):
		data, _, err := parse[*powerv14.MinerRawPowerParams, *powerv14.MinerRawPowerReturn](msg, raw, rawReturn, true)
		return data, err
	case tools.V15.IsSupported(network, height):
		data, _, err := parse[*powerv15.MinerRawPowerParams, *powerv15.MinerRawPowerReturn](msg, raw, rawReturn, true)
		return data, err
	}
	return nil, nil
}

func MinerCount(network string, msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(network, height):
		return nil, fmt.Errorf("not supported")
	case tools.V9.IsSupported(network, height):
		return nil, fmt.Errorf("not supported")
	case tools.V10.IsSupported(network, height):
		data, _, err := parse[*powerv10.MinerCountReturn, *powerv10.MinerCountReturn](msg, raw, rawReturn, false)
		return data, err
	case tools.V11.IsSupported(network, height):
		data, _, err := parse[*powerv11.MinerCountReturn, *powerv11.MinerCountReturn](msg, raw, rawReturn, false)
		return data, err
	case tools.V12.IsSupported(network, height):
		data, _, err := parse[*powerv12.MinerCountReturn, *powerv12.MinerCountReturn](msg, raw, rawReturn, false)
		return data, err
	case tools.V13.IsSupported(network, height):
		data, _, err := parse[*powerv13.MinerCountReturn, *powerv13.MinerCountReturn](msg, raw, rawReturn, false)
		return data, err
	case tools.V14.IsSupported(network, height):
		data, _, err := parse[*powerv14.MinerCountReturn, *powerv14.MinerCountReturn](msg, raw, rawReturn, false)
		return data, err
	case tools.V15.IsSupported(network, height):
		data, _, err := parse[*powerv15.MinerCountReturn, *powerv15.MinerCountReturn](msg, raw, rawReturn, false)
		return data, err
	}
	return nil, nil
}

func MinerConsensusCount(network string, msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(network, height):
		return nil, fmt.Errorf("not supported")
	case tools.V9.IsSupported(network, height):
		return nil, fmt.Errorf("not supported")
	case tools.V10.IsSupported(network, height):
		data, _, err := parse[*powerv10.MinerConsensusCountReturn, *powerv10.MinerConsensusCountReturn](msg, raw, rawReturn, false)
		return data, err
	case tools.V11.IsSupported(network, height):
		data, _, err := parse[*powerv11.MinerConsensusCountReturn, *powerv11.MinerConsensusCountReturn](msg, raw, rawReturn, false)
		return data, err
	case tools.V12.IsSupported(network, height):
		data, _, err := parse[*powerv12.MinerConsensusCountReturn, *powerv12.MinerConsensusCountReturn](msg, raw, rawReturn, false)
		return data, err
	case tools.V13.IsSupported(network, height):
		data, _, err := parse[*powerv13.MinerConsensusCountReturn, *powerv13.MinerConsensusCountReturn](msg, raw, rawReturn, false)
		return data, err
	case tools.V14.IsSupported(network, height):
		data, _, err := parse[*powerv14.MinerConsensusCountReturn, *powerv14.MinerConsensusCountReturn](msg, raw, rawReturn, false)
		return data, err
	case tools.V15.IsSupported(network, height):
		data, _, err := parse[*powerv15.MinerConsensusCountReturn, *powerv15.MinerConsensusCountReturn](msg, raw, rawReturn, false)
		return data, err
	}
	return nil, nil
}
