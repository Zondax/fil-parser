package eam

import (
	"context"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/golem/pkg/logger"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	"github.com/ipfs/go-cid"

	eamv10 "github.com/filecoin-project/go-state-types/builtin/v10/eam"
	eamv11 "github.com/filecoin-project/go-state-types/builtin/v11/eam"
	eamv12 "github.com/filecoin-project/go-state-types/builtin/v12/eam"
	eamv13 "github.com/filecoin-project/go-state-types/builtin/v13/eam"
	eamv14 "github.com/filecoin-project/go-state-types/builtin/v14/eam"
	eamv15 "github.com/filecoin-project/go-state-types/builtin/v15/eam"
	eamv16 "github.com/filecoin-project/go-state-types/builtin/v16/eam"

	typegen "github.com/whyrusleeping/cbor-gen"

	"github.com/zondax/fil-parser/actors"
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

var methods = map[string]map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
	tools.V18.String(): actors.CopyMethods(eamv10.Methods),
	tools.V19.String(): actors.CopyMethods(eamv11.Methods),
	tools.V20.String(): actors.CopyMethods(eamv11.Methods),
	tools.V21.String(): actors.CopyMethods(eamv12.Methods),
	tools.V22.String(): actors.CopyMethods(eamv13.Methods),
	tools.V23.String(): actors.CopyMethods(eamv14.Methods),
	tools.V24.String(): actors.CopyMethods(eamv15.Methods),
	tools.V25.String(): actors.CopyMethods(eamv16.Methods),
}

func (e *Eam) Methods(_ context.Context, network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error) {
	version := tools.VersionFromHeight(network, height)
	methods, ok := methods[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return methods, nil
}

func (e *Eam) newEamCreate(r typegen.CBORUnmarshaler, msgCid cid.Cid) (string, *types.AddressInfo, parser.EamCreateReturn, error) {
	getReturnStruct := func(actorID uint64, robustAddress *address.Address, ethAddress string) (string, *types.AddressInfo, parser.EamCreateReturn, error) {
		createReturn := parser.EamCreateReturn{
			ActorId:       actorID,
			RobustAddress: robustAddress,
			EthAddress:    ethAddress,
		}

		var ethHashStr string
		var robustAddressStr string
		ethHash, err := ethtypes.EthHashFromCid(msgCid)
		if err != nil {
			e.logger.Warnf("error getting ethHash msgCid: %s actorID: %d ethAddress: %s", msgCid.String(), actorID, ethAddress)
		} else {
			ethHashStr = ethHash.String()
		}

		if robustAddress != nil {
			if robustAddress.Empty() {
				robustAddressStr = ""
			} else {
				robustAddressStr = robustAddress.String()
			}
		}

		return ethHashStr, &types.AddressInfo{
			Short:         parser.FilPrefix + strconv.FormatUint(actorID, 10),
			Robust:        robustAddressStr,
			EthAddress:    parser.EthPrefix + ethAddress,
			ActorType:     "evm",
			CreationTxCid: msgCid.String(),
		}, createReturn, nil

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
		return "", nil, parser.EamCreateReturn{}, fmt.Errorf("invalid create return type")
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
			// the robust address can be nil
			// calibration: bafy2bzacedinpapwbjgevxivauhnqlunr6m2hsckgz2mmxdf67smxq4x47jn6
			return nil
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
