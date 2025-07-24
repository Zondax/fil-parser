package types

import (
	"math/big"
	"time"
)

type DataCapEvents struct {
	DataCapInfo           []*DataCapInfo
	DataCapTokenEvent     []*DataCapTokenEvent
	DataCapAllowanceEvent []*DataCapAllowanceEvent
}

type DataCapInfo struct {
	ID          string    `json:"id"`
	Address     string    `json:"address"`
	Height      uint64    `json:"height"`
	TxCid       string    `json:"tx_cid"`
	ActionType  string    `json:"action_type"`
	Data        string    `json:"data"`
	TxTimestamp time.Time `json:"tx_timestamp"`
}

type DataCapTokenEvent struct {
	ID          string    `json:"id"`
	Address     string    `json:"address"`
	Height      uint64    `json:"height"`
	TxCid       string    `json:"tx_cid"`
	ActionType  string    `json:"action_type"`
	Balance     *big.Int  `json:"balance"`
	Supply      *big.Int  `json:"supply"`
	Data        string    `json:"data"`
	TxTimestamp time.Time `json:"tx_timestamp"`
}

type DataCapAllowanceEvent struct {
	ID               string    `json:"id"`
	Owner            string    `json:"owner"`
	Operator         string    `json:"operator"`
	Height           uint64    `json:"height"`
	TxCid            string    `json:"tx_cid"`
	ActionType       string    `json:"action_type"`
	AllowanceBalance *big.Int  `json:"allowance_balance"`
	Data             string    `json:"data"`
	TxTimestamp      time.Time `json:"tx_timestamp"`
}
