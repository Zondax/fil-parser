package types

import (
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"time"
)

type ComputeStateOutputV23 struct {
	Root  cid.Cid
	Trace []*InvocResultV23
}

// InvocResult This is a copy of native lotus InvocResult type. We need to copy it because
// we need a modified ExecutionTrace field, and we can't do that in lotus codebase.
type InvocResultV23 struct {
	MsgCid         cid.Cid
	Msg            *types.Message
	MsgRct         *types.MessageReceipt
	GasCost        api.MsgGasCost
	Error          string
	Duration       time.Duration
	ExecutionTrace ExecutionTraceV23
}

// ExecutionTrace This is a copy of native lotus ExecutionTrace type
type ExecutionTraceV23 struct {
	Msg        types.MessageTrace
	MsgRct     types.ReturnTrace
	GasCharges []*types.GasTrace   `cborgen:"maxlen=1000000000" json:"-"`
	Subcalls   []ExecutionTraceV23 `cborgen:"maxlen=1000000000"`
}
