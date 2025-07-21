package verifreg

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/tools/common"
	"github.com/zondax/fil-parser/types"
)

const (
	KeyAddress           = "Address"
	KeyAllowance         = "Allowance"
	KeyParams            = "Params"
	KeyVerifiedClient    = "VerifiedClient"
	KeyDataCapRemoved    = "DataCapRemoved"
	KeyVerifierRequest1  = "VerifierRequest1"
	KeyVerifierRequest2  = "VerifierRequest2"
	KeyVerifier          = "Verifier"
	KeyVerifierSignature = "VerifierSignature"
	KeyData              = "Data"
)

func getAddressAllowance(value map[string]interface{}) (string, *big.Int, error) {
	params, err := common.GetItem[map[string]interface{}](value, parser.ParamsKey, false)
	if err != nil {
		return "", nil, fmt.Errorf("error parsing params: %w", err)
	}

	addr, err := common.GetItem[string](params, KeyAddress, false)
	if err != nil {
		return "", nil, fmt.Errorf("error parsing address: %w", err)
	}

	allowance, err := common.GetBigInt(params, KeyAllowance, false)
	if err != nil {
		return "", nil, fmt.Errorf("error parsing allowance: %w", err)
	}

	return addr, allowance, nil
}

func parseRemoveVerifier(metadata map[string]interface{}) (string, error) {
	addr, err := common.GetItem[string](metadata, KeyParams, false)
	if err != nil {
		return "", fmt.Errorf("error getting params from metadata: %w", err)
	}

	return addr, nil
}

// returns clientToRemove, verifiers, dataCapAmountToRemove, error
func parseRemoveVerifiedClient(metadata map[string]interface{}, network string, height int64) (string, []string, *big.Int, error) {
	params, err := common.GetItem[map[string]interface{}](metadata, parser.ParamsKey, false)
	if err != nil {
		return "", nil, nil, fmt.Errorf("error getting params from metadata: %w", err)
	}

	returnValue, err := common.GetItem[map[string]interface{}](metadata, parser.ReturnKey, false)
	if err != nil {
		return "", nil, nil, fmt.Errorf("error getting params from metadata: %w", err)
	}

	clientToRemove, err := common.GetItem[string](returnValue, KeyVerifiedClient, false)
	if err != nil {
		return "", nil, nil, fmt.Errorf("error getting params from metadata: %w", err)
	}

	dataCapAmountToRemove, err := common.GetBigInt(returnValue, KeyDataCapRemoved, false)
	if err != nil {
		return "", nil, nil, fmt.Errorf("error getting params from metadata: %w", err)
	}

	verifiers := []string{}
	verifierRequests := []string{KeyVerifierRequest1, KeyVerifierRequest2}
	for _, verifierRequest := range verifierRequests {
		verifierAddress, _, err := getVerifierFromVerifierRequest(params, verifierRequest)
		if err != nil {
			return "", nil, nil, fmt.Errorf("error getting params from metadata: %w", err)
		}
		verifiers = append(verifiers, verifierAddress)
	}

	return clientToRemove, verifiers, dataCapAmountToRemove, nil
}

// returns the verifierAddress, verifierSignature, error
func getVerifierFromVerifierRequest(value map[string]interface{}, key string) (string, string, error) {
	verifierRequest, err := common.GetItem[map[string]interface{}](value, key, false)
	if err != nil {
		return "", "", fmt.Errorf("error getting params from metadata: %w", err)
	}

	verifierAddress, err := common.GetItem[string](verifierRequest, KeyVerifier, false)
	if err != nil {
		return "", "", fmt.Errorf("error getting params from metadata: %w", err)
	}

	verifierSignature, err := common.GetItem[map[string]interface{}](verifierRequest, KeyVerifierSignature, false)
	if err != nil {
		return "", "", fmt.Errorf("error getting params from metadata: %w", err)
	}

	verifierSignatureData, err := common.GetItem[string](verifierSignature, KeyData, false)
	if err != nil {
		return "", "", fmt.Errorf("error getting params from metadata: %w", err)
	}

	return verifierAddress, verifierSignatureData, nil
}

func parserUniversalReceiverHook(tx *types.Transaction, tipsetCid string) (string, []*types.VerifregDeal, error) {
	// Parse the FRC46 transaction metadata
	params, returnData, err := ParseFRC46TransactionMetadata(tx.TxMetadata)
	if err != nil {
		return "", nil, fmt.Errorf("error parsing FRC46 transaction metadata: %w", err)
	}

	// TODO: what happens when fail deal
	if len(params.OperatorData.Allocations) != len(returnData.NewAllocations) {
		return "", nil, errors.New("invalid number of allocations")
	}

	allocations := make([]AllocationDataWithDealID, len(params.OperatorData.Allocations))
	deals := make([]*types.VerifregDeal, len(returnData.NewAllocations))
	for i := range params.OperatorData.Allocations {
		dealId := fmt.Sprintf("%d", returnData.NewAllocations[i])

		allocations[i].AllocationData = params.OperatorData.Allocations[i]
		allocations[i].DealID = dealId

		allocBytes, err := json.Marshal(allocations[i])
		if err != nil {
			return "", nil, fmt.Errorf("error marshalling allocation: %w", err)
		}

		deals[i] = &types.VerifregDeal{
			ID:          tools.BuildId(dealId, tx.TxCid, fmt.Sprint(tx.Height)),
			DeadID:      dealId,
			TxCid:       tx.TxCid,
			Height:      tx.Height,
			Value:       string(allocBytes),
			TxTimestamp: tx.TxTimestamp,
		}
	}

	clientValue, err := json.Marshal(allocations)
	if err != nil {
		return "", nil, fmt.Errorf("error marshalling allocations: %w", err)
	}

	return string(clientValue), deals, nil
}

// ParseFRC46TransactionMetadata parses FRC46 token transaction metadata
func ParseFRC46TransactionMetadata(metadata string) (*FRC46TransactionParams, *FRC46TransactionReturn, error) {
	var txMetadata FRC46TransactionMetadata
	err := json.Unmarshal([]byte(metadata), &txMetadata)
	if err != nil {
		return nil, nil, fmt.Errorf("error unmarshalling FRC46 transaction metadata: %w", err)
	}

	return &txMetadata.Params, &txMetadata.Return, nil
}
