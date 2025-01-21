package paymentchannel

import (
	"bytes"
	"io"

	paychv10 "github.com/filecoin-project/go-state-types/builtin/v10/paych"
	paychv11 "github.com/filecoin-project/go-state-types/builtin/v11/paych"
	paychv12 "github.com/filecoin-project/go-state-types/builtin/v12/paych"
	paychv13 "github.com/filecoin-project/go-state-types/builtin/v13/paych"
	paychv14 "github.com/filecoin-project/go-state-types/builtin/v14/paych"
	paychv15 "github.com/filecoin-project/go-state-types/builtin/v15/paych"
	paychv8 "github.com/filecoin-project/go-state-types/builtin/v8/paych"
	paychv9 "github.com/filecoin-project/go-state-types/builtin/v9/paych"
	"github.com/zondax/fil-parser/parser"
)

type paymentChannelParams interface {
	UnmarshalCBOR(io.Reader) error
}

func PaymentChannelConstructor(height int64, raw []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		return parse[*paychv8.ConstructorParams](raw)
	case 9:
		return parse[*paychv9.ConstructorParams](raw)
	case 10:
		return parse[*paychv10.ConstructorParams](raw)
	case 11:
		return parse[*paychv11.ConstructorParams](raw)
	case 12:
		return parse[*paychv12.ConstructorParams](raw)
	case 13:
		return parse[*paychv13.ConstructorParams](raw)
	case 14:
		return parse[*paychv14.ConstructorParams](raw)
	case 15:
		return parse[*paychv15.ConstructorParams](raw)
	}
	return nil, nil
}

func UpdateChannelState(height int64, raw []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		return parse[*paychv8.UpdateChannelStateParams](raw)
	case 9:
		return parse[*paychv9.UpdateChannelStateParams](raw)
	case 10:
		return parse[*paychv10.UpdateChannelStateParams](raw)
	case 11:
		return parse[*paychv11.UpdateChannelStateParams](raw)
	case 12:
		return parse[*paychv12.UpdateChannelStateParams](raw)
	case 13:
		return parse[*paychv13.UpdateChannelStateParams](raw)
	case 14:
		return parse[*paychv14.UpdateChannelStateParams](raw)
	case 15:
		return parse[*paychv15.UpdateChannelStateParams](raw)
	}
	return nil, nil
}

func parse[T paymentChannelParams](raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var constructor T
	err := constructor.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = constructor
	return metadata, nil
}
