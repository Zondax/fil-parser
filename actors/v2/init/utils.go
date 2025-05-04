package init

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/filecoin-project/go-address"
	builtinInitv10 "github.com/filecoin-project/go-state-types/builtin/v10/init"
	builtinInitv11 "github.com/filecoin-project/go-state-types/builtin/v11/init"
	builtinInitv12 "github.com/filecoin-project/go-state-types/builtin/v12/init"
	builtinInitv13 "github.com/filecoin-project/go-state-types/builtin/v13/init"
	builtinInitv14 "github.com/filecoin-project/go-state-types/builtin/v14/init"
	builtinInitv15 "github.com/filecoin-project/go-state-types/builtin/v15/init"
	builtinInitv16 "github.com/filecoin-project/go-state-types/builtin/v16/init"
	builtinInitv8 "github.com/filecoin-project/go-state-types/builtin/v8/init"
	builtinInitv9 "github.com/filecoin-project/go-state-types/builtin/v9/init"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"

	legacyInitv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/init"

	"github.com/ipfs/go-cid"
	typegen "github.com/whyrusleeping/cbor-gen"
	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
)

func setExecParams(params typegen.CBORUnmarshaler) (cid.Cid, any, error) {
	setParams := func(codeCid cid.Cid, constructorParams []byte) (cid.Cid, any, error) {
		return cid.Undef, parser.ExecParams{
			CodeCid:           codeCid.String(),
			ConstructorParams: base64.StdEncoding.EncodeToString(constructorParams),
		}, nil
	}

	setExec4Params := func(codeCid cid.Cid, constructorParams []byte, subAddressBytes []byte) (cid.Cid, any, error) {
		var subAddressStr string
		// Exec4 parses f4 addresses which are currently only used by evm based actors until webassembly support is added
		// the else case is added to support future actors that may use f4 addresses
		if len(subAddressBytes) == ethtypes.EthAddressLength {
			var addr ethtypes.EthAddress
			copy(addr[:], subAddressBytes)
			subAddressStr = addr.String()
		} else {
			subAddress, err := address.NewFromBytes(subAddressBytes)
			if err != nil {
				return cid.Undef, nil, fmt.Errorf("error parsing subaddress for Exec4: %w", err)
			}
			subAddressStr = subAddress.String()
		}
		return cid.Undef, parser.Exec4Params{
			CodeCid:           codeCid.String(),
			ConstructorParams: base64.StdEncoding.EncodeToString(constructorParams),
			SubAddress:        subAddressStr,
		}, nil
	}

	switch v := params.(type) {
	case *builtinInitv16.ExecParams:
		return setParams(v.CodeCID, v.ConstructorParams)
	case *builtinInitv15.ExecParams:
		return setParams(v.CodeCID, v.ConstructorParams)
	case *builtinInitv14.ExecParams:
		return setParams(v.CodeCID, v.ConstructorParams)
	case *builtinInitv13.ExecParams:
		return setParams(v.CodeCID, v.ConstructorParams)
	case *builtinInitv12.ExecParams:
		return setParams(v.CodeCID, v.ConstructorParams)
	case *builtinInitv11.ExecParams:
		return setParams(v.CodeCID, v.ConstructorParams)
	case *builtinInitv10.ExecParams:
		return setParams(v.CodeCID, v.ConstructorParams)
	case *builtinInitv9.ExecParams:
		return setParams(v.CodeCID, v.ConstructorParams)
	case *builtinInitv8.ExecParams:
		return setParams(v.CodeCID, v.ConstructorParams)
	case *legacyInitv7.ExecParams:
		return setParams(v.CodeCID, v.ConstructorParams)

		// Code commented out as types are the same as v7, and compiler complains
		/*
			case *legacyInitv6.ExecParams:
				return setParams(v.CodeCID, v.ConstructorParams)
			case *legacyInitv5.ExecParams:
				return setParams(v.CodeCID, v.ConstructorParams)
			case *legacyInitv4.ExecParams:
				return setParams(v.CodeCID, v.ConstructorParams)
			case *legacyInitv3.ExecParams:
				return setParams(v.CodeCID, v.ConstructorParams)
			case *legacyInitv2.ExecParams:
				return setParams(v.CodeCID, v.ConstructorParams)
			case *legacyInitv1.ExecParams:
				return setParams(v.CodeCID, v.ConstructorParams)
		*/
	case *builtinInitv16.Exec4Params:
		return setParams(v.CodeCID, v.ConstructorParams)
	case *builtinInitv15.Exec4Params:
		return setExec4Params(v.CodeCID, v.ConstructorParams, v.SubAddress)
	case *builtinInitv14.Exec4Params:
		return setExec4Params(v.CodeCID, v.ConstructorParams, v.SubAddress)
	case *builtinInitv13.Exec4Params:
		return setExec4Params(v.CodeCID, v.ConstructorParams, v.SubAddress)
	case *builtinInitv12.Exec4Params:
		return setExec4Params(v.CodeCID, v.ConstructorParams, v.SubAddress)
	case *builtinInitv11.Exec4Params:
		return setExec4Params(v.CodeCID, v.ConstructorParams, v.SubAddress)
	case *builtinInitv10.Exec4Params:
		return setExec4Params(v.CodeCID, v.ConstructorParams, v.SubAddress)

	}

	return cid.Undef, nil, actors.ErrUnsupportedHeight
}

