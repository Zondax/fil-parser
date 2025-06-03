package cache

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/zondax/fil-parser/metrics"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/golem/pkg/logger"

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
	"f080":  true,
	"f090":  true,
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

const combinedCacheImpl = "combined"

func SetupActorsCache(dataSource common.DataSource, logger *logger.Logger, metricsClient metrics.MetricsClient, backoff backoff.BackOff) (*ActorsCache, error) {
	var setupMu sync.Mutex
	setupMu.Lock()
	defer setupMu.Unlock()

	var offChainCache IActorsCache
	var onChainCache impl.OnChain

	logger = logger2.GetSafeLogger(logger)
	metrics := cacheMetrics.NewClient(metricsClient, "actorsCache")

	err := onChainCache.NewImpl(dataSource, logger, metrics)
	if err != nil {
		return nil, err
	}

	var combinedCache impl.ZCache
	if err = combinedCache.NewImpl(dataSource, logger, metrics); err != nil {
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

func (a *ActorsCache) GetActorCode(add address.Address, key filTypes.TipSetKey, onChainOnly bool) (string, error) {
	addrStr := add.String()
	// Check if this address is flagged as bad
	if a.isBadAddress(add) {
		return "", fmt.Errorf("address %s is flagged as bad", addrStr)
	}

	if !onChainOnly {
		actorCode, err := a.offChainCache.GetActorCode(add, key, onChainOnly)
		if err == nil {
			return actorCode, nil
		}
	}

	a.logger.Debugf("[ActorsCache] - Unable to retrieve actor code from offchain cache for address %s. Trying on-chain cache", add.String())
	// Try on-chain cache
	actorCode, err := a.onChainCache.GetActorCode(add, key, onChainOnly)
	if err != nil {
		a.logger.Debugf("[ActorsCache] - Unable to retrieve actor code from node: %s", err.Error())
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
		a.logger.Errorf("[ActorsCache] - Unable to store address info: %s", err.Error())
		return "", err
	}

	return actorCode, nil
}

func (a *ActorsCache) GetRobustAddress(add address.Address) (string, error) {
	// check if the address is a system actor ( no robust address)
	if _, ok := SystemActorsId[add.String()]; ok {
		return add.String(), nil
	}
	// check if the address is a genesis actor ( no robust address)
	if _, ok := GenesisActorsId[add.String()]; ok {
		return add.String(), nil
	}
	// check if the address is a calibration genesis actor ( no robust address)
	if tools.ParseRawNetworkName(a.networkName) == tools.CalibrationNetwork {
		if _, ok := CalibrationActorsId[add.String()]; ok {
			return add.String(), nil
		}
	}

	// Try offline store cache
	robust, err := a.offChainCache.GetRobustAddress(add)
	if err == nil {
		return robust, nil
	}

	// Check if this is a flagged address
	if a.isBadAddress(add) {
		return "", fmt.Errorf("address %s is flagged as bad", add.String())
	}

	a.logger.Debugf("[ActorsCache] - Unable to retrieve robust address from offchain cache for address %s. Trying on-chain cache", add.String())

	// Try on-chain cache
	robust, err = a.onChainCache.GetRobustAddress(add)
	if err != nil {
		a.logger.Debugf("[ActorsCache] - Unable to retrieve actor code from node: %s", err.Error())
		return "", err
	}

	// Robust address is not cached, store it
	err = a.storeRobustAddress(add, types.AddressInfo{
		Robust: robust,
	})

	if err != nil {
		a.logger.Errorf("[ActorsCache] - Unable to store address info: %s", err.Error())
		return "", err
	}

	return robust, nil
}

func (a *ActorsCache) GetShortAddress(add address.Address) (string, error) {
	// Try kv store cache
	short, err := a.offChainCache.GetShortAddress(add)
	if err == nil {
		return short, nil
	}

	// Check if this is a flagged address
	if a.isBadAddress(add) {
		return "", fmt.Errorf("address %s is flagged as bad", add.String())
	}

	a.logger.Debugf("[ActorsCache] - Unable to retrieve short address from offchain cache for address %s. Trying on-chain cache", add.String())

	// Try on-chain cache
	short, err = a.onChainCache.GetShortAddress(add)
	if err != nil {
		a.logger.Debugf("[ActorsCache] - Unable to retrieve actor code from node: %s", err.Error())
		return "", err
	}

	// Robust address is not cached, store it
	err = a.storeShortAddress(add, types.AddressInfo{
		Short: short,
	})

	if err != nil {
		a.logger.Errorf("[ActorsCache] - Unable to store address info: %s", err.Error())
		return "", err
	}

	return short, nil
}

func (a *ActorsCache) GetEVMSelectorSig(ctx context.Context, selectorID string) (string, error) {
	selectorSig, err := a.offChainCache.GetEVMSelectorSig(ctx, selectorID)
	if err != nil {
		return "", err
	}

	if selectorSig != "" {
		return selectorSig, nil
	}

	// not found in cache
	resp, err := a.httpClient.NewRequest().
		SetQueryParam("hex_signature", selectorID).
		SetContext(ctx).
		SetResult(&FourBytesSignatureResult{}).
		Get(SignatureDBURL)
	if err != nil {
		return selectorSig, err
	}

	if resp.StatusCode() != http.StatusOK {
		return selectorSig, fmt.Errorf("error from 4bytes: %v", resp.Error())
	}

	signatureData, ok := resp.Result().(*FourBytesSignatureResult)
	if !ok {
		return selectorSig, errors.New("error asserting result to SignatureResult")
	}

	if len(signatureData.Results) == 0 {
		return selectorSig, fmt.Errorf("signature not found: %s", selectorID)
	}

	sig := signatureData.Results[0].TextSignature

	if err := a.offChainCache.StoreEVMSelectorSig(ctx, selectorID, sig); err != nil {
		return selectorSig, fmt.Errorf("error adding selector_sig to cache: %w", err)
	}
	return sig, nil
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

func (a *ActorsCache) StoreEVMSelectorSig(ctx context.Context, selectorID string, sig string) error {
	return a.offChainCache.StoreEVMSelectorSig(ctx, selectorID, sig)
}

func (a *ActorsCache) storeActorCode(add address.Address, info types.AddressInfo) error {
	shortAddress, err := a.GetShortAddress(add)
	if err != nil {
		return err
	}

	a.offChainCache.StoreAddressInfo(types.AddressInfo{
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

	a.offChainCache.StoreAddressInfo(types.AddressInfo{
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

	a.offChainCache.StoreAddressInfo(types.AddressInfo{
		Short:  shortAddress,
		Robust: info.Robust,
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

func (a *ActorsCache) NewImpl(dataSource common.DataSource, logger *logger.Logger, metrics *cacheMetrics.ActorsCacheMetricsClient) error {
	return nil
}

func (a *ActorsCache) StoreAddressInfo(addInfo types.AddressInfo) {
	a.offChainCache.StoreAddressInfo(addInfo)
}
