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

	legacyv1 "github.com/filecoin-project/specs-actors/actors/builtin/miner"
	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/miner"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/miner"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/miner"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/miner"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/miner"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/miner"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

func (*Miner) ChangeMultiaddrsExported(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.ChangeMultiaddrsParams{}, &miner15.ChangeMultiaddrsParams{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.ChangeMultiaddrsParams{}, &miner14.ChangeMultiaddrsParams{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.ChangeMultiaddrsParams{}, &miner13.ChangeMultiaddrsParams{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.ChangeMultiaddrsParams{}, &miner12.ChangeMultiaddrsParams{}, parser.ParamsKey)

	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.ChangeMultiaddrsParams{}, &miner11.ChangeMultiaddrsParams{}, parser.ParamsKey)

	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.ChangeMultiaddrsParams{}, &miner10.ChangeMultiaddrsParams{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.ChangeMultiaddrsParams{}, &miner9.ChangeMultiaddrsParams{}, parser.ParamsKey)
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner8.ChangeMultiaddrsParams{}, &miner8.ChangeMultiaddrsParams{}, parser.ParamsKey)
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv7.ChangeMultiaddrsParams{}, &legacyv7.ChangeMultiaddrsParams{}, parser.ParamsKey)
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv6.ChangeMultiaddrsParams{}, &legacyv6.ChangeMultiaddrsParams{}, parser.ParamsKey)
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv5.ChangeMultiaddrsParams{}, &legacyv5.ChangeMultiaddrsParams{}, parser.ParamsKey)
	case tools.V12.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv4.ChangeMultiaddrsParams{}, &legacyv4.ChangeMultiaddrsParams{}, parser.ParamsKey)

	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parseGeneric(rawParams, nil, false, &legacyv3.ChangeMultiaddrsParams{}, &legacyv3.ChangeMultiaddrsParams{}, parser.ParamsKey)

	case tools.AnyIsSupported(network, height, tools.V8, tools.V9):
		return parseGeneric(rawParams, nil, false, &legacyv2.ChangeMultiaddrsParams{}, &legacyv2.ChangeMultiaddrsParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return parseGeneric(rawParams, nil, false, &legacyv1.ChangeMultiaddrsParams{}, &legacyv1.ChangeMultiaddrsParams{}, parser.ParamsKey)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) ChangePeerIDExported(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.ChangePeerIDParams{}, &miner15.ChangePeerIDParams{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.ChangePeerIDParams{}, &miner14.ChangePeerIDParams{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.ChangePeerIDParams{}, &miner13.ChangePeerIDParams{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.ChangePeerIDParams{}, &miner12.ChangePeerIDParams{}, parser.ParamsKey)

	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.ChangePeerIDParams{}, &miner11.ChangePeerIDParams{}, parser.ParamsKey)

	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.ChangePeerIDParams{}, &miner10.ChangePeerIDParams{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.ChangePeerIDParams{}, &miner9.ChangePeerIDParams{}, parser.ParamsKey)
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner8.ChangePeerIDParams{}, &miner8.ChangePeerIDParams{}, parser.ParamsKey)
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv7.ChangePeerIDParams{}, &legacyv7.ChangePeerIDParams{}, parser.ParamsKey)
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv6.ChangePeerIDParams{}, &legacyv6.ChangePeerIDParams{}, parser.ParamsKey)
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv5.ChangePeerIDParams{}, &legacyv5.ChangePeerIDParams{}, parser.ParamsKey)
	case tools.V12.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv4.ChangePeerIDParams{}, &legacyv4.ChangePeerIDParams{}, parser.ParamsKey)

	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parseGeneric(rawParams, nil, false, &legacyv3.ChangePeerIDParams{}, &legacyv3.ChangePeerIDParams{}, parser.ParamsKey)

	case tools.AnyIsSupported(network, height, tools.V8, tools.V9):
		return parseGeneric(rawParams, nil, false, &legacyv2.ChangePeerIDParams{}, &legacyv2.ChangePeerIDParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return parseGeneric(rawParams, nil, false, &legacyv1.ChangePeerIDParams{}, &legacyv1.ChangePeerIDParams{}, parser.ParamsKey)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) ChangeWorkerAddressExported(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.ChangeWorkerAddressParams{}, &miner15.ChangeWorkerAddressParams{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.ChangeWorkerAddressParams{}, &miner14.ChangeWorkerAddressParams{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.ChangeWorkerAddressParams{}, &miner13.ChangeWorkerAddressParams{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.ChangeWorkerAddressParams{}, &miner12.ChangeWorkerAddressParams{}, parser.ParamsKey)

	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.ChangeWorkerAddressParams{}, &miner11.ChangeWorkerAddressParams{}, parser.ParamsKey)

	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.ChangeWorkerAddressParams{}, &miner10.ChangeWorkerAddressParams{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.ChangeWorkerAddressParams{}, &miner9.ChangeWorkerAddressParams{}, parser.ParamsKey)
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner8.ChangeWorkerAddressParams{}, &miner8.ChangeWorkerAddressParams{}, parser.ParamsKey)
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv7.ChangeWorkerAddressParams{}, &legacyv7.ChangeWorkerAddressParams{}, parser.ParamsKey)
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv6.ChangeWorkerAddressParams{}, &legacyv6.ChangeWorkerAddressParams{}, parser.ParamsKey)
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv5.ChangeWorkerAddressParams{}, &legacyv5.ChangeWorkerAddressParams{}, parser.ParamsKey)
	case tools.V12.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv4.ChangeWorkerAddressParams{}, &legacyv4.ChangeWorkerAddressParams{}, parser.ParamsKey)

	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parseGeneric(rawParams, nil, false, &legacyv3.ChangeWorkerAddressParams{}, &legacyv3.ChangeWorkerAddressParams{}, parser.ParamsKey)

	case tools.AnyIsSupported(network, height, tools.V8, tools.V9):
		return parseGeneric(rawParams, nil, false, &legacyv2.ChangeWorkerAddressParams{}, &legacyv2.ChangeWorkerAddressParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return parseGeneric(rawParams, nil, false, &legacyv1.ChangeWorkerAddressParams{}, &legacyv1.ChangeWorkerAddressParams{}, parser.ParamsKey)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) ChangeOwnerAddressExported(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	return parseGeneric(rawParams, nil, false, &address.Address{}, &address.Address{}, parser.ParamsKey)
}

