package types

import (
	"bytes"
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSectorContentChangedReturn_UnmarshalCBOR(t *testing.T) {
	tests := []struct {
		name     string
		raw      []byte
		expected SectorContentChangedReturn
		wantErr  bool
	}{
		{
			name:     "no sectors",
			raw:      []byte{0x80},
			expected: SectorContentChangedReturn{},
		},
		{
			// calibration 4100979, ProveReplicaUpdates3 bafy2bzaceddlcdlambexdwjc2sjgzxc4yg5xsfzdle6udkuyfpnyzohncfkn6
			name:     "one sector one piece",
			raw:      []byte{0x81, 0x81, 0xf5},
			expected: SectorContentChangedReturn{{true}},
		},
		{
			name:     "several sectors, mixed flags and an empty sector",
			raw:      []byte{0x83, 0x82, 0xf5, 0xf4, 0x80, 0x81, 0xf4},
			expected: SectorContentChangedReturn{{true, false}, {}, {false}},
		},
		{
			name:    "not an array",
			raw:     []byte{0xf5},
			wantErr: true,
		},
		{
			name:    "sector is not an array",
			raw:     []byte{0x81, 0xf5},
			wantErr: true,
		},
		{
			name:    "piece is not a bool",
			raw:     []byte{0x81, 0x81, 0x01},
			wantErr: true,
		},
		{
			name:    "truncated",
			raw:     []byte{0x82, 0x81, 0xf5},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got SectorContentChangedReturn
			err := got.UnmarshalCBOR(bytes.NewReader(tt.raw))
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.expected, got)
		})
	}
}

// mainnet 6559711, ProveCommitSectors3 bafy2bzaced7dq5gbepjalokgod675ya53ea35m4cnb2zi3uvqavmmzgkekmhy:
// one market notification for 32 sectors, 136 pieces, all accepted.
func TestSectorContentChangedReturn_MultiSector(t *testing.T) {
	raw, err := base64.StdEncoding.DecodeString("mCCB9YH1gvX1gvX1gvX1gvX1gvX1gvX1gvX1gvX1gvX1gvX1gvX1gvX1hPX19fWE9fX19YT19fX1hPX19fWE9fX19YT19fX1hPX19fWF9fX19fWH9fX19fX19Yj19fX19fX19Yj19fX19fX19Yj19fX19fX19Yn19fX19fX19fWP9fX19fX19fX19fX19fX1k/X19fX19fX19fX19fX19fX19fWB9YH1gfU=")
	require.NoError(t, err)

	var got SectorContentChangedReturn
	require.NoError(t, got.UnmarshalCBOR(bytes.NewReader(raw)))
	require.Len(t, got, 32)
	pieces := 0
	for _, sector := range got {
		for _, accepted := range sector {
			require.True(t, accepted)
			pieces++
		}
	}
	require.Equal(t, 136, pieces)
}
