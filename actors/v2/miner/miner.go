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
	"github.com/filecoin-project/go-state-types/manifest"
	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/miner"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/miner"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/miner"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/miner"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/miner"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/miner"
	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/tools"
)

func (m *Miner) Name() string {
	return manifest.MinerKey
}

func (*Miner) TerminateSectors(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &miner15.TerminateSectorsParams{}, &miner15.TerminateSectorsReturn{})
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &miner14.TerminateSectorsParams{}, &miner14.TerminateSectorsReturn{})
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &miner13.TerminateSectorsParams{}, &miner13.TerminateSectorsReturn{})
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &miner12.TerminateSectorsParams{}, &miner12.TerminateSectorsReturn{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, rawReturn, true, &miner11.TerminateSectorsParams{}, &miner11.TerminateSectorsReturn{})
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &miner10.TerminateSectorsParams{}, &miner10.TerminateSectorsReturn{})
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &miner9.TerminateSectorsParams{}, &miner9.TerminateSectorsReturn{})
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &miner8.TerminateSectorsParams{}, &miner8.TerminateSectorsReturn{})
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &legacyv7.TerminateSectorsParams{}, &legacyv7.TerminateSectorsReturn{})
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &legacyv6.TerminateSectorsParams{}, &legacyv6.TerminateSectorsReturn{})
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &legacyv5.TerminateSectorsParams{}, &legacyv5.TerminateSectorsReturn{})
	case tools.V12.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &legacyv4.TerminateSectorsParams{}, &legacyv4.TerminateSectorsReturn{})
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parseGeneric(rawParams, rawReturn, true, &legacyv3.TerminateSectorsParams{}, &legacyv3.TerminateSectorsReturn{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V9)...):
		return parseGeneric(rawParams, rawReturn, true, &legacyv2.TerminateSectorsParams{}, &legacyv2.TerminateSectorsReturn{})
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) DeclareFaults(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.DeclareFaultsParams{}, &miner15.DeclareFaultsParams{})
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.DeclareFaultsParams{}, &miner14.DeclareFaultsParams{})
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.DeclareFaultsParams{}, &miner13.DeclareFaultsParams{})
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.DeclareFaultsParams{}, &miner12.DeclareFaultsParams{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.DeclareFaultsParams{}, &miner11.DeclareFaultsParams{})
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.DeclareFaultsParams{}, &miner10.DeclareFaultsParams{})
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.DeclareFaultsParams{}, &miner9.DeclareFaultsParams{})
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner8.DeclareFaultsParams{}, &miner8.DeclareFaultsParams{})
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv7.DeclareFaultsParams{}, &legacyv7.DeclareFaultsParams{})
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv6.DeclareFaultsParams{}, &legacyv6.DeclareFaultsParams{})
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv5.DeclareFaultsParams{}, &legacyv5.DeclareFaultsParams{})
	case tools.V12.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv4.DeclareFaultsParams{}, &legacyv4.DeclareFaultsParams{})
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parseGeneric(rawParams, nil, false, &legacyv3.DeclareFaultsParams{}, &legacyv3.DeclareFaultsParams{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V9)...):
		return parseGeneric(rawParams, nil, false, &legacyv2.DeclareFaultsParams{}, &legacyv2.DeclareFaultsParams{})
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func (*Miner) DeclareFaultsRecovered(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.DeclareFaultsRecoveredParams{}, &miner15.DeclareFaultsRecoveredParams{})
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.DeclareFaultsRecoveredParams{}, &miner14.DeclareFaultsRecoveredParams{})
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.DeclareFaultsRecoveredParams{}, &miner13.DeclareFaultsRecoveredParams{})
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.DeclareFaultsRecoveredParams{}, &miner12.DeclareFaultsRecoveredParams{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.DeclareFaultsRecoveredParams{}, &miner11.DeclareFaultsRecoveredParams{})
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.DeclareFaultsRecoveredParams{}, &miner10.DeclareFaultsRecoveredParams{})
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.DeclareFaultsRecoveredParams{}, &miner9.DeclareFaultsRecoveredParams{})
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner8.DeclareFaultsRecoveredParams{}, &miner8.DeclareFaultsRecoveredParams{})
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv7.DeclareFaultsRecoveredParams{}, &legacyv7.DeclareFaultsRecoveredParams{})
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv6.DeclareFaultsRecoveredParams{}, &legacyv6.DeclareFaultsRecoveredParams{})
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv5.DeclareFaultsRecoveredParams{}, &legacyv5.DeclareFaultsRecoveredParams{})
	case tools.V12.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv4.DeclareFaultsRecoveredParams{}, &legacyv4.DeclareFaultsRecoveredParams{})
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parseGeneric(rawParams, nil, false, &legacyv3.DeclareFaultsRecoveredParams{}, &legacyv3.DeclareFaultsRecoveredParams{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V9)...):
		return parseGeneric(rawParams, nil, false, &legacyv2.DeclareFaultsRecoveredParams{}, &legacyv2.DeclareFaultsRecoveredParams{})
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) ProveReplicaUpdates(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.ProveReplicaUpdatesParams{}, &miner15.ProveReplicaUpdatesParams{})
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.ProveReplicaUpdatesParams{}, &miner14.ProveReplicaUpdatesParams{})
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.ProveReplicaUpdatesParams{}, &miner13.ProveReplicaUpdatesParams{})
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.ProveReplicaUpdatesParams{}, &miner12.ProveReplicaUpdatesParams{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.ProveReplicaUpdatesParams{}, &miner11.ProveReplicaUpdatesParams{})
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.ProveReplicaUpdatesParams{}, &miner10.ProveReplicaUpdatesParams{})
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.ProveReplicaUpdatesParams{}, &miner9.ProveReplicaUpdatesParams{})
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner8.ProveReplicaUpdatesParams{}, &miner8.ProveReplicaUpdatesParams{})
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv7.ProveReplicaUpdatesParams{}, &legacyv7.ProveReplicaUpdatesParams{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V14)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) PreCommitSectorBatch2(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.PreCommitSectorBatchParams2{}, &miner15.PreCommitSectorBatchParams2{})
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.PreCommitSectorBatchParams2{}, &miner14.PreCommitSectorBatchParams2{})
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.PreCommitSectorBatchParams2{}, &miner13.PreCommitSectorBatchParams2{})
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.PreCommitSectorBatchParams2{}, &miner12.PreCommitSectorBatchParams2{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.PreCommitSectorBatchParams2{}, &miner11.PreCommitSectorBatchParams2{})
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.PreCommitSectorBatchParams2{}, &miner10.PreCommitSectorBatchParams2{})
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.PreCommitSectorBatchParams2{}, &miner9.PreCommitSectorBatchParams2{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V16)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) ProveReplicaUpdates2(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &miner15.ProveReplicaUpdatesParams2{}, &miner15.ProveReplicaUpdatesParams2{})
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &miner14.ProveReplicaUpdatesParams2{}, &miner14.ProveReplicaUpdatesParams2{})
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &miner13.ProveReplicaUpdatesParams2{}, &miner13.ProveReplicaUpdatesParams2{})
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &miner12.ProveReplicaUpdatesParams2{}, &miner12.ProveReplicaUpdatesParams2{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, rawReturn, true, &miner11.ProveReplicaUpdatesParams2{}, &miner11.ProveReplicaUpdatesParams2{})
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &miner10.ProveReplicaUpdatesParams2{}, &miner10.ProveReplicaUpdatesParams2{})
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &miner9.ProveReplicaUpdatesParams2{}, &miner9.ProveReplicaUpdatesParams2{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V16)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) ProveCommitAggregate(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.ProveCommitAggregateParams{}, &miner15.ProveCommitAggregateParams{})
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.ProveCommitAggregateParams{}, &miner14.ProveCommitAggregateParams{})
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.ProveCommitAggregateParams{}, &miner13.ProveCommitAggregateParams{})
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.ProveCommitAggregateParams{}, &miner12.ProveCommitAggregateParams{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.ProveCommitAggregateParams{}, &miner11.ProveCommitAggregateParams{})
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.ProveCommitAggregateParams{}, &miner10.ProveCommitAggregateParams{})
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.ProveCommitAggregateParams{}, &miner9.ProveCommitAggregateParams{})
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner8.ProveCommitAggregateParams{}, &miner8.ProveCommitAggregateParams{})
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv7.ProveCommitAggregateParams{}, &legacyv7.ProveCommitAggregateParams{})
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv6.ProveCommitAggregateParams{}, &legacyv6.ProveCommitAggregateParams{})
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv5.ProveCommitAggregateParams{}, &legacyv5.ProveCommitAggregateParams{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V12)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) DisputeWindowedPoSt(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.DisputeWindowedPoStParams{}, &miner15.DisputeWindowedPoStParams{})
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.DisputeWindowedPoStParams{}, &miner14.DisputeWindowedPoStParams{})
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.DisputeWindowedPoStParams{}, &miner13.DisputeWindowedPoStParams{})
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.DisputeWindowedPoStParams{}, &miner12.DisputeWindowedPoStParams{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.DisputeWindowedPoStParams{}, &miner11.DisputeWindowedPoStParams{})
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.DisputeWindowedPoStParams{}, &miner10.DisputeWindowedPoStParams{})
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.DisputeWindowedPoStParams{}, &miner9.DisputeWindowedPoStParams{})
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner8.DisputeWindowedPoStParams{}, &miner8.DisputeWindowedPoStParams{})
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv7.DisputeWindowedPoStParams{}, &legacyv7.DisputeWindowedPoStParams{})
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv6.DisputeWindowedPoStParams{}, &legacyv6.DisputeWindowedPoStParams{})
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv5.DisputeWindowedPoStParams{}, &legacyv5.DisputeWindowedPoStParams{})
	case tools.V12.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv4.DisputeWindowedPoStParams{}, &legacyv4.DisputeWindowedPoStParams{})
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parseGeneric(rawParams, nil, false, &legacyv3.DisputeWindowedPoStParams{}, &legacyv3.DisputeWindowedPoStParams{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V9)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) ReportConsensusFault(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.ReportConsensusFaultParams{}, &miner15.ReportConsensusFaultParams{})
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.ReportConsensusFaultParams{}, &miner14.ReportConsensusFaultParams{})
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.ReportConsensusFaultParams{}, &miner13.ReportConsensusFaultParams{})
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.ReportConsensusFaultParams{}, &miner12.ReportConsensusFaultParams{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.ReportConsensusFaultParams{}, &miner11.ReportConsensusFaultParams{})
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.ReportConsensusFaultParams{}, &miner10.ReportConsensusFaultParams{})
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.ReportConsensusFaultParams{}, &miner9.ReportConsensusFaultParams{})
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner8.ReportConsensusFaultParams{}, &miner8.ReportConsensusFaultParams{})
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv7.ReportConsensusFaultParams{}, &legacyv7.ReportConsensusFaultParams{})
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv6.ReportConsensusFaultParams{}, &legacyv6.ReportConsensusFaultParams{})
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv5.ReportConsensusFaultParams{}, &legacyv5.ReportConsensusFaultParams{})
	case tools.V12.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv4.ReportConsensusFaultParams{}, &legacyv4.ReportConsensusFaultParams{})
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parseGeneric(rawParams, nil, false, &legacyv3.ReportConsensusFaultParams{}, &legacyv3.ReportConsensusFaultParams{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V9)...):
		return parseGeneric(rawParams, nil, false, &legacyv2.ReportConsensusFaultParams{}, &legacyv2.ReportConsensusFaultParams{})
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) ChangeBeneficiaryExported(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.ChangeBeneficiaryParams{}, &miner15.ChangeBeneficiaryParams{})
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.ChangeBeneficiaryParams{}, &miner14.ChangeBeneficiaryParams{})
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.ChangeBeneficiaryParams{}, &miner13.ChangeBeneficiaryParams{})
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.ChangeBeneficiaryParams{}, &miner12.ChangeBeneficiaryParams{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.ChangeBeneficiaryParams{}, &miner11.ChangeBeneficiaryParams{})
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.ChangeBeneficiaryParams{}, &miner10.ChangeBeneficiaryParams{})
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.ChangeBeneficiaryParams{}, &miner9.ChangeBeneficiaryParams{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V16)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) Constructor(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.MinerConstructorParams{}, &miner15.MinerConstructorParams{})
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.MinerConstructorParams{}, &miner14.MinerConstructorParams{})
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.MinerConstructorParams{}, &miner13.MinerConstructorParams{})
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.MinerConstructorParams{}, &miner12.MinerConstructorParams{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.MinerConstructorParams{}, &miner11.MinerConstructorParams{})
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.MinerConstructorParams{}, &miner10.MinerConstructorParams{})
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.MinerConstructorParams{}, &miner9.MinerConstructorParams{})
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner8.MinerConstructorParams{}, &miner8.MinerConstructorParams{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V15)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) ApplyRewards(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.ApplyRewardParams{}, &miner15.ApplyRewardParams{})
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.ApplyRewardParams{}, &miner14.ApplyRewardParams{})
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.ApplyRewardParams{}, &miner13.ApplyRewardParams{})
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.ApplyRewardParams{}, &miner12.ApplyRewardParams{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.ApplyRewardParams{}, &miner11.ApplyRewardParams{})
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.ApplyRewardParams{}, &miner10.ApplyRewardParams{})
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.ApplyRewardParams{}, &miner9.ApplyRewardParams{})
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner8.ApplyRewardParams{}, &miner8.ApplyRewardParams{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V15)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) OnDeferredCronEvent(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.DeferredCronEventParams{}, &miner15.DeferredCronEventParams{})
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.DeferredCronEventParams{}, &miner14.DeferredCronEventParams{})
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.DeferredCronEventParams{}, &miner13.DeferredCronEventParams{})
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.DeferredCronEventParams{}, &miner12.DeferredCronEventParams{})
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.DeferredCronEventParams{}, &miner11.DeferredCronEventParams{})
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.DeferredCronEventParams{}, &miner10.DeferredCronEventParams{})
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.DeferredCronEventParams{}, &miner9.DeferredCronEventParams{})
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner8.DeferredCronEventParams{}, &miner8.DeferredCronEventParams{})
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V15)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}
