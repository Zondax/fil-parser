package impl

import (
	"context"
	"fmt"

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
	DummyTtl        = -1
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

		if m.shortCidMap, err = zcache.NewLocalCache(&zcache.LocalConfig{Prefix: Short2CidMapPrefix, EvictionInSeconds: m.ttl}); err != nil {
			return fmt.Errorf("error creating shortCidMap for local zcache, err: %s", err)
		}
		if m.robustShortMap, err = zcache.NewLocalCache(&zcache.LocalConfig{Prefix: Robust2ShortMapPrefix, EvictionInSeconds: m.ttl}); err != nil {
			return fmt.Errorf("error creating robustShortMap for local zcache, err: %s", err)
		}
		if m.shortRobustMap, err = zcache.NewLocalCache(&zcache.LocalConfig{Prefix: Short2RobustMapPrefix, EvictionInSeconds: m.ttl}); err != nil {
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

		shortCidMapConfig := &zcache.CombinedConfig{
			GlobalPrefix:       fmt.Sprintf("%s%s", prefix, Short2CidMapPrefix),
			GlobalTtlSeconds:   cacheConfig.GlobalTtlSeconds,
			IsRemoteBestEffort: cacheConfig.IsRemoteBestEffort,
			Local:              cacheConfig.Local,
			Remote:             cacheConfig.Remote,
		}

		robustShortMapConfig := &zcache.CombinedConfig{
			GlobalPrefix:       fmt.Sprintf("%s%s", prefix, Robust2ShortMapPrefix),
			GlobalTtlSeconds:   cacheConfig.GlobalTtlSeconds,
			IsRemoteBestEffort: cacheConfig.IsRemoteBestEffort,
			Local:              cacheConfig.Local,
			Remote:             cacheConfig.Remote,
		}

		shortRobustMapConfig := &zcache.CombinedConfig{
			GlobalPrefix:       fmt.Sprintf("%s%s", prefix, Short2RobustMapPrefix),
			GlobalTtlSeconds:   cacheConfig.GlobalTtlSeconds,
			IsRemoteBestEffort: cacheConfig.IsRemoteBestEffort,
			Local:              cacheConfig.Local,
			Remote:             cacheConfig.Remote,
		}

		if m.shortCidMap, err = zcache.NewCombinedCache(shortCidMapConfig); err != nil {
			return fmt.Errorf("error creating shortCidMap for combined zcache, err: %s", err)
		}
		if m.robustShortMap, err = zcache.NewCombinedCache(robustShortMapConfig); err != nil {
			return fmt.Errorf("error creating robustShortMap for combined zcache, err: %s", err)
		}
		if m.shortRobustMap, err = zcache.NewCombinedCache(shortRobustMapConfig); err != nil {
			return fmt.Errorf("error creating shortRobustMap for combined zcache, err: %s", err)
		}
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

	if isRobustAddress {
		// Already a robust address
		return address.String(), nil
	}

	// This is a short address, get the robust one
	var robustAdd string
	ctx := context.Background()
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
	_ = m.robustShortMap.Set(ctx, robust, short, DummyTtl)
}

func (m *ZCache) storeShortRobust(short string, robust string) {
	if robust == "" || short == "" {
		m.logger.Sugar().Debugf("[ActorsCache] - Trying to store empty robust or short address")
		return
	}

	// Possible ZCache types can be Local or Combined. Both types set the TTL at instantiation time
	// The ttl here is pointless
	ctx := context.Background()
	_ = m.shortRobustMap.Set(ctx, short, robust, DummyTtl)
}

func (m *ZCache) StoreAddressInfo(info types.AddressInfo) {
	m.storeRobustShort(info.Robust, info.Short)
	m.storeShortRobust(info.Short, info.Robust)
	m.storeActorCode(info.Short, info.ActorCid)
}

func (m *ZCache) storeActorCode(shortAddress string, cid string) {
	if shortAddress == "" || cid == "" {
		m.logger.Sugar().Debugf("[ActorsCache] - Trying to store empty cid or short address")
		return
	}

	// Possible ZCache types can be Local or Combined. Both types set the TTL at instantiation time
	// The ttl here is pointless
	ctx := context.Background()
	_ = m.shortCidMap.Set(ctx, shortAddress, cid, DummyTtl)
}
