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
	datacapv16 "github.com/filecoin-project/go-state-types/builtin/v16/datacap"
	datacapv9 "github.com/filecoin-project/go-state-types/builtin/v9/datacap"

	typegen "github.com/whyrusleeping/cbor-gen"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

func transferParams() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V17.String(): &datacapv9.TransferParams{},
		tools.V18.String(): &datacapv10.TransferParams{},
		tools.V19.String(): &datacapv11.TransferParams{},
		tools.V20.String(): &datacapv11.TransferParams{},
		tools.V21.String(): &datacapv12.TransferParams{},
		tools.V22.String(): &datacapv13.TransferParams{},
		tools.V23.String(): &datacapv14.TransferParams{},
		tools.V24.String(): &datacapv15.TransferParams{},
		tools.V25.String(): &datacapv16.TransferParams{},
	}
}

func transferReturn() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V17.String(): &datacapv9.TransferReturn{},
		tools.V18.String(): &datacapv10.TransferReturn{},
		tools.V19.String(): &datacapv11.TransferReturn{},
		tools.V20.String(): &datacapv11.TransferReturn{},
		tools.V21.String(): &datacapv12.TransferReturn{},
		tools.V22.String(): &datacapv13.TransferReturn{},
		tools.V23.String(): &datacapv14.TransferReturn{},
		tools.V24.String(): &datacapv15.TransferReturn{},
		tools.V25.String(): &datacapv16.TransferReturn{},
	}
}

func transferFromParams() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V17.String(): &datacapv9.TransferFromParams{},
		tools.V18.String(): &datacapv10.TransferFromParams{},
		tools.V19.String(): &datacapv11.TransferFromParams{},
		tools.V20.String(): &datacapv11.TransferFromParams{},
		tools.V21.String(): &datacapv12.TransferFromParams{},
		tools.V22.String(): &datacapv13.TransferFromParams{},
		tools.V23.String(): &datacapv14.TransferFromParams{},
		tools.V24.String(): &datacapv15.TransferFromParams{},
		tools.V25.String(): &datacapv16.TransferFromParams{},
	}
}

func transferFromReturn() map[string]typegen.CBORUnmarshaler {
	return map[string]typegen.CBORUnmarshaler{
		tools.V17.String(): &datacapv9.TransferFromReturn{},
		tools.V18.String(): &datacapv10.TransferFromReturn{},
		tools.V19.String(): &datacapv11.TransferFromReturn{},
		tools.V20.String(): &datacapv11.TransferFromReturn{},
		tools.V21.String(): &datacapv12.TransferFromReturn{},
		tools.V22.String(): &datacapv13.TransferFromReturn{},
		tools.V23.String(): &datacapv14.TransferFromReturn{},
		tools.V24.String(): &datacapv15.TransferFromReturn{},
		tools.V25.String(): &datacapv16.TransferFromReturn{},
	}
}

func (*Datacap) TransferExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := transferParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := transferReturn()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return parse(raw, rawReturn, true, params, returnValue, parser.ParamsKey)
}

func (*Datacap) BalanceExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	return parse(raw, rawReturn, true, &address.Address{}, &abi.TokenAmount{}, parser.ParamsKey)
}

func (*Datacap) TransferFromExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := transferFromParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := transferFromReturn()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return parse(raw, rawReturn, true, params, returnValue, parser.ParamsKey)
}

func (*Datacap) BalanceOf(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	return parse(raw, rawReturn, true, &address.Address{}, &abi.TokenAmount{}, parser.ParamsKey)
}
