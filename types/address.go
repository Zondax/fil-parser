package types

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
	// CreationTxHash is the tx hash were this actor was created (if applicable)
	CreationTxHash string `json:"creation_tx_hash" gorm:"index:idx_addresses_creation_tx_hash"`
}

type AddressInfoMap map[string]*AddressInfo

func NewAddressInfoMap() AddressInfoMap {
	return make(AddressInfoMap)
}
