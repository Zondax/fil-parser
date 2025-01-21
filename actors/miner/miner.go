package miner

import (
	"fmt"

	miner10 "github.com/filecoin-project/go-state-types/builtin/v10/miner"
	miner11 "github.com/filecoin-project/go-state-types/builtin/v11/miner"
	miner12 "github.com/filecoin-project/go-state-types/builtin/v12/miner"
	miner13 "github.com/filecoin-project/go-state-types/builtin/v13/miner"
	miner14 "github.com/filecoin-project/go-state-types/builtin/v14/miner"
	miner15 "github.com/filecoin-project/go-state-types/builtin/v15/miner"
	miner8 "github.com/filecoin-project/go-state-types/builtin/v8/miner"
	miner9 "github.com/filecoin-project/go-state-types/builtin/v9/miner"
	"github.com/zondax/fil-parser/tools"
)

func TerminateSectors(height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		return parseGeneric[*miner15.TerminateSectorsParams, *miner15.TerminateSectorsReturn](rawParams, rawReturn, true)
	case tools.V14.IsSupported(height):
		return parseGeneric[*miner14.TerminateSectorsParams, *miner14.TerminateSectorsReturn](rawParams, rawReturn, true)
	case tools.V13.IsSupported(height):
		return parseGeneric[*miner13.TerminateSectorsParams, *miner13.TerminateSectorsReturn](rawParams, rawReturn, true)
	case tools.V12.IsSupported(height):
		return parseGeneric[*miner12.TerminateSectorsParams, *miner12.TerminateSectorsReturn](rawParams, rawReturn, true)
	case tools.V11.IsSupported(height):
		return parseGeneric[*miner11.TerminateSectorsParams, *miner11.TerminateSectorsReturn](rawParams, rawReturn, true)
	case tools.V10.IsSupported(height):
		return parseGeneric[*miner10.TerminateSectorsParams, *miner10.TerminateSectorsReturn](rawParams, rawReturn, true)
	case tools.V9.IsSupported(height):
		return parseGeneric[*miner9.TerminateSectorsParams, *miner9.TerminateSectorsReturn](rawParams, rawReturn, true)
	case tools.V8.IsSupported(height):
		return parseGeneric[*miner8.TerminateSectorsParams, *miner8.TerminateSectorsReturn](rawParams, rawReturn, true)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func DeclareFaults(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		return parseGeneric[*miner15.DeclareFaultsParams, *miner15.DeclareFaultsParams](rawParams, nil, false)
	case tools.V14.IsSupported(height):
		return parseGeneric[*miner14.DeclareFaultsParams, *miner14.DeclareFaultsParams](rawParams, nil, false)
	case tools.V13.IsSupported(height):
		return parseGeneric[*miner13.DeclareFaultsParams, *miner13.DeclareFaultsParams](rawParams, nil, false)
	case tools.V12.IsSupported(height):
		return parseGeneric[*miner12.DeclareFaultsParams, *miner12.DeclareFaultsParams](rawParams, nil, false)
	case tools.V11.IsSupported(height):
		return parseGeneric[*miner11.DeclareFaultsParams, *miner11.DeclareFaultsParams](rawParams, nil, false)
	case tools.V10.IsSupported(height):
		return parseGeneric[*miner10.DeclareFaultsParams, *miner10.DeclareFaultsParams](rawParams, nil, false)
	case tools.V9.IsSupported(height):
		return parseGeneric[*miner9.DeclareFaultsParams, *miner9.DeclareFaultsParams](rawParams, nil, false)
	case tools.V8.IsSupported(height):
		return parseGeneric[*miner8.DeclareFaultsParams, *miner8.DeclareFaultsParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func DeclareFaultsRecovered(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		return parseGeneric[*miner15.DeclareFaultsRecoveredParams, *miner15.DeclareFaultsRecoveredParams](rawParams, nil, false)
	case tools.V14.IsSupported(height):
		return parseGeneric[*miner14.DeclareFaultsRecoveredParams, *miner14.DeclareFaultsRecoveredParams](rawParams, nil, false)
	case tools.V13.IsSupported(height):
		return parseGeneric[*miner13.DeclareFaultsRecoveredParams, *miner13.DeclareFaultsRecoveredParams](rawParams, nil, false)
	case tools.V12.IsSupported(height):
		return parseGeneric[*miner12.DeclareFaultsRecoveredParams, *miner12.DeclareFaultsRecoveredParams](rawParams, nil, false)
	case tools.V11.IsSupported(height):
		return parseGeneric[*miner11.DeclareFaultsRecoveredParams, *miner11.DeclareFaultsRecoveredParams](rawParams, nil, false)
	case tools.V10.IsSupported(height):
		return parseGeneric[*miner10.DeclareFaultsRecoveredParams, *miner10.DeclareFaultsRecoveredParams](rawParams, nil, false)
	case tools.V9.IsSupported(height):
		return parseGeneric[*miner9.DeclareFaultsRecoveredParams, *miner9.DeclareFaultsRecoveredParams](rawParams, nil, false)
	case tools.V8.IsSupported(height):
		return parseGeneric[*miner8.DeclareFaultsRecoveredParams, *miner8.DeclareFaultsRecoveredParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func ProveReplicaUpdates(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		return parseGeneric[*miner15.ProveReplicaUpdatesParams, *miner15.ProveReplicaUpdatesParams](rawParams, nil, false)
	case tools.V14.IsSupported(height):
		return parseGeneric[*miner14.ProveReplicaUpdatesParams, *miner14.ProveReplicaUpdatesParams](rawParams, nil, false)
	case tools.V13.IsSupported(height):
		return parseGeneric[*miner13.ProveReplicaUpdatesParams, *miner13.ProveReplicaUpdatesParams](rawParams, nil, false)
	case tools.V12.IsSupported(height):
		return parseGeneric[*miner12.ProveReplicaUpdatesParams, *miner12.ProveReplicaUpdatesParams](rawParams, nil, false)
	case tools.V11.IsSupported(height):
		return parseGeneric[*miner11.ProveReplicaUpdatesParams, *miner11.ProveReplicaUpdatesParams](rawParams, nil, false)
	case tools.V10.IsSupported(height):
		return parseGeneric[*miner10.ProveReplicaUpdatesParams, *miner10.ProveReplicaUpdatesParams](rawParams, nil, false)
	case tools.V9.IsSupported(height):
		return parseGeneric[*miner9.ProveReplicaUpdatesParams, *miner9.ProveReplicaUpdatesParams](rawParams, nil, false)
	case tools.V8.IsSupported(height):
		return parseGeneric[*miner8.ProveReplicaUpdatesParams, *miner8.ProveReplicaUpdatesParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func PreCommitSectorBatch2(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		return parseGeneric[*miner15.PreCommitSectorBatchParams2, *miner15.PreCommitSectorBatchParams2](rawParams, nil, false)
	case tools.V14.IsSupported(height):
		return parseGeneric[*miner14.PreCommitSectorBatchParams2, *miner14.PreCommitSectorBatchParams2](rawParams, nil, false)
	case tools.V13.IsSupported(height):
		return parseGeneric[*miner13.PreCommitSectorBatchParams2, *miner13.PreCommitSectorBatchParams2](rawParams, nil, false)
	case tools.V12.IsSupported(height):
		return parseGeneric[*miner12.PreCommitSectorBatchParams2, *miner12.PreCommitSectorBatchParams2](rawParams, nil, false)
	case tools.V11.IsSupported(height):
		return parseGeneric[*miner11.PreCommitSectorBatchParams2, *miner11.PreCommitSectorBatchParams2](rawParams, nil, false)
	case tools.V10.IsSupported(height):
		return parseGeneric[*miner10.PreCommitSectorBatchParams2, *miner10.PreCommitSectorBatchParams2](rawParams, nil, false)
	case tools.V9.IsSupported(height):
		return parseGeneric[*miner9.PreCommitSectorBatchParams2, *miner9.PreCommitSectorBatchParams2](rawParams, nil, false)
	case tools.V8.IsSupported(height):
		return nil, fmt.Errorf("unsupported height: %d", height)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func ProveReplicaUpdates2(height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		return parseGeneric[*miner15.ProveReplicaUpdatesParams2, *miner15.ProveReplicaUpdatesParams2](rawParams, rawReturn, true)
	case tools.V14.IsSupported(height):
		return parseGeneric[*miner14.ProveReplicaUpdatesParams2, *miner14.ProveReplicaUpdatesParams2](rawParams, rawReturn, true)
	case tools.V13.IsSupported(height):
		return parseGeneric[*miner13.ProveReplicaUpdatesParams2, *miner13.ProveReplicaUpdatesParams2](rawParams, rawReturn, true)
	case tools.V12.IsSupported(height):
		return parseGeneric[*miner12.ProveReplicaUpdatesParams2, *miner12.ProveReplicaUpdatesParams2](rawParams, rawReturn, true)
	case tools.V11.IsSupported(height):
		return parseGeneric[*miner11.ProveReplicaUpdatesParams2, *miner11.ProveReplicaUpdatesParams2](rawParams, rawReturn, true)
	case tools.V10.IsSupported(height):
		return parseGeneric[*miner10.ProveReplicaUpdatesParams2, *miner10.ProveReplicaUpdatesParams2](rawParams, rawReturn, true)
	case tools.V9.IsSupported(height):
		return parseGeneric[*miner9.ProveReplicaUpdatesParams2, *miner9.ProveReplicaUpdatesParams2](rawParams, rawReturn, true)
	case tools.V8.IsSupported(height):
		return nil, fmt.Errorf("unsupported height: %d", height)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func ProveCommitAggregate(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		return parseGeneric[*miner15.ProveCommitAggregateParams, *miner15.ProveCommitAggregateParams](rawParams, nil, false)
	case tools.V14.IsSupported(height):
		return parseGeneric[*miner14.ProveCommitAggregateParams, *miner14.ProveCommitAggregateParams](rawParams, nil, false)
	case tools.V13.IsSupported(height):
		return parseGeneric[*miner13.ProveCommitAggregateParams, *miner13.ProveCommitAggregateParams](rawParams, nil, false)
	case tools.V12.IsSupported(height):
		return parseGeneric[*miner12.ProveCommitAggregateParams, *miner12.ProveCommitAggregateParams](rawParams, nil, false)
	case tools.V11.IsSupported(height):
		return parseGeneric[*miner11.ProveCommitAggregateParams, *miner11.ProveCommitAggregateParams](rawParams, nil, false)
	case tools.V10.IsSupported(height):
		return parseGeneric[*miner10.ProveCommitAggregateParams, *miner10.ProveCommitAggregateParams](rawParams, nil, false)
	case tools.V9.IsSupported(height):
		return parseGeneric[*miner9.ProveCommitAggregateParams, *miner9.ProveCommitAggregateParams](rawParams, nil, false)
	case tools.V8.IsSupported(height):
		return parseGeneric[*miner8.ProveCommitAggregateParams, *miner8.ProveCommitAggregateParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func DisputeWindowedPoSt(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		return parseGeneric[*miner15.DisputeWindowedPoStParams, *miner15.DisputeWindowedPoStParams](rawParams, nil, false)
	case tools.V14.IsSupported(height):
		return parseGeneric[*miner14.DisputeWindowedPoStParams, *miner14.DisputeWindowedPoStParams](rawParams, nil, false)
	case tools.V13.IsSupported(height):
		return parseGeneric[*miner13.DisputeWindowedPoStParams, *miner13.DisputeWindowedPoStParams](rawParams, nil, false)
	case tools.V12.IsSupported(height):
		return parseGeneric[*miner12.DisputeWindowedPoStParams, *miner12.DisputeWindowedPoStParams](rawParams, nil, false)
	case tools.V11.IsSupported(height):
		return parseGeneric[*miner11.DisputeWindowedPoStParams, *miner11.DisputeWindowedPoStParams](rawParams, nil, false)
	case tools.V10.IsSupported(height):
		return parseGeneric[*miner10.DisputeWindowedPoStParams, *miner10.DisputeWindowedPoStParams](rawParams, nil, false)
	case tools.V9.IsSupported(height):
		return parseGeneric[*miner9.DisputeWindowedPoStParams, *miner9.DisputeWindowedPoStParams](rawParams, nil, false)
	case tools.V8.IsSupported(height):
		return parseGeneric[*miner8.DisputeWindowedPoStParams, *miner8.DisputeWindowedPoStParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func ReportConsensusFault(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		return parseGeneric[*miner15.ReportConsensusFaultParams, *miner15.ReportConsensusFaultParams](rawParams, nil, false)
	case tools.V14.IsSupported(height):
		return parseGeneric[*miner14.ReportConsensusFaultParams, *miner14.ReportConsensusFaultParams](rawParams, nil, false)
	case tools.V13.IsSupported(height):
		return parseGeneric[*miner13.ReportConsensusFaultParams, *miner13.ReportConsensusFaultParams](rawParams, nil, false)
	case tools.V12.IsSupported(height):
		return parseGeneric[*miner12.ReportConsensusFaultParams, *miner12.ReportConsensusFaultParams](rawParams, nil, false)
	case tools.V11.IsSupported(height):
		return parseGeneric[*miner11.ReportConsensusFaultParams, *miner11.ReportConsensusFaultParams](rawParams, nil, false)
	case tools.V10.IsSupported(height):
		return parseGeneric[*miner10.ReportConsensusFaultParams, *miner10.ReportConsensusFaultParams](rawParams, nil, false)
	case tools.V9.IsSupported(height):
		return parseGeneric[*miner9.ReportConsensusFaultParams, *miner9.ReportConsensusFaultParams](rawParams, nil, false)
	case tools.V8.IsSupported(height):
		return parseGeneric[*miner8.ReportConsensusFaultParams, *miner8.ReportConsensusFaultParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func ChangeBeneficiary(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		return parseGeneric[*miner15.ChangeBeneficiaryParams, *miner15.ChangeBeneficiaryParams](rawParams, nil, false)
	case tools.V14.IsSupported(height):
		return parseGeneric[*miner14.ChangeBeneficiaryParams, *miner14.ChangeBeneficiaryParams](rawParams, nil, false)
	case tools.V13.IsSupported(height):
		return parseGeneric[*miner13.ChangeBeneficiaryParams, *miner13.ChangeBeneficiaryParams](rawParams, nil, false)
	case tools.V12.IsSupported(height):
		return parseGeneric[*miner12.ChangeBeneficiaryParams, *miner12.ChangeBeneficiaryParams](rawParams, nil, false)
	case tools.V11.IsSupported(height):
		return parseGeneric[*miner11.ChangeBeneficiaryParams, *miner11.ChangeBeneficiaryParams](rawParams, nil, false)
	case tools.V10.IsSupported(height):
		return parseGeneric[*miner10.ChangeBeneficiaryParams, *miner10.ChangeBeneficiaryParams](rawParams, nil, false)
	case tools.V9.IsSupported(height):
		return parseGeneric[*miner9.ChangeBeneficiaryParams, *miner9.ChangeBeneficiaryParams](rawParams, nil, false)
	case tools.V8.IsSupported(height):
		return parseGeneric[*miner8.ChangeBeneficiaryParams, *miner8.ChangeBeneficiaryParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func MinerConstructor(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		return parseGeneric[*miner15.MinerConstructorParams, *miner15.MinerConstructorParams](rawParams, nil, false)
	case tools.V14.IsSupported(height):
		return parseGeneric[*miner14.MinerConstructorParams, *miner14.MinerConstructorParams](rawParams, nil, false)
	case tools.V13.IsSupported(height):
		return parseGeneric[*miner13.MinerConstructorParams, *miner13.MinerConstructorParams](rawParams, nil, false)
	case tools.V12.IsSupported(height):
		return parseGeneric[*miner12.MinerConstructorParams, *miner12.MinerConstructorParams](rawParams, nil, false)
	case tools.V11.IsSupported(height):
		return parseGeneric[*miner11.MinerConstructorParams, *miner11.MinerConstructorParams](rawParams, nil, false)
	case tools.V10.IsSupported(height):
		return parseGeneric[*miner10.MinerConstructorParams, *miner10.MinerConstructorParams](rawParams, nil, false)
	case tools.V9.IsSupported(height):
		return parseGeneric[*miner9.MinerConstructorParams, *miner9.MinerConstructorParams](rawParams, nil, false)
	case tools.V8.IsSupported(height):
		return parseGeneric[*miner8.MinerConstructorParams, *miner8.MinerConstructorParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func ApplyRewards(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		return parseGeneric[*miner15.ApplyRewardParams, *miner15.ApplyRewardParams](rawParams, nil, false)
	case tools.V14.IsSupported(height):
		return parseGeneric[*miner14.ApplyRewardParams, *miner14.ApplyRewardParams](rawParams, nil, false)
	case tools.V13.IsSupported(height):
		return parseGeneric[*miner13.ApplyRewardParams, *miner13.ApplyRewardParams](rawParams, nil, false)
	case tools.V12.IsSupported(height):
		return parseGeneric[*miner12.ApplyRewardParams, *miner12.ApplyRewardParams](rawParams, nil, false)
	case tools.V11.IsSupported(height):
		return parseGeneric[*miner11.ApplyRewardParams, *miner11.ApplyRewardParams](rawParams, nil, false)
	case tools.V10.IsSupported(height):
		return parseGeneric[*miner10.ApplyRewardParams, *miner10.ApplyRewardParams](rawParams, nil, false)
	case tools.V9.IsSupported(height):
		return parseGeneric[*miner9.ApplyRewardParams, *miner9.ApplyRewardParams](rawParams, nil, false)
	case tools.V8.IsSupported(height):
		return parseGeneric[*miner8.ApplyRewardParams, *miner8.ApplyRewardParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func OnDeferredCronEvent(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		return parseGeneric[*miner15.DeferredCronEventParams, *miner15.DeferredCronEventParams](rawParams, nil, false)
	case tools.V14.IsSupported(height):
		return parseGeneric[*miner14.DeferredCronEventParams, *miner14.DeferredCronEventParams](rawParams, nil, false)
	case tools.V13.IsSupported(height):
		return parseGeneric[*miner13.DeferredCronEventParams, *miner13.DeferredCronEventParams](rawParams, nil, false)
	case tools.V12.IsSupported(height):
		return parseGeneric[*miner12.DeferredCronEventParams, *miner12.DeferredCronEventParams](rawParams, nil, false)
	case tools.V11.IsSupported(height):
		return parseGeneric[*miner11.DeferredCronEventParams, *miner11.DeferredCronEventParams](rawParams, nil, false)
	case tools.V10.IsSupported(height):
		return parseGeneric[*miner10.DeferredCronEventParams, *miner10.DeferredCronEventParams](rawParams, nil, false)
	case tools.V9.IsSupported(height):
		return parseGeneric[*miner9.DeferredCronEventParams, *miner9.DeferredCronEventParams](rawParams, nil, false)
	case tools.V8.IsSupported(height):
		return parseGeneric[*miner8.DeferredCronEventParams, *miner8.DeferredCronEventParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}
