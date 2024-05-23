package actors

import (
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"fmt"
	multisig2 "github.com/filecoin-project/go-state-types/builtin/v14/multisig"
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/parser"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestActorParser_approve(t *testing.T) {
	p := getActorParser()
	tests := []struct {
		name   string
		method string
	}{
		{
			name:   "Approve",
			method: parser.MethodApprove,
		},
		{
			name:   "Approve",
			method: parser.MethodApproveExported,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, rawReturn, err := getParamsAndReturn(manifest.MultisigKey, tt.method)
			require.NoError(t, err)
			require.NotNil(t, rawParams)
			require.NotNil(t, rawReturn)
			msg, err := deserializeMessage(manifest.MultisigKey, tt.method)
			require.NoError(t, err)
			require.NotNil(t, msg)
			tipSet, err := deserializeTipset(manifest.MultisigKey, tt.method)
			require.NoError(t, err)

			got, err := p.approve(msg, rawReturn, int64(tipSet.Height()), tipSet.Key())
			require.NoError(t, err)
			require.NotNil(t, got)
		})
	}
}

func TestActorParser_multisigWithParamsAndReturn(t *testing.T) {
	p := getActorParser()
	tests := []struct {
		name   string
		txType string
		f      func([]byte, []byte) (map[string]interface{}, error)
	}{
		{
			name:   "Propose",
			txType: parser.MethodPropose,
			f:      p.propose,
		},
		{
			name:   "Propose Exported",
			txType: parser.MethodProposeExported,
			f:      p.propose,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, rawReturn, err := getParamsAndReturn(manifest.MultisigKey, tt.txType)
			require.NoError(t, err)
			require.NotNil(t, rawParams)
			require.NotNil(t, rawReturn)

			got, err := tt.f(rawParams, rawReturn)
			require.NoError(t, err)
			require.NotNil(t, got)
			require.Contains(t, got, parser.ParamsKey, "Params could no be found in metadata")
			require.NotNil(t, got[parser.ParamsKey])
			require.Contains(t, got, parser.ReturnKey, "Return could no be found in metadata")
			require.NotNil(t, got[parser.ReturnKey])
		})
	}
}

func TestActorParser_multiSigParams(t *testing.T) {
	p := getActorParser()
	tests := []struct {
		name   string
		txType string
		f      func(msg *parser.LotusMessage, height int64, key filTypes.TipSetKey) (map[string]interface{}, error)
	}{
		{
			name:   "Add Signer",
			txType: parser.MethodAddSigner,
			f:      p.msigParams,
		},
		{
			name:   "Add Signer Exported",
			txType: parser.MethodAddSignerExported,
			f:      p.msigParams,
		},
		{
			name:   "Remove Signer",
			txType: parser.MethodRemoveSigner,
			f:      p.removeSigner,
		},
		{
			name:   "Remove Signer Exported",
			txType: parser.MethodRemoveSignerExported,
			f:      p.removeSigner,
		},
		{
			name:   "Cancel",
			txType: parser.MethodCancel,
			f:      p.cancel,
		},
		/* // TODO: will fail until https://github.com/Zondax/rosetta-filecoin-lib/pull/109 is merged
		{
			name:   "Cancel Exported",
			txType: parser.MethodCancelExported,
			f:      p.cancel,
		},
		*/
		{
			name:   "Swap Signer",
			txType: parser.MethodSwapSigner,
			f:      p.cancel,
		},
		{
			name:   "Swap Signer Exported",
			txType: parser.MethodSwapSignerExported,
			f:      p.cancel,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg, err := deserializeMessage(manifest.MultisigKey, tt.txType)
			require.NoError(t, err)
			require.NotNil(t, msg)

			tipset, err := deserializeTipset(manifest.MultisigKey, tt.txType)
			require.NoError(t, err)

			got, err := tt.f(msg, int64(tipset.Height()), tipset.Key())
			require.NoError(t, err)
			require.NotNil(t, got)
		})
	}
}

