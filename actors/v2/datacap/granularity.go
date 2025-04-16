package datacap

import (
	"fmt"

	"github.com/filecoin-project/go-state-types/abi"
	datacapv10 "github.com/filecoin-project/go-state-types/builtin/v10/datacap"
	datacapv11 "github.com/filecoin-project/go-state-types/builtin/v11/datacap"
	datacapv12 "github.com/filecoin-project/go-state-types/builtin/v12/datacap"
	datacapv13 "github.com/filecoin-project/go-state-types/builtin/v13/datacap"
	datacapv14 "github.com/filecoin-project/go-state-types/builtin/v14/datacap"
	datacapv15 "github.com/filecoin-project/go-state-types/builtin/v15/datacap"
	datacapv16 "github.com/filecoin-project/go-state-types/builtin/v16/datacap"

	typegen "github.com/whyrusleeping/cbor-gen"
	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

func granuralityReturn() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V18.String(): new(datacapv10.GranularityReturn),
		tools.V19.String(): new(datacapv11.GranularityReturn),
		tools.V20.String(): new(datacapv11.GranularityReturn),
		tools.V21.String(): new(datacapv12.GranularityReturn),
		tools.V22.String(): new(datacapv13.GranularityReturn),
		tools.V23.String(): new(datacapv14.GranularityReturn),
		tools.V24.String(): new(datacapv15.GranularityReturn),
		tools.V25.String(): new(datacapv16.GranularityReturn),
	}
}

func (*Datacap) GranularityExported(network string, height int64, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	returnValue, ok := granuralityReturn()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parse(rawReturn, nil, false, returnValue, &abi.EmptyValue{}, parser.ReturnKey)
}
