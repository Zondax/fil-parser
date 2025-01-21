package power

import (
	"bytes"
	"io"

	"github.com/filecoin-project/go-address"
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
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
)

type powerParams interface {
	UnmarshalCBOR(io.Reader) error
}

type powerReturn interface {
	UnmarshalCBOR(io.Reader) error
}

func getAddressInfo(r powerReturn, msg *parser.LotusMessage) *types.AddressInfo {
	createAddressInfo := func(idAddress, robustAddress address.Address, cid cid.Cid) *types.AddressInfo {
		return &types.AddressInfo{
			Short:         idAddress.String(),
			Robust:        robustAddress.String(),
			ActorType:     "miner",
			CreationTxCid: cid.String(),
		}
	}
	switch r := r.(type) {
	case *powerv8.CreateMinerReturn:
		return createAddressInfo(r.IDAddress, r.RobustAddress, msg.Cid)
	case *powerv9.CreateMinerReturn:
		return createAddressInfo(r.IDAddress, r.RobustAddress, msg.Cid)
	case *powerv10.CreateMinerReturn:
		return createAddressInfo(r.IDAddress, r.RobustAddress, msg.Cid)
	case *powerv11.CreateMinerReturn:
		return createAddressInfo(r.IDAddress, r.RobustAddress, msg.Cid)
	case *powerv12.CreateMinerReturn:
		return createAddressInfo(r.IDAddress, r.RobustAddress, msg.Cid)
	case *powerv13.CreateMinerReturn:
		return createAddressInfo(r.IDAddress, r.RobustAddress, msg.Cid)
	case *powerv14.CreateMinerReturn:
		return createAddressInfo(r.IDAddress, r.RobustAddress, msg.Cid)
	case *powerv15.CreateMinerReturn:
		return createAddressInfo(r.IDAddress, r.RobustAddress, msg.Cid)
	}
	return nil
}

func parse[T powerParams, R powerReturn](msg *parser.LotusMessage, raw, rawReturn []byte, customReturn bool) (map[string]interface{}, *types.AddressInfo, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var constructor T
	err := constructor.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, nil, err
	}

	metadata[parser.ParamsKey] = constructor
	if !customReturn {
		return metadata, nil, nil
	}

	reader = bytes.NewReader(rawReturn)
	var r R
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, nil, err
	}
	createdActor := getAddressInfo(r, msg)
	metadata[parser.ReturnKey] = createdActor
	return metadata, createdActor, nil
}

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
