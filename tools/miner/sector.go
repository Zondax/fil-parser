package miner

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
)

const (
	KeySectorNumber      = "SectorNumber"
	KeySectorNumbers     = "SectorNumbers"
	KeySectors           = "Sectors"
	KeyExpiration        = "Expiration"
	KeySectorSize        = "SectorSize"
	KeyDealIDs           = "DealIDs"
	KeyNewExpiration     = "NewExpiration"
	KeyParams            = "Params"
	KeyTerminations      = "Terminations"
	KeyFaults            = "Faults"
	KeyRecoveries        = "Recoveries"
	KeyExtensions        = "Extensions"
	KeySealProof         = "SealProof"
	KeySectorActivations = "SectorActivations"
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
	err := json.Unmarshal([]byte(tx.TxMetadata), &value)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling tx metadata: %w", err)
	}

	params, err := getItem[map[string]interface{}](value, KeyParams, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing params: %w", err)
	}

	switch tx.TxType {
	case parser.MethodPreCommitSector, parser.MethodPreCommitSectorBatch, parser.MethodPreCommitSectorBatch2:
		sectorEvents, err := eg.parsePreCommitStage(ctx, tx, tipsetCid, params)
		if err != nil {
			return nil, fmt.Errorf("error parsing pre-commit stage: %w", err)
		}
		return sectorEvents, nil

	case parser.MethodConfirmSectorProofsValid, parser.MethodProveCommitAggregate, parser.MethodProveCommitSector, parser.MethodProveCommitSectors3, parser.MethodProveCommitSectorsNI:
		sectorEvents, err := eg.parseProveCommitStage(ctx, tx, tipsetCid, params)
		if err != nil {
			return nil, fmt.Errorf("error parsing prove commit stage: %w", err)
		}
		return sectorEvents, nil

	case parser.MethodTerminateSectors, parser.MethodDeclareFaults, parser.MethodDeclareFaultsRecovered:
		sectorEvents, err := eg.parseSectorTerminationFaultAndRecoveries(ctx, tx, tipsetCid, params)
		if err != nil {
			return nil, fmt.Errorf("error parsing sector termination fault and recoveries: %w", err)
		}
		return sectorEvents, nil

	case parser.MethodExtendSectorExpiration, parser.MethodExtendSectorExpiration2:
		sectorEvents, err := eg.parseSectorExpiryExtensions(ctx, tx, tipsetCid, params)
		if err != nil {
			return nil, fmt.Errorf("error parsing sector expiry extensions: %w", err)
		}
		return sectorEvents, nil

	}

	return nil, nil
}

func (eg *eventGenerator) parsePreCommitStage(_ context.Context, tx *types.Transaction, tipsetCid string, params map[string]interface{}) ([]*types.MinerSectorEvent, error) {
	var sectorEvents []*types.MinerSectorEvent

	addEvent := func(params map[string]interface{}) error {
		sealProof, err := getInteger[int64](params, KeySealProof, false)
		if err != nil {
			return fmt.Errorf("error parsing seal proof: %w", err)
		}

		sectorNumber, err := getInteger[uint64](params, KeySectorNumber, false)
		if err != nil {
			return fmt.Errorf("error parsing sector number: %w", err)
		}

		dealIDs, err := getIntegerSlice[uint64](params, KeyDealIDs, true)
		if err != nil {
			return fmt.Errorf("error parsing deal ids: %w", err)
		}

		expiration, err := getInteger[int64](params, KeyExpiration, false)
		if err != nil {
			return fmt.Errorf("error parsing expiration: %w", err)
		}
		jsonData, err := json.Marshal(map[string]interface{}{
			KeySectorNumber: sectorNumber,
			KeyExpiration:   expiration,
			KeySectorSize:   sectorProofToBigInt(sealProof).Uint64(),
			KeyDealIDs:      dealIDs,
		})
		if err != nil {
			return fmt.Errorf("error marshaling event: %w", err)
		}
		sectorEvents = append(sectorEvents, createSectorEvent(tipsetCid, tx, sectorNumber, jsonData))
		return nil
	}

	if tx.TxType == parser.MethodPreCommitSector {
		err := addEvent(params)
		if err != nil {
			return nil, err
		}
	} else {
		// parser.MethodPreCommitSectorBatch, parser.MethodPreCommitSectorBatch2,
		sectors, err := getSlice[map[string]interface{}](params, KeySectors, false)
		if err != nil {
			return nil, fmt.Errorf("error parsing sectors: %w", err)
		}
		for _, sector := range sectors {
			err := addEvent(sector)
			if err != nil {
				return nil, fmt.Errorf("error adding event: %w", err)
			}
		}
	}
	return sectorEvents, nil
}

