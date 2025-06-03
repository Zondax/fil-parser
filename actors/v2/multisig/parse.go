package multisig

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/zondax/golem/pkg/logger"

	"github.com/filecoin-project/go-state-types/abi"
	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
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
	"github.com/zondax/fil-parser/actors/v2/evm"
	"github.com/zondax/fil-parser/actors/v2/miner"
	"github.com/zondax/fil-parser/actors/v2/verifiedRegistry"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
)

type Msig struct {
	helper   *helper.Helper
	logger   *logger.Logger
	metrics  *metrics.ActorsMetricsClient
	miner    *miner.Miner
	verifreg *verifiedRegistry.VerifiedRegistry
	evm      *evm.Evm

	methodNameFn actors.MethodNameFn
}

func New(helper *helper.Helper, logger *logger.Logger, metrics *metrics.ActorsMetricsClient, methodNameFn actors.MethodNameFn) *Msig {
	return &Msig{
		helper:       helper,
		logger:       logger,
		metrics:      metrics,
		miner:        miner.New(logger),
		verifreg:     verifiedRegistry.New(logger),
		evm:          evm.New(logger, metrics),
		methodNameFn: methodNameFn,
	}
}

func (p *Msig) Name() string {
	return manifest.MultisigKey
}

func (*Msig) StartNetworkHeight() int64 {
	return tools.V1.Height()
}

// All methods can be found in the Actor.Exports method in
// the correct version package for "github.com/filecoin-project/specs-actors/actors/builtin/multisig"

func v1Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	m := &Msig{}
	return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
		1: {
			Name:   parser.MethodConstructor,
			Method: actors.ParseConstructor,
		},
		2: {
			Name:   parser.MethodPropose,
			Method: m.Propose,
		},
		3: {
			Name:   parser.MethodApprove,
			Method: m.Approve,
		},
		4: {
			Name:   parser.MethodCancel,
			Method: m.Cancel,
		},
		5: {
			Name:   parser.MethodAddSigner,
			Method: m.AddSigner,
		},
		6: {
			Name:   parser.MethodRemoveSigner,
			Method: m.RemoveSigner,
		},
		7: {
			Name:   parser.MethodSwapSigner,
			Method: m.SwapSigner,
		},
		8: {
			Name:   parser.MethodChangeNumApprovalsThreshold,
			Method: m.ChangeNumApprovalsThreshold,
		},
		9: {
			Name:   parser.MethodLockBalance,
			Method: m.LockBalance,
		},
	}
}
func v2Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return v1Methods()
}
func v3Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return v1Methods()
}
func v4Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return v1Methods()
}
func v5Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return v1Methods()
}
func v6Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return v1Methods()
}
func v7Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return v1Methods()
}

var methods = map[string]map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
	tools.V0.String(): v1Methods(),
	tools.V1.String(): v1Methods(),
	tools.V2.String(): v1Methods(),
	tools.V3.String(): v1Methods(),

	tools.V4.String(): v2Methods(),
	tools.V5.String(): v2Methods(),
	tools.V6.String(): v2Methods(),
	tools.V7.String(): v2Methods(),
	tools.V8.String(): v2Methods(),
	tools.V9.String(): v2Methods(),

	tools.V10.String(): v3Methods(),
	tools.V11.String(): v3Methods(),

	tools.V12.String(): v4Methods(),
	tools.V13.String(): v5Methods(),
	tools.V14.String(): v6Methods(),
	tools.V15.String(): v7Methods(),

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
		ret, err = p.Propose(network, msg, height, txType, key, msg.Params, msgRct.Return)
	case parser.MethodApprove, parser.MethodApproveExported:
		ret, err = p.Approve(network, msg, height, key, msg.Params, msgRct.Return)
	case parser.MethodCancel, parser.MethodCancelExported:
		ret, err = p.Cancel(network, msg, height, key, msg.Params)
	case parser.MethodAddSigner, parser.MethodAddSignerExported:
		ret, err = p.AddSigner(network, msg, height, key, msg.Params, msgRct.Return)
	case parser.MethodRemoveSigner, parser.MethodRemoveSignerExported:
		ret, err = p.RemoveSigner(network, msg, height, key, msg.Params)
	case parser.MethodSwapSigner, parser.MethodSwapSignerExported:
		ret, err = p.SwapSigner(network, msg, height, key, msg.Params, msgRct.Return)
	case parser.MethodChangeNumApprovalsThreshold, parser.MethodChangeNumApprovalsThresholdExported:
		ret, err = p.ChangeNumApprovalsThreshold(network, msg, height, key, msg.Params)
	case parser.MethodLockBalance, parser.MethodLockBalanceExported:
		ret, err = p.LockBalance(network, msg, height, key, msg.Params)
	case parser.MethodAddVerifier:
		ret, err = p.AddVerifier(network, msg, height, key, msg.Params, msgRct.Return)
	case parser.MethodChangeOwnerAddress, parser.MethodChangeOwnerAddressExported:
		ret, err = p.ChangeOwnerAddress(network, msg, height, key, msg.Params, msgRct.Return)
	case parser.MethodWithdrawBalance, parser.MethodWithdrawBalanceExported:
		ret, err = p.WithdrawBalance(network, msg, height, key, msg.Params, msgRct.Return)
	case parser.MethodInvokeContract:
		ret, err = p.InvokeContract(network, msg, height, key, msg.Params, msgRct.Return)
	case parser.MethodMsigUniversalReceiverHook: // TODO: not tested
		ret, err = p.UniversalReceiverHook(network, msg, height, key, msg.Params)
	case parser.MethodFallback:
		ret, err = p.Fallback(network, height, msg.Params)
	case parser.UnknownStr:
		resp, err := actors.ParseUnknownMetadata(msg.Params, msgRct.Return)
		return resp, nil, err
	}

	return ret, nil, err
}

func (p *Msig) Fallback(network string, height int64, raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	metadata[parser.ParamsRawKey] = hex.EncodeToString(raw)
	return metadata, nil
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
		parser.MethodAddSigner:                           p.AddSigner,
		parser.MethodAddSignerExported:                   p.AddSigner,
		parser.MethodSwapSigner:                          p.SwapSigner,
		parser.MethodSwapSignerExported:                  p.SwapSigner,
		parser.MethodRemoveSigner:                        p.RemoveSigner,
		parser.MethodRemoveSignerExported:                p.RemoveSigner,
		parser.MethodChangeNumApprovalsThreshold:         p.ChangeNumApprovalsThreshold,
		parser.MethodChangeNumApprovalsThresholdExported: p.ChangeNumApprovalsThreshold,
		parser.MethodLockBalance:                         p.LockBalance,
		parser.MethodLockBalanceExported:                 p.LockBalance,
		parser.MethodMsigUniversalReceiverHook:           p.UniversalReceiverHook,
	}
}
