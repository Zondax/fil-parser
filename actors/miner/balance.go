package miner

import (
	"fmt"

	miner15 "github.com/filecoin-project/go-state-types/builtin/v15/miner"
)

func GetAvailableBalance(height int64, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*miner15.GetAvailableBalanceReturn, *miner15.GetAvailableBalanceReturn](rawReturn, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func GetVestingFunds(height int64, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*miner15.GetVestingFundsReturn, *miner15.GetVestingFundsReturn](rawReturn, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func ParseWithdrawBalance(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch height {
	case 15:
		return parseGeneric[*miner15.WithdrawBalanceParams, *miner15.WithdrawBalanceParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}
