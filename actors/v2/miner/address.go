package miner

import (
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	miner10 "github.com/filecoin-project/go-state-types/builtin/v10/miner"
	miner11 "github.com/filecoin-project/go-state-types/builtin/v11/miner"
	miner12 "github.com/filecoin-project/go-state-types/builtin/v12/miner"
	miner13 "github.com/filecoin-project/go-state-types/builtin/v13/miner"
	miner14 "github.com/filecoin-project/go-state-types/builtin/v14/miner"
	miner15 "github.com/filecoin-project/go-state-types/builtin/v15/miner"
	miner16 "github.com/filecoin-project/go-state-types/builtin/v16/miner"
	miner8 "github.com/filecoin-project/go-state-types/builtin/v8/miner"
	miner9 "github.com/filecoin-project/go-state-types/builtin/v9/miner"

	legacyv1 "github.com/filecoin-project/specs-actors/actors/builtin/miner"
	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/miner"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/miner"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/miner"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/miner"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/miner"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/miner"

	cbg "github.com/whyrusleeping/cbor-gen"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

func changeMultiaddrsParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.ChangeMultiaddrsParams{},

		tools.V8.String(): &legacyv2.ChangeMultiaddrsParams{},
		tools.V9.String(): &legacyv2.ChangeMultiaddrsParams{},

		tools.V10.String(): &legacyv3.ChangeMultiaddrsParams{},
		tools.V11.String(): &legacyv3.ChangeMultiaddrsParams{},

		tools.V12.String(): &legacyv4.ChangeMultiaddrsParams{},
		tools.V13.String(): &legacyv5.ChangeMultiaddrsParams{},
		tools.V14.String(): &legacyv6.ChangeMultiaddrsParams{},
		tools.V15.String(): &legacyv7.ChangeMultiaddrsParams{},
		tools.V16.String(): &miner8.ChangeMultiaddrsParams{},
		tools.V17.String(): &miner9.ChangeMultiaddrsParams{},
		tools.V18.String(): &miner10.ChangeMultiaddrsParams{},

		tools.V19.String(): &miner11.ChangeMultiaddrsParams{},
		tools.V20.String(): &miner11.ChangeMultiaddrsParams{},

		tools.V21.String(): &miner12.ChangeMultiaddrsParams{},
		tools.V22.String(): &miner13.ChangeMultiaddrsParams{},
		tools.V23.String(): &miner14.ChangeMultiaddrsParams{},
		tools.V24.String(): &miner15.ChangeMultiaddrsParams{},
		tools.V25.String(): &miner16.ChangeMultiaddrsParams{},
	}
}

func changePeerIDParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.ChangePeerIDParams{},

		tools.V8.String(): &legacyv2.ChangePeerIDParams{},
		tools.V9.String(): &legacyv2.ChangePeerIDParams{},

		tools.V10.String(): &legacyv3.ChangePeerIDParams{},
		tools.V11.String(): &legacyv3.ChangePeerIDParams{},

		tools.V12.String(): &legacyv4.ChangePeerIDParams{},
		tools.V13.String(): &legacyv5.ChangePeerIDParams{},
		tools.V14.String(): &legacyv6.ChangePeerIDParams{},
		tools.V15.String(): &legacyv7.ChangePeerIDParams{},
		tools.V16.String(): &miner8.ChangePeerIDParams{},
		tools.V17.String(): &miner9.ChangePeerIDParams{},
		tools.V18.String(): &miner10.ChangePeerIDParams{},

		tools.V19.String(): &miner11.ChangePeerIDParams{},
		tools.V20.String(): &miner11.ChangePeerIDParams{},

		tools.V21.String(): &miner12.ChangePeerIDParams{},
		tools.V22.String(): &miner13.ChangePeerIDParams{},
		tools.V23.String(): &miner14.ChangePeerIDParams{},
		tools.V24.String(): &miner15.ChangePeerIDParams{},
		tools.V25.String(): &miner16.ChangePeerIDParams{},
	}
}

func changeWorkerAddressParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.ChangeWorkerAddressParams{},

		tools.V8.String(): &legacyv2.ChangeWorkerAddressParams{},
		tools.V9.String(): &legacyv2.ChangeWorkerAddressParams{},

		tools.V10.String(): &legacyv3.ChangeWorkerAddressParams{},
		tools.V11.String(): &legacyv3.ChangeWorkerAddressParams{},

		tools.V12.String(): &legacyv4.ChangeWorkerAddressParams{},
		tools.V13.String(): &legacyv5.ChangeWorkerAddressParams{},
		tools.V14.String(): &legacyv6.ChangeWorkerAddressParams{},
		tools.V15.String(): &legacyv7.ChangeWorkerAddressParams{},
		tools.V16.String(): &miner8.ChangeWorkerAddressParams{},
		tools.V17.String(): &miner9.ChangeWorkerAddressParams{},
		tools.V18.String(): &miner10.ChangeWorkerAddressParams{},

		tools.V19.String(): &miner11.ChangeWorkerAddressParams{},
		tools.V20.String(): &miner11.ChangeWorkerAddressParams{},

		tools.V21.String(): &miner12.ChangeWorkerAddressParams{},
		tools.V22.String(): &miner13.ChangeWorkerAddressParams{},
		tools.V23.String(): &miner14.ChangeWorkerAddressParams{},
		tools.V24.String(): &miner15.ChangeWorkerAddressParams{},
		tools.V25.String(): &miner16.ChangeWorkerAddressParams{},
	}
}

func isControllingAddressParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V18.String(): &miner10.IsControllingAddressParams{},

		tools.V19.String(): &miner11.IsControllingAddressParams{},
		tools.V20.String(): &miner11.IsControllingAddressParams{},

		tools.V21.String(): &miner12.IsControllingAddressParams{},
		tools.V22.String(): &miner13.IsControllingAddressParams{},
		tools.V23.String(): &miner14.IsControllingAddressParams{},
		tools.V24.String(): &miner15.IsControllingAddressParams{},
		tools.V25.String(): &miner16.IsControllingAddressParams{},
	}
}

func isControllingAddressReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V18.String(): new(miner10.IsControllingAddressReturn),

		tools.V19.String(): new(miner11.IsControllingAddressReturn),
		tools.V20.String(): new(miner11.IsControllingAddressReturn),

		tools.V21.String(): new(miner12.IsControllingAddressReturn),
		tools.V22.String(): new(miner13.IsControllingAddressReturn),
		tools.V23.String(): new(miner14.IsControllingAddressReturn),
		tools.V24.String(): new(miner15.IsControllingAddressReturn),
		tools.V25.String(): new(miner16.IsControllingAddressReturn),
	}
}

func getOwnerReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V18.String(): &miner10.GetOwnerReturn{},

		tools.V19.String(): &miner11.GetOwnerReturn{},
		tools.V20.String(): &miner11.GetOwnerReturn{},

		tools.V21.String(): &miner12.GetOwnerReturn{},
		tools.V22.String(): &miner13.GetOwnerReturn{},
		tools.V23.String(): &miner14.GetOwnerReturn{},
		tools.V24.String(): &miner15.GetOwnerReturn{},
		tools.V25.String(): &miner16.GetOwnerReturn{},
	}
}

func getPeerIDReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V18.String(): &miner10.GetPeerIDReturn{},

		tools.V19.String(): &miner11.GetPeerIDReturn{},
		tools.V20.String(): &miner11.GetPeerIDReturn{},

		tools.V21.String(): &miner12.GetPeerIDReturn{},
		tools.V22.String(): &miner13.GetPeerIDReturn{},
		tools.V23.String(): &miner14.GetPeerIDReturn{},
		tools.V24.String(): &miner15.GetPeerIDReturn{},
		tools.V25.String(): &miner16.GetPeerIDReturn{},
	}
}

func getMultiAddrsReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V18.String(): &miner10.GetMultiAddrsReturn{},

		tools.V19.String(): &miner11.GetMultiAddrsReturn{},
		tools.V20.String(): &miner11.GetMultiAddrsReturn{},

		tools.V21.String(): &miner12.GetMultiAddrsReturn{},
		tools.V22.String(): &miner13.GetMultiAddrsReturn{},
		tools.V23.String(): &miner14.GetMultiAddrsReturn{},
		tools.V24.String(): &miner15.GetMultiAddrsReturn{},
		tools.V25.String(): &miner16.GetMultiAddrsReturn{},
	}
}

func getControlAddressesReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.GetControlAddressesReturn{},

		tools.V8.String(): &legacyv2.GetControlAddressesReturn{},
		tools.V9.String(): &legacyv2.GetControlAddressesReturn{},

		tools.V10.String(): &legacyv3.GetControlAddressesReturn{},
		tools.V11.String(): &legacyv3.GetControlAddressesReturn{},

		tools.V12.String(): &legacyv4.GetControlAddressesReturn{},
		tools.V13.String(): &legacyv5.GetControlAddressesReturn{},
		tools.V14.String(): &legacyv6.GetControlAddressesReturn{},
		tools.V15.String(): &legacyv7.GetControlAddressesReturn{},
		tools.V16.String(): &miner8.GetControlAddressesReturn{},
		tools.V17.String(): &miner9.GetControlAddressesReturn{},
		tools.V18.String(): &miner10.GetControlAddressesReturn{},

		tools.V19.String(): &miner11.GetControlAddressesReturn{},
		tools.V20.String(): &miner11.GetControlAddressesReturn{},

		tools.V21.String(): &miner12.GetControlAddressesReturn{},
		tools.V22.String(): &miner13.GetControlAddressesReturn{},
		tools.V23.String(): &miner14.GetControlAddressesReturn{},
		tools.V24.String(): &miner15.GetControlAddressesReturn{},
		tools.V25.String(): &miner16.GetControlAddressesReturn{},
	}
}

func (*Miner) ChangeMultiaddrsExported(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := changeMultiaddrsParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parseGeneric(rawParams, nil, false, params, &abi.EmptyValue{}, parser.ParamsKey)
}

func (*Miner) ChangePeerIDExported(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := changePeerIDParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parseGeneric(rawParams, nil, false, params, &abi.EmptyValue{}, parser.ParamsKey)
}

func (*Miner) ChangeWorkerAddressExported(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := changeWorkerAddressParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parseGeneric(rawParams, nil, false, params, &abi.EmptyValue{}, parser.ParamsKey)
}

func (*Miner) ChangeOwnerAddressExported(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	return parseGeneric(rawParams, nil, false, &address.Address{}, &address.Address{}, parser.ParamsKey)
}

func (*Miner) IsControllingAddressExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := isControllingAddressParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := isControllingAddressReturn()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parseGeneric(rawParams, rawReturn, true, params, returnValue, parser.ParamsKey)
}

func (*Miner) GetOwnerExported(network string, height int64, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	returnValue, ok := getOwnerReturn()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parseGeneric(rawReturn, nil, false, returnValue, &abi.EmptyValue{}, parser.ReturnKey)
}

func (*Miner) GetPeerIDExported(network string, height int64, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	returnValue, ok := getPeerIDReturn()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parseGeneric(rawReturn, nil, false, returnValue, &abi.EmptyValue{}, parser.ReturnKey)
}

func (*Miner) GetMultiaddrsExported(network string, height int64, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	returnValue, ok := getMultiAddrsReturn()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parseGeneric(rawReturn, nil, false, returnValue, &abi.EmptyValue{}, parser.ReturnKey)
}

func (*Miner) ControlAddresses(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	returnValue, ok := getControlAddressesReturn()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parseControlReturn(rawParams, rawReturn, returnValue)
}
