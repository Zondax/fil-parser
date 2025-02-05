package datacap

import (
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
	"go.uber.org/zap"
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
func (p *Datacap) Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, _ cid.Cid, _ filTypes.TipSetKey) (map[string]interface{}, *types.AddressInfo, error) {
	switch txType {
	case parser.MethodConstructor:
		// resp, err := p.Constructor(network, height, msg.Params)
		// return resp, nil, err
	case parser.MethodMintExported:
		resp, err := p.MintExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodDestroyExported:
		resp, err := p.DestroyExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodNameExported:
		resp, err := p.NameExported(msgRct.Return)
		return resp, nil, err
	case parser.MethodSymbolExported:
		resp, err := p.SymbolExported(msgRct.Return)
		return resp, nil, err
	case parser.MethodTotalSupplyExported:
		resp, err := p.TotalSupplyExported(msgRct.Return)
		return resp, nil, err
	case parser.MethodBalanceExported:
		resp, err := p.BalanceExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodTransferExported:
		resp, err := p.TransferExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodTransferFromExported:
		resp, err := p.TransferFromExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodIncreaseAllowanceExported:
		resp, err := p.IncreaseAllowanceExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodDecreaseAllowanceExported:
		resp, err := p.DecreaseAllowanceExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodRevokeAllowanceExported:
		resp, err := p.RevokeAllowanceExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodBurnExported:
		resp, err := p.BurnExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodBurnFromExported:
		resp, err := p.BurnFromExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodAllowanceExported:
		resp, err := p.AllowanceExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodGranularityExported:
		resp, err := p.GranularityExported(network, height, msgRct.Return)
		return resp, nil, err
	case parser.UnknownStr:
		// return p.unknownMetadata(msg.Params, msgRct.Return)
	}
	return map[string]interface{}{}, nil, parser.ErrUnknownMethod
}

func (d *Datacap) TransactionTypes() map[string]any {
	return map[string]any{
		parser.MethodConstructor:               nil,
		parser.MethodMintExported:              d.MintExported,
		parser.MethodDestroyExported:           d.DestroyExported,
		parser.MethodNameExported:              d.NameExported,
		parser.MethodSymbolExported:            d.SymbolExported,
		parser.MethodTotalSupplyExported:       d.TotalSupplyExported,
		parser.MethodBalanceExported:           nil,
		parser.MethodTransferExported:          d.TransferExported,
		parser.MethodTransferFromExported:      nil,
		parser.MethodIncreaseAllowanceExported: d.IncreaseAllowanceExported,
		parser.MethodDecreaseAllowanceExported: d.DecreaseAllowanceExported,
		parser.MethodRevokeAllowanceExported:   d.RevokeAllowanceExported,
		parser.MethodBurnExported:              d.BurnExported,
		parser.MethodBurnFromExported:          d.BurnFromExported,
		parser.MethodAllowanceExported:         d.AllowanceExported,
		parser.MethodGranularityExported:       d.GranularityExported,
	}
}
