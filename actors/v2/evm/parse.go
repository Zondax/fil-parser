package evm

import (
	"github.com/zondax/fil-parser/parser"
)

func (p *Evm) Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	switch txType {
	case parser.MethodConstructor:
		return p.Constructor(network, height, msg.Params)
	case parser.MethodResurrect: // TODO: not tested
		return p.Resurrect(network, height, msg.Params)
	case parser.MethodInvokeContract, parser.MethodInvokeContractReadOnly:
		return p.InvokeContract(network, height, msg.Params, msgRct.Return)
	case parser.MethodInvokeContractDelegate:
		return p.InvokeContractDelegate(network, height, msg.Params, msgRct.Return)
	case parser.MethodGetBytecode:
		return p.GetBytecode(network, height, msgRct.Return)
	case parser.MethodGetBytecodeHash: // TODO: not tested
		return p.GetBytecodeHash(network, height, msgRct.Return)
	case parser.MethodGetStorageAt: // TODO: not tested
		return p.GetStorageAt(network, height, msg.Params, msgRct.Return)
	case parser.UnknownStr:
		// return p.unknownMetadata(msg.Params, msgRct.Return)
	}
	return metadata, nil
}
