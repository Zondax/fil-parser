package cache

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/zondax/fil-parser/metrics"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/golem/pkg/logger"
	golemBackoff "github.com/zondax/golem/pkg/zhttpclient/backoff"

	"github.com/filecoin-project/go-address"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/go-resty/resty/v2"
	cmap "github.com/orcaman/concurrent-map"
	"github.com/zondax/fil-parser/actors/cache/impl"
	"github.com/zondax/fil-parser/actors/cache/impl/common"
	cacheMetrics "github.com/zondax/fil-parser/actors/cache/metrics"

	logger2 "github.com/zondax/fil-parser/logger"
	"github.com/zondax/fil-parser/types"
)

// SystemActorsId Map to identify system actors which don't have an associated robust address.
// https://github.com/filecoin-project/go-state-types/blob/571b84617a4b7fe032cf63c25e0c079f90e2f8a7/builtin/singletons.go#L9
var SystemActorsId = map[string]bool{
	// system actor
	"f00": true,
	// init actor
	"f01": true,
	// reward actor
	"f02": true,
	// cron actor
	"f03": true,
	// storagepower actor
	"f04": true,
	// storagemarket actor
	"f05": true,
	// verifiedregistry actor
	"f06": true,
	// datacap actor
	"f07": true,
	// eam actor
	"f010": true,
}

// GenesisActorsId Map to identify actors created at genesis and don't have an associated robust address.
// Check data/genesis/{network}_genesis_balances.json
var GenesisActorsId = map[string]bool{
	// multisig
	"f080": true,
	// keyless account actor
	"f090": true,

	"f0115": true,
	"f0116": true,
	"f0117": true,
	"f0121": true,
	"f0118": true,
	"f0119": true,
	"f0120": true,
	"f0122": true,
	// burn address
	"f099": true,
}

// CalibrationActorsId Map to identify actors created at genesis which don't have an associated robust address in the calibration network.
// These are storage miners and multisig addresses that initiated the calibration network.
// Check data/genesis/calibration_genesis_balances.json
var CalibrationActorsId = map[string]bool{
	// miners
	"f01000": true,
	"f01001": true,
	"f01002": true,
	// multisig
	"f080": true,
}

var ErrBadAddress = errors.New("ErrBadAddress")

const combinedCacheImpl = "combined"

func SetupActorsCache(dataSource common.DataSource, logger *logger.Logger, metricsClient metrics.MetricsClient, backoff *golemBackoff.BackOff) (*ActorsCache, error) {
	var setupMu sync.Mutex
	setupMu.Lock()
	defer setupMu.Unlock()

	var offChainCache IActorsCache

	var onChainCache impl.OnChain

	logger = logger2.GetSafeLogger(logger)
	metrics := cacheMetrics.NewClient(metricsClient, "actorsCache")

	err := onChainCache.NewImpl(dataSource, logger, metrics, backoff)
	if err != nil {
		return nil, err
	}

	var combinedCache impl.ZCacheBlockConfirmation
	if err = combinedCache.NewImpl(dataSource, logger, metrics, backoff); err != nil {
		logger.Errorf("[ActorsCache] - Unable to initialize combined cache: %s", err.Error())
		return nil, err
	}

	offChainCache = &combinedCache

	logger.Infof("[ActorsCache] - Actors cache initialized. Off chain cache implementation: %s", offChainCache.ImplementationType())

	return &ActorsCache{
		offChainCache: offChainCache,
		onChainCache:  &onChainCache,
		badAddress:    cmap.New(),
		logger:        logger,
		httpClient:    resty.New().SetTimeout(30 * time.Second),
		metrics:       metrics,
		networkName:   dataSource.Config.NetworkName,
	}, nil
}

func (a *ActorsCache) ClearBadAddressCache() {
	a.badAddress.Clear()
}

func (a *ActorsCache) GetActorCode(add address.Address, key filTypes.TipSetKey, onChainOnly, canonical bool) (string, error) {
	addStr := add.String()

	actorCode, err := a.getActorCode(add, key, onChainOnly, canonical)
	if err != nil {
		a.logger.Errorf("[ActorsCache] - Unable to retrieve actor code from node: %s", err.Error())
		if strings.Contains(err.Error(), "actor not found") {
			a.badAddress.Set(addStr, true)
		}

		return "", err
	}

	// Code is not cached, store it
	err = a.storeActorCode(add, types.AddressInfo{
		ActorCid:    actorCode,
		IsCanonical: canonical,
	})

	if err != nil {
		a.logger.Errorf("[ActorsCache] - Unable to store address info: %s", err.Error())
		return "", err
	}

	return actorCode, nil
}

