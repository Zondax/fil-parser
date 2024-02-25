package types

import (
	"crypto/sha256"
	"encoding/json"
	"math/big"
	"reflect"
	"time"

	"github.com/filecoin-project/lotus/chain/types/ethtypes"
)

// Transaction parses transaction heights into the desired format for reports
type Transaction struct {
	TxBasicBlockData `gorm:"embedded"`
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
	// ParserVersion is the parser version used to parse this tx
	ParserVersion string `json:"parser_version"`
	// FeeData is the fee data
	FeeData string `json:"fee_data,omitempty"`
	NodeInfo
}

func (t Transaction) Equal(b Transaction) bool {
	b.ParentId = t.ParentId
	b.Id = t.Id
	b.ParserVersion = t.ParserVersion
	b.NodeMajorMinorVersion = t.NodeMajorMinorVersion
	b.NodeFullVersion = t.NodeFullVersion
	return reflect.DeepEqual(t, b)
}

type EthLog struct {
	ethtypes.EthLog
	TransactionCid string `json:"transactionCid"`
}

func (t EthLog) GetId() (string, error) {
	h := sha256.New()
	rawData, err := json.Marshal(t)
	if err != nil {
		return "", err
	}

	h.Write(rawData)
	hash := h.Sum(nil)
	return string(hash), nil
}

type GenesisBalances struct {
	Actors struct {
		All []struct {
			Key   string `json:"Key"`
			Value struct {
				Balance string `json:"Balance"`
			}
		}
	}
}
