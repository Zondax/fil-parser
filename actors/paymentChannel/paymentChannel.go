package paymentchannel

import (
	"io"

	paychv10 "github.com/filecoin-project/go-state-types/builtin/v10/paych"
	paychv11 "github.com/filecoin-project/go-state-types/builtin/v11/paych"
	paychv12 "github.com/filecoin-project/go-state-types/builtin/v12/paych"
	paychv13 "github.com/filecoin-project/go-state-types/builtin/v13/paych"
	paychv14 "github.com/filecoin-project/go-state-types/builtin/v14/paych"
	paychv15 "github.com/filecoin-project/go-state-types/builtin/v15/paych"
	paychv8 "github.com/filecoin-project/go-state-types/builtin/v8/paych"
	paychv9 "github.com/filecoin-project/go-state-types/builtin/v9/paych"
	"github.com/zondax/fil-parser/tools"
)

type PaymentChannel struct{}
type paymentChannelParams interface {
	UnmarshalCBOR(io.Reader) error
}

func (*PaymentChannel) PaymentChannelConstructor(network string, height int64, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.V16.IsSupported(network, height):
		return parse[*paychv8.ConstructorParams](raw)
	case tools.V17.IsSupported(network, height):
		return parse[*paychv9.ConstructorParams](raw)
	case tools.V18.IsSupported(network, height):
		return parse[*paychv10.ConstructorParams](raw)
	case tools.V20.IsSupported(network, height):
		return parse[*paychv11.ConstructorParams](raw)
	case tools.V21.IsSupported(network, height):
		return parse[*paychv12.ConstructorParams](raw)
	case tools.V22.IsSupported(network, height):
		return parse[*paychv13.ConstructorParams](raw)
	case tools.V23.IsSupported(network, height):
		return parse[*paychv14.ConstructorParams](raw)
	case tools.V24.IsSupported(network, height):
		return parse[*paychv15.ConstructorParams](raw)
	}
	return nil, nil
}

func (*PaymentChannel) UpdateChannelState(network string, height int64, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.V16.IsSupported(network, height):
		return parse[*paychv8.UpdateChannelStateParams](raw)
	case tools.V17.IsSupported(network, height):
		return parse[*paychv9.UpdateChannelStateParams](raw)
	case tools.V18.IsSupported(network, height):
		return parse[*paychv10.UpdateChannelStateParams](raw)
	case tools.V20.IsSupported(network, height):
		return parse[*paychv11.UpdateChannelStateParams](raw)
	case tools.V21.IsSupported(network, height):
		return parse[*paychv12.UpdateChannelStateParams](raw)
	case tools.V22.IsSupported(network, height):
		return parse[*paychv13.UpdateChannelStateParams](raw)
	case tools.V23.IsSupported(network, height):
		return parse[*paychv14.UpdateChannelStateParams](raw)
	case tools.V24.IsSupported(network, height):
		return parse[*paychv15.UpdateChannelStateParams](raw)
	}
	return nil, nil
}
