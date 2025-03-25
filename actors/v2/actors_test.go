package v2_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/zondax/golem/pkg/logger"
	"os"
	"testing"

	"github.com/zondax/fil-parser/actors/metrics"
	metrics2 "github.com/zondax/fil-parser/metrics"

	builtinActors "github.com/filecoin-project/go-state-types/actors"
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/assert"
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
	"github.com/zondax/fil-parser/actors/v2/placeholder"
	"github.com/zondax/fil-parser/actors/v2/power"
	"github.com/zondax/fil-parser/actors/v2/reward"
	"github.com/zondax/fil-parser/actors/v2/system"
	verifiedregistry "github.com/zondax/fil-parser/actors/v2/verifiedRegistry"
	logger2 "github.com/zondax/fil-parser/logger"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"

	// all multisig version imports
	multisig10 "github.com/filecoin-project/go-state-types/builtin/v10/multisig"
	multisig11 "github.com/filecoin-project/go-state-types/builtin/v11/multisig"
	multisig12 "github.com/filecoin-project/go-state-types/builtin/v12/multisig"
	multisig13 "github.com/filecoin-project/go-state-types/builtin/v13/multisig"
	multisig14 "github.com/filecoin-project/go-state-types/builtin/v14/multisig"
	multisig15 "github.com/filecoin-project/go-state-types/builtin/v15/multisig"
	multisig8 "github.com/filecoin-project/go-state-types/builtin/v8/multisig"
	multisig9 "github.com/filecoin-project/go-state-types/builtin/v9/multisig"
	legacymultisig1 "github.com/filecoin-project/specs-actors/actors/builtin/multisig"
	legacymultisig2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/multisig"
	legacymultisig3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/multisig"
	legacymultisig4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/multisig"
	legacymultisig5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/multisig"
	legacymultisig6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/multisig"
	legacymultisig7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/multisig"

	// all account version imports
	legacyaccountv1 "github.com/filecoin-project/specs-actors/actors/builtin/account"
	legacyaccountv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/account"
	legacyaccountv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/account"
	legacyaccountv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/account"
	legacyaccountv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/account"
	legacyaccountv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/account"
	legacyaccountv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/account"

	accountv10 "github.com/filecoin-project/go-state-types/builtin/v10/account"
	accountv11 "github.com/filecoin-project/go-state-types/builtin/v11/account"
	accountv12 "github.com/filecoin-project/go-state-types/builtin/v12/account"
	accountv13 "github.com/filecoin-project/go-state-types/builtin/v13/account"
	accountv14 "github.com/filecoin-project/go-state-types/builtin/v14/account"
	accountv15 "github.com/filecoin-project/go-state-types/builtin/v15/account"
	accountv8 "github.com/filecoin-project/go-state-types/builtin/v8/account"
	accountv9 "github.com/filecoin-project/go-state-types/builtin/v9/account"

	// all cron version imports
	legacycronv1 "github.com/filecoin-project/specs-actors/actors/builtin/cron"
	legacycronv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/cron"
	legacycronv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/cron"
	legacycronv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/cron"
	legacycronv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/cron"
	legacycronv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/cron"
	legacycronv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/cron"

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
	datacapv9 "github.com/filecoin-project/go-state-types/builtin/v9/datacap"

	// all eam version imports
	eamv10 "github.com/filecoin-project/go-state-types/builtin/v10/eam"
	eamv11 "github.com/filecoin-project/go-state-types/builtin/v11/eam"
	eamv12 "github.com/filecoin-project/go-state-types/builtin/v12/eam"
	eamv13 "github.com/filecoin-project/go-state-types/builtin/v13/eam"
	eamv14 "github.com/filecoin-project/go-state-types/builtin/v14/eam"
	eamv15 "github.com/filecoin-project/go-state-types/builtin/v15/eam"

	// all ethaccount version imports
	ethaccountActorv10 "github.com/filecoin-project/go-state-types/builtin/v10/ethaccount"
	ethaccountActorv11 "github.com/filecoin-project/go-state-types/builtin/v11/ethaccount"
	ethaccountActorv12 "github.com/filecoin-project/go-state-types/builtin/v12/ethaccount"
	ethaccountActorv13 "github.com/filecoin-project/go-state-types/builtin/v13/ethaccount"
	ethaccountActorv14 "github.com/filecoin-project/go-state-types/builtin/v14/ethaccount"
	ethaccountActorv15 "github.com/filecoin-project/go-state-types/builtin/v15/ethaccount"

	// all evm version imports
	evmv10 "github.com/filecoin-project/go-state-types/builtin/v10/evm"
	evmv11 "github.com/filecoin-project/go-state-types/builtin/v11/evm"
	evmv12 "github.com/filecoin-project/go-state-types/builtin/v12/evm"
	evmv13 "github.com/filecoin-project/go-state-types/builtin/v13/evm"
	evmv14 "github.com/filecoin-project/go-state-types/builtin/v14/evm"
	evmv15 "github.com/filecoin-project/go-state-types/builtin/v15/evm"

	// all init version imports
	legacyInitv1 "github.com/filecoin-project/specs-actors/actors/builtin/init"
	legacyInitv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/init"
	legacyInitv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/init"
	legacyInitv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/init"
	legacyInitv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/init"
	legacyInitv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/init"
	legacyInitv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/init"

	builtinInitv10 "github.com/filecoin-project/go-state-types/builtin/v10/init"
	builtinInitv11 "github.com/filecoin-project/go-state-types/builtin/v11/init"
	builtinInitv12 "github.com/filecoin-project/go-state-types/builtin/v12/init"
	builtinInitv13 "github.com/filecoin-project/go-state-types/builtin/v13/init"
	builtinInitv14 "github.com/filecoin-project/go-state-types/builtin/v14/init"
	builtinInitv15 "github.com/filecoin-project/go-state-types/builtin/v15/init"
	builtinInitv8 "github.com/filecoin-project/go-state-types/builtin/v8/init"
	builtinInitv9 "github.com/filecoin-project/go-state-types/builtin/v9/init"

	// all market version imports
	legacymarketv1 "github.com/filecoin-project/specs-actors/actors/builtin/market"
	legacymarketv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/market"
	legacymarketv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/market"
	legacymarketv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/market"
	legacymarketv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/market"
	legacymarketv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/market"
	legacymarketv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/market"

	marketv10 "github.com/filecoin-project/go-state-types/builtin/v10/market"
	marketv11 "github.com/filecoin-project/go-state-types/builtin/v11/market"
	marketv12 "github.com/filecoin-project/go-state-types/builtin/v12/market"
	marketv13 "github.com/filecoin-project/go-state-types/builtin/v13/market"
	marketv14 "github.com/filecoin-project/go-state-types/builtin/v14/market"
	marketv15 "github.com/filecoin-project/go-state-types/builtin/v15/market"
	marketv8 "github.com/filecoin-project/go-state-types/builtin/v8/market"
	marketv9 "github.com/filecoin-project/go-state-types/builtin/v9/market"

	// all miner version imports
	legacyminerv1 "github.com/filecoin-project/specs-actors/actors/builtin/miner"
	legacyminerv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/miner"
	legacyminerv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/miner"
	legacyminerv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/miner"
	legacyminerv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/miner"
	legacyminerv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/miner"
	legacyminerv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/miner"

	minerv10 "github.com/filecoin-project/go-state-types/builtin/v10/miner"
	minerv11 "github.com/filecoin-project/go-state-types/builtin/v11/miner"
	minerv12 "github.com/filecoin-project/go-state-types/builtin/v12/miner"
	minerv13 "github.com/filecoin-project/go-state-types/builtin/v13/miner"
	minerv14 "github.com/filecoin-project/go-state-types/builtin/v14/miner"
	minerv15 "github.com/filecoin-project/go-state-types/builtin/v15/miner"
	minerv8 "github.com/filecoin-project/go-state-types/builtin/v8/miner"
	minerv9 "github.com/filecoin-project/go-state-types/builtin/v9/miner"

	// all paych version imports
	legacypaychv1 "github.com/filecoin-project/specs-actors/actors/builtin/paych"
	legacypaychv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/paych"
	legacypaychv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/paych"
	legacypaychv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/paych"
	legacypaychv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/paych"
	legacypaychv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/paych"
	legacypaychv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/paych"

	paychv10 "github.com/filecoin-project/go-state-types/builtin/v10/paych"
	paychv11 "github.com/filecoin-project/go-state-types/builtin/v11/paych"
	paychv12 "github.com/filecoin-project/go-state-types/builtin/v12/paych"
	paychv13 "github.com/filecoin-project/go-state-types/builtin/v13/paych"
	paychv14 "github.com/filecoin-project/go-state-types/builtin/v14/paych"
	paychv15 "github.com/filecoin-project/go-state-types/builtin/v15/paych"
	paychv8 "github.com/filecoin-project/go-state-types/builtin/v8/paych"
	paychv9 "github.com/filecoin-project/go-state-types/builtin/v9/paych"

	// all placeholder version imports
	placeholderv10 "github.com/filecoin-project/go-state-types/builtin/v10/placeholder"
	placeholderv11 "github.com/filecoin-project/go-state-types/builtin/v11/placeholder"
	placeholderv12 "github.com/filecoin-project/go-state-types/builtin/v12/placeholder"
	placeholderv13 "github.com/filecoin-project/go-state-types/builtin/v13/placeholder"
	placeholderv14 "github.com/filecoin-project/go-state-types/builtin/v14/placeholder"
	placeholderv15 "github.com/filecoin-project/go-state-types/builtin/v15/placeholder"

	// all power version imports
	legacypowerv1 "github.com/filecoin-project/specs-actors/actors/builtin/power"
	legacypowerv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/power"
	legacypowerv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/power"
	legacypowerv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/power"
	legacypowerv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/power"
	legacypowerv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/power"
	legacypowerv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/power"

	powerv10 "github.com/filecoin-project/go-state-types/builtin/v10/power"
	powerv11 "github.com/filecoin-project/go-state-types/builtin/v11/power"
	powerv12 "github.com/filecoin-project/go-state-types/builtin/v12/power"
	powerv13 "github.com/filecoin-project/go-state-types/builtin/v13/power"
	powerv14 "github.com/filecoin-project/go-state-types/builtin/v14/power"
	powerv15 "github.com/filecoin-project/go-state-types/builtin/v15/power"
	powerv8 "github.com/filecoin-project/go-state-types/builtin/v8/power"
	powerv9 "github.com/filecoin-project/go-state-types/builtin/v9/power"

	// all reward version imports
	legacyrewardv1 "github.com/filecoin-project/specs-actors/actors/builtin/reward"
	legacyrewardv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/reward"
	legacyrewardv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/reward"
	legacyrewardv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/reward"
	legacyrewardv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/reward"
	legacyrewardv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/reward"
	legacyrewardv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/reward"

	rewardv10 "github.com/filecoin-project/go-state-types/builtin/v10/reward"
	rewardv11 "github.com/filecoin-project/go-state-types/builtin/v11/reward"
	rewardv12 "github.com/filecoin-project/go-state-types/builtin/v12/reward"
	rewardv13 "github.com/filecoin-project/go-state-types/builtin/v13/reward"
	rewardv14 "github.com/filecoin-project/go-state-types/builtin/v14/reward"
	rewardv15 "github.com/filecoin-project/go-state-types/builtin/v15/reward"
	rewardv8 "github.com/filecoin-project/go-state-types/builtin/v8/reward"
	rewardv9 "github.com/filecoin-project/go-state-types/builtin/v9/reward"

	// all system version imports
	legacysystemv1 "github.com/filecoin-project/specs-actors/actors/builtin/system"
	legacysystemv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/system"
	legacysystemv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/system"
	legacysystemv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/system"
	legacysystemv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/system"
	legacysystemv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/system"
	legacysystemv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/system"

	systemv10 "github.com/filecoin-project/go-state-types/builtin/v10/system"
	systemv11 "github.com/filecoin-project/go-state-types/builtin/v11/system"
	systemv12 "github.com/filecoin-project/go-state-types/builtin/v12/system"
	systemv13 "github.com/filecoin-project/go-state-types/builtin/v13/system"
	systemv14 "github.com/filecoin-project/go-state-types/builtin/v14/system"
	systemv15 "github.com/filecoin-project/go-state-types/builtin/v15/system"
	systemv8 "github.com/filecoin-project/go-state-types/builtin/v8/system"
	systemv9 "github.com/filecoin-project/go-state-types/builtin/v9/system"

	// all verifiedregistry version imports
	legacyverifregv1 "github.com/filecoin-project/specs-actors/actors/builtin/verifreg"
	legacyverifregv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/verifreg"
	legacyverifregv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/verifreg"
	legacyverifregv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/verifreg"
	legacyverifregv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/verifreg"
	legacyverifregv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/verifreg"
	legacyverifregv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/verifreg"

	verifregv10 "github.com/filecoin-project/go-state-types/builtin/v10/verifreg"
	verifregv11 "github.com/filecoin-project/go-state-types/builtin/v11/verifreg"
	verifregv12 "github.com/filecoin-project/go-state-types/builtin/v12/verifreg"
	verifregv13 "github.com/filecoin-project/go-state-types/builtin/v13/verifreg"
	verifregv14 "github.com/filecoin-project/go-state-types/builtin/v14/verifreg"
	verifregv15 "github.com/filecoin-project/go-state-types/builtin/v15/verifreg"
	verifregv8 "github.com/filecoin-project/go-state-types/builtin/v8/verifreg"
	verifregv9 "github.com/filecoin-project/go-state-types/builtin/v9/verifreg"
)

