package deals

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/tools/common"
	"github.com/zondax/fil-parser/types"
)

const (
	KeySectors              = "Sectors"
	KeyActivations          = "Activations"
	KeyActivationResults    = "ActivationResults"
	KeyFailCodes            = "FailCodes"
	KeyIdx                  = "Idx"
	KeyDealWeight           = "DealWeight"
	KeyVerifiedDealWeight   = "VerifiedDealWeight"
	KeyNonVerifiedDealSpace = "NonVerifiedDealSpace"
	KeyVerifiedInfos        = "VerifiedInfos"
	KeySize                 = "Size"
	KeySectorNumber         = "SectorNumber"
)

func (eg *eventGenerator) createDealActivations(_ context.Context, tx *types.Transaction) ([]*types.DealsActivations, []*types.DealsSpaceInfo, error) {
	var (
		err             error
		dealActivations []*types.DealsActivations
		dealSpaceInfo   []*types.DealsSpaceInfo
		metadata        map[string]interface{}
	)

	if err = json.Unmarshal([]byte(tx.TxMetadata), &metadata); err != nil {
		return nil, nil, err
	}
	// The same metadata with numbers kept as json.Number, for the fields added with
	// SectorContentChanged support (sector numbers can exceed 2^53, which float64 rounds).
	// The legacy fields keep their float64 decoding.
	exactMetadata, err := decodeExact(tx.TxMetadata)
	if err != nil {
		return nil, nil, err
	}

	// SectorContentChanged params are an array of sectors, not a map, so it is handled before
	// the map-shaped params of the legacy methods are read.
	if tx.TxType == parser.MethodSectorContentChanged {
		dealActivations, err = eg.parseSectorContentChanged(tx, exactMetadata)
		if err != nil {
			return nil, nil, err
		}
		return dealActivations, nil, nil
	}

	params, err := common.GetItem[map[string]interface{}](metadata, KeyParams, false)
	if err != nil {
		return nil, nil, err
	}
	exactParams, err := common.GetItem[map[string]interface{}](exactMetadata, KeyParams, false)
	if err != nil {
		return nil, nil, err
	}

	switch tx.TxType {
	case parser.MethodVerifyDealsForActivation:
		ret, err := common.GetItem[map[string]interface{}](metadata, KeyReturn, false)
		if err != nil {
			return nil, nil, err
		}
		dealSpaceInfo, err = eg.parseVerifyDealsForActivation(tx, params, ret)
		if err != nil {
			return nil, nil, err
		}
	case parser.MethodActivateDeals, parser.MethodBatchActivateDeals:
		// return is empty for earlier network versions
		ret, err := common.GetItem[map[string]interface{}](metadata, KeyReturn, true)
		if err != nil {
			return nil, nil, err
		}
		dealActivations, dealSpaceInfo, err = eg.parseActivateDeals(tx, params, exactParams, ret)
		if err != nil {
			return nil, nil, err
		}
	}

	return dealActivations, dealSpaceInfo, nil
}

