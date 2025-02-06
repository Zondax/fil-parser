package actortest

import (
	"fmt"
	"testing"

	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
	actorsV1 "github.com/zondax/fil-parser/actors/v1"
	actorsV2 "github.com/zondax/fil-parser/actors/v2"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

var minerWithParamsOrReturnTests = []struct {
	name   string
	txType string
	key    string
}{
	{
		name:   "Constructor",
		txType: parser.MethodConstructor,
		key:    parser.ParamsKey,
	},
	{
		name:   "Apply Rewards",
		txType: parser.MethodApplyRewards,
		key:    parser.ParamsKey,
	},
	{
		name:   "Change Beneficiary",
		txType: parser.MethodChangeBeneficiary,
		key:    parser.ParamsKey,
	},
	{
		name:   "Change Beneficiary Exported",
		txType: parser.MethodChangeBeneficiaryExported,
		key:    parser.ParamsKey,
	},
	{
		name:   "Change Multiaddrs",
		txType: parser.MethodChangeMultiaddrs,
		key:    parser.ParamsKey,
	},
	{
		name:   "Change Multiaddrs Exported",
		txType: parser.MethodChangeMultiaddrsExported,
		key:    parser.ParamsKey,
	},
	{
		name:   "Change Owner Address",
		txType: parser.MethodChangeOwnerAddress,
		key:    parser.ParamsKey,
	},
	{
		name:   "Change Peer ID",
		txType: parser.MethodChangePeerID,
		key:    parser.ParamsKey,
	},
	{
		name:   "Change Peer ID Exported",
		txType: parser.MethodChangePeerIDExported,
		key:    parser.ParamsKey,
	},
	{
		name:   "Change Worker Address",
		txType: parser.MethodChangeWorkerAddress,
		key:    parser.ParamsKey,
	},
	{
		name:   "Confirm Sector Proofs Valid",
		txType: parser.MethodConfirmSectorProofsValid,
		key:    parser.ParamsKey,
	},
	{
		name:   "Declare Faults Recovered",
		txType: parser.MethodDeclareFaultsRecovered,
		key:    parser.ParamsKey,
	},
	{
		name:   "Dispute Windowed Post",
		txType: parser.MethodDisputeWindowedPoSt,
		key:    parser.ParamsKey,
	},
	{
		name:   "Extend Sector Expiration",
		txType: parser.MethodExtendSectorExpiration,
		key:    parser.ParamsKey,
	},
	{
		name:   "Extend Sector Expiration2",
		txType: parser.MethodExtendSectorExpiration2,
		key:    parser.ParamsKey,
	},
	{
		name:   "On Deferred Cron Event",
		txType: parser.MethodOnDeferredCronEvent,
		key:    parser.ParamsKey,
	},
	{
		name:   "PreCommit Sector",
		txType: parser.MethodPreCommitSector,
		key:    parser.ParamsKey,
	},
	{
		name:   "PreCommit Sector Batch",
		txType: parser.MethodPreCommitSectorBatch,
		key:    parser.ParamsKey,
	},
	{
		name:   "PreCommit Sector Batch2",
		txType: parser.MethodPreCommitSectorBatch2,
		key:    parser.ParamsKey,
	},
	{
		name:   "Prove Commit Aggregate",
		txType: parser.MethodProveCommitAggregate,
		key:    parser.ParamsKey,
	},
	{
		name:   "Prove Commit Sector",
		txType: parser.MethodProveCommitSector,
		key:    parser.ParamsKey,
	},
	{
		name:   "Prove Replica Updated",
		txType: parser.MethodProveReplicaUpdates,
		key:    parser.ParamsKey,
	},
	{
		name:   "Submit Windowed Post",
		txType: parser.MethodSubmitWindowedPoSt,
		key:    parser.ParamsKey,
	},
	{
		name:   "Withdraw Balance",
		txType: parser.MethodWithdrawBalance,
		key:    parser.ParamsKey,
	},
	{
		name:   "Withdraw Balance Exported",
		txType: parser.MethodWithdrawBalanceExported,
		key:    parser.ParamsKey,
	},
	{
		name:   "Declare Faults",
		txType: parser.MethodDeclareFaults,
		key:    parser.ParamsKey,
	},
	{
		name:   "Report Consensus Fault",
		txType: parser.MethodReportConsensusFault,
		key:    parser.ParamsKey,
	},
	{
		name:   "Compact Partitions",
		txType: parser.MethodCompactPartitions,
		key:    parser.ParamsKey,
	},
	{
		name:   "Compact Sector Numbers",
		txType: parser.MethodCompactSectorNumbers,
		key:    parser.ParamsKey,
	},
	{
		name:   "Change Owner Address",
		txType: parser.MethodChangeOwnerAddress,
		key:    parser.ParamsKey,
	},
	{
		name:   "Change Owner Address Exported",
		txType: parser.MethodChangeOwnerAddressExported,
		key:    parser.ParamsKey,
	},
	{
		name:   "Get Owner",
		txType: parser.MethodGetOwner,
		key:    parser.ReturnKey,
	},
	{
		name:   "Get Available Balance",
		txType: parser.MethodGetAvailableBalance,
		key:    parser.ReturnKey,
	},
	{
		name:   "Check Sector Proven",
		txType: parser.MethodCheckSectorProven,
		key:    parser.ParamsKey,
	},
	{
		name:   "Get Vesting Funds",
		txType: parser.MethodGetVestingFunds,
		key:    parser.ReturnKey,
	},
	{
		name:   "Get Peer ID",
		txType: parser.MethodGetPeerID,
		key:    parser.ReturnKey,
	},
	{
		name:   "Multiaddrs",
		txType: parser.MethodGetMultiaddrs,
		key:    parser.ReturnKey,
	},
}

var minerWithParamsAndReturnTests = []struct {
	name   string
	txType string
}{
	{
		name:   "Control Addresses",
		txType: parser.MethodControlAddresses,
	},
	{
		name:   "Is Controlling Addresses Exported",
		txType: parser.MethodIsControllingAddressExported,
	},
	{
		name:   "Terminate Sectors",
		txType: parser.MethodTerminateSectors,
	},
	{
		name:   "Control Addresses",
		txType: parser.MethodControlAddresses,
	},
	{
		name:   "Prove Replica Updates2",
		txType: parser.MethodProveReplicaUpdates2,
	},
	{
		name:   "Get Beneficiary",
		txType: parser.MethodGetBeneficiary,
	},
	// { TODO: Get file
	//	name:   "Prove Commit Sectors 3",
	//	txType: parser.MethodProveCommitSectors3,
	//	f:      p.proveCommitSectors3,
	// },
}

func TestActorParserV1_MinerWithParamsOrReturn(t *testing.T) {
	p := getActorParser(actorsV1.NewActorParser).(*actorsV1.ActorParser)

	for _, tt := range minerWithParamsOrReturnTests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, err := loadFile(manifest.MinerKey, tt.txType, tt.key)
			require.NoError(t, err)
			require.NotNil(t, rawParams)

			msg := &parser.LotusMessage{}
			msgRct := &parser.LotusMessageReceipt{}

			if tt.key == parser.ReturnKey {
				msgRct.Return = rawParams
			} else {
				msg.Params = rawParams
			}

			got, err := p.ParseStorageminer(tt.txType, msg, msgRct)

			require.NoError(t, err)
			require.NotNil(t, got)
			require.Contains(t, got, tt.key, fmt.Sprintf("%s could no be found in metadata", tt.key))
			require.NotNil(t, got[tt.key])
		})
	}
}