var latestBuiltinActorVersion uint64
var l *logger.Logger

func TestMain(m *testing.M) {
	version, err := v2.LatestBuiltinActorVersion()
	if err != nil {
		panic(fmt.Sprintf("failed to get latest builtin actor version: %v", err))
	}
	fmt.Printf("latestBuiltinActorVersion: %d\n", version)
	latestBuiltinActorVersion = version

	l = logger2.GetSafeLogger(logger.NewDevelopmentLogger())
	os.Exit(m.Run())
}

// TestVersionCoverage tests that all actor methods are supported for all supported network versions
func TestVersionCoverage(t *testing.T) {
	network := "mainnet"
	versions := tools.GetSupportedVersions(network)
	actorParsers := getActors(t)

	for _, version := range versions {
		height := tools.DeterministicTestHeight(version)
		for _, actor := range actorParsers {
			transactionTypes := actor.TransactionTypes()
			for txType := range transactionTypes {
				_, _, err := actor.Parse(context.Background(), network, tools.DeterministicTestHeight(version), txType, &parser.LotusMessage{}, &parser.LotusMessageReceipt{}, cid.Undef, filTypes.TipSetKey{})
				require.Falsef(t, errors.Is(err, actors.ErrUnsupportedHeight), "Missing support for txType: %s, actor: %s version: %s height: %d", txType, actor.Name(), version, height)
				require.Falsef(t, errors.Is(err, parser.ErrUnknownMethod), "Method missing in actor.Parse: %s, actor: %s version: %s height: %d", txType, actor.Name(), version, height)
			}
		}
	}
}

