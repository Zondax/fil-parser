package miner

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
)

const (
	KeySectorNumber  = "SectorNumber"
	KeySectorNumbers = "SectorNumbers"
	KeySectors       = "Sectors"
	KeyExpiration    = "Expiration"
	KeySectorSize    = "SectorSize"
	KeyNewExpiration = "NewExpiration"
	KeyParams        = "Params"
	KeyTerminations  = "Terminations"
	KeyFaults        = "Faults"
	KeyRecoveries    = "Recoveries"
	KeyExtensions    = "Extensions"
	KeySealProof     = "SealProof"
)

func (eg *eventGenerator) isMinerSectorMessage(actorName, txType string) bool {
	if actorName != manifest.MinerKey {
		return false
	}
	switch txType {
	case
		// pre-commit stage
		parser.MethodPreCommitSector,
		parser.MethodPreCommitSectorBatch,
		parser.MethodPreCommitSectorBatch2,

		// prove-commit stage
		parser.MethodProveCommitSector,
		parser.MethodProveCommitSectors3,
		parser.MethodProveCommitSectorsNI,
		parser.MethodConfirmSectorProofsValid,
		parser.MethodProveCommitAggregate,

		// termination and recovery stage
		parser.MethodTerminateSectors,
		parser.MethodDeclareFaults,
		parser.MethodDeclareFaultsRecovered,

		// expiry extension stage
		parser.MethodExtendSectorExpiration,
		parser.MethodExtendSectorExpiration2:

		return true
	}
	return false
}

func (eg *eventGenerator) createSectorEvents(ctx context.Context, tx *types.Transaction, tipsetCid string) ([]*types.MinerSectorEvent, error) {
	var value map[string]interface{}
	json.Unmarshal([]byte(tx.TxMetadata), &value)

	params, err := getItem[map[string]interface{}](value, KeyParams, false)
	if err != nil {
		return nil, err
	}

	switch tx.TxType {
	case parser.MethodPreCommitSector, parser.MethodPreCommitSectorBatch, parser.MethodPreCommitSectorBatch2:
		sectorEvents, err := eg.parsePreCommitStage(ctx, tx, tipsetCid, params)
		if err != nil {
			return nil, err
		}
		return sectorEvents, nil

	case parser.MethodConfirmSectorProofsValid, parser.MethodProveCommitAggregate, parser.MethodProveCommitSector, parser.MethodProveCommitSectors3, parser.MethodProveCommitSectorsNI:
		sectorEvents, err := eg.parseProveCommitStage(ctx, tx, tipsetCid, params)
		if err != nil {
			return nil, err
		}
		return sectorEvents, nil

	case parser.MethodTerminateSectors, parser.MethodDeclareFaults, parser.MethodDeclareFaultsRecovered:
		sectorEvents, err := eg.parseSectorTerminationFaultAndRecoveries(ctx, tx, tipsetCid, params)
		if err != nil {
			return nil, err
		}
		return sectorEvents, nil

	case parser.MethodExtendSectorExpiration, parser.MethodExtendSectorExpiration2:
		sectorEvents, err := eg.parseSectorExpiryExtensions(ctx, tx, tipsetCid, params)
		if err != nil {
			return nil, err
		}
		return sectorEvents, nil

	}

	return nil, nil
}

