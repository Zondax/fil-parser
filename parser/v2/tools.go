package v2

import (
	filTypes "github.com/filecoin-project/lotus/chain/types"
)

// LotusMsgToExecutionTraceMsg copies the fields from the lotus Message to the ExecutionTrace Message
// We lose: Version, Nonce, GasFeeCap, GasPremium
func LotusMsgToExecutionTraceMsg(msg *filTypes.Message) *filTypes.MessageTrace {
	return &filTypes.MessageTrace{
		From:     msg.From,
		To:       msg.To,
		Value:    msg.Value,
		Method:   msg.Method,
		Params:   msg.Params,
		GasLimit: uint64(msg.GasLimit),
	}
}

// LotusMsgRctToExecutionTraceMsgRct copies the fields from the lotus MessageReceipt to the ExecutionTrace MessageReceipt
// We lose: GasUsed, EventsRoot, version
func LotusMsgRctToExecutionTraceMsgRct(msgRct *filTypes.MessageReceipt) *filTypes.ReturnTrace {
	return &filTypes.ReturnTrace{
		ExitCode: msgRct.ExitCode,
		Return:   msgRct.Return,
	}
}
