package miner

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin"
	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"
	miner10 "github.com/filecoin-project/go-state-types/builtin/v10/miner"
	miner11 "github.com/filecoin-project/go-state-types/builtin/v11/miner"
	miner12 "github.com/filecoin-project/go-state-types/builtin/v12/miner"
	miner13 "github.com/filecoin-project/go-state-types/builtin/v13/miner"
	miner14 "github.com/filecoin-project/go-state-types/builtin/v14/miner"
	miner15 "github.com/filecoin-project/go-state-types/builtin/v15/miner"
	miner16 "github.com/filecoin-project/go-state-types/builtin/v16/miner"
	miner8 "github.com/filecoin-project/go-state-types/builtin/v8/miner"
	miner9 "github.com/filecoin-project/go-state-types/builtin/v9/miner"

	builtinv1 "github.com/filecoin-project/specs-actors/actors/builtin"
	legacyv1 "github.com/filecoin-project/specs-actors/actors/builtin/miner"
	builtinv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/miner"
	builtinv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/miner"
	builtinv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/miner"
	builtinv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/miner"
	builtinv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/miner"
	builtinv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/miner"
	cbg "github.com/whyrusleeping/cbor-gen"
	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

// All methods can be found in the Actor.Exports method in
// the correct version package for "github.com/filecoin-project/specs-actors/actors/builtin/miner"

func v1Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	m := &Miner{}
	return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
		builtin.MethodConstructor: {
			Name:   parser.MethodConstructor,
			Method: m.Constructor,
		},
		2: {
			Name:   parser.MethodControlAddresses,
			Method: m.ControlAddresses,
		},
		3: {
			Name:   parser.MethodChangeWorkerAddress,
			Method: m.ChangeWorkerAddressExported,
		},
		4: {
			Name:   parser.MethodChangePeerID,
			Method: m.ChangePeerIDExported,
		},
		5: {
			Name:   parser.MethodSubmitWindowedPoSt,
			Method: m.SubmitWindowedPoSt,
		},
		6: {
			Name:   parser.MethodPreCommitSector,
			Method: m.PreCommitSector,
		},
		7: {
			Name:   parser.MethodProveCommitSector,
			Method: m.ProveCommitSector,
		},
		8: {
			Name:   parser.MethodExtendSectorExpiration,
			Method: m.ExtendSectorExpiration,
		},
		9: {
			Name:   parser.MethodTerminateSectors,
			Method: m.TerminateSectors,
		},
		10: {
			Name:   parser.MethodDeclareFaults,
			Method: m.DeclareFaults,
		},
		11: {
			Name:   parser.MethodDeclareFaultsRecovered,
			Method: m.DeclareFaultsRecovered,
		},
		12: {
			Name:   parser.MethodOnDeferredCronEvent,
			Method: m.OnDeferredCronEvent,
		},
		13: {
			Name:   parser.MethodCheckSectorProven,
			Method: m.CheckSectorProven,
		},
		14: {
			Name:   parser.MethodAddLockedFund,
			Method: m.AddLockedFund,
		},
		15: {
			Name:   parser.MethodReportConsensusFault,
			Method: m.ReportConsensusFault,
		},
		16: {
			Name:   parser.MethodWithdrawBalance,
			Method: m.WithdrawBalanceExported,
		},
		17: {
			Name:   parser.MethodConfirmSectorProofsValid,
			Method: m.ConfirmSectorProofsValid,
		},
		18: {
			Name:   parser.MethodChangeMultiaddrs,
			Method: m.ChangeMultiaddrsExported,
		},
		19: {
			Name:   parser.MethodCompactPartitions,
			Method: m.CompactPartitions,
		},
		20: {
			Name:   parser.MethodCompactSectorNumbers,
			Method: m.CompactSectorNumbers,
		},
	}
}

// method number
func v2Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	m := &Miner{}
	methods := v1Methods()
	// Method 14 changed to ApplyRewards
	methods[14] = nonLegacyBuiltin.MethodMeta{
		Name:   parser.MethodApplyRewards,
		Method: m.ApplyRewards,
	}

	methods[21] = nonLegacyBuiltin.MethodMeta{
		Name:   parser.MethodConfirmUpdateWorkerKey,
		Method: m.ConfirmUpdateWorkerKey,
	}
	methods[22] = nonLegacyBuiltin.MethodMeta{
		Name:   parser.MethodRepayDebt,
		Method: actors.ParseEmptyParamsAndReturn,
	}
	methods[23] = nonLegacyBuiltin.MethodMeta{
		Name:   parser.MethodChangeOwnerAddress,
		Method: m.ChangeOwnerAddressExported,
	}
	return methods
}

