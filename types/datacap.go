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
	ID           string    `json:"id"`
	ActorAddress string    `json:"actor_address"`
	Height       uint64    `json:"height"`
	TxCid        string    `json:"tx_cid"`
	ActionType   string    `json:"action_type"`
	Data         string    `json:"data"`
	TxTimestamp  time.Time `json:"tx_timestamp"`
}

type DataCapTokenEvent struct {
	ID           string    `json:"id"`
	ActorAddress string    `json:"actor_address"`
	Height       uint64    `json:"height"`
	TxCid        string    `json:"tx_cid"`
	ActionType   string    `json:"action_type"`
	Balance      *big.Int  `json:"balance" gorm:"column:balance;type:Int256"`
	Supply       *big.Int  `json:"supply" gorm:"column:supply;type:Int256"`
	Data         string    `json:"data"`
	TxTimestamp  time.Time `json:"tx_timestamp"`
}

type DataCapAllowanceEvent struct {
	ID               string    `json:"id"`
	OwnerAddress     string    `json:"owner_address"`
	OperatorAddress  string    `json:"operator_address"`
	Height           uint64    `json:"height"`
	TxCid            string    `json:"tx_cid"`
	ActionType       string    `json:"action_type"`
	AllowanceBalance *big.Int  `json:"allowance_balance" gorm:"column:allowance_balance;type:Int256"`
	Data             string    `json:"data"`
	TxTimestamp      time.Time `json:"tx_timestamp"`
}
