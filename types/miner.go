package types

import "time"

type MinerEvents struct {
	MinerInfo    []*MinerInfo
	MinerSectors []*MinerSectorEvent
}
type MinerInfo struct {
	ID           string    `json:"id"`
	ActorAddress string    `json:"address"`
	Height       uint64    `json:"height"`
	TxCid        string    `json:"tx_cid"`
	ActionType   string    `json:"action_type"`
	Value        string    `json:"value"`
	TxTimestamp  time.Time `json:"tx_timestamp"`
}

type MinerSectorEvent struct {
	ID           string    `json:"id"`
	MinerAddress string    `json:"miner_address"`
	SectorNumber uint64    `json:"sector_number"`
	Height       uint64    `json:"height"`
	TxCid        string    `json:"tx_cid"`
	ActionType   string    `json:"action_type"`
	Data         string    `json:"data"`
	TxTimestamp  time.Time `json:"tx_timestamp"`
}
