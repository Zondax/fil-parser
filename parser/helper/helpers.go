package helper

import (
	"context"
	"errors"
	"strings"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/filecoin-project/lotus/api"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/actors/cache"
	logger2 "github.com/zondax/fil-parser/logger"
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

type Helper struct {
	lib        *rosettaFilecoinLib.RosettaConstructionFilecoin
	node       api.FullNode
	actorCache *cache.ActorsCache
	logger     *zap.Logger
}

func NewHelper(lib *rosettaFilecoinLib.RosettaConstructionFilecoin, actorsCache *cache.ActorsCache, node api.FullNode, logger *zap.Logger) *Helper {
	return &Helper{lib: lib, actorCache: actorsCache, node: node, logger: logger2.GetSafeLogger(logger)}
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

	actorName, _ := h.GetActorNameFromAddress(msg.To, height, key)

	actorMethods, err := allMethods(height, "", actorName)
	if err != nil {
		return "", err
	}

	method, ok := actorMethods[msg.Method]
	if !ok {
		return parser.UnknownStr, nil
	}
	return method.Name, nil
}

func (h *Helper) GetEVMSelectorSig(ctx context.Context, selectorID string) (string, error) {
	return h.actorCache.GetEVMSelectorSig(ctx, selectorID)
}

func (h *Helper) FilterTxsByActorType(ctx context.Context, txs []*types.Transaction, actorType string, tipsetKey filTypes.TipSetKey) ([]*types.Transaction, error) {
	var result []*types.Transaction
	for _, tx := range txs {
		addrTo, err := address.NewFromString(tx.TxTo)
		if err != nil {
			h.logger.Sugar().Errorf("could not parse address. Err: %s", err)
			continue
		}
		addrFrom, err := address.NewFromString(tx.TxFrom)
		if err != nil {
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