func (eg *eventGenerator) parseProveCommitStage(ctx context.Context, tx *types.Transaction, tipsetCid string, params map[string]interface{}) ([]*types.MinerSectorEvent, error) {
	var sectorEvents []*types.MinerSectorEvent

	switch tx.TxType {
	case parser.MethodProveCommitSector:
		sectorNumber, err := getInteger[uint64](params, KeySectorNumber, false)
		if err != nil {
			return nil, fmt.Errorf("error parsing sector number: %w", err)
		}
		jsonData, err := json.Marshal(map[string]interface{}{
			KeySectorNumber: sectorNumber,
		})
		if err != nil {
			return nil, fmt.Errorf("error marshaling event: %w", err)
		}
		sectorEvents = append(sectorEvents, createSectorEvent(tipsetCid, tx, sectorNumber, jsonData))

		return sectorEvents, nil
	case parser.MethodProveCommitAggregate:
		return eg.parseProveCommitAggregate(ctx, tx, tipsetCid, params)
	case parser.MethodConfirmSectorProofsValid:
		return eg.parseConfirmSectorProofsValid(ctx, tx, tipsetCid, params)
	case parser.MethodProveCommitSectors3:
		return eg.parseProveCommitSectors3(ctx, tx, tipsetCid, params)
	case parser.MethodProveCommitSectorsNI:
		return eg.parseProveCommitSectorsNI(ctx, tx, tipsetCid, params)
	}
	return nil, fmt.Errorf("unexpected method: %s", tx.TxType)
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
		return nil, fmt.Errorf("error parsing events: %w", err)
	}
	for _, event := range events {
		sectorBitField, err := getIntegerSlice[int](event, KeySectors, false)
		if err != nil {
			return nil, fmt.Errorf("error parsing integer slice: %w", err)
		}

		sectorNumbers, err := jsonEncodedBitfieldToSectorNumbers(sectorBitField)
		if err != nil {
			return nil, fmt.Errorf("error parsing sector bitfield: %w", err)
		}
		event[KeySectorNumbers] = sectorNumbers
		jsonData, err := json.Marshal(event)
		if err != nil {
			return nil, fmt.Errorf("error marshaling event: %w", err)
		}
		for _, sectorNumber := range sectorNumbers {
			sectorEvents = append(sectorEvents, createSectorEvent(tipsetCid, tx, sectorNumber, jsonData))
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
			return nil, fmt.Errorf("error parsing integer slice: %w", err)
		}
		newExpiration, err := getInteger[int64](extension, KeyNewExpiration, false)
		if err != nil {
			return nil, fmt.Errorf("error parsing new expiration: %w", err)
		}

		sectorNumbers, err := jsonEncodedBitfieldToSectorNumbers(sectorBitField)
		if err != nil {
			return nil, fmt.Errorf("error parsing sector bitfield: %w", err)
		}
		jsonData, err := json.Marshal(map[string]interface{}{
			KeySectors:       sectorBitField,
			KeySectorNumbers: sectorNumbers,
			KeyNewExpiration: newExpiration,
		})
		if err != nil {
			return nil, fmt.Errorf("error marshaling event: %w", err)
		}
		for _, sectorNumber := range sectorNumbers {
			sectorEvents = append(sectorEvents, createSectorEvent(tipsetCid, tx, sectorNumber, jsonData))
		}
	}
	return sectorEvents, nil
}

func (eg *eventGenerator) parseProveCommitSectorsNI(_ context.Context, tx *types.Transaction, tipsetCid string, params map[string]interface{}) ([]*types.MinerSectorEvent, error) {
	sectorActivations, err := getSlice[map[string]interface{}](params, KeySectors, false)
	if err != nil {
		return nil, err
	}
	var sectorEvents []*types.MinerSectorEvent
	for _, sectorActivation := range sectorActivations {
		sectorNumber, err := getInteger[uint64](sectorActivation, KeySectorNumber, false)
		if err != nil {
			return nil, fmt.Errorf("error parsing sector number: %w", err)
		}
		jsonData, err := json.Marshal(sectorActivation)
		if err != nil {
			return nil, fmt.Errorf("error marshaling event: %w", err)
		}
		sectorEvents = append(sectorEvents, createSectorEvent(tipsetCid, tx, sectorNumber, jsonData))
	}
	return sectorEvents, nil
}

