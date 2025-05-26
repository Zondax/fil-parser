package types

import (
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"time"
)

type ComputeStateOutputV1 struct {
	Root  cid.Cid
	Trace []*InvocResultV1
}

// InvocResultV1 This is a copy of native lotus InvocResult type at version v1.22.
// We need to copy it because we cannot have more than one
type InvocResultV1 struct {
	MsgCid         cid.Cid
	Msg            *types.Message
	MsgRct         *types.MessageReceipt
	GasCost        api.MsgGasCost
	Error          string
	Duration       time.Duration
	ExecutionTrace ExecutionTraceV1
}

type ExecutionTraceV1 struct {
	Msg        *types.Message
	MsgRct     *types.MessageReceipt
	Error      string
	Duration   time.Duration
	GasCharges []*types.GasTrace `json:"-"` // Ignoring this field increases the performance of the json unmarshalling

	Subcalls []ExecutionTraceV1
}
