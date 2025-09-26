package datacap

import (
	"context"
	"fmt"

	"github.com/zondax/golem/pkg/logger"

	"github.com/filecoin-project/go-state-types/abi"
	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"

	datacapv10 "github.com/filecoin-project/go-state-types/builtin/v10/datacap"
	datacapv11 "github.com/filecoin-project/go-state-types/builtin/v11/datacap"
	datacapv12 "github.com/filecoin-project/go-state-types/builtin/v12/datacap"
	datacapv13 "github.com/filecoin-project/go-state-types/builtin/v13/datacap"
	datacapv14 "github.com/filecoin-project/go-state-types/builtin/v14/datacap"
	datacapv15 "github.com/filecoin-project/go-state-types/builtin/v15/datacap"
	datacapv16 "github.com/filecoin-project/go-state-types/builtin/v16/datacap"
	datacapv17 "github.com/filecoin-project/go-state-types/builtin/v17/datacap"
	datacapv9 "github.com/filecoin-project/go-state-types/builtin/v9/datacap"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
)

type Datacap struct {
	logger *logger.Logger
}

func New(logger *logger.Logger) *Datacap {
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

func legacyMethods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
		abi.MethodNum(1):  datacapv9.Methods[abi.MethodNum(1)],
		abi.MethodNum(2):  datacapv9.Methods[abi.MethodNum(2)],
		abi.MethodNum(3):  datacapv9.Methods[abi.MethodNum(3)],
		abi.MethodNum(10): datacapv9.Methods[abi.MethodNum(10)],
		abi.MethodNum(11): datacapv9.Methods[abi.MethodNum(11)],
		abi.MethodNum(12): datacapv9.Methods[abi.MethodNum(12)],
		abi.MethodNum(13): datacapv9.Methods[abi.MethodNum(13)],
		abi.MethodNum(14): datacapv9.Methods[abi.MethodNum(14)],
		abi.MethodNum(15): datacapv9.Methods[abi.MethodNum(15)],
		abi.MethodNum(16): datacapv9.Methods[abi.MethodNum(16)],
		abi.MethodNum(17): datacapv9.Methods[abi.MethodNum(17)],
		abi.MethodNum(18): datacapv9.Methods[abi.MethodNum(18)],
		abi.MethodNum(19): datacapv9.Methods[abi.MethodNum(19)],
		abi.MethodNum(20): datacapv9.Methods[abi.MethodNum(20)],
		abi.MethodNum(21): datacapv9.Methods[abi.MethodNum(21)],
	}
}

var methods = map[string]map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
	tools.V17.String(): actors.CopyMethods(datacapv9.Methods, legacyMethods()),
	tools.V18.String(): actors.CopyMethods(datacapv10.Methods, legacyMethods()),
	tools.V19.String(): actors.CopyMethods(datacapv11.Methods, legacyMethods()),
	tools.V20.String(): actors.CopyMethods(datacapv11.Methods, legacyMethods()),
	tools.V21.String(): actors.CopyMethods(datacapv12.Methods, legacyMethods()),
	tools.V22.String(): actors.CopyMethods(datacapv13.Methods, legacyMethods()),
	tools.V23.String(): actors.CopyMethods(datacapv14.Methods, legacyMethods()),
	tools.V24.String(): actors.CopyMethods(datacapv15.Methods, legacyMethods()),
	tools.V25.String(): actors.CopyMethods(datacapv16.Methods, legacyMethods()),
	tools.V26.String(): actors.CopyMethods(datacapv16.Methods, legacyMethods()),
	tools.V27.String(): actors.CopyMethods(datacapv17.Methods, legacyMethods()),
}

func (d *Datacap) Methods(_ context.Context, network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error) {
	version := tools.VersionFromHeight(network, height)
	methods, ok := methods[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return methods, nil
}

func (p *Datacap) Parse(_ context.Context, network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, _ cid.Cid, _ filTypes.TipSetKey, canonical bool) (map[string]interface{}, *types.AddressInfo, error) {
	switch txType {
	case parser.MethodSend:
		resp := actors.ParseSend(msg)
		return resp, nil, nil
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
		parser.MethodSend:                      actors.ParseSend,
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
