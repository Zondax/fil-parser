package types

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
