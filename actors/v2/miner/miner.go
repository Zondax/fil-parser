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
	version := tools.VersionFromHeight(network, height)
	params, ok := terminateSectorsParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := terminateSectorsReturn()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, rawReturn, true, params, returnValue, parser.ParamsKey)
}

func (*Miner) DeclareFaults(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)

	params, ok := declareFaultsParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parseGeneric(rawParams, nil, false, params, &abi.EmptyValue{}, parser.ParamsKey)
}

func (*Miner) DeclareFaultsRecovered(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := declareFaultsRecoveredParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, nil, false, params, &abi.EmptyValue{}, parser.ParamsKey)
}

func (*Miner) ProveReplicaUpdates(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := proveReplicaUpdatesParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, nil, false, params, &abi.EmptyValue{}, parser.ParamsKey)
}

func (*Miner) PreCommitSectorBatch2(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := preCommitSectorBatchParams2()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, nil, false, params, &abi.EmptyValue{}, parser.ParamsKey)
}

func (*Miner) ProveReplicaUpdates2(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := proveReplicaUpdatesParams2()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, rawReturn, true, params, &bitfield.BitField{}, parser.ParamsKey)
}

func (*Miner) ProveReplicaUpdates3(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := proveReplicaUpdates3Params()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := proveReplicaUpdates3Return()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, rawReturn, true, params, returnValue, parser.ParamsKey)
}

func (*Miner) ProveCommitAggregate(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := proveCommitAggregateParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, nil, false, params, &abi.EmptyValue{}, parser.ParamsKey)
}

func (*Miner) DisputeWindowedPoSt(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := disputeWindowedPoStParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, nil, false, params, &abi.EmptyValue{}, parser.ParamsKey)
}

func (*Miner) ReportConsensusFault(network string, height int64, rawParams []byte) (map[string]interface{}, error) {

	version := tools.VersionFromHeight(network, height)
	params, ok := reportConsensusFaultParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, nil, false, params, &abi.EmptyValue{}, parser.ParamsKey)
}

func (*Miner) ChangeBeneficiaryExported(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := changeBeneficiaryParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, nil, false, params, &abi.EmptyValue{}, parser.ParamsKey)
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
	version := tools.VersionFromHeight(network, height)
	params, ok := minerConstructorParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, nil, false, params, &abi.EmptyValue{}, parser.ParamsKey)
}

func (*Miner) ApplyRewards(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := applyRewardParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, nil, false, params, &abi.EmptyValue{}, parser.ParamsKey)
}

func (*Miner) OnDeferredCronEvent(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := deferredCronEventParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, nil, false, params, &abi.EmptyValue{}, parser.ParamsKey)
}

func (*Miner) InitialPledgeExported(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	return parseGeneric(rawParams, nil, false, &InitialPledgeReturn{}, &InitialPledgeReturn{}, parser.ReturnKey)
}
