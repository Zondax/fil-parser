package deals

import (
	"fmt"
	"strings"

	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/tools/common"
	"github.com/zondax/fil-parser/types"
)

const (
	KeyUpdates     = "Updates"
	KeySectorID    = "SectorID"
	KeyUpdateDeals = "Deals"

	// maxReplicaUpdates is a sanity bound on the updates read from one ProveReplicaUpdates(2) call.
	// The actors cap a call at PROVE_REPLICA_UPDATES_MAX_SIZE = PRE_COMMIT_SECTOR_BATCH_MAX_SIZE
	// (256, specs-actors v7 and builtin-actors v8-v12); the bound only protects against corrupt
	// metadata.
	maxReplicaUpdates = 10_000
)

// isSectorMappingMessage reports whether tx is a miner method whose params map sectors to deals
// before actors v13 (NV22). From actors v13 on, every activation carries its sector number
// (BatchActivateDeals, SectorContentChanged), and the mapping rows are copied from the activations.
func (eg *eventGenerator) isSectorMappingMessage(actorName string, tx *types.Transaction) bool {
	if !strings.Contains(actorName, manifest.MinerKey) {
		return false
	}
	switch tx.TxType {
	case parser.MethodPreCommitSector,
		parser.MethodPreCommitSectorBatch,
		parser.MethodPreCommitSectorBatch2,
		parser.MethodProveReplicaUpdates,
		parser.MethodProveReplicaUpdates2:
	default:
		return false
	}
	//#nosec G115
	return tools.VersionFromHeight(eg.network, int64(tx.Height)).NodeVersion() < tools.V22.NodeVersion()
}

// createSectorMappings builds the sector -> deal rows of a successful pre-commit or replica update
// sent to a miner before actors v13. The miner is tx.TxTo.
func (eg *eventGenerator) createSectorMappings(tx *types.Transaction) ([]*types.DealsSectorMappings, error) {
	metadata, err := decodeExact(tx.TxMetadata)
	if err != nil {
		return nil, err
	}

	switch tx.TxType {
	case parser.MethodPreCommitSector:
		params, err := common.GetItem[map[string]interface{}](metadata, KeyParams, false)
		if err != nil {
			return nil, err
		}
		return eg.preCommitMappings(tx, []map[string]interface{}{params})
	case parser.MethodPreCommitSectorBatch, parser.MethodPreCommitSectorBatch2:
		params, err := common.GetItem[map[string]interface{}](metadata, KeyParams, false)
		if err != nil {
			return nil, err
		}
		sectors, err := common.GetSlice[map[string]interface{}](params, KeySectors, false)
		if err != nil {
			return nil, err
		}
		return eg.preCommitMappings(tx, sectors)
	case parser.MethodProveReplicaUpdates, parser.MethodProveReplicaUpdates2:
		return eg.replicaUpdateMappings(tx, metadata)
	}
	return nil, nil
}

// preCommitMappings emits one row per deal id of each pre-committed sector.
//
// A pre-commit batch is all-or-nothing: an invalid sector aborts the whole call (specs-actors v7
// PreCommitSectorBatch, builtin-actors v8-v12 pre_commit_sector_batch(_inner) have no per-sector
// skip), so every sector of a successful call was pre-committed. A sector number is allocated once
// and never reused, and a sector whose deals fail to activate at prove-commit is not proven, so a
// sector that is ever terminated holds exactly these deals.
func (eg *eventGenerator) preCommitMappings(tx *types.Transaction, sectors []map[string]interface{}) ([]*types.DealsSectorMappings, error) {
	mappings := []*types.DealsSectorMappings{}
	for i, sector := range sectors {
		dealIDs, err := exactUint64Slice(sector[KeyDealIDs])
		if err != nil {
			return nil, fmt.Errorf("sector %d deal ids: %w", i, err)
		}
		if len(dealIDs) == 0 {
			continue
		}
		sectorNumber, err := exactUint64(sector, KeySectorNumber)
		if err != nil {
			return nil, fmt.Errorf("sector %d: %w", i, err)
		}
		for _, dealID := range dealIDs {
			mappings = append(mappings, newSectorMapping(tx, tx.TxTo, sectorNumber, dealID))
		}
	}
	return mappings, nil
}

