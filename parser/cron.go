package parser

import (
	"bytes"

	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/specs-actors/v8/actors/builtin/cron"
)

func (p *Parser) parseCron(txType string, msg *filTypes.Message) (map[string]interface{}, error) {
	switch txType {
	case MethodConstructor:
		return p.cronConstructor(msg.Params)
	case MethodEpochTick:
	}
	return map[string]interface{}{}, errUnknownMethod
}

func (p *Parser) cronConstructor(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var constructor cron.ConstructorParams
	err := constructor.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = constructor
	return metadata, nil
}
