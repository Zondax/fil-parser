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

type (
	mintParams = unmarshalCBOR
	mintReturn = unmarshalCBOR
)

func (*Datacap) MintExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(network, height):
		return nil, fmt.Errorf("not supported")
	case tools.V17.IsSupported(network, height):
		return parse[*datacapv9.MintParams, *datacapv9.MintReturn](raw, rawReturn, true)
	case tools.V18.IsSupported(network, height):
		return parse[*datacapv10.MintParams, *datacapv10.MintReturn](raw, rawReturn, true)
	case tools.V19.IsSupported(network, height) || tools.V20.IsSupported(network, height):
		return parse[*datacapv11.MintParams, *datacapv11.MintReturn](raw, rawReturn, true)
	case tools.V12.IsSupported(network, height):
		return parse[*datacapv12.MintParams, *datacapv12.MintReturn](raw, rawReturn, true)
	case tools.V21.IsSupported(network, height):
		return parse[*datacapv12.MintParams, *datacapv12.MintReturn](raw, rawReturn, true)
	case tools.V23.IsSupported(network, height):
		return parse[*datacapv14.MintParams, *datacapv14.MintReturn](raw, rawReturn, true)
	case tools.V24.IsSupported(network, height):
		return parse[*datacapv15.MintParams, *datacapv15.MintReturn](raw, rawReturn, true)
	}
	return nil, fmt.Errorf("not supported")
}
