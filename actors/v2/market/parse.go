package market

import (
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
)

func (p *Market) Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, _ cid.Cid) (map[string]interface{}, *types.AddressInfo, error) {
	switch txType {
	case parser.MethodSend:
		// return p.parseSend(msg), nil
	case parser.MethodConstructor:
		// return p.emptyParamsAndReturn()
	case parser.MethodAddBalance, parser.MethodAddBalanceExported:
		resp, err := p.AddBalance(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodWithdrawBalance, parser.MethodWithdrawBalanceExported:
		resp, err := p.WithdrawBalance(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodPublishStorageDeals, parser.MethodPublishStorageDealsExported:
		resp, err := p.PublishStorageDealsExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodVerifyDealsForActivation:
		resp, err := p.VerifyDealsForActivationExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodActivateDeals:
		resp, err := p.ActivateDealsExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodOnMinerSectorsTerminate:
		resp, err := p.OnMinerSectorsTerminateExported(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodComputeDataCommitment:
		resp, err := p.ComputeDataCommitmentExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodCronTick:
		// return p.emptyParamsAndReturn()
	case parser.MethodGetBalance:
		resp, err := p.GetBalanceExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodGetDealDataCommitment:
		resp, err := p.GetDealDataCommitmentExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodGetDealClient:
		resp, err := p.GetDealClientExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodGetDealProvider:
		resp, err := p.GetDealProviderExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodGetDealLabel:
		resp, err := p.GetDealLabelExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodGetDealTerm:
		resp, err := p.GetDealTermExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodGetDealTotalPrice:
		resp, err := p.GetDealTotalPriceExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodGetDealClientCollateral:
		resp, err := p.GetDealClientCollateralExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodGetDealProviderCollateral:
		resp, err := p.GetDealProviderCollateralExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodGetDealVerified:
		resp, err := p.GetDealVerifiedExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodGetDealActivation:
		resp, err := p.GetDealActivationExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.UnknownStr:
		// return p.unknownMetadata(msg.Params, msgRct.Return)
	}
	return map[string]interface{}{}, nil, parser.ErrUnknownMethod
}

func (p *Market) TransactionTypes() map[string]any {
	return map[string]any{
		parser.MethodSend:                        nil,
		parser.MethodConstructor:                 nil,
		parser.MethodAddBalance:                  p.AddBalance,
		parser.MethodAddBalanceExported:          p.AddBalance,
		parser.MethodWithdrawBalance:             p.WithdrawBalance,
		parser.MethodWithdrawBalanceExported:     p.WithdrawBalance,
		parser.MethodPublishStorageDeals:         p.PublishStorageDealsExported,
		parser.MethodPublishStorageDealsExported: p.PublishStorageDealsExported,
		parser.MethodVerifyDealsForActivation:    p.VerifyDealsForActivationExported,
		parser.MethodActivateDeals:               p.ActivateDealsExported,
		parser.MethodOnMinerSectorsTerminate:     p.OnMinerSectorsTerminateExported,
		parser.MethodComputeDataCommitment:       p.ComputeDataCommitmentExported,
		parser.MethodCronTick:                    nil,
		parser.MethodGetBalance:                  p.GetBalanceExported,
		parser.MethodGetDealDataCommitment:       p.GetDealDataCommitmentExported,
		parser.MethodGetDealClient:               p.GetDealClientExported,
		parser.MethodGetDealProvider:             p.GetDealProviderExported,
		parser.MethodGetDealLabel:                p.GetDealLabelExported,
		parser.MethodGetDealTerm:                 p.GetDealTermExported,
		parser.MethodGetDealTotalPrice:           p.GetDealTotalPriceExported,
		parser.MethodGetDealClientCollateral:     p.GetDealClientCollateralExported,
		parser.MethodGetDealProviderCollateral:   p.GetDealProviderCollateralExported,
		parser.MethodGetDealVerified:             p.GetDealVerifiedExported,
		parser.MethodGetDealActivation:           p.GetDealActivationExported,
	}
}
