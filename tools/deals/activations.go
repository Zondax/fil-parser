package deals

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/tools/common"
	"github.com/zondax/fil-parser/types"
)

const (
	KeySectors              = "Sectors"
	KeyActivations          = "Activations"
	KeyDealWeight           = "DealWeight"
	KeyVerifiedDealWeight   = "VerifiedDealWeight"
	KeyNonVerifiedDealSpace = "NonVerifiedDealSpace"
	KeyVerifiedInfos        = "VerifiedInfos"
	KeySize                 = "Size"
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

	params, err := common.GetItem[map[string]interface{}](metadata, KeyParams, false)
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
		dealActivations, dealSpaceInfo, err = eg.parseActivateDeals(tx, params, ret)
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
	if version.NodeVersion() < tools.V3.NodeVersion() {
		dealIDs, nonVerifiedDealWeight, verifiedDealWeight, err := eg.getCommonVerifyDealForActivationFields(params, ret)
		if err != nil {
			return nil, err
		}

		dealSpaceInfo = append(dealSpaceInfo, &types.DealsSpaceInfo{
			ID:                   tools.BuildId(tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType, fmt.Sprint(dealIDs)),
			Height:               tx.Height,
			TxCid:                tx.TxCid,
			DealIDs:              dealIDs,
			NonVerifiedDealSpace: nonVerifiedDealWeight,
			VerifiedDealSpace:    verifiedDealWeight,
			ActionType:           tx.TxType,
			TxTimestamp:          tx.TxTimestamp,
		})

		return dealSpaceInfo, nil
	}

	maxVersion := tools.V8.NodeVersion()
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

		After NV8(mainnet) and NV16(calibration) VerifyDealsForActivation return changed to remove deal space and weight information,
		we get the info from the ActivateDeals method
	*/
	if version.NodeVersion() > tools.V3.NodeVersion() && version.NodeVersion() <= maxVersion {
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

			dealSpaceInfo = append(dealSpaceInfo, &types.DealsSpaceInfo{
				ID:                   tools.BuildId(tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType, fmt.Sprint(dealIDs[i])),
				Height:               tx.Height,
				TxCid:                tx.TxCid,
				DealIDs:              dealIDs,
				NonVerifiedDealSpace: nonVerifiedDealWeight,
				VerifiedDealSpace:    verifiedDealWeight,
				SpaceAsWeight:        true,
				ActionType:           tx.TxType,
				TxTimestamp:          tx.TxTimestamp,
			})
		}
		return dealSpaceInfo, nil
	}

	return nil, nil
}

func (eg *eventGenerator) parseActivateDeals(tx *types.Transaction, params, ret map[string]interface{}) ([]*types.DealsActivations, []*types.DealsSpaceInfo, error) {
	//#nosec G115
	version := tools.VersionFromHeight(eg.network, int64(tx.Height))
	dealActivations := []*types.DealsActivations{}
	dealSpaceInfo := []*types.DealsSpaceInfo{}

	parseDeals := func(params, ret map[string]interface{}) error {
		dealIDs, sectorExpiry, err := eg.getActivationFields(params)
		if err != nil {
			return err
		}

		for _, dealID := range dealIDs {
			dealActivations = append(dealActivations, &types.DealsActivations{
				ID:           tools.BuildId(tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType, fmt.Sprint(dealID)),
				Height:       tx.Height,
				TxCid:        tx.TxCid,
				DealID:       dealID,
				SectorExpiry: sectorExpiry,
				ActionType:   tx.TxType,
				TxTimestamp:  tx.TxTimestamp,
			})
		}
		// Before NV17(<=NV16), ActivateDeals return is empty and we get the deal space info from VerifyDealsForActivation
		if ret == nil {
			return nil
		}
		nonVerifiedDealSpace, verifiedDealSpace, err := eg.getDealSpaceFields(ret)
		if err != nil {
			return err
		}
		dealSpaceInfo = append(dealSpaceInfo, &types.DealsSpaceInfo{
			ID:                   tools.BuildId(tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType, fmt.Sprint(dealIDs)),
			Height:               tx.Height,
			TxCid:                tx.TxCid,
			DealIDs:              dealIDs,
			NonVerifiedDealSpace: nonVerifiedDealSpace,
			VerifiedDealSpace:    verifiedDealSpace,
			ActionType:           tx.TxType,
			TxTimestamp:          tx.TxTimestamp,
		})
		return nil
	}

	// Before NV20, ActivateDeals activates multiple deals in 1 sector
	if version.NodeVersion() <= tools.V20.NodeVersion() {
		if err := parseDeals(params, ret); err != nil {
			return nil, nil, err
		}
		return dealActivations, dealSpaceInfo, nil
	}

	// After NV20, ActivateDeals activates multiple deals in multiple sectors
	if version.NodeVersion() > tools.V20.NodeVersion() {
		sectorDeals, err := common.GetSlice[map[string]interface{}](params, KeySectors, false)
		if err != nil {
			return nil, nil, err
		}
		activations, err := common.GetSlice[map[string]interface{}](ret, KeyActivations, false)
		if err != nil {
			return nil, nil, err
		}
		if len(sectorDeals) != len(activations) {
			return nil, nil, fmt.Errorf("sectorDeals and activations have different lengths: sectorDeals(%d) != activations(%d)", len(sectorDeals), len(activations))
		}

		for i := range sectorDeals {
			if err := parseDeals(sectorDeals[i], activations[i]); err != nil {
				return nil, nil, err
			}
		}
		return dealActivations, dealSpaceInfo, nil
	}

	return nil, nil, nil
}

func (eg *eventGenerator) getCommonVerifyDealForActivationFields(params, ret map[string]interface{}) (dealIDs []uint64, nonVerifiedDealWeight *big.Int, verifiedDealWeight *big.Int, err error) {
	dealIDs, err = common.GetIntegerSlice[uint64](params, KeyDealIDs, false)
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
	dealIDs, err = common.GetIntegerSlice[uint64](params, KeyDealIDs, false)
	if err != nil {
		return
	}
	sectorExpiry, err = common.GetInteger[int64](params, KeySectorExpiry, false)
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
		parser.MethodVerifyDealsForActivation:
		return true
	}

	return false
}
