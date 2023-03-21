package parser

import (
	"bytes"
	"encoding/base64"

	"github.com/filecoin-project/go-state-types/builtin/v9/market" // TODO: v10 does not support ComputeDataCommitmentParams and OnMinerSectorsTerminateParams
	filTypes "github.com/filecoin-project/lotus/chain/types"
)

func (p *Parser) parseStoragemarket(txType string, msg *filTypes.Message, msgRct *filTypes.MessageReceipt) (map[string]interface{}, error) {
	switch txType {
	case MethodSend:
		return p.parseSend(msg), nil
	case MethodConstructor:
	case MethodAddBalance:
	case MethodWithdrawBalance:
		return p.withdrawBalance(msg.Params, msgRct.Return)
	case MethodPublishStorageDeals:
		return p.publishStorageDeals(msg.Params, msgRct.Return)
	case MethodVerifyDealsForActivation:
		return p.verifyDealsForActivation(msg.Params, msgRct.Return)
	case MethodActivateDeals:
		return p.activateDeals(msg.Params)
	case MethodOnMinerSectorsTerminate:
		return p.onMinerSectorsTerminate(msg.Params)
	case MethodComputeDataCommitment:
		return p.computeDataCommitment(msg.Params, msgRct.Return)
	case MethodCronTick:
	case UnknownStr:
		return p.unknownMetadata(msg.Params, msgRct.Return)
	}
	return map[string]interface{}{}, errUnknownMethod
}

func (p *Parser) withdrawBalance(raw, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params market.WithdrawBalanceParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	if rawReturn != nil {
		metadata[ReturnKey] = base64.StdEncoding.EncodeToString(rawReturn)
	}
	return metadata, nil
}

func (p *Parser) publishStorageDeals(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params market.PublishStorageDealsParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var publishReturn market.PublishStorageDealsReturn
	err = publishReturn.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = publishReturn
	return metadata, nil
}

func (p *Parser) verifyDealsForActivation(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params market.VerifyDealsForActivationParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var dealsReturn market.VerifyDealsForActivationReturn
	err = dealsReturn.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = dealsReturn
	return metadata, nil
}

func (p *Parser) activateDeals(rawParams []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params market.ActivateDealsParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) onMinerSectorsTerminate(rawParams []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params market.OnMinerSectorsTerminateParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) computeDataCommitment(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params market.ComputeDataCommitmentParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var computeReturn market.ComputeDataCommitmentReturn
	err = computeReturn.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = computeReturn
	return metadata, nil
}
