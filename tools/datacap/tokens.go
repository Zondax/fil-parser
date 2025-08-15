package datacap

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/tools/common"
	"github.com/zondax/fil-parser/types"
)

const (
	KeyAmount      = "Amount"
	KeyTo          = "To"
	KeyFrom        = "From"
	KeyOwner       = "Owner"
	KeyAllowance   = "Allowance"
	KeyBalance     = "Balance"
	KeySupply      = "Supply"
	KeyOperator    = "Operator"
	KeyFromBalance = "FromBalance"
	KeyToBalance   = "ToBalance"
)

func (eg *eventGenerator) createDataCapTokenEvents(ctx context.Context, tx *types.Transaction, tipsetCid string) ([]*types.DataCapTokenEvent, []*types.DataCapAllowanceEvent, error) {
	tokenEvents := []*types.DataCapTokenEvent{}
	allowanceEvents := []*types.DataCapAllowanceEvent{}

	var value map[string]interface{}
	err := json.Unmarshal([]byte(tx.TxMetadata), &value)
	if err != nil {
		return nil, nil, fmt.Errorf("error unmarshalling tx metadata: %w", err)
	}

	params, err := common.GetItem[map[string]interface{}](value, parser.ParamsKey, false)
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing params: %w", err)
	}

	switch tx.TxType {
	case parser.MethodMint, parser.MethodMintExported:
		returnValue, err := getReturn[map[string]interface{}](value)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing return value: %w", err)
		}

		tokenEvent, err := eg.parseMint(ctx, tx, tipsetCid, params, returnValue)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing mint: %w", err)
		}
		tokenEvents = append(tokenEvents, tokenEvent)
	case parser.MethodDestroy, parser.MethodDestroyExported:
		returnValue, err := getReturn[map[string]interface{}](value)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing return value: %w", err)
		}

		tokenEvent, err := eg.parseDestroy(ctx, tx, tipsetCid, params, returnValue)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing destroy: %w", err)
		}
		tokenEvents = append(tokenEvents, tokenEvent)
	case parser.MethodTransfer, parser.MethodTransferExported:
		returnValue, err := getReturn[map[string]interface{}](value)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing return value: %w", err)
		}

		tokenEvent, err := eg.parseTransfer(ctx, tx, tipsetCid, params, returnValue)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing transfer: %w", err)
		}
		tokenEvents = append(tokenEvents, tokenEvent...)
	case parser.MethodTransferFrom, parser.MethodTransferFromExported:
		returnValue, err := getReturn[map[string]interface{}](value)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing return value: %w", err)
		}

		tokenEvent, allowanceEvent, err := eg.parseTransferFrom(ctx, tx, tipsetCid, params, returnValue)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing transfer from: %w", err)
		}
		tokenEvents = append(tokenEvents, tokenEvent...)
		allowanceEvents = append(allowanceEvents, allowanceEvent)
	case parser.MethodIncreaseAllowance, parser.MethodIncreaseAllowanceExported:
		returnValue, err := common.GetBigInt(value, parser.ReturnKey, false)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing return value: %w", err)
		}

		allowanceEvent, err := eg.parseIncreaseAndDecreaseAllowance(ctx, tx, tipsetCid, params, returnValue)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing increase allowance: %w", err)
		}
		allowanceEvents = append(allowanceEvents, allowanceEvent)
	case parser.MethodDecreaseAllowance, parser.MethodDecreaseAllowanceExported:
		returnValue, err := common.GetBigInt(value, parser.ReturnKey, false)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing return value: %w", err)
		}
		allowanceEvent, err := eg.parseIncreaseAndDecreaseAllowance(ctx, tx, tipsetCid, params, returnValue)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing decrease allowance: %w", err)
		}
		allowanceEvents = append(allowanceEvents, allowanceEvent)
	case parser.MethodRevokeAllowance, parser.MethodRevokeAllowanceExported:
		allowanceEvent, err := eg.parseRevokeAllowance(ctx, tx, tipsetCid, params)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing revoke allowance: %w", err)
		}
		allowanceEvents = append(allowanceEvents, allowanceEvent)
	case parser.MethodBurn, parser.MethodBurnExported:
		returnValue, err := getReturn[map[string]interface{}](value)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing return value: %w", err)
		}
		tokenEvent, err := eg.parseBurn(ctx, tx, tipsetCid, params, returnValue)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing burn: %w", err)
		}
		tokenEvents = append(tokenEvents, tokenEvent)
	case parser.MethodBurnFrom, parser.MethodBurnFromExported:
		returnValue, err := getReturn[map[string]interface{}](value)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing return value: %w", err)
		}
		tokenEvent, allowanceEvent, err := eg.parseBurnFrom(ctx, tx, tipsetCid, params, returnValue)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing burn from: %w", err)
		}
		tokenEvents = append(tokenEvents, tokenEvent)
		allowanceEvents = append(allowanceEvents, allowanceEvent)
	}
	return tokenEvents, allowanceEvents, nil
}

