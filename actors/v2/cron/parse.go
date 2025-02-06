package cron

import (
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
	"go.uber.org/zap"
)

type Cron struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) *Cron {
	return &Cron{
		logger: logger,
	}
}

func (c *Cron) Name() string {
	return manifest.CronKey
}

func (c *Cron) Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, _ cid.Cid, _ filTypes.TipSetKey) (map[string]interface{}, *types.AddressInfo, error) {
	switch txType {
	case parser.MethodConstructor:
		resp, err := c.Constructor(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodEpochTick:
		resp, err := actors.ParseEmptyParamsAndReturn()
		return resp, nil, err
	case parser.UnknownStr:
		resp, err := actors.ParseUnknownMetadata(msg.Params, msgRct.Return)
		return resp, nil, err
	}
	return map[string]interface{}{}, nil, parser.ErrUnknownMethod
}

func (c *Cron) TransactionTypes() map[string]any {
	return map[string]any{
		parser.MethodConstructor: c.Constructor,
		parser.MethodEpochTick:   actors.ParseEmptyParamsAndReturn,
	}
}
