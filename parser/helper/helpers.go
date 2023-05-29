package helper

import (
	"errors"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/builtin/v11/account"
	"github.com/filecoin-project/go-state-types/builtin/v11/cron"
	"github.com/filecoin-project/go-state-types/builtin/v11/datacap"
	"github.com/filecoin-project/go-state-types/builtin/v11/eam"
	"github.com/filecoin-project/go-state-types/builtin/v11/evm"
	filInit "github.com/filecoin-project/go-state-types/builtin/v11/init"
	"github.com/filecoin-project/go-state-types/builtin/v11/market"
	"github.com/filecoin-project/go-state-types/builtin/v11/miner"
	"github.com/filecoin-project/go-state-types/builtin/v11/multisig"
	"github.com/filecoin-project/go-state-types/builtin/v11/paych"
	"github.com/filecoin-project/go-state-types/builtin/v11/power"
	"github.com/filecoin-project/go-state-types/builtin/v11/reward"
	"github.com/filecoin-project/go-state-types/builtin/v11/verifreg"
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	rosettaFilecoinLib "github.com/zondax/rosetta-filecoin-lib"
	"github.com/zondax/rosetta-filecoin-lib/actors"
	"go.uber.org/zap"

	"github.com/zondax/fil-parser/database"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
)

var allMethods = map[string]map[abi.MethodNum]builtin.MethodMeta{
	manifest.InitKey:       filInit.Methods,
	manifest.CronKey:       cron.Methods,
	manifest.AccountKey:    account.Methods,
	manifest.PowerKey:      power.Methods,
	manifest.MinerKey:      miner.Methods,
	manifest.MarketKey:     market.Methods,
	manifest.PaychKey:      paych.Methods,
	manifest.MultisigKey:   multisig.Methods,
	manifest.RewardKey:     reward.Methods,
	manifest.VerifregKey:   verifreg.Methods,
	manifest.EvmKey:        evm.Methods,
	manifest.EamKey:        eam.Methods,
	manifest.DatacapKey:    datacap.Methods,
	manifest.EthAccountKey: evm.Methods, // investigate this bafy2bzacebj3i5ehw2w6veowqisj2ag4wpp25glmmfsvejbwjj2e7axavonm6
}

type Helper struct {
	lib *rosettaFilecoinLib.RosettaConstructionFilecoin
}

func NewHelper(lib *rosettaFilecoinLib.RosettaConstructionFilecoin) *Helper {
	return &Helper{lib: lib}
}

func (h *Helper) GetActorAddressInfo(add address.Address, height int64, key filTypes.TipSetKey) *types.AddressInfo {
	var err error
	addInfo := &types.AddressInfo{}
	addInfo.Robust, err = database.ActorsDB.GetRobustAddress(add)
	if err != nil {
		zap.S().Errorf("could not get robust address for %s. Err: %v", add.String(), err)
	}

	addInfo.Short, err = database.ActorsDB.GetShortAddress(add)
	if err != nil {
		zap.S().Errorf("could not get short address for %s. Err: %v", add.String(), err)
	}

	addInfo.ActorCid, err = database.ActorsDB.GetActorCode(add, height, key)
	if err != nil {
		zap.S().Errorf("could not get actor code from address. Err:", err)
	} else {
		addInfo.ActorType, _ = h.lib.BuiltinActors.GetActorNameFromCid(addInfo.ActorCid)
	}

	return addInfo
}

func (h *Helper) getActorNameFromAddress(address address.Address, height int64, key filTypes.TipSetKey) string {
	var actorCode cid.Cid
	// Search for actor in cache
	var err error
	actorCode, err = database.ActorsDB.GetActorCode(address, height, key)
	if err != nil {
		return actors.UnknownStr
	}

	actorName, err := h.lib.BuiltinActors.GetActorNameFromCid(actorCode)
	if err != nil {
		return actors.UnknownStr
	}

	return actorName
}

func (h *Helper) GetMethodName(msg *parser.LotusMessage, height int64, key filTypes.TipSetKey) (string, error) {

	if msg == nil {
		return "", errors.New("malformed value")
	}

	// Shortcut 1 - Method "0" corresponds to "MethodSend"
	if msg.Method == 0 {
		return parser.MethodSend, nil
	}

	// Shortcut 2 - Method "1" corresponds to "MethodConstructor"
	if msg.Method == 1 {
		return parser.MethodConstructor, nil
	}

	actorName := h.getActorNameFromAddress(msg.To, height, key)

	actorMethods, ok := allMethods[actorName]
	if !ok {
		return "", parser.ErrNotKnownActor
	}
	method, ok := actorMethods[msg.Method]
	if !ok {
		return parser.UnknownStr, nil
	}
	return method.Name, nil
}
