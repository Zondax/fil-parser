package datacap

import (
	"fmt"

	"github.com/filecoin-project/go-state-types/abi"
	datacapv10 "github.com/filecoin-project/go-state-types/builtin/v10/datacap"
	datacapv11 "github.com/filecoin-project/go-state-types/builtin/v11/datacap"
	datacapv14 "github.com/filecoin-project/go-state-types/builtin/v14/datacap"
	datacapv15 "github.com/filecoin-project/go-state-types/builtin/v15/datacap"
	datacapv9 "github.com/filecoin-project/go-state-types/builtin/v9/datacap"
	"github.com/zondax/fil-parser/tools"
)

func (*Datacap) IncreaseAllowance(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(network, height):
		return nil, fmt.Errorf("not supported")
	case tools.V17.IsSupported(network, height):
		return parse[*datacapv9.IncreaseAllowanceParams, *abi.TokenAmount](raw, rawReturn, true)
	case tools.V18.IsSupported(network, height):
		return parse[*datacapv10.IncreaseAllowanceParams, *abi.TokenAmount](raw, rawReturn, true)
	case tools.V19.IsSupported(network, height) || tools.V20.IsSupported(network, height):
		return parse[*datacapv11.IncreaseAllowanceParams, *abi.TokenAmount](raw, rawReturn, true)
	case tools.V21.IsSupported(network, height):
		return parse[*datacapv11.IncreaseAllowanceParams, *abi.TokenAmount](raw, rawReturn, true)
	case tools.V23.IsSupported(network, height):
		return parse[*datacapv14.IncreaseAllowanceParams, *abi.TokenAmount](raw, rawReturn, true)
	case tools.V24.IsSupported(network, height):
		return parse[*datacapv15.IncreaseAllowanceParams, *abi.TokenAmount](raw, rawReturn, true)
	}
	return nil, fmt.Errorf("not supported")
}

func (*Datacap) DecreaseAllowance(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(network, height):
		return nil, fmt.Errorf("not supported")
	case tools.V17.IsSupported(network, height):
		return parse[*datacapv9.DecreaseAllowanceParams, *abi.TokenAmount](raw, rawReturn, true)
	case tools.V18.IsSupported(network, height):
		return parse[*datacapv10.DecreaseAllowanceParams, *abi.TokenAmount](raw, rawReturn, true)
	case tools.V19.IsSupported(network, height) || tools.V20.IsSupported(network, height):
		return parse[*datacapv11.DecreaseAllowanceParams, *abi.TokenAmount](raw, rawReturn, true)
	case tools.V21.IsSupported(network, height):
		return parse[*datacapv11.DecreaseAllowanceParams, *abi.TokenAmount](raw, rawReturn, true)
	case tools.V23.IsSupported(network, height):
		return parse[*datacapv14.DecreaseAllowanceParams, *abi.TokenAmount](raw, rawReturn, true)
	case tools.V24.IsSupported(network, height):
		return parse[*datacapv15.DecreaseAllowanceParams, *abi.TokenAmount](raw, rawReturn, true)
	}
	return nil, fmt.Errorf("not supported")
}

func (*Datacap) RevokeAllowance(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(network, height):
		return nil, fmt.Errorf("not supported")
	case tools.V17.IsSupported(network, height):
		return parse[*datacapv9.RevokeAllowanceParams, *abi.TokenAmount](raw, rawReturn, true)
	case tools.V18.IsSupported(network, height):
		return parse[*datacapv10.RevokeAllowanceParams, *abi.TokenAmount](raw, rawReturn, true)
	case tools.V19.IsSupported(network, height) || tools.V20.IsSupported(network, height):
		return parse[*datacapv11.RevokeAllowanceParams, *abi.TokenAmount](raw, rawReturn, true)
	case tools.V21.IsSupported(network, height):
		return parse[*datacapv11.RevokeAllowanceParams, *abi.TokenAmount](raw, rawReturn, true)
	case tools.V23.IsSupported(network, height):
		return parse[*datacapv14.RevokeAllowanceParams, *abi.TokenAmount](raw, rawReturn, true)
	case tools.V24.IsSupported(network, height):
		return parse[*datacapv15.RevokeAllowanceParams, *abi.TokenAmount](raw, rawReturn, true)
	}
	return nil, fmt.Errorf("not supported")
}

func (*Datacap) GetAllowance(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(network, height):
		return nil, fmt.Errorf("not supported")
	case tools.V17.IsSupported(network, height):
		return parse[*datacapv9.GetAllowanceParams, *abi.TokenAmount](raw, rawReturn, true)
	case tools.V18.IsSupported(network, height):
		return parse[*datacapv10.GetAllowanceParams, *abi.TokenAmount](raw, rawReturn, true)
	case tools.V19.IsSupported(network, height) || tools.V20.IsSupported(network, height):
		return parse[*datacapv11.GetAllowanceParams, *abi.TokenAmount](raw, rawReturn, true)
	case tools.V21.IsSupported(network, height):
		return parse[*datacapv11.GetAllowanceParams, *abi.TokenAmount](raw, rawReturn, true)
	case tools.V23.IsSupported(network, height):
		return parse[*datacapv14.GetAllowanceParams, *abi.TokenAmount](raw, rawReturn, true)
	case tools.V24.IsSupported(network, height):
		return parse[*datacapv15.GetAllowanceParams, *abi.TokenAmount](raw, rawReturn, true)
	}
	return nil, fmt.Errorf("not supported")
}
