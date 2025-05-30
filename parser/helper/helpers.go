package helper

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/zondax/golem/pkg/logger"

	"github.com/ipfs/go-cid"
	// The following import is necessary to ensure that the init() function
	// from the lotus build package is invoked.
	// In a recent refactor (v1.30.0), some build packages were modularized to reduce
	// unnecessary dependencies. As a result, if this package is not explicitly
	// imported, its init() will not be triggered, potentially causing issues
	// with initialization, such as errors when searching for actorNameByCid.
	_ "github.com/filecoin-project/lotus/build"

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

	"github.com/zondax/fil-parser/metrics"
	"github.com/zondax/fil-parser/tools"
	rosettaFilecoinLib "github.com/zondax/rosetta-filecoin-lib"
	"github.com/zondax/rosetta-filecoin-lib/actors"

	"github.com/zondax/fil-parser/actors/cache"
	logger2 "github.com/zondax/fil-parser/logger"
	"github.com/zondax/fil-parser/parser"
	parsermetrics "github.com/zondax/fil-parser/parser/metrics"
	"github.com/zondax/fil-parser/types"
)

// Deprecated: Use v2/tools.ActorMethods instead
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
	actorCache cache.IActorsCache
	logger     *logger.Logger
	metrics    *parsermetrics.ParserMetricsClient
	network    string
}

func NewHelper(lib *rosettaFilecoinLib.RosettaConstructionFilecoin, actorsCache cache.IActorsCache, node api.FullNode, logger *logger.Logger, metrics metrics.MetricsClient) *Helper {
	h := &Helper{
		lib:        lib,
		actorCache: actorsCache,
		node:       node,
		logger:     logger2.GetSafeLogger(logger),
		metrics:    parsermetrics.NewClient(metrics, "helper"),
	}
	network, err := h.node.StateNetworkName(context.Background())
	if err != nil {
		h.logger.Errorf("could not get network name: %v", err)
		return nil
	}
	h.network = tools.ParseRawNetworkName(string(network))
	return h
}

func (h *Helper) GetActorsCache() cache.IActorsCache {
	return h.actorCache
}

func (h *Helper) GetFilecoinLib() *rosettaFilecoinLib.RosettaConstructionFilecoin {
	return h.lib
}

func (h *Helper) GetFilecoinNodeClient() api.FullNode {
	return h.node
}

func (h *Helper) GetActorAddressInfo(add address.Address, key filTypes.TipSetKey, height abi.ChainEpoch) *types.AddressInfo {
	var err error
	addInfo := &types.AddressInfo{}

	if add == address.Undef {
		return addInfo
	}

	version := tools.VersionFromHeight(h.network, int64(height))
	addInfo.ActorCid, err = h.actorCache.GetActorCode(add, key, false)
	if err != nil {
		h.logger.Errorf("could not get actor code from address. Err: %s", err)
	} else {
		c, err := cid.Parse(addInfo.ActorCid)
		if err != nil {
			h.logger.Errorf("Could not parse params. Cannot cid.parse actor code: %v", err)
		}
		addInfo.ActorType, _ = h.lib.BuiltinActors.GetActorNameFromCidByVersion(c, version.FilNetworkVersion())
	}

	addInfo.Short, err = h.actorCache.GetShortAddress(add)
	if err != nil {
		h.logger.Errorf("could not get short address for %s. Err: %v", add.String(), err)
	}

	addInfo.Robust, err = h.actorCache.GetRobustAddress(add)
	if err != nil {
		h.logger.Errorf("could not get robust address for %s. Err: %v", add.String(), err)
	}

	addInfo.IsSystemActor = h.IsSystemActor(add) || h.IsGenesisActor(add)

	return addInfo
}