func (*Miner) IsControllingAddressExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		var returnValue miner15.IsControllingAddressReturn
		return parseGeneric(rawParams, rawReturn, true, &miner15.IsControllingAddressParams{}, &returnValue, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		var returnValue miner14.IsControllingAddressReturn
		return parseGeneric(rawParams, rawReturn, true, &miner14.IsControllingAddressParams{}, &returnValue, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		var returnValue miner13.IsControllingAddressReturn
		return parseGeneric(rawParams, rawReturn, true, &miner13.IsControllingAddressParams{}, &returnValue, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		var returnValue miner12.IsControllingAddressReturn
		return parseGeneric(rawParams, rawReturn, true, &miner12.IsControllingAddressParams{}, &returnValue, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V20, tools.V19):
		var returnValue miner11.IsControllingAddressReturn
		return parseGeneric(rawParams, rawReturn, true, &miner11.IsControllingAddressParams{}, &returnValue, parser.ParamsKey)
	case tools.V18.IsSupported(network, height):
		var returnValue miner10.IsControllingAddressReturn
		return parseGeneric(rawParams, rawReturn, true, &miner10.IsControllingAddressParams{}, &returnValue, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V17)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) GetOwnerExported(network string, height int64, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawReturn, nil, false, &miner15.GetOwnerReturn{}, &miner15.GetOwnerReturn{}, parser.ReturnKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawReturn, nil, false, &miner14.GetOwnerReturn{}, &miner14.GetOwnerReturn{}, parser.ReturnKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawReturn, nil, false, &miner13.GetOwnerReturn{}, &miner13.GetOwnerReturn{}, parser.ReturnKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawReturn, nil, false, &miner12.GetOwnerReturn{}, &miner12.GetOwnerReturn{}, parser.ReturnKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawReturn, nil, false, &miner11.GetOwnerReturn{}, &miner11.GetOwnerReturn{}, parser.ReturnKey)
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawReturn, nil, false, &miner10.GetOwnerReturn{}, &miner10.GetOwnerReturn{}, parser.ReturnKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V17)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) GetPeerIDExported(network string, height int64, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawReturn, nil, false, &miner15.GetPeerIDReturn{}, &miner15.GetPeerIDReturn{}, parser.ReturnKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawReturn, nil, false, &miner14.GetPeerIDReturn{}, &miner14.GetPeerIDReturn{}, parser.ReturnKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawReturn, nil, false, &miner13.GetPeerIDReturn{}, &miner13.GetPeerIDReturn{}, parser.ReturnKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawReturn, nil, false, &miner12.GetPeerIDReturn{}, &miner12.GetPeerIDReturn{}, parser.ReturnKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawReturn, nil, false, &miner11.GetPeerIDReturn{}, &miner11.GetPeerIDReturn{}, parser.ReturnKey)
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawReturn, nil, false, &miner10.GetPeerIDReturn{}, &miner10.GetPeerIDReturn{}, parser.ReturnKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V17)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) GetMultiaddrsExported(network string, height int64, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawReturn, nil, false, &miner15.GetMultiAddrsReturn{}, &miner15.GetMultiAddrsReturn{}, parser.ReturnKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawReturn, nil, false, &miner14.GetMultiAddrsReturn{}, &miner14.GetMultiAddrsReturn{}, parser.ReturnKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawReturn, nil, false, &miner13.GetMultiAddrsReturn{}, &miner13.GetMultiAddrsReturn{}, parser.ReturnKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawReturn, nil, false, &miner12.GetMultiAddrsReturn{}, &miner12.GetMultiAddrsReturn{}, parser.ReturnKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawReturn, nil, false, &miner11.GetMultiAddrsReturn{}, &miner11.GetMultiAddrsReturn{}, parser.ReturnKey)
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawReturn, nil, false, &miner10.GetMultiAddrsReturn{}, &miner10.GetMultiAddrsReturn{}, parser.ReturnKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V17)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) ControlAddresses(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseControlReturn(rawParams, rawReturn, &miner15.GetControlAddressesReturn{})
	case tools.V23.IsSupported(network, height):
		return parseControlReturn(rawParams, rawReturn, &miner14.GetControlAddressesReturn{})
	case tools.V22.IsSupported(network, height):
		return parseControlReturn(rawParams, rawReturn, &miner13.GetControlAddressesReturn{})
	case tools.V21.IsSupported(network, height):
		return parseControlReturn(rawParams, rawReturn, &miner12.GetControlAddressesReturn{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseControlReturn(rawParams, rawReturn, &miner11.GetControlAddressesReturn{})
	case tools.V18.IsSupported(network, height):
		return parseControlReturn(rawParams, rawReturn, &miner10.GetControlAddressesReturn{})
	case tools.V17.IsSupported(network, height):
		return parseControlReturn(rawParams, rawReturn, &miner9.GetControlAddressesReturn{})
	case tools.V16.IsSupported(network, height):
		return parseControlReturn(rawParams, rawReturn, &miner8.GetControlAddressesReturn{})
	case tools.V15.IsSupported(network, height):
		return parseControlReturn(rawParams, rawReturn, &legacyv7.GetControlAddressesReturn{})
	case tools.V14.IsSupported(network, height):
		return parseControlReturn(rawParams, rawReturn, &legacyv6.GetControlAddressesReturn{})
	case tools.V13.IsSupported(network, height):
		return parseControlReturn(rawParams, rawReturn, &legacyv5.GetControlAddressesReturn{})
	case tools.V12.IsSupported(network, height):
		return parseControlReturn(rawParams, rawReturn, &legacyv4.GetControlAddressesReturn{})
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parseControlReturn(rawParams, rawReturn, &legacyv3.GetControlAddressesReturn{})
	case tools.AnyIsSupported(network, height, tools.V8, tools.V9):
		return parseControlReturn(rawParams, rawReturn, &legacyv2.GetControlAddressesReturn{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return parseControlReturn(rawParams, rawReturn, &legacyv1.GetControlAddressesReturn{})
	}

	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}
