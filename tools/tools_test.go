package tools

import (
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestBuildTipSetKeyHash(t *testing.T) {
	tests := []struct {
		name       string
		tipsetPath string
	}{
		{
			name:       "tipset",
			tipsetPath: "../data/tipsets/Tipset",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, err := os.Open(tt.tipsetPath)
			require.NoError(t, err)
			require.NotNil(t, file)
			var tipset filTypes.TipSet
			err = tipset.UnmarshalCBOR(file)
			require.NoError(t, err)
			require.NotNil(t, tipset)
			got, err := tipset.Key().Cid()
			require.NoError(t, err)
			require.NotEmpty(t, got.String())
		})
	}
}
