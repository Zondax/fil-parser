package cache

import (
	"fmt"
	"github.com/filecoin-project/go-address"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	cmap "github.com/orcaman/concurrent-map"
	"github.com/zondax/fil-parser/actors/cache/impl"
	"github.com/zondax/fil-parser/actors/cache/impl/common"
	logger2 "github.com/zondax/fil-parser/logger"
	"github.com/zondax/fil-parser/types"
	"go.uber.org/zap"
	"strings"
)

// SystemActorsId Map to identify system actors which don't have an associated robust address
var SystemActorsId = map[string]bool{
	"f00":  true,
	"f01":  true,
	"f02":  true,
	"f03":  true,
	"f04":  true,
	"f05":  true,
	"f06":  true,
	"f07":  true,
	"f099": true,
}

func SetupActorsCache(dataSource common.DataSource, logger *zap.Logger) (*ActorsCache, error) {
	var offlineCache IActorsCache
	var onChainCache impl.OnChain

	logger = logger2.GetSafeLogger(logger)
	err := onChainCache.NewImpl(dataSource, logger)
	if err != nil {
		return nil, err
	}

	// Try kvStore cache, if it fails, on-memory cache
	var kvStoreCache impl.KVStore
	err = kvStoreCache.NewImpl(dataSource, logger)
	if err == nil {
		offlineCache = &kvStoreCache
	} else {
		logger.Sugar().Warn("[ActorsCache] - Unable to initialize kv store cache. Using on-memory cache")
		var inMemoryCache impl.Memory
		err = inMemoryCache.NewImpl(dataSource, logger)
		if err != nil {
			logger.Sugar().Errorf("[ActorsCache] - Unable to initialize on-memory cache: %s", err.Error())
			return nil, err
		}
		offlineCache = &inMemoryCache
	}

	logger.Sugar().Infof("[ActorsCache] - Actors cache initialized. Offline cache implementation: %s", offlineCache.ImplementationType())

	return &ActorsCache{
		offlineCache: offlineCache,
		onChainCache: &onChainCache,
		badAddress:   cmap.New(),
		logger:       logger,
	}, nil
}

func (a *ActorsCache) ClearBadAddressCache() {
	a.badAddress.Clear()
}

func (a *ActorsCache) GetActorCode(add address.Address, key filTypes.TipSetKey) (string, error) {
	// Check if this address is flagged as bad
	if a.isBadAddress(add) {
		return "", fmt.Errorf("address %s is flagged as bad", add.String())
	}

	// Try kv store cache
	actorCode, err := a.offlineCache.GetActorCode(add, key)
	if err == nil {
		return actorCode, nil
	}

	a.logger.Sugar().Debugf("[ActorsCache] - Unable to retrieve actor code from offline cache for address %s. Trying on-chain cache", add.String())
	// Try on-chain cache
	actorCode, err = a.onChainCache.GetActorCode(add, key)
	if err != nil {
		a.logger.Sugar().Error("[ActorsCache] - Unable to retrieve actor code from node: %s", err.Error())
		if strings.Contains(err.Error(), "actor not found") {
			a.badAddress.Set(add.String(), true)
		}

		return "", err
	}

	// Code is not cached, store it
	err = a.storeActorCode(add, types.AddressInfo{
		ActorCid: actorCode,
	})

	if err != nil {
		a.logger.Sugar().Errorf("[ActorsCache] - Unable to store address info: %s", err.Error())
		return "", err
	}

	return actorCode, nil
}

func (a *ActorsCache) GetRobustAddress(add address.Address) (string, error) {
	if _, ok := SystemActorsId[add.String()]; ok {
		return add.String(), nil
	}

	// Try offline store cache
	robust, err := a.offlineCache.GetRobustAddress(add)
	if err == nil {
		return robust, nil
	}

	// Check if this is a flagged address
	if a.isBadAddress(add) {
		return "", fmt.Errorf("address %s is flagged as bad", add.String())
	}

	a.logger.Sugar().Debugf("[ActorsCache] - Unable to retrieve robust address from offline cache for address %s. Trying on-chain cache", add.String())

	// Try on-chain cache
	robust, err = a.onChainCache.GetRobustAddress(add)
	if err != nil {
		a.logger.Sugar().Errorf("[ActorsCache] - Unable to retrieve actor code from node: %s", err.Error())
		return "", err
	}

	// Robust address is not cached, store it
	err = a.storeRobustAddress(add, types.AddressInfo{
		Robust: robust,
	})

	if err != nil {
		a.logger.Sugar().Errorf("[ActorsCache] - Unable to store address info: %s", err.Error())
		return "", err
	}

	return robust, nil
}

func (a *ActorsCache) GetShortAddress(add address.Address) (string, error) {
	// Try kv store cache
	short, err := a.offlineCache.GetShortAddress(add)
	if err == nil {
		return short, nil
	}

	// Check if this is a flagged address
	if a.isBadAddress(add) {
		return "", fmt.Errorf("address %s is flagged as bad", add.String())
	}

	a.logger.Sugar().Debugf("[ActorsCache] - Unable to retrieve short address from offline cache for address %s. Trying on-chain cache", add.String())

	// Try on-chain cache
	short, err = a.onChainCache.GetShortAddress(add)
	if err != nil {
		a.logger.Sugar().Error("[ActorsCache] - Unable to retrieve actor code from node: %s", err.Error())
		return "", err
	}

	// Robust address is not cached, store it
	err = a.storeShortAddress(add, types.AddressInfo{
		Short: short,
	})

	if err != nil {
		a.logger.Sugar().Errorf("[ActorsCache] - Unable to store address info: %s", err.Error())
		return "", err
	}

	return short, nil
}

func (a *ActorsCache) storeActorCode(add address.Address, info types.AddressInfo) error {
	shortAddress, err := a.GetShortAddress(add)
	if err != nil {
		return err
	}

	a.offlineCache.StoreAddressInfo(types.AddressInfo{
		Short:    shortAddress,
		ActorCid: info.ActorCid,
	})

	return nil
}

func (a *ActorsCache) storeShortAddress(add address.Address, info types.AddressInfo) error {
	robustAddress, err := a.GetRobustAddress(add)
	if err != nil {
		return err
	}

	a.offlineCache.StoreAddressInfo(types.AddressInfo{
		Short:  info.Short,
		Robust: robustAddress,
	})

	return nil
}

func (a *ActorsCache) storeRobustAddress(add address.Address, info types.AddressInfo) error {
	shortAddress, err := a.GetShortAddress(add)
	if err != nil {
		return err
	}

	a.offlineCache.StoreAddressInfo(types.AddressInfo{
		Short:  shortAddress,
		Robust: info.Robust,
	})

	return nil
}

func (a *ActorsCache) isBadAddress(add address.Address) bool {
	_, bad := a.badAddress.Get(add.String())
	return bad
}
