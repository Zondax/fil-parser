package eam

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"strconv"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	eamv10 "github.com/filecoin-project/go-state-types/builtin/v10/eam"
	eamv11 "github.com/filecoin-project/go-state-types/builtin/v11/eam"
	eamv12 "github.com/filecoin-project/go-state-types/builtin/v12/eam"
	eamv13 "github.com/filecoin-project/go-state-types/builtin/v13/eam"
	eamv14 "github.com/filecoin-project/go-state-types/builtin/v14/eam"
	eamv15 "github.com/filecoin-project/go-state-types/builtin/v15/eam"
	eamv9 "github.com/filecoin-project/go-state-types/builtin/v9/eam"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
)

type createReturn interface {
	UnmarshalCBOR(io.Reader) error
}

type createReturnStruct[T any] struct {
	CreateReturn T
}

func parseEamReturn[R createReturn](rawReturn []byte) (R, error) {
	var cr R

	reader := bytes.NewReader(rawReturn)
	err := cr.UnmarshalCBOR(reader)
	if err != nil {
		return cr, err
	}

	err = validateEamReturn(cr)
	if err != nil {
		rawString := hex.EncodeToString(rawReturn)
		return cr, fmt.Errorf("[parseEamReturn]- Detected invalid return bytes: %s. Raw: %s", err, rawString)
	}

	return cr, nil
}

