package datacap

import (
	"fmt"

	"github.com/ipfs/go-cid"
	"go.uber.org/zap"

	"github.com/filecoin-project/go-state-types/abi"
	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"

	datacapv10 "github.com/filecoin-project/go-state-types/builtin/v10/datacap"
	datacapv11 "github.com/filecoin-project/go-state-types/builtin/v11/datacap"
	datacapv12 "github.com/filecoin-project/go-state-types/builtin/v12/datacap"
	datacapv13 "github.com/filecoin-project/go-state-types/builtin/v13/datacap"
	datacapv14 "github.com/filecoin-project/go-state-types/builtin/v14/datacap"
	datacapv15 "github.com/filecoin-project/go-state-types/builtin/v15/datacap"
	datacapv9 "github.com/filecoin-project/go-state-types/builtin/v9/datacap"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
)

type Datacap struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) *Datacap {
	return &Datacap{
		logger: logger,
	}
}

func (d *Datacap) Name() string {
	return manifest.DatacapKey
}

func (*Datacap) StartNetworkHeight() int64 {
	return tools.V17.Height()
}

func (d *Datacap) Methods(network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error) {
	switch {
	// all legacy version
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V16)...):
		return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{}, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	case tools.V17.IsSupported(network, height):
		return datacapv9.Methods, nil
	case tools.V18.IsSupported(network, height):
		return datacapv10.Methods, nil
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return datacapv11.Methods, nil
	case tools.V21.IsSupported(network, height):
		return datacapv12.Methods, nil
	case tools.V22.IsSupported(network, height):
		return datacapv13.Methods, nil
	case tools.V23.IsSupported(network, height):
		return datacapv14.Methods, nil
	case tools.V24.IsSupported(network, height):
		return datacapv15.Methods, nil
	default:
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
}

func (p *Datacap) Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, _ cid.Cid, _ filTypes.TipSetKey) (map[string]interface{}, *types.AddressInfo, error) {
	switch txType {
	case parser.MethodConstructor:
		resp, err := actors.ParseConstructor(msg.Params)
		return resp, nil, err
	case parser.MethodMintExported, parser.MethodMint:
		resp, err := p.MintExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodDestroyExported, parser.MethodDestroy:
		resp, err := p.DestroyExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodNameExported, parser.MethodName:
		resp, err := p.NameExported(msgRct.Return)
		return resp, nil, err
	case parser.MethodSymbolExported, parser.MethodSymbol:
		resp, err := p.SymbolExported(msgRct.Return)
		return resp, nil, err
	case parser.MethodTotalSupplyExported, parser.MethodTotalSupply:
		resp, err := p.TotalSupplyExported(msgRct.Return)
		return resp, nil, err
	case parser.MethodBalanceExported:
		resp, err := p.BalanceExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodTransferExported, parser.MethodTransfer:
		resp, err := p.TransferExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodTransferFromExported, parser.MethodTransferFrom:
		resp, err := p.TransferFromExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodIncreaseAllowanceExported, parser.MethodIncreaseAllowance:
		resp, err := p.IncreaseAllowanceExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodDecreaseAllowanceExported, parser.MethodDecreaseAllowance:
		resp, err := p.DecreaseAllowanceExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodRevokeAllowanceExported, parser.MethodRevokeAllowance:
		resp, err := p.RevokeAllowanceExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodBurnExported, parser.MethodBurn:
		resp, err := p.BurnExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodBurnFromExported, parser.MethodBurnFrom:
		resp, err := p.BurnFromExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodAllowanceExported, parser.MethodAllowance:
		resp, err := p.AllowanceExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodGranularityExported:
		resp, err := p.GranularityExported(network, height, msgRct.Return)
		return resp, nil, err
	case parser.MethodBalanceOf:
		resp, err := p.BalanceOf(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.UnknownStr:
		resp, err := actors.ParseUnknownMetadata(msg.Params, msgRct.Return)
		return resp, nil, err
	}
	return map[string]interface{}{}, nil, parser.ErrUnknownMethod
}

func (d *Datacap) TransactionTypes() map[string]any {
	return map[string]any{
		parser.MethodConstructor:               actors.ParseConstructor,
		parser.MethodMint:                      d.MintExported,
		parser.MethodMintExported:              d.MintExported,
		parser.MethodDestroy:                   d.DestroyExported,
		parser.MethodDestroyExported:           d.DestroyExported,
		parser.MethodName:                      d.NameExported,
		parser.MethodNameExported:              d.NameExported,
		parser.MethodSymbol:                    d.SymbolExported,
		parser.MethodSymbolExported:            d.SymbolExported,
		parser.MethodTotalSupply:               d.TotalSupplyExported,
		parser.MethodTotalSupplyExported:       d.TotalSupplyExported,
		parser.MethodBalanceExported:           d.BalanceExported,
		parser.MethodTransfer:                  d.TransferExported,
		parser.MethodTransferExported:          d.TransferExported,
		parser.MethodTransferFromExported:      d.TransferFromExported,
		parser.MethodTransferFrom:              d.TransferFromExported,
		parser.MethodIncreaseAllowance:         d.IncreaseAllowanceExported,
		parser.MethodIncreaseAllowanceExported: d.IncreaseAllowanceExported,
		parser.MethodDecreaseAllowance:         d.DecreaseAllowanceExported,
		parser.MethodDecreaseAllowanceExported: d.DecreaseAllowanceExported,
		parser.MethodRevokeAllowance:           d.RevokeAllowanceExported,
		parser.MethodRevokeAllowanceExported:   d.RevokeAllowanceExported,
		parser.MethodBurn:                      d.BurnExported,
		parser.MethodBurnExported:              d.BurnExported,
		parser.MethodBurnFrom:                  d.BurnFromExported,
		parser.MethodBurnFromExported:          d.BurnFromExported,
		parser.MethodAllowance:                 d.AllowanceExported,
		parser.MethodAllowanceExported:         d.AllowanceExported,
		parser.MethodGranularityExported:       d.GranularityExported,
		parser.MethodBalanceOf:                 d.BalanceOf,
	}
}
