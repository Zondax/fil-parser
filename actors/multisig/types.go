package multisig

import (
	"encoding/json"
	"io"
)

type WithCBOR struct {
}

func (w *WithCBOR) UnmarshalCBOR(reader io.Reader) error {
	return nil
}

func (w *WithCBOR) MarshalCBOR(writer io.Writer) error {
	return nil
}

type ChangeOwnerAddressParams struct {
	WithCBOR
	Params string `json:"Params"`
}

// I decided omit ethLog because it's not needed
type InvokeContractParams struct {
	WithCBOR
	Params string `json:"Params"`
	Return string `json:"Return"`
}

type ApproveValue struct {
	ID           int    `json:"ID"`
	ProposalHash string `json:"ProposalHash"`
	Return       any    `json:"Return"`
}

type CancelValue struct {
	WithCBOR
	ID           int    `json:"ID"`
	ProposalHash string `json:"ProposalHash"`
}

func mapToStruct(m map[string]interface{}, v interface{}) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

type SendValue struct {
	WithCBOR
	Params interface{} `json:"Params"`
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

type UniversalReceiverHookReturnValue struct {
	AllocationResults UniversalReceiverHookResults `json:"AllocationResults"`
	ExtensionResults  UniversalReceiverHookResults `json:"ExtensionResults"`
	NewAllocations    []int                        `json:"NewAllocations"`
}

type UniversalReceiverHookResults struct {
	SuccessCount int         `json:"SuccessCount"`
	FailCodes    interface{} `json:"FailCodes"`
}
