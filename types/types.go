package types

import (
	"math/big"
	"time"

	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	"github.com/ipfs/go-cid"
)

type AddressInfo struct {
	// Short is the address in 'short' format
	Short string `json:"short" gorm:"uniqueIndex:idx_addresses_combination"`
	// Robust is the address in 'robust' format
	Robust string `json:"id_address" gorm:"uniqueIndex:idx_addresses_combination"`
	// EthAddress is the corresponding eth address (if applicable)
	EthAddress string `json:"eth_address" gorm:"index:idx_addresses_eth_address"`
	// ActorCid is the actor's cid for this address
	ActorCid cid.Cid `json:"actor_cid" gorm:"-"`
	// ActorType is the actor's type name of this address
	ActorType string `json:"actor_type"`
	// CreationTxHash is the tx hash were this actor was created (if applicable)
	CreationTxHash string `json:"creation_tx_hash" gorm:"index:idx_addresses_creation_tx_hash"`
}

type AddressInfoMap map[string]AddressInfo

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
	// Hash contains the block hash
	Hash string `json:"hash" gorm:"index:idx_block_hash"`
	// Canonical indicates if this block belongs to the canonical chain
	Canonical bool `json:"canonical"`
}

// Transaction parses transaction data into the desired format for reports
type Transaction struct {
	BasicBlockData `gorm:"embedded"`
	// Level reflects the level that this transaction belongs to inside the trace nest
	Level uint16
	// TxTimestamp is the timestamp of the transaction
	TxTimestamp time.Time `json:"tx_timestamp"`
	// TxHash is the transaction hash
	TxHash string `json:"tx_hash" gorm:"index:idx_transactions_tx_hash"`
	// TxFrom is the sender address
	TxFrom string `json:"tx_from" gorm:"index:idx_transactions_tx_from"`
	// TxTo is the receiver address
	TxTo string `json:"tx_to" gorm:"index:idx_transactions_tx_to"`
	// Amount is the amount of the tx in attoFil
	Amount *big.Int `json:"amount" gorm:"type:numeric"`
	// GasUsed is the total gas used amount in attoFil
	GasUsed int64 `json:"gas_used"`
	// Status
	Status string `json:"status"`
	// TxType is the message type
	TxType string `json:"tx_type" gorm:"index:idx_tx_type"`
	// TxMetadata is the message metadata
	TxMetadata string `json:"tx_metadata"`
	// TxParams contain the transaction params
	TxParams string `json:"tx_params"`
	// TxReturn contains the returned data by the destination actor
	TxReturn string `json:"tx_return"`
}