func (a *ActorsCache) GetRobustAddress(add address.Address, canonical bool) (string, error) {
	addStr := add.String()
	// check if the address is a system actor ( no robust address)
	if _, ok := SystemActorsId[addStr]; ok {
		return addStr, nil
	}
	// check if the address is a genesis actor ( no robust address)
	if _, ok := GenesisActorsId[addStr]; ok {
		return addStr, nil
	}
	// check if the address is a calibration genesis actor ( no robust address)
	if tools.ParseRawNetworkName(a.networkName) == tools.CalibrationNetwork {
		if _, ok := CalibrationActorsId[addStr]; ok {
			return addStr, nil
		}
	}

	robust, err := a.getRobustAddress(add, canonical)
	if err != nil {
		return "", err
	}

	// Robust address is not cached, store it
	err = a.storeRobustAddress(add, types.AddressInfo{
		Robust:      robust,
		IsCanonical: canonical,
	})

	if err != nil {
		a.logger.Errorf("[ActorsCache] - Unable to store address info: %s", err.Error())
		return "", err
	}

	return robust, nil
}

func (a *ActorsCache) GetShortAddress(add address.Address, canonical bool) (string, error) {
	short, err := a.getShortAddress(add, canonical)
	if err != nil {
		return "", err
	}

	// Robust address is not cached, store it
	err = a.storeShortAddress(add, types.AddressInfo{
		Short:       short,
		IsCanonical: canonical,
	})

	if err != nil {
		a.logger.Errorf("[ActorsCache] - Unable to store address info: %s", err.Error())
		return "", err
	}

	return short, nil
}

func (a *ActorsCache) GetEVMSelectorSig(ctx context.Context, selectorID string, canonical bool) (string, error) {
	selectorSig, err := a.getEVMSelectorSig(ctx, selectorID, canonical)
	if err != nil {
		return "", err
	}

	if err := a.offChainCache.StoreEVMSelectorSig(ctx, selectorID, selectorSig, canonical); err != nil {
		return selectorSig, fmt.Errorf("error adding selector_sig to cache: %w", err)
	}
	return selectorSig, nil
}

// IsSystemActor checks if addr is a system actor as defined here:
// https://github.com/filecoin-project/go-state-types/blob/571b84617a4b7fe032cf63c25e0c079f90e2f8a7/builtin/singletons.go#L9
func (a *ActorsCache) IsSystemActor(addr string) bool {
	return SystemActorsId[addr]
}

// IsGenesisActor checks if addr is a genesis actor as defined here:
func (a *ActorsCache) IsGenesisActor(addr string) bool {
	return GenesisActorsId[addr]
}

func (a *ActorsCache) getEVMSelectorSig(ctx context.Context, selectorID string, canonical bool) (string, error) {
	selectorSig, err := a.offChainCache.GetEVMSelectorSig(ctx, selectorID, canonical)
	if err == nil {
		return selectorSig, nil
	}
	a.logger.Debugf("[ActorsCache] - Unable to retrieve selector_sig from offchain cache for selector_id %s. Trying onchain cache", selectorID)
	// Try onchain
	selectorSig, err = a.onChainCache.GetEVMSelectorSig(ctx, selectorID, canonical)
	if err == nil {
		return selectorSig, nil
	}
	a.logger.Debugf("[ActorsCache] - Unable to retrieve selector_sig from onchain cache for selector_id %s", selectorID)
	return "", err
}

func (a *ActorsCache) getShortAddress(add address.Address, canonical bool) (string, error) {
	addStr := add.String()
	// Try canonical cache
	short, err := a.offChainCache.GetShortAddress(add, canonical)
	if err == nil {
		return short, nil
	}
	a.logger.Debugf("[ActorsCache] - Unable to retrieve short address from offchain cache for address %s. Trying onchain cache", addStr)
	// Try onchain
	// Check if this is a flagged address
	if a.isBadAddress(add) {
		return "", fmt.Errorf("address %s is flagged as bad", addStr)
	}
	short, err = a.onChainCache.GetShortAddress(add, canonical)
	if err == nil {
		return short, nil
	}
	a.logger.Debugf("[ActorsCache] - Unable to retrieve short address from onchain cache for address %s.", addStr)

	return "", nil
}

