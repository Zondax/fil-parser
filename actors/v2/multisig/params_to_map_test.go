package multisig

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"math"
	"testing"

	minerV13 "github.com/filecoin-project/go-state-types/builtin/v13/miner"
	"github.com/stretchr/testify/require"

	"github.com/zondax/fil-parser/parser"
)

// Inner params of the multisig Propose in mainnet message
// bafy2bzacedyg4uiezhkabfgvbbtfluie5xiukxdiivnjlfqsn7omti3r63rxy: miner ChangeBeneficiary with
// NewExpiration = MaxInt64 and NewQuota = 2^128-1 (Lotus StateDecodeParams returns both exactly).
const changeBeneficiaryCBOR = "834500a19c81015100ffffffffffffffffffffffffffffffff1b7fffffffffffffff"

func TestParamsToMapKeepsLargeIntegers(t *testing.T) {
	raw, err := hex.DecodeString(changeBeneficiaryCBOR)
	require.NoError(t, err)
	var params minerV13.ChangeBeneficiaryParams
	require.NoError(t, params.UnmarshalCBOR(bytes.NewReader(raw)))

	m := &Msig{}
	got, err := m.paramsToMap(&params)
	require.NoError(t, err)
	require.Equal(t, json.Number("9223372036854775807"), got["NewExpiration"])
	require.Equal(t, "340282366920938463463374607431768211455", got["NewQuota"]) // token amounts stay strings
	require.Equal(t, "f02117153", got["NewBeneficiary"])

	out, err := json.Marshal(got)
	require.NoError(t, err)
	require.Equal(t, `{"NewBeneficiary":"f02117153","NewExpiration":9223372036854775807,"NewQuota":"340282366920938463463374607431768211455"}`, string(out))
}

func TestParamsToMapNonObjectFallback(t *testing.T) {
	m := &Msig{}

	got, err := m.paramsToMap(uint64(math.MaxUint64))
	require.NoError(t, err)
	require.Equal(t, map[string]any{parser.ValueKey: json.Number("18446744073709551615")}, got)

	got, err = m.paramsToMap([]uint64{9007199254740993, 1})
	require.NoError(t, err)
	out, err := json.Marshal(got)
	require.NoError(t, err)
	require.Equal(t, `{"Value":[9007199254740993,1]}`, string(out))

	got, err = m.paramsToMap("text")
	require.NoError(t, err)
	require.Equal(t, map[string]any{parser.ValueKey: "text"}, got)
}

// Small values, negatives, fractions, strings, nested arrays/objects re-encode exactly as with the
// previous float64 decoding.
func TestParamsToMapSmallValuesUnchanged(t *testing.T) {
	type inner struct {
		A []int64        `json:"A"`
		B map[string]int `json:"B"`
	}
	params := struct {
		Small    int64   `json:"Small"`
		Negative int64   `json:"Negative"`
		Fraction float64 `json:"Fraction"`
		Amount   string  `json:"Amount"`
		Nested   inner   `json:"Nested"`
		Nil      *int    `json:"Nil"`
	}{Small: 42, Negative: -7, Fraction: 0.25, Amount: "1000000000000000000", Nested: inner{A: []int64{1, -2, 9007199254740992}, B: map[string]int{"x": 3}}}

	m := &Msig{}
	got, err := m.paramsToMap(params)
	require.NoError(t, err)
	out, err := json.Marshal(got)
	require.NoError(t, err)

	// the previous implementation: json.Unmarshal into map[string]any (float64 numbers)
	tmp, err := json.Marshal(params)
	require.NoError(t, err)
	var legacy map[string]any
	require.NoError(t, json.Unmarshal(tmp, &legacy))
	legacyOut, err := json.Marshal(legacy)
	require.NoError(t, err)

	require.Equal(t, string(legacyOut), string(out))
}
