package init

import (
	"fmt"

	builtinInitv10 "github.com/filecoin-project/go-state-types/builtin/v10/init"
	builtinInitv11 "github.com/filecoin-project/go-state-types/builtin/v11/init"
	builtinInitv12 "github.com/filecoin-project/go-state-types/builtin/v12/init"
	builtinInitv13 "github.com/filecoin-project/go-state-types/builtin/v13/init"
	builtinInitv14 "github.com/filecoin-project/go-state-types/builtin/v14/init"
	builtinInitv15 "github.com/filecoin-project/go-state-types/builtin/v15/init"
	builtinInitv8 "github.com/filecoin-project/go-state-types/builtin/v8/init"
	builtinInitv9 "github.com/filecoin-project/go-state-types/builtin/v9/init"

	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
)

type Init struct{}

func (*Init) InitConstructor(network string, height int64, raw []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return initConstructor[*builtinInitv15.ConstructorParams](raw)
	case tools.V23.IsSupported(network, height):
		return initConstructor[*builtinInitv14.ConstructorParams](raw)
	case tools.V22.IsSupported(network, height):
		return initConstructor[*builtinInitv13.ConstructorParams](raw)
	case tools.V21.IsSupported(network, height):
		return initConstructor[*builtinInitv12.ConstructorParams](raw)
	case tools.V20.IsSupported(network, height):
		return initConstructor[*builtinInitv11.ConstructorParams](raw)
	case tools.V18.IsSupported(network, height):
		return initConstructor[*builtinInitv10.ConstructorParams](raw)
	case tools.V17.IsSupported(network, height):
		return initConstructor[*builtinInitv9.ConstructorParams](raw)
	case tools.V16.IsSupported(network, height):
		return initConstructor[*builtinInitv8.ConstructorParams](raw)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func (*Init) ParseExec(network string, height int64, msg *parser.LotusMessage, raw []byte) (map[string]interface{}, *types.AddressInfo, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseExec[*builtinInitv15.ExecParams, *builtinInitv15.ExecReturn](msg, raw)
	case tools.V23.IsSupported(network, height):
		return parseExec[*builtinInitv14.ExecParams, *builtinInitv14.ExecReturn](msg, raw)
	case tools.V22.IsSupported(network, height):
		return parseExec[*builtinInitv13.ExecParams, *builtinInitv13.ExecReturn](msg, raw)
	case tools.V21.IsSupported(network, height):
		return parseExec[*builtinInitv12.ExecParams, *builtinInitv12.ExecReturn](msg, raw)
	case tools.V20.IsSupported(network, height):
		return parseExec[*builtinInitv11.ExecParams, *builtinInitv11.ExecReturn](msg, raw)
	case tools.V18.IsSupported(network, height):
		return parseExec[*builtinInitv10.ExecParams, *builtinInitv10.ExecReturn](msg, raw)
	case tools.V17.IsSupported(network, height):
		return parseExec[*builtinInitv9.ExecParams, *builtinInitv9.ExecReturn](msg, raw)
	case tools.V16.IsSupported(network, height):
		return parseExec[*builtinInitv8.ExecParams, *builtinInitv8.ExecReturn](msg, raw)
	}
	return nil, nil, fmt.Errorf("unsupported height: %d", height)
}

func (*Init) ParseExec4(network string, height int64, msg *parser.LotusMessage, raw []byte) (map[string]interface{}, *types.AddressInfo, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseExec[*builtinInitv15.Exec4Params, *builtinInitv15.Exec4Return](msg, raw)
	case tools.V23.IsSupported(network, height):
		return parseExec[*builtinInitv14.Exec4Params, *builtinInitv14.Exec4Return](msg, raw)
	case tools.V22.IsSupported(network, height):
		return parseExec[*builtinInitv13.Exec4Params, *builtinInitv13.Exec4Return](msg, raw)
	case tools.V21.IsSupported(network, height):
		return parseExec[*builtinInitv12.Exec4Params, *builtinInitv12.Exec4Return](msg, raw)
	case tools.V20.IsSupported(network, height):
		return parseExec[*builtinInitv11.Exec4Params, *builtinInitv11.Exec4Return](msg, raw)
	case tools.V18.IsSupported(network, height):
		return parseExec[*builtinInitv10.Exec4Params, *builtinInitv10.Exec4Return](msg, raw)
	case tools.V17.IsSupported(network, height):
		return nil, nil, fmt.Errorf("unsupported height: %d", height)
	case tools.V16.IsSupported(network, height):
		return nil, nil, fmt.Errorf("unsupported height: %d", height)
	}
	return nil, nil, fmt.Errorf("unsupported height: %d", height)
}
