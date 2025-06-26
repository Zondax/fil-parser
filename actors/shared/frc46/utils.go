package frc46

import (
	"bytes"
	"fmt"

	"github.com/zondax/fil-parser/actors/shared/frc42"
)

// FRC46_TOKEN_TYPE is the standard token type identifier for FRC46 tokens
var FRC46_TOKEN_TYPE uint32

// TODO: not a big fan of this
func init() {
	methodHash, err := frc42.MethodHash("FRC46")
	if err != nil {
		panic(err)
	}
	FRC46_TOKEN_TYPE = uint32(methodHash)
}

// ParseUniversalReceiverParams parses the raw bytes into UniversalReceiverParams
func ParseUniversalReceiverParams(raw []byte) (*UniversalReceiverParams, error) {
	if len(raw) == 0 {
		return nil, fmt.Errorf("empty raw data")
	}

	var params UniversalReceiverParams
	reader := bytes.NewReader(raw)
	if err := params.UnmarshalCBOR(reader); err != nil {
		return nil, fmt.Errorf("failed to unmarshal UniversalReceiverParams: %w", err)
	}

	return &params, nil
}

// ParseFRC46TokenParams parses the payload from UniversalReceiverParams as FRC46TokenParams
func ParseFRC46TokenParams(payload []byte) (*FRC46TokenParams, error) {
	if len(payload) == 0 {
		return nil, fmt.Errorf("empty payload")
	}

	var params FRC46TokenParams
	reader := bytes.NewReader(payload)
	if err := params.UnmarshalCBOR(reader); err != nil {
		return nil, fmt.Errorf("failed to unmarshal FRC46TokenParams: %w", err)
	}

	return &params, nil
}

// IsFRC46Token checks if the universal receiver params contain an FRC46 token type
func IsFRC46Token(params *UniversalReceiverParams) bool {
	return params != nil && params.Type == TokenType(FRC46_TOKEN_TYPE)
}

// CreateUniversalReceiverParams creates UniversalReceiverParams for FRC46 tokens
func CreateUniversalReceiverParams(payload []byte) *UniversalReceiverParams {
	return &UniversalReceiverParams{
		Type:    TokenType(FRC46_TOKEN_TYPE),
		Payload: payload,
	}
}

// MarshalFRC46TokenParams marshals FRC46TokenParams into bytes
func MarshalFRC46TokenParams(params *FRC46TokenParams) ([]byte, error) {
	if params == nil {
		return nil, fmt.Errorf("nil params")
	}

	var buf bytes.Buffer
	if err := params.MarshalCBOR(&buf); err != nil {
		return nil, fmt.Errorf("failed to marshal FRC46TokenParams: %w", err)
	}

	return buf.Bytes(), nil
}

// MarshalUniversalReceiverParams marshals UniversalReceiverParams into bytes
func MarshalUniversalReceiverParams(params *UniversalReceiverParams) ([]byte, error) {
	if params == nil {
		return nil, fmt.Errorf("nil params")
	}

	var buf bytes.Buffer
	if err := params.MarshalCBOR(&buf); err != nil {
		return nil, fmt.Errorf("failed to marshal UniversalReceiverParams: %w", err)
	}

	return buf.Bytes(), nil
}
