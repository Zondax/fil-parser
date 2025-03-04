package helper

import (

	// all multisig version imports
	multisigv10 "github.com/filecoin-project/go-state-types/builtin/v10/multisig"
	multisigv11 "github.com/filecoin-project/go-state-types/builtin/v11/multisig"
	multisigv12 "github.com/filecoin-project/go-state-types/builtin/v12/multisig"
	multisigv13 "github.com/filecoin-project/go-state-types/builtin/v13/multisig"
	multisigv14 "github.com/filecoin-project/go-state-types/builtin/v14/multisig"
	multisigv15 "github.com/filecoin-project/go-state-types/builtin/v15/multisig"
	multisigv8 "github.com/filecoin-project/go-state-types/builtin/v8/multisig"
	multisigv9 "github.com/filecoin-project/go-state-types/builtin/v9/multisig"

	// all account version imports
	accountv10 "github.com/filecoin-project/go-state-types/builtin/v10/account"
	accountv11 "github.com/filecoin-project/go-state-types/builtin/v11/account"
	accountv12 "github.com/filecoin-project/go-state-types/builtin/v12/account"
	accountv13 "github.com/filecoin-project/go-state-types/builtin/v13/account"
	accountv14 "github.com/filecoin-project/go-state-types/builtin/v14/account"
	accountv15 "github.com/filecoin-project/go-state-types/builtin/v15/account"
	accountv8 "github.com/filecoin-project/go-state-types/builtin/v8/account"
	accountv9 "github.com/filecoin-project/go-state-types/builtin/v9/account"

	// all cron version imports
	cronv10 "github.com/filecoin-project/go-state-types/builtin/v10/cron"
	cronv11 "github.com/filecoin-project/go-state-types/builtin/v11/cron"
	cronv12 "github.com/filecoin-project/go-state-types/builtin/v12/cron"
	cronv13 "github.com/filecoin-project/go-state-types/builtin/v13/cron"
	cronv14 "github.com/filecoin-project/go-state-types/builtin/v14/cron"
	cronv15 "github.com/filecoin-project/go-state-types/builtin/v15/cron"
	cronv8 "github.com/filecoin-project/go-state-types/builtin/v8/cron"
	cronv9 "github.com/filecoin-project/go-state-types/builtin/v9/cron"

	// all datacap version imports
	datacapv10 "github.com/filecoin-project/go-state-types/builtin/v10/datacap"
	datacapv11 "github.com/filecoin-project/go-state-types/builtin/v11/datacap"
	datacapv12 "github.com/filecoin-project/go-state-types/builtin/v12/datacap"
	datacapv13 "github.com/filecoin-project/go-state-types/builtin/v13/datacap"
	datacapv14 "github.com/filecoin-project/go-state-types/builtin/v14/datacap"
	datacapv15 "github.com/filecoin-project/go-state-types/builtin/v15/datacap"

	// all eam version imports
	eamv10 "github.com/filecoin-project/go-state-types/builtin/v10/eam"
	eamv11 "github.com/filecoin-project/go-state-types/builtin/v11/eam"
	eamv12 "github.com/filecoin-project/go-state-types/builtin/v12/eam"
	eamv13 "github.com/filecoin-project/go-state-types/builtin/v13/eam"
	eamv14 "github.com/filecoin-project/go-state-types/builtin/v14/eam"
	eamv15 "github.com/filecoin-project/go-state-types/builtin/v15/eam"

	// all ethaccount version imports

	// all evm version imports
	evmv10 "github.com/filecoin-project/go-state-types/builtin/v10/evm"
	evmv11 "github.com/filecoin-project/go-state-types/builtin/v11/evm"
	evmv12 "github.com/filecoin-project/go-state-types/builtin/v12/evm"
	evmv13 "github.com/filecoin-project/go-state-types/builtin/v13/evm"
	evmv14 "github.com/filecoin-project/go-state-types/builtin/v14/evm"
	evmv15 "github.com/filecoin-project/go-state-types/builtin/v15/evm"

	// all init version imports
	builtinInitv10 "github.com/filecoin-project/go-state-types/builtin/v10/init"
	builtinInitv11 "github.com/filecoin-project/go-state-types/builtin/v11/init"
	builtinInitv12 "github.com/filecoin-project/go-state-types/builtin/v12/init"
	builtinInitv13 "github.com/filecoin-project/go-state-types/builtin/v13/init"
	builtinInitv14 "github.com/filecoin-project/go-state-types/builtin/v14/init"
	builtinInitv15 "github.com/filecoin-project/go-state-types/builtin/v15/init"
	builtinInitv8 "github.com/filecoin-project/go-state-types/builtin/v8/init"
	builtinInitv9 "github.com/filecoin-project/go-state-types/builtin/v9/init"

	// all market version imports
	marketv10 "github.com/filecoin-project/go-state-types/builtin/v10/market"
	marketv11 "github.com/filecoin-project/go-state-types/builtin/v11/market"
	marketv12 "github.com/filecoin-project/go-state-types/builtin/v12/market"
	marketv13 "github.com/filecoin-project/go-state-types/builtin/v13/market"
	marketv14 "github.com/filecoin-project/go-state-types/builtin/v14/market"
	marketv15 "github.com/filecoin-project/go-state-types/builtin/v15/market"
	marketv8 "github.com/filecoin-project/go-state-types/builtin/v8/market"
	marketv9 "github.com/filecoin-project/go-state-types/builtin/v9/market"

	// all miner version imports
	minerv10 "github.com/filecoin-project/go-state-types/builtin/v10/miner"
	minerv11 "github.com/filecoin-project/go-state-types/builtin/v11/miner"
	minerv12 "github.com/filecoin-project/go-state-types/builtin/v12/miner"
	minerv13 "github.com/filecoin-project/go-state-types/builtin/v13/miner"
	minerv14 "github.com/filecoin-project/go-state-types/builtin/v14/miner"
	minerv15 "github.com/filecoin-project/go-state-types/builtin/v15/miner"
	minerv8 "github.com/filecoin-project/go-state-types/builtin/v8/miner"
	minerv9 "github.com/filecoin-project/go-state-types/builtin/v9/miner"

	// all paych version imports
	paychv10 "github.com/filecoin-project/go-state-types/builtin/v10/paych"
	paychv11 "github.com/filecoin-project/go-state-types/builtin/v11/paych"
	paychv12 "github.com/filecoin-project/go-state-types/builtin/v12/paych"
	paychv13 "github.com/filecoin-project/go-state-types/builtin/v13/paych"
	paychv14 "github.com/filecoin-project/go-state-types/builtin/v14/paych"
	paychv15 "github.com/filecoin-project/go-state-types/builtin/v15/paych"
	paychv8 "github.com/filecoin-project/go-state-types/builtin/v8/paych"
	paychv9 "github.com/filecoin-project/go-state-types/builtin/v9/paych"

	// all placeholder version imports

	// all power version imports
	powerv10 "github.com/filecoin-project/go-state-types/builtin/v10/power"
	powerv11 "github.com/filecoin-project/go-state-types/builtin/v11/power"
	powerv12 "github.com/filecoin-project/go-state-types/builtin/v12/power"
	powerv13 "github.com/filecoin-project/go-state-types/builtin/v13/power"
	powerv14 "github.com/filecoin-project/go-state-types/builtin/v14/power"
	powerv15 "github.com/filecoin-project/go-state-types/builtin/v15/power"
	powerv8 "github.com/filecoin-project/go-state-types/builtin/v8/power"
	powerv9 "github.com/filecoin-project/go-state-types/builtin/v9/power"

	// all reward version imports
	rewardv10 "github.com/filecoin-project/go-state-types/builtin/v10/reward"
	rewardv11 "github.com/filecoin-project/go-state-types/builtin/v11/reward"
	rewardv12 "github.com/filecoin-project/go-state-types/builtin/v12/reward"
	rewardv13 "github.com/filecoin-project/go-state-types/builtin/v13/reward"
	rewardv14 "github.com/filecoin-project/go-state-types/builtin/v14/reward"
	rewardv15 "github.com/filecoin-project/go-state-types/builtin/v15/reward"
	rewardv8 "github.com/filecoin-project/go-state-types/builtin/v8/reward"
	rewardv9 "github.com/filecoin-project/go-state-types/builtin/v9/reward"

	// all system version imports
	systemv10 "github.com/filecoin-project/go-state-types/builtin/v10/system"
	systemv11 "github.com/filecoin-project/go-state-types/builtin/v11/system"
	systemv12 "github.com/filecoin-project/go-state-types/builtin/v12/system"
	systemv13 "github.com/filecoin-project/go-state-types/builtin/v13/system"
	systemv14 "github.com/filecoin-project/go-state-types/builtin/v14/system"
	systemv15 "github.com/filecoin-project/go-state-types/builtin/v15/system"
	systemv8 "github.com/filecoin-project/go-state-types/builtin/v8/system"
	systemv9 "github.com/filecoin-project/go-state-types/builtin/v9/system"

	// all verifiedregistry version imports
	verifregv10 "github.com/filecoin-project/go-state-types/builtin/v10/verifreg"
	verifregv11 "github.com/filecoin-project/go-state-types/builtin/v11/verifreg"
	verifregv12 "github.com/filecoin-project/go-state-types/builtin/v12/verifreg"
	verifregv13 "github.com/filecoin-project/go-state-types/builtin/v13/verifreg"
	verifregv14 "github.com/filecoin-project/go-state-types/builtin/v14/verifreg"
	verifregv15 "github.com/filecoin-project/go-state-types/builtin/v15/verifreg"
	verifregv8 "github.com/filecoin-project/go-state-types/builtin/v8/verifreg"
	verifregv9 "github.com/filecoin-project/go-state-types/builtin/v9/verifreg"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/manifest"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

// EthAccount and Placeholder can receive tokens with Send and InvokeEVM methods
// We set evm.Methods instead of empty array of methods. Therefore, we will be able to understand
// this specific method (3844450837) - tx cid example: bafy2bzacedgmcvsp56ieciutvgwza2qpvz7pvbhhu4l5y5tdl35rwfnjn5buk
func allMethods(height int64, network string, actorName string) (map[abi.MethodNum]builtin.MethodMeta, error) {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V15)...):
		return legacyMethods(actorName)
	case tools.V16.IsSupported(network, height):
		return v8Methods(actorName)
	case tools.V17.IsSupported(network, height):
		return v9Methods(actorName)
	case tools.V18.IsSupported(network, height):
		return v10Methods(actorName)
	case tools.AnyIsSupported(network, height, tools.V20, tools.V19):
		return v11Methods(actorName)
	case tools.V21.IsSupported(network, height):
		return v12Methods(actorName)
	case tools.V22.IsSupported(network, height):
		return v13Methods(actorName)
	case tools.V23.IsSupported(network, height):
		return v14Methods(actorName)
	case tools.V24.IsSupported(network, height):
		return v15Methods(actorName)
	}
	return nil, actors.ErrUnsupportedHeight
}