func ParseEamReturn(height int64, rawReturn []byte) (any, error) {
	switch height {
	case 15:
		return parseEamReturn[*eamv15.CreateReturn](rawReturn)
	case 14:
		return parseEamReturn[*eamv14.CreateReturn](rawReturn)
	case 13:
		return parseEamReturn[*eamv13.CreateReturn](rawReturn)
	case 12:
		return parseEamReturn[*eamv12.CreateReturn](rawReturn)
	case 11:
		return parseEamReturn[*eamv11.CreateReturn](rawReturn)
	case 10:
		return parseEamReturn[*eamv10.CreateReturn](rawReturn)
	case 9:
		return parseEamReturn[*eamv9.CreateReturn](rawReturn)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func ParseCreateExternal(height int64, rawParams, rawReturn []byte, msgCid cid.Cid) (map[string]interface{}, *types.AddressInfo, error) {
	external := true
	switch height {
	case 15:
		return parseCreate[*eamv15.CreateExternalReturn](rawParams, rawReturn, msgCid, external)
	case 14:
		return parseCreate[*eamv14.CreateExternalReturn](rawParams, rawReturn, msgCid, external)
	case 13:
		return parseCreate[*eamv13.CreateExternalReturn](rawParams, rawReturn, msgCid, external)
	case 12:
		return parseCreate[*eamv12.CreateExternalReturn](rawParams, rawReturn, msgCid, external)
	case 11:
		return parseCreate[*eamv11.CreateExternalReturn](rawParams, rawReturn, msgCid, external)
	case 10:
		return parseCreate[*eamv10.CreateExternalReturn](rawParams, rawReturn, msgCid, external)
	case 9:
		return parseCreate[*eamv9.CreateExternalReturn](rawParams, rawReturn, msgCid, external)
	}
	return nil, nil, fmt.Errorf("unsupported height: %d", height)
}

func ParseCreate(height int64, rawParams, rawReturn []byte, msgCid cid.Cid) (map[string]interface{}, *types.AddressInfo, error) {
	external := false
	switch height {
	case 15:
		return parseCreate[*eamv15.CreateReturn](rawParams, rawReturn, msgCid, external)
	case 14:
		return parseCreate[*eamv14.CreateReturn](rawParams, rawReturn, msgCid, external)
	case 13:
		return parseCreate[*eamv13.CreateReturn](rawParams, rawReturn, msgCid, external)
	case 12:
		return parseCreate[*eamv12.CreateReturn](rawParams, rawReturn, msgCid, external)
	case 11:
		return parseCreate[*eamv11.CreateReturn](rawParams, rawReturn, msgCid, external)
	case 10:
		return parseCreate[*eamv10.CreateReturn](rawParams, rawReturn, msgCid, external)
	case 9:
		return parseCreate[*eamv9.CreateReturn](rawParams, rawReturn, msgCid, external)
	}
	return nil, nil, fmt.Errorf("unsupported height: %d", height)
}

func ParseCreate2(height int64, rawParams, rawReturn []byte, msgCid cid.Cid) (map[string]interface{}, *types.AddressInfo, error) {
	external := false
	switch height {
	case 15:
		return parseCreate[*eamv15.Create2Return](rawParams, rawReturn, msgCid, external)
	case 14:
		return parseCreate[*eamv14.Create2Return](rawParams, rawReturn, msgCid, external)
	case 13:
		return parseCreate[*eamv13.Create2Return](rawParams, rawReturn, msgCid, external)
	case 12:
		return parseCreate[*eamv12.Create2Return](rawParams, rawReturn, msgCid, external)
	case 11:
		return parseCreate[*eamv11.Create2Return](rawParams, rawReturn, msgCid, external)
	case 10:
		return parseCreate[*eamv10.Create2Return](rawParams, rawReturn, msgCid, external)
	case 9:
		return parseCreate[*eamv9.Create2Return](rawParams, rawReturn, msgCid, external)
	}
	return nil, nil, fmt.Errorf("unsupported height: %d", height)
}

func newEamCreate(r createReturn) parser.EamCreateReturn {
	switch v := r.(type) {
	case *eamv15.CreateReturn:
		return parser.EamCreateReturn{
			ActorId:       v.ActorID,
			RobustAddress: v.RobustAddress,
			EthAddress:    parser.EthPrefix + hex.EncodeToString(v.EthAddress[:]),
		}
	case *eamv14.CreateReturn:
		return parser.EamCreateReturn{
			ActorId:       v.ActorID,
			RobustAddress: v.RobustAddress,
			EthAddress:    parser.EthPrefix + hex.EncodeToString(v.EthAddress[:]),
		}
	case *eamv13.CreateReturn:
		return parser.EamCreateReturn{
			ActorId:       v.ActorID,
			RobustAddress: v.RobustAddress,
			EthAddress:    parser.EthPrefix + hex.EncodeToString(v.EthAddress[:]),
		}
	case *eamv12.CreateReturn:
		return parser.EamCreateReturn{
			ActorId:       v.ActorID,
			RobustAddress: v.RobustAddress,
			EthAddress:    parser.EthPrefix + hex.EncodeToString(v.EthAddress[:]),
		}
	case *eamv11.CreateReturn:
		return parser.EamCreateReturn{
			ActorId:       v.ActorID,
			RobustAddress: v.RobustAddress,
			EthAddress:    parser.EthPrefix + hex.EncodeToString(v.EthAddress[:]),
		}
	case *eamv10.CreateReturn:
		return parser.EamCreateReturn{
			ActorId:       v.ActorID,
			RobustAddress: v.RobustAddress,
			EthAddress:    parser.EthPrefix + hex.EncodeToString(v.EthAddress[:]),
		}
	case *eamv9.CreateReturn:
		return parser.EamCreateReturn{
			ActorId:       v.ActorID,
			RobustAddress: v.RobustAddress,
			EthAddress:    parser.EthPrefix + hex.EncodeToString(v.EthAddress[:]),
		}
	default:
		return parser.EamCreateReturn{}
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
	case *eamv14.CreateReturn:
		return checkAndSetAddress(&v.RobustAddress)
	case *eamv13.CreateReturn:
		return checkAndSetAddress(&v.RobustAddress)
	case *eamv12.CreateReturn:
		return checkAndSetAddress(&v.RobustAddress)
	case *eamv11.CreateReturn:
		return checkAndSetAddress(&v.RobustAddress)
	case *eamv10.CreateReturn:
		return checkAndSetAddress(&v.RobustAddress)
	case *eamv9.CreateReturn:
		return checkAndSetAddress(&v.RobustAddress)
	}

	return nil
}

func parseCreate[T createReturn](rawParams, rawReturn []byte, msgCid cid.Cid, isExternal bool) (map[string]interface{}, *types.AddressInfo, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)

	if isExternal {
		metadata[parser.ParamsKey] = parser.EthPrefix + hex.EncodeToString(rawParams)
		var params abi.CborBytes
		if err := params.UnmarshalCBOR(reader); err != nil {
			return metadata, nil, fmt.Errorf("error deserializing rawParams: %s - hex data: %s", err.Error(), hex.EncodeToString(rawParams))
		}

		if reader.Len() == 0 { // This means that the reader has processed all the bytes
			metadata[parser.ParamsKey] = parser.EthPrefix + hex.EncodeToString(params)
		}
	} else {
		var params T
		err := params.UnmarshalCBOR(reader)
		if err != nil {
			return metadata, nil, err
		}
	}

	createReturn, err := parseEamReturn[T](rawReturn)
	if err != nil {
		return metadata, nil, err
	}

	metadata[parser.ReturnKey] = newEamCreate(createReturn)

	ethHash, err := ethtypes.EthHashFromCid(msgCid)
	if err != nil {
		return metadata, nil, err
	}
	metadata[parser.EthHashKey] = ethHash.String()

	r := newEamCreate(createReturn)
	createdEvmActor := &types.AddressInfo{
		Short:         parser.FilPrefix + strconv.FormatUint(r.ActorId, 10),
		Robust:        r.RobustAddress.String(),
		EthAddress:    parser.EthPrefix + r.EthAddress,
		ActorType:     "evm",
		CreationTxCid: msgCid.String(),
	}
	return metadata, createdEvmActor, nil
}
