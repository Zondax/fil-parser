package datacap

import (
	"fmt"

	datacapv10 "github.com/filecoin-project/go-state-types/builtin/v10/datacap"
	datacapv11 "github.com/filecoin-project/go-state-types/builtin/v11/datacap"
	datacapv12 "github.com/filecoin-project/go-state-types/builtin/v12/datacap"
	datacapv14 "github.com/filecoin-project/go-state-types/builtin/v14/datacap"
	datacapv15 "github.com/filecoin-project/go-state-types/builtin/v15/datacap"
	datacapv9 "github.com/filecoin-project/go-state-types/builtin/v9/datacap"
	"github.com/zondax/fil-parser/tools"
)

func (*Datacap) TransferExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(network, height):
		return nil, fmt.Errorf("not supported")
	case tools.V17.IsSupported(network, height):
		return parse[*datacapv9.TransferParams, *datacapv9.TransferReturn](raw, rawReturn, true)
	case tools.V18.IsSupported(network, height):
		return parse[*datacapv10.TransferParams, *datacapv10.TransferReturn](raw, rawReturn, true)
	case tools.V19.IsSupported(network, height) || tools.V20.IsSupported(network, height):
		return parse[*datacapv11.TransferParams, *datacapv11.TransferReturn](raw, rawReturn, true)
	case tools.V12.IsSupported(network, height):
		return parse[*datacapv12.TransferParams, *datacapv12.TransferReturn](raw, rawReturn, true)
	case tools.V21.IsSupported(network, height):
		return parse[*datacapv12.TransferParams, *datacapv12.TransferReturn](raw, rawReturn, true)
	case tools.V23.IsSupported(network, height):
		return parse[*datacapv14.TransferParams, *datacapv14.TransferReturn](raw, rawReturn, true)
	case tools.V24.IsSupported(network, height):
		return parse[*datacapv15.TransferParams, *datacapv15.TransferReturn](raw, rawReturn, true)
	}
	return nil, fmt.Errorf("not supported")
}
