package miner

import (
	"bytes"
	"encoding/base64"
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
	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/miner"
	"github.com/zondax/fil-parser/parser"
)

func parseGeneric[T minerParam, R minerReturn](rawParams, rawReturn []byte, customReturn bool, params T, r R) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, fmt.Errorf("error unmarshalling params: %w", err)
	}
	metadata[parser.ParamsKey] = params
	if !customReturn {
		return metadata, nil
	}
	if len(rawReturn) > 0 {
		reader = bytes.NewReader(rawReturn)
		err = r.UnmarshalCBOR(reader)
		if err != nil {
			return metadata, err
		}
		metadata[parser.ReturnKey] = r
	}
	return metadata, nil
}

func parseControlReturn[R minerReturn](rawParams, rawReturn []byte, controlReturn R) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	if rawParams != nil {
		metadata[parser.ParamsKey] = base64.StdEncoding.EncodeToString(rawParams)
	}
	reader := bytes.NewReader(rawReturn)
	err := controlReturn.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	controlAddress, err := getControlAddress(controlReturn)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = controlAddress
	return metadata, nil
}

func getControlAddress(controlReturn any) (parser.ControlAddress, error) {
	controlAddress := parser.ControlAddress{}
	setControlReturn := func(owner, worker string, controlAddrs []string) {
		controlAddress.Owner = owner
		controlAddress.Worker = worker
		controlAddress.ControlAddrs = controlAddrs
	}

	switch v := controlReturn.(type) {
	case *legacyv2.GetControlAddressesReturn:
		setControlReturn(v.Owner.String(), v.Worker.String(), getControlAddrs(v.ControlAddrs))
	// *legacyv2.GetControlAddressesReturn is the same type in v2,v3,v4,v5,v6,v7
	// case *legacyv3.GetControlAddressesReturn:
	// case *legacyv4.GetControlAddressesReturn:
	// case *legacyv5.GetControlAddressesReturn:
	// case *legacyv6.GetControlAddressesReturn:
	// case *legacyv7.GetControlAddressesReturn:
	case *miner8.GetControlAddressesReturn:
		setControlReturn(v.Owner.String(), v.Worker.String(), getControlAddrs(v.ControlAddrs))
	case *miner9.GetControlAddressesReturn:
		setControlReturn(v.Owner.String(), v.Worker.String(), getControlAddrs(v.ControlAddrs))
	case *miner10.GetControlAddressesReturn:
		setControlReturn(v.Owner.String(), v.Worker.String(), getControlAddrs(v.ControlAddrs))
	case *miner11.GetControlAddressesReturn:
		setControlReturn(v.Owner.String(), v.Worker.String(), getControlAddrs(v.ControlAddrs))
	case *miner12.GetControlAddressesReturn:
		setControlReturn(v.Owner.String(), v.Worker.String(), getControlAddrs(v.ControlAddrs))
	case *miner13.GetControlAddressesReturn:
		setControlReturn(v.Owner.String(), v.Worker.String(), getControlAddrs(v.ControlAddrs))
	case *miner14.GetControlAddressesReturn:
		setControlReturn(v.Owner.String(), v.Worker.String(), getControlAddrs(v.ControlAddrs))
	case *miner15.GetControlAddressesReturn:
		setControlReturn(v.Owner.String(), v.Worker.String(), getControlAddrs(v.ControlAddrs))
	default:
		return controlAddress, fmt.Errorf("unsupported control return type: %T", v)
	}
	return controlAddress, nil

}

func getControlAddrs(addrs []address.Address) []string {
	r := make([]string, len(addrs))
	for i, addr := range addrs {
		r[i] = addr.String()
	}
	return r
}