func (eg *eventGenerator) parseMint(ctx context.Context, tx *types.Transaction, tipsetCid string, params, ret map[string]interface{}) (*types.DataCapTokenEvent, error) {
	to, err := common.GetItem[string](params, KeyTo, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing to: %w", err)
	}

	balance, err := common.GetBigInt(ret, KeyBalance, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing balance: %w", err)
	}
	supply, err := common.GetBigInt(ret, KeySupply, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing supply: %w", err)
	}

	return &types.DataCapTokenEvent{
		ID:           tools.BuildId(tipsetCid, tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType),
		ActorAddress: to,
		Height:       tx.Height,
		TxCid:        tx.TxCid,
		ActionType:   tx.TxType,
		Data:         tx.TxMetadata,
		Balance:      balance,
		Supply:       supply,
		TxTimestamp:  tx.TxTimestamp,
	}, nil
}

func (eg *eventGenerator) parseDestroy(ctx context.Context, tx *types.Transaction, tipsetCid string, params, ret map[string]interface{}) (*types.DataCapTokenEvent, error) {
	owner, err := common.GetItem[string](params, KeyOwner, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing owner: %w", err)
	}

	balance, err := common.GetBigInt(ret, KeyBalance, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing balance: %w", err)
	}

	return &types.DataCapTokenEvent{
		ID:           tools.BuildId(tipsetCid, tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType),
		ActorAddress: owner,
		Height:       tx.Height,
		TxCid:        tx.TxCid,
		ActionType:   tx.TxType,
		Data:         tx.TxMetadata,
		Balance:      balance,
		TxTimestamp:  tx.TxTimestamp,
	}, nil
}

func (eg *eventGenerator) parseTransfer(ctx context.Context, tx *types.Transaction, tipsetCid string, params, ret map[string]interface{}) ([]*types.DataCapTokenEvent, error) {
	to, err := common.GetItem[string](params, KeyTo, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing to: %w", err)
	}

	fromBalance, err := common.GetBigInt(ret, KeyFromBalance, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing balance: %w", err)
	}
	toBalance, err := common.GetBigInt(ret, KeyToBalance, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing balance: %w", err)
	}

	return []*types.DataCapTokenEvent{
			{
				ID:           tools.BuildId(tipsetCid, tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType, "from"),
				ActorAddress: tx.TxFrom,
				Height:       tx.Height,
				TxCid:        tx.TxCid,
				ActionType:   tx.TxType,
				Data:         tx.TxMetadata,
				Balance:      fromBalance,
				TxTimestamp:  tx.TxTimestamp,
			},
			{
				ID:           tools.BuildId(tipsetCid, tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType, "to"),
				ActorAddress: to,
				Height:       tx.Height,
				TxCid:        tx.TxCid,
				ActionType:   tx.TxType,
				Data:         tx.TxMetadata,
				Balance:      toBalance,
				TxTimestamp:  tx.TxTimestamp,
			},
		},
		nil
}

func (eg *eventGenerator) parseTransferFrom(ctx context.Context, tx *types.Transaction, tipsetCid string, params, ret map[string]interface{}) ([]*types.DataCapTokenEvent, *types.DataCapAllowanceEvent, error) {
	from, err := common.GetItem[string](params, KeyFrom, false)
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing from: %w", err)
	}
	to, err := common.GetItem[string](params, KeyTo, false)
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing to: %w", err)
	}
	fromBalance, err := common.GetBigInt(ret, KeyFromBalance, false)
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing fromBalance: %w", err)
	}
	toBalance, err := common.GetBigInt(ret, KeyToBalance, false)
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing toBalance: %w", err)
	}
	allowance, err := common.GetBigInt(ret, KeyAllowance, false)
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing allowance: %w", err)
	}

	return []*types.DataCapTokenEvent{
			{
				ID:           tools.BuildId(tipsetCid, tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType, "from"),
				ActorAddress: from,
				Height:       tx.Height,
				TxCid:        tx.TxCid,
				ActionType:   tx.TxType,
				Data:         tx.TxMetadata,
				Balance:      fromBalance,
				TxTimestamp:  tx.TxTimestamp,
			},
			{
				ID:           tools.BuildId(tipsetCid, tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType, "to"),
				ActorAddress: to,
				Height:       tx.Height,
				TxCid:        tx.TxCid,
				ActionType:   tx.TxType,
				Data:         tx.TxMetadata,
				Balance:      toBalance,
				TxTimestamp:  tx.TxTimestamp,
			},
		}, &types.DataCapAllowanceEvent{
			ID:               tools.BuildId(tipsetCid, tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType),
			OwnerAddress:     from,
			OperatorAddress:  tx.TxFrom,
			Height:           tx.Height,
			TxCid:            tx.TxCid,
			ActionType:       tx.TxType,
			AllowanceBalance: allowance,
			Data:             tx.TxMetadata,
			TxTimestamp:      tx.TxTimestamp,
		}, nil
}

