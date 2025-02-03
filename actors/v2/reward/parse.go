package reward

import (
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/zondax/fil-parser/parser"
)

type Reward struct{}

func (p *Reward) Name() string {
	return manifest.RewardKey
}

func (p *Reward) Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt) (map[string]interface{}, error) {

	return map[string]interface{}{}, parser.ErrUnknownMethod
}
