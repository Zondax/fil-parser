package verifiedRegistry

import (
	"bytes"
	"fmt"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/tools"
)

// ParseFRC46TokenOperatorDataReq parses the operatorData and returns the actual verifreg.AllocationsResponse
func (v *VerifiedRegistry) ParseFRC46TokenOperatorDataReq(network string, height int64, operatorData []byte) (any, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := allocationsResponse[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	p := params()
	err := p.UnmarshalCBOR(bytes.NewReader(operatorData))
	if err != nil {
		return nil, err
	}

	return p, nil
}

// ParseFRC46TokenOperatorDataResp parses the operatorData and returns the actual verifreg.AllocationsResponse
func (v *VerifiedRegistry) ParseFRC46TokenOperatorDataResp(network string, height int64, operatorData []byte) (any, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := allocationsResponse[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	p := params()
	err := p.UnmarshalCBOR(bytes.NewReader(operatorData))
	if err != nil {
		return nil, err
	}

	return p, nil
}
