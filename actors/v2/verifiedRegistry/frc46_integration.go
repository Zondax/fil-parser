package verifiedRegistry

import (
	"bytes"
	"fmt"

	"github.com/filecoin-project/go-address"
	verifreg10 "github.com/filecoin-project/go-state-types/builtin/v10/verifreg"
	cbg "github.com/whyrusleeping/cbor-gen"
	"github.com/zondax/fil-parser/actors/shared/frc46"
	"github.com/zondax/fil-parser/parser"
)

// ParseFRC46UniversalReceiverHook parses FRC46 token data from UniversalReceiverHook
func (v *VerifiedRegistry) ParseFRC46UniversalReceiverHook(network string, height int64, raw, rawReturn []byte, r cbg.CBORUnmarshaler) (map[string]interface{}, error) {
	// Parse the universal receiver parameters
	params, err := frc46.ParseUniversalReceiverParams(raw)
	if err != nil {
		return nil, fmt.Errorf("failed to parse universal receiver params: %w", err)
	}

	// Check if this is an FRC46 token
	if !frc46.IsFRC46Token(params) {
		return nil, fmt.Errorf("not an FRC46 token, got type: %d", params.Type)
	}

	// Parse the FRC46 token payload
	tokenParams, err := frc46.ParseFRC46TokenParams(params.Payload)
	if err != nil {
		return nil, fmt.Errorf("failed to parse FRC46 token params: %w", err)
	}

	from, err := address.NewIDAddress(tokenParams.From)
	if err != nil {
		return nil, fmt.Errorf("failed to parse from address: %w", err)
	}
	to, err := address.NewIDAddress(tokenParams.To)
	if err != nil {
		return nil, fmt.Errorf("failed to parse to address: %w", err)
	}
	operator, err := address.NewIDAddress(tokenParams.Operator)
	if err != nil {
		return nil, fmt.Errorf("failed to parse operator address: %w", err)
	}

	var operatorData interface{} = tokenParams.OperatorData
	var allocReq verifreg10.AllocationRequests // TODO: get the correct version
	if err = allocReq.UnmarshalCBOR(bytes.NewReader(tokenParams.OperatorData)); err == nil {
		operatorData = allocReq
	}

	metadata := map[string]interface{}{
		parser.ParamsKey: map[string]interface{}{
			"type":          "frc46_token",
			"from":          from.String(),
			"to":            to.String(),
			"operator":      operator.String(),
			"amount":        tokenParams.Amount.String(),
			"operator_data": operatorData,
			"token_data":    tokenParams.TokenData,
		},
	}

	if len(rawReturn) > 0 {
		reader := bytes.NewReader(rawReturn)
		err := r.UnmarshalCBOR(reader)
		if err != nil {
			return metadata, fmt.Errorf("error unmarshaling return: %w", err)
		}
		metadata[parser.ReturnKey] = r
	}

	return metadata, nil
}
