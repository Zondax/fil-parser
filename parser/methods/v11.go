package methods

import (
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
)

var v11methods = map[string]map[abi.MethodNum]builtin.MethodMeta{
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

func V11Methods() map[string]map[abi.MethodNum]builtin.MethodMeta {
	return v11methods
}
