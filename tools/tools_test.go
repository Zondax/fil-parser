package tools

import (
	"fmt"
	"os"
	"testing"

	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestBuildCidFromMessageTrace(t *testing.T) {
	h1, err := multihash.Sum([]byte("TEST"), multihash.SHA2_256, -1)
	assert.NoError(t, err)

	defaultCid := cid.NewCidV1(7, h1)

	tb := []struct {
		name         string
		actor        filTypes.ActorTrace
		parentMsgCID string
		wantCID      string
		wantErr      error
	}{
		{
			name:         "error marshaling actor state",
			actor:        filTypes.ActorTrace{},
			parentMsgCID: "bafy2bzaceab3xcn7qkcuj5oyifa6dn3ihke55bdmerphef4r6aorjdhk3uriq",
			wantErr:      fmt.Errorf("failed to write cid field t.Head: undefined cid"),
		},
		{
			name: "use defaultCodeCid when actor codeCID is undefined",
			actor: filTypes.ActorTrace{
				State: filTypes.Actor{
					Head: defaultCid,
				},
			},
			parentMsgCID: "bafy2bzaceab3xcn7qkcuj5oyifa6dn3ihke55bdmerphef4r6aorjdhk3uriq",
			wantCID:      "bafy2bzacebtrro4733sdya5vxtv2deuqeqznyaw4lngz4umdypniwitraz4fs",
		},
		{
			name: "use existing actor codeCID",
			actor: filTypes.ActorTrace{
				State: filTypes.Actor{
					Head: defaultCid,
					Code: defaultCid,
				},
			},
			parentMsgCID: "bafy2bzaceab3xcn7qkcuj5oyifa6dn3ihke55bdmerphef4r6aorjdhk3uriq",
			wantCID:      "bafy2bzacecczeqpns7edzaaz6tyuprzhho6pviz4yb4hedwepvnaklgbzn3ig",
		},
	}

	for i := range tb {
		tt := tb[i]
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gotCID, gotErr := BuildCidFromMessageTrace(tt.actor, tt.parentMsgCID)
			if tt.wantErr != nil {
				assert.Error(t, gotErr)
				assert.Equal(t, gotErr.Error(), tt.wantErr.Error())
				return
			}
			assert.NoError(t, gotErr)
			assert.Equal(t, tt.wantCID, gotCID)
		})
	}
}
