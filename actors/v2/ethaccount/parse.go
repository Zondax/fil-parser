package ethaccount

import (
	"fmt"

	"github.com/ipfs/go-cid"
	"go.uber.org/zap"

	"github.com/filecoin-project/go-state-types/abi"
	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"

	ethaccountv10 "github.com/filecoin-project/go-state-types/builtin/v10/ethaccount"
	ethaccountv11 "github.com/filecoin-project/go-state-types/builtin/v11/ethaccount"
	ethaccountv12 "github.com/filecoin-project/go-state-types/builtin/v12/ethaccount"
	ethaccountv13 "github.com/filecoin-project/go-state-types/builtin/v13/ethaccount"
	ethaccountv14 "github.com/filecoin-project/go-state-types/builtin/v14/ethaccount"
	ethaccountv15 "github.com/filecoin-project/go-state-types/builtin/v15/ethaccount"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
)

type EthAccount struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) *EthAccount {
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

func (e *EthAccount) Methods(network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error) {
	switch {
	// all legacy version
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V17)...):
		return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{}, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	case tools.V18.IsSupported(network, height):
		return ethaccountv10.Methods, nil
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return ethaccountv11.Methods, nil
	case tools.V21.IsSupported(network, height):
		return ethaccountv12.Methods, nil
	case tools.V22.IsSupported(network, height):
		return ethaccountv13.Methods, nil
	case tools.V23.IsSupported(network, height):
		return ethaccountv14.Methods, nil
	case tools.V24.IsSupported(network, height):
		return ethaccountv15.Methods, nil
	default:
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
}

func (e *EthAccount) Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, mainMsgCid cid.Cid, key filTypes.TipSetKey) (map[string]interface{}, *types.AddressInfo, error) {
	var resp map[string]interface{}
	var err error
	switch txType {
	case parser.MethodConstructor:
		resp, err = e.Constructor()
	default:
		resp, err = e.parseEthAccountAny(msg.Params, msgRct.Return)
	}

	return resp, nil, err
}

func (e *EthAccount) TransactionTypes() map[string]any {
	return map[string]any{
		parser.MethodConstructor: e.Constructor,
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
