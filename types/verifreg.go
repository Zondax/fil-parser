package types

import (
	"math/big"
	"time"
)

type VerifregEvents struct {
	VerifierInfo []*VerifregEvent
	ClientInfo   []*VerifregClientInfo
	Deals        []*VerifregDeal
}

type VerifregEvent struct {
	ID           string    `json:"id"`
	ActorAddress string    `json:"actor_address"`
	TxCid        string    `json:"tx_cid"`
	Height       uint64    `json:"height"`
	ActionType   string    `json:"action_type"`
	Data         string    `json:"data"`
	TxTimestamp  time.Time `json:"tx_timestamp"`
}
type VerifregClientInfo struct {
	ID            string    `json:"id"`
	ActorAddress  string    `json:"actor_address"`
	ClientAddress string    `json:"client_address"`
	TxCid         string    `json:"tx_cid"`
	Height        uint64    `json:"height"`
	ActionType    string    `json:"action_type"`
	Data          string    `json:"data"`
	Verifiers     []string  `json:"verifiers"`
	DataCap       *big.Int  `json:"datacap" gorm:"column:datacap;type:Int256"`
	TxTimestamp   time.Time `json:"tx_timestamp"`
}

type VerifregDeal struct {
	ID              string    `json:"id"`
	ProviderAddress string    `json:"provider_address"`
	DealID          string    `json:"deal_id"`
	TxCid           string    `json:"tx_cid"`
	Height          uint64    `json:"height"`
	Data            string    `json:"data"`
	TxTimestamp     time.Time `json:"tx_timestamp"`
}
