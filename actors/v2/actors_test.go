package v2_test

import (
	"errors"
	"testing"

	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/actors"
	v2 "github.com/zondax/fil-parser/actors/v2"
	"github.com/zondax/fil-parser/actors/v2/account"
	"github.com/zondax/fil-parser/actors/v2/cron"
	"github.com/zondax/fil-parser/actors/v2/datacap"
	"github.com/zondax/fil-parser/actors/v2/eam"
	"github.com/zondax/fil-parser/actors/v2/ethaccount"
	"github.com/zondax/fil-parser/actors/v2/evm"
	initActor "github.com/zondax/fil-parser/actors/v2/init"
	"github.com/zondax/fil-parser/actors/v2/market"
	"github.com/zondax/fil-parser/actors/v2/miner"
	"github.com/zondax/fil-parser/actors/v2/multisig"
	paymentchannel "github.com/zondax/fil-parser/actors/v2/paymentChannel"
	"github.com/zondax/fil-parser/actors/v2/power"
	"github.com/zondax/fil-parser/actors/v2/reward"
	verifiedregistry "github.com/zondax/fil-parser/actors/v2/verifiedRegistry"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

var actorParsers = []v2.Actor{
	&account.Account{},
	&cron.Cron{},
	&initActor.Init{},
	&datacap.Datacap{},
	&eam.Eam{},
	&evm.Evm{},
	&market.Market{},
	&miner.Miner{},
	&multisig.Msig{},
	&paymentchannel.PaymentChannel{},
	&power.Power{},
	&reward.Reward{},
	&verifiedregistry.VerifiedRegistry{},
	&ethaccount.EthAccount{},
}

// TestVersionCoverage tests that all actor methods are supported for all supported network versions
func TestVersionCoverage(t *testing.T) {
	network := "mainnet"
	versions := tools.GetSupportedVersions(network)

	for _, version := range versions {
		height := tools.DeterministicTestHeight(version)
		for _, actor := range actorParsers {
			transactionTypes := actor.TransactionTypes()
			for txType := range transactionTypes {
				_, _, err := actor.Parse(network, tools.DeterministicTestHeight(version), txType, &parser.LotusMessage{}, &parser.LotusMessageReceipt{}, cid.Undef, filTypes.TipSetKey{})
				require.Falsef(t, errors.Is(err, actors.ErrUnsupportedHeight), "Missing support for txType: %s, actor: %s version: %s height: %d", txType, actor.Name(), version, height)
			}
		}
	}
}

func TestAllActorsSupported(t *testing.T) {
	// we use v10 to ensure all evm methods are retrieved as well
	filActors := manifest.GetBuiltinActorsKeys(10)
	exclude := map[string]bool{
		manifest.SystemKey:      true,
		manifest.PlaceholderKey: true,
	}

	for _, filActor := range filActors {
		if exclude[filActor] {
			continue
		}
		found := false
		for _, actor := range actorParsers {
			if actor.Name() == filActor {
				found = true
				break
			}
		}
		require.True(t, found, "Actor %s is not supported", filActor)
	}
}