func TestActorParser_multisigWithParamsOrReturn(t *testing.T) {
	p := getActorParser()
	tests := []struct {
		name   string
		txType string
		f      func([]byte) (map[string]interface{}, error)
		key    string
	}{
		{
			name:   "Constructor",
			txType: parser.MethodConstructor,
			f:      p.msigConstructor,
			key:    parser.ParamsKey,
		},
		{
			name:   "Change Num Approvals Threshold",
			txType: parser.MethodChangeNumApprovalsThreshold,
			f:      p.changeNumApprovalsThreshold,
			key:    parser.ParamsKey,
		},
		{
			name:   "Change Num Approvals Threshold",
			txType: parser.MethodChangeNumApprovalsThresholdExported,
			f:      p.changeNumApprovalsThreshold,
			key:    parser.ParamsKey,
		},
		{
			name:   "Lock Balance",
			txType: parser.MethodLockBalance,
			f:      p.lockBalance,
			key:    parser.ParamsKey,
		},
		{
			name:   "Lock Balance Exported",
			txType: parser.MethodLockBalanceExported,
			f:      p.lockBalance,
			key:    parser.ParamsKey,
		},
		{
			name:   "Msig Universal Receiver Hook",
			txType: parser.MethodMsigUniversalReceiverHook,
			f:      p.universalReceiverHook,
			key:    parser.ParamsKey,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, err := loadFile(manifest.MultisigKey, tt.txType, tt.key)
			require.NoError(t, err)
			require.NotNil(t, rawParams)

			got, err := tt.f(rawParams)
			require.NoError(t, err)
			require.NotNil(t, got)
			require.Contains(t, got, tt.key, fmt.Sprintf("%s could no be found in metadata", tt.key))
			require.NotNil(t, got[tt.key])
		})
	}
}

func TestParseMultisigMetadata(t *testing.T) {
	filePath := filepath.Join("..", "data", "actors", "multisig", "Metadata", "metadata_test.csv")
	file, err := os.Open(filePath)
	require.NoError(t, err, "Error opening CSV file")
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	require.NoError(t, err, "Error reading CSV file")

	for _, record := range records {
		require.Len(t, record, 3, "Invalid record")

		txType := record[0]
		txMetadata := record[1]
		expectedStr := record[2]

		result, err := ParseMultisigMetadata(txType, txMetadata)
		require.NoError(t, err, "Error parsing metadata for txType %s", txType)

		expected, err := unmarshalExpected(txType, expectedStr)
		require.NoError(t, err, "Error unmarshaling expected for txType %s", txType)

		resultJSON, err := json.Marshal(result)
		require.NoError(t, err, "Error marshaling result to JSON")

		expectedJSON, err := json.Marshal(expected)
		require.NoError(t, err, "Error marshaling expected to JSON")

		var resultMap map[string]interface{}
		var expectedMap map[string]interface{}
		require.NoError(t, json.Unmarshal(resultJSON, &resultMap), "Error unmarshaling result JSON")
		require.NoError(t, json.Unmarshal(expectedJSON, &expectedMap), "Error unmarshaling expected JSON")

		compareRetField(t, txType, resultMap, expectedMap)
		assert.Equal(t, expectedMap, resultMap, "Mismatch for other fields in txType %s", txType)
	}
}

func unmarshalExpected(txType, jsonStr string) (interface{}, error) {
	var v interface{}
	switch txType {
	case parser.MethodAddSigner:
		v = &multisig2.AddSignerParams{}
	case parser.MethodApprove:
		v = &ApproveValue{}
	case parser.MethodCancel:
		v = &CancelValue{}
	case parser.MethodChangeNumApprovalsThreshold:
		v = &multisig2.ChangeNumApprovalsThresholdParams{}
	case parser.MethodConstructor:
		v = &multisig2.ConstructorParams{}
	case parser.MethodLockBalance:
		v = &multisig2.LockBalanceParams{}
	case parser.MethodRemoveSigner:
		v = &multisig2.RemoveSignerParams{}
	case parser.MethodSend:
		v = &SendValue{}
	case parser.MethodSwapSigner:
		v = &multisig2.SwapSignerParams{}
	case parser.MethodMsigUniversalReceiverHook:
		v = &UniversalReceiverHookValue{}
	default:
		return nil, fmt.Errorf("unknown txType: %s", txType)
	}
	err := json.Unmarshal([]byte(jsonStr), v)
	return v, err
}

func compareRetField(t *testing.T, txType string, resultMap, expectedMap map[string]interface{}) {
	if strings.EqualFold(txType, parser.MethodApprove) {
		resultReturn := resultMap["Return"].(map[string]interface{})
		expectedReturn := expectedMap["Return"].(map[string]interface{})

		resultRet, resultRetExists := resultReturn["Ret"].(string)
		expectedRet, expectedRetExists := expectedReturn["Ret"].(string)

		if resultRetExists && expectedRetExists {
			resultRetBytes, err := base64.StdEncoding.DecodeString(resultRet)
			require.NoError(t, err, "Error decoding result Ret from Base64")

			assert.Equal(t, string(resultRetBytes), expectedRet, "Mismatch for Ret field in txType %s", txType)
		}

		delete(resultReturn, "Ret")
		delete(expectedReturn, "Ret")
	}
}