func v8Methods(actorName string) (map[abi.MethodNum]builtin.MethodMeta, error) {
	var actorMethods map[abi.MethodNum]builtin.MethodMeta
	switch actorName {
	case manifest.AccountKey:
		actorMethods = accountv8.Methods
	case manifest.InitKey:
		actorMethods = builtinInitv8.Methods
	case manifest.CronKey:
		actorMethods = cronv8.Methods
	case manifest.RewardKey:
		actorMethods = rewardv8.Methods
	case manifest.MultisigKey:
		actorMethods = multisigv8.Methods
	case manifest.PaychKey:
		actorMethods = paychv8.Methods
	case manifest.MarketKey:
		actorMethods = marketv8.Methods
	case manifest.PowerKey:
		actorMethods = powerv8.Methods
	case manifest.MinerKey:
		actorMethods = minerv8.Methods
	case manifest.VerifregKey:
		actorMethods = verifregv8.Methods
	case manifest.SystemKey:
		actorMethods = systemv8.Methods
	default:
		return nil, parser.ErrNotKnownActor
	}
	return actorMethods, nil
}

func v9Methods(actorName string) (map[abi.MethodNum]builtin.MethodMeta, error) {
	var actorMethods map[abi.MethodNum]builtin.MethodMeta
	switch actorName {
	case manifest.AccountKey:
		actorMethods = accountv9.Methods
	case manifest.InitKey:
		actorMethods = builtinInitv9.Methods
	case manifest.CronKey:
		actorMethods = cronv9.Methods
	case manifest.RewardKey:
		actorMethods = rewardv9.Methods
	case manifest.MultisigKey:
		actorMethods = multisigv9.Methods
	case manifest.PaychKey:
		actorMethods = paychv9.Methods
	case manifest.MarketKey:
		actorMethods = marketv9.Methods
	case manifest.PowerKey:
		actorMethods = powerv9.Methods
	case manifest.MinerKey:
		actorMethods = minerv9.Methods
	case manifest.VerifregKey:
		actorMethods = verifregv9.Methods
	case manifest.SystemKey:
		actorMethods = systemv9.Methods
	default:
		return nil, parser.ErrNotKnownActor
	}
	return actorMethods, nil
}

