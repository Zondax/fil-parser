package market

import (
	"github.com/zondax/fil-parser/parser"
)

func (p *Market) Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt) (map[string]interface{}, error) {
	switch txType {
	case parser.MethodSend:
		// return p.parseSend(msg), nil
	case parser.MethodConstructor:
		// return p.emptyParamsAndReturn()
	case parser.MethodAddBalance, parser.MethodAddBalanceExported:
		return p.AddBalance(network, height, msg.Params)
	case parser.MethodWithdrawBalance, parser.MethodWithdrawBalanceExported:
		return p.WithdrawBalance(network, height, msg.Params)
	case parser.MethodPublishStorageDeals, parser.MethodPublishStorageDealsExported:
		return p.PublishStorageDealsExported(network, height, msg.Params, msgRct.Return)
	case parser.MethodVerifyDealsForActivation:
		return p.VerifyDealsForActivationExported(network, height, msg.Params, msgRct.Return)
	case parser.MethodActivateDeals:
		return p.ActivateDealsExported(network, height, msg.Params, msgRct.Return)
	case parser.MethodOnMinerSectorsTerminate:
		return p.OnMinerSectorsTerminateExported(network, height, msg.Params)
	case parser.MethodComputeDataCommitment:
		return p.ComputeDataCommitmentExported(network, height, msg.Params, msgRct.Return)
	case parser.MethodCronTick:
		// return p.emptyParamsAndReturn()
	case parser.MethodGetBalance:
		return p.GetBalanceExported(network, height, msg.Params, msgRct.Return)
	case parser.MethodGetDealDataCommitment:
		return p.GetDealDataCommitmentExported(network, height, msg.Params, msgRct.Return)
	case parser.MethodGetDealClient:
		return p.GetDealClientExported(network, height, msg.Params, msgRct.Return)
	case parser.MethodGetDealProvider:
		return p.GetDealProviderExported(network, height, msg.Params, msgRct.Return)
	case parser.MethodGetDealLabel:
		return p.GetDealLabelExported(network, height, msg.Params, msgRct.Return)
	case parser.MethodGetDealTerm:
		return p.GetDealTermExported(network, height, msg.Params, msgRct.Return)
	case parser.MethodGetDealTotalPrice:
		return p.GetDealTotalPriceExported(network, height, msg.Params, msgRct.Return)
	case parser.MethodGetDealClientCollateral:
		return p.GetDealClientCollateralExported(network, height, msg.Params, msgRct.Return)
	case parser.MethodGetDealProviderCollateral:
		return p.GetDealProviderCollateralExported(network, height, msg.Params, msgRct.Return)
	case parser.MethodGetDealVerified:
		return p.GetDealVerifiedExported(network, height, msg.Params, msgRct.Return)
	case parser.MethodGetDealActivation:
		return p.GetDealActivationExported(network, height, msg.Params, msgRct.Return)
	case parser.UnknownStr:
		// return p.unknownMetadata(msg.Params, msgRct.Return)
	}
	return map[string]interface{}{}, parser.ErrUnknownMethod
}
