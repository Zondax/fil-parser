package multisig

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/stretchr/testify/require"
	cbg "github.com/whyrusleeping/cbor-gen"

	"github.com/zondax/fil-parser/tools"
)

// getProposeParams and getProposeReturn decode with the per-version maps, then dispatch on the
// concrete go-state-types type. Encode a real payload with each version's own type and decode it
// at that version's height, so a version missing from either type switch fails here.
func TestProposeHelpersHandleEveryCalibrationVersion(t *testing.T) {
	to, err := address.NewIDAddress(1234)
	require.NoError(t, err)

	for _, v := range tools.GetSupportedVersions(tools.CalibrationNetwork) {
		if v.NodeVersion() < tools.V16.NodeVersion() {
			continue // calibration parsing starts at V16
		}
		height := tools.DeterministicTestHeight(v)
		require.Equal(t, v.String(), tools.VersionFromHeight(tools.CalibrationNetwork, height).String())

		newParams, ok := proposeParams[v.String()]
		require.True(t, ok, "no ProposeParams entry for %s", v)
		p := newParams()
		pv := reflect.ValueOf(p).Elem()
		pv.FieldByName("To").Set(reflect.ValueOf(to))
		pv.FieldByName("Value").Set(reflect.ValueOf(big.NewInt(7)))
		var pbuf bytes.Buffer
		require.NoError(t, p.(cbg.CBORMarshaler).MarshalCBOR(&pbuf))

		_, _, gotTo, gotValue, _, err := getProposeParams(tools.CalibrationNetwork, height, pbuf.Bytes())
		require.NoError(t, err, "ProposeParams at %s (%T) is not handled", v, p)
		require.Equal(t, to, gotTo)
		require.Equal(t, "7", gotValue)

		newReturn, ok := proposeReturn[v.String()]
		require.True(t, ok, "no ProposeReturn entry for %s", v)
		r := newReturn()
		reflect.ValueOf(r).Elem().FieldByName("Applied").SetBool(true)
		var rbuf bytes.Buffer
		require.NoError(t, r.(cbg.CBORMarshaler).MarshalCBOR(&rbuf))

		applied, _, _, _, err := getProposeReturn(tools.CalibrationNetwork, height, rbuf.Bytes())
		require.NoError(t, err, "ProposeReturn at %s (%T) is not handled", v, r)
		require.True(t, applied)
	}
}
