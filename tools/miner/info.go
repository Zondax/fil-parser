package miner

import (
	"fmt"

	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
)

func (eg *eventGenerator) createMinerInfo(tx *types.Transaction, tipsetCid string) (*types.MinerInfo, error) {
	minerInfo := &types.MinerInfo{
		ID:           tools.BuildId(tipsetCid, tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType),
		ActorAddress: tx.TxTo,
		Height:       tx.Height,
		TxCid:        tx.TxCid,
		ActionType:   tx.TxType,
		Value:        tx.TxMetadata,
	}

	return minerInfo, nil
}
