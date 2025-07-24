package datacap

import (
	"fmt"

	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
)

func (eg *eventGenerator) createDataCapInfo(tx *types.Transaction, tipsetCid string, actorAddress string) (*types.DataCapInfo, error) {
	return &types.DataCapInfo{
		ID:          tools.BuildId(tipsetCid, tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType),
		Address:     actorAddress,
		Height:      tx.Height,
		TxCid:       tipsetCid,
		ActionType:  tx.TxType,
		Data:        tx.TxMetadata,
		TxTimestamp: tx.TxTimestamp,
	}, nil
}
