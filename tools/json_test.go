package tools

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnmarshalJSONUseNumber_KeepsLargeIntegers(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{name: "MaxInt64", input: `{"NewExpiration":9223372036854775807}`},
		{name: "MinInt64", input: `{"v":-9223372036854775808}`},
		{name: "MaxUint64", input: `{"v":18446744073709551615}`},
		{name: "2^53+1", input: `{"v":9007199254740993}`},
		{name: "beyond uint64", input: `{"v":340282366920938463463374607431768211455}`},
		{name: "nested arrays and objects", input: `{"a":[9007199254740993,{"b":[18446744073709551615]}],"c":{"d":9223372036854775807}}`},
		{name: "top level array", input: `[9007199254740993,1,-2]`},
		{name: "top level number", input: `9223372036854775807`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got any
			require.NoError(t, UnmarshalJSONUseNumber([]byte(tt.input), &got))
			out, err := json.Marshal(got)
			require.NoError(t, err)
			assert.Equal(t, tt.input, string(out))

			// json.Unmarshal (float64) changes these values
			var lossy any
			require.NoError(t, json.Unmarshal([]byte(tt.input), &lossy))
			lossyOut, err := json.Marshal(lossy)
			require.NoError(t, err)
			assert.NotEqual(t, tt.input, string(lossyOut))
		})
	}
}

// Values float64 represents exactly must re-encode byte-identically to the json.Unmarshal path,
// so existing outputs do not change.
func TestUnmarshalJSONUseNumber_SameOutputAsUnmarshalForExactValues(t *testing.T) {
	inputs := []string{
		`{}`,
		`[]`,
		`null`,
		`{"a":0,"b":1,"c":-1,"d":42,"e":9007199254740992,"f":-9007199254740992}`,
		`{"fraction":0.5,"neg":-1.25,"small":1e-7,"tiny":5e-324,"big":1e+21,"max":1.7976931348623157e+308}`,
		`{"amount":"340282366920938463463374607431768211455","zero":"0","neg":"-1"}`,
		`{"s":"text","t":true,"f":false,"n":null,"cid":{"/":"bafy2bzacedyg4uiezhkabfgvbbtfluie5xiukxdiivnjlfqsn7omti3r63rxy"}}`,
		`{"nested":{"arr":[1,[2,[3,{"x":4.5}]],"y"],"obj":{"z":[]}}}`,
		`[{"Params":{"Signers":["f01","f02"],"NumApprovalsThreshold":2,"UnlockDuration":5600000000000000,"StartEpoch":0}}]`,
		`"string"`,
		`123`,
		`true`,
	}
	for _, input := range inputs {
		t.Run(input, func(t *testing.T) {
			var exact, lossy any
			require.NoError(t, UnmarshalJSONUseNumber([]byte(input), &exact))
			require.NoError(t, json.Unmarshal([]byte(input), &lossy))
			exactOut, err := json.Marshal(exact)
			require.NoError(t, err)
			lossyOut, err := json.Marshal(lossy)
			require.NoError(t, err)
			assert.Equal(t, string(lossyOut), string(exactOut))
		})
	}
}

func TestUnmarshalJSONUseNumber_TypedTargets(t *testing.T) {
	type typed struct {
		I   int64   `json:"I"`
		U   uint64  `json:"U"`
		F   float64 `json:"F"`
		S   string  `json:"S"`
		Any any     `json:"Any"`
	}
	input := `{"I":9223372036854775807,"U":18446744073709551615,"F":0.5,"S":"x","Any":9007199254740993}`

	var got, want typed
	require.NoError(t, UnmarshalJSONUseNumber([]byte(input), &got))
	require.NoError(t, json.Unmarshal([]byte(input), &want))
	assert.Equal(t, want.I, got.I)
	assert.Equal(t, want.U, got.U)
	assert.Equal(t, want.F, got.F)
	assert.Equal(t, want.S, got.S)
	assert.Equal(t, json.Number("9007199254740993"), got.Any)

	var m map[string]any
	require.NoError(t, UnmarshalJSONUseNumber([]byte(input), &m))
	assert.Equal(t, json.Number("9223372036854775807"), m["I"])
}

