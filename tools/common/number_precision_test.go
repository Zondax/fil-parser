package common_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/tools/common"
)

// GetInteger/GetIntegerSlice read json.Number (tools.UnmarshalJSONUseNumber) exactly and keep
// accepting float64 (json.Unmarshal).
func TestGetInteger_JSONNumber(t *testing.T) {
	var m map[string]interface{}
	require.NoError(t, tools.UnmarshalJSONUseNumber([]byte(`{"maxInt64":9223372036854775807,"minInt64":-9223372036854775808,"maxUint64":18446744073709551615,"above2pow53":9007199254740993,"small":7,"neg":-3,"str":"1","null":null}`), &m))

	gotInt64, err := common.GetInteger[int64](m, "maxInt64", false)
	require.NoError(t, err)
	assert.Equal(t, int64(math.MaxInt64), gotInt64)

	gotInt64, err = common.GetInteger[int64](m, "minInt64", false)
	require.NoError(t, err)
	assert.Equal(t, int64(math.MinInt64), gotInt64)

	gotUint64, err := common.GetInteger[uint64](m, "maxUint64", false)
	require.NoError(t, err)
	assert.Equal(t, uint64(math.MaxUint64), gotUint64)

	gotUint64, err = common.GetInteger[uint64](m, "above2pow53", false)
	require.NoError(t, err)
	assert.Equal(t, uint64(9007199254740993), gotUint64)

	gotInt, err := common.GetInteger[int](m, "small", false)
	require.NoError(t, err)
	assert.Equal(t, 7, gotInt)

	gotInt, err = common.GetInteger[int](m, "neg", false)
	require.NoError(t, err)
	assert.Equal(t, -3, gotInt)

	gotUint32, err := common.GetInteger[uint32](m, "small", false)
	require.NoError(t, err)
	assert.Equal(t, uint32(7), gotUint32)

	_, err = common.GetInteger[int64](m, "str", false)
	require.Error(t, err)

	_, err = common.GetInteger[int64](m, "missing", false)
	require.Error(t, err)

	gotInt64, err = common.GetInteger[int64](m, "null", true)
	require.NoError(t, err)
	assert.Equal(t, int64(0), gotInt64)

	gotInt64, err = common.GetInteger[int64](map[string]interface{}{"f": float64(9007199254740992)}, "f", false)
	require.NoError(t, err)
	assert.Equal(t, int64(9007199254740992), gotInt64)
}

func TestGetIntegerSlice_JSONNumber(t *testing.T) {
	var m map[string]interface{}
	require.NoError(t, tools.UnmarshalJSONUseNumber([]byte(`{"ids":[18446744073709551615,9007199254740993,0,67798],"mixed":[1,"2"],"null":null}`), &m))

	got, err := common.GetIntegerSlice[uint64](m, "ids", false)
	require.NoError(t, err)
	assert.Equal(t, []uint64{math.MaxUint64, 9007199254740993, 0, 67798}, got)

	_, err = common.GetIntegerSlice[uint64](m, "mixed", false)
	require.Error(t, err)

	// absent / null keys keep returning an empty, non-nil slice (it is marshalled as [])
	for _, k := range []string{"null", "missing"} {
		got, err = common.GetIntegerSlice[uint64](m, k, true)
		require.NoError(t, err)
		require.NotNil(t, got)
		assert.Empty(t, got)
		out, err := json.Marshal(got)
		require.NoError(t, err)
		assert.Equal(t, "[]", string(out))
	}

	_, err = common.GetIntegerSlice[uint64](m, "missing", false)
	require.Error(t, err)

	// float64 values from json.Unmarshal are still accepted
	gotInt, err := common.GetIntegerSlice[int](map[string]interface{}{"k": []interface{}{float64(1), float64(2)}}, "k", false)
	require.NoError(t, err)
	assert.Equal(t, []int{1, 2}, gotInt)
}
