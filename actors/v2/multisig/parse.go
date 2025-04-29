package multisig

import (
	"context"
	"fmt"

	"github.com/zondax/golem/pkg/logger"

	"github.com/filecoin-project/go-state-types/abi"
	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	legacyBuiltin "github.com/filecoin-project/specs-actors/actors/builtin"
	"github.com/ipfs/go-cid"

	multisigv10 "github.com/filecoin-project/go-state-types/builtin/v10/multisig"
	multisigv11 "github.com/filecoin-project/go-state-types/builtin/v11/multisig"
	multisigv12 "github.com/filecoin-project/go-state-types/builtin/v12/multisig"
	multisigv13 "github.com/filecoin-project/go-state-types/builtin/v13/multisig"
	multisigv14 "github.com/filecoin-project/go-state-types/builtin/v14/multisig"
	multisigv15 "github.com/filecoin-project/go-state-types/builtin/v15/multisig"
	multisigv16 "github.com/filecoin-project/go-state-types/builtin/v16/multisig"
	multisigv8 "github.com/filecoin-project/go-state-types/builtin/v8/multisig"
	multisigv9 "github.com/filecoin-project/go-state-types/builtin/v9/multisig"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/actors/metrics"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
)

type Msig struct {
	helper  *helper.Helper
	logger  *logger.Logger
	metrics *metrics.ActorsMetricsClient
}

func New(helper *helper.Helper, logger *logger.Logger, metrics *metrics.ActorsMetricsClient) *Msig {
	return &Msig{
		helper:  helper,
		logger:  logger,
		metrics: metrics,
	}
}

func (p *Msig) Name() string {
	return manifest.MultisigKey
}

func (*Msig) StartNetworkHeight() int64 {
	return tools.V1.Height()
}

func legacyMethods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	m := &Msig{}
	return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
		legacyBuiltin.MethodsMultisig.Constructor: {
			Name:   parser.MethodConstructor,
			Method: actors.ParseConstructor,
		},
		legacyBuiltin.MethodsMultisig.Propose: {
			Name:   parser.MethodPropose,
			Method: m.Propose,
		},
		legacyBuiltin.MethodsMultisig.Approve: {
			Name:   parser.MethodApprove,
			Method: m.Approve,
		},
		legacyBuiltin.MethodsMultisig.Cancel: {
			Name:   parser.MethodCancel,
			Method: m.Cancel,
		},
		legacyBuiltin.MethodsMultisig.AddSigner: {
			Name:   parser.MethodAddSigner,
			Method: m.MsigParams,
		},
		legacyBuiltin.MethodsMultisig.RemoveSigner: {
			Name:   parser.MethodRemoveSigner,
			Method: m.RemoveSigner,
		},
		legacyBuiltin.MethodsMultisig.SwapSigner: {
			Name:   parser.MethodSwapSigner,
			Method: m.MsigParams,
		},
		legacyBuiltin.MethodsMultisig.ChangeNumApprovalsThreshold: {
			Name:   parser.MethodChangeNumApprovalsThreshold,
			Method: m.ChangeNumApprovalsThreshold,
		},
		legacyBuiltin.MethodsMultisig.LockBalance: {
			Name:   parser.MethodLockBalance,
			Method: m.LockBalance,
		},
	}
}

var methods = map[string]map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
	tools.V1.String():  legacyMethods(),
	tools.V2.String():  legacyMethods(),
	tools.V3.String():  legacyMethods(),
	tools.V4.String():  legacyMethods(),
	tools.V5.String():  legacyMethods(),
	tools.V6.String():  legacyMethods(),
	tools.V7.String():  legacyMethods(),
	tools.V8.String():  legacyMethods(),
	tools.V9.String():  legacyMethods(),
	tools.V10.String(): legacyMethods(),
	tools.V11.String(): legacyMethods(),
	tools.V12.String(): legacyMethods(),
	tools.V13.String(): legacyMethods(),
	tools.V14.String(): legacyMethods(),
	tools.V15.String(): legacyMethods(),
	tools.V15.String(): legacyMethods(),
	tools.V16.String(): actors.CopyMethods(multisigv8.Methods),
	tools.V17.String(): actors.CopyMethods(multisigv9.Methods),
	tools.V18.String(): actors.CopyMethods(multisigv10.Methods),
	tools.V19.String(): actors.CopyMethods(multisigv11.Methods),
	tools.V20.String(): actors.CopyMethods(multisigv11.Methods),
	tools.V21.String(): actors.CopyMethods(multisigv12.Methods),
	tools.V22.String(): actors.CopyMethods(multisigv13.Methods),
	tools.V23.String(): actors.CopyMethods(multisigv14.Methods),
	tools.V24.String(): actors.CopyMethods(multisigv15.Methods),
	tools.V25.String(): actors.CopyMethods(multisigv16.Methods),
}

func (m *Msig) Methods(_ context.Context, network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error) {
	version := tools.VersionFromHeight(network, height)
	methods, ok := methods[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return methods, nil
}

/*
Still needs to parse:

	Receive
*/
func (p *Msig) Parse(_ context.Context, network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, _ cid.Cid, key filTypes.TipSetKey) (map[string]interface{}, *types.AddressInfo, error) {
	var ret map[string]interface{}
	var err error
	switch txType {
	case parser.MethodConstructor: // TODO: not tested
		ret, err = p.MsigConstructor(network, height, msg.Params)
	case parser.MethodSend:
		resp := actors.ParseSend(msg)
		return resp, nil, nil
	case parser.MethodPropose, parser.MethodProposeExported:
		ret, err = p.Propose(network, msg, height, txType, key, msg.Params, msgRct.Return, p.parseMsigParams)
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
