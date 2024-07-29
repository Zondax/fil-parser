package actors

import (
	"fmt"
	"testing"

	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/parser"
)

func TestActorParser_minerWithParamsOrReturn(t *testing.T) {
	p := getActorParser()
	tests := []struct {
		name   string
		txType string
		f      func([]byte) (map[string]interface{}, error)
		key    string
	}{
		{
			name:   "Constructor",
			txType: parser.MethodConstructor,
			f:      p.minerConstructor,
			key:    parser.ParamsKey,
		},
		{
			name:   "Apply Rewards",
			txType: parser.MethodApplyRewards,
			f:      p.applyRewards,
			key:    parser.ParamsKey,
		},
		{
			name:   "Change Beneficiary",
			txType: parser.MethodChangeBeneficiary,
			f:      p.changeBeneficiary,
			key:    parser.ParamsKey,
		},
		{
			name:   "Change Beneficiary Exported",
			txType: parser.MethodChangeBeneficiaryExported,
			f:      p.changeBeneficiary,
			key:    parser.ParamsKey,
		},
		{
			name:   "Change Multiaddrs",
			txType: parser.MethodChangeMultiaddrs,
			f:      p.changeMultiaddrs,
			key:    parser.ParamsKey,
		},
		{
			name:   "Change Multiaddrs Exported",
			txType: parser.MethodChangeMultiaddrsExported,
			f:      p.changeMultiaddrs,
			key:    parser.ParamsKey,
		},
		{
			name:   "Change Owner Address",
			txType: parser.MethodChangeOwnerAddress,
			f:      p.changeOwnerAddress,
			key:    parser.ParamsKey,
		},
		{
			name:   "Change Peer ID",
			txType: parser.MethodChangePeerID,
			f:      p.changePeerID,
			key:    parser.ParamsKey,
		},
		{
			name:   "Change Peer ID Exported",
			txType: parser.MethodChangePeerIDExported,
			f:      p.changePeerID,
			key:    parser.ParamsKey,
		},
		{
			name:   "Change Worker Address",
			txType: parser.MethodChangeWorkerAddress,
			f:      p.changeWorkerAddress,
			key:    parser.ParamsKey,
		},
		{
			name:   "Confirm Sector Proofs Valid",
			txType: parser.MethodConfirmSectorProofsValid,
			f:      p.confirmSectorProofsValid,
			key:    parser.ParamsKey,
		},
		{
			name:   "Declare Faults Recovered",
			txType: parser.MethodDeclareFaultsRecovered,
			f:      p.declareFaultsRecovered,
			key:    parser.ParamsKey,
		},
		{
			name:   "Dispute Windowed Post",
			txType: parser.MethodDisputeWindowedPoSt,
			f:      p.disputeWindowedPoSt,
			key:    parser.ParamsKey,
		},
		{
			name:   "Extend Sector Expiration",
			txType: parser.MethodExtendSectorExpiration,
			f:      p.extendSectorExpiration,
			key:    parser.ParamsKey,
		},
		{
			name:   "Extend Sector Expiration2",
			txType: parser.MethodExtendSectorExpiration2,
			f:      p.extendSectorExpiration2,
			key:    parser.ParamsKey,
		},
		{
			name:   "On Deferred Cron Event",
			txType: parser.MethodOnDeferredCronEvent,
			f:      p.onDeferredCronEvent,
			key:    parser.ParamsKey,
		},
		{
			name:   "PreCommit Sector",
			txType: parser.MethodPreCommitSector,
			f:      p.preCommitSector,
			key:    parser.ParamsKey,
		},
		{
			name:   "PreCommit Sector Batch",
			txType: parser.MethodPreCommitSectorBatch,
			f:      p.preCommitSectorBatch,
			key:    parser.ParamsKey,
		},
		{
			name:   "PreCommit Sector Batch2",
			txType: parser.MethodPreCommitSectorBatch2,
			f:      p.preCommitSectorBatch2,
			key:    parser.ParamsKey,
		},
		{
			name:   "Prove Commit Aggregate",
			txType: parser.MethodProveCommitAggregate,
			f:      p.proveCommitAggregate,
			key:    parser.ParamsKey,
		},
		{
			name:   "Prove Commit Sector",
			txType: parser.MethodProveCommitSector,
			f:      p.proveCommitSector,
			key:    parser.ParamsKey,
		},
		{
			name:   "Prove Replica Updated",
			txType: parser.MethodProveReplicaUpdates,
			f:      p.proveReplicaUpdates,
			key:    parser.ParamsKey,
		},
		{
			name:   "Submit Windowed Post",
			txType: parser.MethodSubmitWindowedPoSt,
			f:      p.submitWindowedPoSt,
			key:    parser.ParamsKey,
		},
		{
			name:   "Withdraw Balance",
			txType: parser.MethodWithdrawBalance,
			f:      p.parseWithdrawBalance,
			key:    parser.ParamsKey,
		},
		{
			name:   "Withdraw Balance Exported",
			txType: parser.MethodWithdrawBalanceExported,
			f:      p.parseWithdrawBalance,
			key:    parser.ParamsKey,
		},
		{
			name:   "Declare Faults",
			txType: parser.MethodDeclareFaults,
			f:      p.declareFaults,
			key:    parser.ParamsKey,
		},
		{
			name:   "Report Consensus Fault",
			txType: parser.MethodReportConsensusFault,
			f:      p.reportConsensusFault,
			key:    parser.ParamsKey,
		},
		{
			name:   "Compact Partitions",
			txType: parser.MethodCompactPartitions,
			f:      p.compactPartitions,
			key:    parser.ParamsKey,
		},
		{
			name:   "Compact Sector Numbers",
			txType: parser.MethodCompactSectorNumbers,
			f:      p.compactSectorNumbers,
			key:    parser.ParamsKey,
		},
		{
			name:   "Change Owner Address",
			txType: parser.MethodChangeOwnerAddress,
			f:      p.changeOwnerAddress,
			key:    parser.ParamsKey,
		},
		{
			name:   "Change Owner Address Exported",
			txType: parser.MethodChangeOwnerAddressExported,
			f:      p.changeOwnerAddress,
			key:    parser.ParamsKey,
		},
		{
			name:   "Get Owner",
			txType: parser.MethodGetOwner,
			f:      p.getOwner,
			key:    parser.ReturnKey,
		},
		{
			name:   "Get Available Balance",
			txType: parser.MethodGetAvailableBalance,
			f:      p.getAvailableBalance,
			key:    parser.ReturnKey,
		},
		{
			name:   "Check Sector Proven",
			txType: parser.MethodCheckSectorProven,
			f:      p.checkSectorProven,
			key:    parser.ParamsKey,
		},
		{
			name:   "Get Vesting Funds",
			txType: parser.MethodGetVestingFunds,
			f:      p.getVestingFunds,
			key:    parser.ReturnKey,
		},
		{
			name:   "Get Peer ID",
			txType: parser.MethodGetPeerID,
			f:      p.getPeerID,
			key:    parser.ReturnKey,
		},
		{
			name:   "Multiaddrs",
			txType: parser.MethodGetMultiaddrs,
			f:      p.getMultiaddrs,
			key:    parser.ReturnKey,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, err := loadFile(manifest.MinerKey, tt.txType, tt.key)
			require.NoError(t, err)
			require.NotNil(t, rawParams)

			got, err := tt.f(rawParams)
			require.NoError(t, err)
			require.NotNil(t, got)
			require.Contains(t, got, tt.key, fmt.Sprintf("%s could no be found in metadata", tt.key))
			require.NotNil(t, got[tt.key])
		})
	}
}

