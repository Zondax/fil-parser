package miner

import (
	"fmt"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

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
