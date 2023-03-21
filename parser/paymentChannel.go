package parser

import (
	"bytes"

	"github.com/filecoin-project/go-state-types/builtin/v8/paych"
	filTypes "github.com/filecoin-project/lotus/chain/types"
)

func (p *Parser) parsePaymentchannel(txType string, msg *filTypes.Message, msgRct *filTypes.MessageReceipt) (map[string]interface{}, error) {
	switch txType {
	case MethodSend:
		return p.parseSend(msg), nil
	case MethodConstructor:
		return p.paymentChannelConstructor(msg.Params)
	case MethodUpdateChannelState:
		return p.updateChannelState(msg.Params)
	case MethodSettle:
	case MethodCollect:
	case UnknownStr:
		return p.unkmownMetadata(msg.Params, msgRct.Return)
	}
	return map[string]interface{}{}, errUnknownMethod
}

func (p *Parser) paymentChannelConstructor(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var constructor paych.ConstructorParams
	err := constructor.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = constructor
	return metadata, nil
}

func (p *Parser) updateChannelState(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var constructor paych.UpdateChannelStateParams
	err := constructor.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = constructor
	return metadata, nil
}
