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
	"github.com/filecoin-project/lotus/chain/actors/builtin/tools"
)

func ExtendSectorExpiration2(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		return parseGeneric[*miner15.ExtendSectorExpiration2Params, *miner15.ExtendSectorExpiration2Params](rawParams, nil, false)
	case tools.V14.IsSupported(height):
		return parseGeneric[*miner14.ExtendSectorExpiration2Params, *miner14.ExtendSectorExpiration2Params](rawParams, nil, false)
	case tools.V13.IsSupported(height):
		return parseGeneric[*miner13.ExtendSectorExpiration2Params, *miner13.ExtendSectorExpiration2Params](rawParams, nil, false)
	case tools.V12.IsSupported(height):
		return parseGeneric[*miner12.ExtendSectorExpiration2Params, *miner12.ExtendSectorExpiration2Params](rawParams, nil, false)
	case tools.V11.IsSupported(height):
		return parseGeneric[*miner11.ExtendSectorExpiration2Params, *miner11.ExtendSectorExpiration2Params](rawParams, nil, false)
	case tools.V10.IsSupported(height):
		return parseGeneric[*miner10.ExtendSectorExpiration2Params, *miner10.ExtendSectorExpiration2Params](rawParams, nil, false)
	case tools.V9.IsSupported(height):
		return parseGeneric[*miner9.ExtendSectorExpiration2Params, *miner9.ExtendSectorExpiration2Params](rawParams, nil, false)
	case tools.V8.IsSupported(height):
		return parseGeneric[*miner8.ExtendSectorExpiration2Params, *miner8.ExtendSectorExpiration2Params](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func PreCommitSector(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		return parseGeneric[*miner15.PreCommitSectorParams, *miner15.PreCommitSectorParams](rawParams, nil, false)
	case tools.V14.IsSupported(height):
		return parseGeneric[*miner14.PreCommitSectorParams, *miner14.PreCommitSectorParams](rawParams, nil, false)
	case tools.V13.IsSupported(height):
		return parseGeneric[*miner13.PreCommitSectorParams, *miner13.PreCommitSectorParams](rawParams, nil, false)
	case tools.V12.IsSupported(height):
		return parseGeneric[*miner12.PreCommitSectorParams, *miner12.PreCommitSectorParams](rawParams, nil, false)
	case tools.V11.IsSupported(height):
		return parseGeneric[*miner11.PreCommitSectorParams, *miner11.PreCommitSectorParams](rawParams, nil, false)
	case tools.V10.IsSupported(height):
		return parseGeneric[*miner10.PreCommitSectorParams, *miner10.PreCommitSectorParams](rawParams, nil, false)
	case tools.V9.IsSupported(height):
		return parseGeneric[*miner9.PreCommitSectorParams, *miner9.PreCommitSectorParams](rawParams, nil, false)
	case tools.V8.IsSupported(height):
		return parseGeneric[*miner8.PreCommitSectorParams, *miner8.PreCommitSectorParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func ProveCommitSector(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		return parseGeneric[*miner15.ProveCommitSectorParams, *miner15.ProveCommitSectorParams](rawParams, nil, false)
	case tools.V14.IsSupported(height):
		return parseGeneric[*miner14.ProveCommitSectorParams, *miner14.ProveCommitSectorParams](rawParams, nil, false)
	case tools.V13.IsSupported(height):
		return parseGeneric[*miner13.ProveCommitSectorParams, *miner13.ProveCommitSectorParams](rawParams, nil, false)
	case tools.V12.IsSupported(height):
		return parseGeneric[*miner12.ProveCommitSectorParams, *miner12.ProveCommitSectorParams](rawParams, nil, false)
	case tools.V11.IsSupported(height):
		return parseGeneric[*miner11.ProveCommitSectorParams, *miner11.ProveCommitSectorParams](rawParams, nil, false)
	case tools.V10.IsSupported(height):
		return parseGeneric[*miner10.ProveCommitSectorParams, *miner10.ProveCommitSectorParams](rawParams, nil, false)
	case tools.V9.IsSupported(height):
		return parseGeneric[*miner9.ProveCommitSectorParams, *miner9.ProveCommitSectorParams](rawParams, nil, false)
	case tools.V8.IsSupported(height):
		return parseGeneric[*miner8.ProveCommitSectorParams, *miner8.ProveCommitSectorParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func ProveCommitSectors3(height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		return parseGeneric[*miner15.ProveCommitSectors3Params, *miner15.ProveCommitSectors3Return](rawParams, rawReturn, true)
	case tools.V14.IsSupported(height):
		return parseGeneric[*miner14.ProveCommitSectors3Params, *miner14.ProveCommitSectors3Return](rawParams, rawReturn, true)
	case tools.V13.IsSupported(height):
		return parseGeneric[*miner13.ProveCommitSectors3Params, *miner13.ProveCommitSectors3Return](rawParams, rawReturn, true)
	case tools.V12.IsSupported(height):
		return parseGeneric[*miner12.ProveCommitSectors3Params, *miner12.ProveCommitSectors3Return](rawParams, rawReturn, true)
	case tools.V11.IsSupported(height):
		return parseGeneric[*miner11.ProveCommitSectors3Params, *miner11.ProveCommitSectors3Return](rawParams, rawReturn, true)
	case tools.V10.IsSupported(height):
		return parseGeneric[*miner10.ProveCommitSectors3Params, *miner10.ProveCommitSectors3Return](rawParams, rawReturn, true)
	case tools.V9.IsSupported(height):
		return parseGeneric[*miner9.ProveCommitSectors3Params, *miner9.ProveCommitSectors3Return](rawParams, rawReturn, true)
	case tools.V8.IsSupported(height):
		return parseGeneric[*miner8.ProveCommitSectors3Params, *miner8.ProveCommitSectors3Return](rawParams, rawReturn, true)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func SubmitWindowedPoSt(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		return parseGeneric[*miner15.SubmitWindowedPoStParams, *miner15.SubmitWindowedPoStParams](rawParams, nil, false)
	case tools.V14.IsSupported(height):
		return parseGeneric[*miner14.SubmitWindowedPoStParams, *miner14.SubmitWindowedPoStParams](rawParams, nil, false)
	case tools.V13.IsSupported(height):
		return parseGeneric[*miner13.SubmitWindowedPoStParams, *miner13.SubmitWindowedPoStParams](rawParams, nil, false)
	case tools.V12.IsSupported(height):
		return parseGeneric[*miner12.SubmitWindowedPoStParams, *miner12.SubmitWindowedPoStParams](rawParams, nil, false)
	case tools.V11.IsSupported(height):
		return parseGeneric[*miner11.SubmitWindowedPoStParams, *miner11.SubmitWindowedPoStParams](rawParams, nil, false)
	case tools.V10.IsSupported(height):
		return parseGeneric[*miner10.SubmitWindowedPoStParams, *miner10.SubmitWindowedPoStParams](rawParams, nil, false)
	case tools.V9.IsSupported(height):
		return parseGeneric[*miner9.SubmitWindowedPoStParams, *miner9.SubmitWindowedPoStParams](rawParams, nil, false)
	case tools.V8.IsSupported(height):
		return parseGeneric[*miner8.SubmitWindowedPoStParams, *miner8.SubmitWindowedPoStParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func ConfirmSectorProofsValid(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		return parseGeneric[*miner15.ConfirmSectorProofsParams, *miner15.ConfirmSectorProofsParams](rawParams, nil, false)
	case tools.V14.IsSupported(height):
		return parseGeneric[*miner14.ConfirmSectorProofsParams, *miner14.ConfirmSectorProofsParams](rawParams, nil, false)
	case tools.V13.IsSupported(height):
		return parseGeneric[*miner13.ConfirmSectorProofsParams, *miner13.ConfirmSectorProofsParams](rawParams, nil, false)
	case tools.V12.IsSupported(height):
		return parseGeneric[*miner12.ConfirmSectorProofsParams, *miner12.ConfirmSectorProofsParams](rawParams, nil, false)
	case tools.V11.IsSupported(height):
		return parseGeneric[*miner11.ConfirmSectorProofsParams, *miner11.ConfirmSectorProofsParams](rawParams, nil, false)
	case tools.V10.IsSupported(height):
		return parseGeneric[*miner10.ConfirmSectorProofsParams, *miner10.ConfirmSectorProofsParams](rawParams, nil, false)
	case tools.V9.IsSupported(height):
		return parseGeneric[*miner9.ConfirmSectorProofsParams, *miner9.ConfirmSectorProofsParams](rawParams, nil, false)
	case tools.V8.IsSupported(height):
		return parseGeneric[*miner8.ConfirmSectorProofsParams, *miner8.ConfirmSectorProofsParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func CheckSectorProven(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		return parseGeneric[*miner15.CheckSectorProvenParams, *miner15.CheckSectorProvenParams](rawParams, nil, false)
	case tools.V14.IsSupported(height):
		return parseGeneric[*miner14.CheckSectorProvenParams, *miner14.CheckSectorProvenParams](rawParams, nil, false)
	case tools.V13.IsSupported(height):
		return parseGeneric[*miner13.CheckSectorProvenParams, *miner13.CheckSectorProvenParams](rawParams, nil, false)
	case tools.V12.IsSupported(height):
		return parseGeneric[*miner12.CheckSectorProvenParams, *miner12.CheckSectorProvenParams](rawParams, nil, false)
	case tools.V11.IsSupported(height):
		return parseGeneric[*miner11.CheckSectorProvenParams, *miner11.CheckSectorProvenParams](rawParams, nil, false)
	case tools.V10.IsSupported(height):
		return parseGeneric[*miner10.CheckSectorProvenParams, *miner10.CheckSectorProvenParams](rawParams, nil, false)
	case tools.V9.IsSupported(height):
		return parseGeneric[*miner9.CheckSectorProvenParams, *miner9.CheckSectorProvenParams](rawParams, nil, false)
	case tools.V8.IsSupported(height):
		return parseGeneric[*miner8.CheckSectorProvenParams, *miner8.CheckSectorProvenParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func ExtendSectorExpiration(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		return parseGeneric[*miner15.ExtendSectorExpirationParams, *miner15.ExtendSectorExpirationParams](rawParams, nil, false)
	case tools.V14.IsSupported(height):
		return parseGeneric[*miner14.ExtendSectorExpirationParams, *miner14.ExtendSectorExpirationParams](rawParams, nil, false)
	case tools.V13.IsSupported(height):
		return parseGeneric[*miner13.ExtendSectorExpirationParams, *miner13.ExtendSectorExpirationParams](rawParams, nil, false)
	case tools.V12.IsSupported(height):
		return parseGeneric[*miner12.ExtendSectorExpirationParams, *miner12.ExtendSectorExpirationParams](rawParams, nil, false)
	case tools.V11.IsSupported(height):
		return parseGeneric[*miner11.ExtendSectorExpirationParams, *miner11.ExtendSectorExpirationParams](rawParams, nil, false)
	case tools.V10.IsSupported(height):
		return parseGeneric[*miner10.ExtendSectorExpirationParams, *miner10.ExtendSectorExpirationParams](rawParams, nil, false)
	case tools.V9.IsSupported(height):
		return parseGeneric[*miner9.ExtendSectorExpirationParams, *miner9.ExtendSectorExpirationParams](rawParams, nil, false)
	case tools.V8.IsSupported(height):
		return parseGeneric[*miner8.ExtendSectorExpirationParams, *miner8.ExtendSectorExpirationParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func CompactSectorNumbers(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		return parseGeneric[*miner15.CompactSectorNumbersParams, *miner15.CompactSectorNumbersParams](rawParams, nil, false)
	case tools.V14.IsSupported(height):
		return parseGeneric[*miner14.CompactSectorNumbersParams, *miner14.CompactSectorNumbersParams](rawParams, nil, false)
	case tools.V13.IsSupported(height):
		return parseGeneric[*miner13.CompactSectorNumbersParams, *miner13.CompactSectorNumbersParams](rawParams, nil, false)
	case tools.V12.IsSupported(height):
		return parseGeneric[*miner12.CompactSectorNumbersParams, *miner12.CompactSectorNumbersParams](rawParams, nil, false)
	case tools.V11.IsSupported(height):
		return parseGeneric[*miner11.CompactSectorNumbersParams, *miner11.CompactSectorNumbersParams](rawParams, nil, false)
	case tools.V10.IsSupported(height):
		return parseGeneric[*miner10.CompactSectorNumbersParams, *miner10.CompactSectorNumbersParams](rawParams, nil, false)
	case tools.V9.IsSupported(height):
		return parseGeneric[*miner9.CompactSectorNumbersParams, *miner9.CompactSectorNumbersParams](rawParams, nil, false)
	case tools.V8.IsSupported(height):
		return parseGeneric[*miner8.CompactSectorNumbersParams, *miner8.CompactSectorNumbersParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func CompactPartitions(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		return parseGeneric[*miner15.CompactPartitionsParams, *miner15.CompactPartitionsParams](rawParams, nil, false)
	case tools.V14.IsSupported(height):
		return parseGeneric[*miner14.CompactPartitionsParams, *miner14.CompactPartitionsParams](rawParams, nil, false)
	case tools.V13.IsSupported(height):
		return parseGeneric[*miner13.CompactPartitionsParams, *miner13.CompactPartitionsParams](rawParams, nil, false)
	case tools.V12.IsSupported(height):
		return parseGeneric[*miner12.CompactPartitionsParams, *miner12.CompactPartitionsParams](rawParams, nil, false)
	case tools.V11.IsSupported(height):
		return parseGeneric[*miner11.CompactPartitionsParams, *miner11.CompactPartitionsParams](rawParams, nil, false)
	case tools.V10.IsSupported(height):
		return parseGeneric[*miner10.CompactPartitionsParams, *miner10.CompactPartitionsParams](rawParams, nil, false)
	case tools.V9.IsSupported(height):
		return parseGeneric[*miner9.CompactPartitionsParams, *miner9.CompactPartitionsParams](rawParams, nil, false)
	case tools.V8.IsSupported(height):
		return parseGeneric[*miner8.CompactPartitionsParams, *miner8.CompactPartitionsParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func PreCommitSectorBatch(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		return parseGeneric[*miner15.PreCommitSectorBatchParams, *miner15.PreCommitSectorBatchParams](rawParams, nil, false)
	case tools.V14.IsSupported(height):
		return parseGeneric[*miner14.PreCommitSectorBatchParams, *miner14.PreCommitSectorBatchParams](rawParams, nil, false)
	case tools.V13.IsSupported(height):
		return parseGeneric[*miner13.PreCommitSectorBatchParams, *miner13.PreCommitSectorBatchParams](rawParams, nil, false)
	case tools.V12.IsSupported(height):
		return parseGeneric[*miner12.PreCommitSectorBatchParams, *miner12.PreCommitSectorBatchParams](rawParams, nil, false)
	case tools.V11.IsSupported(height):
		return parseGeneric[*miner11.PreCommitSectorBatchParams, *miner11.PreCommitSectorBatchParams](rawParams, nil, false)
	case tools.V10.IsSupported(height):
		return parseGeneric[*miner10.PreCommitSectorBatchParams, *miner10.PreCommitSectorBatchParams](rawParams, nil, false)
	case tools.V9.IsSupported(height):
		return parseGeneric[*miner9.PreCommitSectorBatchParams, *miner9.PreCommitSectorBatchParams](rawParams, nil, false)
	case tools.V8.IsSupported(height):
		return parseGeneric[*miner8.PreCommitSectorBatchParams, *miner8.PreCommitSectorBatchParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func GetSectorSize(height int64, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		return parseGeneric[*miner15.GetSectorSizeReturn, *miner15.GetSectorSizeReturn](rawReturn, nil, false)
	case tools.V14.IsSupported(height):
		return parseGeneric[*miner14.GetSectorSizeReturn, *miner14.GetSectorSizeReturn](rawReturn, nil, false)
	case tools.V13.IsSupported(height):
		return parseGeneric[*miner13.GetSectorSizeReturn, *miner13.GetSectorSizeReturn](rawReturn, nil, false)
	case tools.V12.IsSupported(height):
		return parseGeneric[*miner12.GetSectorSizeReturn, *miner12.GetSectorSizeReturn](rawReturn, nil, false)
	case tools.V11.IsSupported(height):
		return parseGeneric[*miner11.GetSectorSizeReturn, *miner11.GetSectorSizeReturn](rawReturn, nil, false)
	case tools.V10.IsSupported(height):
		return parseGeneric[*miner10.GetSectorSizeReturn, *miner10.GetSectorSizeReturn](rawReturn, nil, false)
	case tools.V9.IsSupported(height):
		return parseGeneric[*miner9.GetSectorSizeReturn, *miner9.GetSectorSizeReturn](rawReturn, nil, false)
	case tools.V8.IsSupported(height):
		return parseGeneric[*miner8.GetSectorSizeReturn, *miner8.GetSectorSizeReturn](rawReturn, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}
