package parser

import (
	"encoding/json"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/cbor"
	"github.com/filecoin-project/go-state-types/exitcode"
	"github.com/filecoin-project/lotus/api"
	"github.com/ipfs/go-cid"
)

type ControlAddress struct {
	Owner        string   `json:"owner"`
	Worker       string   `json:"worker"`
	ControlAddrs []string `json:"controlAddrs"`
}

type ExecParams struct {
	CodeCid           string `json:"CodeCid"`
	ConstructorParams string `json:"constructorParams"`
}

type Exec4Params struct {
	CodeCid           string `json:"CodeCid"`
	ConstructorParams string `json:"constructorParams"`
	SubAddress        string `json:"subAddress"`
}

type BeneficiaryTerm struct {
	Quota      string `json:"quota"`
	UsedQuota  string `json:"usedQuota"`
	Expiration int64  `json:"expiration"`
}
type ActiveBeneficiary struct {
	Beneficiary string          `json:"beneficiary"`
	Term        BeneficiaryTerm `json:"term"`
}

type Proposed struct {
	NewBeneficiary        string `json:"newBeneficiary"`
	NewQuota              string `json:"newQuota"`
	NewExpiration         int64  `json:"newExpiration"`
	ApprovedByBeneficiary bool   `json:"approvedByBeneficiary"`
	ApprovedByNominee     bool   `json:"approvedByNominee"`
}

type GetBeneficiaryReturn struct {
	Active   ActiveBeneficiary `json:"active"`
	Proposed Proposed          `json:"proposed"`
}

// TODO: look how to combine these two proposes
type MultisigPropose struct {
	To     string
	Value  string
	Method string
	Params map[string]interface{}
}

type Propose struct {
	To     string
	Value  string
	Method string
	Params cbor.Unmarshaler
}

type EamCreateReturn struct {
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
	TxType                string `json:"TxType,omitempty"`
	MinerFee              MinerFee
	OverEstimationBurnFee OverEstimationBurnFee
	BurnFee               BurnFee
	TotalCost             string
}

type LotusMessage struct {
	To     address.Address
	From   address.Address
	Method abi.MethodNum
	Cid    cid.Cid
	Params []byte
}

type RawLotusMessage LotusMessage

type mCid struct {
	*RawLotusMessage
	CID cid.Cid
}

func (m *LotusMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(&mCid{
		RawLotusMessage: (*RawLotusMessage)(m),
		CID:             m.Cid,
	})
}

type LotusMessageReceipt struct {
	ExitCode exitcode.ExitCode
	Return   []byte
}

type ComputeOutputVersioned struct {
	api.ComputeStateOutput
	Version string
}
