package evm

import (
	evmv10 "github.com/filecoin-project/go-state-types/builtin/v10/evm"
	evmv11 "github.com/filecoin-project/go-state-types/builtin/v11/evm"
	evmv12 "github.com/filecoin-project/go-state-types/builtin/v12/evm"
	evmv13 "github.com/filecoin-project/go-state-types/builtin/v13/evm"
	evmv14 "github.com/filecoin-project/go-state-types/builtin/v14/evm"
	evmv15 "github.com/filecoin-project/go-state-types/builtin/v15/evm"
	evmv16 "github.com/filecoin-project/go-state-types/builtin/v16/evm"
	evmv17 "github.com/filecoin-project/go-state-types/builtin/v17/evm"
	typegen "github.com/whyrusleeping/cbor-gen"
	"github.com/zondax/fil-parser/tools"
)

var resurrectParams = map[string]func() typegen.CBORUnmarshaler{
	tools.V18.String(): func() typegen.CBORUnmarshaler { return new(evmv10.ResurrectParams) },
	tools.V19.String(): func() typegen.CBORUnmarshaler { return new(evmv11.ResurrectParams) },
	tools.V20.String(): func() typegen.CBORUnmarshaler { return new(evmv11.ResurrectParams) },
	tools.V21.String(): func() typegen.CBORUnmarshaler { return new(evmv12.ResurrectParams) },
	tools.V22.String(): func() typegen.CBORUnmarshaler { return new(evmv13.ResurrectParams) },
	tools.V23.String(): func() typegen.CBORUnmarshaler { return new(evmv14.ResurrectParams) },
	tools.V24.String(): func() typegen.CBORUnmarshaler { return new(evmv15.ResurrectParams) },
	tools.V25.String(): func() typegen.CBORUnmarshaler { return new(evmv16.ResurrectParams) },
	tools.V26.String(): func() typegen.CBORUnmarshaler { return new(evmv17.ResurrectParams) },
}

var delegateCallParams = map[string]func() typegen.CBORUnmarshaler{
	tools.V18.String(): func() typegen.CBORUnmarshaler { return new(evmv10.DelegateCallParams) },
	tools.V19.String(): func() typegen.CBORUnmarshaler { return new(evmv11.DelegateCallParams) },
	tools.V20.String(): func() typegen.CBORUnmarshaler { return new(evmv11.DelegateCallParams) },
	tools.V21.String(): func() typegen.CBORUnmarshaler { return new(evmv12.DelegateCallParams) },
	tools.V22.String(): func() typegen.CBORUnmarshaler { return new(evmv13.DelegateCallParams) },
	tools.V23.String(): func() typegen.CBORUnmarshaler { return new(evmv14.DelegateCallParams) },
	tools.V24.String(): func() typegen.CBORUnmarshaler { return new(evmv15.DelegateCallParams) },
	tools.V25.String(): func() typegen.CBORUnmarshaler { return new(evmv16.DelegateCallParams) },
	tools.V26.String(): func() typegen.CBORUnmarshaler { return new(evmv17.DelegateCallParams) },
}

var getBytecodeReturn = map[string]func() typegen.CBORUnmarshaler{
	tools.V18.String(): func() typegen.CBORUnmarshaler { return new(evmv10.GetBytecodeReturn) },
	tools.V19.String(): func() typegen.CBORUnmarshaler { return new(evmv11.GetBytecodeReturn) },
	tools.V20.String(): func() typegen.CBORUnmarshaler { return new(evmv11.GetBytecodeReturn) },
	tools.V21.String(): func() typegen.CBORUnmarshaler { return new(evmv12.GetBytecodeReturn) },
	tools.V22.String(): func() typegen.CBORUnmarshaler { return new(evmv13.GetBytecodeReturn) },
	tools.V23.String(): func() typegen.CBORUnmarshaler { return new(evmv14.GetBytecodeReturn) },
	tools.V24.String(): func() typegen.CBORUnmarshaler { return new(evmv15.GetBytecodeReturn) },
	tools.V25.String(): func() typegen.CBORUnmarshaler { return new(evmv16.GetBytecodeReturn) },
	tools.V26.String(): func() typegen.CBORUnmarshaler { return new(evmv17.GetBytecodeReturn) },
}

var constructorParams = map[string]func() typegen.CBORUnmarshaler{
	tools.V18.String(): func() typegen.CBORUnmarshaler { return new(evmv10.ConstructorParams) },
	tools.V19.String(): func() typegen.CBORUnmarshaler { return new(evmv11.ConstructorParams) },
	tools.V20.String(): func() typegen.CBORUnmarshaler { return new(evmv11.ConstructorParams) },
	tools.V21.String(): func() typegen.CBORUnmarshaler { return new(evmv12.ConstructorParams) },
	tools.V22.String(): func() typegen.CBORUnmarshaler { return new(evmv13.ConstructorParams) },
	tools.V23.String(): func() typegen.CBORUnmarshaler { return new(evmv14.ConstructorParams) },
	tools.V24.String(): func() typegen.CBORUnmarshaler { return new(evmv15.ConstructorParams) },
	tools.V25.String(): func() typegen.CBORUnmarshaler { return new(evmv16.ConstructorParams) },
	tools.V26.String(): func() typegen.CBORUnmarshaler { return new(evmv17.ConstructorParams) },
}

var getStorageAtParams = map[string]func() typegen.CBORUnmarshaler{
	tools.V18.String(): func() typegen.CBORUnmarshaler { return new(evmv10.GetStorageAtParams) },
	tools.V19.String(): func() typegen.CBORUnmarshaler { return new(evmv11.GetStorageAtParams) },
	tools.V20.String(): func() typegen.CBORUnmarshaler { return new(evmv11.GetStorageAtParams) },
	tools.V21.String(): func() typegen.CBORUnmarshaler { return new(evmv12.GetStorageAtParams) },
	tools.V22.String(): func() typegen.CBORUnmarshaler { return new(evmv13.GetStorageAtParams) },
	tools.V23.String(): func() typegen.CBORUnmarshaler { return new(evmv14.GetStorageAtParams) },
	tools.V24.String(): func() typegen.CBORUnmarshaler { return new(evmv15.GetStorageAtParams) },
	tools.V25.String(): func() typegen.CBORUnmarshaler { return new(evmv16.GetStorageAtParams) },
	tools.V26.String(): func() typegen.CBORUnmarshaler { return new(evmv17.GetStorageAtParams) },
}
