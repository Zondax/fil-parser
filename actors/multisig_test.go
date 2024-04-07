package actors

import (
	"fmt"
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/parser"
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
