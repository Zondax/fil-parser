package types

import "time"

type VerifregEvents struct {
	VerifierInfo []*VerifierInfo
	ClientInfo   []*ClientInfo
	Deals        []*VerifregDeal
}

type VerifierInfo struct {
	Address     string    `json:"address"`
	Allowance   uint64    `json:"allowance"`
	TxCid       string    `json:"tx_cid"`
	Height      uint64    `json:"height"`
	TxTimestamp time.Time `json:"tx_timestamp"`
}

type ClientInfo struct {
	VerifierAddress string `json:"verifier_address"`
	Address         string `json:"address"`
	Allowance       uint64 `json:"allowance"`
}

type VerifregDeal struct {
}
