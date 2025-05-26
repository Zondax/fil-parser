package common

import (
	"errors"
	"github.com/filecoin-project/go-address"
)

var (
	ErrKeyNotFound       = errors.New("key not found")
	ErrUnkownAddressType = errors.New("unknown address type")
	ErrEmptyValue        = errors.New("empty value")
)

func IsRobustAddress(add address.Address) (bool, error) {
	switch add.Protocol() {
	case address.BLS, address.SECP256K1, address.Actor, address.Delegated:
		return true, nil
	case address.ID:
		return false, nil
	default:
		// Consider unknown type as robust
		return true, ErrUnkownAddressType
	}
}
