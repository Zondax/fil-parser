package miner_test

import (
	_ "embed"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/actors/miner"
	typesV2 "github.com/zondax/fil-parser/parser/v2/types"
	"github.com/zondax/fil-parser/tools"
)

type testFn func(network string, height int64, rawReturn []byte) (map[string]interface{}, error)

//go:embed expected.json
var expected []byte
var expectedData map[string]any

var network string

func TestMain(m *testing.M) {
	network = "mainnet"
	if err := json.Unmarshal(expected, &expectedData); err != nil {
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
		"ChangeBeneficiary":        miner.ChangeBeneficiary,
		"MinerConstructor":         miner.MinerConstructor,
		"ApplyRewards":             miner.ApplyRewards,
		"OnDeferredCronEvent":      miner.OnDeferredCronEvent,
		"ChangeMultiaddrs":         miner.ChangeMultiaddrs,
		"ChangePeerID":             miner.ChangePeerID,
		"ChangeWorkerAddress":      miner.ChangeWorkerAddress,
		"ChangeOwnerAddress":       miner.ChangeOwnerAddress,
		"GetOwner":                 miner.GetOwner,
		"GetPeerID":                miner.GetPeerID,
		"GetMultiaddrs":            miner.GetMultiaddrs,
		"GetAvailableBalance":      miner.GetAvailableBalance,
		"GetVestingFunds":          miner.GetVestingFunds,
		"ParseWithdrawBalance":     miner.ParseWithdrawBalance,
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
