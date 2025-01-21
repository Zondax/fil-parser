package eam

import (
	"encoding/hex"
	"fmt"

	"github.com/filecoin-project/go-address"
	eamv10 "github.com/filecoin-project/go-state-types/builtin/v10/eam"
	eamv11 "github.com/filecoin-project/go-state-types/builtin/v11/eam"
	eamv12 "github.com/filecoin-project/go-state-types/builtin/v12/eam"
	eamv13 "github.com/filecoin-project/go-state-types/builtin/v13/eam"
	eamv14 "github.com/filecoin-project/go-state-types/builtin/v14/eam"
	eamv15 "github.com/filecoin-project/go-state-types/builtin/v15/eam"
	eamv9 "github.com/filecoin-project/go-state-types/builtin/v9/eam"
	"github.com/zondax/fil-parser/parser"
)

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
