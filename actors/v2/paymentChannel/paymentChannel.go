package paymentchannel

import (
	"fmt"
	"io"

	paychv10 "github.com/filecoin-project/go-state-types/builtin/v10/paych"
	paychv11 "github.com/filecoin-project/go-state-types/builtin/v11/paych"
	paychv12 "github.com/filecoin-project/go-state-types/builtin/v12/paych"
	paychv13 "github.com/filecoin-project/go-state-types/builtin/v13/paych"
	paychv14 "github.com/filecoin-project/go-state-types/builtin/v14/paych"
	paychv15 "github.com/filecoin-project/go-state-types/builtin/v15/paych"
	paychv8 "github.com/filecoin-project/go-state-types/builtin/v8/paych"
	paychv9 "github.com/filecoin-project/go-state-types/builtin/v9/paych"
	"github.com/filecoin-project/go-state-types/manifest"
	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/paych"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/paych"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/paych"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/paych"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/paych"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/paych"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/tools"
	"go.uber.org/zap"
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

type paymentChannelParams interface {
	UnmarshalCBOR(io.Reader) error
}

func (*PaymentChannel) Constructor(network string, height int64, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V9)...):
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
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*PaymentChannel) UpdateChannelState(network string, height int64, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V9)...):
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
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}
