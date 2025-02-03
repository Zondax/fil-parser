package account

import (
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/zondax/fil-parser/parser"
)

type Account struct{}

func (a *Account) Name() string {
	return manifest.AccountKey
}

func (a *Account) Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt) (map[string]interface{}, error) {
	return map[string]interface{}{}, parser.ErrUnknownMethod
}
