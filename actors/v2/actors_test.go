package v2_test

import (
	"errors"
	"testing"

	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/actors"
	v2 "github.com/zondax/fil-parser/actors/v2"
	"github.com/zondax/fil-parser/actors/v2/account"
	"github.com/zondax/fil-parser/actors/v2/cron"
	"github.com/zondax/fil-parser/actors/v2/datacap"
	"github.com/zondax/fil-parser/actors/v2/eam"
	"github.com/zondax/fil-parser/actors/v2/evm"
	initActor "github.com/zondax/fil-parser/actors/v2/init"
	"github.com/zondax/fil-parser/actors/v2/market"
	"github.com/zondax/fil-parser/actors/v2/miner"
	paymentchannel "github.com/zondax/fil-parser/actors/v2/paymentChannel"
	"github.com/zondax/fil-parser/actors/v2/power"
	"github.com/zondax/fil-parser/actors/v2/reward"
	verifiedregistry "github.com/zondax/fil-parser/actors/v2/verifiedRegistry"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

// TestVersionCoverage tests that all actor methods are supported for all supported network versions
func TestVersionCoverage(t *testing.T) {
	network := "mainnet"
	actorParsers := []v2.Actor{
		&account.Account{},
		&cron.Cron{},
		&initActor.Init{},
		&datacap.Datacap{},
		&eam.Eam{},
		&evm.Evm{},
		&market.Market{},
		&miner.Miner{},
		// &multisig.Msig{},
		&paymentchannel.PaymentChannel{},
		&power.Power{},
		&reward.Reward{},
		&verifiedregistry.VerifiedRegistry{},
	}
	versions := tools.GetSupportedVersions(network)

	for _, version := range versions {
		height := tools.DeterministicTestHeight(version)
		for _, actor := range actorParsers {
			transactionTypes := actor.TransactionTypes()
			for txType := range transactionTypes {
				_, _, err := actor.Parse(network, tools.DeterministicTestHeight(version), txType, &parser.LotusMessage{}, &parser.LotusMessageReceipt{}, cid.Undef)
				require.Falsef(t, errors.Is(err, actors.ErrUnsupportedHeight), "Missing support for txType: %s, actor: %s version: %s height: %d", txType, actor.Name(), version, height)
			}
		}
	}
}
