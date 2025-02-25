package datacap

import (
	"fmt"

	datacapv10 "github.com/filecoin-project/go-state-types/builtin/v10/datacap"
	datacapv11 "github.com/filecoin-project/go-state-types/builtin/v11/datacap"
	datacapv12 "github.com/filecoin-project/go-state-types/builtin/v12/datacap"
	datacapv13 "github.com/filecoin-project/go-state-types/builtin/v13/datacap"
	datacapv14 "github.com/filecoin-project/go-state-types/builtin/v14/datacap"
	datacapv15 "github.com/filecoin-project/go-state-types/builtin/v15/datacap"
	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

func (*Datacap) GranularityExported(network string, height int64, rawReturn []byte) (map[string]interface{}, error) {
	var data map[string]interface{}
	var err error
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V17)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	case tools.V18.IsSupported(network, height):
		var params datacapv10.GranularityReturn
		var r datacapv10.GranularityReturn
		data, err = parse(rawReturn, nil, false, &params, &r, parser.ReturnKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		var params datacapv11.GranularityReturn
		var r datacapv11.GranularityReturn
		data, err = parse(rawReturn, nil, false, &params, &r, parser.ReturnKey)
	case tools.V21.IsSupported(network, height):
		var params datacapv12.GranularityReturn
		var r datacapv12.GranularityReturn
		data, err = parse(rawReturn, nil, false, &params, &r, parser.ReturnKey)
	case tools.V22.IsSupported(network, height):
		var params datacapv13.GranularityReturn
		var r datacapv13.GranularityReturn
		data, err = parse(rawReturn, nil, false, &params, &r, parser.ReturnKey)
	case tools.V23.IsSupported(network, height):
		var params datacapv14.GranularityReturn
		var r datacapv14.GranularityReturn
		data, err = parse(rawReturn, nil, false, &params, &r, parser.ReturnKey)
	case tools.V24.IsSupported(network, height):
		var params datacapv15.GranularityReturn
		var r datacapv15.GranularityReturn
		data, err = parse(rawReturn, nil, false, &params, &r, parser.ReturnKey)
	default:
		err = actors.ErrUnsupportedHeight
	}
	if err != nil {
		return nil, err
	}

	return data, nil
}
