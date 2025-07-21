package miner

import (
	"encoding/json"
	"fmt"

	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/tools/common"
	"github.com/zondax/fil-parser/types"
)

const (
	KeyMiner           = "Miner"
	KeyOwnerAddr       = "OwnerAddr"
	KeyWorkerAddr      = "WorkerAddr"
	KeyNewWorker       = "NewWorker"
	KeyControlAddrs    = "ControlAddrs"
	KeyNewControlAddrs = "NewControlAddrs"
	KeyMultiaddrs      = "Multiaddrs"
	KeyNewMultiAddrs   = "NewMultiaddrs"
	KeyNewBeneficiary  = "NewBeneficiary"
)

func (eg *eventGenerator) createMinerInfo(tx *types.Transaction, tipsetCid, actorAddress string) (*types.MinerInfo, error) {
	// for these tx types we need to consolidate the addresses in the parameters
	switch tx.TxType {
	case parser.MethodAwardBlockReward:
		return eg.parseAwardBlockReward(tx, tipsetCid)
	case parser.MethodConstructor:
		return eg.parseConstructor(tx, tipsetCid, actorAddress)
	case parser.MethodChangeWorkerAddress:
		return eg.parseChangeWorkerAddress(tx, tipsetCid, actorAddress)
	case parser.MethodChangeMultiaddrs:
		return eg.parseChangeMultiaddrs(tx, tipsetCid, actorAddress)
	case parser.MethodChangeBeneficiary:
		return eg.parseChangeBeneficiary(tx, tipsetCid, actorAddress)
	case parser.MethodChangeOwnerAddress:
		return eg.parseChangeOwnerAddress(tx, tipsetCid, actorAddress)
	}

	minerInfo := &types.MinerInfo{
		ID:           tools.BuildId(tipsetCid, tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType),
		MinerAddress: actorAddress,
		Height:       tx.Height,
		TxCid:        tx.TxCid,
		ActionType:   tx.TxType,
		Data:         tx.TxMetadata,
		TxTimestamp:  tx.TxTimestamp,
	}

	return minerInfo, nil
}

func (eg *eventGenerator) parseAwardBlockReward(tx *types.Transaction, tipsetCid string) (*types.MinerInfo, error) {
	var value map[string]interface{}
	err := json.Unmarshal([]byte(tx.TxMetadata), &value)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling tx metadata: %w", err)
	}

	params, err := common.GetItem[map[string]interface{}](value, KeyParams, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing params: %w", err)
	}
	minerAddress, err := common.GetItem[string](params, KeyMiner, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing miner address: %w", err)
	}
	if eg.config.ConsolidateRobustAddress {
		if minerAddress != "" {
			parsedMinerAddress, err := eg.consolidateAddress(minerAddress)
			if err != nil {
				eg.logger.Errorf("error consolidating miner address: %w", err)
			} else {
				minerAddress = parsedMinerAddress
				params[KeyMiner] = minerAddress
			}
		}
		value[KeyParams] = params
	}
	return &types.MinerInfo{
		ID:           tools.BuildId(tipsetCid, tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType),
		MinerAddress: minerAddress,
		Height:       tx.Height,
		TxCid:        tx.TxCid,
		ActionType:   tx.TxType,
		Data:         tx.TxMetadata,
		TxTimestamp:  tx.TxTimestamp,
	}, nil
}

func (eg *eventGenerator) parseConstructor(tx *types.Transaction, tipsetCid, actorAddress string) (*types.MinerInfo, error) {
	var value map[string]interface{}
	err := json.Unmarshal([]byte(tx.TxMetadata), &value)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling tx metadata: %w", err)
	}

	params, err := common.GetItem[map[string]interface{}](value, KeyParams, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing params: %w", err)
	}

	if eg.config.ConsolidateRobustAddress {
		if err := eg.consolidateConstructorAddresses(params); err != nil {
			eg.logger.Errorf("error consolidating constructor addresses: %w", err)
		} else {
			value[KeyParams] = params
		}
	}

	jsonData, err := json.Marshal(value)
	if err != nil {
		return nil, fmt.Errorf("error marshalling tx metadata: %w", err)
	}
	return &types.MinerInfo{
		ID:           tools.BuildId(tipsetCid, tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType),
		MinerAddress: actorAddress,
		Height:       tx.Height,
		TxCid:        tx.TxCid,
		ActionType:   tx.TxType,
		Data:         string(jsonData),
		TxTimestamp:  tx.TxTimestamp,
	}, nil
}

