package types

import (
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"time"
)

type ComputeStateOutputV22 struct {
	Root  cid.Cid
	Trace []*InvocResultV22
}

// InvocResultV22 This is a copy of native lotus InvocResult type. We need to copy it because
// we need a modified ExecutionTrace field, and we can't do that in lotus codebase.
type InvocResultV22 struct {
	MsgCid         cid.Cid
	Msg            *types.Message
	MsgRct         *types.MessageReceipt
	GasCost        api.MsgGasCost
	Error          string
	Duration       time.Duration
	ExecutionTrace ExecutionTraceV22
}

type ExecutionTraceV22 struct {
	Msg        *types.Message
	MsgRct     *types.MessageReceipt
	Error      string
	Duration   time.Duration
	GasCharges []*types.GasTrace `json:"-"` // Ignoring this field increases the performance of the json unmarshalling

	Subcalls []ExecutionTraceV22
}
