package cron

import (
	"context"
	"fmt"

	"github.com/zondax/golem/pkg/logger"

	"github.com/filecoin-project/go-state-types/abi"
	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"

	cronv10 "github.com/filecoin-project/go-state-types/builtin/v10/cron"
	cronv11 "github.com/filecoin-project/go-state-types/builtin/v11/cron"
	cronv12 "github.com/filecoin-project/go-state-types/builtin/v12/cron"
	cronv13 "github.com/filecoin-project/go-state-types/builtin/v13/cron"
	cronv14 "github.com/filecoin-project/go-state-types/builtin/v14/cron"
	cronv15 "github.com/filecoin-project/go-state-types/builtin/v15/cron"
	cronv16 "github.com/filecoin-project/go-state-types/builtin/v16/cron"
	cronv17 "github.com/filecoin-project/go-state-types/builtin/v17/cron"
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

var methods = map[string]map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
	tools.V0.String(): v1Methods(),
	tools.V1.String(): v1Methods(),
	tools.V2.String(): v1Methods(),
	tools.V3.String(): v1Methods(),

	tools.V4.String(): v2Methods(),
	tools.V5.String(): v2Methods(),
	tools.V6.String(): v2Methods(),
	tools.V7.String(): v2Methods(),
	tools.V8.String(): v2Methods(),
	tools.V9.String(): v2Methods(),

	tools.V10.String(): v3Methods(),
	tools.V11.String(): v3Methods(),

	tools.V12.String(): v4Methods(),
	tools.V13.String(): v5Methods(),
	tools.V14.String(): v6Methods(),
	tools.V15.String(): v7Methods(),
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
	tools.V26.String(): actors.CopyMethods(cronv16.Methods),
	tools.V27.String(): actors.CopyMethods(cronv17.Methods),
}

func (c *Cron) Methods(_ context.Context, network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error) {
	version := tools.VersionFromHeight(network, height)
	methods, ok := methods[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return methods, nil
}

func (c *Cron) Parse(_ context.Context, network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, _ cid.Cid, _ filTypes.TipSetKey, canonical bool) (map[string]interface{}, *types.AddressInfo, error) {
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
