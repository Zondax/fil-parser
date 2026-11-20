package deals

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

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
	return strings.EqualFold(tx.TxType, parser.MethodOnMinerSectorsTerminate)
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
// `DealIDs` for the first shape and `SectorBitField` for the second. The tipset at the NV22 upgrade
// height still ran actors v12, and the actor parser decodes it with the v12 type
// (tools.ExecutedVersionFromHeight), so it yields `DealIDs`. Metadata produced by an older parser
// decoded that tipset with the v13 type, which reads the v12 deal ids as a CBOR array
// (`SectorNumbers`); that shape is still accepted. No v13+ market accepts an array (its params are
// a BitField), so in a successful call an array is always the v12 deal ids. Duplicate ids in one call are emitted once
// (the v12 market skips the repeat as already slashed).
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

	dealIDsKey := ""
	if _, ok := params[KeyDealIDs]; ok {
		dealIDsKey = KeyDealIDs
	} else if numbers, ok := params[KeySectorNumbers].([]interface{}); ok && len(numbers) > 0 {
		dealIDsKey = KeySectorNumbers
	}
	if dealIDsKey != "" {
		dealIDs, err := exactUint64Slice(params[dealIDsKey])
		if err != nil {
			return nil, fmt.Errorf("error parsing deal ids: %w", err)
		}
		terminations := make([]*types.DealsTerminations, 0, len(dealIDs))
		seen := make(map[uint64]bool, len(dealIDs))
		for _, dealID := range dealIDs {
			if seen[dealID] {
				continue
			}
			seen[dealID] = true
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
// params must come from decodeExact.
func terminatedSectors(params map[string]interface{}) ([]uint64, error) {
	raw, ok := params[KeySectorBitField]
	if !ok || raw == nil {
		return nil, fmt.Errorf("params have neither %s, %s nor %s", KeyDealIDs, KeySectorNumbers, KeySectorBitField)
	}
	sectors, err := decodeBitfieldJSON(raw, maxTerminatedSectors)
	if err != nil {
		return nil, fmt.Errorf("sector bitfield: %w", err)
	}
	return sectors, nil
}

// decodeBitfieldJSON reads a bitfield stored by the actor parser in its JSON form (run lengths,
// go-bitfield RLE MarshalJSON) and returns its set bits. It fails, without expanding the
// bitfield, when more than max bits are set.
func decodeBitfieldJSON(raw interface{}, max uint64) ([]uint64, error) {
	encoded, err := json.Marshal(raw)
	if err != nil {
		return nil, fmt.Errorf("error encoding: %w", err)
	}
	var bf bitfield.BitField
	if err := bf.UnmarshalJSON(encoded); err != nil {
		return nil, fmt.Errorf("error decoding: %w", err)
	}
	count, err := bf.Count()
	if err != nil {
		return nil, fmt.Errorf("error counting: %w", err)
	}
	if count > max {
		return nil, fmt.Errorf("has %d bits set, more than %d", count, max)
	}
	bits, err := bf.All(max)
	if err != nil {
		return nil, fmt.Errorf("error reading: %w", err)
	}
	return bits, nil
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
