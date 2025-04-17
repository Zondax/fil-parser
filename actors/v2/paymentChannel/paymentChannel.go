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

func (p *PaymentChannel) Methods(_ context.Context, network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error) {
	switch {
	// all legacy version
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V15)...):
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
		}, nil
	case tools.V16.IsSupported(network, height):
		return paychv8.Methods, nil
	case tools.V17.IsSupported(network, height):
		return paychv9.Methods, nil
	case tools.V18.IsSupported(network, height):
		return paychv10.Methods, nil
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return paychv11.Methods, nil
	case tools.V21.IsSupported(network, height):
		return paychv12.Methods, nil
	case tools.V22.IsSupported(network, height):
		return paychv13.Methods, nil
	case tools.V23.IsSupported(network, height):
		return paychv14.Methods, nil
	case tools.V24.IsSupported(network, height):
		return paychv15.Methods, nil
	case tools.V25.IsSupported(network, height):
		return paychv16.Methods, nil
	default:
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
}

type paymentChannelParams interface {
	UnmarshalCBOR(io.Reader) error
}

func (*PaymentChannel) Constructor(network string, height int64, raw []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := constructorParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parse(raw, params)
}

func (*PaymentChannel) UpdateChannelState(network string, height int64, raw []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := updateChannelStateParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parse(raw, params)
}
