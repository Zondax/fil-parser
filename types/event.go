package types

import "time"

const (
	EventTypeEVM    = "evm"
	EventTypeNative = "native"
)

type Event struct {
	BasicBlockData
	ID             string    `json:"id"`
	TxCid          string    `json:"tx_cid"`
	LogIndex       uint64    `json:"log_index"`
	Emitter        string    `json:"emitter"`
	Type           string    `json:"type"`
	SelectorID     string    `json:"selector_id"`
	SelectorSig    string    `json:"selector_sig"`
	Reverted       bool      `json:"reverted"`
	Metadata       string    `json:"metadata"`
	EventTimestamp time.Time `json:"event_timestamp"`
	// ParserVersion is the parser version used to parse this event
	ParserVersion string `json:"parser_version"`
	NodeInfo
}
