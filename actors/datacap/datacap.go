package datacap

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/zondax/fil-parser/parser"
)

type Datacap struct{}

func (d *Datacap) NameExported(rawReturn []byte) (map[string]interface{}, error) {
	data, err := parse[*abi.CborString, *abi.CborString](nil, rawReturn, false)
	if err != nil {
		return nil, err
	}
	data[parser.ReturnKey] = data[parser.ParamsKey]
	return data, nil
}

func (d *Datacap) SymbolExported(rawReturn []byte) (map[string]interface{}, error) {
	data, err := parse[*abi.CborString, *abi.CborString](nil, rawReturn, false)
	if err != nil {
		return nil, err
	}
	data[parser.ReturnKey] = data[parser.ParamsKey]
	return data, nil
}

func (d *Datacap) TotalSupplyExported(rawReturn []byte) (map[string]interface{}, error) {
	data, err := parse[*abi.TokenAmount, *abi.TokenAmount](nil, rawReturn, false)
	if err != nil {
		return nil, err
	}
	data[parser.ReturnKey] = data[parser.ParamsKey]
	return data, nil
}