// Errors must match json.Unmarshal, including the rejection of trailing data.
func TestUnmarshalJSONUseNumber_ErrorsLikeUnmarshal(t *testing.T) {
	inputs := []string{
		``,
		` `,
		`{`,
		`{"a":}`,
		`{} x`,
		`{}{}`,
		`[1,2]]`,
		`{"a":1}garbage`,
		`nul`,
		`01`,
	}
	for _, input := range inputs {
		t.Run(input, func(t *testing.T) {
			got := map[string]any{"untouched": true}
			want := map[string]any{"untouched": true}
			gotErr := UnmarshalJSONUseNumber([]byte(input), &got)
			wantErr := json.Unmarshal([]byte(input), &want)
			require.Error(t, wantErr)
			require.Error(t, gotErr)
			assert.Equal(t, wantErr.Error(), gotErr.Error())
			assert.Equal(t, want, got)
		})
	}

	// type mismatch: same error type as json.Unmarshal
	var m map[string]any
	err := UnmarshalJSONUseNumber([]byte(`[1]`), &m)
	var typeErr *json.UnmarshalTypeError
	require.ErrorAs(t, err, &typeErr)

	// trailing whitespace is fine, as with json.Unmarshal
	require.NoError(t, UnmarshalJSONUseNumber([]byte(" {\"a\":1} \n\t\r"), &m))
	assert.Equal(t, json.Number("1"), m["a"])
}

func TestJSONNumberToInt64(t *testing.T) {
	tests := []struct {
		name   string
		input  any
		want   int64
		wantOk bool
	}{
		{name: "json.Number MaxInt64", input: json.Number("9223372036854775807"), want: math.MaxInt64, wantOk: true},
		{name: "json.Number MinInt64", input: json.Number("-9223372036854775808"), want: math.MinInt64, wantOk: true},
		{name: "json.Number 2^53+1", input: json.Number("9007199254740993"), want: 9007199254740993, wantOk: true},
		{name: "json.Number small", input: json.Number("11"), want: 11, wantOk: true},
		{name: "json.Number negative", input: json.Number("-5"), want: -5, wantOk: true},
		{name: "json.Number exponent", input: json.Number("1e3"), want: 1000, wantOk: true},
		{name: "json.Number fraction truncates like float64", input: json.Number("2.9"), want: 2, wantOk: true},
		{name: "json.Number invalid", input: json.Number("abc"), wantOk: false},
		{name: "float64", input: float64(42), want: 42, wantOk: true},
		{name: "int", input: 7, want: 7, wantOk: true},
		{name: "uint64", input: uint64(8), want: 8, wantOk: true},
		{name: "string", input: "1", wantOk: false},
		{name: "nil", input: nil, wantOk: false},
		{name: "bool", input: true, wantOk: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := JSONNumberToInt64(tt.input)
			require.Equal(t, tt.wantOk, ok)
			if ok {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestJSONNumberToUint64(t *testing.T) {
	tests := []struct {
		name   string
		input  any
		want   uint64
		wantOk bool
	}{
		{name: "json.Number MaxUint64", input: json.Number("18446744073709551615"), want: math.MaxUint64, wantOk: true},
		{name: "json.Number MaxInt64", input: json.Number("9223372036854775807"), want: math.MaxInt64, wantOk: true},
		{name: "json.Number 2^53+1", input: json.Number("9007199254740993"), want: 9007199254740993, wantOk: true},
		{name: "json.Number small", input: json.Number("3"), want: 3, wantOk: true},
		{name: "json.Number exponent", input: json.Number("1e3"), want: 1000, wantOk: true},
		{name: "json.Number invalid", input: json.Number(""), wantOk: false},
		{name: "float64", input: float64(42), want: 42, wantOk: true},
		{name: "uint32", input: uint32(9), want: 9, wantOk: true},
		{name: "int64", input: int64(10), want: 10, wantOk: true},
		{name: "string", input: "1", wantOk: false},
		{name: "nil", input: nil, wantOk: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := JSONNumberToUint64(tt.input)
			require.Equal(t, tt.wantOk, ok)
			if ok {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
