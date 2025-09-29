package impl

import (
	"context"
	"fmt"

	"github.com/filecoin-project/go-address"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/zondax/fil-parser/actors/cache/impl/common"
	cacheMetrics "github.com/zondax/fil-parser/actors/cache/metrics"
	"github.com/zondax/fil-parser/types"
	"github.com/zondax/golem/pkg/logger"
	golemBackoff "github.com/zondax/golem/pkg/zhttpclient/backoff"
)

type ZCacheBlockConfirmation struct {
	offChainLatest    *ZCache
	offChainCanonical *ZCache
}

func (m *ZCacheBlockConfirmation) NewImpl(source common.DataSource, logger *logger.Logger, metrics *cacheMetrics.ActorsCacheMetricsClient, _ *golemBackoff.BackOff) error {
	m.offChainCanonical = &ZCache{}
	if err := m.offChainCanonical.NewImpl(source, logger, metrics); err != nil {
		return err
	}
	source.Config.Cache.Ttl = source.Config.Cache.LatestCacheTTL
	source.Config.Cache.GlobalPrefix = fmt.Sprintf("%s-%s", "latest", source.Config.Cache.GlobalPrefix)
	m.offChainLatest = &ZCache{}
	if err := m.offChainLatest.NewImpl(source, logger, metrics); err != nil {
		return err
	}

	return nil
}

func (m *ZCacheBlockConfirmation) BackFill() error {
	// Nothing to do
	return nil
}

// IsSystemActor returns false for all ZCache implementations as the system actors list is maintained by the helper.
// Use the ActorsCache directly.
// Only required to satisfy IActorsCache.
func (m *ZCacheBlockConfirmation) IsSystemActor(_ string) bool {
	return false
}

// IsGenesisActor returns false for all ZCache implementations as the genesis actors list is maintained by the helper.
// Use the ActorsCache directly.
// Only required to satisfy IActorsCache.
func (m *ZCacheBlockConfirmation) IsGenesisActor(_ string) bool {
	return false
}

func (m *ZCacheBlockConfirmation) StoreAddressInfo(info types.AddressInfo) {
	if info.IsCanonical {
		m.offChainCanonical.StoreAddressInfo(info)
	} else {
		m.offChainLatest.StoreAddressInfo(info)
	}
}

func (m *ZCacheBlockConfirmation) GetActorCode(address address.Address, key filTypes.TipSetKey, _, canonical bool) (string, error) {
	if canonical {
		return m.offChainCanonical.GetActorCode(address, key)
	}
	return m.offChainLatest.GetActorCode(address, key)
}

func (m *ZCacheBlockConfirmation) GetRobustAddress(address address.Address, canonical bool) (string, error) {
	if canonical {
		return m.offChainCanonical.GetRobustAddress(address)
	}
	return m.offChainLatest.GetRobustAddress(address)
}

func (m *ZCacheBlockConfirmation) GetShortAddress(address address.Address, canonical bool) (string, error) {
	if canonical {
		return m.offChainCanonical.GetShortAddress(address)
	}

	return m.offChainLatest.GetShortAddress(address)
}

func (m *ZCacheBlockConfirmation) GetEVMSelectorSig(ctx context.Context, selectorHash string, canonical bool) (string, error) {
	if canonical {
		return m.offChainCanonical.GetEVMSelectorSig(ctx, selectorHash)
	}

	return m.offChainLatest.GetEVMSelectorSig(ctx, selectorHash)
}

func (m *ZCacheBlockConfirmation) StoreEVMSelectorSig(ctx context.Context, selectorHash, selectorSig string, canonical bool) error {
	if canonical {
		return m.offChainCanonical.StoreEVMSelectorSig(ctx, selectorHash, selectorSig)
	}
	return m.offChainLatest.StoreEVMSelectorSig(ctx, selectorHash, selectorSig)
}

func (m *ZCacheBlockConfirmation) ClearBadAddressCache() {
	// Nothing to do
}

func (m *ZCacheBlockConfirmation) ImplementationType() string {
	return ZCacheImpl + "/" + ZCacheLocalOnly
}
