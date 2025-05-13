package ethaccount

import (
	"context"
	"encoding/hex"
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
	"github.com/zondax/fil-parser/actors/v2/miner"
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

func customMethods(e *EthAccount) map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
		// This is a miner method verified from testing with CID: bafy2bzacealgb5zr5g2cc5emi7yc2mpragoufvt5lm54xzdkhdorpfgjhbshi on calibration.
		abi.MethodNum(23): {
			Name:   parser.MethodChangeOwnerAddressExported,
			Method: e.ChangeOwnerAddress,
		},
		// This is a miner method verified from testing with CID: f3vmqpcytevkwn6fktjd2zelo4lftq6xzsb2vnmp2r3qarbr4vnso7c7y3nqi5gmxifp22m2pbqdctfxrwkmga on calibration.
		abi.MethodNum(18): {
			Name:   parser.MethodChangeMultiaddrsExported,
			Method: e.ChangeMultiAddrs,
		},
		// https://github.com/filecoin-project/ref-fvm/issues/835#issuecomment-1236096270
		abi.MethodNum(0): {
			Name:   parser.MethodValueTransfer,
			Method: e.ValueTransfer,
		},
	}
}

var methods = map[string]map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
	tools.V18.String(): actors.CopyMethods(ethaccountv10.Methods, customMethods(&EthAccount{})),
	tools.V19.String(): actors.CopyMethods(ethaccountv11.Methods, customMethods(&EthAccount{})),
	tools.V20.String(): actors.CopyMethods(ethaccountv11.Methods, customMethods(&EthAccount{})),
	tools.V21.String(): actors.CopyMethods(ethaccountv12.Methods, customMethods(&EthAccount{})),
	tools.V22.String(): actors.CopyMethods(ethaccountv13.Methods, customMethods(&EthAccount{})),
	tools.V23.String(): actors.CopyMethods(ethaccountv14.Methods, customMethods(&EthAccount{})),
	tools.V24.String(): actors.CopyMethods(ethaccountv15.Methods, customMethods(&EthAccount{})),
	tools.V25.String(): actors.CopyMethods(ethaccountv16.Methods, customMethods(&EthAccount{})),
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
	case parser.MethodChangeOwnerAddressExported:
		resp, err = e.ChangeOwnerAddress(network, height, msg.Params)
	case parser.MethodChangeMultiaddrsExported:
		resp, err = e.ChangeMultiAddrs(network, height, msg.Params)
	case parser.MethodValueTransfer:
		resp, err = e.ValueTransfer(network, height, msg.Params)
	default:
		resp, err = e.parseEthAccountAny(msg.Params, msgRct.Return)
	}

	return resp, nil, err
}

func (e *EthAccount) TransactionTypes() map[string]any {
	return map[string]any{
		parser.MethodConstructor:                e.Constructor,
		parser.MethodSend:                       actors.ParseSend,
		parser.MethodChangeOwnerAddressExported: e.ChangeOwnerAddress,
		parser.MethodChangeMultiaddrsExported:   e.ChangeMultiAddrs,
		parser.MethodValueTransfer:              e.ValueTransfer,
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
	metadata[parser.ParamsRawKey] = hex.EncodeToString(raw)
	return metadata, nil
}

func (e *EthAccount) ChangeOwnerAddress(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	m := miner.New(e.logger)
	return m.ChangeOwnerAddressExported(network, height, rawParams)
}

func (e *EthAccount) ChangeMultiAddrs(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	m := miner.New(e.logger)
	return m.ChangeMultiaddrsExported(network, height, rawParams)
}

func (e *EthAccount) ValueTransfer(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	return map[string]interface{}{}, nil
}
