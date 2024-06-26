package types

type MultisigInfo struct {
	ID              string `json:"id"`
	MultisigAddress string `json:"multisig_address"`
	Height          uint64 `json:"height"`
	TxCid           string `json:"tx_cid"`
	ActionType      string `json:"action_type"`
	Value           string `json:"value"`
	Signer          string `json:"signer"`
}

type MultisigProposal struct {
	ID              string `json:"id"`
	MultisigAddress string `json:"multisig_address"`
	ProposalID      int64  `json:"proposal_id"`
	Height          uint64 `json:"height"`
	TxCid           string `json:"tx_cid"`
	Signer          string `json:"signer"`
	ActionType      string `json:"action_type"`
	TxTypeToExecute string `json:"tx_type_to_execute"`
	Value           string `json:"value"`
}

type MultisigEvents struct {
	Proposals    []*MultisigProposal
	MultisigInfo []*MultisigInfo
}
