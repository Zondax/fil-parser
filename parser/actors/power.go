package actors

import (
	"bytes"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/zondax/fil-parser/parser"

	"github.com/filecoin-project/go-state-types/builtin/v11/power"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/specs-actors/actors/runtime/proof"

	"github.com/zondax/fil-parser/types"
)

func ParseStoragepower(txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt) (map[string]interface{}, error) {
	switch txType {
	case parser.MethodSend:
		return parseSend(msg), nil
	case parser.MethodConstructor:
		return powerConstructor(msg.Params)
	case parser.MethodCreateMiner, parser.MethodCreateMinerExported:
		return parseCreateMiner(msg, msgRct)
	case parser.MethodUpdateClaimedPower:
		return updateClaimedPower(msg.Params)
	case parser.MethodEnrollCronEvent:
		return enrollCronEvent(msg.Params)
	case parser.MethodCronTick:
		return emptyParamsAndReturn()
	case parser.MethodUpdatePledgeTotal:
		return updatePledgeTotal(msg.Params)
	case parser.MethodPowerDeprecated1: // OnConsensusFault
	case parser.MethodSubmitPoRepForBulkVerify:
		return submitPoRepForBulkVerify(msg.Params)
	case parser.MethodCurrentTotalPower:
		return currentTotalPower(msgRct.Return)
	case parser.MethodNetworkRawPowerExported:
		return networkRawPower(msgRct.Return)
	case parser.MethodMinerRawPowerExported:
		return minerRawPower(msg.Params, msgRct.Return)
	case parser.MethodMinerCountExported:
		return minerCount(msgRct.Return)
	case parser.MethodMinerConsensusCountExported:
		return minerConsensusCount(msgRct.Return)
	case parser.UnknownStr:
		return unknownMetadata(msg.Params, msgRct.Return)

	}
	return map[string]interface{}{}, parser.ErrUnknownMethod
}

func currentTotalPower(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params power.CurrentTotalPowerReturn
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = params
	return metadata, nil
}

func submitPoRepForBulkVerify(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params proof.SealVerifyInfo
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func powerConstructor(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params power.MinerConstructorParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func parseCreateMiner(msg *parser.LotusMessage, msgRct *filTypes.MessageReceipt) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(msg.Params)
	var params power.CreateMinerParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params

	reader = bytes.NewReader(msgRct.Return)
	var r power.CreateMinerReturn
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	createdActor := &types.AddressInfo{
		Short:          r.IDAddress.String(),
		Robust:         r.RobustAddress.String(),
		ActorType:      "miner",
		CreationTxHash: msg.Cid.String(),
	}
	metadata[parser.ReturnKey] = createdActor
	appendToAddresses(*createdActor)
	return metadata, nil
}

func enrollCronEvent(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params power.EnrollCronEventParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func updateClaimedPower(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params power.UpdateClaimedPowerParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func updatePledgeTotal(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params abi.TokenAmount
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func networkRawPower(rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawReturn)
	var r power.NetworkRawPowerReturn
	err := r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r
	return metadata, nil
}

func minerRawPower(raw, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params power.MinerRawPowerParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params

	reader = bytes.NewReader(rawReturn)
	var r power.MinerRawPowerReturn
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r
	return metadata, nil
}

func minerCount(rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawReturn)
	var r power.MinerCountReturn
	err := r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r
	return metadata, nil
}

func minerConsensusCount(rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawReturn)
	var r power.MinerConsensusCountReturn
	err := r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r
	return metadata, nil
}