func (eg *eventGenerator) parseProveCommitStage(_ context.Context, tx *types.Transaction, tipsetCid string, params map[string]interface{}) ([]*types.MinerSectorEvent, error) {
	var sectorEvents []*types.MinerSectorEvent
	if tx.TxType == parser.MethodProveCommitAggregate {
		sectorBitField, err := getIntegerSlice[int](params, "Sectors", false)
		if err != nil {
			return nil, err
		}
		sectorNumbers := jsonEncodedBitfieldToSectorNumbers(sectorBitField)
		jsonData, err := json.Marshal(map[string]interface{}{
			KeySectorNumbers: sectorNumbers,
			KeySectors:       sectorBitField,
		})
		if err != nil {
			return nil, err
		}
		for _, sectorNumber := range sectorNumbers {
			sectorEvents = append(sectorEvents, &types.MinerSectorEvent{
				ID:           tools.BuildId(tipsetCid, tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType),
				MinerAddress: tx.TxTo,
				SectorNumber: uint64(sectorNumber),
				Height:       tx.Height,
				TxCid:        tx.TxCid,
				ActionType:   tx.TxType,
				Data:         string(jsonData),
			})
		}
		return sectorEvents, nil
	}

	if tx.TxType == parser.MethodConfirmSectorProofsValid {
		sectors, err := getIntegerSlice[int64](params, KeySectors, false)
		if err != nil {
			return nil, err
		}
		jsonData, err := json.Marshal(map[string]interface{}{
			KeySectorNumbers: sectors,
		})
		if err != nil {
			return nil, err
		}
		for _, sector := range sectors {
			sectorEvents = append(sectorEvents, &types.MinerSectorEvent{
				ID:           tools.BuildId(tipsetCid, tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType),
				MinerAddress: tx.TxTo,
				SectorNumber: uint64(sector),
				Height:       tx.Height,
				TxCid:        tx.TxCid,
				ActionType:   tx.TxType,
				Data:         string(jsonData),
			})
		}
	}

	sectorNumber, err := getInteger[int64](params, KeySectorNumber, false)
	if err != nil {
		return nil, err
	}
	jsonData, err := json.Marshal(map[string]interface{}{
		KeySectorNumber: sectorNumber,
	})
	if err != nil {
		return nil, err
	}
	sectorEvents = append(sectorEvents, &types.MinerSectorEvent{
		ID:           tools.BuildId(tipsetCid, tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType),
		MinerAddress: tx.TxTo,
		SectorNumber: uint64(sectorNumber),
		Height:       tx.Height,
		TxCid:        tx.TxCid,
		ActionType:   tx.TxType,
		Data:         string(jsonData),
	})

	return sectorEvents, nil
}

func (eg *eventGenerator) parsePreCommitStage(_ context.Context, tx *types.Transaction, tipsetCid string, params map[string]interface{}) ([]*types.MinerSectorEvent, error) {
	var sectorEvents []*types.MinerSectorEvent

	addEvent := func(params map[string]interface{}) error {
		sealProof, err := getInteger[int64](params, KeySealProof, false)
		if err != nil {
			return err
		}

		sectorNumber, err := getInteger[uint64](params, KeySectorNumber, false)
		if err != nil {
			return err
		}

		expiration, err := getInteger[int64](params, KeyExpiration, false)
		if err != nil {
			return err
		}
		jsonData, err := json.Marshal(map[string]interface{}{
			KeySectorNumber: sectorNumber,
			KeyExpiration:   expiration,
			KeySectorSize:   sectorProofToBigInt(sealProof).Uint64(),
		})
		if err != nil {
			return err
		}
		sectorEvents = append(sectorEvents, &types.MinerSectorEvent{
			ID:           tools.BuildId(tipsetCid, tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType),
			TxCid:        tx.TxCid,
			Height:       tx.Height,
			ActionType:   tx.TxType,
			MinerAddress: tx.TxTo,
			SectorNumber: sectorNumber,
			Data:         string(jsonData),
		})
		return nil
	}

	if tx.TxType == parser.MethodPreCommitSector {
		err := addEvent(params)
		if err != nil {
			return nil, err
		}
	} else {
		sectors, err := getSlice[map[string]interface{}](params, KeySectors, false)
		if err != nil {
			return nil, err
		}
		for _, sector := range sectors {
			err := addEvent(sector)
			if err != nil {
				return nil, err
			}
		}
	}
	return sectorEvents, nil
}

