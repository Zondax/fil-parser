package ethaccount

import (
	"context"
	"fmt"

	"github.com/zondax/golem/pkg/logger"

	"github.com/filecoin-project/go-state-types/abi"
	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"

	ethaccountv10 "github.com/filecoin-project/go-state-types/builtin/v10/ethaccount"
	ethaccountv11 "github.com/filecoin-project/go-state-types/builtin/v11/ethaccount"
	ethaccountv12 "github.com/filecoin-project/go-state-types/builtin/v12/ethaccount"
	ethaccountv13 "github.com/filecoin-project/go-state-types/builtin/v13/ethaccount"
	ethaccountv14 "github.com/filecoin-project/go-state-types/builtin/v14/ethaccount"
	ethaccountv15 "github.com/filecoin-project/go-state-types/builtin/v15/ethaccount"
	ethaccountv16 "github.com/filecoin-project/go-state-types/builtin/v16/ethaccount"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
)

type EthAccount struct {
	logger *logger.Logger
}

func New(logger *logger.Logger) *EthAccount {
	return &EthAccount{
		logger: logger,
	}
}

func (e *EthAccount) Name() string {
	return manifest.EthAccountKey
}

func (*EthAccount) StartNetworkHeight() int64 {
	return tools.V18.Height()
}

var methods = map[string]map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
	tools.V18.String(): actors.CopyMethods(ethaccountv10.Methods),
	tools.V19.String(): actors.CopyMethods(ethaccountv11.Methods),
	tools.V20.String(): actors.CopyMethods(ethaccountv11.Methods),
	tools.V21.String(): actors.CopyMethods(ethaccountv12.Methods),
	tools.V22.String(): actors.CopyMethods(ethaccountv13.Methods),
	tools.V23.String(): actors.CopyMethods(ethaccountv14.Methods),
	tools.V24.String(): actors.CopyMethods(ethaccountv15.Methods),
	tools.V25.String(): actors.CopyMethods(ethaccountv16.Methods),
}

func (e *EthAccount) Methods(_ context.Context, network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error) {
	version := tools.VersionFromHeight(network, height)
	methods, ok := methods[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return methods, nil
}

func (e *EthAccount) Parse(_ context.Context, network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, mainMsgCid cid.Cid, key filTypes.TipSetKey) (map[string]interface{}, *types.AddressInfo, error) {
	var resp map[string]interface{}
	var err error
	switch txType {
	case parser.MethodConstructor:
		resp, err = e.Constructor()
	case parser.MethodFallback:
		resp, err = e.Fallback(network, height, msg.Params)
	default:
		resp, err = e.parseEthAccountAny(msg.Params, msgRct.Return)
	}

	return resp, nil, err
}

func (e *EthAccount) TransactionTypes() map[string]any {
	return map[string]any{
		parser.MethodConstructor: e.Constructor,
		parser.MethodSend:        actors.ParseSend,
	}
}

func (e *EthAccount) Constructor() (map[string]interface{}, error) {
	return e.parseEthAccountAny(nil, nil)
}

func (e *EthAccount) parseEthAccountAny(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	metadata[parser.ParamsKey] = rawParams
	metadata[parser.ReturnKey] = rawReturn

	return metadata, nil
}

func (e *EthAccount) Fallback(network string, height int64, raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	metadata[parser.ParamsRawKey] = raw
	return metadata, nil
}
