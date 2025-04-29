package paymentChannel

import (
	"context"
	"fmt"
	"io"

	"github.com/zondax/golem/pkg/logger"

	"github.com/filecoin-project/go-state-types/abi"
	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/manifest"
	legacyBuiltin "github.com/filecoin-project/specs-actors/actors/builtin"

	paychv10 "github.com/filecoin-project/go-state-types/builtin/v10/paych"
	paychv11 "github.com/filecoin-project/go-state-types/builtin/v11/paych"
	paychv12 "github.com/filecoin-project/go-state-types/builtin/v12/paych"
	paychv13 "github.com/filecoin-project/go-state-types/builtin/v13/paych"
	paychv14 "github.com/filecoin-project/go-state-types/builtin/v14/paych"
	paychv15 "github.com/filecoin-project/go-state-types/builtin/v15/paych"
	paychv16 "github.com/filecoin-project/go-state-types/builtin/v16/paych"
	paychv8 "github.com/filecoin-project/go-state-types/builtin/v8/paych"
	paychv9 "github.com/filecoin-project/go-state-types/builtin/v9/paych"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

type PaymentChannel struct {
	logger *logger.Logger
}

func New(logger *logger.Logger) *PaymentChannel {
	return &PaymentChannel{
		logger: logger,
	}
}
func (p *PaymentChannel) Name() string {
	return manifest.PaychKey
}

func (*PaymentChannel) StartNetworkHeight() int64 {
	return tools.V1.Height()
}

func legacyMethods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	m := &Msig{}
	return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
		legacyBuiltin.MethodsPaych.Constructor: {
			Name:   parser.MethodConstructor,
			Method: actors.ParseConstructor,
		},
		legacyBuiltin.MethodsPaych.UpdateChannelState: {
			Name:   parser.MethodUpdateChannelState,
			Method: p.UpdateChannelState,
		},
		legacyBuiltin.MethodsPaych.Settle: {
			Name:   parser.MethodSettle,
			Method: actors.ParseEmptyParamsAndReturn,
		},
		legacyBuiltin.MethodsPaych.Collect: {
			Name:   parser.MethodCollect,
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
	tools.V16.String(): actors.CopyMethods(paychv8.Methods),
	tools.V17.String(): actors.CopyMethods(paychv9.Methods),
	tools.V18.String(): actors.CopyMethods(paychv10.Methods),
	tools.V19.String(): actors.CopyMethods(paychv11.Methods),
	tools.V20.String(): actors.CopyMethods(paychv11.Methods),
	tools.V21.String(): actors.CopyMethods(paychv12.Methods),
	tools.V22.String(): actors.CopyMethods(paychv13.Methods),
	tools.V23.String(): actors.CopyMethods(paychv14.Methods),
	tools.V24.String(): actors.CopyMethods(paychv15.Methods),
	tools.V25.String(): actors.CopyMethods(paychv16.Methods),
}

func (p *PaymentChannel) Methods(_ context.Context, network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error) {
	version := tools.VersionFromHeight(network, height)
	methods, ok := methods[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return methods, nil
}

type paymentChannelParams interface {
	UnmarshalCBOR(io.Reader) error
}

func (*PaymentChannel) Constructor(network string, height int64, raw []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := constructorParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parse(raw, params())
}

func (*PaymentChannel) UpdateChannelState(network string, height int64, raw []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := updateChannelStateParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parse(raw, params())
}
