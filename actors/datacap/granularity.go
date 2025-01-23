package datacap

import (
	"fmt"

	datacapv10 "github.com/filecoin-project/go-state-types/builtin/v10/datacap"
	datacapv11 "github.com/filecoin-project/go-state-types/builtin/v11/datacap"
	datacapv12 "github.com/filecoin-project/go-state-types/builtin/v12/datacap"
	datacapv14 "github.com/filecoin-project/go-state-types/builtin/v14/datacap"
	datacapv15 "github.com/filecoin-project/go-state-types/builtin/v15/datacap"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

func (*Datacap) GranularityExported(network string, height int64, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V8.IsSupported(network, height):
		return nil, fmt.Errorf("not supported")
	case tools.V18.IsSupported(network, height):
		data, err := parse[*datacapv10.GranularityReturn, *datacapv10.GranularityReturn](rawReturn, nil, false)
		if err != nil {
			return nil, err
		}
		data[parser.ReturnKey] = data[parser.ParamsKey]
		return data, nil
	case tools.V19.IsSupported(network, height) || tools.V20.IsSupported(network, height):
		data, err := parse[*datacapv11.GranularityReturn, *datacapv11.GranularityReturn](rawReturn, nil, false)
		if err != nil {
			return nil, err
		}
		data[parser.ReturnKey] = data[parser.ParamsKey]
		return data, nil
	case tools.V21.IsSupported(network, height):
		data, err := parse[*datacapv12.GranularityReturn, *datacapv12.GranularityReturn](rawReturn, nil, false)
		if err != nil {
			return nil, err
		}
		data[parser.ReturnKey] = data[parser.ParamsKey]
		return data, nil
	case tools.V23.IsSupported(network, height):
		data, err := parse[*datacapv14.GranularityReturn, *datacapv14.GranularityReturn](rawReturn, nil, false)
		if err != nil {
			return nil, err
		}
		data[parser.ReturnKey] = data[parser.ParamsKey]
		return data, nil
	case tools.V24.IsSupported(network, height):
		data, err := parse[*datacapv15.GranularityReturn, *datacapv15.GranularityReturn](rawReturn, nil, false)
		if err != nil {
			return nil, err
		}
		data[parser.ReturnKey] = data[parser.ParamsKey]
		return data, nil
	}
	return nil, fmt.Errorf("not supported")
}
