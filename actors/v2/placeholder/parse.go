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
	placeholderv16 "github.com/filecoin-project/go-state-types/builtin/v16/placeholder"
	placeholderv17 "github.com/filecoin-project/go-state-types/builtin/v17/placeholder"
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

var methods = map[string]map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
	tools.V18.String(): actors.CopyMethods(placeholderv10.Methods),
	tools.V19.String(): actors.CopyMethods(placeholderv11.Methods),
	tools.V20.String(): actors.CopyMethods(placeholderv11.Methods),
	tools.V21.String(): actors.CopyMethods(placeholderv12.Methods),
	tools.V22.String(): actors.CopyMethods(placeholderv13.Methods),
	tools.V23.String(): actors.CopyMethods(placeholderv14.Methods),
	tools.V24.String(): actors.CopyMethods(placeholderv15.Methods),
	tools.V25.String(): actors.CopyMethods(placeholderv16.Methods),
	tools.V26.String(): actors.CopyMethods(placeholderv16.Methods),
	tools.V27.String(): actors.CopyMethods(placeholderv17.Methods),
}

func (*Placeholder) Methods(_ context.Context, network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error) {
	version := tools.VersionFromHeight(network, height)
	methods, ok := methods[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return methods, nil
}

func (p *Placeholder) Parse(_ context.Context, network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, mainMsgCid cid.Cid, key filTypes.TipSetKey, canonical bool) (map[string]interface{}, *types.AddressInfo, error) {
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
