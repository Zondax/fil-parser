package impl

import (
	"context"
	"fmt"

	"github.com/filecoin-project/go-address"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/actors/cache/impl/common"
	logger2 "github.com/zondax/fil-parser/logger"
	metrics2 "github.com/zondax/golem/pkg/metrics"

	"github.com/zondax/fil-parser/types"
	"github.com/zondax/golem/pkg/logger"
	"github.com/zondax/golem/pkg/zcache"
)

const (
	ZCacheImpl      = "zcache"
	ZCacheLocalOnly = "in-memory"
	ZCacheCombined  = "combined"
	DummyTtl        = -1
	PrefixSplitter  = "/"
)

// ZCache In-Memory database
type ZCache struct {
	shortCidMap        zcache.ZCache
	robustShortMap     zcache.ZCache
	shortRobustMap     zcache.ZCache
	selectorHashSigMap zcache.ZCache
	logger             *logger.Logger
	cacheType          string
}

func (m *ZCache) NewImpl(source common.DataSource, logger *logger.Logger) error {
	var err error
	m.logger = logger2.GetSafeLogger(logger)

	// If no config was provided, the combined cache is configured as
	// remote best effort, as the remote cache will fail. However, the cache will
	// work anyway
	cacheConfig := source.Config.Cache
	if cacheConfig == nil {
		m.cacheType = ZCacheLocalOnly
		if err := m.initMapsLocalCache(); err != nil {
			return err
		}
	} else {
		m.cacheType = ZCacheCombined

		prefix := ""
		if cacheConfig.GlobalPrefix != "" {
			prefix = fmt.Sprintf("%s%s", cacheConfig.GlobalPrefix, PrefixSplitter)
		}
		if source.Config.NetworkName != "" {
			prefix = fmt.Sprintf("%s%s%s", prefix, source.Config.NetworkName, PrefixSplitter)
		}

		if err := m.initMapsCombinedCache(prefix, cacheConfig); err != nil {
			return err
		}
	}

	if m.shortCidMap, err = zcache.NewLocalCache(&zcache.LocalConfig{
		Prefix:       Short2CidMapPrefix,
		Logger:       m.logger,
		MetricServer: metrics2.NewNoopMetrics()},
	); err != nil {
		return fmt.Errorf("error creating shortCidMap for local zcache, err: %s", err)
	}

	return nil
}

func (m *ZCache) initMapsCombinedCache(prefix string, cacheConfig *zcache.CombinedConfig) error {
	var err error
	robustShortMapConfig := &zcache.CombinedConfig{
		GlobalPrefix:       fmt.Sprintf("%s%s", prefix, Robust2ShortMapPrefix),
		IsRemoteBestEffort: cacheConfig.IsRemoteBestEffort,
		Local:              cacheConfig.Local,
		Remote:             cacheConfig.Remote,
		GlobalLogger:       m.logger,
	}

	shortRobustMapConfig := &zcache.CombinedConfig{
		GlobalPrefix:       fmt.Sprintf("%s%s", prefix, Short2RobustMapPrefix),
		IsRemoteBestEffort: cacheConfig.IsRemoteBestEffort,
		Local:              cacheConfig.Local,
		Remote:             cacheConfig.Remote,
		GlobalLogger:       m.logger,
	}

	selectorHashSigMapConfig := &zcache.CombinedConfig{
		GlobalPrefix:       fmt.Sprintf("%s%s", prefix, SelectorHash2SigMapPrefix),
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
	if m.selectorHashSigMap, err = zcache.NewCombinedCache(selectorHashSigMapConfig); err != nil {
		return fmt.Errorf("error creating selectorHashSigMap for combined zcache, err: %s", err)
	}
	return nil
}

func (m *ZCache) initMapsLocalCache() error {
	var err error

	if m.robustShortMap, err = zcache.NewLocalCache(&zcache.LocalConfig{
		Prefix:       Robust2ShortMapPrefix,
		Logger:       m.logger,
		MetricServer: metrics2.NewNoopMetrics()},
	); err != nil {
		return fmt.Errorf("error creating robustShortMap for local zcache, err: %s", err)
	}
	if m.shortRobustMap, err = zcache.NewLocalCache(&zcache.LocalConfig{
		Prefix:       Short2RobustMapPrefix,
		Logger:       m.logger,
		MetricServer: metrics2.NewNoopMetrics()},
	); err != nil {
		return fmt.Errorf("error creating shortRobustMap for local zcache, err: %s", err)
	}
	if m.selectorHashSigMap, err = zcache.NewLocalCache(&zcache.LocalConfig{
		Prefix:       SelectorHash2SigMapPrefix,
		Logger:       m.logger,
		MetricServer: metrics2.NewNoopMetrics()},
	); err != nil {
		return fmt.Errorf("error creating selectorHashSigMap for local zcache, err: %s", err)
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
		m.logger.Debugf("[ActorsCache] - short address [%s] not found, err: %s\n", address.String(), err.Error())
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

func (m *ZCache) GetEVMSelectorSig(ctx context.Context, selectorHash string) (string, error) {
	var selectorSig string
	if err := m.selectorHashSigMap.Get(ctx, selectorHash, &selectorSig); err != nil {
		if !m.selectorHashSigMap.IsNotFoundError(err) {
			return "", err
		}
	}
	return selectorSig, nil
}

func (m *ZCache) StoreEVMSelectorSig(ctx context.Context, selectorHash, selectorSig string) error {
	if err := m.selectorHashSigMap.Set(ctx, selectorHash, selectorSig, 0); err != nil {
		return fmt.Errorf("error adding selector_sig to cache: %w", err)
	}
	return nil
}

func (m *ZCache) storeRobustShort(robust string, short string) {
	if robust == "" || short == "" {
		m.logger.Debugf("[ActorsCache] - Trying to store empty robust or short address")
		return
	}

	// Possible ZCache types can be Local or Combined. Both types set the TTL at instantiation time
	// The ttl here is pointless
	ctx := context.Background()
	_ = m.robustShortMap.Set(ctx, robust, short, DummyTtl)
}

func (m *ZCache) storeShortRobust(short string, robust string) {
	if robust == "" || short == "" {
		m.logger.Debugf("[ActorsCache] - Trying to store empty robust or short address")
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
		m.logger.Debugf("[ActorsCache] - Trying to store empty cid or short address")
		return
	}

	// Possible ZCache types can be Local or Combined. Both types set the TTL at instantiation time
	// The ttl here is pointless
	ctx := context.Background()
	_ = m.shortCidMap.Set(ctx, shortAddress, cid, DummyTtl)
}
