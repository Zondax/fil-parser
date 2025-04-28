package evm

import (
	"context"

	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"

	actor_tools "github.com/zondax/fil-parser/actors/v2/tools"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
)

func (p *Evm) Parse(_ context.Context, network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, msgCid cid.Cid, tipsetKey filTypes.TipSetKey) (map[string]interface{}, *types.AddressInfo, error) {
	metadata := make(map[string]interface{})
	switch txType {
	case parser.MethodSend:
		resp := actor_tools.ParseSend(msg)
		return resp, nil, nil
	case parser.MethodConstructor:
		resp, err := p.Constructor(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodResurrect: // TODO: not tested
		resp, err := p.Resurrect(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodInvokeContract, parser.MethodInvokeContractReadOnly:
		resp, err := p.InvokeContract(network, height, txType, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodInvokeContractDelegate:
		resp, err := p.InvokeContractDelegate(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodGetBytecode:
		resp, err := p.GetBytecode(network, height, msgRct.Return)
		return resp, nil, err
	case parser.MethodGetBytecodeHash: // TODO: not tested
		resp, err := p.GetBytecodeHash(network, height, msgRct.Return)
		return resp, nil, err
	case parser.MethodGetStorageAt: // TODO: not tested
		resp, err := p.GetStorageAt(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodInvokeContractFilecoinHandler:
		resp, err := p.InvokeContractFilecoinHandler(network, height, msgCid, msg.From, tipsetKey, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.UnknownStr:
		resp, err := actor_tools.ParseUnknownMetadata(msg.Params, msgRct.Return)
		return resp, nil, err
	}
	return metadata, nil, parser.ErrUnknownMethod
}

func (p *Evm) TransactionTypes() map[string]any {
	return map[string]any{
		parser.MethodSend:                   actor_tools.ParseSend,
		parser.MethodConstructor:            p.Constructor,
		parser.MethodResurrect:              p.Resurrect,
		parser.MethodInvokeContract:         p.InvokeContract,
		parser.MethodInvokeContractReadOnly: p.InvokeContract,
		parser.MethodInvokeContractDelegate: p.InvokeContractDelegate,
		parser.MethodGetBytecode:            p.GetBytecode,
		parser.MethodGetBytecodeHash:        p.GetBytecodeHash,
		parser.MethodGetStorageAt:           p.GetStorageAt,
	}
}
