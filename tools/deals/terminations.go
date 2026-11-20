package deals

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/filecoin-project/go-bitfield"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/tools/common"
	"github.com/zondax/fil-parser/types"
)

const (
	KeyEpoch          = "Epoch"
	KeySectorBitField = "SectorBitField"
	KeySectorNumbers  = "SectorNumbers"

	// maxTerminatedSectors bounds the sectors read from one OnMinerSectorsTerminate call. The miner
	// sends at most policy.addressed_sectors_max (25,000) sectors per call, and the 1 MiB FVM block
	// limit caps it at ~130K, so anything above this is corrupt metadata, never expanded.
	maxTerminatedSectors = 1 << 20
)

func (eg *eventGenerator) isDealTermination(tx *types.Transaction) bool {
	return tx.TxType == parser.MethodOnMinerSectorsTerminate
}

// createDealTerminations builds the rows of a market OnMinerSectorsTerminate call. The caller has
// already checked that the tx goes to the market actor and succeeded (IsTxSuccess).
//
// The params shape depends on the actors version (builtin-actors actors/market/src/types.rs):
//   - up to v12 (NV21): {epoch, deal_ids}. The miner sends the deal ids of all the terminated
//     sectors, epoch being the epoch the sectors were terminated at (actors/miner/src/lib.rs
//     process_early_terminations / request_terminate_deals, chunks of 8192). One row per deal id.
//   - from v13 (NV22): {epoch, sectors}, sectors being a bitfield of the terminated sectors that
//     hold data, epoch the current epoch. The market looks the deals up in its own sector -> deals
//     map. One row per sector, with no deal id.
//
// In both cases the market skips deals whose end_epoch <= epoch ("do not slash expired deals"),
// so a row is a candidate early termination; the end_epoch filter is left to downstream, which
// has the proposals.
//
// The shape the actor parser produced is used, not the network version: the metadata holds
// `DealIDs` for the first shape and `SectorBitField` (or `SectorNumbers`) for the second.
func (eg *eventGenerator) createDealTerminations(tx *types.Transaction) ([]*types.DealsTerminations, error) {
	metadata, err := decodeExact(tx.TxMetadata)
	if err != nil {
		return nil, err
	}
	params, err := common.GetItem[map[string]interface{}](metadata, KeyParams, false)
	if err != nil {
		return nil, err
	}
	epoch, err := exactInt64(params, KeyEpoch)
	if err != nil {
		return nil, err
	}

	newRow := func(kind string, value uint64) *types.DealsTerminations {
		return &types.DealsTerminations{
			ID:               tools.BuildId(tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType, kind, fmt.Sprint(value)),
			Height:           tx.Height,
			TxCid:            tx.TxCid,
			ActorAddress:     tx.TxFrom,
			TerminationEpoch: epoch,
			ActionType:       tx.TxType,
			TxTimestamp:      tx.TxTimestamp,
		}
	}

	if _, ok := params[KeyDealIDs]; ok {
		dealIDs, err := exactUint64Slice(params[KeyDealIDs])
		if err != nil {
			return nil, fmt.Errorf("error parsing deal ids: %w", err)
		}
		terminations := make([]*types.DealsTerminations, 0, len(dealIDs))
		for _, dealID := range dealIDs {
			row := newRow("deal", dealID)
			row.DealID = &dealID
			terminations = append(terminations, row)
		}
		return terminations, nil
	}

	sectorNumbers, err := terminatedSectors(params)
	if err != nil {
		return nil, err
	}
	terminations := make([]*types.DealsTerminations, 0, len(sectorNumbers))
	for _, sectorNumber := range sectorNumbers {
		row := newRow("sector", sectorNumber)
		row.SectorNumber = &sectorNumber
		terminations = append(terminations, row)
	}
	return terminations, nil
}

// terminatedSectors reads the sectors of the actors v13+ params, which the actor parser stores as
// `SectorBitField` in the bitfield's JSON form (run lengths, go-bitfield RLE MarshalJSON).
//
// The actor parser's v13+ type also decodes a CBOR array, as `SectorNumbers`. No v13+ market can
// accept that (its params are a BitField), so a successful call with an array was executed with
// the pre-v13 {epoch, deal_ids} params and the array would be deal ids, not sectors. That means
// the height's actors version disagrees with the chain; it is reported as an error (the tx is
// skipped and counted), never guessed.
// params must come from decodeExact.
func terminatedSectors(params map[string]interface{}) ([]uint64, error) {
	if numbers, ok := params[KeySectorNumbers].([]interface{}); ok && len(numbers) > 0 {
		return nil, fmt.Errorf("params carry an array (%s) instead of a sector bitfield", KeySectorNumbers)
	}

	raw, ok := params[KeySectorBitField]
	if !ok || raw == nil {
		return nil, fmt.Errorf("params have neither %s, %s nor %s", KeyDealIDs, KeySectorBitField, KeySectorNumbers)
	}
	encoded, err := json.Marshal(raw)
	if err != nil {
		return nil, fmt.Errorf("error encoding sector bitfield: %w", err)
	}
	var bf bitfield.BitField
	if err := bf.UnmarshalJSON(encoded); err != nil {
		return nil, fmt.Errorf("error decoding sector bitfield: %w", err)
	}
	count, err := bf.Count()
	if err != nil {
		return nil, fmt.Errorf("error counting sector bitfield: %w", err)
	}
	if count > maxTerminatedSectors {
		return nil, fmt.Errorf("sector bitfield has %d sectors, more than %d", count, maxTerminatedSectors)
	}
	sectors, err := bf.All(maxTerminatedSectors)
	if err != nil {
		return nil, fmt.Errorf("error reading sector bitfield: %w", err)
	}
	return sectors, nil
}

// exactUint64Slice reads a JSON array decoded by decodeExact as uint64s. null is an empty slice.
func exactUint64Slice(value interface{}) ([]uint64, error) {
	if value == nil {
		return nil, nil
	}
	values, ok := value.([]interface{})
	if !ok {
		return nil, fmt.Errorf("not an array: %T", value)
	}
	out := make([]uint64, 0, len(values))
	for i, v := range values {
		number, ok := v.(json.Number)
		if !ok {
			return nil, fmt.Errorf("item %d is not a number: %T", i, v)
		}
		parsed, err := strconv.ParseUint(number.String(), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("item %d: %w", i, err)
		}
		out = append(out, parsed)
	}
	return out, nil
}
