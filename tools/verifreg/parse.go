package verifreg

import (
	"encoding/json"
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
)

// Type-safe structs for RemoveVerifier metadata
type VerifierSignature struct {
	Type int    `json:"Type"`
	Data string `json:"Data"`
}

type VerifierRequest struct {
	Verifier          string            `json:"Verifier"`
	VerifierSignature VerifierSignature `json:"VerifierSignature"`
}

type RemoveVerifierParams struct {
	VerifiedClientToRemove string          `json:"VerifiedClientToRemove"`
	DataCapAmountToRemove  string          `json:"DataCapAmountToRemove"`
	VerifierRequest1       VerifierRequest `json:"VerifierRequest1"`
	VerifierRequest2       VerifierRequest `json:"VerifierRequest2"`
}

type RemoveVerifierMetadata struct {
	MethodNum string               `json:"MethodNum"`
	Params    RemoveVerifierParams `json:"Params"`
}

// FRC46 Token transaction structures
type AllocationData struct {
	Provider   int64             `json:"Provider"`
	Data       map[string]string `json:"Data"`
	Size       int64             `json:"Size"`
	TermMin    int64             `json:"TermMin"`
	TermMax    int64             `json:"TermMax"`
	Expiration int64             `json:"Expiration"`
}

type OperatorData struct {
	Allocations []AllocationData `json:"Allocations"`
	Extensions  interface{}      `json:"Extensions"`
}

type FRC46TransactionParams struct {
	Amount       string       `json:"amount"`
	From         string       `json:"from"`
	Operator     string       `json:"operator"`
	OperatorData OperatorData `json:"operator_data"`
	To           string       `json:"to"`
	TokenData    string       `json:"token_data"`
	Type         string       `json:"type"`
}

type AllocationResults struct {
	SuccessCount int         `json:"SuccessCount"`
	FailCodes    interface{} `json:"FailCodes"`
}

type ExtensionResults struct {
	SuccessCount int         `json:"SuccessCount"`
	FailCodes    interface{} `json:"FailCodes"`
}

type FRC46TransactionReturn struct {
	AllocationResults AllocationResults `json:"AllocationResults"`
	ExtensionResults  ExtensionResults  `json:"ExtensionResults"`
	NewAllocations    []int64           `json:"NewAllocations"`
}

type FRC46TransactionMetadata struct {
	MethodNum string                 `json:"MethodNum"`
	Params    FRC46TransactionParams `json:"Params"`
	Return    FRC46TransactionReturn `json:"Return"`
}

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

func parserUniversalReceiverHook(metadata string) (string, error) {
	// Parse the FRC46 transaction metadata
	params, returnData, err := ParseFRC46TransactionMetadata(metadata)
	if err != nil {
		return "", fmt.Errorf("error parsing FRC46 transaction metadata: %w", err)
	}

	// You can now access the parsed data
	// For example, accessing allocation information
	if len(params.OperatorData.Allocations) > 0 {
		allocation := params.OperatorData.Allocations[0]
		fmt.Printf("Provider: %d, Size: %d, Expiration: %d\n",
			allocation.Provider, allocation.Size, allocation.Expiration)
	}

	// Access return data
	if len(returnData.NewAllocations) > 0 {
		fmt.Printf("New Allocations: %v\n", returnData.NewAllocations)
		fmt.Printf("Success Count: %d\n", returnData.AllocationResults.SuccessCount)
	}

	return "", nil
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

// ExampleParseFRC46Transaction demonstrates how to parse the FRC46 transaction metadata
func ExampleParseFRC46Transaction() {
	jsonString := `{"MethodNum":"3726118371","Params":{"amount":"34359738368000000000000000000","from":"f03201686","operator":"f03201686","operator_data":{"Allocations":[{"Provider":3175111,"Data":{"/":"baga6ea4seaqmo4bi4luj3fdif3unw75zwpnek67ay4jafo4754kgypp4o7cf2fi"},"Size":34359738368,"TermMin":518400,"TermMax":5256000,"Expiration":4473813}],"Extensions":null},"to":"f06","token_data":"","type":"frc46_token"},"Return":{"AllocationResults":{"SuccessCount":1,"FailCodes":null},"ExtensionResults":{"SuccessCount":0,"FailCodes":null},"NewAllocations":[81930453]}}`

	params, returnData, err := ParseFRC46TransactionMetadata(jsonString)
	if err != nil {
		fmt.Printf("Error parsing: %v\n", err)
		return
	}

	// Access parameters
	fmt.Printf("Method Number: %s\n", "3726118371") // You can extract this separately if needed
	fmt.Printf("Amount: %s\n", params.Amount)
	fmt.Printf("From: %s\n", params.From)
	fmt.Printf("To: %s\n", params.To)
	fmt.Printf("Operator: %s\n", params.Operator)
	fmt.Printf("Type: %s\n", params.Type)

	// Access allocation data
	if len(params.OperatorData.Allocations) > 0 {
		allocation := params.OperatorData.Allocations[0]
		fmt.Printf("Allocation Provider: %d\n", allocation.Provider)
		fmt.Printf("Allocation Size: %d\n", allocation.Size)
		fmt.Printf("Term Min: %d\n", allocation.TermMin)
		fmt.Printf("Term Max: %d\n", allocation.TermMax)
		fmt.Printf("Expiration: %d\n", allocation.Expiration)
		fmt.Printf("Data CID: %s\n", allocation.Data["/"])
	}

	// Access return data
	fmt.Printf("Allocation Success Count: %d\n", returnData.AllocationResults.SuccessCount)
	fmt.Printf("Extension Success Count: %d\n", returnData.ExtensionResults.SuccessCount)
	if len(returnData.NewAllocations) > 0 {
		fmt.Printf("New Allocations: %v\n", returnData.NewAllocations)
	}
}
