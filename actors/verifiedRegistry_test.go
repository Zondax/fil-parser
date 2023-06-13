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
		name     string
		txType   string
		f        func([]byte) (map[string]interface{}, error)
		fileName string
		key      string
	}{
		{
			name:     "Add Verified Client",
			txType:   parser.MethodAddVerifiedClient,
			f:        p.addVerifiedClient,
			fileName: "params",
			key:      parser.ParamsKey,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, err := loadFile(manifest.VerifregKey, tt.txType, tt.fileName)
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
