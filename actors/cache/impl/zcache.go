package impl

import (
	"context"
	"fmt"
	"github.com/zondax/fil-parser/actors/constants"
	"strings"

	"github.com/filecoin-project/go-address"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/actors/cache/impl/common"
	logger2 "github.com/zondax/fil-parser/logger"
	"github.com/zondax/fil-parser/types"
	"github.com/zondax/golem/pkg/zcache"
	"go.uber.org/zap"
)

const (
	ZCacheImpl      = "zcache"
	ZCacheLocalOnly = "in-memory"
	ZCacheCombined  = "combined"
	NoTtl           = -1
	NotExpiringTtl  = -1
	PrefixSplitter  = "/"
)

// ZCache In-Memory database
type ZCache struct {
	shortCidMap    zcache.ZCache
	robustShortMap zcache.ZCache
	shortRobustMap zcache.ZCache
	logger         *zap.Logger
	cacheType      string
	ttl            int
}

func (m *ZCache) NewImpl(source common.DataSource, logger *zap.Logger) error {
	var err error
	m.logger = logger2.GetSafeLogger(logger)

	// If no config was provided, the combined cache is configured as
	// remote best effort, as the remote cache will fail. However, the cache will
	// work anyway
	cacheConfig := source.Config.Cache
	if cacheConfig == nil {
		m.cacheType = ZCacheLocalOnly
		m.ttl = NoTtl

		if m.robustShortMap, err = zcache.NewLocalCache(&zcache.LocalConfig{Prefix: Robust2ShortMapPrefix, Logger: m.logger}); err != nil {
			return fmt.Errorf("error creating robustShortMap for local zcache, err: %s", err)
		}
		if m.shortRobustMap, err = zcache.NewLocalCache(&zcache.LocalConfig{Prefix: Short2RobustMapPrefix, Logger: m.logger}); err != nil {
			return fmt.Errorf("error creating shortRobustMap for local zcache, err: %s", err)
		}
	} else {
		m.cacheType = ZCacheCombined
		m.ttl = cacheConfig.GlobalTtlSeconds

		prefix := ""
		if cacheConfig.GlobalPrefix != "" {
			prefix = fmt.Sprintf("%s%s", cacheConfig.GlobalPrefix, PrefixSplitter)
		}
		if source.Config.NetworkName != "" {
			prefix = fmt.Sprintf("%s%s%s", prefix, source.Config.NetworkName, PrefixSplitter)
		}

		robustShortMapConfig := &zcache.CombinedConfig{
			GlobalPrefix:       fmt.Sprintf("%s%s", prefix, Robust2ShortMapPrefix),
			GlobalTtlSeconds:   cacheConfig.GlobalTtlSeconds,
			IsRemoteBestEffort: cacheConfig.IsRemoteBestEffort,
			Local:              cacheConfig.Local,
			Remote:             cacheConfig.Remote,
			GlobalLogger:       m.logger,
		}

		shortRobustMapConfig := &zcache.CombinedConfig{
			GlobalPrefix:       fmt.Sprintf("%s%s", prefix, Short2RobustMapPrefix),
			GlobalTtlSeconds:   cacheConfig.GlobalTtlSeconds,
			IsRemoteBestEffort: cacheConfig.IsRemoteBestEffort,
			Local:              cacheConfig.Local,
			Remote:             cacheConfig.Remote,
			GlobalLogger:       m.logger,
		}

		if m.robustShortMap, err = zcache.NewCombinedCache(robustShortMapConfig); err != nil {
			return fmt.Errorf("error creating robustShortMap for combined zcache, err: %s", err)
		}
		if m.shortRobustMap, err = zcache.NewCombinedCache(shortRobustMapConfig); err != nil {
			return fmt.Errorf("error creating shortRobustMap for combined zcache, err: %s", err)
		}
	}

	if m.shortCidMap, err = zcache.NewLocalCache(&zcache.LocalConfig{Prefix: Short2CidMapPrefix, Logger: m.logger}); err != nil {
		return fmt.Errorf("error creating shortCidMap for local zcache, err: %s", err)
	}

	return nil
}

func (m *ZCache) ImplementationType() string {
	return ZCacheImpl + "/" + m.cacheType
}

func (m *ZCache) BackFill() error {
	// Nothing to do
	return nil
}

func (m *ZCache) GetActorCode(address address.Address, key filTypes.TipSetKey) (string, error) {
	shortAddress, err := m.GetShortAddress(address)
	if err != nil {
		m.logger.Sugar().Debugf("[ActorsCache] - short address [%s] not found, err: %s\n", address.String(), err.Error())
		return cid.Undef.String(), common.ErrKeyNotFound
	}

	var code string
	ctx := context.Background()
	if err = m.shortCidMap.Get(ctx, shortAddress, &code); err != nil {
		return cid.Undef.String(), common.ErrKeyNotFound
	}

	if code == "" {
		return cid.Undef.String(), common.ErrEmptyValue
	}

	return code, nil
}

