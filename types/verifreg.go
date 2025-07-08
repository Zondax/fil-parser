package types

import "time"

type VerifregEvents struct {
	VerifierInfo []*VerifregEvent
	ClientInfo   []*VerifregEvent
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

type VerifregDeal struct {
	ID          string    `json:"id"`
	DeadID      string    `json:"dead_id"`
	TxCid       string    `json:"tx_cid"`
	Height      uint64    `json:"height"`
	Value       string    `json:"value"`
	TxTimestamp time.Time `json:"tx_timestamp"`
}
