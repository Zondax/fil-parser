package power

import (
	"github.com/filecoin-project/go-state-types/abi"
	powerv8 "github.com/filecoin-project/go-state-types/builtin/v8/power"
	powerv9 "github.com/filecoin-project/go-state-types/builtin/v9/power"
	"github.com/filecoin-project/go-state-types/proof"
	"github.com/zondax/fil-parser/parser"
)

func CurrentTotalPower(msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		data, _, err := parse[*powerv8.CurrentTotalPowerReturn, *powerv8.CurrentTotalPowerReturn](msg, raw, rawReturn, false)
		return data, err
	case 9:
		data, _, err := parse[*powerv9.CurrentTotalPowerReturn, *powerv9.CurrentTotalPowerReturn](msg, raw, rawReturn, false)
		return data, err
	}
	return nil, nil
}

func SubmitPoRepForBulkVerify(msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	data, _, err := parse[*proof.SealVerifyInfo, *proof.SealVerifyInfo](msg, raw, rawReturn, false)
	return data, err
}

func PowerConstructor(height int64, msg *parser.LotusMessage, raw []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		data, _, err := parse[*powerv8.MinerConstructorParams, *powerv8.MinerConstructorParams](msg, raw, nil, false)
		return data, err
	}
	return nil, nil
}

func CreateMiner(msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		data, _, err := parse[*powerv8.CreateMinerParams, *powerv8.CreateMinerReturn](msg, raw, rawReturn, true)
		return data, err
	}
	return nil, nil
}

func EnrollCronEvent(msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		data, _, err := parse[*powerv8.EnrollCronEventParams, *powerv8.EnrollCronEventParams](msg, raw, rawReturn, true)
		return data, err
	}
	return nil, nil
}

func UpdateClaimedPower(msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		data, _, err := parse[*powerv8.UpdateClaimedPowerParams, *powerv8.UpdateClaimedPowerParams](msg, raw, rawReturn, true)
		return data, err
	}
	return nil, nil
}

func UpdatePledgeTotal(msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		data, _, err := parse[*abi.TokenAmount, *abi.TokenAmount](msg, raw, rawReturn, false)
		return data, err
	}
	return nil, nil
}

func NetworkRawPower(msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		data, _, err := parse[*powerv8.NetworkRawPowerReturn, *powerv8.NetworkRawPowerReturn](msg, raw, rawReturn, false)
		return data, err
	}
	return nil, nil
}

func MinerRawPower(msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		data, _, err := parse[*powerv8.MinerRawPowerParams, *powerv8.MinerRawPowerReturn](msg, raw, rawReturn, true)
		return data, err
	}
	return nil, nil
}

func MinerCount(msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		data, _, err := parse[*powerv8.MinerCountReturn, *powerv8.MinerCountReturn](msg, raw, rawReturn, false)
		return data, err
	}
	return nil, nil
}

func MinerConsensusCount(msg *parser.LotusMessage, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		data, _, err := parse[*powerv8.MinerConsensusCountReturn, *powerv8.MinerConsensusCountReturn](msg, raw, rawReturn, false)
		return data, err
	}
	return nil, nil
}
