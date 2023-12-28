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
	CombinedCacheImpl     = "combined-cache"
	Short2CidMapPrefix    = "short2Cid"
	Robust2ShortMapPrefix = "robust2Short"
	Short2RobustMapPrefix = "short2Robust"
)

// Combined In-Memory + Distributed Combined database
type Combined struct {
	shortCidMap    zcache.CombinedCache
	robustShortMap zcache.CombinedCache
	shortRobustMap zcache.CombinedCache
	logger         *zap.Logger
}

func (m *Combined) NewImpl(source common.DataSource, logger *zap.Logger) error {
	m.logger = logger2.GetSafeLogger(logger)

	// If no config was provided, the combined cache is configured as
	// remote best effort, as the remote cache will fail. However, the cache will
	// work anyway
	cacheConfig := source.Config.Cache
	if cacheConfig == nil {
		cacheConfig = &zcache.CombinedConfig{
			IsRemoteBestEffort: true,
			Local:              &zcache.LocalConfig{},
			Remote:             &zcache.RemoteConfig{},
		}
	}

	prefix := ""
	if source.Config.Cache.GlobalPrefix != "" {
		prefix = fmt.Sprintf("%s/", source.Config.Cache.GlobalPrefix)
	}
	if source.Config.NetworkName != "" {
		prefix = fmt.Sprintf("%s%s/", prefix, source.Config.NetworkName)
	}

	shortCidMapConfig := &zcache.CombinedConfig{
		GlobalPrefix:       fmt.Sprintf("%s%s", prefix, Short2CidMapPrefix),
		GlobalTtlSeconds:   source.Config.Cache.GlobalTtlSeconds,
		IsRemoteBestEffort: source.Config.Cache.IsRemoteBestEffort,
		Local:              source.Config.Cache.Local,
		Remote:             source.Config.Cache.Remote,
	}

	robustShortMapConfig := &zcache.CombinedConfig{
		GlobalPrefix:       fmt.Sprintf("%s%s", prefix, Robust2ShortMapPrefix),
		GlobalTtlSeconds:   -1,
		IsRemoteBestEffort: source.Config.Cache.IsRemoteBestEffort,
		Local:              source.Config.Cache.Local,
		Remote:             source.Config.Cache.Remote,
	}

	shortRobustMapConfig := &zcache.CombinedConfig{
		GlobalPrefix:       fmt.Sprintf("%s%s", prefix, Short2RobustMapPrefix),
		GlobalTtlSeconds:   -1,
		IsRemoteBestEffort: source.Config.Cache.IsRemoteBestEffort,
		Local:              source.Config.Cache.Local,
		Remote:             source.Config.Cache.Remote,
	}

	var err error
	if m.shortCidMap, err = zcache.NewCombinedCache(shortCidMapConfig); err != nil {
		return err
	}
	if m.robustShortMap, err = zcache.NewCombinedCache(robustShortMapConfig); err != nil {
		return err
	}
	if m.shortRobustMap, err = zcache.NewCombinedCache(shortRobustMapConfig); err != nil {
		return err
	}

	return nil
}

func (m *Combined) ImplementationType() string {
	return CombinedCacheImpl
}

func (m *Combined) BackFill() error {
	// Nothing to do
	return nil
}

func (m *Combined) GetActorCode(address address.Address, key filTypes.TipSetKey) (string, error) {
	shortAddress, err := m.GetShortAddress(address)
	if err != nil {
		m.logger.Sugar().Infof("[ActorsCache] - Error getting short address: %s\n", err.Error())
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

func (m *Combined) GetRobustAddress(address address.Address) (string, error) {
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

func (m *Combined) GetShortAddress(address address.Address) (string, error) {
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

func (m *Combined) storeRobustShort(robust string, short string) {
	if robust == "" || short == "" {
		// zap.S().Warn("[ActorsCache] - Trying to store empty robust or short address")
		return
	}

	ctx := context.Background()
	_ = m.robustShortMap.Set(ctx, robust, short, -1)
}

func (m *Combined) storeShortRobust(short string, robust string) {
	if robust == "" || short == "" {
		// zap.S().Warn("[ActorsCache] - Trying to store empty robust or short address")
		return
	}

	ctx := context.Background()
	_ = m.shortRobustMap.Set(ctx, short, robust, -1)
}

func (m *Combined) StoreAddressInfo(info types.AddressInfo) {
	m.storeRobustShort(info.Robust, info.Short)
	m.storeShortRobust(info.Short, info.Robust)
	m.storeActorCode(info.Short, info.ActorCid)
}

func (m *Combined) storeActorCode(shortAddress string, cid string) {
	if shortAddress == "" || cid == "" {
		// zap.S().Warn("[ActorsCache] - Trying to store empty cid or short address")
		return
	}

	ctx := context.Background()
	_ = m.shortCidMap.Set(ctx, shortAddress, cid, -1)
}
