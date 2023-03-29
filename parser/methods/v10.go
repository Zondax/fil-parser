package methods

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/builtin/v10/account"
	"github.com/filecoin-project/go-state-types/builtin/v10/cron"
	"github.com/filecoin-project/go-state-types/builtin/v10/datacap"
	"github.com/filecoin-project/go-state-types/builtin/v10/eam"
	"github.com/filecoin-project/go-state-types/builtin/v10/evm"
	filInit "github.com/filecoin-project/go-state-types/builtin/v10/init"
	"github.com/filecoin-project/go-state-types/builtin/v10/market"
	"github.com/filecoin-project/go-state-types/builtin/v10/miner"
	"github.com/filecoin-project/go-state-types/builtin/v10/multisig"
	"github.com/filecoin-project/go-state-types/builtin/v10/paych"
	"github.com/filecoin-project/go-state-types/builtin/v10/power"
	"github.com/filecoin-project/go-state-types/builtin/v10/reward"
	"github.com/filecoin-project/go-state-types/builtin/v10/verifreg"
	"github.com/filecoin-project/go-state-types/manifest"
)

var v10methods = map[string]map[abi.MethodNum]builtin.MethodMeta{
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
	manifest.EthAccountKey: evm.Methods,
}

func V10Methods() map[string]map[abi.MethodNum]builtin.MethodMeta {
	return v10methods
}
