package placeholder

import (
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
	"go.uber.org/zap"
)

type Placeholder struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) *Placeholder {
	return &Placeholder{
		logger: logger,
	}
}
func (p *Placeholder) Name() string {
	return manifest.PlaceholderKey
}

func (p *Placeholder) Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, mainMsgCid cid.Cid, key filTypes.TipSetKey) (map[string]interface{}, *types.AddressInfo, error) {
	var resp map[string]interface{}
	var err error
	switch txType {
	default:
		resp, err = p.parsePlaceholderAny(msg.Params, msgRct.Return)
	}

	return resp, nil, err
}

func (p *Placeholder) TransactionTypes() map[string]any {
	return map[string]any{}
}

func (p *Placeholder) parsePlaceholderAny(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	metadata[parser.ParamsKey] = rawParams
	metadata[parser.ReturnKey] = rawReturn

	return metadata, nil
}
