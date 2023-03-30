package parser

import (
	"bytes"
	"encoding/base64"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/builtin/v11/market"
	v9Market "github.com/filecoin-project/go-state-types/builtin/v9/market" // v0.10.0 does not support ComputeDataCommitmentParams and OnMinerSectorsTerminateParams on v11
	filTypes "github.com/filecoin-project/lotus/chain/types"
)

func (p *Parser) parseStoragemarket(txType string, msg *filTypes.Message, msgRct *filTypes.MessageReceipt) (map[string]interface{}, error) {
	switch txType {
	case MethodSend:
		return p.parseSend(msg), nil
	case MethodConstructor:
		return p.emptyParamsAndReturn()
	case MethodAddBalance, MethodAddBalanceExported:
		return p.addBalance(msg.Params)
	case MethodWithdrawBalance, MethodWithdrawBalanceExported:
		return p.withdrawBalance(msg.Params, msgRct.Return)
	case MethodPublishStorageDeals, MethodPublishStorageDealsExported:
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
		return p.emptyParamsAndReturn()
	case MethodGetBalance:
		return p.getBalance(msg.Params, msgRct.Return)
	case MethodGetDealDataCommitment:
		return p.getDealDataCommitment(msg.Params, msgRct.Return)
	case MethodGetDealClient:
		return p.getDealClient(msg.Params, msgRct.Return)
	case MethodGetDealProvider:
		return p.getDealProvider(msg.Params, msgRct.Return)
	case MethodGetDealLabel:
		return p.getDealLabel(msg.Params, msgRct.Return)
	case MethodGetDealTerm:
		return p.getDealTerm(msg.Params, msgRct.Return)
	case MethodGetDealTotalPrice:
		return p.getDealTotalPrice(msg.Params, msgRct.Return)
	case MethodGetDealClientCollateral:
		return p.getDealClientCollateral(msg.Params, msgRct.Return)
	case MethodGetDealProviderCollateral:
		return p.getDealProviderCollateral(msg.Params, msgRct.Return)
	case MethodGetDealVerified:
		return p.getDealVerified(msg.Params, msgRct.Return)
	case MethodGetDealActivation:
		return p.getDealActivation(msg.Params, msgRct.Return)
	case UnknownStr:
		return p.unknownMetadata(msg.Params, msgRct.Return)
	}
	return map[string]interface{}{}, errUnknownMethod
}

func (p *Parser) addBalance(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params address.Address
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params.String()
	return metadata, nil
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
	var params v9Market.OnMinerSectorsTerminateParams
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
	var params v9Market.ComputeDataCommitmentParams
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

func (p *Parser) getBalance(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params address.Address
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params.String()
	reader = bytes.NewReader(rawReturn)
	var r market.GetBalanceReturn
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = r
	return metadata, nil
}

func (p *Parser) getDealDataCommitment(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params market.GetDealDataCommitmentParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var r market.GetDealDataCommitmentReturn
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = r
	return metadata, nil
}

func (p *Parser) getDealClient(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params market.GetDealClientParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var r market.GetDealClientReturn
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = r
	return metadata, nil
}

func (p *Parser) getDealProvider(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params market.GetDealProviderParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var r market.GetDealProviderReturn
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = r
	return metadata, nil
}

func (p *Parser) getDealLabel(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params market.GetDealLabelParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var r market.GetDealLabelReturn
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = r
	return metadata, nil
}

func (p *Parser) getDealTerm(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params market.GetDealTermParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var r market.GetDealTermReturn
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = r
	return metadata, nil
}

func (p *Parser) getDealTotalPrice(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params market.GetDealTotalPriceParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var r market.GetDealTotalPriceReturn
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = r
	return metadata, nil
}

func (p *Parser) getDealClientCollateral(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params market.GetDealClientCollateralParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var r market.GetDealClientCollateralReturn
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = r
	return metadata, nil
}

func (p *Parser) getDealProviderCollateral(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params market.GetDealProviderCollateralParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var r market.GetDealProviderCollateralReturn
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = r
	return metadata, nil
}

func (p *Parser) getDealVerified(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params market.GetDealVerifiedParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var r market.GetDealVerifiedReturn
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = r
	return metadata, nil
}

func (p *Parser) getDealActivation(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params market.GetDealActivationParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var r market.GetDealActivationReturn
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = r
	return metadata, nil
}
