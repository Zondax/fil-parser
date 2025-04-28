package eam

import (
	"context"
	"encoding/hex"
	"fmt"
	"strconv"

	actor_tools "github.com/zondax/fil-parser/actors/v2/tools"
	"github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/golem/pkg/logger"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/ipfs/go-cid"

	eamv10 "github.com/filecoin-project/go-state-types/builtin/v10/eam"
	eamv11 "github.com/filecoin-project/go-state-types/builtin/v11/eam"
	eamv12 "github.com/filecoin-project/go-state-types/builtin/v12/eam"
	eamv13 "github.com/filecoin-project/go-state-types/builtin/v13/eam"
	eamv14 "github.com/filecoin-project/go-state-types/builtin/v14/eam"
	eamv15 "github.com/filecoin-project/go-state-types/builtin/v15/eam"
	eamv16 "github.com/filecoin-project/go-state-types/builtin/v16/eam"

	typegen "github.com/whyrusleeping/cbor-gen"

	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
)

type Eam struct {
	helper *helper.Helper
	logger *logger.Logger
}

func New(helper *helper.Helper, logger *logger.Logger) *Eam {
	return &Eam{
		logger: logger,
		helper: helper,
	}
}

func (e *Eam) Name() string {
	return manifest.EamKey
}

func (*Eam) StartNetworkHeight() int64 {
	return tools.V18.Height()
}

func (e *Eam) Methods(_ context.Context, network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error) {
	switch {
	// all legacy version
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V17)...):
		return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{}, fmt.Errorf("%w: %d", actor_tools.ErrUnsupportedHeight, height)
	case tools.V18.IsSupported(network, height):
		return eamv10.Methods, nil
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return eamv11.Methods, nil
	case tools.V21.IsSupported(network, height):
		return eamv12.Methods, nil
	case tools.V22.IsSupported(network, height):
		return eamv13.Methods, nil
	case tools.V23.IsSupported(network, height):
		return eamv14.Methods, nil
	case tools.V24.IsSupported(network, height):
		return eamv15.Methods, nil
	case tools.V25.IsSupported(network, height):
		return eamv16.Methods, nil
	default:
		return nil, fmt.Errorf("%w: %d", actor_tools.ErrUnsupportedHeight, height)
	}
}

func newEamCreate(r typegen.CBORUnmarshaler, msgCid cid.Cid) (*types.AddressInfo, parser.EamCreateReturn) {
	getReturnStruct := func(actorID uint64, robustAddress *address.Address, ethAddress string) (*types.AddressInfo, parser.EamCreateReturn) {
		createReturn := parser.EamCreateReturn{
			ActorId:       actorID,
			RobustAddress: robustAddress,
			EthAddress:    ethAddress,
		}

		return &types.AddressInfo{
			Short:         parser.FilPrefix + strconv.FormatUint(actorID, 10),
			Robust:        robustAddress.String(),
			EthAddress:    parser.EthPrefix + ethAddress,
			ActorType:     "evm",
			CreationTxCid: msgCid.String(),
		}, createReturn

	}
	switch v := r.(type) {
	case *eamv16.CreateReturn:
		return getReturnStruct(v.ActorID, v.RobustAddress, parser.EthPrefix+hex.EncodeToString(v.EthAddress[:]))
	case *eamv16.Create2Return:
		return getReturnStruct(v.ActorID, v.RobustAddress, parser.EthPrefix+hex.EncodeToString(v.EthAddress[:]))
	case *eamv16.CreateExternalReturn:
		return getReturnStruct(v.ActorID, v.RobustAddress, parser.EthPrefix+hex.EncodeToString(v.EthAddress[:]))

	case *eamv15.CreateReturn:
		return getReturnStruct(v.ActorID, v.RobustAddress, parser.EthPrefix+hex.EncodeToString(v.EthAddress[:]))
	case *eamv15.Create2Return:
		return getReturnStruct(v.ActorID, v.RobustAddress, parser.EthPrefix+hex.EncodeToString(v.EthAddress[:]))
	case *eamv15.CreateExternalReturn:
		return getReturnStruct(v.ActorID, v.RobustAddress, parser.EthPrefix+hex.EncodeToString(v.EthAddress[:]))

	case *eamv14.CreateReturn:
		return getReturnStruct(v.ActorID, v.RobustAddress, parser.EthPrefix+hex.EncodeToString(v.EthAddress[:]))
	case *eamv14.Create2Return:
		return getReturnStruct(v.ActorID, v.RobustAddress, parser.EthPrefix+hex.EncodeToString(v.EthAddress[:]))
	case *eamv14.CreateExternalReturn:
		return getReturnStruct(v.ActorID, v.RobustAddress, parser.EthPrefix+hex.EncodeToString(v.EthAddress[:]))

	case *eamv13.CreateReturn:
		return getReturnStruct(v.ActorID, v.RobustAddress, parser.EthPrefix+hex.EncodeToString(v.EthAddress[:]))
	case *eamv13.Create2Return:
		return getReturnStruct(v.ActorID, v.RobustAddress, parser.EthPrefix+hex.EncodeToString(v.EthAddress[:]))
	case *eamv13.CreateExternalReturn:
		return getReturnStruct(v.ActorID, v.RobustAddress, parser.EthPrefix+hex.EncodeToString(v.EthAddress[:]))

	case *eamv12.CreateReturn:
		return getReturnStruct(v.ActorID, v.RobustAddress, parser.EthPrefix+hex.EncodeToString(v.EthAddress[:]))
	case *eamv12.Create2Return:
		return getReturnStruct(v.ActorID, v.RobustAddress, parser.EthPrefix+hex.EncodeToString(v.EthAddress[:]))
	case *eamv12.CreateExternalReturn:
		return getReturnStruct(v.ActorID, v.RobustAddress, parser.EthPrefix+hex.EncodeToString(v.EthAddress[:]))

	case *eamv11.CreateReturn:
		return getReturnStruct(v.ActorID, v.RobustAddress, parser.EthPrefix+hex.EncodeToString(v.EthAddress[:]))
	case *eamv11.Create2Return:
		return getReturnStruct(v.ActorID, v.RobustAddress, parser.EthPrefix+hex.EncodeToString(v.EthAddress[:]))
	case *eamv11.CreateExternalReturn:
		return getReturnStruct(v.ActorID, v.RobustAddress, parser.EthPrefix+hex.EncodeToString(v.EthAddress[:]))

	case *eamv10.CreateReturn:
		return getReturnStruct(v.ActorID, v.RobustAddress, parser.EthPrefix+hex.EncodeToString(v.EthAddress[:]))
	case *eamv10.Create2Return:
		return getReturnStruct(v.ActorID, v.RobustAddress, parser.EthPrefix+hex.EncodeToString(v.EthAddress[:]))
	case *eamv10.CreateExternalReturn:
		return getReturnStruct(v.ActorID, v.RobustAddress, parser.EthPrefix+hex.EncodeToString(v.EthAddress[:]))

	default:
		return nil, parser.EamCreateReturn{}
	}
}

