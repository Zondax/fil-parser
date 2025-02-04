package miner_test

import (
	"encoding/json"
	"os"
	"testing"

	miner10 "github.com/filecoin-project/go-state-types/builtin/v10/miner"
	miner11 "github.com/filecoin-project/go-state-types/builtin/v11/miner"
	miner12 "github.com/filecoin-project/go-state-types/builtin/v12/miner"
	miner13 "github.com/filecoin-project/go-state-types/builtin/v13/miner"
	miner14 "github.com/filecoin-project/go-state-types/builtin/v14/miner"
	miner15 "github.com/filecoin-project/go-state-types/builtin/v15/miner"
	miner8 "github.com/filecoin-project/go-state-types/builtin/v8/miner"
	miner9 "github.com/filecoin-project/go-state-types/builtin/v9/miner"
	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/miner"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/miner"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/miner"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/miner"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/miner"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/miner"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	v2 "github.com/zondax/fil-parser/actors/v2"
	"github.com/zondax/fil-parser/actors/v2/miner"
	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
	"github.com/zondax/fil-parser/tools"
)

type testFn func(network string, height int64, rawReturn []byte) (map[string]interface{}, error)

var expected []byte
var expectedData map[string]any

var network string

func TestMain(m *testing.M) {
	network = "mainnet"
	if err := json.Unmarshal(expected, &expectedData); err != nil {
		panic(err)
	}
	var err error
	expected, err = tools.ReadActorSnapshot()
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestMiner(t *testing.T) {
	miner := &miner.Miner{}
	testFns := map[string]testFn{
		"DeclareFaults":            miner.DeclareFaults,
		"DeclareFaultsRecovered":   miner.DeclareFaultsRecovered,
		"ProveReplicaUpdates":      miner.ProveReplicaUpdates,
		"PreCommitSectorBatch2":    miner.PreCommitSectorBatch2,
		"ProveCommitAggregate":     miner.ProveCommitAggregate,
		"DisputeWindowedPoSt":      miner.DisputeWindowedPoSt,
		"ReportConsensusFault":     miner.ReportConsensusFault,
		"ChangeBeneficiary":        miner.ChangeBeneficiaryExported,
		"MinerConstructor":         miner.Constructor,
		"ApplyRewards":             miner.ApplyRewards,
		"OnDeferredCronEvent":      miner.OnDeferredCronEvent,
		"ChangeMultiaddrs":         miner.ChangeMultiaddrsExported,
		"ChangePeerID":             miner.ChangePeerIDExported,
		"ChangeWorkerAddress":      miner.ChangeWorkerAddressExported,
		"ChangeOwnerAddress":       miner.ChangeOwnerAddressExported,
		"GetOwner":                 miner.GetOwnerExported,
		"GetPeerID":                miner.GetPeerIDExported,
		"GetMultiaddrs":            miner.GetMultiaddrsExported,
		"GetAvailableBalance":      miner.GetAvailableBalanceExported,
		"GetVestingFunds":          miner.GetVestingFundsExported,
		"ParseWithdrawBalance":     miner.WithdrawBalanceExported,
		"ExtendSectorExpiration2":  miner.ExtendSectorExpiration2,
		"PreCommitSector":          miner.PreCommitSector,
		"ProveCommitSector":        miner.ProveCommitSector,
		"SubmitWindowedPoSt":       miner.SubmitWindowedPoSt,
		"ConfirmSectorProofsValid": miner.ConfirmSectorProofsValid,
		"CheckSectorProven":        miner.CheckSectorProven,
		"ExtendSectorExpiration":   miner.ExtendSectorExpiration,
		"CompactSectorNumbers":     miner.CompactSectorNumbers,
		"CompactPartitions":        miner.CompactPartitions,
		"PreCommitSectorBatch":     miner.PreCommitSectorBatch,
		"GetSectorSize":            miner.GetSectorSize,
	}
	for name, fn := range testFns {
		t.Run(name, func(t *testing.T) {
			tests, err := tools.LoadTestData[map[string]any](network, name, expectedData)
			require.NoError(t, err)
			runTest(t, fn, tests)
		})
	}
}

func TestTerminateSectors(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "TerminateSectors", expectedData)
	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.Height, tt.Version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}
				miner := &miner.Miner{}
				result, err := miner.TerminateSectors(tt.Network, tt.Height, trace.Msg.Params, trace.MsgRct.Return)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.Expected))

			}
		})
	}
}

func TestProveReplicaUpdates2(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "ProveReplicaUpdates2", expectedData)
	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.Height, tt.Version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}

				miner := &miner.Miner{}
				result, err := miner.ProveReplicaUpdates2(tt.Network, tt.Height, trace.Msg.Params, trace.MsgRct.Return)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.Expected))

			}
		})
	}
}

func TestIsControllingAddressExported(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "IsControllingAddressExported", expectedData)
	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.Height, tt.Version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}
				miner := &miner.Miner{}
				result, err := miner.IsControllingAddressExported(tt.Network, tt.Height, trace.Msg.Params, trace.MsgRct.Return)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.Expected))

			}
		})
	}
}

func TestMethodCoverage(t *testing.T) {
	m := &miner.Miner{}

	actorVersions := []any{
		legacyv2.Actor{},
		legacyv3.Actor{},
		legacyv4.Actor{},
		legacyv5.Actor{},
		legacyv6.Actor{},
		legacyv7.Actor{},
		miner8.Methods,
		miner9.Methods,
		miner10.Methods,
		miner11.Methods,
		miner12.Methods,
		miner13.Methods,
		miner14.Methods,
		miner15.Methods,
	}

	missingMethods := v2.MissingMethods(m, actorVersions)
	assert.Empty(t, missingMethods, "missing methods: %v", missingMethods)
}

func runTest(t *testing.T, fn testFn, tests []tools.TestCase[map[string]any]) {
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.Height, tt.Version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}

				result, err := fn(tt.Network, tt.Height, trace.MsgRct.Return)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.Expected))

			}
		})
	}
}

func TestProveCommitSectors3(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "ProveCommitSectors3", expectedData)
	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			computeState, err := tools.ComputeState[typesV2.ComputeStateOutputV2](tt.Height, tt.Version)
			require.NoError(t, err)

			for _, trace := range computeState.Trace {
				if trace.Msg == nil {
					continue
				}
				miner := &miner.Miner{}
				result, err := miner.ProveCommitSectors3(tt.Network, tt.Height, trace.Msg.Params, trace.MsgRct.Return)
				require.NoError(t, err)
				require.True(t, tools.CompareResult(result, tt.Expected))

			}
		})
	}
}
