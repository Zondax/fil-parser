package types

import "time"

type DealsEvents struct {
	DealsInfo []*DealsInfo
}
type DealsInfo struct {
	ID           string    `json:"id"`
	ActorAddress string    `json:"address"`
	Height       uint64    `json:"height"`
	TxCid        string    `json:"tx_cid"`
	ActionType   string    `json:"action_type"`
	Value        string    `json:"value"`
	TxTimestamp  time.Time `json:"tx_timestamp"`
}
