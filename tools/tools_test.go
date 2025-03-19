package tools

import (
	"fmt"
	"os"
	"testing"

	filTypes "github.com/filecoin-project/lotus/chain/types"
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

func TestIsSupported(t *testing.T) {
	tests := []struct {
		name    string
		version version
		network string
		height  int64
		want    bool
	}{
		// {name: "V13 on calibration", version: V18, network: "calibration", height: 151, want: true},

		{name: "V7 on calibration", version: V7, network: "calibration", height: 2383680, want: false},
		{name: "V7 on mainnet", version: V7, network: "mainnet", height: 170000, want: false},
		{name: "V7 on calibration", version: V7, network: "calibration", height: 0, want: true},
		{name: "V7 on mainnet", version: V7, network: "mainnet", height: 10000, want: true},

		{name: "V9 on calibration", version: V9, network: "calibration", height: 265100, want: false},
		{name: "V9 on mainnet", version: V9, network: "mainnet", height: 265201, want: true},
		{name: "V9 on calibration", version: V9, network: "calibration", height: 265200, want: false},
		{name: "V9 on mainnet", version: V9, network: "mainnet", height: 265200, want: true},

		{name: "V24 on calibration", version: V24, network: "calibration", height: 2081674, want: true},
		{name: "V24 on mainnet", version: V24, network: "mainnet", height: 1427974, want: false},
		{name: "V24 on calibration", version: V24, network: "calibration", height: 2081672, want: false},
		{name: "V24 on mainnet", version: V24, network: "mainnet", height: 4461240, want: true},
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
		{name: "V10 on calibration", version: V10, height: 332640, want: false},
		{name: "V18 on calibration", version: V18, height: 332640, want: true},
		{name: "V22 on calibration", version: V22, height: 2791307, want: false},
		{name: "V25 on calibration", version: V25, height: 2791307, want: true},

		{name: "V16 on calibration", version: V16, height: 1000, want: false},
		{name: "V17 on calibration", version: V17, height: 16900, want: false},
		{name: "V18 on calibration", version: V18, height: 1000, want: true},

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
	case V18.IsSupported("calibration", height):
		t.Fatalf("V18 should be supported on calibration at height %d", height)
	case V19.IsSupported("calibration", height):
		t.Fatalf("V19 should not be supported on calibration at height %d", height)
	case V20.IsSupported("calibration", height):
	case V21.IsSupported("calibration", height):
		t.Fatalf("V21 should not be supported on calibration at height %d", height)
	case V22.IsSupported("calibration", height):
		t.Fatalf("V22 should not be supported on calibration at height %d", height)
	case V23.IsSupported("calibration", height):
		t.Fatalf("V23 should not be supported on calibration at height %d", height)
	case V24.IsSupported("calibration", height):
		t.Fatalf("V24 should be supported on calibration at height %d", height)
	default:
		t.Fatalf("V20 should be supported on calibration at height %d", height)
	}

	height = int64(15800)
	switch {
	case V15.IsSupported("calibration", height):
		t.Fatalf("V15 should not be supported on calibration at height %d", height)
	case V16.IsSupported("calibration", height):
		t.Fatalf("V16 should not be supported on calibration at height %d", height)
	case V17.IsSupported("calibration", height):
		t.Fatalf("V17 should not be supported on calibration at height %d", height)
	case V18.IsSupported("calibration", height):
	default:
		t.Fatalf("V16 should be supported on calibration at height %d", height)
	}
}

func TestVersionIterator(t *testing.T) {
	tests := []struct {
		name             string
		version          version
		network          string
		expectedVersions []version
	}{
		{name: "V1 on calibration", version: V1, network: "calibration",
			expectedVersions: []version{V1, V2, V3, V4, V5, V6, V7, V8, V9, V10, V11, V12, V13, V14, V15, V16, V17, V18, V19, V20, V21, V22, V23, V24, V25}},
		{name: "V1 on mainnet", version: V1, network: "mainnet", expectedVersions: []version{V1, V2, V3, V4, V5, V6, V7, V8, V9, V10, V11, V12, V13, V14, V15, V16, V17, V18, V19, V20, V21, V22, V23, V24}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var versions []version
			iter := NewVersionIterator(tt.version, tt.network)
			nodeVersionStart := tt.version.nodeVersion
			for {
				v, ok := iter.Next()
				if !ok {
					break
				}
				fmt.Printf("v: %d, nodeVersionStart: %d\n", v.nodeVersion, nodeVersionStart)
				require.Equal(t, nodeVersionStart, v.nodeVersion)
				versions = append(versions, v)
				nodeVersionStart++
			}
			require.Equal(t, len(tt.expectedVersions), len(versions))
		})

	}
}

