package types

import (
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	"math/big"
	"time"
)

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

type EthLog struct {
	ethtypes.EthLog
	TransactionCid string `json:"transactionCid"`
}
