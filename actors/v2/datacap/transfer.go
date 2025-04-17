package datacap

import (
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

func (*Datacap) TransferExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := transferParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := transferReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return parse(raw, rawReturn, true, params(), returnValue(), parser.ParamsKey)
}

func (*Datacap) BalanceExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	return parse(raw, rawReturn, true, &address.Address{}, &abi.TokenAmount{}, parser.ParamsKey)
}

func (*Datacap) TransferFromExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := transferFromParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := transferFromReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return parse(raw, rawReturn, true, params(), returnValue(), parser.ParamsKey)
}

func (*Datacap) BalanceOf(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	return parse(raw, rawReturn, true, &address.Address{}, &abi.TokenAmount{}, parser.ParamsKey)
}
