package miner

import (
	"fmt"

	"github.com/filecoin-project/go-address"
	miner10 "github.com/filecoin-project/go-state-types/builtin/v10/miner"
	miner11 "github.com/filecoin-project/go-state-types/builtin/v11/miner"
	miner12 "github.com/filecoin-project/go-state-types/builtin/v12/miner"
	miner13 "github.com/filecoin-project/go-state-types/builtin/v13/miner"
	miner14 "github.com/filecoin-project/go-state-types/builtin/v14/miner"
	miner15 "github.com/filecoin-project/go-state-types/builtin/v15/miner"
	miner8 "github.com/filecoin-project/go-state-types/builtin/v8/miner"
	miner9 "github.com/filecoin-project/go-state-types/builtin/v9/miner"
	"github.com/zondax/fil-parser/tools"
)

func ChangeMultiaddrs(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		return parseGeneric[*miner15.ChangeMultiaddrsParams, *miner15.ChangeMultiaddrsParams](rawParams, nil, false)
	case tools.V14.IsSupported(height):
		return parseGeneric[*miner14.ChangeMultiaddrsParams, *miner14.ChangeMultiaddrsParams](rawParams, nil, false)
	case tools.V13.IsSupported(height):
		return parseGeneric[*miner13.ChangeMultiaddrsParams, *miner13.ChangeMultiaddrsParams](rawParams, nil, false)
	case tools.V12.IsSupported(height):
		return parseGeneric[*miner12.ChangeMultiaddrsParams, *miner12.ChangeMultiaddrsParams](rawParams, nil, false)
	case tools.V11.IsSupported(height):
		return parseGeneric[*miner11.ChangeMultiaddrsParams, *miner11.ChangeMultiaddrsParams](rawParams, nil, false)
	case tools.V10.IsSupported(height):
		return parseGeneric[*miner10.ChangeMultiaddrsParams, *miner10.ChangeMultiaddrsParams](rawParams, nil, false)
	case tools.V9.IsSupported(height):
		return parseGeneric[*miner9.ChangeMultiaddrsParams, *miner9.ChangeMultiaddrsParams](rawParams, nil, false)
	case tools.V8.IsSupported(height):
		return parseGeneric[*miner8.ChangeMultiaddrsParams, *miner8.ChangeMultiaddrsParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func ChangePeerID(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		return parseGeneric[*miner15.ChangePeerIDParams, *miner15.ChangePeerIDParams](rawParams, nil, false)
	case tools.V14.IsSupported(height):
		return parseGeneric[*miner14.ChangePeerIDParams, *miner14.ChangePeerIDParams](rawParams, nil, false)
	case tools.V13.IsSupported(height):
		return parseGeneric[*miner13.ChangePeerIDParams, *miner13.ChangePeerIDParams](rawParams, nil, false)
	case tools.V12.IsSupported(height):
		return parseGeneric[*miner12.ChangePeerIDParams, *miner12.ChangePeerIDParams](rawParams, nil, false)
	case tools.V11.IsSupported(height):
		return parseGeneric[*miner11.ChangePeerIDParams, *miner11.ChangePeerIDParams](rawParams, nil, false)
	case tools.V10.IsSupported(height):
		return parseGeneric[*miner10.ChangePeerIDParams, *miner10.ChangePeerIDParams](rawParams, nil, false)
	case tools.V9.IsSupported(height):
		return parseGeneric[*miner9.ChangePeerIDParams, *miner9.ChangePeerIDParams](rawParams, nil, false)
	case tools.V8.IsSupported(height):
		return parseGeneric[*miner8.ChangePeerIDParams, *miner8.ChangePeerIDParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func ChangeWorkerAddress(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		return parseGeneric[*miner15.ChangeWorkerAddressParams, *miner15.ChangeWorkerAddressParams](rawParams, nil, false)
	case tools.V14.IsSupported(height):
		return parseGeneric[*miner14.ChangeWorkerAddressParams, *miner14.ChangeWorkerAddressParams](rawParams, nil, false)
	case tools.V13.IsSupported(height):
		return parseGeneric[*miner13.ChangeWorkerAddressParams, *miner13.ChangeWorkerAddressParams](rawParams, nil, false)
	case tools.V12.IsSupported(height):
		return parseGeneric[*miner12.ChangeWorkerAddressParams, *miner12.ChangeWorkerAddressParams](rawParams, nil, false)
	case tools.V11.IsSupported(height):
		return parseGeneric[*miner11.ChangeWorkerAddressParams, *miner11.ChangeWorkerAddressParams](rawParams, nil, false)
	case tools.V10.IsSupported(height):
		return parseGeneric[*miner10.ChangeWorkerAddressParams, *miner10.ChangeWorkerAddressParams](rawParams, nil, false)
	case tools.V9.IsSupported(height):
		return parseGeneric[*miner9.ChangeWorkerAddressParams, *miner9.ChangeWorkerAddressParams](rawParams, nil, false)
	case tools.V8.IsSupported(height):
		return parseGeneric[*miner8.ChangeWorkerAddressParams, *miner8.ChangeWorkerAddressParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func ChangeOwnerAddress(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		return parseGeneric[*address.Address, *address.Address](rawParams, nil, false)
	case tools.V14.IsSupported(height):
		return parseGeneric[*address.Address, *address.Address](rawParams, nil, false)
	case tools.V13.IsSupported(height):
		return parseGeneric[*address.Address, *address.Address](rawParams, nil, false)
	case tools.V12.IsSupported(height):
		return parseGeneric[*address.Address, *address.Address](rawParams, nil, false)
	case tools.V11.IsSupported(height):
		return parseGeneric[*address.Address, *address.Address](rawParams, nil, false)
	case tools.V10.IsSupported(height):
		return parseGeneric[*address.Address, *address.Address](rawParams, nil, false)
	case tools.V9.IsSupported(height):
		return parseGeneric[*address.Address, *address.Address](rawParams, nil, false)
	case tools.V8.IsSupported(height):
		return parseGeneric[*address.Address, *address.Address](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func IsControllingAddressExported(height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		return parseGeneric[*miner15.IsControllingAddressParams, *miner15.IsControllingAddressReturn](rawParams, rawReturn, true)
	case tools.V14.IsSupported(height):
		return parseGeneric[*miner14.IsControllingAddressParams, *miner14.IsControllingAddressReturn](rawParams, rawReturn, true)
	case tools.V13.IsSupported(height):
		return parseGeneric[*miner13.IsControllingAddressParams, *miner13.IsControllingAddressReturn](rawParams, rawReturn, true)
	case tools.V12.IsSupported(height):
		return parseGeneric[*miner12.IsControllingAddressParams, *miner12.IsControllingAddressReturn](rawParams, rawReturn, true)
	case tools.V11.IsSupported(height):
		return parseGeneric[*miner11.IsControllingAddressParams, *miner11.IsControllingAddressReturn](rawParams, rawReturn, true)
	case tools.V10.IsSupported(height):
		return parseGeneric[*miner10.IsControllingAddressParams, *miner10.IsControllingAddressReturn](rawParams, rawReturn, true)
	case tools.V9.IsSupported(height):
		return parseGeneric[*miner9.IsControllingAddressParams, *miner9.IsControllingAddressReturn](rawParams, rawReturn, true)
	case tools.V8.IsSupported(height):
		return parseGeneric[*miner8.IsControllingAddressParams, *miner8.IsControllingAddressReturn](rawParams, rawReturn, true)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func GetOwner(height int64, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		return parseGeneric[*miner15.GetOwnerReturn, *miner15.GetOwnerReturn](rawReturn, nil, false)
	case tools.V14.IsSupported(height):
		return parseGeneric[*miner14.GetOwnerReturn, *miner14.GetOwnerReturn](rawReturn, nil, false)
	case tools.V13.IsSupported(height):
		return parseGeneric[*miner13.GetOwnerReturn, *miner13.GetOwnerReturn](rawReturn, nil, false)
	case tools.V12.IsSupported(height):
		return parseGeneric[*miner12.GetOwnerReturn, *miner12.GetOwnerReturn](rawReturn, nil, false)
	case tools.V11.IsSupported(height):
		return parseGeneric[*miner11.GetOwnerReturn, *miner11.GetOwnerReturn](rawReturn, nil, false)
	case tools.V10.IsSupported(height):
		return parseGeneric[*miner10.GetOwnerReturn, *miner10.GetOwnerReturn](rawReturn, nil, false)
	case tools.V9.IsSupported(height):
		return parseGeneric[*miner9.GetOwnerReturn, *miner9.GetOwnerReturn](rawReturn, nil, false)
	case tools.V8.IsSupported(height):
		return parseGeneric[*miner8.GetOwnerReturn, *miner8.GetOwnerReturn](rawReturn, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func GetPeerID(height int64, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		return parseGeneric[*miner15.GetPeerIDReturn, *miner15.GetPeerIDReturn](rawReturn, nil, false)
	case tools.V14.IsSupported(height):
		return parseGeneric[*miner14.GetPeerIDReturn, *miner14.GetPeerIDReturn](rawReturn, nil, false)
	case tools.V13.IsSupported(height):
		return parseGeneric[*miner13.GetPeerIDReturn, *miner13.GetPeerIDReturn](rawReturn, nil, false)
	case tools.V12.IsSupported(height):
		return parseGeneric[*miner12.GetPeerIDReturn, *miner12.GetPeerIDReturn](rawReturn, nil, false)
	case tools.V11.IsSupported(height):
		return parseGeneric[*miner11.GetPeerIDReturn, *miner11.GetPeerIDReturn](rawReturn, nil, false)
	case tools.V10.IsSupported(height):
		return parseGeneric[*miner10.GetPeerIDReturn, *miner10.GetPeerIDReturn](rawReturn, nil, false)
	case tools.V9.IsSupported(height):
		return parseGeneric[*miner9.GetPeerIDReturn, *miner9.GetPeerIDReturn](rawReturn, nil, false)
	case tools.V8.IsSupported(height):
		return parseGeneric[*miner8.GetPeerIDReturn, *miner8.GetPeerIDReturn](rawReturn, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func GetMultiaddrs(height int64, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		return parseGeneric[*miner15.GetMultiAddrsReturn, *miner15.GetMultiAddrsReturn](rawReturn, nil, false)
	case tools.V14.IsSupported(height):
		return parseGeneric[*miner14.GetMultiAddrsReturn, *miner14.GetMultiAddrsReturn](rawReturn, nil, false)
	case tools.V13.IsSupported(height):
		return parseGeneric[*miner13.GetMultiAddrsReturn, *miner13.GetMultiAddrsReturn](rawReturn, nil, false)
	case tools.V12.IsSupported(height):
		return parseGeneric[*miner12.GetMultiAddrsReturn, *miner12.GetMultiAddrsReturn](rawReturn, nil, false)
	case tools.V11.IsSupported(height):
		return parseGeneric[*miner11.GetMultiAddrsReturn, *miner11.GetMultiAddrsReturn](rawReturn, nil, false)
	case tools.V10.IsSupported(height):
		return parseGeneric[*miner10.GetMultiAddrsReturn, *miner10.GetMultiAddrsReturn](rawReturn, nil, false)
	case tools.V9.IsSupported(height):
		return parseGeneric[*miner9.GetMultiAddrsReturn, *miner9.GetMultiAddrsReturn](rawReturn, nil, false)
	case tools.V8.IsSupported(height):
		return parseGeneric[*miner8.GetMultiAddrsReturn, *miner8.GetMultiAddrsReturn](rawReturn, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}
