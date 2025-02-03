package market

import (
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/zondax/fil-parser/parser"
)

type Market struct{}

func (p *Market) Name() string {
	return manifest.MarketKey
}

func (p *Market) ParseStoragemarket(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt) (map[string]interface{}, error) {
	return map[string]interface{}{}, parser.ErrUnknownMethod
}
