package actors

import (
	"bytes"
	"github.com/zondax/fil-parser/parser"

	"github.com/filecoin-project/specs-actors/v8/actors/builtin/cron"
)

func ParseCron(txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt) (map[string]interface{}, error) {
	switch txType {
	case parser.MethodConstructor:
		return cronConstructor(msg.Params)
	case parser.MethodEpochTick:
		return emptyParamsAndReturn()
	case parser.UnknownStr:
		return unknownMetadata(msg.Params, msgRct.Return)
	}
	return map[string]interface{}{}, parser.ErrUnknownMethod
}

func cronConstructor(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var constructor cron.ConstructorParams
	err := constructor.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = constructor
	return metadata, nil
}