func v3Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	m := &Miner{}
	methods := v2Methods()
	methods[24] = nonLegacyBuiltin.MethodMeta{
		Name:   parser.MethodDisputeWindowedPoSt,
		Method: m.DisputeWindowedPoSt,
	}
	return methods
}

func v4Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return v3Methods()
}

func v5Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	m := &Miner{}
	methods := v4Methods()
	methods[25] = nonLegacyBuiltin.MethodMeta{
		Name:   parser.MethodPreCommitSectorBatch,
		Method: m.PreCommitSectorBatch,
	}
	methods[26] = nonLegacyBuiltin.MethodMeta{
		Name:   parser.MethodProveCommitAggregate,
		Method: m.ProveCommitAggregate,
	}
	return methods
}

func v6Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	return v5Methods()
}

func v7Methods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	m := &Miner{}
	methods := v6Methods()
	methods[27] = nonLegacyBuiltin.MethodMeta{
		Name:   parser.MethodProveReplicaUpdates,
		Method: m.ProveReplicaUpdates,
	}
	return methods
}

var terminateSectorsParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.TerminateSectorsParams) },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.TerminateSectorsParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.TerminateSectorsParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.TerminateSectorsParams) },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.TerminateSectorsParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.TerminateSectorsParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.TerminateSectorsParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.TerminateSectorsParams) },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.TerminateSectorsParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.TerminateSectorsParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.TerminateSectorsParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.TerminateSectorsParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.TerminateSectorsParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.TerminateSectorsParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.TerminateSectorsParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.TerminateSectorsParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(miner8.TerminateSectorsParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.TerminateSectorsParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.TerminateSectorsParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.TerminateSectorsParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.TerminateSectorsParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.TerminateSectorsParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.TerminateSectorsParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.TerminateSectorsParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.TerminateSectorsParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.TerminateSectorsParams) },
}

var terminateSectorsReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.TerminateSectorsReturn) },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.TerminateSectorsReturn) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.TerminateSectorsReturn) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.TerminateSectorsReturn) },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.TerminateSectorsReturn) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.TerminateSectorsReturn) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.TerminateSectorsReturn) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.TerminateSectorsReturn) },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.TerminateSectorsReturn) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.TerminateSectorsReturn) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.TerminateSectorsReturn) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.TerminateSectorsReturn) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.TerminateSectorsReturn) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.TerminateSectorsReturn) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.TerminateSectorsReturn) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.TerminateSectorsReturn) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(miner8.TerminateSectorsReturn) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.TerminateSectorsReturn) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.TerminateSectorsReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.TerminateSectorsReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.TerminateSectorsReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.TerminateSectorsReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.TerminateSectorsReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.TerminateSectorsReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.TerminateSectorsReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.TerminateSectorsReturn) },
}

var declareFaultsParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.DeclareFaultsParams) },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.DeclareFaultsParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.DeclareFaultsParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.DeclareFaultsParams) },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.DeclareFaultsParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.DeclareFaultsParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.DeclareFaultsParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.DeclareFaultsParams) },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.DeclareFaultsParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.DeclareFaultsParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.DeclareFaultsParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.DeclareFaultsParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.DeclareFaultsParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.DeclareFaultsParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.DeclareFaultsParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.DeclareFaultsParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(miner8.DeclareFaultsParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.DeclareFaultsParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.DeclareFaultsParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.DeclareFaultsParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.DeclareFaultsParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.DeclareFaultsParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.DeclareFaultsParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.DeclareFaultsParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.DeclareFaultsParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.DeclareFaultsParams) },
}

var declareFaultsRecoveredParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.DeclareFaultsRecoveredParams) },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.DeclareFaultsRecoveredParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.DeclareFaultsRecoveredParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.DeclareFaultsRecoveredParams) },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.DeclareFaultsRecoveredParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.DeclareFaultsRecoveredParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.DeclareFaultsRecoveredParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.DeclareFaultsRecoveredParams) },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.DeclareFaultsRecoveredParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.DeclareFaultsRecoveredParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.DeclareFaultsRecoveredParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.DeclareFaultsRecoveredParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.DeclareFaultsRecoveredParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.DeclareFaultsRecoveredParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.DeclareFaultsRecoveredParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.DeclareFaultsRecoveredParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(miner8.DeclareFaultsRecoveredParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.DeclareFaultsRecoveredParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.DeclareFaultsRecoveredParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.DeclareFaultsRecoveredParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.DeclareFaultsRecoveredParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.DeclareFaultsRecoveredParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.DeclareFaultsRecoveredParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.DeclareFaultsRecoveredParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.DeclareFaultsRecoveredParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.DeclareFaultsRecoveredParams) },
}

