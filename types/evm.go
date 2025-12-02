package types

type EvmEvents struct {
	Events []EvmEvent
}

type EvmEvent struct {
	Height          uint64
	TxCid           string
	ContractAddress string
	FunctionName    string
	ParsedMetadata  map[string]any
}
