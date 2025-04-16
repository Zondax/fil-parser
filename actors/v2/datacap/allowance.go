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
	datacapv9 "github.com/filecoin-project/go-state-types/builtin/v9/datacap"

	typegen "github.com/whyrusleeping/cbor-gen"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

func increaseAllowanceParams() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V17.String(): &datacapv9.IncreaseAllowanceParams{},
		tools.V18.String(): &datacapv10.IncreaseAllowanceParams{},
		tools.V19.String(): &datacapv11.IncreaseAllowanceParams{},
		tools.V20.String(): &datacapv11.IncreaseAllowanceParams{},
		tools.V21.String(): &datacapv12.IncreaseAllowanceParams{},
		tools.V22.String(): &datacapv13.IncreaseAllowanceParams{},
		tools.V23.String(): &datacapv14.IncreaseAllowanceParams{},
		tools.V24.String(): &datacapv15.IncreaseAllowanceParams{},
		tools.V25.String(): &datacapv16.IncreaseAllowanceParams{},
	}
}

func decreaseAllowanceParams() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V17.String(): &datacapv9.DecreaseAllowanceParams{},
		tools.V18.String(): &datacapv10.DecreaseAllowanceParams{},
		tools.V19.String(): &datacapv11.DecreaseAllowanceParams{},
		tools.V20.String(): &datacapv11.DecreaseAllowanceParams{},
		tools.V21.String(): &datacapv12.DecreaseAllowanceParams{},
		tools.V22.String(): &datacapv13.DecreaseAllowanceParams{},
		tools.V23.String(): &datacapv14.DecreaseAllowanceParams{},
		tools.V24.String(): &datacapv15.DecreaseAllowanceParams{},
		tools.V25.String(): &datacapv16.DecreaseAllowanceParams{},
	}
}

func revokeAllowanceParams() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V17.String(): &datacapv9.RevokeAllowanceParams{},
		tools.V18.String(): &datacapv10.RevokeAllowanceParams{},
		tools.V19.String(): &datacapv11.RevokeAllowanceParams{},
		tools.V20.String(): &datacapv11.RevokeAllowanceParams{},
		tools.V21.String(): &datacapv12.RevokeAllowanceParams{},
		tools.V22.String(): &datacapv13.RevokeAllowanceParams{},
		tools.V23.String(): &datacapv14.RevokeAllowanceParams{},
		tools.V24.String(): &datacapv15.RevokeAllowanceParams{},
		tools.V25.String(): &datacapv16.RevokeAllowanceParams{},
	}
}

func allowanceParams() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V17.String(): &datacapv9.GetAllowanceParams{},
		tools.V18.String(): &datacapv10.GetAllowanceParams{},
		tools.V19.String(): &datacapv11.GetAllowanceParams{},
		tools.V20.String(): &datacapv11.GetAllowanceParams{},
		tools.V21.String(): &datacapv12.GetAllowanceParams{},
		tools.V22.String(): &datacapv13.GetAllowanceParams{},
		tools.V23.String(): &datacapv14.GetAllowanceParams{},
		tools.V24.String(): &datacapv15.GetAllowanceParams{},
		tools.V25.String(): &datacapv16.GetAllowanceParams{},
	}
}

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
