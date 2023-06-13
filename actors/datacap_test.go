package actors

import (
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
			name:   "Transfer From Exported",
			txType: parser.MethodTransferFromExported,
			f:      p.transferFromExported,
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
