package actors

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/parser"
	"os"
	"path/filepath"
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
	if err != nil {
		t.Fatalf("Error opening CSV file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		t.Fatalf("Error reading CSV file: %v", err)
	}

	for _, record := range records {
		if len(record) < 3 {
			t.Fatalf("Invalid record: %v", record)
		}
		txType := record[0]
		txMetadata := record[1]
		expectedStr := record[2]

		result, err := ParseMultisigMetadata(txType, txMetadata)
		if err != nil {
			t.Errorf("Error parsing metadata for txType %s: %v", txType, err)
			continue
		}

		expected, err := unmarshalExpected(txType, expectedStr)
		if err != nil {
			t.Errorf("Error unmarshaling expected for txType %s: %v", txType, err)
			continue
		}

		if diff := cmp.Diff(dereference(expected), dereference(result)); diff != "" {
			t.Errorf("Mismatch for txType %s (-expected +result):\n%s", txType, diff)
		}
	}
}

func unmarshalExpected(txType, jsonStr string) (interface{}, error) {
	var v interface{}
	switch txType {
	case parser.MethodAddSigner:
		v = &AddSignerValue{}
	case parser.MethodApprove:
		v = &ApproveValue{}
	case parser.MethodCancel:
		v = &CancelValue{}
	case parser.MethodChangeNumApprovalsThreshold:
		v = &ChangeNumApprovalsThresholdValue{}
	case parser.MethodConstructor:
		v = &ConstructorValue{}
	case parser.MethodLockBalance:
		v = &LockBalanceValue{}
	case parser.MethodRemoveSigner:
		v = &RemoveSignerValue{}
	case parser.MethodSend:
		v = &SendValue{}
	case parser.MethodSwapSigner:
		v = &SwapSignerValue{}
	case parser.MethodMsigUniversalReceiverHook:
		v = &UniversalReceiverHookValue{}
	default:
		return nil, fmt.Errorf("unknown txType: %s", txType)
	}
	err := json.Unmarshal([]byte(jsonStr), v)
	return v, err
}

// TODO: Better this
func dereference(v interface{}) interface{} {
	switch val := v.(type) {
	case *AddSignerValue:
		return *val
	case *ApproveValue:
		return *val
	case *CancelValue:
		return *val
	case *ChangeNumApprovalsThresholdValue:
		return *val
	case *ConstructorValue:
		return *val
	case *LockBalanceValue:
		return *val
	case *RemoveSignerValue:
		return *val
	case *SendValue:
		return *val
	case *SwapSignerValue:
		return *val
	case *UniversalReceiverHookValue:
		return *val
	default:
		return v
	}
}
