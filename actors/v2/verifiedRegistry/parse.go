package verifiedregistry

import (
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/zondax/fil-parser/parser"
)

type VerifiedRegistry struct{}

func (p *VerifiedRegistry) Name() string {
	return manifest.VerifregKey
}

func (p *VerifiedRegistry) Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt) (map[string]interface{}, error) {

	return map[string]interface{}{}, parser.ErrUnknownMethod
}
