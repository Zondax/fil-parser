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

	"github.com/ipfs/go-cid"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
)

func (e *Eam) CreateExternal(network string, height int64, rawParams, rawReturn []byte, msgCid cid.Cid) (map[string]interface{}, *types.AddressInfo, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseCreateExternal(rawParams, rawReturn, msgCid, abi.CborBytes{}, &eamv15.CreateExternalReturn{}, e.helper)
	case tools.V23.IsSupported(network, height):
		return parseCreateExternal(rawParams, rawReturn, msgCid, abi.CborBytes{}, &eamv14.CreateExternalReturn{}, e.helper)
	case tools.V22.IsSupported(network, height):
		return parseCreateExternal(rawParams, rawReturn, msgCid, abi.CborBytes{}, &eamv13.CreateExternalReturn{}, e.helper)
	case tools.V21.IsSupported(network, height):
		return parseCreateExternal(rawParams, rawReturn, msgCid, abi.CborBytes{}, &eamv12.CreateExternalReturn{}, e.helper)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseCreateExternal(rawParams, rawReturn, msgCid, abi.CborBytes{}, &eamv11.CreateExternalReturn{}, e.helper)
	case tools.V18.IsSupported(network, height):
		return parseCreateExternal(rawParams, rawReturn, msgCid, abi.CborBytes{}, &eamv10.CreateExternalReturn{}, e.helper)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V17)...):
		return map[string]interface{}{}, nil, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (e *Eam) Create(network string, height int64, rawParams, rawReturn []byte, msgCid cid.Cid) (map[string]interface{}, *types.AddressInfo, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseCreate(rawParams, rawReturn, msgCid, &eamv15.CreateParams{}, &eamv15.CreateReturn{}, e.helper)
	case tools.V23.IsSupported(network, height):
		return parseCreate(rawParams, rawReturn, msgCid, &eamv14.CreateParams{}, &eamv14.CreateReturn{}, e.helper)
	case tools.V22.IsSupported(network, height):
		return parseCreate(rawParams, rawReturn, msgCid, &eamv13.CreateParams{}, &eamv13.CreateReturn{}, e.helper)
	case tools.V21.IsSupported(network, height):
		return parseCreate(rawParams, rawReturn, msgCid, &eamv12.CreateParams{}, &eamv12.CreateReturn{}, e.helper)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseCreate(rawParams, rawReturn, msgCid, &eamv11.CreateParams{}, &eamv11.CreateReturn{}, e.helper)
	case tools.V18.IsSupported(network, height):
		return parseCreate(rawParams, rawReturn, msgCid, &eamv10.CreateParams{}, &eamv10.CreateReturn{}, e.helper)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V17)...):
		return map[string]interface{}{}, nil, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (e *Eam) Create2(network string, height int64, rawParams, rawReturn []byte, msgCid cid.Cid) (map[string]interface{}, *types.AddressInfo, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseCreate(rawParams, rawReturn, msgCid, &eamv15.Create2Params{}, &eamv15.Create2Return{}, e.helper)
	case tools.V23.IsSupported(network, height):
		return parseCreate(rawParams, rawReturn, msgCid, &eamv14.Create2Params{}, &eamv14.Create2Return{}, e.helper)
	case tools.V22.IsSupported(network, height):
		return parseCreate(rawParams, rawReturn, msgCid, &eamv13.Create2Params{}, &eamv13.Create2Return{}, e.helper)
	case tools.V21.IsSupported(network, height):
		return parseCreate(rawParams, rawReturn, msgCid, &eamv12.Create2Params{}, &eamv12.Create2Return{}, e.helper)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseCreate(rawParams, rawReturn, msgCid, &eamv11.Create2Params{}, &eamv11.Create2Return{}, e.helper)
	case tools.V18.IsSupported(network, height):
		return parseCreate(rawParams, rawReturn, msgCid, &eamv10.Create2Params{}, &eamv10.Create2Return{}, e.helper)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V17)...):
		return map[string]interface{}{}, nil, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}
