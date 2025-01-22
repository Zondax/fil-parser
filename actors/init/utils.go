package init

import (
	"encoding/base64"

	"github.com/filecoin-project/go-address"
	builtinInitv10 "github.com/filecoin-project/go-state-types/builtin/v10/init"
	builtinInitv11 "github.com/filecoin-project/go-state-types/builtin/v11/init"
	builtinInitv12 "github.com/filecoin-project/go-state-types/builtin/v12/init"
	builtinInitv13 "github.com/filecoin-project/go-state-types/builtin/v13/init"
	builtinInitv14 "github.com/filecoin-project/go-state-types/builtin/v14/init"
	builtinInitv15 "github.com/filecoin-project/go-state-types/builtin/v15/init"
	builtinInitv8 "github.com/filecoin-project/go-state-types/builtin/v8/init"
	builtinInitv9 "github.com/filecoin-project/go-state-types/builtin/v9/init"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
)

func execParams(params constructorParams) parser.ExecParams {
	setParams := func(codeCid cid.Cid, constructorParams []byte) parser.ExecParams {
		return parser.ExecParams{
			CodeCid:           codeCid.String(),
			ConstructorParams: base64.StdEncoding.EncodeToString(constructorParams),
		}
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
	}
	return parser.ExecParams{}
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
	}
	return &types.AddressInfo{}

}
