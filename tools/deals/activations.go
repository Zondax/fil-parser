package deals

import (
	"context"
	"encoding/json"
	"fmt"

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
	ret, err := common.GetItem[map[string]interface{}](metadata, KeyReturn, false)
	if err != nil {
		return nil, nil, err
	}

	switch tx.TxType {
	case parser.MethodVerifyDealsForActivation:
		dealSpaceInfo, err = eg.parseVerifyDealsForActivation(tx, params, ret)
		if err != nil {
			return nil, nil, err
		}
	case parser.MethodActivateDeals, parser.MethodBatchActivateDeals:
		dealActivations, dealSpaceInfo, err = eg.parseActivateDeals(tx, params, ret)
		if err != nil {
			return nil, nil, err
		}
	}

	return dealActivations, dealSpaceInfo, nil
}

func (eg *eventGenerator) parseVerifyDealsForActivation(tx *types.Transaction, params, ret map[string]interface{}) ([]*types.DealsSpaceInfo, error) {
	version := tools.VersionFromHeight(eg.network, int64(tx.Height))
	dealSpaceInfo := []*types.DealsSpaceInfo{}

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

	if version.NodeVersion() > tools.V3.NodeVersion() && version.NodeVersion() <= tools.V8.NodeVersion() {
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

	if version.NodeVersion() < tools.V20.NodeVersion() {
		if err := parseDeals(params, ret); err != nil {
			return nil, nil, err
		}
		return dealActivations, dealSpaceInfo, nil
	}

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
			return nil, nil, fmt.Errorf("sectorDeals and activations have different lengths: %d != %d", len(sectorDeals), len(activations))
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

func (eg *eventGenerator) getCommonVerifyDealForActivationFields(params, ret map[string]interface{}) (dealIDs []uint64, nonVerifiedDealWeight uint64, verifiedDealWeight uint64, err error) {
	dealIDs, err = common.GetIntegerSlice[uint64](params, KeyDealIDs, false)
	if err != nil {
		return
	}

	nonVerifiedDealWeight, err = common.GetInteger[uint64](ret, KeyDealWeight, false)
	if err != nil {
		return
	}

	verifiedDealWeight, err = common.GetInteger[uint64](ret, KeyVerifiedDealWeight, false)
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

func (eg *eventGenerator) getDealSpaceFields(ret map[string]interface{}) (nonVerifiedDealSpace uint64, verifiedDealSpace uint64, err error) {
	nonVerifiedDealSpace, err = common.GetInteger[uint64](ret, KeyNonVerifiedDealSpace, false)
	if err != nil {
		return
	}

	verifiedInfos, err := common.GetSlice[map[string]interface{}](ret, KeyVerifiedInfos, false)
	if err != nil {
		return
	}

	for _, verifiedInfo := range verifiedInfos {
		var pieceSize uint64
		pieceSize, err = common.GetInteger[uint64](verifiedInfo, KeySize, false)
		if err != nil {
			return
		}

		verifiedDealSpace += pieceSize
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
