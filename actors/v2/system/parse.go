package system

import (
	"context"
	"fmt"

	"github.com/zondax/golem/pkg/logger"

	"github.com/filecoin-project/go-state-types/abi"
	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"

	systemv10 "github.com/filecoin-project/go-state-types/builtin/v10/system"
	systemv11 "github.com/filecoin-project/go-state-types/builtin/v11/system"
	systemv12 "github.com/filecoin-project/go-state-types/builtin/v12/system"
	systemv13 "github.com/filecoin-project/go-state-types/builtin/v13/system"
	systemv14 "github.com/filecoin-project/go-state-types/builtin/v14/system"
	systemv15 "github.com/filecoin-project/go-state-types/builtin/v15/system"
	systemv16 "github.com/filecoin-project/go-state-types/builtin/v16/system"
	systemv8 "github.com/filecoin-project/go-state-types/builtin/v8/system"
	systemv9 "github.com/filecoin-project/go-state-types/builtin/v9/system"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
)

type System struct {
	logger *logger.Logger
}

func New(logger *logger.Logger) *System {
	return &System{
		logger: logger,
	}
}
func (s *System) Name() string {
	return manifest.SystemKey
}

func (*System) StartNetworkHeight() int64 {
	return tools.V1.Height()
}

func legacyMethods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
		abi.MethodNum(0): {
			Name:   parser.MethodConstructor,
			Method: actors.ParseConstructor,
		},
	}
}

var methods = map[string]map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
	tools.V1.String():  legacyMethods(),
	tools.V2.String():  legacyMethods(),
	tools.V3.String():  legacyMethods(),
	tools.V4.String():  legacyMethods(),
	tools.V5.String():  legacyMethods(),
	tools.V6.String():  legacyMethods(),
	tools.V7.String():  legacyMethods(),
	tools.V8.String():  legacyMethods(),
	tools.V9.String():  legacyMethods(),
	tools.V10.String(): legacyMethods(),
	tools.V11.String(): legacyMethods(),
	tools.V12.String(): legacyMethods(),
	tools.V13.String(): legacyMethods(),
	tools.V14.String(): legacyMethods(),
	tools.V15.String(): legacyMethods(),
	tools.V16.String(): actors.CopyMethods(systemv8.Methods),
	tools.V17.String(): actors.CopyMethods(systemv9.Methods),
	tools.V18.String(): actors.CopyMethods(systemv10.Methods),
	tools.V19.String(): actors.CopyMethods(systemv11.Methods),
	tools.V20.String(): actors.CopyMethods(systemv11.Methods),
	tools.V21.String(): actors.CopyMethods(systemv12.Methods),
	tools.V22.String(): actors.CopyMethods(systemv13.Methods),
	tools.V23.String(): actors.CopyMethods(systemv14.Methods),
	tools.V24.String(): actors.CopyMethods(systemv15.Methods),
	tools.V25.String(): actors.CopyMethods(systemv16.Methods),
}

func (*System) Methods(_ context.Context, network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error) {
	version := tools.VersionFromHeight(network, height)
	methods, ok := methods[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return methods, nil
}

func (s *System) Parse(_ context.Context, network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, mainMsgCid cid.Cid, key filTypes.TipSetKey) (map[string]interface{}, *types.AddressInfo, error) {
	var resp map[string]interface{}
	var err error
	switch txType {
	case parser.MethodSend:
		resp := actors.ParseSend(msg)
		return resp, nil, nil
	case parser.MethodConstructor:
		resp, err = s.Constructor()
	default:
		resp, err = s.parseSystemAny(msg.Params, msgRct.Return)
	}

	return resp, nil, err
}

func (s *System) TransactionTypes() map[string]any {
	return map[string]any{
		parser.MethodSend:        actors.ParseSend,
		parser.MethodConstructor: s.Constructor,
	}
}

func (s *System) Constructor() (map[string]interface{}, error) {
	return s.parseSystemAny(nil, nil)
}

func (s *System) parseSystemAny(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	metadata[parser.ParamsKey] = rawParams
	metadata[parser.ReturnKey] = rawReturn

	return metadata, nil
}
