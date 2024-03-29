package cache

import (
	"github.com/filecoin-project/go-address"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	cmap "github.com/orcaman/concurrent-map"
	"github.com/zondax/fil-parser/actors/cache/impl/common"
	"github.com/zondax/fil-parser/types"
	"go.uber.org/zap"
)

type IActorsCache interface {
	NewImpl(source common.DataSource, logger *zap.Logger) error
	GetActorCode(add address.Address, key filTypes.TipSetKey) (string, error)
	GetRobustAddress(add address.Address) (string, error)
	GetShortAddress(add address.Address) (string, error)
	StoreAddressInfo(info types.AddressInfo)
	BackFill() error
	ImplementationType() string
}

type ActorsCache struct {
	offChainCache IActorsCache
	onChainCache  IActorsCache
	badAddress    cmap.ConcurrentMap
	logger        *zap.Logger
}
