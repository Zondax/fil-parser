package parser

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/cbor"
)

type controlAddress struct {
	Owner        string   `json:"owner"`
	Worker       string   `json:"worker"`
	ControlAddrs []string `json:"controlAddrs"`
}

type execParams struct {
	CodeCid           string `json:"CodeCid"`
	ConstructorParams string `json:"constructorParams"`
}

type beneficiaryTerm struct {
	Quota      string `json:"quota"`
	UsedQuota  string `json:"usedQuota"`
	Expiration int64  `json:"expiration"`
}
type activeBeneficiary struct {
	Beneficiary string          `json:"beneficiary"`
	Term        beneficiaryTerm `json:"term"`
}

type proposed struct {
	NewBeneficiary        string `json:"newBeneficiary"`
	NewQuota              string `json:"newQuota"`
	NewExpiration         int64  `json:"newExpiration"`
	ApprovedByBeneficiary bool   `json:"approvedByBeneficiary"`
	ApprovedByNominee     bool   `json:"approvedByNominee"`
}

type getBeneficiryReturn struct {
	Active   activeBeneficiary `json:"active"`
	Proposed proposed          `json:"proposed"`
}

type propose struct {
	To     string
	Value  string
	Method string
	Params cbor.Unmarshaler
}

type eamCreateReturn struct {
	ActorId       uint64
	RobustAddress *address.Address
	EthAddress    string
}

type MinerFee struct {
	MinerAddress string
	Amount       string
}

type OverEstimationBurnFee struct {
	BurnAddress string
	Amount      string
}

type BurnFee struct {
	BurnAddress string
	Amount      string
}

type FeesMetadata struct {
	TxType                string
	MinerFee              MinerFee
	OverEstimationBurnFee OverEstimationBurnFee
	BurnFee               BurnFee
}
