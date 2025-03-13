package system

import (
	"context"
	"fmt"

	"github.com/ipfs/go-cid"
	"go.uber.org/zap"

	"github.com/filecoin-project/go-state-types/abi"
	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"

	systemv10 "github.com/filecoin-project/go-state-types/builtin/v10/system"
	systemv11 "github.com/filecoin-project/go-state-types/builtin/v11/system"
	systemv12 "github.com/filecoin-project/go-state-types/builtin/v12/system"
	systemv13 "github.com/filecoin-project/go-state-types/builtin/v13/system"
	systemv14 "github.com/filecoin-project/go-state-types/builtin/v14/system"
	systemv15 "github.com/filecoin-project/go-state-types/builtin/v15/system"
	systemv8 "github.com/filecoin-project/go-state-types/builtin/v8/system"
	systemv9 "github.com/filecoin-project/go-state-types/builtin/v9/system"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
)

type System struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) *System {
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

func (*System) Methods(_ context.Context, network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error) {
	switch {
	// all legacy version
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V15)...):
		return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
			abi.MethodNum(0): {
				Name: parser.MethodConstructor,
			},
		}, nil
	case tools.V16.IsSupported(network, height):
		return systemv8.Methods, nil
	case tools.V17.IsSupported(network, height):
		return systemv9.Methods, nil
	case tools.V18.IsSupported(network, height):
		return systemv10.Methods, nil
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return systemv11.Methods, nil
	case tools.V21.IsSupported(network, height):
		return systemv12.Methods, nil
	case tools.V22.IsSupported(network, height):
		return systemv13.Methods, nil
	case tools.V23.IsSupported(network, height):
		return systemv14.Methods, nil
	case tools.V24.IsSupported(network, height):
		return systemv15.Methods, nil
	default:
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
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
