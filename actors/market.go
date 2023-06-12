package actors

import (
	"bytes"
	"encoding/base64"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/builtin/v11/market"
	v9Market "github.com/filecoin-project/go-state-types/builtin/v9/market" // v0.10.0 does not support ComputeDataCommitmentParams and OnMinerSectorsTerminateParams on v11
	"github.com/zondax/fil-parser/parser"
)

func (p *ActorParser) ParseStoragemarket(txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt) (map[string]interface{}, error) {
	switch txType {
	case parser.MethodSend:
		return p.parseSend(msg), nil
	case parser.MethodConstructor:
		return p.emptyParamsAndReturn()
	case parser.MethodAddBalance, parser.MethodAddBalanceExported:
		return p.addBalance(msg.Params)
	case parser.MethodWithdrawBalance, parser.MethodWithdrawBalanceExported:
		return p.withdrawBalance(msg.Params, msgRct.Return)
	case parser.MethodPublishStorageDeals, parser.MethodPublishStorageDealsExported:
		return p.publishStorageDeals(msg.Params, msgRct.Return)
	case parser.MethodVerifyDealsForActivation:
		return p.verifyDealsForActivation(msg.Params, msgRct.Return)
	case parser.MethodActivateDeals:
		return p.activateDeals(msg.Params)
	case parser.MethodOnMinerSectorsTerminate:
		return p.onMinerSectorsTerminate(msg.Params)
	case parser.MethodComputeDataCommitment:
		return p.computeDataCommitment(msg.Params, msgRct.Return)
	case parser.MethodCronTick:
		return p.emptyParamsAndReturn()
	case parser.MethodGetBalance:
		return p.getBalance(msg.Params, msgRct.Return)
	case parser.MethodGetDealDataCommitment:
		return p.getDealDataCommitment(msg.Params, msgRct.Return)
	case parser.MethodGetDealClient:
		return p.getDealClient(msg.Params, msgRct.Return)
	case parser.MethodGetDealProvider:
		return p.getDealProvider(msg.Params, msgRct.Return)
	case parser.MethodGetDealLabel:
		return p.getDealLabel(msg.Params, msgRct.Return)
	case parser.MethodGetDealTerm:
		return p.getDealTerm(msg.Params, msgRct.Return)
	case parser.MethodGetDealTotalPrice:
		return p.getDealTotalPrice(msg.Params, msgRct.Return)
	case parser.MethodGetDealClientCollateral:
		return p.getDealClientCollateral(msg.Params, msgRct.Return)
	case parser.MethodGetDealProviderCollateral:
		return p.getDealProviderCollateral(msg.Params, msgRct.Return)
	case parser.MethodGetDealVerified:
		return p.getDealVerified(msg.Params, msgRct.Return)
	case parser.MethodGetDealActivation:
		return p.getDealActivation(msg.Params, msgRct.Return)
	case parser.UnknownStr:
		return p.unknownMetadata(msg.Params, msgRct.Return)
	}
	return map[string]interface{}{}, parser.ErrUnknownMethod
}

func (p *ActorParser) addBalance(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params address.Address
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params.String()
	return metadata, nil
}

func (p *ActorParser) withdrawBalance(raw, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params market.WithdrawBalanceParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	if rawReturn != nil {
		metadata[parser.ReturnKey] = base64.StdEncoding.EncodeToString(rawReturn)
	}
	return metadata, nil
}

func (p *ActorParser) publishStorageDeals(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params market.PublishStorageDealsParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var publishReturn market.PublishStorageDealsReturn
	err = publishReturn.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = publishReturn
	return metadata, nil
}

func (p *ActorParser) verifyDealsForActivation(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params market.VerifyDealsForActivationParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var dealsReturn market.VerifyDealsForActivationReturn
	err = dealsReturn.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = dealsReturn
	return metadata, nil
}

func (p *ActorParser) activateDeals(rawParams []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params market.ActivateDealsParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func (p *ActorParser) onMinerSectorsTerminate(rawParams []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params v9Market.OnMinerSectorsTerminateParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func (p *ActorParser) computeDataCommitment(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params v9Market.ComputeDataCommitmentParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var computeReturn market.ComputeDataCommitmentReturn
	err = computeReturn.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = computeReturn
	return metadata, nil
}

func (p *ActorParser) getBalance(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params address.Address
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params.String()
	reader = bytes.NewReader(rawReturn)
	var r market.GetBalanceReturn
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r
	return metadata, nil
}

func (p *ActorParser) getDealDataCommitment(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params market.GetDealDataCommitmentParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var r market.GetDealDataCommitmentReturn
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r
	return metadata, nil
}

func (p *ActorParser) getDealClient(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params market.GetDealClientParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var r market.GetDealClientReturn
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r
	return metadata, nil
}

func (p *ActorParser) getDealProvider(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params market.GetDealProviderParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var r market.GetDealProviderReturn
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r
	return metadata, nil
}

func (p *ActorParser) getDealLabel(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params market.GetDealLabelParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var r market.GetDealLabelReturn
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r
	return metadata, nil
}

func (p *ActorParser) getDealTerm(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params market.GetDealTermParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var r market.GetDealTermReturn
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r
	return metadata, nil
}

func (p *ActorParser) getDealTotalPrice(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params market.GetDealTotalPriceParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var r market.GetDealTotalPriceReturn
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r
	return metadata, nil
}

func (p *ActorParser) getDealClientCollateral(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params market.GetDealClientCollateralParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var r market.GetDealClientCollateralReturn
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r
	return metadata, nil
}

func (p *ActorParser) getDealProviderCollateral(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params market.GetDealProviderCollateralParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var r market.GetDealProviderCollateralReturn
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r
	return metadata, nil
}

func (p *ActorParser) getDealVerified(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params market.GetDealVerifiedParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var r market.GetDealVerifiedReturn
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r
	return metadata, nil
}

func (p *ActorParser) getDealActivation(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params market.GetDealActivationParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var r market.GetDealActivationReturn
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r
	return metadata, nil
}
