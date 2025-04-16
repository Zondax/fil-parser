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

func burnParams() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V17.String(): &datacapv9.BurnParams{},
		tools.V18.String(): &datacapv10.BurnParams{},
		tools.V19.String(): &datacapv11.BurnParams{},
		tools.V20.String(): &datacapv11.BurnParams{},
		tools.V21.String(): &datacapv12.BurnParams{},
		tools.V22.String(): &datacapv13.BurnParams{},
		tools.V23.String(): &datacapv14.BurnParams{},
		tools.V24.String(): &datacapv15.BurnParams{},
		tools.V25.String(): &datacapv16.BurnParams{},
	}
}

func burnReturn() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V17.String(): &datacapv9.BurnReturn{},
		tools.V18.String(): &datacapv10.BurnReturn{},
		tools.V19.String(): &datacapv11.BurnReturn{},
		tools.V20.String(): &datacapv11.BurnReturn{},
		tools.V21.String(): &datacapv12.BurnReturn{},
		tools.V22.String(): &datacapv13.BurnReturn{},
		tools.V23.String(): &datacapv14.BurnReturn{},
		tools.V24.String(): &datacapv15.BurnReturn{},
		tools.V25.String(): &datacapv16.BurnReturn{},
	}
}

func burnFromParams() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V17.String(): &datacapv9.BurnFromParams{},
		tools.V18.String(): &datacapv10.BurnFromParams{},
		tools.V19.String(): &datacapv11.BurnFromParams{},
		tools.V20.String(): &datacapv11.BurnFromParams{},
		tools.V21.String(): &datacapv12.BurnFromParams{},
		tools.V22.String(): &datacapv13.BurnFromParams{},
		tools.V23.String(): &datacapv14.BurnFromParams{},
		tools.V24.String(): &datacapv15.BurnFromParams{},
		tools.V25.String(): &datacapv16.BurnFromParams{},
	}
}

func burnFromReturn() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V17.String(): &datacapv9.BurnFromReturn{},
		tools.V18.String(): &datacapv10.BurnFromReturn{},
		tools.V19.String(): &datacapv11.BurnFromReturn{},
		tools.V20.String(): &datacapv11.BurnFromReturn{},
		tools.V21.String(): &datacapv12.BurnFromReturn{},
		tools.V22.String(): &datacapv13.BurnFromReturn{},
		tools.V23.String(): &datacapv14.BurnFromReturn{},
		tools.V24.String(): &datacapv15.BurnFromReturn{},
		tools.V25.String(): &datacapv16.BurnFromReturn{},
	}
}

func destroyParams() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V17.String(): &datacapv9.DestroyParams{},
		tools.V18.String(): &datacapv10.DestroyParams{},
		tools.V19.String(): &datacapv11.DestroyParams{},
		tools.V20.String(): &datacapv11.DestroyParams{},
		tools.V21.String(): &datacapv12.DestroyParams{},
		tools.V22.String(): &datacapv13.DestroyParams{},
		tools.V23.String(): &datacapv14.DestroyParams{},
		tools.V24.String(): &datacapv15.DestroyParams{},
		tools.V25.String(): &datacapv16.DestroyParams{},
	}
}

func (*Datacap) BurnExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := burnParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := burnReturn()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parse(raw, rawReturn, true, params, returnValue, parser.ParamsKey)
}

func (*Datacap) BurnFromExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := burnFromParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := burnFromReturn()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parse(raw, rawReturn, true, params, returnValue, parser.ParamsKey)
}

func (*Datacap) DestroyExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := destroyParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := burnReturn()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parse(raw, rawReturn, true, params, returnValue, parser.ParamsKey)
}