func (m *ZCache) GetRobustAddress(address address.Address) (string, error) {
	isRobustAddress, err := common.IsRobustAddress(address)
	if err != nil {
		return "", err
	}

	ctx := context.Background()

	if isRobustAddress {
		// If already a robust address, we attempt to get a f4 address.
		// This is particularly useful in the case of EVM actors, where a robust f2 address
		// may need to be converted as a f4 address.
		if f4Address := m.tryToGetF4Address(address); f4Address != "" {
			return f4Address, nil
		}

		return address.String(), nil
	}

	var robustAdd string
	if err = m.shortRobustMap.Get(ctx, address.String(), &robustAdd); err != nil {
		return "", common.ErrKeyNotFound
	}

	if robustAdd == "" {
		return "", common.ErrEmptyValue
	}

	return robustAdd, nil
}

func (m *ZCache) GetShortAddress(address address.Address) (string, error) {
	isRobustAddress, err := common.IsRobustAddress(address)
	if err != nil {
		return "", err
	}

	if !isRobustAddress {
		// Already a short address
		return address.String(), nil
	}

	// This is a robust address, get the short one
	var shortAdd string
	ctx := context.Background()

	if err = m.robustShortMap.Get(ctx, address.String(), &shortAdd); err != nil {
		return "", common.ErrKeyNotFound
	}

	if shortAdd == "" {
		return "", common.ErrEmptyValue
	}

	return shortAdd, nil
}

func (m *ZCache) storeRobustShort(robust string, short string) {
	if robust == "" || short == "" {
		m.logger.Sugar().Debugf("[ActorsCache] - Trying to store empty robust or short address")
		return
	}

	// Possible ZCache types can be Local or Combined. Both types set the TTL at instantiation time
	// The ttl here is pointless
	ctx := context.Background()
	_ = m.robustShortMap.Set(ctx, robust, short, NotExpiringTtl)
}

func (m *ZCache) storeShortRobust(short string, robust string) {
	if robust == "" || short == "" {
		m.logger.Sugar().Debugf("[ActorsCache] - Trying to store empty robust or short address")
		return
	}

	// Possible ZCache types can be Local or Combined. Both types set the TTL at instantiation time
	// The ttl here is pointless
	ctx := context.Background()
	_ = m.shortRobustMap.Set(ctx, short, robust, NotExpiringTtl)
}

func (m *ZCache) StoreAddressInfo(info types.AddressInfo) {
	m.storeRobustShort(info.Robust, info.Short)
	m.storeActorCode(info.Short, info.ActorCid)

	isEvm := strings.EqualFold(info.ActorType, constants.ActorTypeEVM)
	isEvmAndAddressIsF4 := isEvm && strings.HasPrefix(info.Robust, constants.AddressTypePrefixF4)

	// Only store the mapping for addresses that are not related to EVM actors,
	// or are associated with EVM actors but use an f4 prefix. We skip storing
	// addresses when they are f2 for EVM actors because only f4 addresses are of interest.
	if !isEvm || isEvmAndAddressIsF4 {
		m.storeShortRobust(info.Short, info.Robust)
	}
}

func (m *ZCache) storeActorCode(shortAddress string, cid string) {
	if shortAddress == "" || cid == "" {
		m.logger.Sugar().Debugf("[ActorsCache] - Trying to store empty cid or short address")
		return
	}

	// Possible ZCache types can be Local or Combined. Both types set the TTL at instantiation time
	// The ttl here is pointless
	ctx := context.Background()
	_ = m.shortCidMap.Set(ctx, shortAddress, cid, NotExpiringTtl)
}

func (m *ZCache) tryToGetF4Address(address address.Address) string {
	ctx := context.Background()

	// Return the address if it's already a f4 address
	if strings.HasPrefix(address.String(), constants.AddressTypePrefixF4) {
		return address.String()
	}

	// If the robust address is not f4, it should be f2
	// Try to get the corresponding f0 for the f2 address
	f0Address, err := m.GetShortAddress(address)
	if err != nil {
		m.logger.Sugar().Errorf("error getting short address for %s: %s", address.String(), err)
		return ""
	}

	// Try to get the f4 address associated with the f0 address
	// If no f4 is found, it implies the address might not be an EVM actor type
	var f4Address string
	err = m.shortRobustMap.Get(ctx, f0Address, &f4Address)
	if err == nil && f4Address != "" {
		return f4Address
	}

	m.logger.Sugar().Infof("no f4 address associated with f0 address: %s. The address might not be an EVM actor type.", f0Address)
	return ""
}
