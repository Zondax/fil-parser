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
	ID          string    `json:"id"`
	Address     string    `json:"address"`
	TxCid       string    `json:"tx_cid"`
	Height      uint64    `json:"height"`
	ActionType  string    `json:"action_type"`
	Value       string    `json:"value"`
	TxTimestamp time.Time `json:"tx_timestamp"`
}
type VerifregClientInfo struct {
	ID          string    `json:"id"`
	Address     string    `json:"address"`
	Client      string    `json:"client"`
	TxCid       string    `json:"tx_cid"`
	Height      uint64    `json:"height"`
	ActionType  string    `json:"action_type"`
	Value       string    `json:"value"`
	Verifiers   []string  `json:"verifiers"`
	DataCap     *big.Int  `json:"data_cap"`
	TxTimestamp time.Time `json:"tx_timestamp"`
}

type VerifregDeal struct {
	ID          string    `json:"id"`
	DeadID      string    `json:"dead_id"`
	TxCid       string    `json:"tx_cid"`
	Height      uint64    `json:"height"`
	Value       string    `json:"value"`
	TxTimestamp time.Time `json:"tx_timestamp"`
}
