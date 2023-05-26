package actors

import (
	"bytes"
	"github.com/zondax/fil-parser/parser"

	"github.com/filecoin-project/go-state-types/builtin/v8/paych"
)

/*
Still needs to parse:

	LockBalance
	Receive
*/
func (p *ActorParser) ParsePaymentchannel(txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt) (map[string]interface{}, error) {
	switch txType {
	case parser.MethodSend:
		return p.parseSend(msg), nil
	case parser.MethodConstructor:
		return p.paymentChannelConstructor(msg.Params)
	case parser.MethodUpdateChannelState:
		return p.updateChannelState(msg.Params)
	case parser.MethodSettle, parser.MethodCollect:
		return p.emptyParamsAndReturn()
	case parser.UnknownStr:
		return p.unknownMetadata(msg.Params, msgRct.Return)
	}
	return map[string]interface{}{}, parser.ErrUnknownMethod
}

func (p *ActorParser) paymentChannelConstructor(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var constructor paych.ConstructorParams
	err := constructor.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = constructor
	return metadata, nil
}

func (p *ActorParser) updateChannelState(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var constructor paych.UpdateChannelStateParams
	err := constructor.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = constructor
	return metadata, nil
}