func (eg *eventGenerator) parseSectorTerminationFaultAndRecoveries(_ context.Context, tx *types.Transaction, tipsetCid string, params map[string]interface{}) ([]*types.MinerSectorEvent, error) {
	var sectorEvents []*types.MinerSectorEvent
	var parameterName string

	switch tx.TxType {
	case parser.MethodTerminateSectors:
		parameterName = KeyTerminations
	case parser.MethodDeclareFaults:
		parameterName = KeyFaults
	case parser.MethodDeclareFaultsRecovered:
		parameterName = KeyRecoveries
	}
	events, err := getSlice[map[string]interface{}](params, parameterName, false)
	if err != nil {
		return nil, err
	}
	for _, event := range events {
		sectorBitField, err := getIntegerSlice[int](event, KeySectors, false)
		if err != nil {
			return nil, err
		}

		sectorNumbers := jsonEncodedBitfieldToSectorNumbers(sectorBitField)
		event[KeySectorNumbers] = sectorNumbers
		jsonData, err := json.Marshal(event)
		if err != nil {
			return nil, err
		}
		for _, sectorNumber := range sectorNumbers {
			sectorEvents = append(sectorEvents, &types.MinerSectorEvent{
				ID:           tools.BuildId(tipsetCid, tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType),
				MinerAddress: tx.TxTo,
				SectorNumber: uint64(sectorNumber),
				Height:       tx.Height,
				TxCid:        tx.TxCid,
				ActionType:   tx.TxType,
				Data:         string(jsonData),
			})
		}
	}
	return sectorEvents, nil
}

func (eg *eventGenerator) parseSectorExpiryExtensions(_ context.Context, tx *types.Transaction, tipsetCid string, params map[string]interface{}) ([]*types.MinerSectorEvent, error) {
	var sectorEvents []*types.MinerSectorEvent
	extensions, err := getSlice[map[string]interface{}](params, KeyExtensions, false)
	if err != nil {
		return nil, err
	}
	for _, extension := range extensions {
		sectorBitField, err := getIntegerSlice[int](extension, KeySectors, false)
		if err != nil {
			return nil, err
		}
		newExpiration, err := getInteger[int64](extension, KeyNewExpiration, false)
		if err != nil {
			return nil, err
		}

		sectorNumbers := jsonEncodedBitfieldToSectorNumbers(sectorBitField)
		jsonData, err := json.Marshal(map[string]interface{}{
			KeySectors:       sectorBitField,
			KeySectorNumbers: sectorNumbers,
			KeyNewExpiration: newExpiration,
		})
		if err != nil {
			return nil, err
		}
		for _, sectorNumber := range sectorNumbers {
			sectorEvents = append(sectorEvents, &types.MinerSectorEvent{
				ID:           tools.BuildId(tipsetCid, tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType),
				MinerAddress: tx.TxTo,
				SectorNumber: uint64(sectorNumber),
				Height:       tx.Height,
				TxCid:        tx.TxCid,
				ActionType:   tx.TxType,
				Data:         string(jsonData),
			})
		}
	}
	return sectorEvents, nil
}

// the sector bit field is a range of bits representing the different sector numbers.
// example: sectors: [ 0 1 2 3] -> bitfield: [1 1 1 1 ] -> JSON: [0,4]
// example: sectors: [0 1 3 4 5] -> bitfield: [1 1 0 1 1 1] -> JSON: [ 0,2,1,3 ]
// the JSON format always starts with a 0 and proceeds with the 0/1 pattern
// see here: https://pkg.go.dev/github.com/filecoin-project/go-bitfield@v0.2.4/rle#RLE.MarshalJSON
func jsonEncodedBitfieldToSectorNumbers(bitField []int) []uint64 {
	var sectorNumbers []uint64

	var sectorNumber int
	for i, num := range bitField {
		for j := 0; j < num; j++ {
			sectorNumber++
			if i%2 != 0 {
				// sector selected
				sectorNumbers = append(sectorNumbers, uint64(sectorNumber))
			}
			// sector is not selected for the operation
		}
	}
	return sectorNumbers
}
