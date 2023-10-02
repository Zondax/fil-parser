package impl

import (
	"github.com/filecoin-project/go-address"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	cmap "github.com/orcaman/concurrent-map"
	"github.com/zondax/fil-parser/actors/cache/impl/common"
	logger2 "github.com/zondax/fil-parser/logger"
	"github.com/zondax/fil-parser/types"
	"go.uber.org/zap"
)

const InMemoryImpl = "in-memory"

// Memory In-memory database
type Memory struct {
	shortCidMap    cmap.ConcurrentMap
	robustShortMap cmap.ConcurrentMap
	shortRobustMap cmap.ConcurrentMap
	logger         *zap.Logger
}

func (m *Memory) NewImpl(source common.DataSource, logger *zap.Logger) error {
	m.logger = logger2.GetSafeLogger(logger)
	m.shortCidMap = cmap.New()
	m.robustShortMap = cmap.New()
	m.shortRobustMap = cmap.New()

	return nil
}

func (m *Memory) ImplementationType() string {
	return InMemoryImpl
}

func (m *Memory) BackFill() error {
	// Nothing to do
	return nil
}

func (m *Memory) GetActorCode(address address.Address, key filTypes.TipSetKey) (string, error) {
	shortAddress, err := m.GetShortAddress(address)
	if err != nil {
		m.logger.Sugar().Infof("[ActorsCache] - Error getting short address: %s\n", err.Error())
		return cid.Undef.String(), common.ErrKeyNotFound
	}

	// Search in memory cache
	code, ok := m.shortCidMap.Get(shortAddress)
	if !ok {
		return cid.Undef.String(), common.ErrKeyNotFound
	}

	if code == "" {
		return cid.Undef.String(), common.ErrEmptyValue
	}

	return code.(string), nil
}

func (m *Memory) GetRobustAddress(address address.Address) (string, error) {
	isRobustAddress, err := common.IsRobustAddress(address)
	if err != nil {
		return "", err
	}

	if isRobustAddress {
		// Already a robust address
		return address.String(), nil
	}

	// This is a short address, get the robust one
	robustAdd, ok := m.shortRobustMap.Get(address.String())
	if !ok {
		return "", common.ErrKeyNotFound
	}

	if robustAdd == "" {
		return "", common.ErrEmptyValue
	}

	return robustAdd.(string), nil
}

func (m *Memory) GetShortAddress(address address.Address) (string, error) {
	isRobustAddress, err := common.IsRobustAddress(address)
	if err != nil {
		return "", err
	}

	if !isRobustAddress {
		// Already a short address
		return address.String(), nil
	}

	// This is a robust address, get the short one
	shortAdd, ok := m.robustShortMap.Get(address.String())

	if !ok {
		return "", common.ErrKeyNotFound
	}

	if shortAdd == "" {
		return "", common.ErrEmptyValue
	}

	return shortAdd.(string), nil
}

func (m *Memory) storeRobustShort(robust string, short string) {
	if robust == "" || short == "" {
		// zap.S().Warn("[ActorsCache] - Trying to store empty robust or short address")
		return
	}

	m.robustShortMap.Set(robust, short)
}

func (m *Memory) storeShortRobust(short string, robust string) {
	if robust == "" || short == "" {
		// zap.S().Warn("[ActorsCache] - Trying to store empty robust or short address")
		return
	}

	m.shortRobustMap.Set(short, robust)
}

func (m *Memory) StoreAddressInfo(info types.AddressInfo) {
	m.storeRobustShort(info.Robust, info.Short)
	m.storeShortRobust(info.Short, info.Robust)
	m.storeActorCode(info.Short, info.ActorCid)
}

func (m *Memory) storeActorCode(shortAddress string, cid string) {
	if shortAddress == "" || cid == "" {
		// zap.S().Warn("[ActorsCache] - Trying to store empty cid or short address")
		return
	}

	m.shortCidMap.Set(shortAddress, cid)
}
