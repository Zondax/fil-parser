package verifreg

import (
	"github.com/filecoin-project/go-state-types/abi"
	cid "github.com/ipfs/go-cid"
)

// Type-safe structs for RemoveVerifier metadata
type VerifierSignature struct {
	Type int    `json:"Type"`
	Data string `json:"Data"`
}

type VerifierRequest struct {
	Verifier          string            `json:"Verifier"`
	VerifierSignature VerifierSignature `json:"VerifierSignature"`
}

type RemoveVerifierParams struct {
	VerifiedClientToRemove string          `json:"VerifiedClientToRemove"`
	DataCapAmountToRemove  string          `json:"DataCapAmountToRemove"`
	VerifierRequest1       VerifierRequest `json:"VerifierRequest1"`
	VerifierRequest2       VerifierRequest `json:"VerifierRequest2"`
}

type RemoveVerifierMetadata struct {
	MethodNum string               `json:"MethodNum"`
	Params    RemoveVerifierParams `json:"Params"`
}

// FRC46 Token transaction structures
type AllocationData struct {
	Provider   interface{} `json:"Provider"`
	Data       cid.Cid     `json:"Data"`
	Size       int64       `json:"Size"`
	TermMin    int64       `json:"TermMin"`
	TermMax    int64       `json:"TermMax"`
	Expiration int64       `json:"Expiration"`
}

type AllocationDataWithDealID struct {
	AllocationData
	DealID string `json:"DealID"`
}

type OperatorData struct {
	Allocations []AllocationData        `json:"Allocations"`
	Extensions  []ClaimExtensionRequest `json:"Extensions"`
}
type ClaimExtensionRequest struct {
	// The provider (miner actor) which may claim the allocation.
	Provider abi.ActorID `json:"Provider"`
	// Identifier of the claim to be extended.
	Claim uint64 `json:"Claim"`
	// The new maximum period for which a provider can earn quality-adjusted power
	// for the piece (epochs).
	TermMax abi.ChainEpoch `json:"TermMax"`
}

type FRC46TransactionParams struct {
	Amount          string       `json:"amount"`
	From            string       `json:"from"`
	Operator        string       `json:"operator"`
	OperatorData    OperatorData `json:"operator_data_parsed"`
	OperatorDataStr string       `json:"operator_data"`
	To              string       `json:"to"`
	TokenData       string       `json:"token_data"`
	Type            string       `json:"type"`
}

type AllocationResults struct {
	SuccessCount int         `json:"SuccessCount"`
	FailCodes    interface{} `json:"FailCodes"`
}

type ExtensionResults struct {
	SuccessCount int         `json:"SuccessCount"`
	FailCodes    interface{} `json:"FailCodes"`
}

type FRC46TransactionReturn struct {
	AllocationResults AllocationResults `json:"AllocationResults"`
	ExtensionResults  ExtensionResults  `json:"ExtensionResults"`
	NewAllocations    []int64           `json:"NewAllocations"`
}

type FRC46TransactionMetadata struct {
	MethodNum string                 `json:"MethodNum"`
	Params    FRC46TransactionParams `json:"Params"`
	Return    FRC46TransactionReturn `json:"Return"`
}
