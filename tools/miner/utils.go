package miner

import (
	"fmt"
	"math/big"

	"github.com/filecoin-project/go-state-types/abi"
	"golang.org/x/exp/constraints"
)

func getBigInt(value map[string]interface{}, key string, canBeNil bool) (*big.Int, error) {
	bigIntString, err := getItem[string](value, key, canBeNil)
	if err != nil {
		return nil, err
	}
	if canBeNil && bigIntString == "" {
		return nil, nil
	}
	bigIntValue, ok := big.NewInt(0).SetString(bigIntString, 10)
	if !ok {
		return nil, fmt.Errorf("failed to convert string %s to big.Int", bigIntString)
	}
	return bigIntValue, nil
}

func getInteger[T constraints.Integer](value map[string]interface{}, key string, canBeNil bool) (T, error) {
	valueAsFloat, err := getItem[float64](value, key, canBeNil)
	if err != nil {
		return 0, err
	}

	return T(valueAsFloat), nil
}

func getIntegerSlice[T constraints.Integer](value map[string]interface{}, key string, canBeNil bool) ([]T, error) {
	valuesAsFloat, err := getSlice[float64](value, key, canBeNil)
	if err != nil {
		return nil, err
	}

	result := make([]T, len(valuesAsFloat))
	for i, v := range valuesAsFloat {
		result[i] = T(v)
	}

	return result, nil
}

func getItem[T any](value map[string]interface{}, key string, canBeNil bool) (T, error) {
	var zero T
	if value == nil {
		if canBeNil {
			return zero, nil
		}
		return zero, fmt.Errorf("value is nil")
	}
	if value[key] == nil {
		if canBeNil {
			return zero, nil
		}
		return zero, fmt.Errorf("key %s not found", key)
	}
	if v, ok := value[key].(T); ok {
		return v, nil
	}
	return zero, fmt.Errorf("key %s not of type %T", key, zero)
}

func getSlice[T any](value map[string]interface{}, key string, canBeNil bool) ([]T, error) {
	var result []T
	if value == nil {
		if canBeNil {
			return result, nil
		}
		return result, fmt.Errorf("value is nil")
	}

	if value[key] == nil {
		if canBeNil {
			return result, nil
		}
		return result, fmt.Errorf("key %s not found", key)
	}

	if v, ok := value[key].([]interface{}); ok {
		for _, item := range v {
			if tmp, ok := item.(T); ok {
				result = append(result, tmp)
			} else {
				return result, fmt.Errorf("item %v not of type %T", item, result)
			}
		}
		return result, nil
	}
	return result, fmt.Errorf("key %s not of type %T , of type %T", key, result, value[key])
}

func sectorProofToBigInt(sectorProof int64) *big.Int {
	proofType := abi.RegisteredSealProof(sectorProof)
	info, ok := abi.SealProofInfos[proofType]
	if !ok {
		fmt.Println("invalid proofType", proofType)
		return big.NewInt(0)
	}

	return big.NewInt(int64(info.SectorSize))
}
