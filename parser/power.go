package parser

import (
	"bytes"
	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/go-state-types/builtin/v11/power"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/specs-actors/actors/runtime/proof"

	"github.com/zondax/fil-parser/types"
)

func (p *Parser) parseStoragepower(txType string, msg *filTypes.Message, msgRct *filTypes.MessageReceipt) (map[string]interface{}, error) {
	switch txType {
	case MethodSend:
		return p.parseSend(msg), nil
	case MethodConstructor:
		return p.powerConstructor(msg.Params)
	case MethodCreateMiner, MethodCreateMinerExported:
		return p.parseCreateMiner(msg, msgRct)
	case MethodUpdateClaimedPower:
		return p.updateClaimedPower(msg.Params)
	case MethodEnrollCronEvent:
		return p.enrollCronEvent(msg.Params)
	case MethodCronTick:
		return p.emptyParamsAndReturn()
	case MethodUpdatePledgeTotal:
		return p.updatePledgeTotal(msg.Params)
	case MethodPowerDeprecated1: // OnConsensusFault
	case MethodSubmitPoRepForBulkVerify:
		return p.submitPoRepForBulkVerify(msg.Params)
	case MethodCurrentTotalPower:
		return p.currentTotalPower(msgRct.Return)
	case MethodNetworkRawPowerExported:
		return p.networkRawPower(msgRct.Return)
	case MethodMinerRawPowerExported:
		return p.minerRawPower(msg.Params, msgRct.Return)
	case MethodMinerCountExported:
		return p.minerCount(msgRct.Return)
	case MethodMinerConsensusCountExported:
		return p.minerConsensusCount(msgRct.Return)
	case UnknownStr:
		return p.unknownMetadata(msg.Params, msgRct.Return)

	}
	return map[string]interface{}{}, errUnknownMethod
}

func (p *Parser) currentTotalPower(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params power.CurrentTotalPowerReturn
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = params
	return metadata, nil
}

func (p *Parser) submitPoRepForBulkVerify(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params proof.SealVerifyInfo
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) powerConstructor(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params power.MinerConstructorParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) parseCreateMiner(msg *filTypes.Message, msgRct *filTypes.MessageReceipt) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(msg.Params)
	var params power.CreateMinerParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params

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
		CreationTxHash: msg.Cid().String(),
	}
	metadata[ReturnKey] = createdActor
	p.appendToAddresses(*createdActor)
	return metadata, nil
}

func (p *Parser) enrollCronEvent(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params power.EnrollCronEventParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) updateClaimedPower(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params power.UpdateClaimedPowerParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) updatePledgeTotal(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params abi.TokenAmount
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) networkRawPower(rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawReturn)
	var r power.NetworkRawPowerReturn
	err := r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = r
	return metadata, nil
}

func (p *Parser) minerRawPower(raw, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params power.MinerRawPowerParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params

	reader = bytes.NewReader(rawReturn)
	var r power.MinerRawPowerReturn
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = r
	return metadata, nil
}

func (p *Parser) minerCount(rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawReturn)
	var r power.MinerCountReturn
	err := r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = r
	return metadata, nil
}

func (p *Parser) minerConsensusCount(rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawReturn)
	var r power.MinerConsensusCountReturn
	err := r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = r
	return metadata, nil
}