func (eg *eventGenerator) parseChangeWorkerAddress(tx *types.Transaction, tipsetCid, actorAddress string) (*types.MinerInfo, error) {
	var value map[string]interface{}
	err := json.Unmarshal([]byte(tx.TxMetadata), &value)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling tx metadata: %w", err)
	}

	params, err := common.GetItem[map[string]interface{}](value, KeyParams, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing params: %w", err)
	}

	if eg.config.ConsolidateRobustAddress {
		workerAddress, _ := common.GetItem[string](params, KeyNewWorker, true)
		if workerAddress != "" {
			parsedWorkerAddress, err := eg.consolidateAddress(workerAddress)
			if err != nil {
				eg.logger.Errorf("error consolidating worker address: %w", err)
			} else {
				params[KeyNewWorker] = parsedWorkerAddress
			}
		}
		controlAddresses, _ := common.GetSlice[string](params, KeyNewControlAddrs, true)
		if len(controlAddresses) > 0 {
			consolidatedControlAddresses := make([]string, 0, len(controlAddresses))
			for _, addrStr := range controlAddresses {
				parsedControlAddress, err := eg.consolidateAddress(addrStr)
				if err != nil {
					eg.logger.Errorf("error consolidating control address: %w", err)
				} else {
					consolidatedControlAddresses = append(consolidatedControlAddresses, parsedControlAddress)
				}
			}
			if len(consolidatedControlAddresses) == len(controlAddresses) {
				params[KeyNewControlAddrs] = consolidatedControlAddresses
			}
		}
		value[KeyParams] = params
	}

	jsonData, err := json.Marshal(value)
	if err != nil {
		return nil, fmt.Errorf("error marshalling tx metadata: %w", err)
	}
	return &types.MinerInfo{
		ID:           tools.BuildId(tipsetCid, tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType),
		MinerAddress: actorAddress,
		Height:       tx.Height,
		TxCid:        tx.TxCid,
		ActionType:   tx.TxType,
		Data:         string(jsonData),
		TxTimestamp:  tx.TxTimestamp,
	}, nil
}

func (eg *eventGenerator) parseChangeMultiaddrs(tx *types.Transaction, tipsetCid, actorAddress string) (*types.MinerInfo, error) {
	var value map[string]interface{}
	err := json.Unmarshal([]byte(tx.TxMetadata), &value)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling tx metadata: %w", err)
	}

	params, err := common.GetItem[map[string]interface{}](value, KeyParams, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing params: %w", err)
	}

	if eg.config.ConsolidateRobustAddress {
		multiaddrs, _ := common.GetSlice[string](params, KeyNewMultiAddrs, true)
		if len(multiaddrs) > 0 {
			consolidatedMultiAddrs := make([]string, 0, len(multiaddrs))
			for _, addrStr := range multiaddrs {
				parsedMultiaddr, err := eg.consolidateAddress(addrStr)
				if err != nil {
					eg.logger.Errorf("error consolidating multiaddr: %w", err)
				} else {
					consolidatedMultiAddrs = append(consolidatedMultiAddrs, parsedMultiaddr)
				}
			}
			if len(consolidatedMultiAddrs) == len(multiaddrs) {
				params[KeyNewMultiAddrs] = consolidatedMultiAddrs
			}
		}
		value[KeyParams] = params
	}

	jsonData, err := json.Marshal(value)
	if err != nil {
		return nil, fmt.Errorf("error marshalling tx metadata: %w", err)
	}
	return &types.MinerInfo{
		ID:           tools.BuildId(tipsetCid, tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType),
		MinerAddress: actorAddress,
		Height:       tx.Height,
		TxCid:        tx.TxCid,
		ActionType:   tx.TxType,
		Data:         string(jsonData),
		TxTimestamp:  tx.TxTimestamp,
	}, nil
}