func setReturnParams(msg *parser.LotusMessage, actorCID string, params typegen.CBORUnmarshaler) *types.AddressInfo {
	setReturn := func(idAddress, robustAddress address.Address) *types.AddressInfo {
		return &types.AddressInfo{
			Short:         idAddress.String(),
			Robust:        robustAddress.String(),
			ActorCid:      actorCID,
			CreationTxCid: msg.Cid.String(),
		}
	}
	switch v := params.(type) {
	case *builtinInitv16.ExecReturn:
		return setReturn(v.IDAddress, v.RobustAddress)
	case *builtinInitv15.ExecReturn:
		return setReturn(v.IDAddress, v.RobustAddress)
	case *builtinInitv14.ExecReturn:
		return setReturn(v.IDAddress, v.RobustAddress)
	case *builtinInitv13.ExecReturn:
		return setReturn(v.IDAddress, v.RobustAddress)
	case *builtinInitv12.ExecReturn:
		return setReturn(v.IDAddress, v.RobustAddress)
	case *builtinInitv11.ExecReturn:
		return setReturn(v.IDAddress, v.RobustAddress)
	case *builtinInitv10.ExecReturn:
		return setReturn(v.IDAddress, v.RobustAddress)
	case *builtinInitv9.ExecReturn:
		return setReturn(v.IDAddress, v.RobustAddress)
	case *builtinInitv8.ExecReturn:
		return setReturn(v.IDAddress, v.RobustAddress)
	case *legacyInitv7.ExecReturn:
		return setReturn(v.IDAddress, v.RobustAddress)

		// Code commented out as types are the same as v1, and compiler complains
		/*
			case *legacyInitv6.ExecReturn:
				return setReturn(v.IDAddress, v.RobustAddress)
			case *legacyInitv5.ExecReturn:
				return setReturn(v.IDAddress, v.RobustAddress)
			case *legacyInitv4.ExecReturn:
				return setReturn(v.IDAddress, v.RobustAddress)
			case *legacyInitv3.ExecReturn:
				return setReturn(v.IDAddress, v.RobustAddress)
			case *legacyInitv2.ExecReturn:
				return setReturn(v.IDAddress, v.RobustAddress)
			case *legacyInitv1.ExecReturn:
				return setReturn(v.IDAddress, v.RobustAddress)
		*/
	}
	return &types.AddressInfo{}

}

func parseExecActor(actor string) string {
	s := strings.Split(actor, "/")
	if len(s) < 1 {
		return actor
	}
	return s[len(s)-1]
}
