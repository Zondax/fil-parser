package methods

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/builtin/v8/account"
	"github.com/filecoin-project/go-state-types/builtin/v8/cron"
	filInit "github.com/filecoin-project/go-state-types/builtin/v8/init"
	"github.com/filecoin-project/go-state-types/builtin/v8/market"
	"github.com/filecoin-project/go-state-types/builtin/v8/miner"
	"github.com/filecoin-project/go-state-types/builtin/v8/multisig"
	"github.com/filecoin-project/go-state-types/builtin/v8/paych"
	"github.com/filecoin-project/go-state-types/builtin/v8/power"
	"github.com/filecoin-project/go-state-types/builtin/v8/reward"
	"github.com/filecoin-project/go-state-types/builtin/v8/verifreg"
	"github.com/filecoin-project/go-state-types/manifest"
)

var v8methods = map[string]map[abi.MethodNum]builtin.MethodMeta{
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
}

func V8Methods() map[string]map[abi.MethodNum]builtin.MethodMeta {
	return v8methods
}
