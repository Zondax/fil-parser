package miner

import (
	"bytes"
	"context"
	"testing"

	miner19 "github.com/filecoin-project/go-state-types/builtin/v19/miner"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"

	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

// TestNV29UpgradeSectorQualityRoundTrip proves miner method 37 (added in builtin-actors v19,
// NV29 / Solstice) is wired to the correct param type, not merely registered. The coverage
// tests push empty payloads and so would not catch a wrong struct.
func TestNV29UpgradeSectorQualityRoundTrip(t *testing.T) {
	const nv29Height = int64(999999999999999)
	require.Equal(t, "V29", tools.VersionFromHeight(tools.CalibrationNetwork, nv29Height).String())

	var buf bytes.Buffer
	require.NoError(t, (&miner19.UpgradeSectorQualityParams{}).MarshalCBOR(&buf))

	got, _, err := New(nil).Parse(context.Background(), tools.CalibrationNetwork, nv29Height,
		parser.MethodUpgradeSectorQuality,
		&parser.LotusMessage{Params: buf.Bytes()},
		&parser.LotusMessageReceipt{},
		cid.Undef, filTypes.EmptyTSK, true)
	require.NoError(t, err)
	require.Contains(t, got, parser.ParamsKey)
}
