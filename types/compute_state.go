package types

import (
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"time"
)

type FilteredeComputeStateOutput struct {
	Root  cid.Cid
	Trace []*InvocResult
}

type InvocResult struct {
	MsgCid         cid.Cid
	Msg            *types.Message
	MsgRct         *types.MessageReceipt
	GasCost        api.MsgGasCost
	Error          string
	Duration       time.Duration
	ExecutionTrace ExecutionTrace
}

type ExecutionTrace struct {
	Msg        *types.Message
	MsgRct     *types.MessageReceipt
	Error      string
	Duration   time.Duration
	GasCharges []*types.GasTrace `json:"-"` // Ignoring this field increases the performance of the json unmarshalling

	Subcalls []ExecutionTrace
}
