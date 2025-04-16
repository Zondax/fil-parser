package miner

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/zondax/golem/pkg/logger"

	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/go-state-types/abi"
	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/manifest"
	legacyBuiltin "github.com/filecoin-project/specs-actors/actors/builtin"

	miner10 "github.com/filecoin-project/go-state-types/builtin/v10/miner"
	miner11 "github.com/filecoin-project/go-state-types/builtin/v11/miner"
	miner12 "github.com/filecoin-project/go-state-types/builtin/v12/miner"
	miner13 "github.com/filecoin-project/go-state-types/builtin/v13/miner"
	miner14 "github.com/filecoin-project/go-state-types/builtin/v14/miner"
	miner15 "github.com/filecoin-project/go-state-types/builtin/v15/miner"
	miner16 "github.com/filecoin-project/go-state-types/builtin/v16/miner"
	miner8 "github.com/filecoin-project/go-state-types/builtin/v8/miner"
	miner9 "github.com/filecoin-project/go-state-types/builtin/v9/miner"

	legacyv1 "github.com/filecoin-project/specs-actors/actors/builtin/miner"
	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/miner"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/miner"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/miner"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/miner"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/miner"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/miner"

	builtinv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	builtinv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin"
	builtinv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
	builtinv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin"
	builtinv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin"
	builtinv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin"

	cbg "github.com/whyrusleeping/cbor-gen"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

type Miner struct {
	logger *logger.Logger
}

func New(logger *logger.Logger) *Miner {
	return &Miner{
		logger: logger,
	}
}

func (m *Miner) Name() string {
	return manifest.MinerKey
}

func (*Miner) StartNetworkHeight() int64 {
	return tools.V1.Height()
}

func terminateSectorsParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.TerminateSectorsParams{},

		tools.V8.String(): &legacyv2.TerminateSectorsParams{},
		tools.V9.String(): &legacyv2.TerminateSectorsParams{},

		tools.V10.String(): &legacyv3.TerminateSectorsParams{},
		tools.V11.String(): &legacyv3.TerminateSectorsParams{},

		tools.V12.String(): &legacyv4.TerminateSectorsParams{},
		tools.V13.String(): &legacyv5.TerminateSectorsParams{},
		tools.V14.String(): &legacyv6.TerminateSectorsParams{},
		tools.V15.String(): &legacyv7.TerminateSectorsParams{},
		tools.V16.String(): &miner8.TerminateSectorsParams{},
		tools.V17.String(): &miner9.TerminateSectorsParams{},
		tools.V18.String(): &miner10.TerminateSectorsParams{},

		tools.V19.String(): &miner11.TerminateSectorsParams{},
		tools.V20.String(): &miner11.TerminateSectorsParams{},

		tools.V21.String(): &miner12.TerminateSectorsParams{},
		tools.V22.String(): &miner13.TerminateSectorsParams{},
		tools.V23.String(): &miner14.TerminateSectorsParams{},
		tools.V24.String(): &miner15.TerminateSectorsParams{},
		tools.V25.String(): &miner16.TerminateSectorsParams{},
	}
}

func terminateSectorsReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.TerminateSectorsReturn{},

		tools.V8.String(): &legacyv2.TerminateSectorsReturn{},
		tools.V9.String(): &legacyv2.TerminateSectorsReturn{},

		tools.V10.String(): &legacyv3.TerminateSectorsReturn{},
		tools.V11.String(): &legacyv3.TerminateSectorsReturn{},

		tools.V12.String(): &legacyv4.TerminateSectorsReturn{},
		tools.V13.String(): &legacyv5.TerminateSectorsReturn{},
		tools.V14.String(): &legacyv6.TerminateSectorsReturn{},
		tools.V15.String(): &legacyv7.TerminateSectorsReturn{},
		tools.V16.String(): &miner8.TerminateSectorsReturn{},
		tools.V17.String(): &miner9.TerminateSectorsReturn{},
		tools.V18.String(): &miner10.TerminateSectorsReturn{},

		tools.V19.String(): &miner11.TerminateSectorsReturn{},
		tools.V20.String(): &miner11.TerminateSectorsReturn{},

		tools.V21.String(): &miner12.TerminateSectorsReturn{},
		tools.V22.String(): &miner13.TerminateSectorsReturn{},
		tools.V23.String(): &miner14.TerminateSectorsReturn{},
		tools.V24.String(): &miner15.TerminateSectorsReturn{},
		tools.V25.String(): &miner16.TerminateSectorsReturn{},
	}
}

func declareFaultsParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.DeclareFaultsParams{},

		tools.V8.String(): &legacyv2.DeclareFaultsParams{},
		tools.V9.String(): &legacyv2.DeclareFaultsParams{},

		tools.V10.String(): &legacyv3.DeclareFaultsParams{},
		tools.V11.String(): &legacyv3.DeclareFaultsParams{},

		tools.V12.String(): &legacyv4.DeclareFaultsParams{},
		tools.V13.String(): &legacyv5.DeclareFaultsParams{},
		tools.V14.String(): &legacyv6.DeclareFaultsParams{},
		tools.V15.String(): &legacyv7.DeclareFaultsParams{},
		tools.V16.String(): &miner8.DeclareFaultsParams{},
		tools.V17.String(): &miner9.DeclareFaultsParams{},
		tools.V18.String(): &miner10.DeclareFaultsParams{},

		tools.V19.String(): &miner11.DeclareFaultsParams{},
		tools.V20.String(): &miner11.DeclareFaultsParams{},

		tools.V21.String(): &miner12.DeclareFaultsParams{},
		tools.V22.String(): &miner13.DeclareFaultsParams{},
		tools.V23.String(): &miner14.DeclareFaultsParams{},
		tools.V24.String(): &miner15.DeclareFaultsParams{},
		tools.V25.String(): &miner16.DeclareFaultsParams{},
	}
}

func declareFaultsRecoveredParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.DeclareFaultsRecoveredParams{},

		tools.V8.String(): &legacyv2.DeclareFaultsRecoveredParams{},
		tools.V9.String(): &legacyv2.DeclareFaultsRecoveredParams{},

		tools.V10.String(): &legacyv3.DeclareFaultsRecoveredParams{},
		tools.V11.String(): &legacyv3.DeclareFaultsRecoveredParams{},

		tools.V12.String(): &legacyv4.DeclareFaultsRecoveredParams{},
		tools.V13.String(): &legacyv5.DeclareFaultsRecoveredParams{},
		tools.V14.String(): &legacyv6.DeclareFaultsRecoveredParams{},
		tools.V15.String(): &legacyv7.DeclareFaultsRecoveredParams{},
		tools.V16.String(): &miner8.DeclareFaultsRecoveredParams{},
		tools.V17.String(): &miner9.DeclareFaultsRecoveredParams{},
		tools.V18.String(): &miner10.DeclareFaultsRecoveredParams{},

		tools.V19.String(): &miner11.DeclareFaultsRecoveredParams{},
		tools.V20.String(): &miner11.DeclareFaultsRecoveredParams{},

		tools.V21.String(): &miner12.DeclareFaultsRecoveredParams{},
		tools.V22.String(): &miner13.DeclareFaultsRecoveredParams{},
		tools.V23.String(): &miner14.DeclareFaultsRecoveredParams{},
		tools.V24.String(): &miner15.DeclareFaultsRecoveredParams{},
		tools.V25.String(): &miner16.DeclareFaultsRecoveredParams{},
	}
}

func proveReplicaUpdatesParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V15.String(): &legacyv7.ProveReplicaUpdatesParams{},
		tools.V16.String(): &miner8.ProveReplicaUpdatesParams{},
		tools.V17.String(): &miner9.ProveReplicaUpdatesParams{},
		tools.V18.String(): &miner10.ProveReplicaUpdatesParams{},

		tools.V19.String(): &miner11.ProveReplicaUpdatesParams{},
		tools.V20.String(): &miner11.ProveReplicaUpdatesParams{},

		tools.V21.String(): &miner12.ProveReplicaUpdatesParams{},
		tools.V22.String(): &miner13.ProveReplicaUpdatesParams{},
		tools.V23.String(): &miner14.ProveReplicaUpdatesParams{},
		tools.V24.String(): &miner15.ProveReplicaUpdatesParams{},
		tools.V25.String(): &miner16.ProveReplicaUpdatesParams{},
	}
}

func preCommitSectorBatchParams2() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V17.String(): &miner9.PreCommitSectorBatchParams2{},
		tools.V18.String(): &miner10.PreCommitSectorBatchParams2{},

		tools.V19.String(): &miner11.PreCommitSectorBatchParams2{},
		tools.V20.String(): &miner11.PreCommitSectorBatchParams2{},

		tools.V21.String(): &miner12.PreCommitSectorBatchParams2{},
		tools.V22.String(): &miner13.PreCommitSectorBatchParams2{},
		tools.V23.String(): &miner14.PreCommitSectorBatchParams2{},
		tools.V24.String(): &miner15.PreCommitSectorBatchParams2{},
		tools.V25.String(): &miner16.PreCommitSectorBatchParams2{},
	}
}

func proveReplicaUpdatesParams2() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V17.String(): &miner9.ProveReplicaUpdatesParams2{},
		tools.V18.String(): &miner10.ProveReplicaUpdatesParams2{},

		tools.V19.String(): &miner11.ProveReplicaUpdatesParams2{},
		tools.V20.String(): &miner11.ProveReplicaUpdatesParams2{},

		tools.V21.String(): &miner12.ProveReplicaUpdatesParams2{},
		tools.V22.String(): &miner13.ProveReplicaUpdatesParams2{},
		tools.V23.String(): &miner14.ProveReplicaUpdatesParams2{},
		tools.V24.String(): &miner15.ProveReplicaUpdatesParams2{},
		tools.V25.String(): &miner16.ProveReplicaUpdatesParams2{},
	}
}

func proveReplicaUpdates3Params() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V22.String(): &miner13.ProveReplicaUpdates3Params{},
		tools.V23.String(): &miner14.ProveReplicaUpdates3Params{},
		tools.V24.String(): &miner15.ProveReplicaUpdates3Params{},
		tools.V25.String(): &miner16.ProveReplicaUpdates3Params{},
	}
}

func proveReplicaUpdates3Return() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V22.String(): &miner13.ProveReplicaUpdates3Return{},
		tools.V23.String(): &miner14.ProveReplicaUpdates3Return{},
		tools.V24.String(): &miner15.ProveReplicaUpdates3Return{},
		tools.V25.String(): &miner16.ProveReplicaUpdates3Return{},
	}
}

func proveCommitAggregateParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V13.String(): &legacyv5.ProveCommitAggregateParams{},
		tools.V14.String(): &legacyv6.ProveCommitAggregateParams{},
		tools.V15.String(): &legacyv7.ProveCommitAggregateParams{},
		tools.V16.String(): &miner8.ProveCommitAggregateParams{},
		tools.V17.String(): &miner9.ProveCommitAggregateParams{},
		tools.V18.String(): &miner10.ProveCommitAggregateParams{},

		tools.V19.String(): &miner11.ProveCommitAggregateParams{},
		tools.V20.String(): &miner11.ProveCommitAggregateParams{},

		tools.V21.String(): &miner12.ProveCommitAggregateParams{},
		tools.V22.String(): &miner13.ProveCommitAggregateParams{},
		tools.V23.String(): &miner14.ProveCommitAggregateParams{},
		tools.V24.String(): &miner15.ProveCommitAggregateParams{},
		tools.V25.String(): &miner16.ProveCommitAggregateParams{},
	}
}

func disputeWindowedPoStParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V10.String(): &legacyv3.DisputeWindowedPoStParams{},
		tools.V11.String(): &legacyv3.DisputeWindowedPoStParams{},

		tools.V12.String(): &legacyv4.DisputeWindowedPoStParams{},
		tools.V13.String(): &legacyv5.DisputeWindowedPoStParams{},
		tools.V14.String(): &legacyv6.DisputeWindowedPoStParams{},
		tools.V15.String(): &legacyv7.DisputeWindowedPoStParams{},
		tools.V16.String(): &miner8.DisputeWindowedPoStParams{},
		tools.V17.String(): &miner9.DisputeWindowedPoStParams{},
		tools.V18.String(): &miner10.DisputeWindowedPoStParams{},

		tools.V19.String(): &miner11.DisputeWindowedPoStParams{},
		tools.V20.String(): &miner11.DisputeWindowedPoStParams{},

		tools.V21.String(): &miner12.DisputeWindowedPoStParams{},
		tools.V22.String(): &miner13.DisputeWindowedPoStParams{},
		tools.V23.String(): &miner14.DisputeWindowedPoStParams{},
		tools.V24.String(): &miner15.DisputeWindowedPoStParams{},
		tools.V25.String(): &miner16.DisputeWindowedPoStParams{},
	}
}

func reportConsensusFaultParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.ReportConsensusFaultParams{},

		tools.V8.String(): &legacyv2.ReportConsensusFaultParams{},
		tools.V9.String(): &legacyv2.ReportConsensusFaultParams{},

		tools.V10.String(): &legacyv3.ReportConsensusFaultParams{},
		tools.V11.String(): &legacyv3.ReportConsensusFaultParams{},

		tools.V12.String(): &legacyv4.ReportConsensusFaultParams{},
		tools.V13.String(): &legacyv5.ReportConsensusFaultParams{},
		tools.V14.String(): &legacyv6.ReportConsensusFaultParams{},
		tools.V15.String(): &legacyv7.ReportConsensusFaultParams{},
		tools.V16.String(): &miner8.ReportConsensusFaultParams{},
		tools.V17.String(): &miner9.ReportConsensusFaultParams{},
		tools.V18.String(): &miner10.ReportConsensusFaultParams{},

		tools.V19.String(): &miner11.ReportConsensusFaultParams{},
		tools.V20.String(): &miner11.ReportConsensusFaultParams{},

		tools.V21.String(): &miner12.ReportConsensusFaultParams{},
		tools.V22.String(): &miner13.ReportConsensusFaultParams{},
		tools.V23.String(): &miner14.ReportConsensusFaultParams{},
		tools.V24.String(): &miner15.ReportConsensusFaultParams{},
		tools.V25.String(): &miner16.ReportConsensusFaultParams{},
	}
}

func changeBeneficiaryParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V17.String(): &miner9.ChangeBeneficiaryParams{},
		tools.V18.String(): &miner10.ChangeBeneficiaryParams{},

		tools.V19.String(): &miner11.ChangeBeneficiaryParams{},
		tools.V20.String(): &miner11.ChangeBeneficiaryParams{},

		tools.V21.String(): &miner12.ChangeBeneficiaryParams{},
		tools.V22.String(): &miner13.ChangeBeneficiaryParams{},
		tools.V23.String(): &miner14.ChangeBeneficiaryParams{},
		tools.V24.String(): &miner15.ChangeBeneficiaryParams{},
		tools.V25.String(): &miner16.ChangeBeneficiaryParams{},
	}
}

func minerConstructorParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.ConstructorParams{},

		tools.V8.String(): &legacyv2.ConstructorParams{},
		tools.V9.String(): &legacyv2.ConstructorParams{},

		tools.V10.String(): &legacyv3.ConstructorParams{},
		tools.V11.String(): &legacyv3.ConstructorParams{},

		tools.V12.String(): &legacyv4.ConstructorParams{},
		tools.V13.String(): &legacyv5.ConstructorParams{},
		tools.V14.String(): &legacyv6.ConstructorParams{},
		tools.V15.String(): &legacyv7.ConstructorParams{},
		tools.V16.String(): &miner8.MinerConstructorParams{},
		tools.V17.String(): &miner9.MinerConstructorParams{},
		tools.V18.String(): &miner10.MinerConstructorParams{},

		tools.V19.String(): &miner11.MinerConstructorParams{},
		tools.V20.String(): &miner11.MinerConstructorParams{},

		tools.V21.String(): &miner12.MinerConstructorParams{},
		tools.V22.String(): &miner13.MinerConstructorParams{},
		tools.V23.String(): &miner14.MinerConstructorParams{},
		tools.V24.String(): &miner15.MinerConstructorParams{},
		tools.V25.String(): &miner16.MinerConstructorParams{},
	}
}

func applyRewardParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &abi.TokenAmount{},

		tools.V8.String(): &builtinv2.ApplyRewardParams{},
		tools.V9.String(): &builtinv2.ApplyRewardParams{},

		tools.V10.String(): &builtinv3.ApplyRewardParams{},
		tools.V11.String(): &builtinv3.ApplyRewardParams{},

		tools.V12.String(): &builtinv4.ApplyRewardParams{},
		tools.V13.String(): &builtinv5.ApplyRewardParams{},
		tools.V14.String(): &builtinv6.ApplyRewardParams{},
		tools.V15.String(): &builtinv7.ApplyRewardParams{},
		tools.V16.String(): &miner8.ApplyRewardParams{},
		tools.V17.String(): &miner9.ApplyRewardParams{},
		tools.V18.String(): &miner10.ApplyRewardParams{},

		tools.V19.String(): &miner11.ApplyRewardParams{},
		tools.V20.String(): &miner11.ApplyRewardParams{},

		tools.V21.String(): &miner12.ApplyRewardParams{},
		tools.V22.String(): &miner13.ApplyRewardParams{},
		tools.V23.String(): &miner14.ApplyRewardParams{},
		tools.V24.String(): &miner15.ApplyRewardParams{},
		tools.V25.String(): &miner16.ApplyRewardParams{},
	}
}

func deferredCronEventParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.CronEventPayload{},

		tools.V8.String(): &legacyv2.CronEventPayload{},
		tools.V9.String(): &legacyv2.CronEventPayload{},

		tools.V10.String(): &legacyv3.CronEventPayload{},
		tools.V11.String(): &legacyv3.CronEventPayload{},

		tools.V12.String(): &legacyv4.CronEventPayload{},
		tools.V13.String(): &legacyv5.CronEventPayload{},
		tools.V14.String(): &builtinv6.DeferredCronEventParams{},
		tools.V15.String(): &builtinv7.DeferredCronEventParams{},
		tools.V16.String(): &miner8.DeferredCronEventParams{},
		tools.V17.String(): &miner9.DeferredCronEventParams{},
		tools.V18.String(): &miner10.DeferredCronEventParams{},

		tools.V19.String(): &miner11.DeferredCronEventParams{},
		tools.V20.String(): &miner11.DeferredCronEventParams{},

		tools.V21.String(): &miner12.DeferredCronEventParams{},
		tools.V22.String(): &miner13.DeferredCronEventParams{},
		tools.V23.String(): &miner14.DeferredCronEventParams{},
		tools.V24.String(): &miner15.DeferredCronEventParams{},
		tools.V25.String(): &miner16.DeferredCronEventParams{},
	}
}

// implemented in the rust builtin-actors but not the golang version
var initialPledgeMethodNum = abi.MethodNum(nonLegacyBuiltin.MustGenerateFRCMethodNum(parser.MethodInitialPledge))

func (m *Miner) Methods(_ context.Context, network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error) {
	var data map[abi.MethodNum]nonLegacyBuiltin.MethodMeta
	switch {
	// all legacy version
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V15)...):
		data = map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
			legacyBuiltin.MethodsMiner.Constructor: {
				Name:   parser.MethodConstructor,
				Method: m.Constructor,
			},
			legacyBuiltin.MethodsMiner.ControlAddresses: {
				Name:   parser.MethodControlAddresses,
				Method: m.ControlAddresses,
			},
			legacyBuiltin.MethodsMiner.ChangeWorkerAddress: {
				Name:   parser.MethodChangeWorkerAddress,
				Method: m.ChangeWorkerAddressExported,
			},
			legacyBuiltin.MethodsMiner.ChangePeerID: {
				Name:   parser.MethodChangePeerID,
				Method: m.ChangePeerIDExported,
			},
			legacyBuiltin.MethodsMiner.SubmitWindowedPoSt: {
				Name:   parser.MethodSubmitWindowedPoSt,
				Method: m.SubmitWindowedPoSt,
			},
			legacyBuiltin.MethodsMiner.PreCommitSector: {
				Name:   parser.MethodPreCommitSector,
				Method: m.PreCommitSector,
			},
			legacyBuiltin.MethodsMiner.ProveCommitSector: {
				Name:   parser.MethodProveCommitSector,
				Method: m.ProveCommitSector,
			},
			nonLegacyBuiltin.MethodsMiner.ExtendSectorExpiration: {
				Name:   parser.MethodExtendSectorExpiration,
				Method: m.ExtendSectorExpiration,
			},
			legacyBuiltin.MethodsMiner.TerminateSectors: {
				Name:   parser.MethodTerminateSectors,
				Method: m.TerminateSectors,
			},
			legacyBuiltin.MethodsMiner.DeclareFaults: {
				Name:   parser.MethodDeclareFaults,
				Method: m.DeclareFaults,
			},
			legacyBuiltin.MethodsMiner.DeclareFaultsRecovered: {
				Name:   parser.MethodDeclareFaultsRecovered,
				Method: m.DeclareFaultsRecovered,
			},
			legacyBuiltin.MethodsMiner.OnDeferredCronEvent: {
				Name:   parser.MethodOnDeferredCronEvent,
				Method: m.OnDeferredCronEvent,
			},
			legacyBuiltin.MethodsMiner.CheckSectorProven: {
				Name:   parser.MethodCheckSectorProven,
				Method: m.CheckSectorProven,
			},
			legacyBuiltin.MethodsMiner.AddLockedFund: {
				Name:   parser.MethodAddLockedFund,
				Method: m.AddLockedFund,
			},
			legacyBuiltin.MethodsMiner.ReportConsensusFault: {
				Name:   parser.MethodReportConsensusFault,
				Method: m.ReportConsensusFault,
			},
			legacyBuiltin.MethodsMiner.WithdrawBalance: {
				Name:   parser.MethodWithdrawBalance,
				Method: m.WithdrawBalanceExported,
			},
		}
	case tools.V16.IsSupported(network, height):
		data = miner8.Methods
	case tools.V17.IsSupported(network, height):
		data = miner9.Methods
	case tools.V18.IsSupported(network, height):
		data = miner10.Methods
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		data = miner11.Methods
	case tools.V21.IsSupported(network, height):
		data = miner12.Methods
	case tools.V22.IsSupported(network, height):
		data = miner13.Methods
	case tools.V23.IsSupported(network, height):
		data = miner14.Methods
	case tools.V24.IsSupported(network, height):
		data = miner15.Methods
	case tools.V25.IsSupported(network, height):
		data = miner16.Methods
	default:
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	data[initialPledgeMethodNum] = nonLegacyBuiltin.MethodMeta{
		Name:   parser.MethodInitialPledge,
		Method: m.InitialPledgeExported,
	}
	return data, nil
}

