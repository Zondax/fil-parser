package common

import (
	"fmt"
	"math/big"

	"golang.org/x/exp/constraints"
)

func GetBigInt(value map[string]interface{}, key string, canBeNil bool) (*big.Int, error) {
	bigIntString, err := GetItem[string](value, key, canBeNil)
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

func GetInteger[T constraints.Integer](value map[string]interface{}, key string, canBeNil bool) (T, error) {
	valueAsFloat, err := GetItem[float64](value, key, canBeNil)
	if err != nil {
		return 0, err
	}

	return T(valueAsFloat), nil
}

func GetIntegerSlice[T constraints.Integer](value map[string]interface{}, key string, canBeNil bool) ([]T, error) {
	valuesAsFloat, err := GetSlice[float64](value, key, canBeNil)
	if err != nil {
		return nil, err
	}

	result := make([]T, len(valuesAsFloat))
	for i, v := range valuesAsFloat {
		result[i] = T(v)
	}

	return result, nil
}

func GetItem[T any](value map[string]interface{}, key string, canBeNil bool) (T, error) {
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

func GetSlice[T any](value map[string]interface{}, key string, canBeNil bool) ([]T, error) {
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