func (eg *eventGenerator) parseBurn(ctx context.Context, tx *types.Transaction, tipsetCid string, params, ret map[string]interface{}) (*types.DataCapTokenEvent, error) {
	balance, err := common.GetBigInt(ret, KeyBalance, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing balance: %w", err)
	}

	return &types.DataCapTokenEvent{
		ID:           tools.BuildId(tipsetCid, tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType),
		ActorAddress: tx.TxFrom,
		Height:       tx.Height,
		TxCid:        tx.TxCid,
		ActionType:   tx.TxType,
		Data:         tx.TxMetadata,
		Balance:      balance,
		TxTimestamp:  tx.TxTimestamp,
	}, nil
}

func (eg *eventGenerator) parseBurnFrom(ctx context.Context, tx *types.Transaction, tipsetCid string, params, ret map[string]interface{}) (*types.DataCapTokenEvent, *types.DataCapAllowanceEvent, error) {

	owner, err := common.GetItem[string](params, KeyOwner, false)
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing owner: %w", err)
	}

	balance, err := common.GetBigInt(ret, KeyBalance, false)
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing balance: %w", err)
	}
	allowance, err := common.GetBigInt(ret, KeyAllowance, false)
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing allowance: %w", err)
	}

	return &types.DataCapTokenEvent{
			ID:           tools.BuildId(tipsetCid, tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType),
			ActorAddress: owner,
			Height:       tx.Height,
			TxCid:        tx.TxCid,
			ActionType:   tx.TxType,
			Data:         tx.TxMetadata,
			Balance:      balance,
			TxTimestamp:  tx.TxTimestamp,
		}, &types.DataCapAllowanceEvent{
			ID:               tools.BuildId(tipsetCid, tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType),
			OwnerAddress:     owner,
			OperatorAddress:  tx.TxFrom,
			Height:           tx.Height,
			TxCid:            tx.TxCid,
			ActionType:       tx.TxType,
			AllowanceBalance: allowance,
			Data:             tx.TxMetadata,
			TxTimestamp:      tx.TxTimestamp,
		}, nil
}

func (eg *eventGenerator) parseIncreaseAndDecreaseAllowance(ctx context.Context, tx *types.Transaction, tipsetCid string, params map[string]interface{}, ret *big.Int) (*types.DataCapAllowanceEvent, error) {
	operator, err := common.GetItem[string](params, KeyOperator, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing operator: %w", err)
	}

	return &types.DataCapAllowanceEvent{
		ID:               tools.BuildId(tipsetCid, tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType),
		OwnerAddress:     tx.TxFrom,
		OperatorAddress:  operator,
		Height:           tx.Height,
		TxCid:            tx.TxCid,
		ActionType:       tx.TxType,
		AllowanceBalance: ret,
		Data:             tx.TxMetadata,
		TxTimestamp:      tx.TxTimestamp,
	}, nil
}
func (eg *eventGenerator) parseRevokeAllowance(ctx context.Context, tx *types.Transaction, tipsetCid string, params map[string]interface{}) (*types.DataCapAllowanceEvent, error) {
	operator, err := common.GetItem[string](params, KeyOperator, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing operator: %w", err)
	}

	return &types.DataCapAllowanceEvent{
		ID:               tools.BuildId(tipsetCid, tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType),
		OwnerAddress:     tx.TxFrom,
		OperatorAddress:  operator,
		Height:           tx.Height,
		TxCid:            tx.TxCid,
		ActionType:       tx.TxType,
		AllowanceBalance: big.NewInt(0),
		Data:             tx.TxMetadata,
		TxTimestamp:      tx.TxTimestamp,
	}, nil
}

func isDatacapTokenMessage(txType string) bool {
	switch {
	case strings.Contains(txType, parser.MethodMint), strings.Contains(txType, parser.MethodMintExported):
		return true
	case strings.Contains(txType, parser.MethodDestroy), strings.Contains(txType, parser.MethodDestroyExported):
		return true
	case strings.Contains(txType, parser.MethodTransfer), strings.Contains(txType, parser.MethodTransferExported):
		return true
	case strings.Contains(txType, parser.MethodTransferFrom), strings.Contains(txType, parser.MethodTransferFromExported):
		return true
	case strings.Contains(txType, parser.MethodIncreaseAllowance), strings.Contains(txType, parser.MethodIncreaseAllowanceExported):
		return true
	case strings.Contains(txType, parser.MethodDecreaseAllowance), strings.Contains(txType, parser.MethodDecreaseAllowanceExported):
		return true
	case strings.Contains(txType, parser.MethodRevokeAllowance), strings.Contains(txType, parser.MethodRevokeAllowanceExported):
		return true
	case strings.Contains(txType, parser.MethodBurn), strings.Contains(txType, parser.MethodBurnExported):
		return true
	case strings.Contains(txType, parser.MethodBurnFrom), strings.Contains(txType, parser.MethodBurnFromExported):
		return true
	}

	return false
}

func getReturn[T any](value map[string]interface{}) (T, error) {
	ret, err := common.GetItem[T](value, parser.ReturnKey, false)
	if err != nil {
		return ret, fmt.Errorf("error parsing return value: %w", err)
	}
	return ret, nil
}
