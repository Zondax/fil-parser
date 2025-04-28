package datacap

import (
	"fmt"

	actor_tools "github.com/zondax/fil-parser/actors/v2/tools"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

func (*Datacap) BurnExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := burnParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actor_tools.ErrUnsupportedHeight, height)
	}
	returnValue, ok := burnReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actor_tools.ErrUnsupportedHeight, height)
	}
	return parse(raw, rawReturn, true, params(), returnValue(), parser.ParamsKey)
}

func (*Datacap) BurnFromExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := burnFromParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actor_tools.ErrUnsupportedHeight, height)
	}
	returnValue, ok := burnFromReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actor_tools.ErrUnsupportedHeight, height)
	}
	return parse(raw, rawReturn, true, params(), returnValue(), parser.ParamsKey)
}

func (*Datacap) DestroyExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := destroyParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actor_tools.ErrUnsupportedHeight, height)
	}
	returnValue, ok := burnReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actor_tools.ErrUnsupportedHeight, height)
	}
	return parse(raw, rawReturn, true, params(), returnValue(), parser.ParamsKey)
}