func (eg *eventGenerator) parseVerifyDealsForActivation(tx *types.Transaction, params, ret map[string]interface{}) ([]*types.DealsSpaceInfo, error) {
	//#nosec G115
	version := tools.VersionFromHeight(eg.network, int64(tx.Height))
	dealSpaceInfo := []*types.DealsSpaceInfo{}

	// Before V3, VerifyDealsForActivation has flat parameters
	/*
		type VerifyDealsForActivationParams struct {
			DealIDs      []abi.DealID
			SectorExpiry abi.ChainEpoch
			SectorStart  abi.ChainEpoch
		}

		type VerifyDealsForActivationReturn struct {
			DealWeight         abi.DealWeight
			VerifiedDealWeight abi.DealWeight
		}
	*/
	minVersion := tools.V9.NodeVersion()
	// calibration network min version is always V16 and this format is not used on the network
	if eg.network == tools.MainnetNetwork && version.NodeVersion() <= minVersion {
		dealIDs, nonVerifiedDealWeight, verifiedDealWeight, err := eg.getCommonVerifyDealForActivationFields(params, ret)
		if err != nil {
			return nil, err
		}

		for _, dealID := range dealIDs {
			dealSpaceInfo = append(dealSpaceInfo, &types.DealsSpaceInfo{
				ID:                   tools.BuildId(tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType, fmt.Sprint(dealID)),
				Height:               tx.Height,
				ActorAddress:         tx.TxFrom,
				TxCid:                tx.TxCid,
				DealID:               dealID,
				GroupDealIDs:         dealIDs,
				NonVerifiedDealSpace: nonVerifiedDealWeight,
				VerifiedDealSpace:    verifiedDealWeight,
				ActionType:           tx.TxType,
				TxTimestamp:          tx.TxTimestamp,
			})
		}

		return dealSpaceInfo, nil
	}

	// After NV17(mainnet) and NV16(calibration) VerifyDealsForActivation return changed to remove deal space and weight information,
	// we get the info from the ActivateDeals method
	maxVersion := tools.V17.NodeVersion()
	if eg.network == tools.CalibrationNetwork {
		maxVersion = tools.V16.NodeVersion()
	}
	/*
		From NV3(mainnet and calibration), to NV8(mainnet) and NV16(calibration) VerifyDealsForActivation params and return changed
		// - Array of sectors rather than just one
		// - Removed SectorStart (which is unknown at call time)
		type VerifyDealsForActivationParams struct {
			Sectors []SectorDeals
		}
		type SectorDeals struct {
			SectorExpiry abi.ChainEpoch
			DealIDs      []abi.DealID
		}
		type VerifyDealsForActivationReturn struct {
			Sectors []SectorWeights
		}

		type SectorWeights struct {
			DealSpace          uint64         // Total space in bytes of submitted deals.
			DealWeight         abi.DealWeight // Total space*time of submitted deals.
			VerifiedDealWeight abi.DealWeight // Total space*time of submitted verified deals.
		}
	*/
	if version.NodeVersion() > minVersion && version.NodeVersion() < maxVersion {
		// number of SectorDeals and SectorWeights will always be the same are they are processed in an all or nothing manner
		sectorDeals, err := common.GetSlice[map[string]interface{}](params, KeySectors, false)
		if err != nil {
			return nil, err
		}
		sectorWeights, err := common.GetSlice[map[string]interface{}](ret, KeySectors, false)
		if err != nil {
			return nil, err
		}

		for i := 0; i < len(sectorDeals); i++ {
			dealIDs, nonVerifiedDealWeight, verifiedDealWeight, err := eg.getCommonVerifyDealForActivationFields(sectorDeals[i], sectorWeights[i])
			if err != nil {
				return nil, err
			}

			for _, dealID := range dealIDs {
				dealSpaceInfo = append(dealSpaceInfo, &types.DealsSpaceInfo{
					ID:                   tools.BuildId(tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType, fmt.Sprint(dealID)),
					Height:               tx.Height,
					ActorAddress:         tx.TxFrom,
					TxCid:                tx.TxCid,
					DealID:               dealID,
					GroupDealIDs:         dealIDs,
					NonVerifiedDealSpace: nonVerifiedDealWeight,
					VerifiedDealSpace:    verifiedDealWeight,
					SpaceAsWeight:        true,
					ActionType:           tx.TxType,
					TxTimestamp:          tx.TxTimestamp,
				})
			}
		}
		return dealSpaceInfo, nil
	}

	return nil, nil
}