func TestActorParserV1_MinerWithParamsAndReturn(t *testing.T) {
	p := getActorParser(actorsV1.NewActorParser).(*actorsV1.ActorParser)

	for _, tt := range minerWithParamsAndReturnTests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, rawReturn, err := getParamsAndReturn(manifest.MinerKey, tt.txType)
			require.NoError(t, err)
			require.NotNil(t, rawParams)
			require.NotNil(t, rawReturn)

			got, err := p.ParseStorageminer(tt.txType, &parser.LotusMessage{
				Params: rawParams,
			}, &parser.LotusMessageReceipt{
				Return: rawReturn,
			})
			require.NoError(t, err)
			require.NotNil(t, got)
			require.Contains(t, got, parser.ParamsKey, "Params could no be found in metadata")
			require.NotNil(t, got[parser.ParamsKey])
			require.Contains(t, got, parser.ReturnKey, "Return could no be found in metadata")
			require.NotNil(t, got[parser.ReturnKey])
		})
	}
}

func TestActorParserV2_MinerWithParamsOrReturn(t *testing.T) {
	p := getActorParser(actorsV2.NewActorParser).(*actorsV2.ActorParser)
	actor, err := p.GetActor(manifest.MinerKey)
	require.NoError(t, err)
	require.NotNil(t, actor)

	for _, tt := range minerWithParamsOrReturnTests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, err := loadFile(manifest.MinerKey, tt.txType, tt.key)
			require.NoError(t, err)
			require.NotNil(t, rawParams)

			msg := &parser.LotusMessage{}
			msgRct := &parser.LotusMessageReceipt{}

			if tt.key == parser.ReturnKey {
				msgRct.Return = rawParams
			} else {
				msg.Params = rawParams
			}

			got, _, err := actor.Parse(network, tools.V20.Height(), tt.txType, msg, msgRct, cid.Undef, filTypes.EmptyTSK)
			require.NoError(t, err)
			require.NotNil(t, got)
			require.Contains(t, got, tt.key, fmt.Sprintf("%s could no be found in metadata", tt.key))
			require.NotNil(t, got[tt.key])
		})
	}
}

func TestActorParserV2_MinerWithParamsAndReturn(t *testing.T) {
	p := getActorParser(actorsV2.NewActorParser).(*actorsV2.ActorParser)
	actor, err := p.GetActor(manifest.MinerKey)
	require.NoError(t, err)
	require.NotNil(t, actor)

	for _, tt := range minerWithParamsAndReturnTests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, rawReturn, err := getParamsAndReturn(manifest.MinerKey, tt.txType)
			require.NoError(t, err)
			require.NotNil(t, rawParams)
			require.NotNil(t, rawReturn)

			got, _, err := actor.Parse(network, tools.LatestVersion.Height(), tt.txType, &parser.LotusMessage{
				Params: rawParams,
			}, &parser.LotusMessageReceipt{
				Return: rawReturn,
			}, cid.Undef, filTypes.EmptyTSK)
			require.NoError(t, err)
			require.NotNil(t, got)
			require.Contains(t, got, parser.ParamsKey, "Params could no be found in metadata")
			require.NotNil(t, got[parser.ParamsKey])
			require.Contains(t, got, parser.ReturnKey, "Return could no be found in metadata")
			require.NotNil(t, got[parser.ReturnKey])
		})
	}
}
