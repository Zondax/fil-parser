package miner

import (
	"fmt"

	miner15 "github.com/filecoin-project/go-state-types/builtin/v15/miner"
)

func ExtendSectorExpiration2(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*miner15.ExtendSectorExpiration2Params, *miner15.ExtendSectorExpiration2Params](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func PreCommitSector(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*miner15.PreCommitSectorParams, *miner15.PreCommitSectorParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func ProveCommitSector(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*miner15.ProveCommitSectorParams, *miner15.ProveCommitSectorParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func ProveCommitSectors3(height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*miner15.ProveCommitSectors3Params, *miner15.ProveCommitSectors3Return](rawParams, rawReturn, true)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func SubmitWindowedPoSt(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*miner15.SubmitWindowedPoStParams, *miner15.SubmitWindowedPoStParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func ConfirmSectorProofsValid(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*miner15.ConfirmSectorProofsParams, *miner15.ConfirmSectorProofsParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func CheckSectorProven(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*miner15.CheckSectorProvenParams, *miner15.CheckSectorProvenParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func ExtendSectorExpiration(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*miner15.ExtendSectorExpirationParams, *miner15.ExtendSectorExpirationParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func CompactSectorNumbers(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*miner15.CompactSectorNumbersParams, *miner15.CompactSectorNumbersParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func CompactPartitions(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*miner15.CompactPartitionsParams, *miner15.CompactPartitionsParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func PreCommitSectorBatch(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*miner15.PreCommitSectorBatchParams, *miner15.PreCommitSectorBatchParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func GetSectorSize(height int64, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*miner15.GetSectorSizeReturn, *miner15.GetSectorSizeReturn](rawReturn, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}
