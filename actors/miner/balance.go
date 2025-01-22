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

func GetAvailableBalance(height int64, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		return parseGeneric[*miner15.GetAvailableBalanceReturn, *miner15.GetAvailableBalanceReturn](rawReturn, nil, false)
	case tools.V14.IsSupported(height):
		return parseGeneric[*miner14.GetAvailableBalanceReturn, *miner14.GetAvailableBalanceReturn](rawReturn, nil, false)
	case tools.V13.IsSupported(height):
		return parseGeneric[*miner13.GetAvailableBalanceReturn, *miner13.GetAvailableBalanceReturn](rawReturn, nil, false)
	case tools.V12.IsSupported(height):
		return parseGeneric[*miner12.GetAvailableBalanceReturn, *miner12.GetAvailableBalanceReturn](rawReturn, nil, false)
	case tools.V11.IsSupported(height):
		return parseGeneric[*miner11.GetAvailableBalanceReturn, *miner11.GetAvailableBalanceReturn](rawReturn, nil, false)
	case tools.V10.IsSupported(height):
		return parseGeneric[*miner10.GetAvailableBalanceReturn, *miner10.GetAvailableBalanceReturn](rawReturn, nil, false)
	case tools.V9.IsSupported(height):
		return nil, fmt.Errorf("not supported")
	case tools.V8.IsSupported(height):
		return nil, fmt.Errorf("not supported")
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func GetVestingFunds(height int64, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		return parseGeneric[*miner15.GetVestingFundsReturn, *miner15.GetVestingFundsReturn](rawReturn, nil, false)
	case tools.V14.IsSupported(height):
		return parseGeneric[*miner14.GetVestingFundsReturn, *miner14.GetVestingFundsReturn](rawReturn, nil, false)
	case tools.V13.IsSupported(height):
		return parseGeneric[*miner13.GetVestingFundsReturn, *miner13.GetVestingFundsReturn](rawReturn, nil, false)
	case tools.V12.IsSupported(height):
		return parseGeneric[*miner12.GetVestingFundsReturn, *miner12.GetVestingFundsReturn](rawReturn, nil, false)
	case tools.V11.IsSupported(height):
		return parseGeneric[*miner11.GetVestingFundsReturn, *miner11.GetVestingFundsReturn](rawReturn, nil, false)
	case tools.V10.IsSupported(height):
		return parseGeneric[*miner10.GetVestingFundsReturn, *miner10.GetVestingFundsReturn](rawReturn, nil, false)
	case tools.V9.IsSupported(height):
		return nil, fmt.Errorf("not supported")
	case tools.V8.IsSupported(height):
		return nil, fmt.Errorf("not supported")
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func ParseWithdrawBalance(height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V15.IsSupported(height):
		return parseGeneric[*miner15.WithdrawBalanceParams, *miner15.WithdrawBalanceParams](rawParams, nil, false)
	case tools.V14.IsSupported(height):
		return parseGeneric[*miner14.WithdrawBalanceParams, *miner14.WithdrawBalanceParams](rawParams, nil, false)
	case tools.V13.IsSupported(height):
		return parseGeneric[*miner13.WithdrawBalanceParams, *miner13.WithdrawBalanceParams](rawParams, nil, false)
	case tools.V12.IsSupported(height):
		return parseGeneric[*miner12.WithdrawBalanceParams, *miner12.WithdrawBalanceParams](rawParams, nil, false)
	case tools.V11.IsSupported(height):
		return parseGeneric[*miner11.WithdrawBalanceParams, *miner11.WithdrawBalanceParams](rawParams, nil, false)
	case tools.V10.IsSupported(height):
		return parseGeneric[*miner10.WithdrawBalanceParams, *miner10.WithdrawBalanceParams](rawParams, nil, false)
	case tools.V9.IsSupported(height):
		return parseGeneric[*miner9.WithdrawBalanceParams, *miner9.WithdrawBalanceParams](rawParams, nil, false)
	case tools.V8.IsSupported(height):
		return parseGeneric[*miner8.WithdrawBalanceParams, *miner8.WithdrawBalanceParams](rawParams, nil, false)
	}
	return nil, fmt.Errorf("unsupported height: %d", height)
}
