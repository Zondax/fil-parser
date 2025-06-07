package datacap

import (
	"fmt"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

func (d *Datacap) MintExported(network string, height int64, raw, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := mintParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := mintReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	metadata, err := parse(raw, rawReturn, true, params(), returnValue(), parser.ParamsKey)
	if err != nil {
		return metadata, err
	}

	balance, supply, recipientData, err := getMintReturnFields(params())
	if err != nil {
		return metadata, err
	}

	allocations, err := d.verifreg.ParseFRC46TokenOperatorDataReq(network, height, recipientData)
	if err != nil {
		return metadata, err
	}

	metadata[parser.ParamsKey] = map[string]interface{}{
		"balance":       balance,
		"supply":        supply,
		"recipientData": allocations,
	}
	return metadata, nil
}
