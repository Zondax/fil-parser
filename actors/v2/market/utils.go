package market

import (
	"bytes"
	"fmt"

	"github.com/filecoin-project/go-address"
	cbg "github.com/whyrusleeping/cbor-gen"
)

func getAddressAsString(addr any) string {
	if address, ok := addr.(*address.Address); ok {
		return address.String()
	}
	return ""
}

func parseCBORArray(network string, height int64, raw []byte, getParam func(network string, height int64) (cbg.CBORUnmarshaler, error)) ([]cbg.CBORUnmarshaler, error) {
	cborReader := cbg.NewCborReader(bytes.NewReader(raw))
	maj, l, err := cborReader.ReadHeader()
	if err != nil {
		return nil, fmt.Errorf("error reading CBOR header: %w", err)
	}
	if maj != cbg.MajArray {
		return nil, fmt.Errorf("expected array, got %d", maj)
	}

	elements := []cbg.CBORUnmarshaler{}
	for i := uint64(0); i < l; i++ {
		parsed, err := getParam(network, height)
		if err != nil {
			return nil, fmt.Errorf("error getting sector content changed return: %w", err)
		}
		if err := parsed.UnmarshalCBOR(cborReader); err != nil {
			return nil, fmt.Errorf("error unmarshaling element %d: %w", i, err)
		}
		elements = append(elements, parsed)
	}
	return elements, nil
}
