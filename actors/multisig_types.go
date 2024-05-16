package actors

type AddSignerValue struct {
	Increase bool   `json:"Increase"`
	Signer   string `json:"Signer"`
}

type ApproveReturn struct {
	Applied bool   `json:"Applied"`
	Code    int    `json:"Code"`
	Ret     string `json:"Ret"`
}

type ApproveValue struct {
	ID           int           `json:"ID"`
	ProposalHash string        `json:"ProposalHash"`
	Return       ApproveReturn `json:"Return"`
}

type CancelValue struct {
	ID           int    `json:"ID"`
	ProposalHash string `json:"ProposalHash"`
}

type ChangeNumApprovalsThresholdValue struct {
	NewThreshold int `json:"NewThreshold"`
}

type ConstructorValue struct {
	Signers               []string `json:"Signers"`
	NumApprovalsThreshold int      `json:"NumApprovalsThreshold"`
	UnlockDuration        int      `json:"UnlockDuration"`
	StartEpoch            int      `json:"StartEpoch"`
}

type LockBalanceValue struct {
	StartEpoch     int    `json:"StartEpoch"`
	UnlockDuration int    `json:"UnlockDuration"`
	Amount         string `json:"Amount"`
}

type RemoveSignerValue struct {
	Signer   string `json:"Signer"`
	Decrease bool   `json:"Decrease"`
}

type SendValue struct {
	Params interface{} `json:"Params"`
}

type SwapSignerValue struct {
	From string `json:"From"`
	To   string `json:"To"`
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
	Type    int                              `json:"Type"`
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
