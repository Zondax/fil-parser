package internal

import (
	"fmt"
	"strings"

	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/zondax/fil-parser/actors"
	actormetrics "github.com/zondax/fil-parser/actors/metrics"
	"github.com/zondax/fil-parser/actors/v2/account"
	"github.com/zondax/fil-parser/actors/v2/cron"
	"github.com/zondax/fil-parser/actors/v2/datacap"
	"github.com/zondax/fil-parser/actors/v2/eam"
	"github.com/zondax/fil-parser/actors/v2/ethaccount"
	"github.com/zondax/fil-parser/actors/v2/evm"
	initActor "github.com/zondax/fil-parser/actors/v2/init"
	"github.com/zondax/fil-parser/actors/v2/market"
	"github.com/zondax/fil-parser/actors/v2/miner"
	paymentchannel "github.com/zondax/fil-parser/actors/v2/paymentChannel"
	"github.com/zondax/fil-parser/actors/v2/placeholder"
	"github.com/zondax/fil-parser/actors/v2/power"
	"github.com/zondax/fil-parser/actors/v2/reward"
	"github.com/zondax/fil-parser/actors/v2/system"
	verifiedregistry "github.com/zondax/fil-parser/actors/v2/verifiedRegistry"
	"github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/golem/pkg/logger"
)

// GetActor returns a new instance of the specified actor type. It does not return multisig actors to avoid
// circular dependencies, as multisig also needs all actors to parse 'propose'.
func GetActor(actor string, logger *logger.Logger, helper *helper.Helper, metrics *actormetrics.ActorsMetricsClient) (actors.Actor, error) {
	actorName := actor
	if strings.Contains(actor, "/") {
		parts := strings.Split(actor, "/")
		actorName = parts[len(parts)-1]
	}
	switch actorName {
	case manifest.AccountKey:
		return account.New(logger), nil
	case manifest.CronKey:
		return cron.New(logger), nil
	case manifest.DatacapKey:
		return datacap.New(logger), nil
	case manifest.EamKey:
		return eam.New(helper, logger), nil
	case manifest.EthAccountKey:
		return ethaccount.New(logger), nil
	case manifest.EvmKey:
		return evm.New(logger, metrics), nil
	case manifest.InitKey:
		return initActor.New(helper, logger), nil
	case manifest.MarketKey:
		return market.New(logger), nil
	case manifest.MinerKey:
		return miner.New(logger), nil
	case manifest.PaychKey:
		return paymentchannel.New(logger), nil
	case manifest.PowerKey:
		return power.New(helper, logger), nil
	case manifest.RewardKey:
		return reward.New(logger), nil
	case manifest.VerifregKey:
		return verifiedregistry.New(logger), nil
	case manifest.PlaceholderKey:
		return placeholder.New(logger), nil
	case manifest.SystemKey:
		return system.New(logger), nil
	default:
		return nil, fmt.Errorf("actor %s not found", actor)
	}
}
