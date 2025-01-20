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

func NameExported(rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawReturn)
	var r abi.CborString
	err := r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r
	return metadata, nil
}

func SymbolExported(rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawReturn)
	var r abi.CborString
	err := r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r
	return metadata, nil
}

func TotalSupplyExported(rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawReturn)
	var r abi.TokenAmount
	err := r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r // TODO: .uint64()??
	return metadata, nil
}
