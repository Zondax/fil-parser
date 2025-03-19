package paymentChannel

import (
	"context"
	"fmt"
	"io"

	"go.uber.org/zap"

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

	legacyv1 "github.com/filecoin-project/specs-actors/actors/builtin/paych"
	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/paych"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/paych"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/paych"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/paych"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/paych"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/paych"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

type PaymentChannel struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) *PaymentChannel {
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

func (*PaymentChannel) Methods(_ context.Context, network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error) {
	switch {
	// all legacy version
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V15)...):
		return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
			legacyBuiltin.MethodsPaych.Constructor: {
				Name: parser.MethodConstructor,
			},
			legacyBuiltin.MethodsPaych.UpdateChannelState: {
				Name: parser.MethodUpdateChannelState,
			},
			legacyBuiltin.MethodsPaych.Settle: {
				Name: parser.MethodSettle,
			},
			legacyBuiltin.MethodsPaych.Collect: {
				Name: parser.MethodCollect,
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
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return parse(raw, &legacyv1.ConstructorParams{})
	case tools.AnyIsSupported(network, height, tools.V8, tools.V9):
		return parse(raw, &legacyv2.ConstructorParams{})
	case tools.AnyIsSupported(network, height, tools.V10, tools.V11):
		return parse(raw, &legacyv3.ConstructorParams{})
	case tools.V12.IsSupported(network, height):
		return parse(raw, &legacyv4.ConstructorParams{})
	case tools.V13.IsSupported(network, height):
		return parse(raw, &legacyv5.ConstructorParams{})
	case tools.V14.IsSupported(network, height):
		return parse(raw, &legacyv6.ConstructorParams{})
	case tools.V15.IsSupported(network, height):
		return parse(raw, &legacyv7.ConstructorParams{})
	case tools.V16.IsSupported(network, height):
		return parse(raw, &paychv8.ConstructorParams{})
	case tools.V17.IsSupported(network, height):
		return parse(raw, &paychv9.ConstructorParams{})
	case tools.V18.IsSupported(network, height):
		return parse(raw, &paychv10.ConstructorParams{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parse(raw, &paychv11.ConstructorParams{})
	case tools.V21.IsSupported(network, height):
		return parse(raw, &paychv12.ConstructorParams{})
	case tools.V22.IsSupported(network, height):
		return parse(raw, &paychv13.ConstructorParams{})
	case tools.V23.IsSupported(network, height):
		return parse(raw, &paychv14.ConstructorParams{})
	case tools.V24.IsSupported(network, height):
		return parse(raw, &paychv15.ConstructorParams{})
	case tools.V25.IsSupported(network, height):
		return parse(raw, &paychv16.ConstructorParams{})
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*PaymentChannel) UpdateChannelState(network string, height int64, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return parse(raw, &legacyv1.UpdateChannelStateParams{})
	case tools.AnyIsSupported(network, height, tools.V8, tools.V9):
		return parse(raw, &legacyv2.UpdateChannelStateParams{})
	case tools.AnyIsSupported(network, height, tools.V10, tools.V11):
		return parse(raw, &legacyv3.UpdateChannelStateParams{})
	case tools.V12.IsSupported(network, height):
		return parse(raw, &legacyv4.UpdateChannelStateParams{})
	case tools.V13.IsSupported(network, height):
		return parse(raw, &legacyv5.UpdateChannelStateParams{})
	case tools.V14.IsSupported(network, height):
		return parse(raw, &legacyv6.UpdateChannelStateParams{})
	case tools.V15.IsSupported(network, height):
		return parse(raw, &legacyv7.UpdateChannelStateParams{})
	case tools.V16.IsSupported(network, height):
		return parse(raw, &paychv8.UpdateChannelStateParams{})
	case tools.V17.IsSupported(network, height):
		return parse(raw, &paychv9.UpdateChannelStateParams{})
	case tools.V18.IsSupported(network, height):
		return parse(raw, &paychv10.UpdateChannelStateParams{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parse(raw, &paychv11.UpdateChannelStateParams{})
	case tools.V21.IsSupported(network, height):
		return parse(raw, &paychv12.UpdateChannelStateParams{})
	case tools.V22.IsSupported(network, height):
		return parse(raw, &paychv13.UpdateChannelStateParams{})
	case tools.V23.IsSupported(network, height):
		return parse(raw, &paychv14.UpdateChannelStateParams{})
	case tools.V24.IsSupported(network, height):
		return parse(raw, &paychv15.UpdateChannelStateParams{})
	case tools.V25.IsSupported(network, height):
		return parse(raw, &paychv16.UpdateChannelStateParams{})
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}
