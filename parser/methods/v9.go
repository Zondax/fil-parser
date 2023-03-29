package methods

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/builtin/v9/account"
	"github.com/filecoin-project/go-state-types/builtin/v9/cron"
	"github.com/filecoin-project/go-state-types/builtin/v9/datacap"
	filInit "github.com/filecoin-project/go-state-types/builtin/v9/init"
	"github.com/filecoin-project/go-state-types/builtin/v9/market"
	"github.com/filecoin-project/go-state-types/builtin/v9/miner"
	"github.com/filecoin-project/go-state-types/builtin/v9/multisig"
	"github.com/filecoin-project/go-state-types/builtin/v9/paych"
	"github.com/filecoin-project/go-state-types/builtin/v9/power"
	"github.com/filecoin-project/go-state-types/builtin/v9/reward"
	"github.com/filecoin-project/go-state-types/builtin/v9/verifreg"
	"github.com/filecoin-project/go-state-types/manifest"
)

var v9methods = map[string]map[abi.MethodNum]builtin.MethodMeta{
	manifest.InitKey:     filInit.Methods,
	manifest.CronKey:     cron.Methods,
	manifest.AccountKey:  account.Methods,
	manifest.PowerKey:    power.Methods,
	manifest.MinerKey:    miner.Methods,
	manifest.MarketKey:   market.Methods,
	manifest.PaychKey:    paych.Methods,
	manifest.MultisigKey: multisig.Methods,
	manifest.RewardKey:   reward.Methods,
	manifest.VerifregKey: verifreg.Methods,
	manifest.DatacapKey:  datacap.Methods,
}

func V9Methods() map[string]map[abi.MethodNum]builtin.MethodMeta {
	return v9methods
}
