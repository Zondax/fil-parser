package miner

import (
	"bytes"
	"fmt"
	"io"

	miner10 "github.com/filecoin-project/go-state-types/builtin/v10/miner"
	miner11 "github.com/filecoin-project/go-state-types/builtin/v11/miner"
	miner12 "github.com/filecoin-project/go-state-types/builtin/v12/miner"
	miner13 "github.com/filecoin-project/go-state-types/builtin/v13/miner"
	miner14 "github.com/filecoin-project/go-state-types/builtin/v14/miner"
	miner15 "github.com/filecoin-project/go-state-types/builtin/v15/miner"
	miner8 "github.com/filecoin-project/go-state-types/builtin/v8/miner"
	miner9 "github.com/filecoin-project/go-state-types/builtin/v9/miner"
	"github.com/zondax/fil-parser/parser"
)

type minerParam interface {
	UnmarshalCBOR(io.Reader) error
}

type minerReturn interface {
	UnmarshalCBOR(io.Reader) error
}

func parseGeneric[T minerParam, R minerReturn](rawParams, rawReturn []byte, customReturn bool) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params T
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	if !customReturn {
		return metadata, nil
	}
	reader = bytes.NewReader(rawReturn)
	var r R
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r
	return metadata, nil
}

func TerminateSectors(height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*miner15.TerminateSectorsParams, *miner15.TerminateSectorsReturn](rawParams, rawReturn, true)
	case 14:
		return parseGeneric[*miner14.TerminateSectorsParams, *miner14.TerminateSectorsReturn](rawParams, rawReturn, true)
	case 13:
		return parseGeneric[*miner13.TerminateSectorsParams, *miner13.TerminateSectorsReturn](rawParams, rawReturn, true)
	case 12:
		return parseGeneric[*miner12.TerminateSectorsParams, *miner12.TerminateSectorsReturn](rawParams, rawReturn, true)
	case 11:
		return parseGeneric[*miner11.TerminateSectorsParams, *miner11.TerminateSectorsReturn](rawParams, rawReturn, true)
	case 10:
		return parseGeneric[*miner10.TerminateSectorsParams, *miner10.TerminateSectorsReturn](rawParams, rawReturn, true)
	case 9:
		return parseGeneric[*miner9.TerminateSectorsParams, *miner9.TerminateSectorsReturn](rawParams, rawReturn, true)
	case 8:
		return parseGeneric[*miner8.TerminateSectorsParams, *miner8.TerminateSectorsReturn](rawParams, rawReturn, true)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func DeclareFaults(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*miner15.DeclareFaultsParams, *miner15.DeclareFaultsParams](rawParams, nil, false)
	case 14:
		return parseGeneric[*miner14.DeclareFaultsParams, *miner14.DeclareFaultsParams](rawParams, nil, false)
	case 13:
		return parseGeneric[*miner13.DeclareFaultsParams, *miner13.DeclareFaultsParams](rawParams, nil, false)
	case 12:
		return parseGeneric[*miner12.DeclareFaultsParams, *miner12.DeclareFaultsParams](rawParams, nil, false)
	case 11:
		return parseGeneric[*miner11.DeclareFaultsParams, *miner11.DeclareFaultsParams](rawParams, nil, false)
	case 10:
		return parseGeneric[*miner10.DeclareFaultsParams, *miner10.DeclareFaultsParams](rawParams, nil, false)
	case 9:
		return parseGeneric[*miner9.DeclareFaultsParams, *miner9.DeclareFaultsParams](rawParams, nil, false)
	case 8:
		return parseGeneric[*miner8.DeclareFaultsParams, *miner8.DeclareFaultsParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func DeclareFaultsRecovered(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*miner15.DeclareFaultsRecoveredParams, *miner15.DeclareFaultsRecoveredParams](rawParams, nil, false)
	case 14:
		return parseGeneric[*miner14.DeclareFaultsRecoveredParams, *miner14.DeclareFaultsRecoveredParams](rawParams, nil, false)
	case 13:
		return parseGeneric[*miner13.DeclareFaultsRecoveredParams, *miner13.DeclareFaultsRecoveredParams](rawParams, nil, false)
	case 12:
		return parseGeneric[*miner12.DeclareFaultsRecoveredParams, *miner12.DeclareFaultsRecoveredParams](rawParams, nil, false)
	case 11:
		return parseGeneric[*miner11.DeclareFaultsRecoveredParams, *miner11.DeclareFaultsRecoveredParams](rawParams, nil, false)
	case 10:
		return parseGeneric[*miner10.DeclareFaultsRecoveredParams, *miner10.DeclareFaultsRecoveredParams](rawParams, nil, false)
	case 9:
		return parseGeneric[*miner9.DeclareFaultsRecoveredParams, *miner9.DeclareFaultsRecoveredParams](rawParams, nil, false)
	case 8:
		return parseGeneric[*miner8.DeclareFaultsRecoveredParams, *miner8.DeclareFaultsRecoveredParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func ProveReplicaUpdates(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*miner15.ProveReplicaUpdatesParams, *miner15.ProveReplicaUpdatesParams](rawParams, nil, false)
	case 14:
		return parseGeneric[*miner14.ProveReplicaUpdatesParams, *miner14.ProveReplicaUpdatesParams](rawParams, nil, false)
	case 13:
		return parseGeneric[*miner13.ProveReplicaUpdatesParams, *miner13.ProveReplicaUpdatesParams](rawParams, nil, false)
	case 12:
		return parseGeneric[*miner12.ProveReplicaUpdatesParams, *miner12.ProveReplicaUpdatesParams](rawParams, nil, false)
	case 11:
		return parseGeneric[*miner11.ProveReplicaUpdatesParams, *miner11.ProveReplicaUpdatesParams](rawParams, nil, false)
	case 10:
		return parseGeneric[*miner10.ProveReplicaUpdatesParams, *miner10.ProveReplicaUpdatesParams](rawParams, nil, false)
	case 9:
		return parseGeneric[*miner9.ProveReplicaUpdatesParams, *miner9.ProveReplicaUpdatesParams](rawParams, nil, false)
	case 8:
		return parseGeneric[*miner8.ProveReplicaUpdatesParams, *miner8.ProveReplicaUpdatesParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func PreCommitSectorBatch2(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*miner15.PreCommitSectorBatchParams2, *miner15.PreCommitSectorBatchParams2](rawParams, nil, false)
	case 14:
		return parseGeneric[*miner14.PreCommitSectorBatchParams2, *miner14.PreCommitSectorBatchParams2](rawParams, nil, false)
	case 13:
		return parseGeneric[*miner13.PreCommitSectorBatchParams2, *miner13.PreCommitSectorBatchParams2](rawParams, nil, false)
	case 12:
		return parseGeneric[*miner12.PreCommitSectorBatchParams2, *miner12.PreCommitSectorBatchParams2](rawParams, nil, false)
	case 11:
		return parseGeneric[*miner11.PreCommitSectorBatchParams2, *miner11.PreCommitSectorBatchParams2](rawParams, nil, false)
	case 10:
		return parseGeneric[*miner10.PreCommitSectorBatchParams2, *miner10.PreCommitSectorBatchParams2](rawParams, nil, false)
	case 9:
		return parseGeneric[*miner9.PreCommitSectorBatchParams2, *miner9.PreCommitSectorBatchParams2](rawParams, nil, false)
	case 8:
		// return parseGeneric[*miner8.PreCommitSectorBatchParams2, *miner8.PreCommitSectorBatchParams2](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func ProveReplicaUpdates2(height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*miner15.ProveReplicaUpdatesParams2, *miner15.ProveReplicaUpdatesParams2](rawParams, rawReturn, true)
	case 14:
		return parseGeneric[*miner14.ProveReplicaUpdatesParams2, *miner14.ProveReplicaUpdatesParams2](rawParams, rawReturn, true)
	case 13:
		return parseGeneric[*miner13.ProveReplicaUpdatesParams2, *miner13.ProveReplicaUpdatesParams2](rawParams, rawReturn, true)
	case 12:
		return parseGeneric[*miner12.ProveReplicaUpdatesParams2, *miner12.ProveReplicaUpdatesParams2](rawParams, rawReturn, true)
	case 11:
		return parseGeneric[*miner11.ProveReplicaUpdatesParams2, *miner11.ProveReplicaUpdatesParams2](rawParams, rawReturn, true)
	case 10:
		return parseGeneric[*miner10.ProveReplicaUpdatesParams2, *miner10.ProveReplicaUpdatesParams2](rawParams, rawReturn, true)
	case 9:
		return parseGeneric[*miner9.ProveReplicaUpdatesParams2, *miner9.ProveReplicaUpdatesParams2](rawParams, rawReturn, true)
	case 8:
		// return parseGeneric[*miner8.ProveReplicaUpdatesParams2, *miner8.ProveReplicaUpdatesParams2](rawParams, rawReturn, true)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func ProveCommitAggregate(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*miner15.ProveCommitAggregateParams, *miner15.ProveCommitAggregateParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func DisputeWindowedPoSt(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*miner15.DisputeWindowedPoStParams, *miner15.DisputeWindowedPoStParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func ReportConsensusFault(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*miner15.ReportConsensusFaultParams, *miner15.ReportConsensusFaultParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func ChangeBeneficiary(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*miner15.ChangeBeneficiaryParams, *miner15.ChangeBeneficiaryParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func MinerConstructor(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*miner15.MinerConstructorParams, *miner15.MinerConstructorParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func ApplyRewards(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*miner15.ApplyRewardParams, *miner15.ApplyRewardParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func OnDeferredCronEvent(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*miner15.DeferredCronEventParams, *miner15.DeferredCronEventParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}
