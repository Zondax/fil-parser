package types

import "sync"

type AddressInfo struct {
	// Short is the address in 'short' format
	Short string `json:"short" gorm:"uniqueIndex:idx_addresses_combination"`
	// Robust is the address in 'robust' format
	Robust string `json:"robust" gorm:"uniqueIndex:idx_addresses_combination"`
	// EthAddress is the corresponding eth address (if applicable)
	EthAddress string `json:"eth_address" gorm:"index:idx_addresses_eth_address"`
	// ActorCid is the actor's cid for this address
	ActorCid string `json:"actor_cid"`
	// ActorType is the actor's type name of this address
	ActorType string `json:"actor_type"`
	// CreationTxCid is the tx cid were this actor was created (if applicable)
	CreationTxCid string `json:"creation_tx_cid" gorm:"index:idx_addresses_creation_tx_cid"`
}

type AddressInfoMap struct {
	sync.Mutex
	m map[string]*AddressInfo
}

func NewAddressInfoMap() *AddressInfoMap {
	return &AddressInfoMap{
		m: make(map[string]*AddressInfo),
	}
}

func (a *AddressInfoMap) Set(key string, value *AddressInfo) {
	a.Lock()
	defer a.Unlock()
	a.m[key] = value
}

func (a *AddressInfoMap) Get(key string) (*AddressInfo, bool) {
	a.Lock()
	defer a.Unlock()
	val, ok := a.m[key]
	return val, ok
}

func (a *AddressInfoMap) Len() int {
	a.Lock()
	defer a.Unlock()
	return len(a.m)
}

func (a *AddressInfoMap) Range(f func(key string, value *AddressInfo) bool) {
	a.Lock()
	defer a.Unlock()

	for k, v := range a.m {
		if !f(k, v) {
			break
		}
	}
}

func (a *AddressInfoMap) Copy() map[string]*AddressInfo {
	a.Lock()
	defer a.Unlock()

	result := make(map[string]*AddressInfo)
	for k, v := range a.m {
		result[k] = v
	}

	return result
}
