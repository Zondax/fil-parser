package datacap

import (
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/zondax/fil-parser/parser"
)

type Datacap struct{}

func (d *Datacap) Name() string {
	return manifest.DatacapKey
}

func (p *Datacap) Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt) (map[string]interface{}, error) {
	switch txType {
	case parser.MethodConstructor:
		// return p.Constructor(network, height, msg.Params)
	case parser.MethodMintExported:
		return p.MintExported(network, height, msg.Params, msgRct.Return)
	case parser.MethodDestroyExported:
		return p.DestroyExported(network, height, msg.Params, msgRct.Return)
	case parser.MethodNameExported:
		return p.NameExported(msgRct.Return)
	case parser.MethodSymbolExported:
		return p.SymbolExported(msgRct.Return)
	case parser.MethodTotalSupplyExported:
		return p.TotalSupplyExported(msgRct.Return)
	case parser.MethodBalanceExported:
		// return p.BalanceExported(network, height, msg.Params, msgRct.Return)
	case parser.MethodTransferExported:
		return p.TransferExported(network, height, msg.Params, msgRct.Return)
	case parser.MethodTransferFromExported:
		// return p.TransferFromExported(network, height, msg.Params, msgRct.Return)
	case parser.MethodIncreaseAllowanceExported:
		return p.IncreaseAllowanceExported(network, height, msg.Params, msgRct.Return)
	case parser.MethodDecreaseAllowanceExported:
		return p.DecreaseAllowanceExported(network, height, msg.Params, msgRct.Return)
	case parser.MethodRevokeAllowanceExported:
		return p.RevokeAllowanceExported(network, height, msg.Params, msgRct.Return)
	case parser.MethodBurnExported:
		return p.BurnExported(network, height, msg.Params, msgRct.Return)
	case parser.MethodBurnFromExported:
		return p.BurnFromExported(network, height, msg.Params, msgRct.Return)
	case parser.MethodAllowanceExported:
		return p.AllowanceExported(network, height, msg.Params, msgRct.Return)
	case parser.MethodGranularityExported:
		return p.GranularityExported(network, height, msgRct.Return)
	case parser.UnknownStr:
		// return p.unknownMetadata(msg.Params, msgRct.Return)
	}
	return map[string]interface{}{}, parser.ErrUnknownMethod
}
