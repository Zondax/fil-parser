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
	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

func (*Datacap) MintExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V16)...):
		return nil, actors.ErrInvalidHeightForMethod
	case tools.V17.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv9.MintParams{}, &datacapv9.MintReturn{}, parser.ParamsKey)
	case tools.V18.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv10.MintParams{}, &datacapv10.MintReturn{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parse(raw, rawReturn, true, &datacapv11.MintParams{}, &datacapv11.MintReturn{}, parser.ParamsKey)
	case tools.V12.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv12.MintParams{}, &datacapv12.MintReturn{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv12.MintParams{}, &datacapv12.MintReturn{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv13.MintParams{}, &datacapv13.MintReturn{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv14.MintParams{}, &datacapv14.MintReturn{}, parser.ParamsKey)
	case tools.V24.IsSupported(network, height):
		return parse(raw, rawReturn, true, &datacapv15.MintParams{}, &datacapv15.MintReturn{}, parser.ParamsKey)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}
