package miner

import (
	"encoding/json"
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
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

	params, err := getItem[map[string]interface{}](value, KeyParams, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing params: %w", err)
	}
	minerAddress, err := getItem[string](params, KeyMiner, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing miner address: %w", err)
	}
	if eg.config.ConsolidateRobustAddress {
		addr, err := address.NewFromString(minerAddress)
		if err != nil {
			return nil, fmt.Errorf("error parsing miner address: %w", err)
		}
		minerAddress, err = actors.ConsolidateToRobustAddress(addr, eg.helper, eg.logger, eg.config.RobustAddressBestEffort)
		if err != nil {
			return nil, fmt.Errorf("error consolidating miner address: %w", err)
		}
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

	params, err := getItem[map[string]interface{}](value, KeyParams, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing params: %w", err)
	}
	ownerAddress, err := getItem[string](params, KeyOwnerAddr, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing miner address: %w", err)
	}
	workerAddress, err := getItem[string](params, KeyWorkerAddr, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing miner address: %w", err)
	}
	controlAddresses, err := getSlice[string](params, KeyControlAddrs, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing miner address: %w", err)
	}
	multiaddrs, err := getSlice[string](params, KeyMultiaddrs, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing miner address: %w", err)
	}

	if eg.config.ConsolidateRobustAddress {
		addr, err := address.NewFromString(ownerAddress)
		if err != nil {
			return nil, fmt.Errorf("error parsing miner address: %w", err)
		}
		ownerAddress, err = actors.ConsolidateToRobustAddress(addr, eg.helper, eg.logger, eg.config.RobustAddressBestEffort)
		if err != nil {
			return nil, fmt.Errorf("error consolidating miner address: %w", err)
		}
		addr, err = address.NewFromString(workerAddress)
		if err != nil {
			return nil, fmt.Errorf("error parsing miner address: %w", err)
		}
		workerAddress, err = actors.ConsolidateToRobustAddress(addr, eg.helper, eg.logger, eg.config.RobustAddressBestEffort)
		if err != nil {
			return nil, fmt.Errorf("error consolidating miner address: %w", err)
		}
		consolidatedControlAddrs := make([]string, 0, len(controlAddresses))
		consolidatedMultiAddrs := make([]string, 0, len(multiaddrs))
		for _, addrStr := range controlAddresses {
			addr, err = address.NewFromString(addrStr)
			if err != nil {
				return nil, fmt.Errorf("error parsing miner address: %w", err)
			}
			controlAddress, err := actors.ConsolidateToRobustAddress(addr, eg.helper, eg.logger, eg.config.RobustAddressBestEffort)
			if err != nil {
				return nil, fmt.Errorf("error consolidating miner address: %w", err)
			}
			consolidatedControlAddrs = append(consolidatedControlAddrs, controlAddress)
		}
		for _, addrStr := range multiaddrs {
			addr, err = address.NewFromString(addrStr)
			if err != nil {
				return nil, fmt.Errorf("error parsing miner address: %w", err)
			}
			multiAddr, err := actors.ConsolidateToRobustAddress(addr, eg.helper, eg.logger, eg.config.RobustAddressBestEffort)
			if err != nil {
				return nil, fmt.Errorf("error consolidating miner address: %w", err)
			}
			consolidatedMultiAddrs = append(consolidatedMultiAddrs, multiAddr)
		}
		controlAddresses = consolidatedControlAddrs
		multiaddrs = consolidatedMultiAddrs
	}
	value[KeyOwnerAddr] = ownerAddress
	value[KeyWorkerAddr] = workerAddress
	value[KeyControlAddrs] = controlAddresses
	value[KeyMultiaddrs] = multiaddrs

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

	params, err := getItem[map[string]interface{}](value, KeyParams, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing params: %w", err)
	}

	workerAddress, err := getItem[string](params, KeyNewWorker, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing miner address: %w", err)
	}
	controlAddresses, err := getSlice[string](params, KeyNewControlAddrs, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing miner address: %w", err)
	}

	if eg.config.ConsolidateRobustAddress {
		addr, err := address.NewFromString(workerAddress)
		if err != nil {
			return nil, fmt.Errorf("error parsing miner address: %w", err)
		}
		workerAddress, err = actors.ConsolidateToRobustAddress(addr, eg.helper, eg.logger, eg.config.RobustAddressBestEffort)
		if err != nil {
			return nil, fmt.Errorf("error consolidating miner address: %w", err)
		}
		consolidatedControlAddrs := make([]string, 0, len(controlAddresses))
		for _, addrStr := range controlAddresses {
			addr, err = address.NewFromString(addrStr)
			if err != nil {
				return nil, fmt.Errorf("error parsing miner address: %w", err)
			}
			controlAddress, err := actors.ConsolidateToRobustAddress(addr, eg.helper, eg.logger, eg.config.RobustAddressBestEffort)
			if err != nil {
				return nil, fmt.Errorf("error consolidating miner address: %w", err)
			}
			consolidatedControlAddrs = append(consolidatedControlAddrs, controlAddress)
		}

		controlAddresses = consolidatedControlAddrs
	}
	value[KeyNewWorker] = workerAddress
	value[KeyNewControlAddrs] = controlAddresses

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

	params, err := getItem[map[string]interface{}](value, KeyParams, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing params: %w", err)
	}

	multiaddrs, err := getSlice[string](params, KeyNewMultiAddrs, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing miner address: %w", err)
	}

	if eg.config.ConsolidateRobustAddress {
		consolidatedMultiAddrs := make([]string, 0, len(multiaddrs))
		for _, addrStr := range multiaddrs {
			addr, err := address.NewFromString(addrStr)
			if err != nil {
				return nil, fmt.Errorf("error parsing miner address: %w", err)
			}
			multiAddr, err := actors.ConsolidateToRobustAddress(addr, eg.helper, eg.logger, eg.config.RobustAddressBestEffort)
			if err != nil {
				return nil, fmt.Errorf("error consolidating miner address: %w", err)
			}
			consolidatedMultiAddrs = append(consolidatedMultiAddrs, multiAddr)
		}

		multiaddrs = consolidatedMultiAddrs
	}
	value[KeyNewMultiAddrs] = multiaddrs

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

	params, err := getItem[map[string]interface{}](value, KeyParams, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing params: %w", err)
	}

	beneficiary, err := getItem[string](params, KeyNewBeneficiary, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing miner address: %w", err)
	}

	if eg.config.ConsolidateRobustAddress {
		addr, err := address.NewFromString(beneficiary)
		if err != nil {
			return nil, fmt.Errorf("error parsing miner address: %w", err)
		}
		beneficiary, err = actors.ConsolidateToRobustAddress(addr, eg.helper, eg.logger, eg.config.RobustAddressBestEffort)
		if err != nil {
			return nil, fmt.Errorf("error consolidating miner address: %w", err)
		}
	}
	value[KeyNewBeneficiary] = beneficiary

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

	ownerAddress, err := getItem[string](value, KeyParams, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing params: %w", err)
	}

	if eg.config.ConsolidateRobustAddress {
		addr, err := address.NewFromString(ownerAddress)
		if err != nil {
			return nil, fmt.Errorf("error parsing miner address: %w", err)
		}
		ownerAddress, err = actors.ConsolidateToRobustAddress(addr, eg.helper, eg.logger, eg.config.RobustAddressBestEffort)
		if err != nil {
			return nil, fmt.Errorf("error consolidating miner address: %w", err)
		}
	}
	value[KeyParams] = ownerAddress

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
