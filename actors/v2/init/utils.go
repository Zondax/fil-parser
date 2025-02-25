package init

import (
	"encoding/base64"
	"fmt"

	"github.com/filecoin-project/go-address"
	builtinInitv10 "github.com/filecoin-project/go-state-types/builtin/v10/init"
	builtinInitv11 "github.com/filecoin-project/go-state-types/builtin/v11/init"
	builtinInitv12 "github.com/filecoin-project/go-state-types/builtin/v12/init"
	builtinInitv13 "github.com/filecoin-project/go-state-types/builtin/v13/init"
	builtinInitv14 "github.com/filecoin-project/go-state-types/builtin/v14/init"
	builtinInitv15 "github.com/filecoin-project/go-state-types/builtin/v15/init"
	builtinInitv8 "github.com/filecoin-project/go-state-types/builtin/v8/init"
	builtinInitv9 "github.com/filecoin-project/go-state-types/builtin/v9/init"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"

	legacyInitv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/init"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
)

func execParams(params constructorParams) (cid.Cid, any, error) {
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
		// all previous legacy versions are the same exact type, adding to the switch case will cause a compile time error
	case *legacyInitv7.ExecParams:
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

func returnParams(msg *parser.LotusMessage, actorCID string, params execReturn) *types.AddressInfo {
	setReturn := func(idAddress, robustAddress address.Address) *types.AddressInfo {
		return &types.AddressInfo{
			Short:         idAddress.String(),
			Robust:        robustAddress.String(),
			ActorCid:      actorCID,
			CreationTxCid: msg.Cid.String(),
		}
	}
	switch v := params.(type) {
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
	}
	return &types.AddressInfo{}

}
