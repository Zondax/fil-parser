package miner

import (
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/zondax/fil-parser/parser"
)

type Miner struct{}

func (m *Miner) Name() string {
	return manifest.MinerKey
}

func (m *Miner) Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt) (map[string]interface{}, error) {
	return map[string]interface{}{}, parser.ErrUnknownMethod
}
