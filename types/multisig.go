package types

type MultisigInfo struct {
	ID                 string `json:"id"`
	MultisigAddress    string `json:"multisig_address"`
	Height             uint64 `json:"height"`
	TxCid              string `json:"tx_cid"`
	Signer             string `json:"signer"`
	ActionType         string `json:"action_type"`
	Value              string `json:"value"`
	Direction          string `json:"direction"`
	InteractingAddress string `json:"interacting_address"`
}

type MultisigProposal struct {
	ID              string `json:"id"`
	MultisigAddress string `json:"multisig_address"`
	ProposalID      int64  `json:"proposal_id"`
	Height          uint64 `json:"height"`
	TxCid           string `json:"tx_cid"`
	Signer          string `json:"signer"`
	ActionType      string `json:"action_type"`
	Value           string `json:"value"`
}

type MultisigEvents struct {
	Proposals    []*MultisigProposal
	MultisigInfo []*MultisigInfo
}