// TestAllActorsSupported tests that all actors are supported for the latest actor version
func TestAllActorsSupported(t *testing.T) {
	getActors(t)
}

func getActors(t *testing.T) []v2.Actor {
	actorParser := v2.NewActorParser("mainnet", nil, l, metrics2.NewNoopMetricsClient()).(*v2.ActorParser)
	filActors := manifest.GetBuiltinActorsKeys(builtinActors.Version(latestBuiltinActorVersion))
	actors := []v2.Actor{}
	for _, filActor := range filActors {
		actor, err := actorParser.GetActor(filActor, &metrics.ActorsMetricsClient{MetricsClient: metrics2.NewNoopMetricsClient()})
		require.NoErrorf(t, err, "Actor %s is not supported", filActor)
		actors = append(actors, actor)
	}
	return actors
}

// TestABIMethodNumberToMethodName tests that the method number is mapped to the correct method name for every version
func TestABIMethodNumberToMethodName(t *testing.T) {
	network := "mainnet"

	versions := tools.GetSupportedVersions(network)
	require.NotEmpty(t, versions)

	actorParsers := getActors(t)
	require.NotEmpty(t, actorParsers)

	for _, version := range versions {
		height := tools.DeterministicTestHeight(version)
		for _, actor := range actorParsers {
			transactionTypes := actor.TransactionTypes()
			// Placeholder actor has no methods
			if actor.Name() == manifest.PlaceholderKey {
				continue
			}
			require.NotEmptyf(t, transactionTypes, "Transaction types are empty for actor: %s version: %s height: %d", actor.Name(), version, height)
			methods, err := actor.Methods(context.Background(), network, height)
			if actor.StartNetworkHeight() > height {
				continue
			}
			require.NoErrorf(t, err, "Failed to get methods for actor: %s version: %s height: %d", actor.Name(), version, height)
			require.NotEmptyf(t, methods, "Methods are empty for actor: %s version: %s height: %d", actor.Name(), version, height)

			for methodNum := range methods {
				methodName := methods[methodNum].Name
				assert.Containsf(t, transactionTypes, methodName, "Method name: %s is not in transaction types for actor: %s version: %s height: %d", methodName, actor.Name(), version, height)
			}
		}
	}
}

