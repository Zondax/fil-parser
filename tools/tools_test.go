package tools

import (
	"os"
	"testing"

	filTypes "github.com/filecoin-project/lotus/chain/types"
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

func TestIsSupported(t *testing.T) {
	tests := []struct {
		name    string
		version version
		network string
		height  int64
		want    bool
	}{

		//278543
		{name: "V13 on calibration", version: V18, network: "calibration", height: 151, want: true},

		// {name: "V7 on calibration", version: V7, network: "calibration", height: 2383680, want: false},
		// {name: "V7 on mainnet", version: V7, network: "mainnet", height: 170000, want: false},
		// {name: "V7 on calibration", version: V7, network: "calibration", height: 0, want: true},
		// {name: "V7 on mainnet", version: V7, network: "mainnet", height: 10000, want: true},

		// {name: "V9 on calibration", version: V9, network: "calibration", height: 265100, want: false},
		// {name: "V9 on mainnet", version: V9, network: "mainnet", height: 265201, want: true},
		// {name: "V9 on calibration", version: V9, network: "calibration", height: 265200, want: false},
		// {name: "V9 on mainnet", version: V9, network: "mainnet", height: 265200, want: true},

		// {name: "V24 on calibration", version: V24, network: "calibration", height: 2081674, want: true},
		// {name: "V24 on mainnet", version: V24, network: "mainnet", height: 1427974, want: false},
		// {name: "V24 on calibration", version: V24, network: "calibration", height: 2081672, want: false},
		// {name: "V24 on mainnet", version: V24, network: "mainnet", height: 4461240, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, tt.version.IsSupported(tt.network, tt.height))
		})
	}
}

func TestIsSupportedCalibrationEdgeCases(t *testing.T) {
	network := "calibration"

	tests := []struct {
		name    string
		version version
		height  int64
		want    bool
	}{
		{name: "V1 on calibration", version: V1, height: 0, want: true},
		{name: "V5 on calibration", version: V5, height: 193700, want: false},
		{name: "V11 on calibration", version: V11, height: 193789, want: false},
		{name: "V12 on calibration", version: V12, height: 193780, want: false},
		{name: "V13 on calibration", version: V13, height: 312700, want: false},
		{name: "V14 on calibration", version: V14, height: 312846, want: false},
		{name: "V16 on calibration", version: V16, height: 1044661, want: false},

		{name: "V16 on calibration", version: V16, height: 1000, want: true},
		{name: "V17 on calibration", version: V17, height: 16900, want: true},

		{name: "V23 on calibration", version: V23, height: 1779094, want: true},
		{name: "V23 on calibration", version: V21, height: 1419335, want: true},
		{name: "V15 on calibration", version: V17, height: 1419335, want: false},
		{name: "V14 on calibration", version: V17, height: 1419335, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, tt.version.IsSupported(network, tt.height))
		})
	}

	require.False(t, AnyIsSupported("calibration", 1419335, VersionsBefore(V17)...))
	height := int64(782006)
	switch {
	case V15.IsSupported("calibration", height):
		t.Fatalf("V15 should not be supported on calibration at height %d", height)
	case V16.IsSupported("calibration", height):
		t.Fatalf("V16 should not be supported on calibration at height %d", height)
	case V17.IsSupported("calibration", height):
		t.Fatalf("V17 should not be supported on calibration at height %d", height)
	case V20.IsSupported("calibration", height):
	default:
		t.Fatalf("V20 should be supported on calibration at height %d", height)
	}

	height = int64(15800)
	switch {
	case V15.IsSupported("calibration", height):
		t.Fatalf("V15 should not be supported on calibration at height %d", height)
	case V16.IsSupported("calibration", height):
	case V17.IsSupported("calibration", height):
		t.Fatalf("V17 should not be supported on calibration at height %d", height)
	default:
		t.Fatalf("V16 should be supported on calibration at height %d", height)
	}
}

/*
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
}*/