func validateEamReturn(ret typegen.CBORUnmarshaler) error {
	if ret == nil {
		return fmt.Errorf("input is nil")
	}

	checkAndSetAddress := func(addr **address.Address) error {
		if *addr == nil {
			emptyAdd, _ := address.NewFromString("")
			*addr = &emptyAdd
			return fmt.Errorf("RobustAddress field is nil. Replaced with empty address")
		}
		return nil
	}

	switch v := ret.(type) {
	case *eamv16.CreateReturn:
		return checkAndSetAddress(&v.RobustAddress)
	case *eamv16.Create2Return:
		return checkAndSetAddress(&v.RobustAddress)
	case *eamv16.CreateExternalReturn:
		return checkAndSetAddress(&v.RobustAddress)

	case *eamv15.CreateReturn:
		return checkAndSetAddress(&v.RobustAddress)
	case *eamv15.Create2Return:
		return checkAndSetAddress(&v.RobustAddress)
	case *eamv15.CreateExternalReturn:
		return checkAndSetAddress(&v.RobustAddress)

	case *eamv14.CreateReturn:
		return checkAndSetAddress(&v.RobustAddress)
	case *eamv14.Create2Return:
		return checkAndSetAddress(&v.RobustAddress)
	case *eamv14.CreateExternalReturn:
		return checkAndSetAddress(&v.RobustAddress)

	case *eamv13.CreateReturn:
		return checkAndSetAddress(&v.RobustAddress)
	case *eamv13.Create2Return:
		return checkAndSetAddress(&v.RobustAddress)
	case *eamv13.CreateExternalReturn:
		return checkAndSetAddress(&v.RobustAddress)

	case *eamv12.CreateReturn:
		return checkAndSetAddress(&v.RobustAddress)
	case *eamv12.Create2Return:
		return checkAndSetAddress(&v.RobustAddress)
	case *eamv12.CreateExternalReturn:
		return checkAndSetAddress(&v.RobustAddress)

	case *eamv11.CreateReturn:
		return checkAndSetAddress(&v.RobustAddress)
	case *eamv11.Create2Return:
		return checkAndSetAddress(&v.RobustAddress)
	case *eamv11.CreateExternalReturn:
		return checkAndSetAddress(&v.RobustAddress)

	case *eamv10.CreateReturn:
		return checkAndSetAddress(&v.RobustAddress)
	case *eamv10.Create2Return:
		return checkAndSetAddress(&v.RobustAddress)
	case *eamv10.CreateExternalReturn:
		return checkAndSetAddress(&v.RobustAddress)
	}

	return nil
}
