package eam

import (
	"fmt"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
)

func (e *Eam) CreateExternal(network string, height int64, rawParams, rawReturn []byte, msgCid cid.Cid) (map[string]interface{}, *types.AddressInfo, error) {
	version := tools.VersionFromHeight(network, height)
	params := abi.CborBytes{}
	returnValue, ok := createExternalReturn[version.String()]
	if !ok {
		return nil, nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseCreateExternal(rawParams, rawReturn, msgCid, params, returnValue(), e.helper)
}

func (e *Eam) Create(network string, height int64, rawParams, rawReturn []byte, msgCid cid.Cid) (map[string]interface{}, *types.AddressInfo, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := createParams[version.String()]
	if !ok {
		return nil, nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := createReturn[version.String()]
	if !ok {
		return nil, nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseCreate(rawParams, rawReturn, msgCid, params(), returnValue(), e.helper)
}

func (e *Eam) Create2(network string, height int64, rawParams, rawReturn []byte, msgCid cid.Cid) (map[string]interface{}, *types.AddressInfo, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := create2Params[version.String()]
	if !ok {
		return nil, nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := create2Return[version.String()]
	if !ok {
		return nil, nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseCreate(rawParams, rawReturn, msgCid, params(), returnValue(), e.helper)
}
