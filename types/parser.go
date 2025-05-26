package types

import (
	filTypes "github.com/filecoin-project/lotus/chain/types"
)

type TxsData struct {
	Traces   []byte
	Tipset   *ExtendedTipSet
	EthLogs  []EthLog
	Metadata BlockMetadata
}

type TxsParsedResult struct {
	Txs       []*Transaction
	Addresses *AddressInfoMap
	TxCids    []TxCidTranslation
}

type EventsData struct {
	Tipset    *ExtendedTipSet
	NativeLog []*filTypes.ActorEvent
	EthLogs   []EthLog
	Metadata  BlockMetadata
}

type EventsParsedResult struct {
	EVMEvents    int
	NativeEvents int
	ParsedEvents []*Event
}
