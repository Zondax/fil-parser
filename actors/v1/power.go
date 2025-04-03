package actors

import (
	"bytes"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin/v11/power"
	"github.com/filecoin-project/specs-actors/actors/runtime/proof"

	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
)

func (p *ActorParser) ParseStoragepower(txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt) (map[string]interface{}, *types.AddressInfo, error) {
	var err error
	var addressInfo *types.AddressInfo
	metadata := make(map[string]interface{})
	switch txType {
	case parser.MethodSend:
		metadata = p.parseSend(msg)
	case parser.MethodConstructor:
		metadata, err = p.powerConstructor(msg.Params)
	case parser.MethodCreateMiner, parser.MethodCreateMinerExported:
		return p.parseCreateMiner(msg, msgRct.Return)
	case parser.MethodUpdateClaimedPower:
		metadata, err = p.updateClaimedPower(msg.Params)
	case parser.MethodEnrollCronEvent:
		metadata, err = p.enrollCronEvent(msg.Params)
	case parser.MethodCronTick:
		metadata, err = p.emptyParamsAndReturn()
	case parser.MethodUpdatePledgeTotal:
		metadata, err = p.updatePledgeTotal(msg.Params)
	case parser.MethodSubmitPoRepForBulkVerify:
		metadata, err = p.submitPoRepForBulkVerify(msg.Params)
	case parser.MethodCurrentTotalPower:
		metadata, err = p.currentTotalPower(msgRct.Return)
	case parser.MethodNetworkRawPowerExported:
		metadata, err = p.networkRawPower(msgRct.Return)
	case parser.MethodMinerRawPowerExported:
		metadata, err = p.minerRawPower(msg.Params, msgRct.Return)
	case parser.MethodMinerCountExported:
		metadata, err = p.minerCount(msgRct.Return)
	case parser.MethodMinerConsensusCountExported:
		metadata, err = p.minerConsensusCount(msgRct.Return)
	case parser.UnknownStr:
		metadata, err = p.unknownMetadata(msg.Params, msgRct.Return)
	default:
		err = parser.ErrUnknownMethod
	}
	return metadata, addressInfo, err
}

func (p *ActorParser) currentTotalPower(raw []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) submitPoRepForBulkVerify(raw []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) powerConstructor(raw []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) parseCreateMiner(msg *parser.LotusMessage, rawReturn []byte) (map[string]interface{}, *types.AddressInfo, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(msg.Params)
	var params power.CreateMinerParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, nil, err
	}
	metadata[parser.ParamsKey] = params

	reader = bytes.NewReader(rawReturn)
	var r power.CreateMinerReturn
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, nil, err
	}
	createdActor := &types.AddressInfo{
		Short:         r.IDAddress.String(),
		Robust:        r.RobustAddress.String(),
		ActorType:     "miner",
		CreationTxCid: msg.Cid.String(),
	}
	metadata[parser.ReturnKey] = createdActor

	p.helper.GetActorsCache().StoreAddressInfoAddress(*createdActor)

	return metadata, createdActor, nil
}

func (p *ActorParser) enrollCronEvent(raw []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) updateClaimedPower(raw []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) updatePledgeTotal(raw []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) networkRawPower(rawReturn []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) minerRawPower(raw, rawReturn []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) minerCount(rawReturn []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) minerConsensusCount(rawReturn []byte) (map[string]interface{}, error) {
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
