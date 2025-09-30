package miner

import (
	"fmt"
	"math/big"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/zondax/fil-parser/actors"
)

func sectorProofToBigInt(sectorProof int64) *big.Int {
	proofType := abi.RegisteredSealProof(sectorProof)
	info, ok := abi.SealProofInfos[proofType]
	if !ok {
		fmt.Println("invalid proofType", proofType)
		return big.NewInt(0)
	}

	return big.NewInt(0).SetUint64(uint64(info.SectorSize))
}

func (eg *eventGenerator) consolidateIDAddress(idAddress uint64) (string, error) {
	addr, err := address.NewIDAddress(idAddress)
	if err != nil {
		return "", fmt.Errorf("error parsing id address: %w", err)
	}
	consolidatedIDAddress, err := actors.ConsolidateToRobustAddress(addr, eg.helper, eg.logger, eg.config.RobustAddressBestEffort, true)
	if err != nil {
		return "", fmt.Errorf("error consolidating id address: %w", err)
	}
	return consolidatedIDAddress, nil
}

func (eg *eventGenerator) consolidateAddress(addrStr string) (string, error) {
	addr, err := address.NewFromString(addrStr)
	if err != nil {
		return "", fmt.Errorf("error parsing address: %w", err)
	}
	consolidatedAddress, err := actors.ConsolidateToRobustAddress(addr, eg.helper, eg.logger, eg.config.RobustAddressBestEffort, true)
	if err != nil {
		return "", fmt.Errorf("error consolidating address: %w", err)
	}
	return consolidatedAddress, nil
}