func (*Miner) ConfirmUpdateWorkerKey(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	return parseGeneric(rawParams, nil, false, &abi.EmptyValue{}, &abi.EmptyValue{}, parser.ParamsKey)
}

func (*Miner) TerminateSectors(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V25.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &miner16.TerminateSectorsParams{}, &miner16.TerminateSectorsReturn{}, parser.ParamsKey)
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &miner15.TerminateSectorsParams{}, &miner15.TerminateSectorsReturn{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &miner14.TerminateSectorsParams{}, &miner14.TerminateSectorsReturn{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &miner13.TerminateSectorsParams{}, &miner13.TerminateSectorsReturn{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &miner12.TerminateSectorsParams{}, &miner12.TerminateSectorsReturn{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, rawReturn, true, &miner11.TerminateSectorsParams{}, &miner11.TerminateSectorsReturn{}, parser.ParamsKey)
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &miner10.TerminateSectorsParams{}, &miner10.TerminateSectorsReturn{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &miner9.TerminateSectorsParams{}, &miner9.TerminateSectorsReturn{}, parser.ParamsKey)
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &miner8.TerminateSectorsParams{}, &miner8.TerminateSectorsReturn{}, parser.ParamsKey)
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &legacyv7.TerminateSectorsParams{}, &legacyv7.TerminateSectorsReturn{}, parser.ParamsKey)
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &legacyv6.TerminateSectorsParams{}, &legacyv6.TerminateSectorsReturn{}, parser.ParamsKey)
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &legacyv5.TerminateSectorsParams{}, &legacyv5.TerminateSectorsReturn{}, parser.ParamsKey)
	case tools.V12.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &legacyv4.TerminateSectorsParams{}, &legacyv4.TerminateSectorsReturn{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parseGeneric(rawParams, rawReturn, true, &legacyv3.TerminateSectorsParams{}, &legacyv3.TerminateSectorsReturn{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V8, tools.V9):
		return parseGeneric(rawParams, rawReturn, true, &legacyv2.TerminateSectorsParams{}, &legacyv2.TerminateSectorsReturn{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return parseGeneric(rawParams, rawReturn, true, &legacyv1.TerminateSectorsParams{}, &legacyv1.TerminateSectorsReturn{}, parser.ParamsKey)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) DeclareFaults(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V25.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner16.DeclareFaultsParams{}, &miner16.DeclareFaultsParams{}, parser.ParamsKey)
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.DeclareFaultsParams{}, &miner15.DeclareFaultsParams{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.DeclareFaultsParams{}, &miner14.DeclareFaultsParams{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.DeclareFaultsParams{}, &miner13.DeclareFaultsParams{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.DeclareFaultsParams{}, &miner12.DeclareFaultsParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.DeclareFaultsParams{}, &miner11.DeclareFaultsParams{}, parser.ParamsKey)
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.DeclareFaultsParams{}, &miner10.DeclareFaultsParams{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.DeclareFaultsParams{}, &miner9.DeclareFaultsParams{}, parser.ParamsKey)
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner8.DeclareFaultsParams{}, &miner8.DeclareFaultsParams{}, parser.ParamsKey)
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv7.DeclareFaultsParams{}, &legacyv7.DeclareFaultsParams{}, parser.ParamsKey)
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv6.DeclareFaultsParams{}, &legacyv6.DeclareFaultsParams{}, parser.ParamsKey)
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv5.DeclareFaultsParams{}, &legacyv5.DeclareFaultsParams{}, parser.ParamsKey)
	case tools.V12.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv4.DeclareFaultsParams{}, &legacyv4.DeclareFaultsParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parseGeneric(rawParams, nil, false, &legacyv3.DeclareFaultsParams{}, &legacyv3.DeclareFaultsParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V8, tools.V9):
		return parseGeneric(rawParams, nil, false, &legacyv2.DeclareFaultsParams{}, &legacyv2.DeclareFaultsParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return parseGeneric(rawParams, nil, false, &legacyv1.DeclareFaultsParams{}, &legacyv1.DeclareFaultsParams{}, parser.ParamsKey)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func (*Miner) DeclareFaultsRecovered(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V25.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner16.DeclareFaultsRecoveredParams{}, &miner16.DeclareFaultsRecoveredParams{}, parser.ParamsKey)
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.DeclareFaultsRecoveredParams{}, &miner15.DeclareFaultsRecoveredParams{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.DeclareFaultsRecoveredParams{}, &miner14.DeclareFaultsRecoveredParams{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.DeclareFaultsRecoveredParams{}, &miner13.DeclareFaultsRecoveredParams{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.DeclareFaultsRecoveredParams{}, &miner12.DeclareFaultsRecoveredParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.DeclareFaultsRecoveredParams{}, &miner11.DeclareFaultsRecoveredParams{}, parser.ParamsKey)
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.DeclareFaultsRecoveredParams{}, &miner10.DeclareFaultsRecoveredParams{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.DeclareFaultsRecoveredParams{}, &miner9.DeclareFaultsRecoveredParams{}, parser.ParamsKey)
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner8.DeclareFaultsRecoveredParams{}, &miner8.DeclareFaultsRecoveredParams{}, parser.ParamsKey)
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv7.DeclareFaultsRecoveredParams{}, &legacyv7.DeclareFaultsRecoveredParams{}, parser.ParamsKey)
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv6.DeclareFaultsRecoveredParams{}, &legacyv6.DeclareFaultsRecoveredParams{}, parser.ParamsKey)
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv5.DeclareFaultsRecoveredParams{}, &legacyv5.DeclareFaultsRecoveredParams{}, parser.ParamsKey)
	case tools.V12.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv4.DeclareFaultsRecoveredParams{}, &legacyv4.DeclareFaultsRecoveredParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parseGeneric(rawParams, nil, false, &legacyv3.DeclareFaultsRecoveredParams{}, &legacyv3.DeclareFaultsRecoveredParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V8, tools.V9):
		return parseGeneric(rawParams, nil, false, &legacyv2.DeclareFaultsRecoveredParams{}, &legacyv2.DeclareFaultsRecoveredParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return parseGeneric(rawParams, nil, false, &legacyv1.DeclareFaultsRecoveredParams{}, &legacyv1.DeclareFaultsRecoveredParams{}, parser.ParamsKey)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) ProveReplicaUpdates(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V25.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner16.ProveReplicaUpdatesParams{}, &miner16.ProveReplicaUpdatesParams{}, parser.ParamsKey)
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.ProveReplicaUpdatesParams{}, &miner15.ProveReplicaUpdatesParams{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.ProveReplicaUpdatesParams{}, &miner14.ProveReplicaUpdatesParams{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.ProveReplicaUpdatesParams{}, &miner13.ProveReplicaUpdatesParams{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.ProveReplicaUpdatesParams{}, &miner12.ProveReplicaUpdatesParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.ProveReplicaUpdatesParams{}, &miner11.ProveReplicaUpdatesParams{}, parser.ParamsKey)
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.ProveReplicaUpdatesParams{}, &miner10.ProveReplicaUpdatesParams{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.ProveReplicaUpdatesParams{}, &miner9.ProveReplicaUpdatesParams{}, parser.ParamsKey)
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner8.ProveReplicaUpdatesParams{}, &miner8.ProveReplicaUpdatesParams{}, parser.ParamsKey)
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv7.ProveReplicaUpdatesParams{}, &legacyv7.ProveReplicaUpdatesParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V14)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) PreCommitSectorBatch2(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V25.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner16.PreCommitSectorBatchParams2{}, &miner16.PreCommitSectorBatchParams2{}, parser.ParamsKey)
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.PreCommitSectorBatchParams2{}, &miner15.PreCommitSectorBatchParams2{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.PreCommitSectorBatchParams2{}, &miner14.PreCommitSectorBatchParams2{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.PreCommitSectorBatchParams2{}, &miner13.PreCommitSectorBatchParams2{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.PreCommitSectorBatchParams2{}, &miner12.PreCommitSectorBatchParams2{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.PreCommitSectorBatchParams2{}, &miner11.PreCommitSectorBatchParams2{}, parser.ParamsKey)
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.PreCommitSectorBatchParams2{}, &miner10.PreCommitSectorBatchParams2{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.PreCommitSectorBatchParams2{}, &miner9.PreCommitSectorBatchParams2{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V16)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) ProveReplicaUpdates2(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V25.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &miner16.ProveReplicaUpdatesParams2{}, &bitfield.BitField{}, parser.ParamsKey)
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &miner15.ProveReplicaUpdatesParams2{}, &bitfield.BitField{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &miner14.ProveReplicaUpdatesParams2{}, &bitfield.BitField{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &miner13.ProveReplicaUpdatesParams2{}, &bitfield.BitField{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &miner12.ProveReplicaUpdatesParams2{}, &bitfield.BitField{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, rawReturn, true, &miner11.ProveReplicaUpdatesParams2{}, &bitfield.BitField{}, parser.ParamsKey)
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &miner10.ProveReplicaUpdatesParams2{}, &bitfield.BitField{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &miner9.ProveReplicaUpdatesParams2{}, &bitfield.BitField{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V16)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) ProveReplicaUpdates3(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V25.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &miner16.ProveReplicaUpdates3Params{}, &miner16.ProveReplicaUpdates3Return{}, parser.ParamsKey)
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &miner15.ProveReplicaUpdates3Params{}, &miner15.ProveReplicaUpdates3Return{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &miner14.ProveReplicaUpdates3Params{}, &miner14.ProveReplicaUpdates3Return{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &miner13.ProveReplicaUpdates3Params{}, &miner13.ProveReplicaUpdates3Return{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V21)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) ProveCommitAggregate(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V25.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner16.ProveCommitAggregateParams{}, &miner16.ProveCommitAggregateParams{}, parser.ParamsKey)
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.ProveCommitAggregateParams{}, &miner15.ProveCommitAggregateParams{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.ProveCommitAggregateParams{}, &miner14.ProveCommitAggregateParams{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.ProveCommitAggregateParams{}, &miner13.ProveCommitAggregateParams{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.ProveCommitAggregateParams{}, &miner12.ProveCommitAggregateParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.ProveCommitAggregateParams{}, &miner11.ProveCommitAggregateParams{}, parser.ParamsKey)
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.ProveCommitAggregateParams{}, &miner10.ProveCommitAggregateParams{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.ProveCommitAggregateParams{}, &miner9.ProveCommitAggregateParams{}, parser.ParamsKey)
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner8.ProveCommitAggregateParams{}, &miner8.ProveCommitAggregateParams{}, parser.ParamsKey)
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv7.ProveCommitAggregateParams{}, &legacyv7.ProveCommitAggregateParams{}, parser.ParamsKey)
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv6.ProveCommitAggregateParams{}, &legacyv6.ProveCommitAggregateParams{}, parser.ParamsKey)
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv5.ProveCommitAggregateParams{}, &legacyv5.ProveCommitAggregateParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V12)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) DisputeWindowedPoSt(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V25.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner16.DisputeWindowedPoStParams{}, &miner16.DisputeWindowedPoStParams{}, parser.ParamsKey)
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.DisputeWindowedPoStParams{}, &miner15.DisputeWindowedPoStParams{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.DisputeWindowedPoStParams{}, &miner14.DisputeWindowedPoStParams{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.DisputeWindowedPoStParams{}, &miner13.DisputeWindowedPoStParams{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.DisputeWindowedPoStParams{}, &miner12.DisputeWindowedPoStParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.DisputeWindowedPoStParams{}, &miner11.DisputeWindowedPoStParams{}, parser.ParamsKey)
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.DisputeWindowedPoStParams{}, &miner10.DisputeWindowedPoStParams{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.DisputeWindowedPoStParams{}, &miner9.DisputeWindowedPoStParams{}, parser.ParamsKey)
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner8.DisputeWindowedPoStParams{}, &miner8.DisputeWindowedPoStParams{}, parser.ParamsKey)
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv7.DisputeWindowedPoStParams{}, &legacyv7.DisputeWindowedPoStParams{}, parser.ParamsKey)
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv6.DisputeWindowedPoStParams{}, &legacyv6.DisputeWindowedPoStParams{}, parser.ParamsKey)
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv5.DisputeWindowedPoStParams{}, &legacyv5.DisputeWindowedPoStParams{}, parser.ParamsKey)
	case tools.V12.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv4.DisputeWindowedPoStParams{}, &legacyv4.DisputeWindowedPoStParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parseGeneric(rawParams, nil, false, &legacyv3.DisputeWindowedPoStParams{}, &legacyv3.DisputeWindowedPoStParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V9)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) ReportConsensusFault(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V25.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner16.ReportConsensusFaultParams{}, &miner16.ReportConsensusFaultParams{}, parser.ParamsKey)
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.ReportConsensusFaultParams{}, &miner15.ReportConsensusFaultParams{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.ReportConsensusFaultParams{}, &miner14.ReportConsensusFaultParams{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.ReportConsensusFaultParams{}, &miner13.ReportConsensusFaultParams{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.ReportConsensusFaultParams{}, &miner12.ReportConsensusFaultParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.ReportConsensusFaultParams{}, &miner11.ReportConsensusFaultParams{}, parser.ParamsKey)
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.ReportConsensusFaultParams{}, &miner10.ReportConsensusFaultParams{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.ReportConsensusFaultParams{}, &miner9.ReportConsensusFaultParams{}, parser.ParamsKey)
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner8.ReportConsensusFaultParams{}, &miner8.ReportConsensusFaultParams{}, parser.ParamsKey)
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv7.ReportConsensusFaultParams{}, &legacyv7.ReportConsensusFaultParams{}, parser.ParamsKey)
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv6.ReportConsensusFaultParams{}, &legacyv6.ReportConsensusFaultParams{}, parser.ParamsKey)
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv5.ReportConsensusFaultParams{}, &legacyv5.ReportConsensusFaultParams{}, parser.ParamsKey)
	case tools.V12.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv4.ReportConsensusFaultParams{}, &legacyv4.ReportConsensusFaultParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parseGeneric(rawParams, nil, false, &legacyv3.ReportConsensusFaultParams{}, &legacyv3.ReportConsensusFaultParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V8, tools.V9):
		return parseGeneric(rawParams, nil, false, &legacyv2.ReportConsensusFaultParams{}, &legacyv2.ReportConsensusFaultParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return parseGeneric(rawParams, nil, false, &legacyv1.ReportConsensusFaultParams{}, &legacyv1.ReportConsensusFaultParams{}, parser.ParamsKey)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) ChangeBeneficiaryExported(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V25.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner16.ChangeBeneficiaryParams{}, &miner16.ChangeBeneficiaryParams{}, parser.ParamsKey)
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.ChangeBeneficiaryParams{}, &miner15.ChangeBeneficiaryParams{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.ChangeBeneficiaryParams{}, &miner14.ChangeBeneficiaryParams{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.ChangeBeneficiaryParams{}, &miner13.ChangeBeneficiaryParams{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.ChangeBeneficiaryParams{}, &miner12.ChangeBeneficiaryParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.ChangeBeneficiaryParams{}, &miner11.ChangeBeneficiaryParams{}, parser.ParamsKey)
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.ChangeBeneficiaryParams{}, &miner10.ChangeBeneficiaryParams{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.ChangeBeneficiaryParams{}, &miner9.ChangeBeneficiaryParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V16)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) GetBeneficiary(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	if rawParams != nil {
		metadata[parser.ParamsKey] = base64.StdEncoding.EncodeToString(rawParams)
	}
	beneficiaryReturn, err := getBeneficiaryReturn(network, height, rawReturn)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = beneficiaryReturn
	return metadata, nil
}

func (*Miner) Constructor(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V25.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner16.MinerConstructorParams{}, &miner16.MinerConstructorParams{}, parser.ParamsKey)
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.MinerConstructorParams{}, &miner15.MinerConstructorParams{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.MinerConstructorParams{}, &miner14.MinerConstructorParams{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.MinerConstructorParams{}, &miner13.MinerConstructorParams{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.MinerConstructorParams{}, &miner12.MinerConstructorParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.MinerConstructorParams{}, &miner11.MinerConstructorParams{}, parser.ParamsKey)
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.MinerConstructorParams{}, &miner10.MinerConstructorParams{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.MinerConstructorParams{}, &miner9.MinerConstructorParams{}, parser.ParamsKey)
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner8.MinerConstructorParams{}, &miner8.MinerConstructorParams{}, parser.ParamsKey)
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv7.ConstructorParams{}, &legacyv7.ConstructorParams{}, parser.ParamsKey)
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv6.ConstructorParams{}, &legacyv6.ConstructorParams{}, parser.ParamsKey)
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv5.ConstructorParams{}, &legacyv5.ConstructorParams{}, parser.ParamsKey)
	case tools.V12.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv4.ConstructorParams{}, &legacyv4.ConstructorParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parseGeneric(rawParams, nil, false, &legacyv3.ConstructorParams{}, &legacyv3.ConstructorParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V8, tools.V9):
		return parseGeneric(rawParams, nil, false, &legacyv2.ConstructorParams{}, &legacyv2.ConstructorParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return parseGeneric(rawParams, nil, false, &legacyv1.ConstructorParams{}, &legacyv1.ConstructorParams{}, parser.ParamsKey)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) ApplyRewards(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V25.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner16.ApplyRewardParams{}, &miner16.ApplyRewardParams{}, parser.ParamsKey)
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.ApplyRewardParams{}, &miner15.ApplyRewardParams{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.ApplyRewardParams{}, &miner14.ApplyRewardParams{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.ApplyRewardParams{}, &miner13.ApplyRewardParams{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.ApplyRewardParams{}, &miner12.ApplyRewardParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.ApplyRewardParams{}, &miner11.ApplyRewardParams{}, parser.ParamsKey)
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.ApplyRewardParams{}, &miner10.ApplyRewardParams{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.ApplyRewardParams{}, &miner9.ApplyRewardParams{}, parser.ParamsKey)
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner8.ApplyRewardParams{}, &miner8.ApplyRewardParams{}, parser.ParamsKey)
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &builtinv7.ApplyRewardParams{}, &builtinv7.ApplyRewardParams{}, parser.ParamsKey)
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &builtinv6.ApplyRewardParams{}, &builtinv6.ApplyRewardParams{}, parser.ParamsKey)
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &builtinv5.ApplyRewardParams{}, &builtinv5.ApplyRewardParams{}, parser.ParamsKey)
	case tools.V12.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &builtinv4.ApplyRewardParams{}, &builtinv4.ApplyRewardParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V10, tools.V11):
		return parseGeneric(rawParams, nil, false, &builtinv3.ApplyRewardParams{}, &builtinv3.ApplyRewardParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V8, tools.V9):
		return parseGeneric(rawParams, nil, false, &builtinv2.ApplyRewardParams{}, &builtinv2.ApplyRewardParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return parseGeneric(rawParams, nil, false, &abi.TokenAmount{}, &abi.TokenAmount{}, parser.ParamsKey)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) OnDeferredCronEvent(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V25.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner16.DeferredCronEventParams{}, &miner16.DeferredCronEventParams{}, parser.ParamsKey)
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.DeferredCronEventParams{}, &miner15.DeferredCronEventParams{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.DeferredCronEventParams{}, &miner14.DeferredCronEventParams{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.DeferredCronEventParams{}, &miner13.DeferredCronEventParams{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.DeferredCronEventParams{}, &miner12.DeferredCronEventParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.DeferredCronEventParams{}, &miner11.DeferredCronEventParams{}, parser.ParamsKey)
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.DeferredCronEventParams{}, &miner10.DeferredCronEventParams{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.DeferredCronEventParams{}, &miner9.DeferredCronEventParams{}, parser.ParamsKey)
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner8.DeferredCronEventParams{}, &miner8.DeferredCronEventParams{}, parser.ParamsKey)

	// the difference in packages (builtin/legacy) is intentional and is how the underlying library is implemented
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &builtinv7.DeferredCronEventParams{}, &builtinv7.DeferredCronEventParams{}, parser.ParamsKey)
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &builtinv6.DeferredCronEventParams{}, &builtinv6.DeferredCronEventParams{}, parser.ParamsKey)
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv5.CronEventPayload{}, &abi.EmptyValue{}, parser.ParamsKey)

	case tools.V12.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv4.CronEventPayload{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V10, tools.V11):
		return parseGeneric(rawParams, nil, false, &legacyv3.CronEventPayload{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V8, tools.V9):
		return parseGeneric(rawParams, nil, false, &legacyv2.CronEventPayload{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return parseGeneric(rawParams, nil, false, &legacyv1.CronEventPayload{}, &abi.EmptyValue{}, parser.ParamsKey)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) InitialPledgeExported(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	return parseGeneric(rawParams, nil, false, &InitialPledgeReturn{}, &InitialPledgeReturn{}, parser.ReturnKey)
}
