package types

import (
	"math/big"
	"time"

	"github.com/filecoin-project/lotus/chain/types/ethtypes"
)

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

type EthLog struct {
	ethtypes.EthLog
	TransactionCid string `json:"transactionCid"`
}

type BasicBlockData struct {
	// Height contains the block height
	Height uint64 `json:"height" gorm:"index:idx_height"`
	// TipsetHash contains the tipset hash
	TipsetCid string `json:"tipset_cid" gorm:"index:idx_tipset_cid"`
	// Block Cid
	BlockCid string `json:"block_cid" gorm:"index:idx_block_cid"`
}

// Transaction parses transaction heights into the desired format for reports
type Transaction struct {
	BasicBlockData `gorm:"embedded"`
	// Id is the unique identifier for this transaction
	Id string `json:"id"`
	// ParentId is the parent transaction id
	ParentId string `json:"parent_id"`
	// Level is the nested level of the transaction
	Level uint16 `json:"level"`
	// TxTimestamp is the timestamp of the transaction
	TxTimestamp time.Time `json:"tx_timestamp"`
	// TxCid is the transaction hash
	TxCid string `json:"tx_cid" gorm:"index:idx_transactions_tx_hash"`
	// TxFrom is the sender address
	TxFrom string `json:"tx_from" gorm:"index:idx_transactions_tx_from"`
	// TxTo is the receiver address
	TxTo string `json:"tx_to" gorm:"index:idx_transactions_tx_to"`
	// Amount is the amount of the tx in attoFil
	Amount *big.Int `json:"amount" gorm:"type:numeric"`
	// GasUsed is the total gas used amount in attoFil
	GasUsed uint64 `json:"gas_used"`
	// Status
	Status string `json:"status"`
	// TxType is the message type
	TxType string `json:"tx_type" gorm:"index:idx_tx_type"`
	// TxMetadata is the message metadata
	TxMetadata string `json:"tx_metadata"`
}
