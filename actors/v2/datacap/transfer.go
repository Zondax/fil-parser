package datacap

import (
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

func (d *Datacap) TransferExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := transferParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := transferReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	metadata, err := parse(raw, rawReturn, true, params(), returnValue(), parser.ParamsKey)
	if err != nil {
		return metadata, err
	}

	to, amount, operatorData, err := getTransferParamsFields(params())
	if err != nil {
		return metadata, err
	}
	paramAllocations, err := d.verifreg.ParseFRC46TokenOperatorDataReq(network, height, operatorData)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = map[string]interface{}{
		"to":           to,
		"amount":       amount,
		"operatorData": paramAllocations,
	}

	fromBalance, toBalance, recipientData, err := getTransferReturnFields(returnValue())
	if err != nil {
		return metadata, err
	}
	returnAllocations, err := d.verifreg.ParseFRC46TokenOperatorDataReq(network, height, recipientData)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = map[string]interface{}{
		"fromBalance":   fromBalance,
		"toBalance":     toBalance,
		"recipientData": returnAllocations,
	}

	return metadata, nil
}

func (*Datacap) BalanceExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	return parse(raw, rawReturn, true, &address.Address{}, &abi.TokenAmount{}, parser.ParamsKey)
}

func (d *Datacap) TransferFromExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := transferFromParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := transferFromReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	metadata, err := parse(raw, rawReturn, true, params(), returnValue(), parser.ParamsKey)
	if err != nil {
		return metadata, err
	}

	from, to, amount, operatorData, err := getTransferFromParamsFields(params())
	if err != nil {
		return metadata, err
	}

	var paramAllocations any
	if len(operatorData) > 0 {
		paramAllocations, err = d.verifreg.ParseFRC46TokenOperatorDataReq(network, height, operatorData)
		if err != nil {
			return metadata, err
		}
	}

	metadata[parser.ParamsKey] = map[string]interface{}{
		"from":         from,
		"to":           to,
		"amount":       amount,
		"operatorData": paramAllocations,
	}

	fromBalance, toBalance, allowance, recipientData, err := getTransferFromReturnFields(returnValue())
	if err != nil {
		return metadata, err
	}
	var returnAllocations any
	if len(recipientData) > 0 {
		returnAllocations, err = d.verifreg.ParseFRC46TokenOperatorDataReq(network, height, recipientData)
		if err != nil {
			return metadata, err
		}
	}
	metadata[parser.ParamsKey] = map[string]interface{}{
		"fromBalance":   fromBalance,
		"toBalance":     toBalance,
		"allowance":     allowance,
		"recipientData": returnAllocations,
	}
	return metadata, nil
}

func (*Datacap) BalanceOf(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	return parse(raw, rawReturn, true, &address.Address{}, &abi.TokenAmount{}, parser.ParamsKey)
}
