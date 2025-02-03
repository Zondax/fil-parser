package datacap

import (
	"fmt"

	"github.com/filecoin-project/go-state-types/abi"
	datacapv10 "github.com/filecoin-project/go-state-types/builtin/v10/datacap"
	datacapv11 "github.com/filecoin-project/go-state-types/builtin/v11/datacap"
	datacapv14 "github.com/filecoin-project/go-state-types/builtin/v14/datacap"
	datacapv15 "github.com/filecoin-project/go-state-types/builtin/v15/datacap"
	datacapv9 "github.com/filecoin-project/go-state-types/builtin/v9/datacap"
	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/tools"
)

func (*Datacap) IncreaseAllowanceExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	var r abi.TokenAmount
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V16)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	case tools.V17.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv9.IncreaseAllowanceParams{}, &r)
	case tools.V18.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv10.IncreaseAllowanceParams{}, &r)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parse(raw, rawReturn, true, &datacapv11.IncreaseAllowanceParams{}, &r)
	case tools.V21.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv11.IncreaseAllowanceParams{}, &r)
	case tools.V23.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv14.IncreaseAllowanceParams{}, &r)
	case tools.V24.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv15.IncreaseAllowanceParams{}, &r)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Datacap) DecreaseAllowanceExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	var r abi.TokenAmount
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V16)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	case tools.V17.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv9.DecreaseAllowanceParams{}, &r)
	case tools.V18.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv10.DecreaseAllowanceParams{}, &r)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parse(raw, rawReturn, true, &datacapv11.DecreaseAllowanceParams{}, &r)
	case tools.V21.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv11.DecreaseAllowanceParams{}, &r)
	case tools.V23.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv14.DecreaseAllowanceParams{}, &r)
	case tools.V24.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv15.DecreaseAllowanceParams{}, &r)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Datacap) RevokeAllowanceExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	var r abi.TokenAmount
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V16)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	case tools.V17.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv9.RevokeAllowanceParams{}, &r)
	case tools.V18.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv10.RevokeAllowanceParams{}, &r)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parse(raw, rawReturn, true, &datacapv11.RevokeAllowanceParams{}, &r)
	case tools.V21.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv11.RevokeAllowanceParams{}, &r)
	case tools.V23.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv14.RevokeAllowanceParams{}, &r)
	case tools.V24.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv15.RevokeAllowanceParams{}, &r)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Datacap) AllowanceExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	var r abi.TokenAmount
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V16)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	case tools.V17.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv9.GetAllowanceParams{}, &r)
	case tools.V18.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv10.GetAllowanceParams{}, &r)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parse(raw, rawReturn, true, &datacapv11.GetAllowanceParams{}, &r)
	case tools.V21.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv11.GetAllowanceParams{}, &r)
	case tools.V23.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv14.GetAllowanceParams{}, &r)
	case tools.V24.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv15.GetAllowanceParams{}, &r)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}
