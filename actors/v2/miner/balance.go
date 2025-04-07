package miner

import (
	"fmt"

	"github.com/filecoin-project/go-state-types/abi"
	miner10 "github.com/filecoin-project/go-state-types/builtin/v10/miner"
	miner11 "github.com/filecoin-project/go-state-types/builtin/v11/miner"
	miner12 "github.com/filecoin-project/go-state-types/builtin/v12/miner"
	miner13 "github.com/filecoin-project/go-state-types/builtin/v13/miner"
	miner14 "github.com/filecoin-project/go-state-types/builtin/v14/miner"
	miner15 "github.com/filecoin-project/go-state-types/builtin/v15/miner"
	miner16 "github.com/filecoin-project/go-state-types/builtin/v16/miner"
	miner8 "github.com/filecoin-project/go-state-types/builtin/v8/miner"
	miner9 "github.com/filecoin-project/go-state-types/builtin/v9/miner"

	legacyv1 "github.com/filecoin-project/specs-actors/actors/builtin/miner"
	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/miner"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/miner"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/miner"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/miner"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/miner"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/miner"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

func (*Miner) GetAvailableBalanceExported(network string, height int64, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V25.IsSupported(network, height):
		return parseGeneric(rawReturn, nil, false, &miner16.GetAvailableBalanceReturn{}, &miner16.GetAvailableBalanceReturn{}, parser.ReturnKey)
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawReturn, nil, false, &miner15.GetAvailableBalanceReturn{}, &miner15.GetAvailableBalanceReturn{}, parser.ReturnKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawReturn, nil, false, &miner14.GetAvailableBalanceReturn{}, &miner14.GetAvailableBalanceReturn{}, parser.ReturnKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawReturn, nil, false, &miner13.GetAvailableBalanceReturn{}, &miner13.GetAvailableBalanceReturn{}, parser.ReturnKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawReturn, nil, false, &miner12.GetAvailableBalanceReturn{}, &miner12.GetAvailableBalanceReturn{}, parser.ReturnKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawReturn, nil, false, &miner11.GetAvailableBalanceReturn{}, &miner11.GetAvailableBalanceReturn{}, parser.ReturnKey)
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawReturn, nil, false, &miner10.GetAvailableBalanceReturn{}, &miner10.GetAvailableBalanceReturn{}, parser.ReturnKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V17)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) GetVestingFundsExported(network string, height int64, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V25.IsSupported(network, height):
		return parseGeneric(rawReturn, nil, false, &miner16.GetVestingFundsReturn{}, &miner16.GetVestingFundsReturn{}, parser.ReturnKey)
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawReturn, nil, false, &miner15.GetVestingFundsReturn{}, &miner15.GetVestingFundsReturn{}, parser.ReturnKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawReturn, nil, false, &miner14.GetVestingFundsReturn{}, &miner14.GetVestingFundsReturn{}, parser.ReturnKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawReturn, nil, false, &miner13.GetVestingFundsReturn{}, &miner13.GetVestingFundsReturn{}, parser.ReturnKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawReturn, nil, false, &miner12.GetVestingFundsReturn{}, &miner12.GetVestingFundsReturn{}, parser.ReturnKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawReturn, nil, false, &miner11.GetVestingFundsReturn{}, &miner11.GetVestingFundsReturn{}, parser.ReturnKey)
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawReturn, nil, false, &miner10.GetVestingFundsReturn{}, &miner10.GetVestingFundsReturn{}, parser.ReturnKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V17)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) WithdrawBalanceExported(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V25.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner16.WithdrawBalanceParams{}, &miner16.WithdrawBalanceParams{}, parser.ParamsKey)
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.WithdrawBalanceParams{}, &miner15.WithdrawBalanceParams{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.WithdrawBalanceParams{}, &miner14.WithdrawBalanceParams{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.WithdrawBalanceParams{}, &miner13.WithdrawBalanceParams{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.WithdrawBalanceParams{}, &miner12.WithdrawBalanceParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.WithdrawBalanceParams{}, &miner11.WithdrawBalanceParams{}, parser.ParamsKey)
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.WithdrawBalanceParams{}, &miner10.WithdrawBalanceParams{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.WithdrawBalanceParams{}, &miner9.WithdrawBalanceParams{}, parser.ParamsKey)
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner8.WithdrawBalanceParams{}, &miner8.WithdrawBalanceParams{}, parser.ParamsKey)
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv7.WithdrawBalanceParams{}, &legacyv7.WithdrawBalanceParams{}, parser.ParamsKey)
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv6.WithdrawBalanceParams{}, &legacyv6.WithdrawBalanceParams{}, parser.ParamsKey)
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv5.WithdrawBalanceParams{}, &legacyv5.WithdrawBalanceParams{}, parser.ParamsKey)
	case tools.V12.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv4.WithdrawBalanceParams{}, &legacyv4.WithdrawBalanceParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parseGeneric(rawParams, nil, false, &legacyv3.WithdrawBalanceParams{}, &legacyv3.WithdrawBalanceParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V8, tools.V9):
		return parseGeneric(rawParams, nil, false, &legacyv2.WithdrawBalanceParams{}, &legacyv2.WithdrawBalanceParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return parseGeneric(rawParams, nil, false, &legacyv1.WithdrawBalanceParams{}, &legacyv1.WithdrawBalanceParams{}, parser.ParamsKey)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) AddLockedFund(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	return parseGeneric(rawParams, nil, false, &abi.TokenAmount{}, &abi.TokenAmount{}, parser.ParamsKey)
}
