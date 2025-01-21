package datacap

import (
	"bytes"
	"io"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/zondax/fil-parser/parser"
)

type unmarshalCBOR interface {
	UnmarshalCBOR(io.Reader) error
}

func datacapGeneric[T unmarshalCBOR](rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawReturn)
	var r T
	err := r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r
	return metadata, nil
}

func NameExported(rawReturn []byte) (map[string]interface{}, error) {
	return datacapGeneric[*abi.CborString](rawReturn)
}

func SymbolExported(rawReturn []byte) (map[string]interface{}, error) {
	return datacapGeneric[*abi.CborString](rawReturn)
}

func TotalSupplyExported(rawReturn []byte) (map[string]interface{}, error) {
	return datacapGeneric[*abi.TokenAmount](rawReturn)
}
