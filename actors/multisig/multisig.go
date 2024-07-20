package multisig

import (
	"context"

	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
)

type MultisigParser interface {
	ParseMultisig(txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, height int64, key filTypes.TipSetKey) (map[string]interface{}, error)
}

type EventGenerator interface {
	GenerateMultisigEvents(ctx context.Context, transactions []*types.Transaction, tipsetCid string, tipsetKey filTypes.TipSetKey) (*types.MultisigEvents, error)
}
