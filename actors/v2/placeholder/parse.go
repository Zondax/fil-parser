package placeholder

import (
	"context"
	"fmt"
	"github.com/zondax/golem/pkg/logger"

	"github.com/filecoin-project/go-state-types/abi"
	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"

	placeholderv10 "github.com/filecoin-project/go-state-types/builtin/v10/placeholder"
	placeholderv11 "github.com/filecoin-project/go-state-types/builtin/v11/placeholder"
	placeholderv12 "github.com/filecoin-project/go-state-types/builtin/v12/placeholder"
	placeholderv13 "github.com/filecoin-project/go-state-types/builtin/v13/placeholder"
	placeholderv14 "github.com/filecoin-project/go-state-types/builtin/v14/placeholder"
	placeholderv15 "github.com/filecoin-project/go-state-types/builtin/v15/placeholder"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
)

type Placeholder struct {
	logger *logger.Logger
}

func New(logger *logger.Logger) *Placeholder {
	return &Placeholder{
		logger: logger,
	}
}
func (p *Placeholder) Name() string {
	return manifest.PlaceholderKey
}

func (*Placeholder) StartNetworkHeight() int64 {
	return tools.V18.Height()
}

func (*Placeholder) Methods(_ context.Context, network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error) {
	switch {
	// all legacy version
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V17)...):
		return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{}, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	case tools.V18.IsSupported(network, height):
		return placeholderv10.Methods, nil
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return placeholderv11.Methods, nil
	case tools.V21.IsSupported(network, height):
		return placeholderv12.Methods, nil
	case tools.V22.IsSupported(network, height):
		return placeholderv13.Methods, nil
	case tools.V23.IsSupported(network, height):
		return placeholderv14.Methods, nil
	case tools.V24.IsSupported(network, height):
		return placeholderv15.Methods, nil
	default:
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
}

func (p *Placeholder) Parse(_ context.Context, network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, mainMsgCid cid.Cid, key filTypes.TipSetKey) (map[string]interface{}, *types.AddressInfo, error) {
	resp, err := p.parsePlaceholderAny(msg.Params, msgRct.Return)
	return resp, nil, err
}

func (p *Placeholder) TransactionTypes() map[string]any {
	return map[string]any{}
}

func (p *Placeholder) parsePlaceholderAny(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	metadata[parser.ParamsKey] = rawParams
	metadata[parser.ReturnKey] = rawReturn

	return metadata, nil
}
