package deals

import (
	"fmt"

	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
)

func (eg *eventGenerator) createDealMessage(tx *types.Transaction, tipsetCid, actorAddress string) (*types.DealsMessages, error) {
	//#nosec G115
	version := tools.VersionFromHeight(eg.network, int64(tx.Height))

	txType := tx.TxType
	if tx.TxType == parser.MethodActivateDeals {
		if version.NodeVersion() >= tools.V20.NodeVersion() {
			txType = parser.MethodBatchActivateDeals
		}
	}

	dealMessage := &types.DealsMessages{
		ID:           tools.BuildId(tipsetCid, tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType),
		ActorAddress: actorAddress,
		Height:       tx.Height,
		TxCid:        tx.TxCid,
		ActionType:   txType,
		Value:        tx.TxMetadata,
		TxTimestamp:  tx.TxTimestamp,
	}

	return dealMessage, nil
}