// parseActivateDeals handles ActivateDeals / BatchActivateDeals. exactParams is params decoded with
// json.Number, used only for the sector number.
func (eg *eventGenerator) parseActivateDeals(tx *types.Transaction, params, exactParams, ret map[string]interface{}) ([]*types.DealsActivations, []*types.DealsSpaceInfo, error) {
	dealActivations := []*types.DealsActivations{}
	dealSpaceInfo := []*types.DealsSpaceInfo{}

	parseDeals := func(params, ret map[string]interface{}, sectorNumber *uint64) error {
		dealIDs, sectorExpiry, err := eg.getActivationFields(params)
		if err != nil {
			return err
		}

		for _, dealID := range dealIDs {
			dealActivations = append(dealActivations, &types.DealsActivations{
				ID:           tools.BuildId(tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType, fmt.Sprint(dealID)),
				ActorAddress: tx.TxFrom,
				Height:       tx.Height,
				TxCid:        tx.TxCid,
				DealID:       dealID,
				SectorExpiry: sectorExpiry,
				SectorNumber: sectorNumber,
				ActionType:   tx.TxType,
				TxTimestamp:  tx.TxTimestamp,
			})
		}
		// Before NV18(mainnet) and NV17(calibration), ActivateDeals return is empty and we get the deal space info from VerifyDealsForActivation
		if len(ret) == 0 {
			return nil
		}
		// the dealIDs can be nil for a single activation
		if len(dealIDs) > 0 {
			nonVerifiedDealSpace, verifiedDealSpace, err := eg.getDealSpaceFields(ret)
			if err != nil {
				return err
			}
			for _, dealID := range dealIDs {
				dealSpaceInfo = append(dealSpaceInfo, &types.DealsSpaceInfo{
					ID:                   tools.BuildId(tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType, fmt.Sprint(dealID)),
					DealID:               dealID,
					Height:               tx.Height,
					ActorAddress:         tx.TxFrom,
					TxCid:                tx.TxCid,
					GroupDealIDs:         dealIDs,
					NonVerifiedDealSpace: nonVerifiedDealSpace,
					VerifiedDealSpace:    verifiedDealSpace,
					ActionType:           tx.TxType,
					TxTimestamp:          tx.TxTimestamp,
				})
			}
		}
		return nil
	}

	// The params are either flat {DealIDs, SectorExpiry} (ActivateDeals up to actors v11) or
	// {Sectors: [{SectorNumber, SectorType, SectorExpiry, DealIDs}], ComputeCid} (BatchActivateDeals,
	// actors v12 / NV21 on). Dispatch on the shape the actor parser produced rather than on the
	// network version: the version cutoff used before was off by one on mainnet (NV21 already has
	// the sectors shape), and the actor parser falls back to other versions' types when the
	// height's type fails to decode, so the shape is what the metadata actually holds.
	if _, ok := params[KeySectors]; !ok {
		if err := parseDeals(params, ret, nil); err != nil {
			return nil, nil, err
		}
		return dealActivations, dealSpaceInfo, nil
	}

	sectorDeals, err := common.GetSlice[map[string]interface{}](params, KeySectors, false)
	if err != nil {
		return nil, nil, err
	}
	exactSectorDeals, err := common.GetSlice[map[string]interface{}](exactParams, KeySectors, false)
	if err != nil {
		return nil, nil, err
	}
	if len(exactSectorDeals) != len(sectorDeals) {
		return nil, nil, fmt.Errorf("sector count mismatch: %d vs %d", len(exactSectorDeals), len(sectorDeals))
	}
	// activations can be nil if all activations failed
	activations, err := common.GetSlice[map[string]interface{}](ret, KeyActivations, true)
	if err != nil {
		return nil, nil, err
	}

	activationResults, err := common.GetItem[map[string]interface{}](ret, KeyActivationResults, true)
	if err != nil {
		return nil, nil, err
	}
	failedActivations := map[int64]bool{}
	if len(activationResults) > 0 {
		failCodes, err := common.GetSlice[map[string]interface{}](activationResults, KeyFailCodes, true)
		if err != nil {
			return nil, nil, err
		}
		for _, failCode := range failCodes {
			idx, err := common.GetInteger[int64](failCode, KeyIdx, false)
			if err != nil {
				return nil, nil, err
			}
			failedActivations[idx] = true
		}
	}

	// one activation per sector that did not fail, in order
	succeeded := 0
	for dealIdx := range sectorDeals {
		if !failedActivations[int64(dealIdx)] {
			succeeded++
		}
	}
	if len(activations) != succeeded {
		return nil, nil, fmt.Errorf("return has %d activations, expected %d", len(activations), succeeded)
	}

	// SectorDeals carries the sector number only from actors v13 (NV22) on
	// (builtin-actors v13.0.0 actors/market/src/types.rs). Actors v12 (NV21) has
	// {sector_type, sector_expiry, deal_ids}, but the Go types used to decode it have a
	// SectorNumber field, so the metadata shows its zero value. 0 is a valid sector number, so
	// it must not be recorded for v12.
	//#nosec G115
	hasSectorNumber := tools.VersionFromHeight(eg.network, int64(tx.Height)).NodeVersion() >= tools.V22.NodeVersion()

	activationIdx := 0
	for dealIdx := range sectorDeals {
		if failedActivations[int64(dealIdx)] {
			continue
		}
		var sectorNumber *uint64
		if hasSectorNumber {
			sectorNumber, err = getOptionalUint64(exactSectorDeals[dealIdx], KeySectorNumber)
			if err != nil {
				return nil, nil, err
			}
		}
		if err := parseDeals(sectorDeals[dealIdx], activations[activationIdx], sectorNumber); err != nil {
			return nil, nil, err
		}
		activationIdx++
	}
	return dealActivations, dealSpaceInfo, nil
}

