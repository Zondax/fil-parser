package eam

import (
	"fmt"

	eamv10 "github.com/filecoin-project/go-state-types/builtin/v10/eam"
	eamv11 "github.com/filecoin-project/go-state-types/builtin/v11/eam"
	eamv12 "github.com/filecoin-project/go-state-types/builtin/v12/eam"
	eamv13 "github.com/filecoin-project/go-state-types/builtin/v13/eam"
	eamv14 "github.com/filecoin-project/go-state-types/builtin/v14/eam"
	eamv15 "github.com/filecoin-project/go-state-types/builtin/v15/eam"
	eamv9 "github.com/filecoin-project/go-state-types/builtin/v9/eam"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
)

func ParseCreateExternal(height int64, rawParams, rawReturn []byte, msgCid cid.Cid) (map[string]interface{}, *types.AddressInfo, error) {
	external := true
	switch {
	case tools.V15.IsSupported(height):
		return parseCreate[*eamv15.CreateExternalReturn](rawParams, rawReturn, msgCid, external)
	case tools.V14.IsSupported(height):
		return parseCreate[*eamv14.CreateExternalReturn](rawParams, rawReturn, msgCid, external)
	case tools.V13.IsSupported(height):
		return parseCreate[*eamv13.CreateExternalReturn](rawParams, rawReturn, msgCid, external)
	case tools.V12.IsSupported(height):
		return parseCreate[*eamv12.CreateExternalReturn](rawParams, rawReturn, msgCid, external)
	case tools.V11.IsSupported(height):
		return parseCreate[*eamv11.CreateExternalReturn](rawParams, rawReturn, msgCid, external)
	case tools.V10.IsSupported(height):
		return parseCreate[*eamv10.CreateExternalReturn](rawParams, rawReturn, msgCid, external)
	case tools.V9.IsSupported(height):
		return parseCreate[*eamv9.CreateExternalReturn](rawParams, rawReturn, msgCid, external)
	}
	return nil, nil, fmt.Errorf("unsupported height: %d", height)
}

func ParseCreate(height int64, rawParams, rawReturn []byte, msgCid cid.Cid) (map[string]interface{}, *types.AddressInfo, error) {
	external := false
	switch {
	case tools.V15.IsSupported(height):
		return parseCreate[*eamv15.CreateReturn](rawParams, rawReturn, msgCid, external)
	case tools.V14.IsSupported(height):
		return parseCreate[*eamv14.CreateReturn](rawParams, rawReturn, msgCid, external)
	case tools.V13.IsSupported(height):
		return parseCreate[*eamv13.CreateReturn](rawParams, rawReturn, msgCid, external)
	case tools.V12.IsSupported(height):
		return parseCreate[*eamv12.CreateReturn](rawParams, rawReturn, msgCid, external)
	case tools.V11.IsSupported(height):
		return parseCreate[*eamv11.CreateReturn](rawParams, rawReturn, msgCid, external)
	case tools.V10.IsSupported(height):
		return parseCreate[*eamv10.CreateReturn](rawParams, rawReturn, msgCid, external)
	case tools.V9.IsSupported(height):
		return parseCreate[*eamv9.CreateReturn](rawParams, rawReturn, msgCid, external)
	}
	return nil, nil, fmt.Errorf("unsupported height: %d", height)
}

func ParseCreate2(height int64, rawParams, rawReturn []byte, msgCid cid.Cid) (map[string]interface{}, *types.AddressInfo, error) {
	external := false
	switch {
	case tools.V15.IsSupported(height):
		return parseCreate[*eamv15.Create2Return](rawParams, rawReturn, msgCid, external)
	case tools.V14.IsSupported(height):
		return parseCreate[*eamv14.Create2Return](rawParams, rawReturn, msgCid, external)
	case tools.V13.IsSupported(height):
		return parseCreate[*eamv13.Create2Return](rawParams, rawReturn, msgCid, external)
	case tools.V12.IsSupported(height):
		return parseCreate[*eamv12.Create2Return](rawParams, rawReturn, msgCid, external)
	case tools.V11.IsSupported(height):
		return parseCreate[*eamv11.Create2Return](rawParams, rawReturn, msgCid, external)
	case tools.V10.IsSupported(height):
		return parseCreate[*eamv10.Create2Return](rawParams, rawReturn, msgCid, external)
	case tools.V9.IsSupported(height):
		return parseCreate[*eamv9.Create2Return](rawParams, rawReturn, msgCid, external)
	}
	return nil, nil, fmt.Errorf("unsupported height: %d", height)
}
