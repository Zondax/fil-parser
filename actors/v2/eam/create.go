package eam

import (
	"fmt"

	"github.com/filecoin-project/go-state-types/abi"
	eamv10 "github.com/filecoin-project/go-state-types/builtin/v10/eam"
	eamv11 "github.com/filecoin-project/go-state-types/builtin/v11/eam"
	eamv12 "github.com/filecoin-project/go-state-types/builtin/v12/eam"
	eamv13 "github.com/filecoin-project/go-state-types/builtin/v13/eam"
	eamv14 "github.com/filecoin-project/go-state-types/builtin/v14/eam"
	eamv15 "github.com/filecoin-project/go-state-types/builtin/v15/eam"
	eamv16 "github.com/filecoin-project/go-state-types/builtin/v16/eam"
	"github.com/ipfs/go-cid"

	typegen "github.com/whyrusleeping/cbor-gen"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
)

func createParams() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V18.String(): &eamv10.CreateParams{},
		tools.V19.String(): &eamv11.CreateParams{},
		tools.V20.String(): &eamv11.CreateParams{},
		tools.V21.String(): &eamv12.CreateParams{},
		tools.V22.String(): &eamv13.CreateParams{},
		tools.V23.String(): &eamv14.CreateParams{},
		tools.V24.String(): &eamv15.CreateParams{},
		tools.V25.String(): &eamv16.CreateParams{},
	}
}

func createExternalReturn() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V18.String(): &eamv10.CreateExternalReturn{},
		tools.V19.String(): &eamv11.CreateExternalReturn{},
		tools.V20.String(): &eamv11.CreateExternalReturn{},
		tools.V21.String(): &eamv12.CreateExternalReturn{},
		tools.V22.String(): &eamv13.CreateExternalReturn{},
		tools.V23.String(): &eamv14.CreateExternalReturn{},
		tools.V24.String(): &eamv15.CreateExternalReturn{},
		tools.V25.String(): &eamv16.CreateExternalReturn{},
	}
}

func createReturn() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V18.String(): &eamv10.CreateReturn{},
		tools.V19.String(): &eamv11.CreateReturn{},
		tools.V20.String(): &eamv11.CreateReturn{},
		tools.V21.String(): &eamv12.CreateReturn{},
		tools.V22.String(): &eamv13.CreateReturn{},
		tools.V23.String(): &eamv14.CreateReturn{},
		tools.V24.String(): &eamv15.CreateReturn{},
		tools.V25.String(): &eamv16.CreateReturn{},
	}
}

func create2Params() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V18.String(): &eamv10.Create2Params{},
		tools.V19.String(): &eamv11.Create2Params{},
		tools.V20.String(): &eamv11.Create2Params{},
		tools.V21.String(): &eamv12.Create2Params{},
		tools.V22.String(): &eamv13.Create2Params{},
		tools.V23.String(): &eamv14.Create2Params{},
		tools.V24.String(): &eamv15.Create2Params{},
		tools.V25.String(): &eamv16.Create2Params{},
	}
}

func create2Return() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V18.String(): &eamv10.Create2Return{},
		tools.V19.String(): &eamv11.Create2Return{},
		tools.V20.String(): &eamv11.Create2Return{},
		tools.V21.String(): &eamv12.Create2Return{},
		tools.V22.String(): &eamv13.Create2Return{},
		tools.V23.String(): &eamv14.Create2Return{},
		tools.V24.String(): &eamv15.Create2Return{},
		tools.V25.String(): &eamv16.Create2Return{},
	}
}

func (e *Eam) CreateExternal(network string, height int64, rawParams, rawReturn []byte, msgCid cid.Cid) (map[string]interface{}, *types.AddressInfo, error) {
	version := tools.VersionFromHeight(network, height)
	params := abi.CborBytes{}
	returnValue, ok := createExternalReturn()[version.String()]
	if !ok {
		return nil, nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseCreateExternal(rawParams, rawReturn, msgCid, params, returnValue, e.helper)
}

func (e *Eam) Create(network string, height int64, rawParams, rawReturn []byte, msgCid cid.Cid) (map[string]interface{}, *types.AddressInfo, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := createParams()[version.String()]
	if !ok {
		return nil, nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := createReturn()[version.String()]
	if !ok {
		return nil, nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseCreate(rawParams, rawReturn, msgCid, params, returnValue, e.helper)
}

func (e *Eam) Create2(network string, height int64, rawParams, rawReturn []byte, msgCid cid.Cid) (map[string]interface{}, *types.AddressInfo, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := create2Params()[version.String()]
	if !ok {
		return nil, nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := create2Return()[version.String()]
	if !ok {
		return nil, nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseCreate(rawParams, rawReturn, msgCid, params, returnValue, e.helper)
}
