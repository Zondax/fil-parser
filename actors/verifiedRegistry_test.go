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
			name:   "Universal Receiver Hook",
			txType: parser.MethodMsigUniversalReceiverHook,
			f:      p.verifregUniversalReceiverHook,
		},
		//{ // TODO: `cbor input had wrong number of fields`
		//	name:   "Remove Expired Allocations",
		//	txType: parser.MethodRemoveExpiredAllocations,
		//	f:      p.removeExpiredClaims,
		//},
		{
			name:   "Get Claims",
			txType: parser.MethodGetClaims,
			f:      p.getClaims,
		},
		{
			name:   "Extend Claims Terms",
			txType: parser.MethodExtendClaimTerms,
			f:      p.extendClaimTerms,
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
