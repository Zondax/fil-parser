package helper

import (
	"context"
	"errors"
	"fmt"
	"strings"

	parsermetrics "github.com/zondax/fil-parser/parser/metrics"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/builtin/v12/account"
	"github.com/filecoin-project/go-state-types/builtin/v12/cron"
	"github.com/filecoin-project/go-state-types/builtin/v12/datacap"
	"github.com/filecoin-project/go-state-types/builtin/v12/eam"
	"github.com/filecoin-project/go-state-types/builtin/v12/evm"
	filInit "github.com/filecoin-project/go-state-types/builtin/v12/init"
	"github.com/filecoin-project/go-state-types/builtin/v12/market"
	"github.com/filecoin-project/go-state-types/builtin/v12/miner"
	"github.com/filecoin-project/go-state-types/builtin/v12/multisig"
	"github.com/filecoin-project/go-state-types/builtin/v12/paych"
	"github.com/filecoin-project/go-state-types/builtin/v12/power"
	"github.com/filecoin-project/go-state-types/builtin/v12/reward"
	"github.com/filecoin-project/go-state-types/builtin/v12/verifreg"
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/filecoin-project/lotus/api"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/actors/cache"
	logger2 "github.com/zondax/fil-parser/logger"
	"github.com/zondax/fil-parser/metrics"
	"github.com/zondax/fil-parser/parser"
	rosettaFilecoinLib "github.com/zondax/rosetta-filecoin-lib"
	"github.com/zondax/rosetta-filecoin-lib/actors"
	"go.uber.org/zap"

	// The following import is necessary to ensure that the init() function
	// from the lotus build package is invoked.
	// In a recent refactor (v1.30.0), some build packages were modularized to reduce
	// unnecessary dependencies. As a result, if this package is not explicitly
	// imported, its init() will not be triggered, potentially causing issues
	// with initialization, such as errors when searching for actorNameByCid.
	_ "github.com/filecoin-project/lotus/build"

	"github.com/zondax/fil-parser/types"
)

var allMethods = map[string]map[abi.MethodNum]builtin.MethodMeta{
	manifest.InitKey:     filInit.Methods,
	manifest.CronKey:     cron.Methods,
	manifest.AccountKey:  account.Methods,
	manifest.PowerKey:    power.Methods,
	manifest.MinerKey:    miner.Methods,
	manifest.MarketKey:   market.Methods,
	manifest.PaychKey:    paych.Methods,
	manifest.MultisigKey: multisig.Methods,
	manifest.RewardKey:   reward.Methods,
	manifest.VerifregKey: verifreg.Methods,
	manifest.EvmKey:      evm.Methods,
	manifest.EamKey:      eam.Methods,
	manifest.DatacapKey:  datacap.Methods,

	// EthAccount and Placeholder can receive tokens with Send and InvokeEVM methods
	// We set evm.Methods instead of empty array of methods. Therefore, we will be able to understand
	// this specific method (3844450837) - tx cid example: bafy2bzacedgmcvsp56ieciutvgwza2qpvz7pvbhhu4l5y5tdl35rwfnjn5buk
	manifest.PlaceholderKey: evm.Methods,
	manifest.EthAccountKey:  evm.Methods,
}

type Helper struct {
	lib        *rosettaFilecoinLib.RosettaConstructionFilecoin
	node       api.FullNode
	actorCache *cache.ActorsCache
	logger     *zap.Logger
	metrics    *parsermetrics.ParserMetricsClient
}

func NewHelper(lib *rosettaFilecoinLib.RosettaConstructionFilecoin, actorsCache *cache.ActorsCache, node api.FullNode, logger *zap.Logger, metrics metrics.MetricsClient) *Helper {
	return &Helper{
		lib:        lib,
		actorCache: actorsCache,
		node:       node,
		logger:     logger2.GetSafeLogger(logger),
		metrics:    parsermetrics.NewClient(metrics, "helper"),
	}
}

func (h *Helper) GetActorsCache() *cache.ActorsCache {
	return h.actorCache
}

func (h *Helper) GetFilecoinLib() *rosettaFilecoinLib.RosettaConstructionFilecoin {
	return h.lib
}

func (h *Helper) GetFilecoinNodeClient() api.FullNode {
	return h.node
}

