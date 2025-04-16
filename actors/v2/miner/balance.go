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

	cbg "github.com/whyrusleeping/cbor-gen"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

func getAvailableBalanceReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V18.String(): &miner10.GetAvailableBalanceReturn{},

		tools.V19.String(): &miner11.GetAvailableBalanceReturn{},
		tools.V20.String(): &miner11.GetAvailableBalanceReturn{},

		tools.V21.String(): &miner12.GetAvailableBalanceReturn{},
		tools.V22.String(): &miner13.GetAvailableBalanceReturn{},
		tools.V23.String(): &miner14.GetAvailableBalanceReturn{},
		tools.V24.String(): &miner15.GetAvailableBalanceReturn{},
		tools.V25.String(): &miner16.GetAvailableBalanceReturn{},
	}
}

func getVestingFundsReturn() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V18.String(): &miner10.GetVestingFundsReturn{},

		tools.V19.String(): &miner11.GetVestingFundsReturn{},
		tools.V20.String(): &miner11.GetVestingFundsReturn{},

		tools.V21.String(): &miner12.GetVestingFundsReturn{},
		tools.V22.String(): &miner13.GetVestingFundsReturn{},
		tools.V23.String(): &miner14.GetVestingFundsReturn{},
		tools.V24.String(): &miner15.GetVestingFundsReturn{},
		tools.V25.String(): &miner16.GetVestingFundsReturn{},
	}
}

func getWithdrawBalanceParams() map[string]cbg.CBORUnmarshaler {
	return map[string]cbg.CBORUnmarshaler{
		tools.V7.String(): &legacyv1.WithdrawBalanceParams{},

		tools.V8.String(): &legacyv2.WithdrawBalanceParams{},
		tools.V9.String(): &legacyv2.WithdrawBalanceParams{},

		tools.V10.String(): &legacyv3.WithdrawBalanceParams{},
		tools.V11.String(): &legacyv3.WithdrawBalanceParams{},

		tools.V12.String(): &legacyv4.WithdrawBalanceParams{},
		tools.V13.String(): &legacyv5.WithdrawBalanceParams{},
		tools.V14.String(): &legacyv6.WithdrawBalanceParams{},
		tools.V15.String(): &legacyv7.WithdrawBalanceParams{},
		tools.V16.String(): &miner8.WithdrawBalanceParams{},
		tools.V17.String(): &miner9.WithdrawBalanceParams{},
		tools.V18.String(): &miner10.WithdrawBalanceParams{},

		tools.V19.String(): &miner11.WithdrawBalanceParams{},
		tools.V20.String(): &miner11.WithdrawBalanceParams{},

		tools.V21.String(): &miner12.WithdrawBalanceParams{},
		tools.V22.String(): &miner13.WithdrawBalanceParams{},
		tools.V23.String(): &miner14.WithdrawBalanceParams{},
		tools.V24.String(): &miner15.WithdrawBalanceParams{},
		tools.V25.String(): &miner16.WithdrawBalanceParams{},
	}
}

func (*Miner) GetAvailableBalanceExported(network string, height int64, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	returnValue, ok := getAvailableBalanceReturn()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawReturn, nil, false, returnValue, &abi.EmptyValue{}, parser.ReturnKey)
}

func (*Miner) GetVestingFundsExported(network string, height int64, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	returnValue, ok := getVestingFundsReturn()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawReturn, nil, false, returnValue, &abi.EmptyValue{}, parser.ReturnKey)
}

func (*Miner) WithdrawBalanceExported(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := getWithdrawBalanceParams()[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, nil, false, params, &abi.EmptyValue{}, parser.ParamsKey)
}

func (*Miner) AddLockedFund(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	return parseGeneric(rawParams, nil, false, &abi.TokenAmount{}, &abi.TokenAmount{}, parser.ParamsKey)
}
