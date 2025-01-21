package multisig

import (
	verifreg10 "github.com/filecoin-project/go-state-types/builtin/v10/verifreg"
	verifreg11 "github.com/filecoin-project/go-state-types/builtin/v11/verifreg"
	verifreg12 "github.com/filecoin-project/go-state-types/builtin/v12/verifreg"
	verifreg13 "github.com/filecoin-project/go-state-types/builtin/v13/verifreg"
	verifreg14 "github.com/filecoin-project/go-state-types/builtin/v14/verifreg"
	verifreg15 "github.com/filecoin-project/go-state-types/builtin/v15/verifreg"
	verifreg8 "github.com/filecoin-project/go-state-types/builtin/v8/verifreg"
	verifreg9 "github.com/filecoin-project/go-state-types/builtin/v9/verifreg"
	"github.com/zondax/fil-parser/tools"
)

func AddVerifierValue(height int64, txMetadata string) (interface{}, error) {
	switch {
	case tools.V8.IsSupported(height):
		return parse[*verifreg8.AddVerifierParams, string](txMetadata, jsonUnmarshaller[*verifreg8.AddVerifierParams])
	case tools.V9.IsSupported(height):
		return parse[*verifreg9.AddVerifierParams, string](txMetadata, jsonUnmarshaller[*verifreg9.AddVerifierParams])
	case tools.V10.IsSupported(height):
		return parse[*verifreg10.AddVerifierParams, string](txMetadata, jsonUnmarshaller[*verifreg10.AddVerifierParams])
	case tools.V11.IsSupported(height):
		return parse[*verifreg11.AddVerifierParams, string](txMetadata, jsonUnmarshaller[*verifreg11.AddVerifierParams])
	case tools.V12.IsSupported(height):
		return parse[*verifreg12.AddVerifierParams, string](txMetadata, jsonUnmarshaller[*verifreg12.AddVerifierParams])
	case tools.V13.IsSupported(height):
		return parse[*verifreg13.AddVerifierParams, string](txMetadata, jsonUnmarshaller[*verifreg13.AddVerifierParams])
	case tools.V14.IsSupported(height):
		return parse[*verifreg14.AddVerifierParams, string](txMetadata, jsonUnmarshaller[*verifreg14.AddVerifierParams])
	case tools.V15.IsSupported(height):
		return parse[*verifreg15.AddVerifierParams, string](txMetadata, jsonUnmarshaller[*verifreg15.AddVerifierParams])
	}
	return nil, nil
}
