package eam

import (
	eamv10 "github.com/filecoin-project/go-state-types/builtin/v10/eam"
	eamv11 "github.com/filecoin-project/go-state-types/builtin/v11/eam"
	eamv12 "github.com/filecoin-project/go-state-types/builtin/v12/eam"
	eamv13 "github.com/filecoin-project/go-state-types/builtin/v13/eam"
	eamv14 "github.com/filecoin-project/go-state-types/builtin/v14/eam"
	eamv15 "github.com/filecoin-project/go-state-types/builtin/v15/eam"
	eamv16 "github.com/filecoin-project/go-state-types/builtin/v16/eam"
	eamv17 "github.com/filecoin-project/go-state-types/builtin/v17/eam"
	typegen "github.com/whyrusleeping/cbor-gen"
	"github.com/zondax/fil-parser/tools"
)

var createParams = map[string]func() typegen.CBORUnmarshaler{
	tools.V18.String(): func() typegen.CBORUnmarshaler { return new(eamv10.CreateParams) },
	tools.V19.String(): func() typegen.CBORUnmarshaler { return new(eamv11.CreateParams) },
	tools.V20.String(): func() typegen.CBORUnmarshaler { return new(eamv11.CreateParams) },
	tools.V21.String(): func() typegen.CBORUnmarshaler { return new(eamv12.CreateParams) },
	tools.V22.String(): func() typegen.CBORUnmarshaler { return new(eamv13.CreateParams) },
	tools.V23.String(): func() typegen.CBORUnmarshaler { return new(eamv14.CreateParams) },
	tools.V24.String(): func() typegen.CBORUnmarshaler { return new(eamv15.CreateParams) },
	tools.V25.String(): func() typegen.CBORUnmarshaler { return new(eamv16.CreateParams) },
	tools.V26.String(): func() typegen.CBORUnmarshaler { return new(eamv16.CreateParams) },
	tools.V27.String(): func() typegen.CBORUnmarshaler { return new(eamv17.CreateParams) },
}

var createExternalReturn = map[string]func() typegen.CBORUnmarshaler{
	tools.V18.String(): func() typegen.CBORUnmarshaler { return new(eamv10.CreateExternalReturn) },
	tools.V19.String(): func() typegen.CBORUnmarshaler { return new(eamv11.CreateExternalReturn) },
	tools.V20.String(): func() typegen.CBORUnmarshaler { return new(eamv11.CreateExternalReturn) },
	tools.V21.String(): func() typegen.CBORUnmarshaler { return new(eamv12.CreateExternalReturn) },
	tools.V22.String(): func() typegen.CBORUnmarshaler { return new(eamv13.CreateExternalReturn) },
	tools.V23.String(): func() typegen.CBORUnmarshaler { return new(eamv14.CreateExternalReturn) },
	tools.V24.String(): func() typegen.CBORUnmarshaler { return new(eamv15.CreateExternalReturn) },
	tools.V25.String(): func() typegen.CBORUnmarshaler { return new(eamv16.CreateExternalReturn) },
	tools.V26.String(): func() typegen.CBORUnmarshaler { return new(eamv16.CreateExternalReturn) },
	tools.V27.String(): func() typegen.CBORUnmarshaler { return new(eamv17.CreateExternalReturn) },
}

var createReturn = map[string]func() typegen.CBORUnmarshaler{
	tools.V18.String(): func() typegen.CBORUnmarshaler { return new(eamv10.CreateReturn) },
	tools.V19.String(): func() typegen.CBORUnmarshaler { return new(eamv11.CreateReturn) },
	tools.V20.String(): func() typegen.CBORUnmarshaler { return new(eamv11.CreateReturn) },
	tools.V21.String(): func() typegen.CBORUnmarshaler { return new(eamv12.CreateReturn) },
	tools.V22.String(): func() typegen.CBORUnmarshaler { return new(eamv13.CreateReturn) },
	tools.V23.String(): func() typegen.CBORUnmarshaler { return new(eamv14.CreateReturn) },
	tools.V24.String(): func() typegen.CBORUnmarshaler { return new(eamv15.CreateReturn) },
	tools.V25.String(): func() typegen.CBORUnmarshaler { return new(eamv16.CreateReturn) },
	tools.V26.String(): func() typegen.CBORUnmarshaler { return new(eamv16.CreateReturn) },
	tools.V27.String(): func() typegen.CBORUnmarshaler { return new(eamv17.CreateReturn) },
}

var create2Params = map[string]func() typegen.CBORUnmarshaler{
	tools.V18.String(): func() typegen.CBORUnmarshaler { return new(eamv10.Create2Params) },
	tools.V19.String(): func() typegen.CBORUnmarshaler { return new(eamv11.Create2Params) },
	tools.V20.String(): func() typegen.CBORUnmarshaler { return new(eamv11.Create2Params) },
	tools.V21.String(): func() typegen.CBORUnmarshaler { return new(eamv12.Create2Params) },
	tools.V22.String(): func() typegen.CBORUnmarshaler { return new(eamv13.Create2Params) },
	tools.V23.String(): func() typegen.CBORUnmarshaler { return new(eamv14.Create2Params) },
	tools.V24.String(): func() typegen.CBORUnmarshaler { return new(eamv15.Create2Params) },
	tools.V25.String(): func() typegen.CBORUnmarshaler { return new(eamv16.Create2Params) },
	tools.V26.String(): func() typegen.CBORUnmarshaler { return new(eamv16.Create2Params) },
	tools.V27.String(): func() typegen.CBORUnmarshaler { return new(eamv17.Create2Params) },
}

var create2Return = map[string]func() typegen.CBORUnmarshaler{
	tools.V18.String(): func() typegen.CBORUnmarshaler { return new(eamv10.Create2Return) },
	tools.V19.String(): func() typegen.CBORUnmarshaler { return new(eamv11.Create2Return) },
	tools.V20.String(): func() typegen.CBORUnmarshaler { return new(eamv11.Create2Return) },
	tools.V21.String(): func() typegen.CBORUnmarshaler { return new(eamv12.Create2Return) },
	tools.V22.String(): func() typegen.CBORUnmarshaler { return new(eamv13.Create2Return) },
	tools.V23.String(): func() typegen.CBORUnmarshaler { return new(eamv14.Create2Return) },
	tools.V24.String(): func() typegen.CBORUnmarshaler { return new(eamv15.Create2Return) },
	tools.V25.String(): func() typegen.CBORUnmarshaler { return new(eamv16.Create2Return) },
	tools.V26.String(): func() typegen.CBORUnmarshaler { return new(eamv16.Create2Return) },
	tools.V27.String(): func() typegen.CBORUnmarshaler { return new(eamv17.Create2Return) },
}
