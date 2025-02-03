package datacap

import (
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/zondax/fil-parser/parser"
)

type Datacap struct{}

func (d *Datacap) Name() string {
	return manifest.DatacapKey
}

func (p *Datacap) Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt) (map[string]interface{}, error) {

	return map[string]interface{}{}, parser.ErrUnknownMethod
}
