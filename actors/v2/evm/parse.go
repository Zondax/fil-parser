package evm

import (
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/zondax/fil-parser/parser"
)

type Evm struct{}

func (p *Evm) Name() string {
	return manifest.EvmKey
}

func (p *Evm) Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})

	return metadata, nil
}
