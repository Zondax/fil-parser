package power

import (
	"github.com/filecoin-project/go-address"
	powerv10 "github.com/filecoin-project/go-state-types/builtin/v10/power"
	powerv11 "github.com/filecoin-project/go-state-types/builtin/v11/power"
	powerv12 "github.com/filecoin-project/go-state-types/builtin/v12/power"
	powerv13 "github.com/filecoin-project/go-state-types/builtin/v13/power"
	powerv14 "github.com/filecoin-project/go-state-types/builtin/v14/power"
	powerv15 "github.com/filecoin-project/go-state-types/builtin/v15/power"
	powerv16 "github.com/filecoin-project/go-state-types/builtin/v16/power"
	powerv8 "github.com/filecoin-project/go-state-types/builtin/v8/power"
	powerv9 "github.com/filecoin-project/go-state-types/builtin/v9/power"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/power"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
)

func getAddressInfo(r powerReturn, msg *parser.LotusMessage) *types.AddressInfo {
	createAddressInfo := func(idAddress, robustAddress address.Address, cid cid.Cid) *types.AddressInfo {
		return &types.AddressInfo{
			Short:         idAddress.String(),
			Robust:        robustAddress.String(),
			ActorType:     "miner",
			CreationTxCid: cid.String(),
		}
	}
	switch r := r.(type) {
	// Code commented out as types are the same as v1, and compiler complains
	/*
		case *legacyv1.CreateMinerReturn:
			return createAddressInfo(r.IDAddress, r.RobustAddress, msg.Cid)
		case *legacyv2.CreateMinerReturn:
			return createAddressInfo(r.IDAddress, r.RobustAddress, msg.Cid)
		case *legacyv3.CreateMinerReturn:
			return createAddressInfo(r.IDAddress, r.RobustAddress, msg.Cid)
		case *legacyv4.CreateMinerReturn:
			return createAddressInfo(r.IDAddress, r.RobustAddress, msg.Cid)
		case *legacyv5.CreateMinerReturn:
			return createAddressInfo(r.IDAddress, r.RobustAddress, msg.Cid)
		case *legacyv6.CreateMinerReturn:
			return createAddressInfo(r.IDAddress, r.RobustAddress, msg.Cid)
	*/
	case *legacyv7.CreateMinerReturn:
		return createAddressInfo(r.IDAddress, r.RobustAddress, msg.Cid)
	case *powerv8.CreateMinerReturn:
		return createAddressInfo(r.IDAddress, r.RobustAddress, msg.Cid)
	case *powerv9.CreateMinerReturn:
		return createAddressInfo(r.IDAddress, r.RobustAddress, msg.Cid)
	case *powerv10.CreateMinerReturn:
		return createAddressInfo(r.IDAddress, r.RobustAddress, msg.Cid)
	case *powerv11.CreateMinerReturn:
		return createAddressInfo(r.IDAddress, r.RobustAddress, msg.Cid)
	case *powerv12.CreateMinerReturn:
		return createAddressInfo(r.IDAddress, r.RobustAddress, msg.Cid)
	case *powerv13.CreateMinerReturn:
		return createAddressInfo(r.IDAddress, r.RobustAddress, msg.Cid)
	case *powerv14.CreateMinerReturn:
		return createAddressInfo(r.IDAddress, r.RobustAddress, msg.Cid)
	case *powerv15.CreateMinerReturn:
		return createAddressInfo(r.IDAddress, r.RobustAddress, msg.Cid)
	case *powerv16.CreateMinerReturn:
		return createAddressInfo(r.IDAddress, r.RobustAddress, msg.Cid)
	}
	return nil
}
