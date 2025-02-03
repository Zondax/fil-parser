package datacap

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/zondax/fil-parser/parser"
)

func (d *Datacap) NameExported(rawReturn []byte) (map[string]interface{}, error) {
	var params abi.CborString
	var r abi.CborString
	data, err := parse(nil, rawReturn, false, &params, &r)
	if err != nil {
		return nil, err
	}
	data[parser.ReturnKey] = data[parser.ParamsKey]
	return data, nil
}

func (d *Datacap) SymbolExported(rawReturn []byte) (map[string]interface{}, error) {
	var params abi.CborString
	var r abi.CborString
	data, err := parse(nil, rawReturn, false, &params, &r)
	if err != nil {
		return nil, err
	}
	data[parser.ReturnKey] = data[parser.ParamsKey]
	return data, nil
}

func (d *Datacap) TotalSupplyExported(rawReturn []byte) (map[string]interface{}, error) {
	var params abi.TokenAmount
	var r abi.TokenAmount
	data, err := parse(nil, rawReturn, false, &params, &r)
	if err != nil {
		return nil, err
	}
	data[parser.ReturnKey] = data[parser.ParamsKey]
	return data, nil
}