func (h *Helper) GetActorNameFromAddress(add address.Address, height int64, key filTypes.TipSetKey) (cid.Cid, string, error) {
	if add == address.Undef {
		return cid.Undef, "", errors.New("address is undefined")
	}

	onChainOnly := false
	for {
		// Search for actor in cache
		actorCode, err := h.actorCache.GetActorCode(add, key, onChainOnly)
		if err != nil {
			return cid.Undef, actors.UnknownStr, err
		}

		c, err := cid.Parse(actorCode)
		if err != nil {
			h.logger.Errorf("Could not parse params. Cannot cid.parse actor code: %v", err)
			return cid.Undef, actors.UnknownStr, err
		}

		actorName, err := h.GetActorNameFromCid(c, height)
		if err != nil {
			return cid.Undef, actors.UnknownStr, err
		}

		if actorName == manifest.PlaceholderKey && !onChainOnly {
			onChainOnly = true
		} else {
			return c, actorName, nil
		}
	}
}

func (h *Helper) GetActorNameFromCid(cid cid.Cid, height int64) (string, error) {
	version := tools.VersionFromHeight(h.network, height)
	actorName, err := h.lib.BuiltinActors.GetActorNameFromCidByVersion(cid, version.FilNetworkVersion())
	if err != nil {
		return "", err
	}
	return actorName, nil
}

// Deprecated: Use v2/tools.GetMethodName instead
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

	_, actorName, err := h.GetActorNameFromAddress(msg.To, height, key)
	if err != nil {
		_ = h.metrics.UpdateActorNameErrorMetric(fmt.Sprint(uint64(msg.Method)))
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

// CheckCommonMethods returns the method name for the given message if Send Or Constructor, otherwise returns an empty string
func (h *Helper) CheckCommonMethods(msg *parser.LotusMessage, height int64, key filTypes.TipSetKey) (string, error) {
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

	return "", nil
}

func (h *Helper) GetEVMSelectorSig(ctx context.Context, selectorID string) (string, error) {
	s, err := h.actorCache.GetEVMSelectorSig(ctx, selectorID)
	if err != nil {
		_ = h.metrics.UpdateGetEvmSelectorSigMetric()
	}
	return s, err
}

func (h *Helper) FilterTxsByActorType(ctx context.Context, txs []*types.Transaction, actorType string, tipsetKey filTypes.TipSetKey) ([]*types.Transaction, error) {
	var result []*types.Transaction
	for _, tx := range txs {
		addrTo, err := address.NewFromString(tx.TxTo)
		if err != nil {
			_ = h.metrics.UpdateParseAddressErrorMetric("to")
			h.logger.Errorf("could not parse address. Err: %s", err)
			continue
		}
		addrFrom, err := address.NewFromString(tx.TxFrom)
		if err != nil {
			_ = h.metrics.UpdateParseAddressErrorMetric("from")
			h.logger.Errorf("could not parse address. Err: %s", err)
			continue
		}

		// #nosec G115
		isType, err := h.isAnyAddressOfType(ctx, []address.Address{addrTo, addrFrom}, int64(tx.Height), tipsetKey, actorType)
		if err != nil {
			h.logger.Errorf("could not get actor type from address. Err: %s", err)
			continue
		}
		if !isType {
			continue
		}

		result = append(result, tx)
	}

	return result, nil
}

func (h *Helper) IsSystemActor(addr address.Address) bool {
	return h.actorCache.IsSystemActor(addr.String())
}
func (h *Helper) IsGenesisActor(addr address.Address) bool {
	return h.actorCache.IsGenesisActor(addr.String())
}

func (h *Helper) IsCronActor(height int64, addr address.Address, tipsetKey filTypes.TipSetKey) bool {
	_, actorName, err := h.GetActorNameFromAddress(addr, height, tipsetKey)
	if err != nil {
		return false
	}
	return strings.Contains(actorName, manifest.CronKey)
}

func (h *Helper) isAnyAddressOfType(_ context.Context, addresses []address.Address, height int64, key filTypes.TipSetKey, actorType string) (bool, error) {
	for _, addr := range addresses {
		if addr == address.Undef {
			continue
		}
		_, actorName, err := h.GetActorNameFromAddress(addr, height, key)
		if err != nil {
			return false, err
		}
		if strings.EqualFold(actorName, actorType) {
			return true, nil
		}
	}
	return false, nil
}
