package datacap

import (
	"fmt"

	"github.com/filecoin-project/go-state-types/abi"

	actor_tools "github.com/zondax/fil-parser/actors/v2/tools"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

func (*Datacap) GranularityExported(network string, height int64, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	returnValue, ok := granularityReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actor_tools.ErrUnsupportedHeight, height)
	}
	return parse(rawReturn, nil, false, returnValue(), &abi.EmptyValue{}, parser.ReturnKey)
}
