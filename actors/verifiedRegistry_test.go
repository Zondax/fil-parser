package actors

import (
	"fmt"
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/parser"
	"testing"
)

func TestActorParser_verifiedWithParamsOrReturn(t *testing.T) {
	p := getActorParser()
	tests := []struct {
		name   string
		txType string
		f      func([]byte) (map[string]interface{}, error)
		key    string
	}{
		{
			name:   "Add Verifier",
			txType: parser.MethodAddVerifier,
			f:      p.addVerifier,
			key:    parser.ParamsKey,
		},
		{
			name:   "Add Verified Client",
			txType: parser.MethodAddVerifiedClient,
			f:      p.addVerifiedClient,
			key:    parser.ParamsKey,
		},
		{
			name:   "Add Verified Client Exported",
			txType: parser.MethodAddVerifiedClientExported,
			f:      p.addVerifiedClient,
			key:    parser.ParamsKey,
		},
		{
			name:   "Use Bytes",
			txType: parser.MethodUseBytes,
			f:      p.useBytes,
			key:    parser.ParamsKey,
		},
		{
			name:   "Restore Bytes",
			txType: parser.MethodRestoreBytes,
			f:      p.restoreBytes,
			key:    parser.ParamsKey,
		},
		{
			name:   "Remove Verified Client DataCap",
			txType: parser.MethodRemoveVerifiedClientDataCap,
			f:      p.removeVerifiedClientDataCap,
			key:    parser.ParamsKey,
		},
		{
			name:   "Deprecated1",
			txType: parser.MethodVerifiedDeprecated1,
			f:      p.deprecated1,
			key:    parser.ParamsKey,
		},
		{
			name:   "Deprecated2",
			txType: parser.MethodVerifiedDeprecated2,
			f:      p.deprecated2,
			key:    parser.ParamsKey,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, err := loadFile(manifest.VerifregKey, tt.txType, tt.key)
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

func TestActorParser_verifiedWithParamsAndReturn(t *testing.T) {
	p := getActorParser()
	tests := []struct {
		name   string
		txType string
		f      func([]byte, []byte) (map[string]interface{}, error)
	}{
		{
			name:   "Claim Allocations",
			txType: parser.MethodClaimAllocations,
			f:      p.claimAllocations,
		},
		{
			name:   "Extend Claim Terms",
			txType: parser.MethodExtendClaimTerms,
			f:      p.extendClaimTerms,
		},
		{
			name:   "Extend Claims Terms Exported",
			txType: parser.MethodExtendClaimTermsExported,
			f:      p.extendClaimTerms,
		},
		{
			name:   "Universal Receiver Hook",
			txType: parser.MethodMsigUniversalReceiverHook,
			f:      p.verifregUniversalReceiverHook,
		},
		{
			name:   "Remove Expired Allocations",
			txType: parser.MethodRemoveExpiredAllocations,
			f:      p.removeExpiredAllocations,
		},
		{
			name:   "Remove Expired Allocations Exported",
			txType: parser.MethodRemoveExpiredAllocationsExported,
			f:      p.removeExpiredAllocations,
		},
		{
			name:   "Get Claims",
			txType: parser.MethodGetClaims,
			f:      p.getClaims,
		},
		{
			name:   "Get Claims Exported",
			txType: parser.MethodGetClaimsExported,
			f:      p.getClaims,
		},
		{
			name:   "Remove Expired Claims",
			txType: parser.MethodRemoveExpiredClaims,
			f:      p.removeExpiredClaims,
		},
		{
			name:   "Remove Expired Claims Exported",
			txType: parser.MethodRemoveExpiredClaimsExported,
			f:      p.removeExpiredClaims,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, rawReturn, err := getParmasAndReturn(manifest.VerifregKey, tt.txType)
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
