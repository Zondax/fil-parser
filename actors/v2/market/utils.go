package market

import "github.com/filecoin-project/go-address"

func getAddressAsString(addr any) string {
	if address, ok := addr.(address.Address); ok {
		return address.String()
	}
	return ""
}
