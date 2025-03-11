package miner

import (
	"context"
	"fmt"

	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
)

func (eg *eventGenerator) isMinerSectorMessage(actorName, txType string) bool {
	// TODO: implement
	return false
}

func (eg *eventGenerator) createMinerSector(_ context.Context, tx *types.Transaction, tipsetCid string) (*types.MinerSector, error) {
	minerSector := &types.MinerSector{
		ID:           tools.BuildId(tipsetCid, tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType),
		MinerAddress: tx.TxTo,
		Height:       tx.Height,
		TxCid:        tx.TxCid,
		ActionType:   tx.TxType,
		Value:        tx.TxMetadata,
	}

	// TODO: get the sector number from the message

	return minerSector, nil
}
