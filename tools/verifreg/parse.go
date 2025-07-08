package verifreg

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	addr "github.com/filecoin-project/go-address"
	verifregv10 "github.com/filecoin-project/go-state-types/builtin/v10/verifreg"
	verifregv11 "github.com/filecoin-project/go-state-types/builtin/v11/verifreg"
	verifregv12 "github.com/filecoin-project/go-state-types/builtin/v12/verifreg"
	verifregv13 "github.com/filecoin-project/go-state-types/builtin/v13/verifreg"
	verifregv14 "github.com/filecoin-project/go-state-types/builtin/v14/verifreg"
	verifregv15 "github.com/filecoin-project/go-state-types/builtin/v15/verifreg"
	verifregv16 "github.com/filecoin-project/go-state-types/builtin/v16/verifreg"
	verifregv8 "github.com/filecoin-project/go-state-types/builtin/v8/verifreg"
	verifregv9 "github.com/filecoin-project/go-state-types/builtin/v9/verifreg"
	verifreg "github.com/zondax/fil-parser/actors/v2/verifiedRegistry"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
)

func getMetadataByKey(metadata, key string) (string, error) {
	var value map[string]interface{}
	err := json.Unmarshal([]byte(metadata), &value)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling tx metadata: %w", err)
	}

	params, ok := value[key].(string)
	if ok {
		return params, nil
	}

	// If params is a map[string]interface{}, marshal it back to JSON string
	if paramsMap, ok := value[key].(map[string]interface{}); ok {
		paramsBytes, err := json.Marshal(paramsMap)
		if err != nil {
			return "", fmt.Errorf("error marshalling params map: %w", err)
		}
		return string(paramsBytes), nil
	}

	return "", fmt.Errorf("params is neither string nor map[string]interface{}")
}

func getMetadataParams(metadata string) (string, error) {
	return getMetadataByKey(metadata, parser.ParamsKey)
}

func getMetadataReturn(metadata string) (string, error) {
	return getMetadataByKey(metadata, parser.ReturnKey)
}

// TODO: this parsing funcs doesn't make much sense. It should be done through
func getAddressAllowance(metadata string) (string, uint64, error) {
	var value map[string]interface{}
	err := json.Unmarshal([]byte(metadata), &value)
	if err != nil {
		return "", 0, fmt.Errorf("error unmarshalling tx metadata: %w", err)
	}

	params, ok := value[parser.ParamsKey].(map[string]interface{})
	if !ok {
		return "", 0, fmt.Errorf("error parsing params: %w", err)
	}

	addr, ok := params["Address"].(string)
	if !ok {
		return "", 0, fmt.Errorf("error parsing address: %w", err)
	}

	allowance, ok := params["Allowance"].(string)
	if !ok {
		return "", 0, fmt.Errorf("error parsing allowance: %w", err)
	}

	intAllowance, err := strconv.ParseUint(allowance, 10, 64)
	if err != nil {
		return "", 0, fmt.Errorf("error parsing allowance string '%s': %w", allowance, err)
	}

	return addr, intAllowance, nil
}

func parseRemoveVerifier(metadata string) (string, error) {
	// Parse the JSON metadata to extract the params, in the case of removeVerifier is just the address
	addr, err := getMetadataParams(metadata)
	if err != nil {
		return "", fmt.Errorf("error getting params from metadata: %w", err)
	}

	return addr, nil
}

func parseRemoveVerifiedClient(metadata, network string, height int64) (string, string, uint64, error) {
	// Parse the JSON metadata to extract the params
	paramsStr, err := getMetadataParams(metadata)
	if err != nil {
		return "", "", 0, fmt.Errorf("error getting params from metadata: %w", err)
	}

	// Get the concrete type constructor based on network version
	fn, ok := verifreg.VerifregTypes[parser.MethodRemoveVerifiedClientDataCap][tools.VersionFromHeight(network, height).String()]
	if !ok {
		return "", "", 0, fmt.Errorf("could not get verified client data")
	}

	// Create an instance of the CBORUnmarshaler interface
	params := fn()

	err = json.Unmarshal([]byte(paramsStr), params)
	if err != nil {
		return "", "", 0, fmt.Errorf("error unmarshalling into RemoveDataCapParams: %w", err)
	}

	// Get the concrete type through type assertion, then unmarshal JSON into it
	switch v := params.(type) {
	case *addr.Address:
		return v.String(), "", 0, nil
	case *verifregv16.RemoveDataCapParams:
		return v.VerifiedClientToRemove.String(),
			v.VerifierRequest1.Verifier.String(),
			v.DataCapAmountToRemove.Int.Uint64(),
			nil
	case *verifregv15.RemoveDataCapParams:
		return v.VerifiedClientToRemove.String(),
			v.VerifierRequest1.Verifier.String(),
			v.DataCapAmountToRemove.Int.Uint64(),
			nil
	case *verifregv14.RemoveDataCapParams:
		return v.VerifiedClientToRemove.String(),
			v.VerifierRequest1.Verifier.String(),
			v.DataCapAmountToRemove.Int.Uint64(),
			nil
	case *verifregv13.RemoveDataCapParams:
		return v.VerifiedClientToRemove.String(),
			fmt.Sprintf(`{"VerifierRequest1": %s, "VerifierRequest2": %s}`, v.VerifierRequest1.Verifier, v.VerifierRequest2.Verifier),
			v.DataCapAmountToRemove.Int.Uint64(),
			nil
	case *verifregv12.RemoveDataCapParams:
		return v.VerifiedClientToRemove.String(),
			v.VerifierRequest1.Verifier.String(),
			v.DataCapAmountToRemove.Int.Uint64(),
			nil
	case *verifregv11.RemoveDataCapParams:
		return v.VerifiedClientToRemove.String(),
			v.VerifierRequest1.Verifier.String(),
			v.DataCapAmountToRemove.Int.Uint64(),
			nil
	case *verifregv10.RemoveDataCapParams:
		return v.VerifiedClientToRemove.String(),
			v.VerifierRequest1.Verifier.String(),
			v.DataCapAmountToRemove.Int.Uint64(),
			nil
	case *verifregv9.RemoveDataCapParams:
		return v.VerifiedClientToRemove.String(),
			v.VerifierRequest1.Verifier.String(),
			v.DataCapAmountToRemove.Int.Uint64(),
			nil
	case *verifregv8.RemoveDataCapParams:
		return v.VerifiedClientToRemove.String(),
			v.VerifierRequest1.Verifier.String(),
			v.DataCapAmountToRemove.Int.Uint64(),
			nil
	}

	return "", "", 0, fmt.Errorf("unsupported concrete type: %T", params)
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
