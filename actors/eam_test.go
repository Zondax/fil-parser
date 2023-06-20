package actors

import (
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
	"testing"
)

func TestActorParser_eamCreates(t *testing.T) {
	p := getActorParser()
	tests := []struct {
		name   string
		txType string
		f      func([]byte, []byte, cid.Cid) (map[string]interface{}, *types.AddressInfo, error)
	}{
		{
			name:   "Create",
			txType: parser.MethodCreate,
			f:      p.parseCreate,
		},
		{
			name:   "Create2",
			txType: parser.MethodCreate2,
			f:      p.parseCreate2,
		},
		{
			name:   "Create External",
			txType: parser.MethodCreateExternal,
			f:      p.parseCreateExternal,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, rawReturn, err := getParmasAndReturn(manifest.EamKey, tt.txType)
			require.NoError(t, err)
			require.NotNil(t, rawParams)
			require.NotNil(t, rawReturn)

			msg, err := deserializeMessage(manifest.EamKey, tt.txType)
			require.NoError(t, err)
			require.NotNil(t, msg)
			got, addr, err := tt.f(rawParams, rawReturn, msg.Cid)
			require.NoError(t, err)
			require.NotNil(t, got)
			require.NotNil(t, addr)
			require.Contains(t, got, parser.ParamsKey, "Params could no be found in metadata")
			require.NotNil(t, got[parser.ParamsKey])
			require.Contains(t, got, parser.ReturnKey, "Return could no be found in metadata")
			require.NotNil(t, got[parser.ReturnKey])
		})
	}
}