func (eg *eventGenerator) parseChangeBeneficiary(tx *types.Transaction, tipsetCid, actorAddress string) (*types.MinerInfo, error) {
	var value map[string]interface{}
	err := json.Unmarshal([]byte(tx.TxMetadata), &value)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling tx metadata: %w", err)
	}

	params, err := common.GetItem[map[string]interface{}](value, KeyParams, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing params: %w", err)
	}

	if eg.config.ConsolidateRobustAddress {
		beneficiary, _ := common.GetItem[string](params, KeyNewBeneficiary, true)
		if beneficiary != "" {
			parsedBeneficiary, err := eg.consolidateAddress(beneficiary)
			if err != nil {
				eg.logger.Errorf("error consolidating beneficiary address: %w", err)
			} else {
				params[KeyNewBeneficiary] = parsedBeneficiary
			}
		}
		value[KeyParams] = params
	}

	jsonData, err := json.Marshal(value)
	if err != nil {
		return nil, fmt.Errorf("error marshalling tx metadata: %w", err)
	}
	return &types.MinerInfo{
		ID:           tools.BuildId(tipsetCid, tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType),
		MinerAddress: actorAddress,
		Height:       tx.Height,
		TxCid:        tx.TxCid,
		ActionType:   tx.TxType,
		Data:         string(jsonData),
		TxTimestamp:  tx.TxTimestamp,
	}, nil
}

func (eg *eventGenerator) parseChangeOwnerAddress(tx *types.Transaction, tipsetCid, actorAddress string) (*types.MinerInfo, error) {
	var value map[string]interface{}
	err := json.Unmarshal([]byte(tx.TxMetadata), &value)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling tx metadata: %w", err)
	}

	if eg.config.ConsolidateRobustAddress {
		ownerAddress, _ := common.GetItem[string](value, KeyParams, true)
		if ownerAddress != "" {
			parsedOwnerAddress, err := eg.consolidateAddress(ownerAddress)
			if err != nil {
				eg.logger.Errorf("error consolidating owner address: %w", err)
			} else {
				value[KeyParams] = parsedOwnerAddress
			}
		}
	}

	jsonData, err := json.Marshal(value)
	if err != nil {
		return nil, fmt.Errorf("error marshalling tx metadata: %w", err)
	}
	return &types.MinerInfo{
		ID:           tools.BuildId(tipsetCid, tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType),
		MinerAddress: actorAddress,
		Height:       tx.Height,
		TxCid:        tx.TxCid,
		ActionType:   tx.TxType,
		Data:         string(jsonData),
		TxTimestamp:  tx.TxTimestamp,
	}, nil
}

func (eg *eventGenerator) consolidateConstructorAddresses(params map[string]interface{}) error {
	ownerAddress, _ := common.GetItem[string](params, KeyOwnerAddr, true)
	if ownerAddress != "" {
		parsedOwnerAddress, err := eg.consolidateAddress(ownerAddress)
		if err != nil {
			eg.logger.Errorf("error consolidating owner address: %w", err)
		} else {
			params[KeyOwnerAddr] = parsedOwnerAddress
		}
	}

	workerAddress, _ := common.GetItem[string](params, KeyWorkerAddr, true)
	if workerAddress != "" {
		parsedWorkerAddress, err := eg.consolidateAddress(workerAddress)
		if err != nil {
			eg.logger.Errorf("error consolidating worker address: %w", err)
		} else {
			params[KeyWorkerAddr] = parsedWorkerAddress
		}
	}
	controlAddresses, _ := common.GetSlice[string](params, KeyControlAddrs, true)
	if len(controlAddresses) > 0 {
		consolidatedControlAddresses := make([]string, 0, len(controlAddresses))
		for _, addrStr := range controlAddresses {
			parsedControlAddress, err := eg.consolidateAddress(addrStr)
			if err != nil {
				eg.logger.Errorf("error consolidating control address: %w", err)
			} else {
				consolidatedControlAddresses = append(consolidatedControlAddresses, parsedControlAddress)
			}
		}
		if len(consolidatedControlAddresses) == len(controlAddresses) {
			params[KeyControlAddrs] = consolidatedControlAddresses
		}
	}
	multiaddrs, _ := common.GetSlice[string](params, KeyMultiaddrs, true)
	if len(multiaddrs) > 0 {
		consolidatedMultiAddrs := make([]string, 0, len(multiaddrs))
		for _, addrStr := range multiaddrs {
			parsedMultiaddr, err := eg.consolidateAddress(addrStr)
			if err != nil {
				eg.logger.Errorf("error consolidating multiaddr: %w", err)
			} else {
				consolidatedMultiAddrs = append(consolidatedMultiAddrs, parsedMultiaddr)
			}
		}
		if len(consolidatedMultiAddrs) == len(multiaddrs) {
			params[KeyMultiaddrs] = consolidatedMultiAddrs
		}
	}

	return nil
}
