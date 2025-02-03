package paymentchannel

import (
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/zondax/fil-parser/parser"
)

type PaymentChannel struct{}

func (p *PaymentChannel) Name() string {
	return manifest.PaychKey
}

/*
Still needs to parse:

	LockBalance
	Receive
*/
func (p *PaymentChannel) Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt) (map[string]interface{}, error) {
	return map[string]interface{}{}, parser.ErrUnknownMethod
}