func v10Methods(actorName string) (map[abi.MethodNum]builtin.MethodMeta, error) {
	var actorMethods map[abi.MethodNum]builtin.MethodMeta
	switch actorName {
	case manifest.AccountKey:
		actorMethods = accountv10.Methods
	case manifest.InitKey:
		actorMethods = builtinInitv10.Methods
	case manifest.CronKey:
		actorMethods = cronv10.Methods
	case manifest.RewardKey:
		actorMethods = rewardv10.Methods
	case manifest.MultisigKey:
		actorMethods = multisigv10.Methods
	case manifest.PaychKey:
		actorMethods = paychv10.Methods
	case manifest.MarketKey:
		actorMethods = marketv10.Methods
	case manifest.PowerKey:
		actorMethods = powerv10.Methods
	case manifest.MinerKey:
		actorMethods = minerv10.Methods
	case manifest.VerifregKey:
		actorMethods = verifregv10.Methods
	case manifest.EamKey:
		actorMethods = eamv10.Methods
	case manifest.EvmKey:
		actorMethods = evmv10.Methods
	case manifest.DatacapKey:
		actorMethods = datacapv10.Methods
	case manifest.EthAccountKey:
		actorMethods = evmv10.Methods
	case manifest.PlaceholderKey:
		actorMethods = evmv10.Methods
	case manifest.SystemKey:
		actorMethods = systemv10.Methods
	default:
		return nil, parser.ErrNotKnownActor
	}
	return actorMethods, nil
}

