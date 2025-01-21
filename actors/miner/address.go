package miner

import (
	"fmt"

	"github.com/filecoin-project/go-address"
	miner15 "github.com/filecoin-project/go-state-types/builtin/v15/miner"
)

func ChangeMultiaddrs(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*miner15.ChangeMultiaddrsParams, *miner15.ChangeMultiaddrsParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func ChangePeerID(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*miner15.ChangePeerIDParams, *miner15.ChangePeerIDParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func ChangeWorkerAddress(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*miner15.ChangeWorkerAddressParams, *miner15.ChangeWorkerAddressParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func ChangeOwnerAddress(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*address.Address, *address.Address](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func IsControllingAddressExported(height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*miner15.IsControllingAddressParams, *miner15.IsControllingAddressReturn](rawParams, rawReturn, true)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func GetOwner(height int64, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*miner15.GetOwnerReturn, *miner15.GetOwnerReturn](rawReturn, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func GetPeerID(height int64, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*miner15.GetPeerIDReturn, *miner15.GetPeerIDReturn](rawReturn, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func GetMultiaddrs(height int64, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*miner15.GetMultiAddrsReturn, *miner15.GetMultiAddrsReturn](rawReturn, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}
