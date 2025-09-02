package common

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-bitfield"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/fil-parser/types"
	"github.com/zondax/golem/pkg/logger"
	"golang.org/x/exp/constraints"
)

const (
	TxStatusOk = "ok"
)

func IsTxSuccess(tx *types.Transaction) bool {
	return strings.EqualFold(tx.Status, TxStatusOk) && strings.EqualFold(tx.SubcallStatus, TxStatusOk)
}

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

func GetCID(value map[string]interface{}, key string, canBeNil bool) (cid.Cid, error) {
	cidMap, err := GetItem[map[string]interface{}](value, key, canBeNil)
	if err != nil {
		return cid.Cid{}, err
	}
	if canBeNil && cidMap == nil {
		return cid.Cid{}, nil
	}
	cidString, err := GetItem[string](cidMap, "/", false)
	if err != nil {
		return cid.Cid{}, err
	}
	return cid.Decode(cidString)
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

// a bit field is a range of bits representing the different ids (numbers).
// example: ids: [ 0 1 2 3] -> bitfield: [1 1 1 1 ] -> JSON: [0,4]
// example: ids: [0 1 3 4 5] -> bitfield: [1 1 0 1 1 1] -> JSON: [ 0,2,1,3 ]
// the JSON format always starts with a 0 and proceeds with the 0/1 pattern
// see here: https://pkg.go.dev/github.com/filecoin-project/go-bitfield@v0.2.4/rle#RLE.MarshalJSON
func JsonEncodedBitfieldToIDs(bitField []int) ([]uint64, error) {
	ids := []uint64{}

	var parsedBitField bitfield.BitField
	bitFieldJSON, err := json.Marshal(bitField)
	if err != nil {
		return nil, fmt.Errorf("error marshaling json encoded bitfield: %w", err)
	}

	err = parsedBitField.UnmarshalJSON(bitFieldJSON)
	if err != nil {
		return nil, fmt.Errorf("error parsing json encoded bitfield: %w", err)
	}

	iter, err := parsedBitField.BitIterator()
	if err != nil {
		return nil, fmt.Errorf("error iterating over bitfield: %w", err)
	}

	for iter.HasNext() {
		id, err := iter.Next()
		if err != nil {
			return nil, fmt.Errorf("error getting next id: %w", err)
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func ConsolidateIDAddress(idAddress uint64, helper *helper.Helper, logger *logger.Logger, config parser.Config) (string, error) {
	addr, err := address.NewIDAddress(idAddress)
	if err != nil {
		return "", fmt.Errorf("error parsing id address: %w", err)
	}
	if config.ConsolidateRobustAddress {
		consolidatedIDAddress, err := actors.ConsolidateToRobustAddress(addr, helper, logger, config.RobustAddressBestEffort)
		if err != nil {
			return addr.String(), fmt.Errorf("error consolidating id address: %w", err)
		}
		return consolidatedIDAddress, nil
	}
	return addr.String(), nil
}

func ConsolidateAddress(addrStr string, helper *helper.Helper, logger *logger.Logger, config parser.Config) (string, error) {
	if config.ConsolidateRobustAddress {
		addr, err := address.NewFromString(addrStr)
		if err != nil {
			return addrStr, fmt.Errorf("error parsing address: %w", err)
		}
		consolidatedAddress, err := actors.ConsolidateToRobustAddress(addr, helper, logger, config.RobustAddressBestEffort)
		if err != nil {
			return addrStr, fmt.Errorf("error consolidating address: %w", err)
		}
		return consolidatedAddress, nil
	}
	return addrStr, nil
}
