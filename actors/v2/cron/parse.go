package cron

import (
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/zondax/fil-parser/parser"
)

type Cron struct{}

func (c *Cron) Name() string {
	return manifest.CronKey
}

func (c *Cron) Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt) (map[string]interface{}, error) {

	return map[string]interface{}{}, parser.ErrUnknownMethod
}
