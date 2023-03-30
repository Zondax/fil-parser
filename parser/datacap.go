package parser

import (
	"bytes"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin/v11/datacap"
	filTypes "github.com/filecoin-project/lotus/chain/types"
)

func (p *Parser) parseDatacap(txType string, msg *filTypes.Message, msgRct *filTypes.MessageReceipt) (map[string]interface{}, error) {
	switch txType {
	case MethodConstructor:
		return p.parseConstructor(msg.Params)
	case MethodMintExported:
		return p.mintExported(msg.Params, msgRct.Return)
	case MethodDestroyExported:
		return p.destroyExported(msg.Params, msgRct.Return)
	case MethodNameExported:
		return p.nameExported(msgRct.Return)
	case MethodSymbolExported:
		return p.symbolExported(msgRct.Return)
	case MethodTotalSupplyExported:
		return p.totalSupplyExported(msgRct.Return)
	case MethodBalanceExported:
		return p.balanceExported(msg.Params, msgRct.Return)
	case MethodTransferExported:
		return p.transferExported(msg.Params, msgRct.Return)
	case MethodTransferFromExported:
		return p.transferFromExported(msg.Params, msgRct.Return)
	case MethodIncreaseAllowanceExported:
		return p.increaseAllowanceExported(msg.Params, msgRct.Return)
	case MethodDecreaseAllowanceExported:
		return p.decreaseAllowanceExported(msg.Params, msgRct.Return)
	case MethodRevokeAllowanceExported:
		return p.revokeExportedAllowanceExported(msg.Params, msgRct.Return)
	case MethodBurnExported:
		return p.burnExported(msg.Params, msgRct.Return)
	case MethodBurnFromExported:
		return p.burnFromExported(msg.Params, msgRct.Return)
	case MethodAllowanceExported:
		return p.allowanceExported(msg.Params, msgRct.Return)
	case MethodGranularityExported:
		return p.granularityExported(msgRct.Return)
	case UnknownStr:
		return p.unknownMetadata(msg.Params, msgRct.Return)
	}
	return map[string]interface{}{}, errUnknownMethod
}

func (p *Parser) mintExported(raw, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params datacap.MintParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params

	reader = bytes.NewReader(rawReturn)
	var r datacap.MintReturn
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = r
	return metadata, nil
}

func (p *Parser) destroyExported(raw, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params datacap.DestroyParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params

	reader = bytes.NewReader(rawReturn)
	var r datacap.BurnReturn
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = r
	return metadata, nil
}

func (p *Parser) nameExported(rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawReturn)
	var r abi.CborString
	err := r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = r
	return metadata, nil
}

func (p *Parser) symbolExported(rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawReturn)
	var r abi.CborString
	err := r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = r
	return metadata, nil
}

func (p *Parser) totalSupplyExported(rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawReturn)
	var r abi.TokenAmount
	err := r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = r // TODO: .uint64()??
	return metadata, nil
}

func (p *Parser) balanceExported(raw, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params address.Address
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params.String()

	reader = bytes.NewReader(rawReturn)
	var r abi.TokenAmount
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = r // TODO: .uint64()??
	return metadata, nil
}

func (p *Parser) transferExported(raw, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params datacap.TransferParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params

	reader = bytes.NewReader(rawReturn)
	var r datacap.TransferReturn
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = r
	return metadata, nil
}

func (p *Parser) transferFromExported(raw, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params datacap.TransferFromParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params

	reader = bytes.NewReader(rawReturn)
	var r datacap.TransferFromReturn
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = r
	return metadata, nil
}

func (p *Parser) increaseAllowanceExported(raw, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params datacap.IncreaseAllowanceParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params

	reader = bytes.NewReader(rawReturn)
	var r abi.TokenAmount
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = r // TODO: .uint64()??
	return metadata, nil
}

func (p *Parser) decreaseAllowanceExported(raw, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params datacap.DecreaseAllowanceParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params

	reader = bytes.NewReader(rawReturn)
	var r abi.TokenAmount
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = r // TODO: .uint64()??
	return metadata, nil
}

func (p *Parser) revokeExportedAllowanceExported(raw, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params datacap.RevokeAllowanceParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params

	reader = bytes.NewReader(rawReturn)
	var r abi.TokenAmount
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = r // TODO: .uint64()??
	return metadata, nil
}

func (p *Parser) burnExported(raw, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params datacap.BurnParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params

	reader = bytes.NewReader(rawReturn)
	var r datacap.BurnReturn
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = r
	return metadata, nil
}

func (p *Parser) burnFromExported(raw, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params datacap.BurnFromParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params

	reader = bytes.NewReader(rawReturn)
	var r datacap.BurnFromReturn
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = r
	return metadata, nil
}

func (p *Parser) allowanceExported(raw, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params datacap.GetAllowanceParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params

	reader = bytes.NewReader(rawReturn)
	var r abi.TokenAmount
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = r // TODO: .uint64()??
	return metadata, nil
}

func (p *Parser) granularityExported(rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawReturn)
	var r datacap.GranularityReturn
	err := r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = r
	return metadata, nil
}