func TestVersionsBefore(t *testing.T) {
	tests := []struct {
		name    string
		version version
		want    []version
	}{
		{name: "V1", version: V1, want: []version{V1}},
		{name: "V2", version: V2, want: []version{V1, V2}},
		{name: "V3", version: V3, want: []version{V1, V2, V3}},
		{name: "V4", version: V4, want: []version{V1, V2, V3, V4}},
		{name: "V5", version: V5, want: []version{V1, V2, V3, V4, V5}},
		{name: "V6", version: V6, want: []version{V1, V2, V3, V4, V5, V6}},
		{name: "V7", version: V7, want: []version{V1, V2, V3, V4, V5, V6, V7}},
		{name: "V8", version: V8, want: []version{V1, V2, V3, V4, V5, V6, V7, V8}}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			versions := VersionsBefore(tt.version)
			require.Equal(t, len(versions), len(tt.want))
			for i, v := range versions {
				require.Equal(t, v.nodeVersion, tt.want[i].nodeVersion)
			}
		})
	}
}

func TestVersionsAfter(t *testing.T) {
	tests := []struct {
		name    string
		version version
		want    []version
	}{
		{name: "V17", version: V17, want: []version{V17, V18, V19, V20, V21, V22, V23, V24}},
		{name: "V18", version: V18, want: []version{V18, V19, V20, V21, V22, V23, V24}},
		{name: "V19", version: V19, want: []version{V19, V20, V21, V22, V23, V24}},
		{name: "V20", version: V20, want: []version{V20, V21, V22, V23, V24}},
		{name: "V21", version: V21, want: []version{V21, V22, V23, V24}},
		{name: "V22", version: V22, want: []version{V22, V23, V24}},
		{name: "V23", version: V23, want: []version{V23, V24}},
		{name: "V24", version: V24, want: []version{V24}}, // mainnet
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			versions := VersionsAfter(tt.version)
			require.Equal(t, len(versions), len(tt.want))
			for i, v := range versions {
				require.Equal(t, v.nodeVersion, tt.want[i].nodeVersion)
			}
		})
	}
}

func TestGetSupportedVersions(t *testing.T) {
	tests := []struct {
		name    string
		network string
		want    []version
	}{
		{name: "calibration", network: "calibration", want: []version{V1, V2, V3, V4, V5, V6, V7, V8, V9, V10, V11, V12, V13, V14, V15, V16, V17, V18, V19, V20, V21, V22, V23, V24, V25}},
		{name: "mainnet", network: "mainnet", want: []version{V1, V2, V3, V4, V5, V6, V7, V8, V9, V10, V11, V12, V13, V14, V15, V16, V17, V18, V19, V20, V21, V22, V23, V24}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			versions := GetSupportedVersions(tt.network)
			fmt.Printf("versions: %v\n", versions)
			fmt.Printf("want: %v\n", tt.want)
			assert.Equal(t, len(versions), len(tt.want))
			for i, v := range versions {
				require.Equal(t, v.nodeVersion, tt.want[i].nodeVersion)
			}
		})
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