func v11Methods(actorName string) (map[abi.MethodNum]builtin.MethodMeta, error) {
	var actorMethods map[abi.MethodNum]builtin.MethodMeta
	switch actorName {
	case manifest.AccountKey:
		actorMethods = accountv11.Methods
	case manifest.InitKey:
		actorMethods = builtinInitv11.Methods
	case manifest.CronKey:
		actorMethods = cronv11.Methods
	case manifest.RewardKey:
		actorMethods = rewardv11.Methods
	case manifest.MultisigKey:
		actorMethods = multisigv11.Methods
	case manifest.PaychKey:
		actorMethods = paychv11.Methods
	case manifest.MarketKey:
		actorMethods = marketv11.Methods
	case manifest.PowerKey:
		actorMethods = powerv11.Methods
	case manifest.MinerKey:
		actorMethods = minerv11.Methods
	case manifest.VerifregKey:
		actorMethods = verifregv11.Methods
	case manifest.EamKey:
		actorMethods = eamv11.Methods
	case manifest.EvmKey:
		actorMethods = evmv11.Methods
	case manifest.DatacapKey:
		actorMethods = datacapv11.Methods
	case manifest.EthAccountKey:
		actorMethods = evmv11.Methods
	case manifest.PlaceholderKey:
		actorMethods = evmv11.Methods
	case manifest.SystemKey:
		actorMethods = systemv11.Methods
	default:
		return nil, parser.ErrNotKnownActor
	}
	return actorMethods, nil
}

func v12Methods(actorName string) (map[abi.MethodNum]builtin.MethodMeta, error) {
	var actorMethods map[abi.MethodNum]builtin.MethodMeta
	switch actorName {
	case manifest.AccountKey:
		actorMethods = accountv12.Methods
	case manifest.InitKey:
		actorMethods = builtinInitv12.Methods
	case manifest.CronKey:
		actorMethods = cronv12.Methods
	case manifest.RewardKey:
		actorMethods = rewardv12.Methods
	case manifest.MultisigKey:
		actorMethods = multisigv12.Methods
	case manifest.PaychKey:
		actorMethods = paychv12.Methods
	case manifest.MarketKey:
		actorMethods = marketv12.Methods
	case manifest.PowerKey:
		actorMethods = powerv12.Methods
	case manifest.MinerKey:
		actorMethods = minerv12.Methods
	case manifest.VerifregKey:
		actorMethods = verifregv12.Methods
	case manifest.EamKey:
		actorMethods = eamv12.Methods
	case manifest.EvmKey:
		actorMethods = evmv12.Methods
	case manifest.DatacapKey:
		actorMethods = datacapv12.Methods
	case manifest.EthAccountKey:
		actorMethods = evmv12.Methods
	case manifest.PlaceholderKey:
		actorMethods = evmv12.Methods
	case manifest.SystemKey:
		actorMethods = systemv12.Methods
	default:
		return nil, parser.ErrNotKnownActor
	}
	return actorMethods, nil
}