var proveReplicaUpdatesParams = map[string]func() cbg.CBORUnmarshaler{
	// SPECIAL CASE:
	// THIS METHOD APPEARS IN V9 BUT THE LIBRARY INTRODUCED IT IN V15
	tools.V9.String():  func() cbg.CBORUnmarshaler { return new(legacyv7.ProveReplicaUpdatesParams) },
	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.ProveReplicaUpdatesParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.ProveReplicaUpdatesParams) },
	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.ProveReplicaUpdatesParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.ProveReplicaUpdatesParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.ProveReplicaUpdatesParams) },

	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.ProveReplicaUpdatesParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(miner8.ProveReplicaUpdatesParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.ProveReplicaUpdatesParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.ProveReplicaUpdatesParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.ProveReplicaUpdatesParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.ProveReplicaUpdatesParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.ProveReplicaUpdatesParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.ProveReplicaUpdatesParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.ProveReplicaUpdatesParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.ProveReplicaUpdatesParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.ProveReplicaUpdatesParams) },
}

var preCommitSectorBatchParams2 = map[string]func() cbg.CBORUnmarshaler{
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.PreCommitSectorBatchParams2) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.PreCommitSectorBatchParams2) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.PreCommitSectorBatchParams2) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.PreCommitSectorBatchParams2) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.PreCommitSectorBatchParams2) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.PreCommitSectorBatchParams2) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.PreCommitSectorBatchParams2) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.PreCommitSectorBatchParams2) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.PreCommitSectorBatchParams2) },
}

var proveReplicaUpdatesParams2 = map[string]func() cbg.CBORUnmarshaler{

	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.ProveReplicaUpdatesParams2) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.ProveReplicaUpdatesParams2) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.ProveReplicaUpdatesParams2) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.ProveReplicaUpdatesParams2) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.ProveReplicaUpdatesParams2) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.ProveReplicaUpdatesParams2) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.ProveReplicaUpdatesParams2) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.ProveReplicaUpdatesParams2) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.ProveReplicaUpdatesParams2) },
}

var proveReplicaUpdates3Params = map[string]func() cbg.CBORUnmarshaler{
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.ProveReplicaUpdates3Params) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.ProveReplicaUpdates3Params) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.ProveReplicaUpdates3Params) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.ProveReplicaUpdates3Params) },
}

var proveReplicaUpdates3Return = map[string]func() cbg.CBORUnmarshaler{

	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.ProveReplicaUpdates3Return) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.ProveReplicaUpdates3Return) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.ProveReplicaUpdates3Return) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.ProveReplicaUpdates3Return) },
}

var proveCommitAggregateParams = map[string]func() cbg.CBORUnmarshaler{

	// SPECIAL CASE:
	// THIS METHOD APPEARS IN V9 BUT THE LIBRARY INTRODUCED IT IN V13
	tools.V9.String():  func() cbg.CBORUnmarshaler { return new(legacyv5.ProveCommitAggregateParams) },
	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.ProveCommitAggregateParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.ProveCommitAggregateParams) },
	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.ProveCommitAggregateParams) },

	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.ProveCommitAggregateParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.ProveCommitAggregateParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.ProveCommitAggregateParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(miner8.ProveCommitAggregateParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.ProveCommitAggregateParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.ProveCommitAggregateParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.ProveCommitAggregateParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.ProveCommitAggregateParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.ProveCommitAggregateParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.ProveCommitAggregateParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.ProveCommitAggregateParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.ProveCommitAggregateParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.ProveCommitAggregateParams) },
}

var disputeWindowedPoStParams = map[string]func() cbg.CBORUnmarshaler{

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.DisputeWindowedPoStParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.DisputeWindowedPoStParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.DisputeWindowedPoStParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.DisputeWindowedPoStParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.DisputeWindowedPoStParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.DisputeWindowedPoStParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(miner8.DisputeWindowedPoStParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.DisputeWindowedPoStParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.DisputeWindowedPoStParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.DisputeWindowedPoStParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.DisputeWindowedPoStParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.DisputeWindowedPoStParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.DisputeWindowedPoStParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.DisputeWindowedPoStParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.DisputeWindowedPoStParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.DisputeWindowedPoStParams) },
}

var reportConsensusFaultParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ReportConsensusFaultParams) },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ReportConsensusFaultParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ReportConsensusFaultParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ReportConsensusFaultParams) },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ReportConsensusFaultParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ReportConsensusFaultParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ReportConsensusFaultParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ReportConsensusFaultParams) },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ReportConsensusFaultParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ReportConsensusFaultParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.ReportConsensusFaultParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.ReportConsensusFaultParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.ReportConsensusFaultParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.ReportConsensusFaultParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.ReportConsensusFaultParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.ReportConsensusFaultParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(miner8.ReportConsensusFaultParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.ReportConsensusFaultParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.ReportConsensusFaultParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.ReportConsensusFaultParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.ReportConsensusFaultParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.ReportConsensusFaultParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.ReportConsensusFaultParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.ReportConsensusFaultParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.ReportConsensusFaultParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.ReportConsensusFaultParams) },
}

var changeBeneficiaryParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.ChangeBeneficiaryParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.ChangeBeneficiaryParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.ChangeBeneficiaryParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.ChangeBeneficiaryParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.ChangeBeneficiaryParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.ChangeBeneficiaryParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.ChangeBeneficiaryParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.ChangeBeneficiaryParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.ChangeBeneficiaryParams) },
}

var minerConstructorParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ConstructorParams) },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ConstructorParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ConstructorParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ConstructorParams) },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ConstructorParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ConstructorParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ConstructorParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ConstructorParams) },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ConstructorParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ConstructorParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.ConstructorParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.ConstructorParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.ConstructorParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.ConstructorParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.ConstructorParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.ConstructorParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(miner8.MinerConstructorParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.MinerConstructorParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.MinerConstructorParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.MinerConstructorParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.MinerConstructorParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.MinerConstructorParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.MinerConstructorParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.MinerConstructorParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.MinerConstructorParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.MinerConstructorParams) },
}

var applyRewardParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return new(abi.TokenAmount) },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(abi.TokenAmount) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(abi.TokenAmount) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(abi.TokenAmount) },
	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(abi.TokenAmount) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(abi.TokenAmount) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(abi.TokenAmount) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(abi.TokenAmount) },

	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(builtinv2.ApplyRewardParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(builtinv2.ApplyRewardParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(builtinv3.ApplyRewardParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(builtinv3.ApplyRewardParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(builtinv4.ApplyRewardParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(builtinv5.ApplyRewardParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(builtinv6.ApplyRewardParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(builtinv7.ApplyRewardParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(miner8.ApplyRewardParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.ApplyRewardParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.ApplyRewardParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.ApplyRewardParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.ApplyRewardParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.ApplyRewardParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.ApplyRewardParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.ApplyRewardParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.ApplyRewardParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.ApplyRewardParams) },
}

var deferredCronEventParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CronEventPayload) },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CronEventPayload) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CronEventPayload) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CronEventPayload) },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CronEventPayload) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CronEventPayload) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CronEventPayload) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CronEventPayload) },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CronEventPayload) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CronEventPayload) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.CronEventPayload) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.CronEventPayload) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.CronEventPayload) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.CronEventPayload) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(builtinv6.DeferredCronEventParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(builtinv7.DeferredCronEventParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(miner8.DeferredCronEventParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.DeferredCronEventParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.DeferredCronEventParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.DeferredCronEventParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.DeferredCronEventParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.DeferredCronEventParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.DeferredCronEventParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.DeferredCronEventParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.DeferredCronEventParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.DeferredCronEventParams) },
}

var changeMultiaddrsParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ChangeMultiaddrsParams) },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ChangeMultiaddrsParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ChangeMultiaddrsParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ChangeMultiaddrsParams) },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ChangeMultiaddrsParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ChangeMultiaddrsParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ChangeMultiaddrsParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ChangeMultiaddrsParams) },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ChangeMultiaddrsParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ChangeMultiaddrsParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.ChangeMultiaddrsParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.ChangeMultiaddrsParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.ChangeMultiaddrsParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.ChangeMultiaddrsParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.ChangeMultiaddrsParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.ChangeMultiaddrsParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(miner8.ChangeMultiaddrsParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.ChangeMultiaddrsParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.ChangeMultiaddrsParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.ChangeMultiaddrsParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.ChangeMultiaddrsParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.ChangeMultiaddrsParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.ChangeMultiaddrsParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.ChangeMultiaddrsParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.ChangeMultiaddrsParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.ChangeMultiaddrsParams) },
}

var changePeerIDParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ChangePeerIDParams) },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ChangePeerIDParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ChangePeerIDParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ChangePeerIDParams) },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ChangePeerIDParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ChangePeerIDParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ChangePeerIDParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ChangePeerIDParams) },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ChangePeerIDParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ChangePeerIDParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.ChangePeerIDParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.ChangePeerIDParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.ChangePeerIDParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.ChangePeerIDParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.ChangePeerIDParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.ChangePeerIDParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(miner8.ChangePeerIDParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.ChangePeerIDParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.ChangePeerIDParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.ChangePeerIDParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.ChangePeerIDParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.ChangePeerIDParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.ChangePeerIDParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.ChangePeerIDParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.ChangePeerIDParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.ChangePeerIDParams) },
}

var changeWorkerAddressParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ChangeWorkerAddressParams) },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ChangeWorkerAddressParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ChangeWorkerAddressParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ChangeWorkerAddressParams) },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ChangeWorkerAddressParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ChangeWorkerAddressParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ChangeWorkerAddressParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ChangeWorkerAddressParams) },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ChangeWorkerAddressParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ChangeWorkerAddressParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.ChangeWorkerAddressParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.ChangeWorkerAddressParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.ChangeWorkerAddressParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.ChangeWorkerAddressParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.ChangeWorkerAddressParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.ChangeWorkerAddressParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(miner8.ChangeWorkerAddressParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.ChangeWorkerAddressParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.ChangeWorkerAddressParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.ChangeWorkerAddressParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.ChangeWorkerAddressParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.ChangeWorkerAddressParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.ChangeWorkerAddressParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.ChangeWorkerAddressParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.ChangeWorkerAddressParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.ChangeWorkerAddressParams) },
}

var isControllingAddressParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.IsControllingAddressParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.IsControllingAddressParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.IsControllingAddressParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.IsControllingAddressParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.IsControllingAddressParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.IsControllingAddressParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.IsControllingAddressParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.IsControllingAddressParams) },
}

var isControllingAddressReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.IsControllingAddressReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.IsControllingAddressReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.IsControllingAddressReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.IsControllingAddressReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.IsControllingAddressReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.IsControllingAddressReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.IsControllingAddressReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.IsControllingAddressReturn) },
}

var getOwnerReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.GetOwnerReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.GetOwnerReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.GetOwnerReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.GetOwnerReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.GetOwnerReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.GetOwnerReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.GetOwnerReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.GetOwnerReturn) },
}

var getPeerIDReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.GetPeerIDReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.GetPeerIDReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.GetPeerIDReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.GetPeerIDReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.GetPeerIDReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.GetPeerIDReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.GetPeerIDReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.GetPeerIDReturn) },
}

var getMultiAddrsReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.GetMultiAddrsReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.GetMultiAddrsReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.GetMultiAddrsReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.GetMultiAddrsReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.GetMultiAddrsReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.GetMultiAddrsReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.GetMultiAddrsReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.GetMultiAddrsReturn) },
}

var getControlAddressesReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.GetControlAddressesReturn) },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.GetControlAddressesReturn) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.GetControlAddressesReturn) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.GetControlAddressesReturn) },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.GetControlAddressesReturn) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.GetControlAddressesReturn) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.GetControlAddressesReturn) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.GetControlAddressesReturn) },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.GetControlAddressesReturn) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.GetControlAddressesReturn) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.GetControlAddressesReturn) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.GetControlAddressesReturn) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.GetControlAddressesReturn) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.GetControlAddressesReturn) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.GetControlAddressesReturn) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.GetControlAddressesReturn) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(miner8.GetControlAddressesReturn) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.GetControlAddressesReturn) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.GetControlAddressesReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.GetControlAddressesReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.GetControlAddressesReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.GetControlAddressesReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.GetControlAddressesReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.GetControlAddressesReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.GetControlAddressesReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.GetControlAddressesReturn) },
}

var getAvailableBalanceReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.GetAvailableBalanceReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.GetAvailableBalanceReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.GetAvailableBalanceReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.GetAvailableBalanceReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.GetAvailableBalanceReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.GetAvailableBalanceReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.GetAvailableBalanceReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.GetAvailableBalanceReturn) },
}

var getVestingFundsReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.GetVestingFundsReturn) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.GetVestingFundsReturn) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.GetVestingFundsReturn) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.GetVestingFundsReturn) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.GetVestingFundsReturn) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.GetVestingFundsReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.GetVestingFundsReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.GetVestingFundsReturn) },
}

var getWithdrawBalanceParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.WithdrawBalanceParams) },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.WithdrawBalanceParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.WithdrawBalanceParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.WithdrawBalanceParams) },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.WithdrawBalanceParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.WithdrawBalanceParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.WithdrawBalanceParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.WithdrawBalanceParams) },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.WithdrawBalanceParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.WithdrawBalanceParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.WithdrawBalanceParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.WithdrawBalanceParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.WithdrawBalanceParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.WithdrawBalanceParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.WithdrawBalanceParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.WithdrawBalanceParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(miner8.WithdrawBalanceParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.WithdrawBalanceParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.WithdrawBalanceParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.WithdrawBalanceParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.WithdrawBalanceParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.WithdrawBalanceParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.WithdrawBalanceParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.WithdrawBalanceParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.WithdrawBalanceParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.WithdrawBalanceParams) },
}

var extendSectorExpiration2Params = map[string]func() cbg.CBORUnmarshaler{
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.ExtendSectorExpiration2Params) },

	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.ExtendSectorExpiration2Params) },
	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.ExtendSectorExpiration2Params) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.ExtendSectorExpiration2Params) },
	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.ExtendSectorExpiration2Params) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.ExtendSectorExpiration2Params) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.ExtendSectorExpiration2Params) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.ExtendSectorExpiration2Params) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.ExtendSectorExpiration2Params) },
}

var preCommitSectorParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.SectorPreCommitInfo) },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.SectorPreCommitInfo) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.SectorPreCommitInfo) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.SectorPreCommitInfo) },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.SectorPreCommitInfo) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.SectorPreCommitInfo) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.SectorPreCommitInfo) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.SectorPreCommitInfo) },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.PreCommitSectorParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.PreCommitSectorParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.PreCommitSectorParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.PreCommitSectorParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.PreCommitSectorParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.PreCommitSectorParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.PreCommitSectorParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.PreCommitSectorParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(miner8.PreCommitSectorParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.PreCommitSectorParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.PreCommitSectorParams) },
	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.PreCommitSectorParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.PreCommitSectorParams) },
	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.PreCommitSectorParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.PreCommitSectorParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.PreCommitSectorParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.PreCommitSectorParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.PreCommitSectorParams) },
}

var proveCommitSectorParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ProveCommitSectorParams) },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ProveCommitSectorParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ProveCommitSectorParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ProveCommitSectorParams) },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ProveCommitSectorParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ProveCommitSectorParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ProveCommitSectorParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ProveCommitSectorParams) },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ProveCommitSectorParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ProveCommitSectorParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.ProveCommitSectorParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.ProveCommitSectorParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.ProveCommitSectorParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.ProveCommitSectorParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.ProveCommitSectorParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.ProveCommitSectorParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(miner8.ProveCommitSectorParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.ProveCommitSectorParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.ProveCommitSectorParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.ProveCommitSectorParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.ProveCommitSectorParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.ProveCommitSectorParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.ProveCommitSectorParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.ProveCommitSectorParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.ProveCommitSectorParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.ProveCommitSectorParams) },
}

var proveCommitSectors3Params = map[string]func() cbg.CBORUnmarshaler{
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.ProveCommitSectors3Params) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.ProveCommitSectors3Params) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.ProveCommitSectors3Params) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.ProveCommitSectors3Params) },
}

var proveCommitSectors3Return = map[string]func() cbg.CBORUnmarshaler{
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.ProveCommitSectors3Return) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.ProveCommitSectors3Return) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.ProveCommitSectors3Return) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.ProveCommitSectors3Return) },
}

var internalSectorSetupForPresealParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.InternalSectorSetupForPresealParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.InternalSectorSetupForPresealParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.InternalSectorSetupForPresealParams) },
}

var submitWindowedPoStParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.SubmitWindowedPoStParams) },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.SubmitWindowedPoStParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.SubmitWindowedPoStParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.SubmitWindowedPoStParams) },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.SubmitWindowedPoStParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.SubmitWindowedPoStParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.SubmitWindowedPoStParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.SubmitWindowedPoStParams) },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.SubmitWindowedPoStParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.SubmitWindowedPoStParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.SubmitWindowedPoStParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.SubmitWindowedPoStParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.SubmitWindowedPoStParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.SubmitWindowedPoStParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.SubmitWindowedPoStParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.SubmitWindowedPoStParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(miner8.SubmitWindowedPoStParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.SubmitWindowedPoStParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.SubmitWindowedPoStParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.SubmitWindowedPoStParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.SubmitWindowedPoStParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.SubmitWindowedPoStParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.SubmitWindowedPoStParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.SubmitWindowedPoStParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.SubmitWindowedPoStParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.SubmitWindowedPoStParams) },
}

var confirmSectorProofsParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return new(builtinv1.ConfirmSectorProofsParams) },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(builtinv1.ConfirmSectorProofsParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(builtinv1.ConfirmSectorProofsParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(builtinv1.ConfirmSectorProofsParams) },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(builtinv2.ConfirmSectorProofsParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(builtinv2.ConfirmSectorProofsParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(builtinv2.ConfirmSectorProofsParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(builtinv2.ConfirmSectorProofsParams) },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(builtinv2.ConfirmSectorProofsParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(builtinv2.ConfirmSectorProofsParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(builtinv3.ConfirmSectorProofsParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(builtinv3.ConfirmSectorProofsParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(builtinv4.ConfirmSectorProofsParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(builtinv5.ConfirmSectorProofsParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(builtinv6.ConfirmSectorProofsParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(builtinv7.ConfirmSectorProofsParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(miner8.ConfirmSectorProofsParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.ConfirmSectorProofsParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.ConfirmSectorProofsParams) },
	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.ConfirmSectorProofsParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.ConfirmSectorProofsParams) },
	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.ConfirmSectorProofsParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.ConfirmSectorProofsParams) },
}

var checkSectorProvenParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CheckSectorProvenParams) },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CheckSectorProvenParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CheckSectorProvenParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CheckSectorProvenParams) },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CheckSectorProvenParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CheckSectorProvenParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CheckSectorProvenParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CheckSectorProvenParams) },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CheckSectorProvenParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CheckSectorProvenParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.CheckSectorProvenParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.CheckSectorProvenParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.CheckSectorProvenParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.CheckSectorProvenParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.CheckSectorProvenParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.CheckSectorProvenParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(miner8.CheckSectorProvenParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.CheckSectorProvenParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.CheckSectorProvenParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.CheckSectorProvenParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.CheckSectorProvenParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.CheckSectorProvenParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.CheckSectorProvenParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.CheckSectorProvenParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.CheckSectorProvenParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.CheckSectorProvenParams) },
}

var extendSectorExpirationParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ExtendSectorExpirationParams) },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ExtendSectorExpirationParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ExtendSectorExpirationParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.ExtendSectorExpirationParams) },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ExtendSectorExpirationParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ExtendSectorExpirationParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ExtendSectorExpirationParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ExtendSectorExpirationParams) },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ExtendSectorExpirationParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.ExtendSectorExpirationParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.ExtendSectorExpirationParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.ExtendSectorExpirationParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.ExtendSectorExpirationParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.ExtendSectorExpirationParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.ExtendSectorExpirationParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.ExtendSectorExpirationParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(miner8.ExtendSectorExpirationParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.ExtendSectorExpirationParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.ExtendSectorExpirationParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.ExtendSectorExpirationParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.ExtendSectorExpirationParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.ExtendSectorExpirationParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.ExtendSectorExpirationParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.ExtendSectorExpirationParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.ExtendSectorExpirationParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.ExtendSectorExpirationParams) },
}

var compactSectorNumbersParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V0.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CompactSectorNumbersParams) },
	tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CompactSectorNumbersParams) },
	tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CompactSectorNumbersParams) },
	tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CompactSectorNumbersParams) },

	tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CompactSectorNumbersParams) },
	tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CompactSectorNumbersParams) },
	tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CompactSectorNumbersParams) },
	tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CompactSectorNumbersParams) },
	tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CompactSectorNumbersParams) },
	tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CompactSectorNumbersParams) },

	tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.CompactSectorNumbersParams) },
	tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.CompactSectorNumbersParams) },

	tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.CompactSectorNumbersParams) },
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.CompactSectorNumbersParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.CompactSectorNumbersParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.CompactSectorNumbersParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(miner8.CompactSectorNumbersParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.CompactSectorNumbersParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.CompactSectorNumbersParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.CompactSectorNumbersParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.CompactSectorNumbersParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.CompactSectorNumbersParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.CompactSectorNumbersParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.CompactSectorNumbersParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.CompactSectorNumbersParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.CompactSectorNumbersParams) },
}

func compactPartitionsParams() map[string]func() cbg.CBORUnmarshaler {
	return map[string]func() cbg.CBORUnmarshaler{
		tools.V0.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CompactPartitionsParams) },
		tools.V1.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CompactPartitionsParams) },
		tools.V2.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CompactPartitionsParams) },
		tools.V3.String(): func() cbg.CBORUnmarshaler { return new(legacyv1.CompactPartitionsParams) },

		tools.V4.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CompactPartitionsParams) },
		tools.V5.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CompactPartitionsParams) },
		tools.V6.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CompactPartitionsParams) },
		tools.V7.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CompactPartitionsParams) },
		tools.V8.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CompactPartitionsParams) },
		tools.V9.String(): func() cbg.CBORUnmarshaler { return new(legacyv2.CompactPartitionsParams) },

		tools.V10.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.CompactPartitionsParams) },
		tools.V11.String(): func() cbg.CBORUnmarshaler { return new(legacyv3.CompactPartitionsParams) },

		tools.V12.String(): func() cbg.CBORUnmarshaler { return new(legacyv4.CompactPartitionsParams) },
		tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.CompactPartitionsParams) },
		tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.CompactPartitionsParams) },
		tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.CompactPartitionsParams) },
		tools.V16.String(): func() cbg.CBORUnmarshaler { return new(miner8.CompactPartitionsParams) },
		tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.CompactPartitionsParams) },
		tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.CompactPartitionsParams) },

		tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.CompactPartitionsParams) },
		tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.CompactPartitionsParams) },

		tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.CompactPartitionsParams) },
		tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.CompactPartitionsParams) },
		tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.CompactPartitionsParams) },
		tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.CompactPartitionsParams) },
		tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.CompactPartitionsParams) },
	}
}

var preCommitSectorBatchParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V13.String(): func() cbg.CBORUnmarshaler { return new(legacyv5.PreCommitSectorBatchParams) },
	tools.V14.String(): func() cbg.CBORUnmarshaler { return new(legacyv6.PreCommitSectorBatchParams) },
	tools.V15.String(): func() cbg.CBORUnmarshaler { return new(legacyv7.PreCommitSectorBatchParams) },
	tools.V16.String(): func() cbg.CBORUnmarshaler { return new(miner8.PreCommitSectorBatchParams) },
	tools.V17.String(): func() cbg.CBORUnmarshaler { return new(miner9.PreCommitSectorBatchParams) },
	tools.V18.String(): func() cbg.CBORUnmarshaler { return new(miner10.PreCommitSectorBatchParams) },

	tools.V19.String(): func() cbg.CBORUnmarshaler { return new(miner11.PreCommitSectorBatchParams) },
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(miner11.PreCommitSectorBatchParams) },

	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(miner12.PreCommitSectorBatchParams) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(miner13.PreCommitSectorBatchParams) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.PreCommitSectorBatchParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.PreCommitSectorBatchParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.PreCommitSectorBatchParams) },
}

var proveCommitSectorsNIParams = map[string]func() cbg.CBORUnmarshaler{
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.ProveCommitSectorsNIParams) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.ProveCommitSectorsNIParams) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.ProveCommitSectorsNIParams) },
}

var proveCommitSectorsNIReturn = map[string]func() cbg.CBORUnmarshaler{
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(miner14.ProveCommitSectorsNIReturn) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(miner15.ProveCommitSectorsNIReturn) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(miner16.ProveCommitSectorsNIReturn) },
}
