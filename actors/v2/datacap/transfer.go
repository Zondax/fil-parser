package datacap

import (
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	datacapv10 "github.com/filecoin-project/go-state-types/builtin/v10/datacap"
	datacapv11 "github.com/filecoin-project/go-state-types/builtin/v11/datacap"
	datacapv12 "github.com/filecoin-project/go-state-types/builtin/v12/datacap"
	datacapv13 "github.com/filecoin-project/go-state-types/builtin/v13/datacap"
	datacapv14 "github.com/filecoin-project/go-state-types/builtin/v14/datacap"
	datacapv15 "github.com/filecoin-project/go-state-types/builtin/v15/datacap"
	datacapv9 "github.com/filecoin-project/go-state-types/builtin/v9/datacap"
	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/tools"
)

func (*Datacap) TransferExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V17)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	case tools.V17.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv9.TransferParams{}, &datacapv9.TransferReturn{})
	case tools.V18.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv10.TransferParams{}, &datacapv10.TransferReturn{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parse(raw, rawReturn, true, &datacapv11.TransferParams{}, &datacapv11.TransferReturn{})
	case tools.V12.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv12.TransferParams{}, &datacapv12.TransferReturn{})
	case tools.V21.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv12.TransferParams{}, &datacapv12.TransferReturn{})
	case tools.V22.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv13.TransferParams{}, &datacapv13.TransferReturn{})
	case tools.V23.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv14.TransferParams{}, &datacapv14.TransferReturn{})
	case tools.V24.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv15.TransferParams{}, &datacapv15.TransferReturn{})
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Datacap) BalanceExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	return parse(raw, rawReturn, true, &address.Address{}, &abi.TokenAmount{})
}

func (*Datacap) TransferFromExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V17)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	case tools.V17.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv9.TransferFromParams{}, &datacapv9.TransferFromReturn{})
	case tools.V18.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv10.TransferFromParams{}, &datacapv10.TransferFromReturn{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parse(raw, rawReturn, true, &datacapv11.TransferFromParams{}, &datacapv11.TransferFromReturn{})
	case tools.V12.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv12.TransferFromParams{}, &datacapv12.TransferFromReturn{})
	case tools.V21.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv12.TransferFromParams{}, &datacapv12.TransferFromReturn{})
	case tools.V22.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv13.TransferFromParams{}, &datacapv13.TransferFromReturn{})
	case tools.V23.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv14.TransferFromParams{}, &datacapv14.TransferFromReturn{})
	case tools.V24.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv15.TransferFromParams{}, &datacapv15.TransferFromReturn{})
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}