// TestMethodCoverage tests that all actor methods are supported for all actor versions
func TestMethodCoverage(t *testing.T) {
	tb := map[v2.Actor][]any{
		&multisig.Msig{}: {
			legacymultisig1.Actor{},
			legacymultisig2.Actor{},
			legacymultisig3.Actor{},
			legacymultisig4.Actor{},
			legacymultisig5.Actor{},
			legacymultisig6.Actor{},
			legacymultisig7.Actor{},
			multisig8.Methods,
			multisig9.Methods,
			multisig10.Methods,
			multisig11.Methods,
			multisig12.Methods,
			multisig13.Methods,
			multisig14.Methods,
			multisig15.Methods,
		},
		&account.Account{}: {
			legacyaccountv1.Actor{},
			legacyaccountv2.Actor{},
			legacyaccountv3.Actor{},
			legacyaccountv4.Actor{},
			legacyaccountv5.Actor{},
			legacyaccountv6.Actor{},
			legacyaccountv7.Actor{},
			accountv8.Methods,
			accountv9.Methods,
			accountv10.Methods,
			accountv11.Methods,
			accountv12.Methods,
			accountv13.Methods,
			accountv14.Methods,
			accountv15.Methods,
		},
		&cron.Cron{}: {
			legacycronv1.Actor{},
			legacycronv2.Actor{},
			legacycronv3.Actor{},
			legacycronv4.Actor{},
			legacycronv5.Actor{},
			legacycronv6.Actor{},
			legacycronv7.Actor{},
			cronv8.Methods,
			cronv9.Methods,
			cronv10.Methods,
			cronv11.Methods,
			cronv12.Methods,
			cronv13.Methods,
			cronv14.Methods,
			cronv15.Methods,
		},
		&datacap.Datacap{}: {
			datacapv9.Methods,
			datacapv10.Methods,
			datacapv11.Methods,
			datacapv12.Methods,
			datacapv13.Methods,
			datacapv14.Methods,
			datacapv15.Methods,
		},
		&eam.Eam{}: {
			eamv10.Methods,
			eamv11.Methods,
			eamv12.Methods,
			eamv13.Methods,
			eamv14.Methods,
			eamv15.Methods,
		},
		&ethaccount.EthAccount{}: {
			ethaccountActorv10.Methods,
			ethaccountActorv11.Methods,
			ethaccountActorv12.Methods,
			ethaccountActorv13.Methods,
			ethaccountActorv14.Methods,
			ethaccountActorv15.Methods,
		},
		&evm.Evm{}: {
			evmv10.Methods,
			evmv11.Methods,
			evmv12.Methods,
			evmv13.Methods,
			evmv14.Methods,
			evmv15.Methods,
		},
		&initActor.Init{}: {
			legacyInitv1.Actor{},
			legacyInitv2.Actor{},
			legacyInitv3.Actor{},
			legacyInitv4.Actor{},
			legacyInitv5.Actor{},
			legacyInitv6.Actor{},
			legacyInitv7.Actor{},
			builtinInitv8.Methods,
			builtinInitv9.Methods,
			builtinInitv10.Methods,
			builtinInitv11.Methods,
			builtinInitv12.Methods,
			builtinInitv13.Methods,
			builtinInitv14.Methods,
			builtinInitv15.Methods,
		},
		&market.Market{}: {
			legacymarketv1.Actor{},
			legacymarketv2.Actor{},
			legacymarketv3.Actor{},
			legacymarketv4.Actor{},
			legacymarketv5.Actor{},
			legacymarketv6.Actor{},
			legacymarketv7.Actor{},
			marketv8.Methods,
			marketv9.Methods,
			marketv10.Methods,
			marketv11.Methods,
			marketv12.Methods,
			marketv13.Methods,
			marketv14.Methods,
			marketv15.Methods,
		},
		&miner.Miner{}: {
			legacyminerv1.Actor{},
			legacyminerv2.Actor{},
			legacyminerv3.Actor{},
			legacyminerv4.Actor{},
			legacyminerv5.Actor{},
			legacyminerv6.Actor{},
			legacyminerv7.Actor{},
			minerv8.Methods,
			minerv9.Methods,
			minerv10.Methods,
			minerv11.Methods,
			minerv12.Methods,
			minerv13.Methods,
			minerv14.Methods,
			minerv15.Methods,
		},
		&paymentchannel.PaymentChannel{}: {
			legacypaychv1.Actor{},
			legacypaychv2.Actor{},
			legacypaychv3.Actor{},
			legacypaychv4.Actor{},
			legacypaychv5.Actor{},
			legacypaychv6.Actor{},
			legacypaychv7.Actor{},
			paychv8.Methods,
			paychv9.Methods,
			paychv10.Methods,
			paychv11.Methods,
			paychv12.Methods,
			paychv13.Methods,
			paychv14.Methods,
			paychv15.Methods,
		},
		&placeholder.Placeholder{}: {
			placeholderv10.Methods,
			placeholderv11.Methods,
			placeholderv12.Methods,
			placeholderv13.Methods,
			placeholderv14.Methods,
			placeholderv15.Methods,
		},
		&power.Power{}: {
			legacypowerv1.Actor{},
			legacypowerv2.Actor{},
			legacypowerv3.Actor{},
			legacypowerv4.Actor{},
			legacypowerv5.Actor{},
			legacypowerv6.Actor{},
			legacypowerv7.Actor{},
			powerv8.Methods,
			powerv9.Methods,
			powerv10.Methods,
			powerv11.Methods,
			powerv12.Methods,
			powerv13.Methods,
			powerv14.Methods,
			powerv15.Methods,
		},
		&reward.Reward{}: {
			legacyrewardv1.Actor{},
			legacyrewardv2.Actor{},
			legacyrewardv3.Actor{},
			legacyrewardv4.Actor{},
			legacyrewardv5.Actor{},
			legacyrewardv6.Actor{},
			legacyrewardv7.Actor{},
			rewardv8.Methods,
			rewardv9.Methods,
			rewardv10.Methods,
			rewardv11.Methods,
			rewardv12.Methods,
			rewardv13.Methods,
			rewardv14.Methods,
			rewardv15.Methods,
		},
		&system.System{}: {
			legacysystemv1.Actor{},
			legacysystemv2.Actor{},
			legacysystemv3.Actor{},
			legacysystemv4.Actor{},
			legacysystemv5.Actor{},
			legacysystemv6.Actor{},
			legacysystemv7.Actor{},
			systemv8.Methods,
			systemv9.Methods,
			systemv10.Methods,
			systemv11.Methods,
			systemv12.Methods,
			systemv13.Methods,
			systemv14.Methods,
			systemv15.Methods,
		},
		&verifiedregistry.VerifiedRegistry{}: {
			legacyverifregv1.Actor{},
			legacyverifregv2.Actor{},
			legacyverifregv3.Actor{},
			legacyverifregv4.Actor{},
			legacyverifregv5.Actor{},
			legacyverifregv6.Actor{},
			legacyverifregv7.Actor{},
			verifregv8.Methods,
			verifregv9.Methods,
			verifregv10.Methods,
			verifregv11.Methods,
			verifregv12.Methods,
			verifregv13.Methods,
			verifregv14.Methods,
			verifregv15.Methods,
		},
	}

	for actor, versions := range tb {
		missingMethods := v2.MissingMethods(actor, versions)
		assert.Emptyf(t, missingMethods, "actor: %s, missing methods: %v", actor.Name(), missingMethods)
	}
}