// replicaUpdateMappings emits one row per deal of each update that succeeded.
//
// In every actors version with replica updates (specs-actors v7 ProveReplicaUpdates; builtin-actors
// v8-v12 prove_replica_updates and, from v9, prove_replica_updates2) the method returns a BitField
// of the sector numbers that were updated. An update that fails validation or deal activation is
// skipped, not reverted; an error after validation aborts the whole call. An updated sector's
// deal ids become exactly that update's deals. A sector listed twice is skipped as a duplicate
// from its second occurrence on (the first occurrence marks it as seen before any other check), so
// a successful sector's deals are those of its first occurrence in the params.
//
// The return is required: without it the successful updates of a batch are unknown and nothing is
// guessed. ProveReplicaUpdates metadata stored before the return was parsed has none.
//
// metadata must come from decodeExact.
func (eg *eventGenerator) replicaUpdateMappings(tx *types.Transaction, metadata map[string]interface{}) ([]*types.DealsSectorMappings, error) {
	params, err := common.GetItem[map[string]interface{}](metadata, KeyParams, false)
	if err != nil {
		return nil, err
	}
	updates, err := common.GetSlice[map[string]interface{}](params, KeyUpdates, false)
	if err != nil {
		return nil, err
	}
	if len(updates) > maxReplicaUpdates {
		return nil, fmt.Errorf("%d updates, more than %d", len(updates), maxReplicaUpdates)
	}
	ret, ok := metadata[KeyReturn]
	if !ok || ret == nil {
		return nil, fmt.Errorf("no return: the updated sectors are unknown")
	}
	// at most one bit per update
	updated, err := decodeBitfieldJSON(ret, uint64(len(updates)))
	if err != nil {
		return nil, fmt.Errorf("updated sectors bitfield: %w", err)
	}

	// deals of the first occurrence of each sector
	deals := make(map[uint64][]uint64, len(updates))
	for i, update := range updates {
		sectorNumber, err := exactUint64(update, KeySectorID)
		if err != nil {
			return nil, fmt.Errorf("update %d: %w", i, err)
		}
		if _, seen := deals[sectorNumber]; seen {
			continue
		}
		dealIDs, err := exactUint64Slice(update[KeyUpdateDeals])
		if err != nil {
			return nil, fmt.Errorf("update %d deals: %w", i, err)
		}
		if dealIDs == nil {
			dealIDs = []uint64{}
		}
		deals[sectorNumber] = dealIDs
	}

	mappings := []*types.DealsSectorMappings{}
	for _, sectorNumber := range updated {
		dealIDs, ok := deals[sectorNumber]
		if !ok {
			return nil, fmt.Errorf("updated sector %d is not in the params", sectorNumber)
		}
		for _, dealID := range dealIDs {
			mappings = append(mappings, newSectorMapping(tx, tx.TxTo, sectorNumber, dealID))
		}
	}
	return mappings, nil
}

// activationMappings copies the sector of each activation that has one.
func activationMappings(tx *types.Transaction, activations []*types.DealsActivations) []*types.DealsSectorMappings {
	mappings := []*types.DealsSectorMappings{}
	for _, activation := range activations {
		if activation.SectorNumber == nil {
			continue
		}
		mappings = append(mappings, newSectorMapping(tx, activation.ActorAddress, *activation.SectorNumber, activation.DealID))
	}
	return mappings
}

func newSectorMapping(tx *types.Transaction, miner string, sectorNumber, dealID uint64) *types.DealsSectorMappings {
	return &types.DealsSectorMappings{
		ID:           tools.BuildId(tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType, "sector-mapping", fmt.Sprint(sectorNumber), fmt.Sprint(dealID)),
		Height:       tx.Height,
		TxCid:        tx.TxCid,
		ActorAddress: miner,
		SectorNumber: sectorNumber,
		DealID:       dealID,
		ActionType:   tx.TxType,
		TxTimestamp:  tx.TxTimestamp,
	}
}
