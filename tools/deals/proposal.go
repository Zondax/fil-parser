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
	KeyIDs                  = "IDs"
	KeyValidDeals           = "ValidDeals"
	KeyDealIDs              = "DealIDs"
	KeyClientSignature      = "ClientSignature"
	KeyData                 = "Data"
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
	KeyStoragePricePerEpoch = "StoragePricePerEpoch"
	KeyProviderCollateral   = "ProviderCollateral"
	KeyClientCollateral     = "ClientCollateral"
)

func (eg *eventGenerator) createDealsInfo(_ context.Context, tx *types.Transaction) ([]*types.DealsProposals, error) {
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

	dealsInfo, err := eg.parsePublishStorageDeals(tx, params, ret)
	if err != nil {
		return nil, fmt.Errorf("error creating events: %w", err)
	}

	return dealsInfo, nil
}

func (eg *eventGenerator) parsePublishStorageDeals(tx *types.Transaction, params, ret map[string]interface{}) ([]*types.DealsProposals, error) {
	dealsInfo := make([]*types.DealsProposals, 0)
	//#nosec G115
	version := tools.VersionFromHeight(eg.network, int64(tx.Height))

	validDealIndexToDealID := make(map[uint64]uint64)

	dealIDs, err := common.GetIntegerSlice[uint64](ret, KeyIDs, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing deal ids: %w", err)
	}

	// use the return to get the deal ids because the actor may drop invalid deals and return less ids than the params
	// From NV0 - NV13 the verified deals are in PublishStorageDealsReturn.IDs
	if version.NodeVersion() < tools.V14.NodeVersion() {
		for i, id := range dealIDs {
			// #nosec G115
			validDealIndexToDealID[uint64(i)] = id
		}
	} else {
		// From >= NV14 the verified deals are in PublishStorageDealsReturn.ValidDeals as a bitfield
		validDeals, err := common.GetIntegerSlice[int](ret, KeyValidDeals, false)
		if err != nil {
			return nil, fmt.Errorf("error parsing verified deal ids: %w", err)
		}
		validDealIndices, err := common.JsonEncodedBitfieldToIDs(validDeals)
		if err != nil {
			return nil, fmt.Errorf("error parsing verified deal ids: %w", err)
		}

		for i, idx := range validDealIndices {
			validDealIndexToDealID[idx] = dealIDs[i]
		}
	}

	deals, err := common.GetSlice[map[string]interface{}](params, KeyDeals, false)
	if err != nil {
		return nil, fmt.Errorf("error parsing deals: %w", err)
	}

	for idx, dealID := range validDealIndexToDealID {
		deal := deals[idx]

		clientSignature, err := common.GetItem[map[string]interface{}](deal, KeyClientSignature, false)
		if err != nil {
			return nil, fmt.Errorf("error parsing client signature: %w", err)
		}
		clientSignatureData, err := common.GetItem[string](clientSignature, KeyData, false)
		if err != nil {
			return nil, fmt.Errorf("error parsing client signature data: %w", err)
		}

		proposal, err := common.GetItem[map[string]interface{}](deal, KeyProposal, false)
		if err != nil {
			return nil, fmt.Errorf("error parsing proposal: %w", err)
		}

		pieceCID, err := common.GetCID(proposal, KeyPieceCID, false)
		if err != nil {
			return nil, fmt.Errorf("error parsing piece cid: %w", err)
		}

		pieceSize, err := common.GetInteger[uint64](proposal, KeyPieceSize, false)
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

		startEpoch, err := common.GetInteger[int64](proposal, KeyStartEpoch, false)
		if err != nil {
			return nil, fmt.Errorf("error parsing start epoch: %w", err)
		}

		endEpoch, err := common.GetInteger[int64](proposal, KeyEndEpoch, false)
		if err != nil {
			return nil, fmt.Errorf("error parsing end epoch: %w", err)
		}

		// encoded as a bigint string
		storagePricePerEpoch, err := common.GetBigInt(proposal, KeyStoragePricePerEpoch, false)
		if err != nil {
			return nil, fmt.Errorf("error parsing storage price per epoch: %w", err)
		}

		// encoded as a bigint string
		providerCollateral, err := common.GetBigInt(proposal, KeyProviderCollateral, false)
		if err != nil {
			return nil, fmt.Errorf("error parsing provider collateral: %w", err)
		}

		// encoded as a bigint string
		clientCollateral, err := common.GetBigInt(proposal, KeyClientCollateral, false)
		if err != nil {
			return nil, fmt.Errorf("error parsing client collateral: %w", err)
		}

		dealsInfo = append(dealsInfo, &types.DealsProposals{
			ID:                 tools.BuildId(tx.TxCid, tx.TxFrom, tx.TxTo, fmt.Sprint(tx.Height), tx.TxType, fmt.Sprint(dealID)),
			ActorAddress:       tx.TxFrom,
			Height:             tx.Height,
			DealID:             dealID,
			TxCid:              tx.TxCid,
			ClientSignature:    clientSignatureData,
			ProviderAddress:    providerAddress,
			ClientAddress:      clientAddress,
			PieceCid:           pieceCID.String(),
			PieceSize:          pieceSize,
			Verified:           verifiedDeal,
			Label:              label,
			StartEpoch:         startEpoch,
			EndEpoch:           endEpoch,
			PricePerEpoch:      storagePricePerEpoch.Uint64(),
			ProviderCollateral: providerCollateral,
			ClientCollateral:   clientCollateral,
			TxTimestamp:        tx.TxTimestamp,
		})

	}

	return dealsInfo, nil
}
