package system

import (
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
	"go.uber.org/zap"
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

func (s *System) Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, mainMsgCid cid.Cid, key filTypes.TipSetKey) (map[string]interface{}, *types.AddressInfo, error) {
	var resp map[string]interface{}
	var err error
	switch txType {
	case parser.MethodConstructor:
		resp, err = s.Constructor()
	default:
		resp, err = s.parseSystemAny(msg.Params, msgRct.Return)
	}

	return resp, nil, err
}

func (s *System) TransactionTypes() map[string]any {
	return map[string]any{
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
