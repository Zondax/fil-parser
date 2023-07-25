package actors

import (
	"fmt"
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/parser"
	"testing"
)

func TestActorParser_datacapWithParamsAndReturn(t *testing.T) {
	p := getActorParser()
	tests := []struct {
		name   string
		txType string
		f      func([]byte, []byte) (map[string]interface{}, error)
	}{
		{
			name:   "Mint Exported",
			txType: parser.MethodMintExported,
			f:      p.mintExported,
		},
		{
			name:   "Burn Exported",
			txType: parser.MethodBurnExported,
			f:      p.burnExported,
		},
		{
			name:   "Balance Exported",
			txType: parser.MethodBalanceExported,
			f:      p.balanceExported,
		},
		{
			name:   "Transfer Exported",
			txType: parser.MethodTransferExported,
			f:      p.transferExported,
		},
		{
			name:   "Transfer From Exported",
			txType: parser.MethodTransferFromExported,
			f:      p.transferFromExported,
		},
		{
			name:   "Destroy Exported",
			txType: parser.MethodDestroyExported,
			f:      p.destroyExported,
		},
		{
			name:   "Increase Allowance Exported",
			txType: parser.MethodIncreaseAllowanceExported,
			f:      p.increaseAllowanceExported,
		},
		{
			name:   "Decrease Allowance Exported",
			txType: parser.MethodDecreaseAllowanceExported,
			f:      p.decreaseAllowanceExported,
		},
		{
			name:   "Revoke Allowance Exported",
			txType: parser.MethodRevokeAllowanceExported,
			f:      p.revokeAllowanceExported,
		},
		{
			name:   "Burn From Exported",
			txType: parser.MethodBurnFromExported,
			f:      p.burnFromExported,
		},
		{
			name:   "Allowance Exported",
			txType: parser.MethodAllowanceExported,
			f:      p.allowanceExported,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, rawReturn, err := getParmasAndReturn(manifest.DatacapKey, tt.txType)
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

func TestActorParser_datacapWithParamsOrReturn(t *testing.T) {
	p := getActorParser()
	tests := []struct {
		name   string
		txType string
		f      func([]byte) (map[string]interface{}, error)
		key    string
	}{
		{
			name:   "Name Exported",
			txType: parser.MethodNameExported,
			f:      p.nameExported,
			key:    parser.ReturnKey,
		},
		{
			name:   "Symbol Exported",
			txType: parser.MethodSymbolExported,
			f:      p.symbolExported,
			key:    parser.ReturnKey,
		},
		{
			name:   "Total Supply Exported",
			txType: parser.MethodTotalSupplyExported,
			f:      p.totalSupplyExported,
			key:    parser.ReturnKey,
		},
		{
			name:   "Granularity Exported",
			txType: parser.MethodGranularityExported,
			f:      p.granularityExported,
			key:    parser.ReturnKey,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, err := loadFile(manifest.DatacapKey, tt.txType, tt.key)
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
