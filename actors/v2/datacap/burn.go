package datacap

import (
	"fmt"

	datacapv10 "github.com/filecoin-project/go-state-types/builtin/v10/datacap"
	datacapv11 "github.com/filecoin-project/go-state-types/builtin/v11/datacap"
	datacapv14 "github.com/filecoin-project/go-state-types/builtin/v14/datacap"
	datacapv15 "github.com/filecoin-project/go-state-types/builtin/v15/datacap"
	datacapv9 "github.com/filecoin-project/go-state-types/builtin/v9/datacap"
	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/tools"
)

func (*Datacap) BurnExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V16)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	case tools.V17.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv9.BurnParams{}, &datacapv9.BurnReturn{})
	case tools.V18.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv10.BurnParams{}, &datacapv10.BurnReturn{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parse(raw, rawReturn, true, &datacapv11.BurnParams{}, &datacapv11.BurnReturn{})
	case tools.V21.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv11.BurnParams{}, &datacapv11.BurnReturn{})
	case tools.V23.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv14.BurnParams{}, &datacapv14.BurnReturn{})
	case tools.V24.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv15.BurnParams{}, &datacapv15.BurnReturn{})
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Datacap) BurnFromExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V16)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	case tools.V17.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv9.BurnFromParams{}, &datacapv9.BurnFromReturn{})
	case tools.V18.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv10.BurnFromParams{}, &datacapv10.BurnFromReturn{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parse(raw, rawReturn, true, &datacapv11.BurnFromParams{}, &datacapv11.BurnFromReturn{})
	case tools.V21.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv11.BurnFromParams{}, &datacapv11.BurnFromReturn{})
	case tools.V23.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv14.BurnFromParams{}, &datacapv14.BurnFromReturn{})
	case tools.V24.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv15.BurnFromParams{}, &datacapv15.BurnFromReturn{})
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Datacap) DestroyExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V16)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	case tools.V17.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv9.DestroyParams{}, &datacapv9.BurnReturn{})
	case tools.V18.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv10.DestroyParams{}, &datacapv10.BurnReturn{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parse(raw, rawReturn, true, &datacapv11.DestroyParams{}, &datacapv11.BurnReturn{})
	case tools.V21.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv11.DestroyParams{}, &datacapv11.BurnReturn{})
	case tools.V23.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv14.DestroyParams{}, &datacapv14.BurnReturn{})
	case tools.V24.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv15.DestroyParams{}, &datacapv15.BurnReturn{})
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}
