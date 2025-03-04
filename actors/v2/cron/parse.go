package cron

import (
	"fmt"

	"github.com/ipfs/go-cid"
	"go.uber.org/zap"

	"github.com/filecoin-project/go-state-types/abi"
	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	legacyBuiltin "github.com/filecoin-project/specs-actors/actors/builtin"

	cronv10 "github.com/filecoin-project/go-state-types/builtin/v10/cron"
	cronv11 "github.com/filecoin-project/go-state-types/builtin/v11/cron"
	cronv12 "github.com/filecoin-project/go-state-types/builtin/v12/cron"
	cronv13 "github.com/filecoin-project/go-state-types/builtin/v13/cron"
	cronv14 "github.com/filecoin-project/go-state-types/builtin/v14/cron"
	cronv15 "github.com/filecoin-project/go-state-types/builtin/v15/cron"
	cronv8 "github.com/filecoin-project/go-state-types/builtin/v8/cron"
	cronv9 "github.com/filecoin-project/go-state-types/builtin/v9/cron"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
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

func (*Cron) StartNetworkHeight() int64 {
	return tools.V1.Height()
}

func (c *Cron) Methods(network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error) {
	switch {
	// all legacy version
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V15)...):
		return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
			legacyBuiltin.MethodsCron.Constructor: {
				Name: parser.MethodConstructor,
			},
			legacyBuiltin.MethodsCron.EpochTick: {
				Name: parser.MethodEpochTick,
			},
		}, nil
	case tools.V16.IsSupported(network, height):
		return cronv8.Methods, nil
	case tools.V17.IsSupported(network, height):
		return cronv9.Methods, nil
	case tools.V18.IsSupported(network, height):
		return cronv10.Methods, nil
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return cronv11.Methods, nil
	case tools.V21.IsSupported(network, height):
		return cronv12.Methods, nil
	case tools.V22.IsSupported(network, height):
		return cronv13.Methods, nil
	case tools.V23.IsSupported(network, height):
		return cronv14.Methods, nil
	case tools.V24.IsSupported(network, height):
		return cronv15.Methods, nil
	default:
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
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
