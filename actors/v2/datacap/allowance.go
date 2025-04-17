package datacap

import (
	"fmt"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

func (*Datacap) IncreaseAllowanceExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := increaseAllowanceParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	var r abi.TokenAmount
	return parse(raw, rawReturn, true, params, &r, parser.ParamsKey)
}

func (*Datacap) DecreaseAllowanceExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := decreaseAllowanceParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	var r abi.TokenAmount
	return parse(raw, rawReturn, true, params, &r, parser.ParamsKey)
}

func (*Datacap) RevokeAllowanceExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := revokeAllowanceParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	var r abi.TokenAmount
	return parse(raw, rawReturn, true, params, &r, parser.ParamsKey)
}

func (*Datacap) AllowanceExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := allowanceParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	var r abi.TokenAmount
	return parse(raw, rawReturn, true, params, &r, parser.ParamsKey)
}
