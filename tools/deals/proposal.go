package deals

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/tools/common"
	"github.com/zondax/fil-parser/types"
)

const (
	KeyDeals                = "Deals"
	KeyDealIDs              = "DealIDs"
	KeyClientSignature      = "ClientSignature"
	KeyProposal             = "Proposal"
	KeyProvider             = "Provider"
	KeySectorExpiry         = "SectorExpiry"
	KeyPieceCID             = "PieceCID"
	KeyPieceSize            = "PieceSize"
	KeyVerifiedDeal         = "VerifiedDeal"
	KeyClient               = "Client"
	KeyLabel                = "Label"
	KeyStartEpoch           = "StartEpoch"
	KeyEndEpoch             = "EndEpoch"
	KeyParams               = "Params"
	KeyReturn               = "Return"
	KeyStoragePricePerEpoch = "PricePerEpoch"
	KeyProviderCollateral   = "ProviderCollateral"
	KeyClientCollateral     = "ClientCollateral"
)

func (eg *eventGenerator) createDealsInfo(_ context.Context, tx *types.Transaction) ([]*types.DealsInfo, error) {
	var value map[string]interface{}
	err := json.Unmarshal([]byte(tx.TxMetadata), &value)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling tx metadata: %w", err)
	}

	params, err := common.GetItem[map[string]interface{}](value, KeyParams, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing params: %w", err)
	}
	ret, err := common.GetItem[map[string]interface{}](value, KeyReturn, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing ret: %w", err)
	}

	dealsInfo, err := parsePublishStorageDeals(tx, params, ret)
	if err != nil {
		return nil, fmt.Errorf("error creating events: %w", err)
	}

	return dealsInfo, nil
}

func parsePublishStorageDeals(tx *types.Transaction, params, ret map[string]interface{}) ([]*types.DealsInfo, error) {
	dealsInfo := make([]*types.DealsInfo, 0)
	dealIDs, err := common.GetIntegerSlice[uint64](ret, KeyDealIDs, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing seal proof: %w", err)
	}

	deals, err := common.GetSlice[map[string]interface{}](params, KeyDeals, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing seal proof: %w", err)
	}

	for idx, deal := range deals {
		clientSignature, err := common.GetItem[string](ret, KeyClientSignature, false)
		if err != nil {
			return nil, fmt.Errorf("error parsing client signature: %w", err)
		}
		proposal, err := common.GetItem[map[string]interface{}](deal, KeyProposal, false)
		if err != nil {
			return nil, fmt.Errorf("error parsing proposal: %w", err)
		}

		pieceCID, err := common.GetItem[string](proposal, KeyPieceCID, false)
		if err != nil {
			return nil, fmt.Errorf("error parsing piece cid: %w", err)
		}

		pieceSize, err := common.GetItem[uint64](proposal, KeyPieceSize, false)
		if err != nil {
			return nil, fmt.Errorf("error parsing piece size: %w", err)
		}

		verifiedDeal, err := common.GetItem[bool](proposal, KeyVerifiedDeal, false)
		if err != nil {
			return nil, fmt.Errorf("error parsing verified deal: %w", err)
		}

		clientAddress, err := common.GetItem[string](proposal, KeyClient, false)
		if err != nil {
			return nil, fmt.Errorf("error parsing client address: %w", err)
		}

		providerAddress, err := common.GetItem[string](proposal, KeyProvider, false)
		if err != nil {
			return nil, fmt.Errorf("error parsing provider address: %w", err)
		}

		label, err := common.GetItem[string](proposal, KeyLabel, false)
		if err != nil {
			return nil, fmt.Errorf("error parsing label: %w", err)
		}

		startEpoch, err := common.GetItem[int64](proposal, KeyStartEpoch, false)
		if err != nil {
			return nil, fmt.Errorf("error parsing start epoch: %w", err)
		}

		endEpoch, err := common.GetItem[int64](proposal, KeyEndEpoch, false)
		if err != nil {
			return nil, fmt.Errorf("error parsing end epoch: %w", err)
		}

		storagePricePerEpoch, err := common.GetItem[uint64](proposal, KeyStoragePricePerEpoch, false)
		if err != nil {
			return nil, fmt.Errorf("error parsing storage price per epoch: %w", err)
		}

		providerCollateral, err := common.GetItem[uint64](proposal, KeyProviderCollateral, false)
		if err != nil {
			return nil, fmt.Errorf("error parsing provider collateral: %w", err)
		}

		clientCollateral, err := common.GetItem[uint64](proposal, KeyClientCollateral, false)
		if err != nil {
			return nil, fmt.Errorf("error parsing client collateral: %w", err)
		}

		dealsInfo = append(dealsInfo, &types.DealsInfo{
			ID:                 tools.BuildID(tx.Height, dealIDs[0]),
			Height:             tx.Height,
			DealID:             dealIDs[idx],
			TxCid:              tx.TxCid,
			ClientSignature:    clientSignature,
			ProviderAddress:    providerAddress,
			ClientAddress:      clientAddress,
			PieceCid:           pieceCID,
			PieceSize:          pieceSize,
			Verified:           verifiedDeal,
			Label:              label,
			StartEpoch:         startEpoch,
			EndEpoch:           endEpoch,
			PricePerEpoch:      storagePricePerEpoch,
			ProviderCollateral: providerCollateral,
			ClientCollateral:   clientCollateral,
			TxTimestamp:        tx.TxTimestamp,
		})

	}

	return dealsInfo, nil
}
