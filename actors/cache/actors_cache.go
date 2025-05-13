package cache

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/cenkalti/backoff/v4"
	actormetrics "github.com/zondax/fil-parser/actors/metrics"
	"github.com/zondax/fil-parser/metrics"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/golem/pkg/logger"

	"github.com/filecoin-project/go-address"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/go-resty/resty/v2"
	cmap "github.com/orcaman/concurrent-map"
	"github.com/zondax/fil-parser/actors/cache/impl"
	"github.com/zondax/fil-parser/actors/cache/impl/common"
	logger2 "github.com/zondax/fil-parser/logger"
	"github.com/zondax/fil-parser/types"
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
	"f010": true,
	"f099": true,
}

// CalibrationActorsId Map to identify system actors which don't have an associated robust address in the calibration network
// These are storage miners that initiated the calibration network
var CalibrationActorsId = map[string]bool{
	"f01000": true,
	"f01001": true,
	"f01002": true,
}

func SetupActorsCache(dataSource common.DataSource, logger *logger.Logger, metrics metrics.MetricsClient, backoff backoff.BackOff) (*ActorsCache, error) {
	var offChainCache IActorsCache
	var onChainCache impl.OnChain

	logger = logger2.GetSafeLogger(logger)

	err := onChainCache.NewImpl(dataSource, logger)
	if err != nil {
		return nil, err
	}

	var combinedCache impl.ZCache
	if err = combinedCache.NewImpl(dataSource, logger); err != nil {
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
		metrics:       actormetrics.NewClient(metrics, "actorsCache"),
		networkName:   dataSource.Config.NetworkName,
	}, nil
}

func (a *ActorsCache) ClearBadAddressCache() {
	a.badAddress.Clear()
}

func (a *ActorsCache) GetActorCode(add address.Address, key filTypes.TipSetKey, onChainOnly bool) (string, error) {
	// Check if this address is flagged as bad
	if a.isBadAddress(add) {
		return "", fmt.Errorf("address %s is flagged as bad", add.String())
	}

	if !onChainOnly {
		actorCode, err := a.offChainCache.GetActorCode(add, key)
		if err == nil {
			return actorCode, nil
		}
	}

	a.logger.Debugf("[ActorsCache] - Unable to retrieve actor code from offchain cache for address %s. Trying on-chain cache", add.String())
	// Try on-chain cache
	actorCode, err := a.onChainCache.GetActorCode(add, key)
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
	if _, ok := SystemActorsId[add.String()]; ok {
		return add.String(), nil
	}

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

func (a *ActorsCache) StoreAddressInfoAddress(addInfo types.AddressInfo) {
	a.offChainCache.StoreAddressInfo(addInfo)
}
