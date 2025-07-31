package miner

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/tools/common"
	"github.com/zondax/fil-parser/types"
)

const (
	KeySectorNumber          = "SectorNumber"
	KeySectorNumbers         = "SectorNumbers"
	KeySectors               = "Sectors"
	KeyExpiration            = "Expiration"
	KeySectorSize            = "SectorSize"
	KeyDealIDs               = "DealIDs"
	KeyNewExpiration         = "NewExpiration"
	KeyParams                = "Params"
	KeyTerminations          = "Terminations"
	KeyFaults                = "Faults"
	KeyRecoveries            = "Recoveries"
	KeyExtensions            = "Extensions"
	KeySealProof             = "SealProof"
	KeySectorActivations     = "SectorActivations"
	KeyPieces                = "Pieces"
	KeyNotify                = "Notify"
	KeyAddress               = "Address"
	KeySealerID              = "SealerID"
	KeyVerifiedAllocationKey = "VerifiedAllocationKey"
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

	params, err := common.GetItem[map[string]interface{}](value, KeyParams, false)
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
		sealProof, err := common.GetInteger[int64](params, KeySealProof, false)
		if err != nil {
			return fmt.Errorf("error parsing seal proof: %w", err)
		}

		sectorNumber, err := common.GetInteger[uint64](params, KeySectorNumber, false)
		if err != nil {
			return fmt.Errorf("error parsing sector number: %w", err)
		}

		dealIDs, err := common.GetIntegerSlice[uint64](params, KeyDealIDs, true)
		if err != nil {
			return fmt.Errorf("error parsing deal ids: %w", err)
		}

		expiration, err := common.GetInteger[int64](params, KeyExpiration, false)
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
		sectors, err := common.GetSlice[map[string]interface{}](params, KeySectors, false)
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
		sectorNumber, err := common.GetInteger[uint64](params, KeySectorNumber, false)
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
	events, err := common.GetSlice[map[string]interface{}](params, parameterName, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing events: %w", err)
	}
	for _, event := range events {
		sectorBitField, err := common.GetIntegerSlice[int](event, KeySectors, false)
		if err != nil {
			return nil, fmt.Errorf("error parsing integer slice: %w", err)
		}

		sectorNumbers, err := common.JsonEncodedBitfieldToIDs(sectorBitField)
		if err != nil {
			return nil, fmt.Errorf("error parsing sector bitfield: %w", err)
		}

		// only keep the Deadline and Partition info
		delete(event, KeySectors)
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
	extensions, err := common.GetSlice[map[string]interface{}](params, KeyExtensions, false)
	if err != nil {
		return nil, err
	}
	for _, extension := range extensions {
		sectorBitField, err := common.GetIntegerSlice[int](extension, KeySectors, false)
		if err != nil {
			return nil, fmt.Errorf("error parsing integer slice: %w", err)
		}
		newExpiration, err := common.GetInteger[int64](extension, KeyNewExpiration, false)
		if err != nil {
			return nil, fmt.Errorf("error parsing new expiration: %w", err)
		}

		sectorNumbers, err := common.JsonEncodedBitfieldToIDs(sectorBitField)
		if err != nil {
			return nil, fmt.Errorf("error parsing sector bitfield: %w", err)
		}
		jsonData, err := json.Marshal(map[string]interface{}{
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
	sectorActivations, err := common.GetSlice[map[string]interface{}](params, KeySectors, false)
	if err != nil {
		return nil, err
	}
	var sectorEvents []*types.MinerSectorEvent
	for _, sectorActivation := range sectorActivations {
		sectorNumber, err := common.GetInteger[uint64](sectorActivation, KeySectorNumber, false)
		if err != nil {
			return nil, fmt.Errorf("error parsing sector number: %w", err)
		}
		sealerID, err := common.GetInteger[uint64](sectorActivation, KeySealerID, false)
		if err != nil {
			return nil, fmt.Errorf("error parsing sealer id: %w", err)
		}

		if eg.config.ConsolidateRobustAddress {
			consolidatedSealerID, err := eg.consolidateIDAddress(sealerID)
			if err != nil {
				eg.logger.Errorf("error consolidating sealer id: %w", err)
			} else {
				sectorActivation[KeySealerID] = consolidatedSealerID
			}
		}

		jsonData, err := json.Marshal(sectorActivation)
		if err != nil {
			return nil, fmt.Errorf("error marshaling event: %w", err)
		}
		sectorEvents = append(sectorEvents, createSectorEvent(tipsetCid, tx, sectorNumber, jsonData))
	}
	return sectorEvents, nil
}

func (eg *eventGenerator) parseProveCommitSectors3(ctx context.Context, tx *types.Transaction, tipsetCid string, params map[string]interface{}) ([]*types.MinerSectorEvent, error) {
	sectorActivations, err := common.GetSlice[map[string]interface{}](params, KeySectorActivations, false)
	if err != nil {
		return nil, err
	}
	var sectorEvents []*types.MinerSectorEvent
	for _, sectorActivation := range sectorActivations {
		sectorNumber, err := common.GetInteger[uint64](sectorActivation, KeySectorNumber, false)
		if err != nil {
			return nil, fmt.Errorf("error parsing sector number: %w", err)
		}

		if eg.config.ConsolidateRobustAddress {
			pieces, err := common.GetSlice[map[string]interface{}](sectorActivation, KeyPieces, true)
			if err != nil {
				return nil, fmt.Errorf("error parsing pieces: %w", err)
			}
			if len(pieces) > 0 {
				parsedPieces, err := eg.consolidatePieceActivationManifests(ctx, pieces)
				if err != nil {
					return nil, fmt.Errorf("error consolidating piece activation manifests: %w", err)
				}
				sectorActivation[KeyPieces] = parsedPieces
			}
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
	sectors, err := common.GetIntegerSlice[int64](params, KeySectors, false)
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
	sectorBitField, err := common.GetIntegerSlice[int](params, KeySectorNumbers, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing integer slice: %w", err)
	}
	var sectorEvents []*types.MinerSectorEvent
	sectorNumbers, err := common.JsonEncodedBitfieldToIDs(sectorBitField)
	if err != nil {
		return nil, fmt.Errorf("error parsing sector bitfield: %w", err)
	}
	jsonData, err := json.Marshal(map[string]interface{}{
		KeySectorNumbers: sectorNumbers,
	})
	if err != nil {
		return nil, fmt.Errorf("error marshaling event: %w", err)
	}
	for _, sectorNumber := range sectorNumbers {
		sectorEvents = append(sectorEvents, createSectorEvent(tipsetCid, tx, sectorNumber, jsonData))
	}
	return sectorEvents, nil
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

func (eg *eventGenerator) consolidatePieceActivationManifests(_ context.Context, pieces []map[string]interface{}) ([]map[string]interface{}, error) {
	parsedPieces := make([]map[string]interface{}, 0, len(pieces))
	for _, piece := range pieces {
		verifiedAllocationKey, err := common.GetItem[map[string]interface{}](piece, KeyVerifiedAllocationKey, true)
		if err != nil {
			return nil, fmt.Errorf("error parsing verified allocation key: %w", err)
		}
		if len(verifiedAllocationKey) > 0 {
			clientIDAddrStr, err := common.GetInteger[uint64](verifiedAllocationKey, KeyAddress, false)
			if err != nil {
				eg.logger.Errorf("error parsing client id address: %w", err)
				break
			}
			consolidatedClientIDAddr, err := eg.consolidateIDAddress(clientIDAddrStr)
			if err != nil {
				eg.logger.Errorf("error consolidating client id address: %w", err)
				break
			}
			verifiedAllocationKey[KeyAddress] = consolidatedClientIDAddr
			piece[KeyVerifiedAllocationKey] = verifiedAllocationKey

		}

		dataActivationNotifications, err := common.GetSlice[map[string]interface{}](piece, KeyNotify, true)
		if err != nil {
			return nil, fmt.Errorf("error parsing notify: %w", err)
		}
		if len(dataActivationNotifications) > 0 {
			parsedDataActivationNotifications := make([]map[string]interface{}, 0, len(dataActivationNotifications))
			for _, notify := range dataActivationNotifications {
				addrStr, err := common.GetItem[string](notify, KeyAddress, false)
				if err != nil {
					eg.logger.Errorf("error parsing notify number: %w", err)
					break
				}
				consolidatedAddr, err := eg.consolidateAddress(addrStr)
				if err != nil {
					eg.logger.Errorf("error consolidating address: %w", err)
					break
				}
				notify[KeyAddress] = consolidatedAddr
				parsedDataActivationNotifications = append(parsedDataActivationNotifications, notify)
			}
			// only add the parsed data activation notifications if all the data activation notifications were parsed successfully
			if len(parsedDataActivationNotifications) == len(dataActivationNotifications) {
				piece[KeyNotify] = parsedDataActivationNotifications
			}

		}
		parsedPieces = append(parsedPieces, piece)
	}
	// only return the parsed pieces if all the pieces were parsed successfully
	if len(parsedPieces) == len(pieces) {
		return parsedPieces, nil
	}

	return pieces, nil
}
