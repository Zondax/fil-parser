package evm

import (
	evmv10 "github.com/filecoin-project/go-state-types/builtin/v10/evm"
	evmv11 "github.com/filecoin-project/go-state-types/builtin/v11/evm"
	evmv12 "github.com/filecoin-project/go-state-types/builtin/v12/evm"
	evmv13 "github.com/filecoin-project/go-state-types/builtin/v13/evm"
	evmv14 "github.com/filecoin-project/go-state-types/builtin/v14/evm"
	evmv15 "github.com/filecoin-project/go-state-types/builtin/v15/evm"
	evmv16 "github.com/filecoin-project/go-state-types/builtin/v16/evm"
	typegen "github.com/whyrusleeping/cbor-gen"
	"github.com/zondax/fil-parser/tools"
)

func resurrectParams() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V18.String(): &evmv10.ResurrectParams{},
		tools.V19.String(): &evmv11.ResurrectParams{},
		tools.V20.String(): &evmv11.ResurrectParams{},
		tools.V21.String(): &evmv12.ResurrectParams{},
		tools.V22.String(): &evmv13.ResurrectParams{},
		tools.V23.String(): &evmv14.ResurrectParams{},
		tools.V24.String(): &evmv15.ResurrectParams{},
		tools.V25.String(): &evmv16.ResurrectParams{},
	}
}

func delegateCallParams() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V18.String(): &evmv10.DelegateCallParams{},
		tools.V19.String(): &evmv11.DelegateCallParams{},
		tools.V20.String(): &evmv11.DelegateCallParams{},
		tools.V21.String(): &evmv12.DelegateCallParams{},
		tools.V22.String(): &evmv13.DelegateCallParams{},
		tools.V23.String(): &evmv14.DelegateCallParams{},
		tools.V24.String(): &evmv15.DelegateCallParams{},
		tools.V25.String(): &evmv16.DelegateCallParams{},
	}
}

func getBytecodeReturn() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V18.String(): &evmv10.GetBytecodeReturn{},
		tools.V19.String(): &evmv11.GetBytecodeReturn{},
		tools.V20.String(): &evmv11.GetBytecodeReturn{},
		tools.V21.String(): &evmv12.GetBytecodeReturn{},
		tools.V22.String(): &evmv13.GetBytecodeReturn{},
		tools.V23.String(): &evmv14.GetBytecodeReturn{},
		tools.V24.String(): &evmv15.GetBytecodeReturn{},
		tools.V25.String(): &evmv16.GetBytecodeReturn{},
	}
}

func constructorParams() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V18.String(): &evmv10.ConstructorParams{},
		tools.V19.String(): &evmv11.ConstructorParams{},
		tools.V20.String(): &evmv11.ConstructorParams{},
		tools.V21.String(): &evmv12.ConstructorParams{},
		tools.V22.String(): &evmv13.ConstructorParams{},
		tools.V23.String(): &evmv14.ConstructorParams{},
		tools.V24.String(): &evmv15.ConstructorParams{},
		tools.V25.String(): &evmv16.ConstructorParams{},
	}
}

func getStorageAtParams() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V18.String(): &evmv10.GetStorageAtParams{},
		tools.V19.String(): &evmv11.GetStorageAtParams{},
		tools.V20.String(): &evmv11.GetStorageAtParams{},
		tools.V21.String(): &evmv12.GetStorageAtParams{},
		tools.V22.String(): &evmv13.GetStorageAtParams{},
		tools.V23.String(): &evmv14.GetStorageAtParams{},
		tools.V24.String(): &evmv15.GetStorageAtParams{},
		tools.V25.String(): &evmv16.GetStorageAtParams{},
	}
}
