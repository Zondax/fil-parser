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

type paymentChannelParams interface {
	UnmarshalCBOR(io.Reader) error
}

func PaymentChannelConstructor(height int64, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(height):
		return parse[*paychv8.ConstructorParams](raw)
	case tools.V9.IsSupported(height):
		return parse[*paychv9.ConstructorParams](raw)
	case tools.V10.IsSupported(height):
		return parse[*paychv10.ConstructorParams](raw)
	case tools.V11.IsSupported(height):
		return parse[*paychv11.ConstructorParams](raw)
	case tools.V12.IsSupported(height):
		return parse[*paychv12.ConstructorParams](raw)
	case tools.V13.IsSupported(height):
		return parse[*paychv13.ConstructorParams](raw)
	case tools.V14.IsSupported(height):
		return parse[*paychv14.ConstructorParams](raw)
	case tools.V15.IsSupported(height):
		return parse[*paychv15.ConstructorParams](raw)
	}
	return nil, nil
}

func UpdateChannelState(height int64, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(height):
		return parse[*paychv8.UpdateChannelStateParams](raw)
	case tools.V9.IsSupported(height):
		return parse[*paychv9.UpdateChannelStateParams](raw)
	case tools.V10.IsSupported(height):
		return parse[*paychv10.UpdateChannelStateParams](raw)
	case tools.V11.IsSupported(height):
		return parse[*paychv11.UpdateChannelStateParams](raw)
	case tools.V12.IsSupported(height):
		return parse[*paychv12.UpdateChannelStateParams](raw)
	case tools.V13.IsSupported(height):
		return parse[*paychv13.UpdateChannelStateParams](raw)
	case tools.V14.IsSupported(height):
		return parse[*paychv14.UpdateChannelStateParams](raw)
	case tools.V15.IsSupported(height):
		return parse[*paychv15.UpdateChannelStateParams](raw)
	}
	return nil, nil
}