func (eg *eventGenerator) parseProveCommitSectors3(_ context.Context, tx *types.Transaction, tipsetCid string, params map[string]interface{}) ([]*types.MinerSectorEvent, error) {
	sectorActivations, err := getSlice[map[string]interface{}](params, KeySectorActivations, false)
	if err != nil {
		return nil, err
	}
	var sectorEvents []*types.MinerSectorEvent
	for _, sectorActivation := range sectorActivations {
		sectorNumber, err := getInteger[uint64](sectorActivation, KeySectorNumber, false)
		if err != nil {
			return nil, fmt.Errorf("error parsing sector number: %w", err)
		}
		jsonData, err := json.Marshal(sectorActivation)
		if err != nil {
			return nil, fmt.Errorf("error marshaling event: %w", err)
		}
		sectorEvents = append(sectorEvents, createSectorEvent(tipsetCid, tx, sectorNumber, jsonData))
	}
	return sectorEvents, nil
}

func (eg *eventGenerator) parseConfirmSectorProofsValid(_ context.Context, tx *types.Transaction, tipsetCid string, params map[string]interface{}) ([]*types.MinerSectorEvent, error) {
	sectors, err := getIntegerSlice[int64](params, KeySectors, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing integer slice: %w", err)
	}
	var sectorEvents []*types.MinerSectorEvent
	jsonData, err := json.Marshal(map[string]interface{}{
		KeySectorNumbers: sectors,
	})
	if err != nil {
		return nil, fmt.Errorf("error marshaling event: %w", err)
	}
	for _, sector := range sectors {
		//nolint:gosec
		sectorEvents = append(sectorEvents, createSectorEvent(tipsetCid, tx, uint64(sector), jsonData))
	}
	return sectorEvents, nil
}

func (eg *eventGenerator) parseProveCommitAggregate(_ context.Context, tx *types.Transaction, tipsetCid string, params map[string]interface{}) ([]*types.MinerSectorEvent, error) {
	sectorBitField, err := getIntegerSlice[int](params, KeySectorNumbers, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing integer slice: %w", err)
	}
	var sectorEvents []*types.MinerSectorEvent
	sectorNumbers, err := jsonEncodedBitfieldToSectorNumbers(sectorBitField)
	if err != nil {
		return nil, fmt.Errorf("error parsing sector bitfield: %w", err)
	}
	jsonData, err := json.Marshal(map[string]interface{}{
		KeySectorNumbers: sectorNumbers,
		KeySectors:       sectorBitField,
	})
	if err != nil {
		return nil, fmt.Errorf("error marshaling event: %w", err)
	}
	for _, sectorNumber := range sectorNumbers {
		sectorEvents = append(sectorEvents, createSectorEvent(tipsetCid, tx, sectorNumber, jsonData))
	}
	return sectorEvents, nil
}

// the sector bit field is a range of bits representing the different sector numbers.
// example: sectors: [ 0 1 2 3] -> bitfield: [1 1 1 1 ] -> JSON: [0,4]
// example: sectors: [0 1 3 4 5] -> bitfield: [1 1 0 1 1 1] -> JSON: [ 0,2,1,3 ]
// the JSON format always starts with a 0 and proceeds with the 0/1 pattern
// see here: https://pkg.go.dev/github.com/filecoin-project/go-bitfield@v0.2.4/rle#RLE.MarshalJSON
func jsonEncodedBitfieldToSectorNumbers(bitField []int) ([]uint64, error) {
	// pre-allocate for worst-case
	sectorNumbers := make([]uint64, 0, len(bitField))

	var parsedBitField bitfield.BitField
	bitFieldJSON, err := json.Marshal(bitField)
	if err != nil {
		return nil, fmt.Errorf("error marshaling json encoded bitfield: %w", err)
	}

	err = parsedBitField.UnmarshalJSON(bitFieldJSON)
	if err != nil {
		return nil, fmt.Errorf("error parsing json encoded bitfield: %w", err)
	}

	iter, err := parsedBitField.BitIterator()
	if err != nil {
		return nil, fmt.Errorf("error iterating over bitfield: %w", err)
	}

	for iter.HasNext() {
		sectorNumber, err := iter.Next()
		if err != nil {
			return nil, fmt.Errorf("error getting next sector number: %w", err)
		}
		sectorNumbers = append(sectorNumbers, sectorNumber)
	}
	return sectorNumbers, nil
}

func createSectorEvent(tipsetCid string, tx *types.Transaction, sectorNumber uint64, jsonData []byte) *types.MinerSectorEvent {
	return &types.MinerSectorEvent{
		ID:           tools.BuildId(tipsetCid, tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType),
		MinerAddress: tx.TxTo,
		SectorNumber: sectorNumber,
		Height:       tx.Height,
		TxCid:        tx.TxCid,
		ActionType:   tx.TxType,
		Data:         string(jsonData),
		TxTimestamp:  tx.TxTimestamp,
	}
}
