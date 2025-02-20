package multisig

import (
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/actors/metrics"
	metrics2 "github.com/zondax/fil-parser/metrics"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/fil-parser/types"
	"go.uber.org/zap"
)

type Msig struct {
	helper  *helper.Helper
	logger  *zap.Logger
	metrics *metrics.ActorsMetricsClient
}

func New(helper *helper.Helper, logger *zap.Logger, metrics *metrics.ActorsMetricsClient) *Msig {
	return &Msig{
		helper:  helper,
		logger:  logger,
		metrics: metrics,
	}
}

func (p *Msig) Name() string {
	return manifest.MultisigKey
}

/*
Still needs to parse:

	Receive
*/
func (p *Msig) Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, _ cid.Cid, key filTypes.TipSetKey) (map[string]interface{}, *types.AddressInfo, error) {
	var ret map[string]interface{}
	var err error
	switch txType {
	case parser.MethodConstructor: // TODO: not tested
		ret, err = p.MsigConstructor(network, height, msg.Params)
	case parser.MethodSend:
		resp := actors.ParseSend(msg)
		return resp, nil, nil
	case parser.MethodPropose, parser.MethodProposeExported:
		ret, err = p.Propose(network, msg, height, key, msg.Params, msgRct.Return, p.parseMsigParams)
	case parser.MethodApprove, parser.MethodApproveExported:
		ret, err = p.Approve(network, msg, height, key, msgRct.Return, p.parseMsigParams)
	case parser.MethodCancel, parser.MethodCancelExported:
		ret, err = p.Cancel(network, msg, height, key, msgRct.Return, p.parseMsigParams)
	case parser.MethodAddSigner, parser.MethodAddSignerExported, parser.MethodSwapSigner, parser.MethodSwapSignerExported:
		ret, err = p.MsigParams(network, msg, height, key, p.parseMsigParams)
	case parser.MethodRemoveSigner, parser.MethodRemoveSignerExported:
		ret, err = p.RemoveSigner(network, msg, height, key, msgRct.Return, p.parseMsigParams)
	case parser.MethodChangeNumApprovalsThreshold, parser.MethodChangeNumApprovalsThresholdExported:
		ret, err = p.ChangeNumApprovalsThreshold(network, msg, height, key, msg.Params, p.parseMsigParams)
	case parser.MethodLockBalance, parser.MethodLockBalanceExported:
		ret, err = p.LockBalance(network, msg, height, key, msg.Params, p.parseMsigParams)
	case parser.MethodMsigUniversalReceiverHook: // TODO: not tested
		ret, err = p.UniversalReceiverHook(network, msg, height, key, msgRct.Return, p.parseMsigParams)
	case parser.UnknownStr:
		resp, err := actors.ParseUnknownMetadata(msg.Params, msgRct.Return)
		return resp, nil, err
	}

	return ret, nil, err
}

func (p *Msig) TransactionTypes() map[string]any {
	return map[string]any{
		parser.MethodConstructor:                         p.MsigConstructor,
		parser.MethodSend:                                actors.ParseSend,
		parser.MethodPropose:                             p.Propose,
		parser.MethodProposeExported:                     p.Propose,
		parser.MethodApprove:                             p.Approve,
		parser.MethodApproveExported:                     p.Approve,
		parser.MethodCancel:                              p.Cancel,
		parser.MethodCancelExported:                      p.Cancel,
		parser.MethodAddSigner:                           p.MsigParams,
		parser.MethodAddSignerExported:                   p.MsigParams,
		parser.MethodSwapSigner:                          p.MsigParams,
		parser.MethodSwapSignerExported:                  p.MsigParams,
		parser.MethodRemoveSigner:                        p.RemoveSigner,
		parser.MethodRemoveSignerExported:                p.RemoveSigner,
		parser.MethodChangeNumApprovalsThreshold:         p.ChangeNumApprovalsThreshold,
		parser.MethodChangeNumApprovalsThresholdExported: p.ChangeNumApprovalsThreshold,
		parser.MethodLockBalance:                         p.LockBalance,
		parser.MethodLockBalanceExported:                 p.LockBalance,
		parser.MethodMsigUniversalReceiverHook:           p.UniversalReceiverHook,
	}
}