func (h *Helper) GetActorAddressInfo(add address.Address, key filTypes.TipSetKey) *types.AddressInfo {
	var err error
	addInfo := &types.AddressInfo{}

	addInfo.ActorCid, err = h.actorCache.GetActorCode(add, key, false)
	if err != nil {
		h.logger.Sugar().Errorf("could not get actor code from address. Err: %s", err)
	} else {
		c, err := cid.Parse(addInfo.ActorCid)
		if err != nil {
			h.logger.Sugar().Errorf("Could not parse params. Cannot cid.parse actor code: %v", err)
		}
		addInfo.ActorType, _ = h.lib.BuiltinActors.GetActorNameFromCid(c)
	}

	addInfo.Short, err = h.actorCache.GetShortAddress(add)
	if err != nil {
		h.logger.Sugar().Errorf("could not get short address for %s. Err: %v", add.String(), err)
	}

	// Ignore searching robust addresses for Msig and miners
	if addInfo.ActorType == manifest.MinerKey || addInfo.ActorType == manifest.MultisigKey {
		return addInfo
	}

	addInfo.Robust, err = h.actorCache.GetRobustAddress(add)
	if err != nil {
		h.logger.Sugar().Errorf("could not get robust address for %s. Err: %v", add.String(), err)
	}

	return addInfo
}

func (h *Helper) GetActorNameFromAddress(address address.Address, height int64, key filTypes.TipSetKey) (string, error) {
	onChainOnly := false
	for {
		// Search for actor in cache
		actorCode, err := h.actorCache.GetActorCode(address, key, onChainOnly)
		if err != nil {
			return actors.UnknownStr, err
		}

		c, err := cid.Parse(actorCode)
		if err != nil {
			h.logger.Sugar().Errorf("Could not parse params. Cannot cid.parse actor code: %v", err)
			return actors.UnknownStr, err
		}

		actorName, err := h.lib.BuiltinActors.GetActorNameFromCid(c)
		if err != nil {
			return actors.UnknownStr, err
		}

		if actorName == manifest.PlaceholderKey && !onChainOnly {
			onChainOnly = true
		} else {
			return actorName, nil
		}
	}
}

func (h *Helper) GetMethodName(msg *parser.LotusMessage, height int64, key filTypes.TipSetKey) (string, error) {
	if msg == nil {
		return "", errors.New("malformed value")
	}

	// Shortcut 1 - Method "0" corresponds to "MethodSend"
	if msg.Method == 0 {
		return parser.MethodSend, nil
	}

	// Shortcut 2 - Method "1" corresponds to "MethodConstructor"
	if msg.Method == 1 {
		return parser.MethodConstructor, nil
	}

	actorName, err := h.GetActorNameFromAddress(msg.To, height, key)
	if err != nil {
		_ = h.metrics.UpdateActorNameErrorMetric(fmt.Sprint(uint64(msg.Method)), err)
	}

	actorMethods, ok := allMethods[actorName]
	if !ok {
		return "", parser.ErrNotKnownActor
	}

	method, ok := actorMethods[msg.Method]
	if !ok {
		return parser.UnknownStr, nil
	}

	return method.Name, nil
}

func (h *Helper) GetEVMSelectorSig(ctx context.Context, selectorID string) (string, error) {
	s, err := h.actorCache.GetEVMSelectorSig(ctx, selectorID)
	if err != nil {
		_ = h.metrics.UpdateGetEvmSelectorSigMetric(err)
	}
	return s, err
}

func (h *Helper) FilterTxsByActorType(ctx context.Context, txs []*types.Transaction, actorType string, tipsetKey filTypes.TipSetKey) ([]*types.Transaction, error) {
	var result []*types.Transaction
	for _, tx := range txs {
		addrTo, err := address.NewFromString(tx.TxTo)
		if err != nil {
			_ = h.metrics.UpdateParseAddressErrorMetric("to", "decode error")
			h.logger.Sugar().Errorf("could not parse address. Err: %s", err)
			continue
		}
		addrFrom, err := address.NewFromString(tx.TxFrom)
		if err != nil {
			_ = h.metrics.UpdateParseAddressErrorMetric("from", "decode error")
			h.logger.Sugar().Errorf("could not parse address. Err: %s", err)
			continue
		}

		isType, err := h.isAnyAddressOfType(ctx, []address.Address{addrTo, addrFrom}, int64(tx.Height), tipsetKey, actorType)
		if err != nil {
			h.logger.Sugar().Errorf("could not get actor type from address. Err: %s", err)
			continue
		}
		if !isType {
			continue
		}

		result = append(result, tx)
	}

	return result, nil
}

func (h *Helper) isAnyAddressOfType(_ context.Context, addresses []address.Address, height int64, key filTypes.TipSetKey, actorType string) (bool, error) {
	for _, addr := range addresses {
		actorName, err := h.GetActorNameFromAddress(addr, height, key)
		if err != nil {
			return false, err
		}
		if strings.EqualFold(actorName, actorType) {
			return true, nil
		}
	}
	return false, nil
}
