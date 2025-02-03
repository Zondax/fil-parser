package eam

import (
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/filecoin-project/go-address"
	eamv10 "github.com/filecoin-project/go-state-types/builtin/v10/eam"
	eamv11 "github.com/filecoin-project/go-state-types/builtin/v11/eam"
	eamv12 "github.com/filecoin-project/go-state-types/builtin/v12/eam"
	eamv13 "github.com/filecoin-project/go-state-types/builtin/v13/eam"
	eamv14 "github.com/filecoin-project/go-state-types/builtin/v14/eam"
	eamv15 "github.com/filecoin-project/go-state-types/builtin/v15/eam"
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	"github.com/ipfs/go-cid"

	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
)

type Eam struct{}

func (e *Eam) Name() string {
	return manifest.EamKey
}

func newEamCreate(r createReturn, msgCid cid.Cid) (string, *types.AddressInfo, parser.EamCreateReturn, error) {
	getReturnStruct := func(actorID uint64, robustAddress *address.Address, ethAddress string) (string, *types.AddressInfo, parser.EamCreateReturn, error) {
		createReturn := parser.EamCreateReturn{
			ActorId:       actorID,
			RobustAddress: robustAddress,
			EthAddress:    ethAddress,
		}

		ethHash, err := ethtypes.EthHashFromCid(msgCid)
		if err != nil {
			return "", nil, createReturn, err
		}

		return ethHash.String(), &types.AddressInfo{
			Short:         parser.FilPrefix + strconv.FormatUint(actorID, 10),
			Robust:        robustAddress.String(),
			EthAddress:    parser.EthPrefix + ethAddress,
			ActorType:     "evm",
			CreationTxCid: msgCid.String(),
		}, createReturn, nil

	}
	switch v := r.(type) {
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

func validateEamReturn(ret createReturn) error {
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
