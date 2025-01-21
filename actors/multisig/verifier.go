package multisig

import verifreg11 "github.com/filecoin-project/go-state-types/builtin/v11/verifreg"

func AddVerifierValue(height int64, txMetadata string) (interface{}, error) {
	switch height {
	case 11:
		return parse[*verifreg11.AddVerifierParams, string](txMetadata, jsonUnmarshaller[*verifreg11.AddVerifierParams])
	}
	return nil, nil
}
