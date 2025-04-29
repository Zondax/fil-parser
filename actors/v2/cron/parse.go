package cron

import (
	"context"
	"fmt"

	"github.com/zondax/golem/pkg/logger"

	"github.com/filecoin-project/go-state-types/abi"
	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	legacyBuiltin "github.com/filecoin-project/specs-actors/actors/builtin"
	"github.com/ipfs/go-cid"

	cronv10 "github.com/filecoin-project/go-state-types/builtin/v10/cron"
	cronv11 "github.com/filecoin-project/go-state-types/builtin/v11/cron"
	cronv12 "github.com/filecoin-project/go-state-types/builtin/v12/cron"
	cronv13 "github.com/filecoin-project/go-state-types/builtin/v13/cron"
	cronv14 "github.com/filecoin-project/go-state-types/builtin/v14/cron"
	cronv15 "github.com/filecoin-project/go-state-types/builtin/v15/cron"
	cronv16 "github.com/filecoin-project/go-state-types/builtin/v16/cron"
	cronv8 "github.com/filecoin-project/go-state-types/builtin/v8/cron"
	cronv9 "github.com/filecoin-project/go-state-types/builtin/v9/cron"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
)

type Cron struct {
	logger *logger.Logger
}

func New(logger *logger.Logger) *Cron {
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

func legacyMethods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
		legacyBuiltin.MethodsCron.Constructor: {
			Name:   parser.MethodConstructor,
			Method: actors.ParseConstructor,
		},
		legacyBuiltin.MethodsCron.EpochTick: {
			Name:   parser.MethodEpochTick,
			Method: actors.ParseEmptyParamsAndReturn,
		},
	}
}

var methods = map[string]map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
	tools.V1.String():  legacyMethods(),
	tools.V2.String():  legacyMethods(),
	tools.V3.String():  legacyMethods(),
	tools.V4.String():  legacyMethods(),
	tools.V5.String():  legacyMethods(),
	tools.V6.String():  legacyMethods(),
	tools.V7.String():  legacyMethods(),
	tools.V8.String():  legacyMethods(),
	tools.V9.String():  legacyMethods(),
	tools.V10.String(): legacyMethods(),
	tools.V11.String(): legacyMethods(),
	tools.V12.String(): legacyMethods(),
	tools.V13.String(): legacyMethods(),
	tools.V14.String(): legacyMethods(),
	tools.V15.String(): legacyMethods(),
	tools.V16.String(): actors.CopyMethods(cronv8.Methods),
	tools.V17.String(): actors.CopyMethods(cronv9.Methods),
	tools.V18.String(): actors.CopyMethods(cronv10.Methods),
	tools.V19.String(): actors.CopyMethods(cronv11.Methods),
	tools.V20.String(): actors.CopyMethods(cronv11.Methods),
	tools.V21.String(): actors.CopyMethods(cronv12.Methods),
	tools.V22.String(): actors.CopyMethods(cronv13.Methods),
	tools.V23.String(): actors.CopyMethods(cronv14.Methods),
	tools.V24.String(): actors.CopyMethods(cronv15.Methods),
	tools.V25.String(): actors.CopyMethods(cronv16.Methods),
}

func (c *Cron) Methods(_ context.Context, network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error) {
	version := tools.VersionFromHeight(network, height)
	methods, ok := methods[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return methods, nil
}

func (c *Cron) Parse(_ context.Context, network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, _ cid.Cid, _ filTypes.TipSetKey) (map[string]interface{}, *types.AddressInfo, error) {
	switch txType {
	case parser.MethodSend:
		resp := actors.ParseSend(msg)
		return resp, nil, nil
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
		parser.MethodSend:        actors.ParseSend,
		parser.MethodConstructor: c.Constructor,
		parser.MethodEpochTick:   actors.ParseEmptyParamsAndReturn,
	}
}