func (eg *eventGenerator) getCommonVerifyDealForActivationFields(params, ret map[string]interface{}) (dealIDs []uint64, nonVerifiedDealWeight *big.Int, verifiedDealWeight *big.Int, err error) {
	// the DealIDs in the activation can be nil
	// see txcid: bafy2bzacecahrkpit4hgkudjhzhqu2tzvnatlsh5uh42hlzetyjbi45iclp3g on mainnet
	dealIDs, err = common.GetIntegerSlice[uint64](params, KeyDealIDs, true)
	if err != nil {
		return
	}

	nonVerifiedDealWeight, err = common.GetBigInt(ret, KeyDealWeight, false)
	if err != nil {
		return
	}

	verifiedDealWeight, err = common.GetBigInt(ret, KeyVerifiedDealWeight, false)
	if err != nil {
		return
	}

	return dealIDs, nonVerifiedDealWeight, verifiedDealWeight, nil
}

func (eg *eventGenerator) getActivationFields(params map[string]interface{}) (dealIDs []uint64, sectorExpiry int64, err error) {
	// the DealIDs in the activation can be nil
	// see txcid: bafy2bzacear4nqxca2qpt4p5sh3ic3qfnm4mnsutoj25kw2g3u3xkpbqdvpfe on calibration
	dealIDs, err = common.GetIntegerSlice[uint64](params, KeyDealIDs, true)
	if err != nil {
		return
	}
	sectorExpiry, err = common.GetInteger[int64](params, KeySectorExpiry, true)
	if err != nil {
		return
	}

	return dealIDs, sectorExpiry, nil
}

func (eg *eventGenerator) getDealSpaceFields(ret map[string]interface{}) (nonVerifiedDealSpace *big.Int, verifiedDealSpace *big.Int, err error) {
	nonVerifiedDealSpace, err = common.GetBigInt(ret, KeyNonVerifiedDealSpace, false)
	if err != nil {
		return
	}

	verifiedInfos, err := common.GetSlice[map[string]interface{}](ret, KeyVerifiedInfos, true)
	if err != nil {
		return
	}

	verifiedDealSpace = big.NewInt(0)
	for _, verifiedInfo := range verifiedInfos {
		var pieceSize uint64
		pieceSize, err = common.GetInteger[uint64](verifiedInfo, KeySize, false)
		if err != nil {
			return
		}

		verifiedDealSpace.Add(verifiedDealSpace, big.NewInt(0).SetUint64(pieceSize))
	}

	return nonVerifiedDealSpace, verifiedDealSpace, nil
}

func (eg *eventGenerator) isDealActivation(tx *types.Transaction) bool {
	switch tx.TxType {
	case parser.MethodActivateDeals,
		parser.MethodBatchActivateDeals,
		parser.MethodVerifyDealsForActivation,
		parser.MethodSectorContentChanged:
		return true
	}

	return false
}

// decodeExact decodes metadata keeping numbers as json.Number.
func decodeExact(metadata string) (map[string]interface{}, error) {
	decoder := json.NewDecoder(strings.NewReader(metadata))
	decoder.UseNumber()
	var out map[string]interface{}
	if err := decoder.Decode(&out); err != nil {
		return nil, err
	}
	return out, nil
}

// exactUint64 reads a json.Number (see decodeExact) as a uint64 without losing precision.
func exactUint64(value map[string]interface{}, key string) (uint64, error) {
	number, ok := value[key].(json.Number)
	if !ok {
		return 0, fmt.Errorf("key %s is not a number: %T", key, value[key])
	}
	v, err := strconv.ParseUint(number.String(), 10, 64)
	if err != nil {
		return 0, fmt.Errorf("key %s: %w", key, err)
	}
	return v, nil
}

// exactInt64 reads a json.Number (see decodeExact) as an int64 without losing precision.
func exactInt64(value map[string]interface{}, key string) (int64, error) {
	number, ok := value[key].(json.Number)
	if !ok {
		return 0, fmt.Errorf("key %s is not a number: %T", key, value[key])
	}
	v, err := strconv.ParseInt(number.String(), 10, 64)
	if err != nil {
		return 0, fmt.Errorf("key %s: %w", key, err)
	}
	return v, nil
}

// getOptionalUint64 returns nil when key is absent. value must come from decodeExact.
func getOptionalUint64(value map[string]interface{}, key string) (*uint64, error) {
	if value == nil || value[key] == nil {
		return nil, nil
	}
	v, err := exactUint64(value, key)
	if err != nil {
		return nil, err
	}
	return &v, nil
}
