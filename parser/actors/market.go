package actors

import (
	"bytes"
	"encoding/base64"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/builtin/v11/market"
	v9Market "github.com/filecoin-project/go-state-types/builtin/v9/market" // v0.10.0 does not support ComputeDataCommitmentParams and OnMinerSectorsTerminateParams on v11
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/zondax/fil-parser/parser"
)

func ParseStoragemarket(txType string, msg *parser.LotusMessage, msgRct *filTypes.MessageReceipt) (map[string]interface{}, error) {
	switch txType {
	case parser.MethodSend:
		return parseSend(msg), nil
	case parser.MethodConstructor:
		return emptyParamsAndReturn()
	case parser.MethodAddBalance, parser.MethodAddBalanceExported:
		return addBalance(msg.Params)
	case parser.MethodWithdrawBalance, parser.MethodWithdrawBalanceExported:
		return withdrawBalance(msg.Params, msgRct.Return)
	case parser.MethodPublishStorageDeals, parser.MethodPublishStorageDealsExported:
		return publishStorageDeals(msg.Params, msgRct.Return)
	case parser.MethodVerifyDealsForActivation:
		return verifyDealsForActivation(msg.Params, msgRct.Return)
	case parser.MethodActivateDeals:
		return activateDeals(msg.Params)
	case parser.MethodOnMinerSectorsTerminate:
		return onMinerSectorsTerminate(msg.Params)
	case parser.MethodComputeDataCommitment:
		return computeDataCommitment(msg.Params, msgRct.Return)
	case parser.MethodCronTick:
		return emptyParamsAndReturn()
	case parser.MethodGetBalance:
		return getBalance(msg.Params, msgRct.Return)
	case parser.MethodGetDealDataCommitment:
		return getDealDataCommitment(msg.Params, msgRct.Return)
	case parser.MethodGetDealClient:
		return getDealClient(msg.Params, msgRct.Return)
	case parser.MethodGetDealProvider:
		return getDealProvider(msg.Params, msgRct.Return)
	case parser.MethodGetDealLabel:
		return getDealLabel(msg.Params, msgRct.Return)
	case parser.MethodGetDealTerm:
		return getDealTerm(msg.Params, msgRct.Return)
	case parser.MethodGetDealTotalPrice:
		return getDealTotalPrice(msg.Params, msgRct.Return)
	case parser.MethodGetDealClientCollateral:
		return getDealClientCollateral(msg.Params, msgRct.Return)
	case parser.MethodGetDealProviderCollateral:
		return getDealProviderCollateral(msg.Params, msgRct.Return)
	case parser.MethodGetDealVerified:
		return getDealVerified(msg.Params, msgRct.Return)
	case parser.MethodGetDealActivation:
		return getDealActivation(msg.Params, msgRct.Return)
	case parser.UnknownStr:
		return unknownMetadata(msg.Params, msgRct.Return)
	}
	return map[string]interface{}{}, parser.ErrUnknownMethod
}

func addBalance(raw []byte) (map[string]interface{}, error) {
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

func withdrawBalance(raw, rawReturn []byte) (map[string]interface{}, error) {
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

func publishStorageDeals(rawParams, rawReturn []byte) (map[string]interface{}, error) {
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

func verifyDealsForActivation(rawParams, rawReturn []byte) (map[string]interface{}, error) {
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

func activateDeals(rawParams []byte) (map[string]interface{}, error) {
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

func onMinerSectorsTerminate(rawParams []byte) (map[string]interface{}, error) {
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

func computeDataCommitment(rawParams, rawReturn []byte) (map[string]interface{}, error) {
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

func getBalance(rawParams, rawReturn []byte) (map[string]interface{}, error) {
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

func getDealDataCommitment(rawParams, rawReturn []byte) (map[string]interface{}, error) {
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

func getDealClient(rawParams, rawReturn []byte) (map[string]interface{}, error) {
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

func getDealProvider(rawParams, rawReturn []byte) (map[string]interface{}, error) {
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

func getDealLabel(rawParams, rawReturn []byte) (map[string]interface{}, error) {
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

func getDealTerm(rawParams, rawReturn []byte) (map[string]interface{}, error) {
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

func getDealTotalPrice(rawParams, rawReturn []byte) (map[string]interface{}, error) {
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

func getDealClientCollateral(rawParams, rawReturn []byte) (map[string]interface{}, error) {
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

func getDealProviderCollateral(rawParams, rawReturn []byte) (map[string]interface{}, error) {
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

func getDealVerified(rawParams, rawReturn []byte) (map[string]interface{}, error) {
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

func getDealActivation(rawParams, rawReturn []byte) (map[string]interface{}, error) {
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
