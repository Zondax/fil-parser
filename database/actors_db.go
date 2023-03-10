package database

import (
	"context"
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	cmap "github.com/orcaman/concurrent-map"

	"github.com/zondax/fil-parser/types"
)

var ActorsDB Database

func SetupActorsDatabase(api *api.FullNode) {
	var db Database = &Cache{}
	db.NewImpl(api)
	ActorsDB = db
}

type Database interface {
	NewImpl(*api.FullNode)
	GetActorCode(robustAdd address.Address, height int64, key filTypes.TipSetKey) (cid.Cid, error)
	GetRobustAddress(shortAdd address.Address) (string, error)
	GetShortAddress(robustAdd address.Address) (string, error)
	StoreAddressInfo(info types.AddressInfo)
}

// Cache In-memory database
type Cache struct {
	shortCidMap    cmap.ConcurrentMap
	robustShortMap cmap.ConcurrentMap
	shortRobustMap cmap.ConcurrentMap
	Node           *api.FullNode
}

func (m *Cache) NewImpl(node *api.FullNode) {
	m.shortCidMap = cmap.New()
	m.robustShortMap = cmap.New()
	m.shortRobustMap = cmap.New()
	m.Node = node
}

func (m *Cache) GetActorCode(address address.Address, height int64, key filTypes.TipSetKey) (cid.Cid, error) {
	shortAddress, _ := m.GetShortAddress(address)
	code, ok := m.shortCidMap.Get(shortAddress)
	if !ok {
		var err error
		code, err = m.retrieveActorFromLotus(address, key)
		if err != nil {
			return cid.Cid{}, err
		}
		m.storeActorCode(shortAddress, code.(cid.Cid))
	}

	return code.(cid.Cid), nil
}

func (m *Cache) GetRobustAddress(address address.Address) (string, error) {
	isRobustAddress, err := IsRobustAddress(address)
	if err != nil {
		return "", err
	}

	if isRobustAddress {
		// Already a robust address
		return address.String(), nil
	}

	// This is a short address, get the robust one
	robustAdd, ok := m.shortRobustMap.Get(address.String())
	if ok {
		return robustAdd.(string), nil
	}

	// Address is not in cache, get robust address from lotus
	robustAdd, err = m.retrieveActorPubKeyFromLotus(address, false)
	if err != nil {
		return "", err
	}
	// Must check here because if lotus cannot find the pair, it will return the same address as result
	m.StoreShortRobust(address.String(), robustAdd.(string))
	m.StoreRobustShort(robustAdd.(string), address.String())
	return robustAdd.(string), nil
}

func (m *Cache) GetShortAddress(address address.Address) (string, error) {
	isRobustAddress, err := IsRobustAddress(address)
	if err != nil {
		return "", err
	}

	if !isRobustAddress {
		// Already a short address
		return address.String(), nil
	}

	// This is a robust address, get the short one
	shortAdd, ok := m.robustShortMap.Get(address.String())

	if ok {
		return shortAdd.(string), nil
	}

	// Address is not in cache, get short address from lotus
	shortAdd, err = m.retrieveActorPubKeyFromLotus(address, true)
	if err != nil {
		return address.String(), err
	}

	// Must check here because if lotus cannot find the pair, it will return the same address as result
	m.StoreRobustShort(address.String(), shortAdd.(string))
	m.StoreShortRobust(shortAdd.(string), address.String())
	return shortAdd.(string), nil
}

func (m *Cache) StoreRobustShort(robust string, short string) {
	m.robustShortMap.Set(robust, short)
}

func (m *Cache) StoreShortRobust(short string, robust string) {
	m.shortRobustMap.Set(short, robust)
}

func (m Cache) StoreAddressInfo(info types.AddressInfo) {
	m.StoreRobustShort(info.Robust, info.Short)
	m.StoreShortRobust(info.Short, info.Robust)
	m.storeActorCode(info.Robust, info.ActorCid)
}

func (m *Cache) storeActorCode(shortAddress string, cid cid.Cid) {
	m.shortCidMap.Set(shortAddress, cid)
}

func (m *Cache) retrieveActorFromLotus(add address.Address, key filTypes.TipSetKey) (cid.Cid, error) {
	actor, err := (*m.Node).StateGetActor(context.Background(), add, filTypes.EmptyTSK)
	if err != nil {
		// Try again but using the corresponding tipset Key
		actor, err = (*m.Node).StateGetActor(context.Background(), add, key)
		if err != nil {
			return cid.Cid{}, err
		}
	}

	return actor.Code, nil
}

func (m *Cache) retrieveActorPubKeyFromLotus(add address.Address, reverse bool) (string, error) {
	var key address.Address
	var err error
	if reverse {
		key, err = (*m.Node).StateLookupID(context.Background(), add, filTypes.EmptyTSK)
	} else {
		key, err = (*m.Node).StateAccountKey(context.Background(), add, filTypes.EmptyTSK)
	}

	if err != nil {
		return add.String(), nil
	}
	return key.String(), nil
}

func IsRobustAddress(add address.Address) (bool, error) {
	switch add.Protocol() {
	case address.BLS, address.SECP256K1, address.Actor, address.Delegated:
		return true, nil
	case address.ID:
		return false, nil
	default:
		// Consider unknown type as robust
		return false, fmt.Errorf("unknown address type for %s", add.String())
	}
}
