package impl

import (
	"fmt"
	nats2 "github.com/nats-io/nats.go"
	"github.com/zondax/fil-parser/actors/cache/impl/common"
	"go.uber.org/zap"
	"strconv"

	"github.com/filecoin-project/go-address"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/zondax/znats/znats"

	"github.com/zondax/fil-parser/types"
	"gorm.io/gorm"
)

const (
	backFillInProgressKey = "backfill_in_progress"
	ShortRobustStoreName  = "map_actors_short_robust"
	RobustShortStoreName  = "map_actors_robust_short"
	ShortCidStoreName     = "map_actors_short_cid"
	StateStoreName        = "map_actors_backfill_state"

	KvStoreImpl = "kv-store"
)

// KVStore Key-value store cache
type KVStore struct {
	db     *gorm.DB
	nats   *znats.ComponentNats
	Config common.DataSourceConfig
}

func (m *KVStore) ImplementationType() string {
	return KvStoreImpl
}

func (m *KVStore) NewImpl(source common.DataSource) error {
	// Database is mandatory
	if source.Db == nil {
		zap.S().Warn("[ActorsCache] - Database ptr is nil. Database cache is disabled")
		return fmt.Errorf("database ptr is nil")
	}

	m.db = source.Db
	m.Config = source.Config

	// Nats is mandatory
	if source.Config.Nats == nil {
		zap.S().Errorf("[ActorsCache] - Nats ptr is nil. Nats cache is disabled")
		return fmt.Errorf("nats ptr is nil")
	}

	nats, err := znats.NewNatsComponent(*source.Config.Nats)
	if err != nil {
		zap.S().Panicf("[ActorsCache] - Error creating nats component: %s", err.Error())
	}

	m.nats = nats
	// Create kv stores if it does not exist
	commonResourceConfig := znats.CommonResourceConfig{
		Prefixes: []string{m.Config.NetworkName},
		Category: znats.CategorySystem,
	}

	err = m.nats.CreateKVStore(znats.ConfigKVStore{
		CommonResourceConfig: commonResourceConfig,
		KVConfig: &nats2.KeyValueConfig{
			Bucket:      ShortRobustStoreName,
			Description: "Short to Robust address map",
		},
	})

	if err != nil {
		zap.S().Errorf("[ActorsCache] - Error creating short robust kv store: %s", err.Error())
		return err
	}

	err = m.nats.CreateKVStore(znats.ConfigKVStore{
		CommonResourceConfig: commonResourceConfig,
		KVConfig: &nats2.KeyValueConfig{
			Bucket:      RobustShortStoreName,
			Description: "Robust to Short address map",
		},
	})

	if err != nil {
		zap.S().Errorf("[ActorsCache] - Error creating robust short kv store: %s", err.Error())
		return err
	}

	err = m.nats.CreateKVStore(znats.ConfigKVStore{
		CommonResourceConfig: commonResourceConfig,
		KVConfig: &nats2.KeyValueConfig{
			Bucket:      ShortCidStoreName,
			Description: "Short to Actor CID map",
		},
	})

	if err != nil {
		zap.S().Errorf("[ActorsCache] - Error creating short actor type kv store: %s", err.Error())
		return err
	}

	err = m.nats.CreateKVStore(znats.ConfigKVStore{
		CommonResourceConfig: commonResourceConfig,
		KVConfig: &nats2.KeyValueConfig{
			Bucket:      StateStoreName,
			Description: "Actors map state",
		},
	})

	if err != nil {
		zap.S().Errorf("[ActorsCache] - Error creating state kv store: %s", err.Error())
		return err
	}

	m.nats.MapKVStore[StateStoreName].Store.Put(backFillInProgressKey, []byte("false"))

	if m.isNatsCacheEmpty() {
		err = m.BackFill()
		if err != nil {
			zap.S().Errorf("[ActorsCache] - Error backfilling cache: %s", err.Error())
		}
	}

	return nil
}

func (m *KVStore) isNatsCacheEmpty() bool {
	if m.nats == nil {
		return true
	}

	var stores = []string{ShortRobustStoreName, RobustShortStoreName, ShortCidStoreName}
	empty := true

	// All stores should be empty to consider nats cache empty
	for _, store := range stores {
		keys, err := m.nats.MapKVStore[store].Store.Keys()
		if err != nil {
			zap.S().Errorf("[ActorsCache] - Error getting keys from store %s: %s", store, err.Error())
			continue
		}

		if len(keys) > 0 {
			empty = false
			break
		}
	}

	return empty
}

