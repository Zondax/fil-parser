package actors

import (
	"bytes"
	"encoding/hex"
	"strings"

	"github.com/filecoin-project/go-address"
	multisig2 "github.com/filecoin-project/go-state-types/builtin/v14/multisig"
	"github.com/zondax/fil-parser/parser"
)

func (p *ActorParser) parseSend(msg *parser.LotusMessage) map[string]interface{} {
	metadata := make(map[string]interface{})
	metadata[parser.ParamsKey] = msg.Params
	return metadata
}

// parseConstructor parse methods with format: *new(func(*address.Address) *abi.EmptyValue)
func (p *ActorParser) parseConstructor(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params address.Address
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params.String()
	return metadata, nil
}

func (p *ActorParser) unknownMetadata(msgParams, msgReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	var params multisig2.ProposeParams

	// Try to parse as ProposeParams
	if err := params.UnmarshalCBOR(bytes.NewReader(msgParams)); err == nil {
		metadata[parser.ParamsKey] = map[string]interface{}{
			"To":     params.To.String(),
			"Value":  params.Value.String(),
			"Method": params.Method,
			"Params": params.Params,
		}
		metadata[parser.ReturnKey] = hex.EncodeToString(msgReturn)

		// Use parser package constants for method types
		if strings.HasPrefix(params.To.String(), "f410") {
			metadata["TxTypeToExecute"] = parser.MethodInvokeEVM
		} else {
			metadata["TxTypeToExecute"] = parser.MethodSend
		}
		return metadata, nil
	}

	// Fallback to original behavior
	if len(msgParams) > 0 {
		metadata[parser.ParamsKey] = hex.EncodeToString(msgParams)
	}
	if len(msgReturn) > 0 {
		metadata[parser.ReturnKey] = hex.EncodeToString(msgReturn)
	}
	return metadata, nil
}

func (p *ActorParser) emptyParamsAndReturn() (map[string]interface{}, error) {
	return make(map[string]interface{}), nil
}
