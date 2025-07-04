package miner

import (
	"fmt"
	"math/big"

	"github.com/filecoin-project/go-state-types/abi"
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
