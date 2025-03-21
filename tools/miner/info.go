package miner

import (
	"encoding/json"
	"fmt"

	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
)

const (
	KeyMiner = "miner"
)

func (eg *eventGenerator) createMinerInfo(tx *types.Transaction, tipsetCid, actorAddress string) (*types.MinerInfo, error) {
	if tx.TxType == parser.MethodAwardBlockReward {
		return eg.parseAwardBlockReward(tx, tipsetCid)
	}
	minerInfo := &types.MinerInfo{
		ID:           tools.BuildId(tipsetCid, tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType),
		ActorAddress: actorAddress,
		Height:       tx.Height,
		TxCid:        tx.TxCid,
		ActionType:   tx.TxType,
		Value:        tx.TxMetadata,
	}

	return minerInfo, nil
}

func (eg *eventGenerator) parseAwardBlockReward(tx *types.Transaction, tipsetCid string) (*types.MinerInfo, error) {
	var value map[string]interface{}
	err := json.Unmarshal([]byte(tx.TxMetadata), &value)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling tx metadata: %w", err)
	}

	params, err := getItem[map[string]interface{}](value, KeyParams, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing params: %w", err)
	}
	minerAddress, err := getItem[string](params, KeyMiner, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing miner address: %w", err)
	}
	return &types.MinerInfo{
		ID:           tools.BuildId(tipsetCid, tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType),
		ActorAddress: minerAddress,
		Height:       tx.Height,
		TxCid:        tx.TxCid,
		ActionType:   tx.TxType,
		Value:        tx.TxMetadata,
	}, nil
}
