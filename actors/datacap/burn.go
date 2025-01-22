package datacap

import (
	"fmt"

	datacapv10 "github.com/filecoin-project/go-state-types/builtin/v10/datacap"
	datacapv11 "github.com/filecoin-project/go-state-types/builtin/v11/datacap"
	datacapv12 "github.com/filecoin-project/go-state-types/builtin/v12/datacap"
	datacapv13 "github.com/filecoin-project/go-state-types/builtin/v13/datacap"
	datacapv14 "github.com/filecoin-project/go-state-types/builtin/v14/datacap"
	datacapv15 "github.com/filecoin-project/go-state-types/builtin/v15/datacap"
	datacapv9 "github.com/filecoin-project/go-state-types/builtin/v9/datacap"
	"github.com/zondax/fil-parser/tools"
)

func BurnExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(network, height):
		return nil, fmt.Errorf("not supported")
	case tools.V9.IsSupported(network, height):
		return parse[*datacapv9.BurnParams, *datacapv9.BurnReturn](raw, rawReturn, true)
	case tools.V10.IsSupported(network, height):
		return parse[*datacapv10.BurnParams, *datacapv10.BurnReturn](raw, rawReturn, true)
	case tools.V11.IsSupported(network, height):
		return parse[*datacapv11.BurnParams, *datacapv11.BurnReturn](raw, rawReturn, true)
	case tools.V12.IsSupported(network, height):
		return parse[*datacapv12.BurnParams, *datacapv12.BurnReturn](raw, rawReturn, true)
	case tools.V13.IsSupported(network, height):
		return parse[*datacapv13.BurnParams, *datacapv13.BurnReturn](raw, rawReturn, true)
	case tools.V14.IsSupported(network, height):
		return parse[*datacapv14.BurnParams, *datacapv14.BurnReturn](raw, rawReturn, true)
	case tools.V15.IsSupported(network, height):
		return parse[*datacapv15.BurnParams, *datacapv15.BurnReturn](raw, rawReturn, true)
	}
	return nil, fmt.Errorf("not supported")
}

func BurnFromExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(network, height):
		return nil, fmt.Errorf("not supported")
	case tools.V9.IsSupported(network, height):
		return parse[*datacapv9.BurnFromParams, *datacapv9.BurnFromReturn](raw, rawReturn, true)
	case tools.V10.IsSupported(network, height):
		return parse[*datacapv10.BurnFromParams, *datacapv10.BurnFromReturn](raw, rawReturn, true)
	case tools.V11.IsSupported(network, height):
		return parse[*datacapv11.BurnFromParams, *datacapv11.BurnFromReturn](raw, rawReturn, true)
	case tools.V12.IsSupported(network, height):
		return parse[*datacapv12.BurnFromParams, *datacapv12.BurnFromReturn](raw, rawReturn, true)
	case tools.V13.IsSupported(network, height):
		return parse[*datacapv13.BurnFromParams, *datacapv13.BurnFromReturn](raw, rawReturn, true)
	case tools.V14.IsSupported(network, height):
		return parse[*datacapv14.BurnFromParams, *datacapv14.BurnFromReturn](raw, rawReturn, true)
	case tools.V15.IsSupported(network, height):
		return parse[*datacapv15.BurnFromParams, *datacapv15.BurnFromReturn](raw, rawReturn, true)
	}
	return nil, fmt.Errorf("not supported")
}

func DestroyExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(network, height):
		return nil, fmt.Errorf("not supported")
	case tools.V9.IsSupported(network, height):
		return parse[*datacapv9.DestroyParams, *datacapv9.BurnReturn](raw, rawReturn, true)
	case tools.V10.IsSupported(network, height):
		return parse[*datacapv10.DestroyParams, *datacapv10.BurnReturn](raw, rawReturn, true)
	case tools.V11.IsSupported(network, height):
		return parse[*datacapv11.DestroyParams, *datacapv11.BurnReturn](raw, rawReturn, true)
	case tools.V12.IsSupported(network, height):
		return parse[*datacapv12.DestroyParams, *datacapv12.BurnReturn](raw, rawReturn, true)
	case tools.V13.IsSupported(network, height):
		return parse[*datacapv13.DestroyParams, *datacapv13.BurnReturn](raw, rawReturn, true)
	case tools.V14.IsSupported(network, height):
		return parse[*datacapv14.DestroyParams, *datacapv14.BurnReturn](raw, rawReturn, true)
	case tools.V15.IsSupported(network, height):
		return parse[*datacapv15.DestroyParams, *datacapv15.BurnReturn](raw, rawReturn, true)
	}
	return nil, fmt.Errorf("not supported")
}