func v13Methods(actorName string) (map[abi.MethodNum]builtin.MethodMeta, error) {
	var actorMethods map[abi.MethodNum]builtin.MethodMeta
	switch actorName {
	case manifest.AccountKey:
		actorMethods = accountv13.Methods
	case manifest.InitKey:
		actorMethods = builtinInitv13.Methods
	case manifest.CronKey:
		actorMethods = cronv13.Methods
	case manifest.RewardKey:
		actorMethods = rewardv13.Methods
	case manifest.MultisigKey:
		actorMethods = multisigv13.Methods
	case manifest.PaychKey:
		actorMethods = paychv13.Methods
	case manifest.MarketKey:
		actorMethods = marketv13.Methods
	case manifest.PowerKey:
		actorMethods = powerv13.Methods
	case manifest.MinerKey:
		actorMethods = minerv13.Methods
	case manifest.VerifregKey:
		actorMethods = verifregv13.Methods
	case manifest.EamKey:
		actorMethods = eamv13.Methods
	case manifest.EvmKey:
		actorMethods = evmv13.Methods
	case manifest.DatacapKey:
		actorMethods = datacapv13.Methods
	case manifest.EthAccountKey:
		actorMethods = evmv13.Methods
	case manifest.PlaceholderKey:
		actorMethods = evmv13.Methods
	case manifest.SystemKey:
		actorMethods = systemv13.Methods
	default:
		return nil, parser.ErrNotKnownActor
	}
	return actorMethods, nil
}

func v14Methods(actorName string) (map[abi.MethodNum]builtin.MethodMeta, error) {
	var actorMethods map[abi.MethodNum]builtin.MethodMeta
	switch actorName {
	case manifest.AccountKey:
		actorMethods = accountv14.Methods
	case manifest.InitKey:
		actorMethods = builtinInitv14.Methods
	case manifest.CronKey:
		actorMethods = cronv14.Methods
	case manifest.RewardKey:
		actorMethods = rewardv14.Methods
	case manifest.MultisigKey:
		actorMethods = multisigv14.Methods
	case manifest.PaychKey:
		actorMethods = paychv14.Methods
	case manifest.MarketKey:
		actorMethods = marketv14.Methods
	case manifest.PowerKey:
		actorMethods = powerv14.Methods
	case manifest.MinerKey:
		actorMethods = minerv14.Methods
	case manifest.VerifregKey:
		actorMethods = verifregv14.Methods
	case manifest.EamKey:
		actorMethods = eamv14.Methods
	case manifest.EvmKey:
		actorMethods = evmv14.Methods
	case manifest.DatacapKey:
		actorMethods = datacapv14.Methods
	case manifest.EthAccountKey:
		actorMethods = evmv14.Methods
	case manifest.PlaceholderKey:
		actorMethods = evmv14.Methods
	case manifest.SystemKey:
		actorMethods = systemv14.Methods
	default:
		return nil, parser.ErrNotKnownActor
	}
	return actorMethods, nil
}

func v15Methods(actorName string) (map[abi.MethodNum]builtin.MethodMeta, error) {
	var actorMethods map[abi.MethodNum]builtin.MethodMeta
	switch actorName {
	case manifest.AccountKey:
		actorMethods = accountv15.Methods
	case manifest.InitKey:
		actorMethods = builtinInitv15.Methods
	case manifest.CronKey:
		actorMethods = cronv15.Methods
	case manifest.RewardKey:
		actorMethods = rewardv15.Methods
	case manifest.MultisigKey:
		actorMethods = multisigv15.Methods
	case manifest.PaychKey:
		actorMethods = paychv15.Methods
	case manifest.MarketKey:
		actorMethods = marketv15.Methods
	case manifest.PowerKey:
		actorMethods = powerv15.Methods
	case manifest.MinerKey:
		actorMethods = minerv15.Methods
	case manifest.VerifregKey:
		actorMethods = verifregv15.Methods
	case manifest.EamKey:
		actorMethods = eamv15.Methods
	case manifest.EvmKey:
		actorMethods = evmv15.Methods
	case manifest.DatacapKey:
		actorMethods = datacapv15.Methods
	case manifest.EthAccountKey:
		actorMethods = evmv15.Methods
	case manifest.PlaceholderKey:
		actorMethods = evmv15.Methods
	case manifest.SystemKey:
		actorMethods = systemv15.Methods
	default:
		return nil, parser.ErrNotKnownActor
	}
	return actorMethods, nil
}