func TestActorParser_minerWithParamsAndReturn(t *testing.T) {
	p := getActorParser()
	tests := []struct {
		name   string
		txType string
		f      func([]byte, []byte) (map[string]interface{}, error)
	}{
		{
			name:   "Control Addresses",
			txType: parser.MethodControlAddresses,
			f:      p.controlAddresses,
		},
		{
			name:   "Is Controlling Addresses Exported",
			txType: parser.MethodIsControllingAddressExported,
			f:      p.isControllingAddressExported,
		},
		{
			name:   "Terminate Sectors",
			txType: parser.MethodTerminateSectors,
			f:      p.terminateSectors,
		},
		{
			name:   "Control Addresses",
			txType: parser.MethodControlAddresses,
			f:      p.controlAddresses,
		},
		{
			name:   "Prove Replica Updates2",
			txType: parser.MethodProveReplicaUpdates2,
			f:      p.proveReplicaUpdates2,
		},
		{
			name:   "Get Beneficiary",
			txType: parser.MethodGetBeneficiary,
			f:      p.getBeneficiary,
		},
		//{ // TODO: Add test after upgrade
		//	name:   "Prove Commit Sectors 3",
		//	txType: parser.MethodProveCommitSectors3,
		//	f:      p.proveCommitSectors3,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, rawReturn, err := getParamsAndReturn(manifest.MinerKey, tt.txType)
			require.NoError(t, err)
			require.NotNil(t, rawParams)
			require.NotNil(t, rawReturn)

			got, err := tt.f(rawParams, rawReturn)
			require.NoError(t, err)
			require.NotNil(t, got)
			require.Contains(t, got, parser.ParamsKey, "Params could no be found in metadata")
			require.NotNil(t, got[parser.ParamsKey])
			require.Contains(t, got, parser.ReturnKey, "Return could no be found in metadata")
			require.NotNil(t, got[parser.ReturnKey])
		})
	}
}
