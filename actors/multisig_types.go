package actors

import (
	"io"

	multisig2 "github.com/filecoin-project/go-state-types/builtin/v14/multisig"
)

type ApproveValue struct {
	ID           int                     `json:"ID"`
	ProposalHash string                  `json:"ProposalHash"`
	Return       multisig2.ApproveReturn `json:"Return"`
}

type CancelValue struct {
	ID           int    `json:"ID"`
	ProposalHash string `json:"ProposalHash"`
}

type SendValue struct {
	Params interface{} `json:"Params"`
}

type UniversalReceiverHookReturnValue struct {
	AllocationResults UniversalReceiverHookResults `json:"AllocationResults"`
	ExtensionResults  UniversalReceiverHookResults `json:"ExtensionResults"`
	NewAllocations    []int                        `json:"NewAllocations"`
}

type UniversalReceiverHookResults struct {
	SuccessCount int         `json:"SuccessCount"`
	FailCodes    interface{} `json:"FailCodes"`
}

type UniversalReceiverHookValue struct {
	Type    uint64                           `json:"Type"`
	Payload string                           `json:"Payload"`
	Return  UniversalReceiverHookReturnValue `json:"Return"`
}

type UniversalReceiverHookParams struct {
	Type_   int    `json:"Type_"`
	Payload string `json:"Payload"`
}

type TransactionUniversalReceiverHookMetadata struct {
	Params string                           `json:"Params"`
	Return UniversalReceiverHookReturnValue `json:"Return"`
}

type ChangeOwnerAddressParams struct {
	WithCBOR
	Params string `json:"Params"`
}

// I decided omit ethLog because it's not needed
type InvokeContractParams struct {
	Params string `json:"Params"`
	Return string `json:"Return"`
}

type WithCBOR struct {
}

func (w *WithCBOR) UnmarshalCBOR(reader io.Reader) error {
	return nil
}

func (w *WithCBOR) MarshalCBOR(writer io.Writer) error {
	return nil
}