func (m *KVStore) BackFill() error {
	if m.nats == nil {
		return fmt.Errorf("nats is not configured")
	}

	if m.db == nil {
		return fmt.Errorf("database is not configured")
	}

	// Check if backfill is already in progress
	backfillInProgress, err := m.nats.MapKVStore[StateStoreName].Store.Get(backFillInProgressKey)
	if err != nil {
		zap.S().Errorf("[ActorsCache] - Error getting backfill in progress key: %s", err.Error())
		return err
	}

	inProgress, err := strconv.ParseBool(string(backfillInProgress.Value()))
	if err != nil {
		zap.S().Errorf("[ActorsCache] - Error parsing backfill in progress key: %s", err.Error())
		return err
	}

	if inProgress {
		zap.S().Info("[ActorsCache] - Backfill already in progress")
		return nil
	}

	// Set backfill in progress
	_, err = m.nats.MapKVStore[StateStoreName].Store.Put(backFillInProgressKey, []byte("true"))
	if err != nil {
		zap.S().Errorf("[ActorsCache] - Error setting backfill in progress key: %s", err.Error())
		return err
	}

	// Copy the content of the database into the kv store
	addresses := make([]types.AddressInfo, 0)
	m.db.Table(m.Config.InputTableName).Find(&addresses)
	zap.S().Infof("[ActorsCache] - Backfilling %d addresses", len(addresses))

	for _, add := range addresses {
		m.storeShortRobust(add.Short, add.Robust)
		m.storeRobustShort(add.Robust, add.Short)
		m.storeActorCode(add.Short, add.ActorCid)
	}

	// Set backfill finished
	_, err = m.nats.MapKVStore[StateStoreName].Store.Put(backFillInProgressKey, []byte("false"))
	if err != nil {
		zap.S().Errorf("[ActorsCache] - Error setting backfill in progress key: %s", err.Error())
		return err
	}

	zap.S().Info("[ActorsCache] - Backfill finished")
	return nil
}

func (m *KVStore) GetActorCode(address address.Address, key filTypes.TipSetKey) (string, error) {
	shortAddress, err := m.GetShortAddress(address)
	if err != nil {
		fmt.Printf("[ActorsCache] - Error getting short address: %s\n", err.Error())
		return cid.Undef.String(), err
	}

	cid, err := m.nats.MapKVStore[ShortCidStoreName].Store.Get(shortAddress)

	if err != nil {
		if err == nats2.ErrKeyNotFound {
			return "", common.ErrKeyNotFound
		}
		return "", err
	}

	return string(cid.Value()), nil
}

func (m *KVStore) GetRobustAddress(address address.Address) (string, error) {
	isRobustAddress, err := common.IsRobustAddress(address)
	if err != nil {
		return "", err
	}

	if isRobustAddress {
		// Already a robust address
		return address.String(), nil
	}

	// This is a short address, get the robust one
	robustAdd, err := m.nats.MapKVStore[ShortRobustStoreName].Store.Get(address.String())
	if err != nil {
		if err == nats2.ErrKeyNotFound {
			return "", common.ErrKeyNotFound
		}
		return "", err
	}

	return string(robustAdd.Value()), nil
}

func (m *KVStore) GetShortAddress(address address.Address) (string, error) {
	isRobustAddress, err := common.IsRobustAddress(address)
	if err != nil {
		return "", err
	}

	if !isRobustAddress {
		// Already a short address
		return address.String(), nil
	}

	// This is a robust address, get the short one
	shortAdd, err := m.nats.MapKVStore[RobustShortStoreName].Store.Get(address.String())
	if err != nil {
		if err == nats2.ErrKeyNotFound {
			return "", common.ErrKeyNotFound
		}
		return "", err
	}

	return string(shortAdd.Value()), nil
}

func (m *KVStore) storeRobustShort(robust string, short string) {
	if robust == "" || short == "" {
		zap.S().Warn("[ActorsCache] - Trying to store empty robust or short address")
		return
	}

	_, err := m.nats.MapKVStore[RobustShortStoreName].Store.Put(robust, []byte(short))
	if err != nil {
		zap.S().Errorf("[ActorsCache] - Error storing robust short in kv store: %s", err.Error())
	}
}

func (m *KVStore) storeShortRobust(short string, robust string) {
	if robust == "" || short == "" {
		zap.S().Warn("[ActorsCache] - Trying to store empty robust or short address")
		return
	}

	_, err := m.nats.MapKVStore[ShortRobustStoreName].Store.Put(short, []byte(robust))
	if err != nil {
		zap.S().Errorf("[ActorsCache] - Error storing short robust in kv store: %s", err.Error())
	}

}

func (m *KVStore) StoreAddressInfo(info types.AddressInfo) {
	m.storeRobustShort(info.Robust, info.Short)
	m.storeShortRobust(info.Short, info.Robust)
	m.storeActorCode(info.Short, info.ActorCid)
}

func (m *KVStore) storeActorCode(shortAddress string, cid string) {
	if shortAddress == "" || cid == "" {
		zap.S().Warn("[ActorsCache] - Trying to store empty cid or short address")
		return
	}

	_, err := m.nats.MapKVStore[ShortCidStoreName].Store.Put(shortAddress, []byte(cid))
	if err != nil {
		zap.S().Errorf("[ActorsCache] - Error storing actor code in kv store: %s", err.Error())
	}
}
