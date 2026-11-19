package deals

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"math/big"

	cbg "github.com/whyrusleeping/cbor-gen"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/tools/common"
	"github.com/zondax/fil-parser/types"
)

const (
	KeySector                 = "Sector"
	KeyMinimumCommitmentEpoch = "MinimumCommitmentEpoch"
	KeyAdded                  = "Added"
	KeyPayload                = "Payload"
)

// parseSectorContentChanged builds deal activations from the market actor's SectorContentChanged
// method, which is how deals are activated by ProveCommitSectors3 / ProveReplicaUpdates3
// (actors v13 / NV22 onwards).
//
// The miner sends one call per notifee with one entry per sector, each listing the pieces added
// to it; for the market the piece payload is a CBOR-encoded DealID. The market accepts or rejects
// every piece independently and returns one `accepted` flag per piece, per sector, in params order.
// A piece is accepted only once the deal passed every check (payload decodes as a DealID, deal not
// already activated in this call or before, proposal found and not expired, provider is the
// calling miner, deal end <= minimum_commitment_epoch, piece CID and size match the proposal), and
// exactly then the market emits a `deal-activated` event and stores the deal state.
// See builtin-actors actors/market/src/lib.rs sector_content_changed (same rule v13 to v19).
//
// So a row is emitted for each piece whose flag is true and whose payload decodes as a DealID.
// Rejected pieces are skipped, as are pieces whose payload is not a DealID (the market never
// accepts those).
//
// SectorExpiry is the sector's minimum_commitment_epoch, which the miner fills with the sector
// expiration (actors/miner/src/notifications.rs), i.e. the same value as the legacy SectorExpiry.
//
// The space row carries the piece's padded size in NonVerifiedDealSpace, which on these rows is the
// deal's whole size whether it is verified or not, and VerifiedDealSpace is 0. Whether a deal is
// verified is not derivable from this call without a state lookup (from v19 the proposal flag no
// longer determines verified power at all); use the deal proposal for that.
func (eg *eventGenerator) parseSectorContentChanged(tx *types.Transaction, metadata map[string]interface{}) ([]*types.DealsActivations, []*types.DealsSpaceInfo, error) {
	sectors, err := common.GetSlice[map[string]interface{}](metadata, KeyParams, true)
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing sector content changed params: %w", err)
	}
	if len(sectors) == 0 {
		return nil, nil, nil
	}

	accepted, err := sectorContentChangedAccepted(metadata[KeyReturn], len(sectors))
	if err != nil {
		return nil, nil, err
	}

	dealActivations := []*types.DealsActivations{}
	dealSpaceInfo := []*types.DealsSpaceInfo{}
	for i, sector := range sectors {
		sectorNumber, err := common.GetInteger[uint64](sector, KeySector, false)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing sector %d number: %w", i, err)
		}
		minimumCommitmentEpoch, err := common.GetInteger[int64](sector, KeyMinimumCommitmentEpoch, false)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing sector %d minimum commitment epoch: %w", i, err)
		}
		pieces, err := common.GetSlice[map[string]interface{}](sector, KeyAdded, true)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing sector %d pieces: %w", i, err)
		}
		if i >= len(accepted) {
			return nil, nil, fmt.Errorf("sector content changed return has no entry for sector %d", i)
		}
		if len(accepted[i]) != len(pieces) {
			return nil, nil, fmt.Errorf("sector content changed return for sector %d has %d entries, expected %d", i, len(accepted[i]), len(pieces))
		}

		for j, piece := range pieces {
			if !accepted[i][j] {
				continue
			}
			dealID, ok, err := pieceDealID(piece)
			if err != nil {
				return nil, nil, fmt.Errorf("error parsing sector %d piece %d: %w", i, j, err)
			}
			if !ok {
				eg.logger.Warnf("accepted piece without a deal id payload, tx cid %s sector %d piece %d", tx.TxCid, i, j)
				continue
			}
			size, err := common.GetInteger[uint64](piece, KeySize, false)
			if err != nil {
				return nil, nil, fmt.Errorf("error parsing sector %d piece %d size: %w", i, j, err)
			}

			id := tools.BuildId(tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType, fmt.Sprint(dealID))
			dealActivations = append(dealActivations, &types.DealsActivations{
				ID:           id,
				ActorAddress: tx.TxFrom,
				Height:       tx.Height,
				TxCid:        tx.TxCid,
				DealID:       dealID,
				SectorExpiry: minimumCommitmentEpoch,
				SectorNumber: &sectorNumber,
				ActionType:   tx.TxType,
				TxTimestamp:  tx.TxTimestamp,
			})
			dealSpaceInfo = append(dealSpaceInfo, &types.DealsSpaceInfo{
				ID:                   id,
				Height:               tx.Height,
				ActorAddress:         tx.TxFrom,
				TxCid:                tx.TxCid,
				DealID:               dealID,
				GroupDealIDs:         []uint64{dealID},
				NonVerifiedDealSpace: new(big.Int).SetUint64(size),
				VerifiedDealSpace:    big.NewInt(0),
				ActionType:           tx.TxType,
				TxTimestamp:          tx.TxTimestamp,
			})
		}
	}

	return dealActivations, dealSpaceInfo, nil
}

// sectorContentChangedAccepted reads the per-sector, per-piece accepted flags.
// The return is [[bool]] (one array per sector). Parsers before the fix stored only the first
// sector's flags as a flat [bool]; that shape is only unambiguous for a single sector.
func sectorContentChangedAccepted(ret interface{}, sectors int) ([][]bool, error) {
	values, ok := ret.([]interface{})
	if !ok {
		return nil, fmt.Errorf("sector content changed return is not an array: %T", ret)
	}

	if len(values) > 0 {
		if _, flat := values[0].(bool); flat {
			if sectors != 1 {
				return nil, fmt.Errorf("sector content changed return is a flat array but there are %d sectors", sectors)
			}
			flags, err := toBools(values)
			if err != nil {
				return nil, err
			}
			return [][]bool{flags}, nil
		}
	}

	accepted := make([][]bool, 0, len(values))
	for i, value := range values {
		pieces, ok := value.([]interface{})
		if !ok && value != nil {
			return nil, fmt.Errorf("sector content changed return for sector %d is not an array: %T", i, value)
		}
		flags, err := toBools(pieces)
		if err != nil {
			return nil, err
		}
		accepted = append(accepted, flags)
	}
	if len(accepted) != sectors {
		return nil, fmt.Errorf("sector content changed return has %d sectors, expected %d", len(accepted), sectors)
	}
	return accepted, nil
}

func toBools(values []interface{}) ([]bool, error) {
	flags := make([]bool, len(values))
	for i, value := range values {
		flag, ok := value.(bool)
		if !ok {
			return nil, fmt.Errorf("piece return %d is not a bool: %T", i, value)
		}
		flags[i] = flag
	}
	return flags, nil
}

// pieceDealID decodes the piece payload as a CBOR unsigned integer (a DealID), as the market does.
// ok is false when the payload is not exactly one CBOR unsigned integer.
func pieceDealID(piece map[string]interface{}) (dealID uint64, ok bool, err error) {
	encoded, err := common.GetItem[string](piece, KeyPayload, true)
	if err != nil {
		return 0, false, err
	}
	payload, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return 0, false, fmt.Errorf("error decoding payload: %w", err)
	}

	reader := bytes.NewReader(payload)
	maj, value, err := cbg.CborReadHeader(reader)
	if err != nil || maj != cbg.MajUnsignedInt || reader.Len() != 0 {
		return 0, false, nil
	}
	return value, true, nil
}
