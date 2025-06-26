package types

import "time"

type VerifregEvents struct {
	VerifierInfo []*VerifierInfo
	ClientInfo   []*ClientInfo
	Deals        []*VerifregDeal
}

type VerifierInfo struct {
	ID          string    `json:"id"`
	Address     string    `json:"address"`
	Allowance   uint64    `json:"allowance"`
	TxCid       string    `json:"tx_cid"`
	Height      uint64    `json:"height"`
	TxTimestamp time.Time `json:"tx_timestamp"`
}

type ClientInfo struct {
	ID          string    `json:"id"`
	Verifier    string    `json:"verifier"`
	Address     string    `json:"address"`
	Allowance   uint64    `json:"allowance"`
	TxCid       string    `json:"tx_cid"`
	Height      uint64    `json:"height"`
	TxTimestamp time.Time `json:"tx_timestamp"`
}

type VerifregDeal struct {
}
