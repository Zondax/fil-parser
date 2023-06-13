package actors

import (
	"fmt"
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/parser"
	"testing"
)

func TestActorParser_minerWithParamsOrReturn(t *testing.T) {
	p := getActorParser()
	tests := []struct {
		name     string
		txType   string
		f        func([]byte) (map[string]interface{}, error)
		fileName string
		key      string
	}{
		{
			name:     "Constructor",
			txType:   parser.MethodConstructor,
			f:        p.minerConstructor,
			fileName: "params",
			key:      parser.ParamsKey,
		},
		{
			name:     "Apply Rewards",
			txType:   parser.MethodApplyRewards,
			f:        p.applyRewards,
			fileName: "params",
			key:      parser.ParamsKey,
		},
		{
			name:     "Change Beneficiary",
			txType:   parser.MethodChangeBeneficiary,
			f:        p.changeBeneficiary,
			fileName: "params",
			key:      parser.ParamsKey,
		},
		{
			name:     "Change Multiaddrs",
			txType:   parser.MethodChangeMultiaddrs,
			f:        p.changeMultiaddrs,
			fileName: "params",
			key:      parser.ParamsKey,
		},
		{
			name:     "Change Owner Address",
			txType:   parser.MethodChangeOwnerAddress,
			f:        p.changeOwnerAddress,
			fileName: "params",
			key:      parser.ParamsKey,
		},
		{
			name:     "Change Peer ID",
			txType:   parser.MethodChangePeerID,
			f:        p.changePeerID,
			fileName: "params",
			key:      parser.ParamsKey,
		},
		{
			name:     "Change Worker Address",
			txType:   parser.MethodChangeWorkerAddress,
			f:        p.changeWorkerAddress,
			fileName: "params",
			key:      parser.ParamsKey,
		},
		{
			name:     "Confirm Sector Proofs Valid",
			txType:   parser.MethodConfirmSectorProofsValid,
			f:        p.confirmSectorProofsValid,
			fileName: "params",
			key:      parser.ParamsKey,
		},
		{
			name:     "Declare FaultsRecovered",
			txType:   parser.MethodDeclareFaultsRecovered,
			f:        p.declareFaultsRecovered,
			fileName: "params",
			key:      parser.ParamsKey,
		},
		{
			name:     "Extend Sector Expiration",
			txType:   parser.MethodExtendSectorExpiration,
			f:        p.extendSectorExpiration,
			fileName: "params",
			key:      parser.ParamsKey,
		},
		{
			name:     "Extend Sector Expiration2",
			txType:   parser.MethodExtendSectorExpiration2,
			f:        p.extendSectorExpiration2,
			fileName: "params",
			key:      parser.ParamsKey,
		},
		{
			name:     "On Deferred Cron Event",
			txType:   parser.MethodOnDeferredCronEvent,
			f:        p.onDeferredCronEvent,
			fileName: "params",
			key:      parser.ParamsKey,
		},
		{
			name:     "ProCommit Sector",
			txType:   parser.MethodPreCommitSector,
			f:        p.preCommitSector,
			fileName: "params",
			key:      parser.ParamsKey,
		},
		{
			name:     "ProCommit Sector Batch",
			txType:   parser.MethodPreCommitSectorBatch,
			f:        p.preCommitSectorBatch,
			fileName: "params",
			key:      parser.ParamsKey,
		},
		{
			name:     "Prove Commit Aggregate",
			txType:   parser.MethodProveCommitAggregate,
			f:        p.proveCommitAggregate,
			fileName: "params",
			key:      parser.ParamsKey,
		},
		{
			name:     "Prove Commit Sector",
			txType:   parser.MethodProveCommitSector,
			f:        p.proveCommitSector,
			fileName: "params",
			key:      parser.ParamsKey,
		},
		{
			name:     "Prove Replica Updated",
			txType:   parser.MethodProveReplicaUpdates,
			f:        p.proveReplicaUpdates,
			fileName: "params",
			key:      parser.ParamsKey,
		},
		{
			name:     "Submit Windowed Post",
			txType:   parser.MethodSubmitWindowedPoSt,
			f:        p.submitWindowedPoSt,
			fileName: "params",
			key:      parser.ParamsKey,
		},
		{
			name:     "Withdraw Balance",
			txType:   parser.MethodWithdrawBalance,
			f:        p.parseWithdrawBalance,
			fileName: "params",
			key:      parser.ParamsKey,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, err := loadFile(manifest.MinerKey, tt.txType, tt.fileName)
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
			name:   "Publish Storage Deals",
			txType: parser.MethodPublishStorageDeals,
			f:      p.publishStorageDeals,
		},
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, rawReturn, err := getParmasAndReturn(manifest.MarketKey, tt.txType)
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
