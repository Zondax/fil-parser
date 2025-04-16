package datacap

import (
	"fmt"

	datacapv10 "github.com/filecoin-project/go-state-types/builtin/v10/datacap"
	datacapv11 "github.com/filecoin-project/go-state-types/builtin/v11/datacap"
	datacapv12 "github.com/filecoin-project/go-state-types/builtin/v12/datacap"
	datacapv13 "github.com/filecoin-project/go-state-types/builtin/v13/datacap"
	datacapv14 "github.com/filecoin-project/go-state-types/builtin/v14/datacap"
	datacapv15 "github.com/filecoin-project/go-state-types/builtin/v15/datacap"
	datacapv16 "github.com/filecoin-project/go-state-types/builtin/v16/datacap"
	datacapv9 "github.com/filecoin-project/go-state-types/builtin/v9/datacap"

	typegen "github.com/whyrusleeping/cbor-gen"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

func mintParams() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V17.String(): &datacapv9.MintParams{},
		tools.V18.String(): &datacapv10.MintParams{},
		tools.V19.String(): &datacapv11.MintParams{},
		tools.V20.String(): &datacapv11.MintParams{},
		tools.V21.String(): &datacapv12.MintParams{},
		tools.V22.String(): &datacapv13.MintParams{},
		tools.V23.String(): &datacapv14.MintParams{},
		tools.V24.String(): &datacapv15.MintParams{},
		tools.V25.String(): &datacapv16.MintParams{},
	}
}

func mintReturn() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V17.String(): &datacapv9.MintReturn{},
		tools.V18.String(): &datacapv10.MintReturn{},
		tools.V19.String(): &datacapv11.MintReturn{},
		tools.V20.String(): &datacapv11.MintReturn{},
		tools.V21.String(): &datacapv12.MintReturn{},
		tools.V22.String(): &datacapv13.MintReturn{},
		tools.V23.String(): &datacapv14.MintReturn{},
		tools.V24.String(): &datacapv15.MintReturn{},
		tools.V25.String(): &datacapv16.MintReturn{},
	}
}

func (*Datacap) MintExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := mintParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := mintReturn()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parse(raw, rawReturn, true, params, returnValue, parser.ParamsKey)
}
