package cache

import (
	"context"

	"github.com/filecoin-project/go-address"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/go-resty/resty/v2"
	cmap "github.com/orcaman/concurrent-map"
	"github.com/zondax/fil-parser/actors/cache/impl/common"
	cacheMetrics "github.com/zondax/fil-parser/actors/cache/metrics"
	"github.com/zondax/fil-parser/types"
	"github.com/zondax/golem/pkg/logger"
	golemBackoff "github.com/zondax/golem/pkg/zhttpclient/backoff"
)

type IActorsCache interface {
	NewImpl(source common.DataSource, logger *logger.Logger, metrics *cacheMetrics.ActorsCacheMetricsClient, backoff *golemBackoff.BackOff) error
	GetActorCode(add address.Address, key filTypes.TipSetKey, onChainOnly, canonical bool) (string, error)
	GetRobustAddress(add address.Address, canonical bool) (string, error)
	GetShortAddress(add address.Address, canonical bool) (string, error)
	StoreAddressInfo(info types.AddressInfo)
	GetEVMSelectorSig(ctx context.Context, selectorHash string, canonical bool) (string, error)
	StoreEVMSelectorSig(ctx context.Context, selectorHash, selectorSig string, canonical bool) error
	IsSystemActor(addr string) bool
	IsGenesisActor(addr string) bool
	BackFill() error
	ClearBadAddressCache()
	ImplementationType() string
}

type ActorsCache struct {
	offChainCache IActorsCache
	onChainCache  IActorsCache
	badAddress    cmap.ConcurrentMap
	logger        *logger.Logger
	httpClient    *resty.Client
	networkName   string
	metrics       *cacheMetrics.ActorsCacheMetricsClient
}