func (a *ActorsCache) getRobustAddress(add address.Address, canonical bool) (string, error) {
	addStr := add.String()
	robust, err := a.offChainCache.GetRobustAddress(add, canonical)
	if err == nil {
		return robust, nil
	}
	a.logger.Debugf("[ActorsCache] - Unable to retrieve robust address from offchain cache for address %s. Trying latest cache", addStr)
	// Try onchain
	if a.isBadAddress(add) {
		return "", fmt.Errorf("%w: address %s is flagged as bad", ErrBadAddress, addStr)
	}
	robust, err = a.onChainCache.GetRobustAddress(add, canonical)
	if err == nil {
		return robust, nil
	}
	a.logger.Debugf("[ActorsCache] - Unable to retrieve robust address from onchain cache for address %s.", addStr)

	return "", err
}

func (a *ActorsCache) getActorCode(add address.Address, key filTypes.TipSetKey, onChainOnly, canonical bool) (string, error) {
	addStr := add.String()
	actorCode, err := a.offChainCache.GetActorCode(add, key, onChainOnly, true)
	if err == nil {
		return actorCode, nil
	}
	a.logger.Debugf("[ActorsCache] - Unable to retrieve actor code from offchain cache for address %s. Trying on-chain cache", addStr)
	// Try onchain
	if a.isBadAddress(add) {
		return "", fmt.Errorf(" %w : %s is flagged as bad", ErrBadAddress, addStr)
	}

	actorCode, err = a.onChainCache.GetActorCode(add, key, onChainOnly, true)
	if err == nil {
		return actorCode, nil
	}
	a.logger.Debugf("[ActorsCache] - Unable to retrieve actor code from onchain cache for address %s.", addStr)
	return "", err
}

func (a *ActorsCache) StoreEVMSelectorSig(ctx context.Context, selectorID string, sig string, canonical bool) error {
	return a.offChainCache.StoreEVMSelectorSig(ctx, selectorID, sig, canonical)
}

func (a *ActorsCache) storeActorCode(add address.Address, info types.AddressInfo) error {
	shortAddress, err := a.GetShortAddress(add, info.IsCanonical)
	if err != nil {
		return err
	}

	a.offChainCache.StoreAddressInfo(types.AddressInfo{
		Short:       shortAddress,
		ActorCid:    info.ActorCid,
		IsCanonical: info.IsCanonical,
	})

	return nil
}

func (a *ActorsCache) storeShortAddress(add address.Address, info types.AddressInfo) error {
	robustAddress, err := a.getRobustAddress(add, info.IsCanonical)
	if err != nil {
		return err
	}

	a.offChainCache.StoreAddressInfo(types.AddressInfo{
		Short:       info.Short,
		Robust:      robustAddress,
		IsCanonical: info.IsCanonical,
	})

	return nil
}

func (a *ActorsCache) storeRobustAddress(add address.Address, info types.AddressInfo) error {
	shortAddress, err := a.getShortAddress(add, info.IsCanonical)
	if err != nil {
		return err
	}

	a.offChainCache.StoreAddressInfo(types.AddressInfo{
		Short:       shortAddress,
		Robust:      info.Robust,
		IsCanonical: info.IsCanonical,
	})
	return nil

}

func (a *ActorsCache) isBadAddress(add address.Address) bool {
	_, bad := a.badAddress.Get(add.String())
	return bad
}

func (a *ActorsCache) BackFill() error {
	return a.offChainCache.BackFill()
}

func (a *ActorsCache) ImplementationType() string {
	return combinedCacheImpl
}

func (a *ActorsCache) NewImpl(dataSource common.DataSource, logger *logger.Logger, metrics *cacheMetrics.ActorsCacheMetricsClient, backoff *golemBackoff.BackOff) error {
	return nil
}

func (a *ActorsCache) StoreAddressInfo(addInfo types.AddressInfo) {
	a.offChainCache.StoreAddressInfo(addInfo)
}
