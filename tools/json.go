package tools

import (
	"bytes"
	"encoding/json"
	"strconv"
)

// UnmarshalJSONUseNumber behaves like json.Unmarshal, except that a JSON number decoded into an
// interface value (any, map[string]any, []any, or an interface{} struct field) is kept as a
// json.Number instead of being converted to float64.
//
// float64 only represents integers exactly up to 2^53, so json.Unmarshal silently rounds larger
// values (epochs such as MaxInt64, uint64 ids, ...): 9223372036854775807 comes back as
// 9223372036854775808 and is re-encoded as 9223372036854776000. A json.Number keeps the original
// text and json.Marshal writes it back verbatim. Decoding into typed fields is unchanged.
//
// Code reading numbers out of the decoded value must accept json.Number; see JSONNumberToInt64
// and JSONNumberToUint64.
func UnmarshalJSONUseNumber(data []byte, v any) error {
	// json.Unmarshal checks that the whole input is a single valid JSON value before decoding
	// anything. A json.Decoder does not reject trailing data and reports syntax errors
	// differently, so invalid input is delegated to json.Unmarshal: the error is the same as
	// before and v is left untouched.
	if !json.Valid(data) {
		return json.Unmarshal(data, v)
	}
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	return decoder.Decode(v)
}

// JSONNumberToInt64 converts a number decoded from JSON to an int64. It accepts a json.Number
// (UnmarshalJSONUseNumber), a float64 (json.Unmarshal) and Go integer types (values that were
// never serialized). Integer json.Numbers are parsed exactly; any other json.Number (fraction,
// exponent, out of range) falls back to the float64 conversion json.Unmarshal would have used.
func JSONNumberToInt64(value any) (int64, bool) {
	switch n := value.(type) {
	case json.Number:
		if v, err := strconv.ParseInt(n.String(), 10, 64); err == nil {
			return v, true
		}
		f, err := n.Float64()
		if err != nil && !isRangeError(err) {
			return 0, false
		}
		return int64(f), true
	case float64:
		return int64(n), true
	case int:
		return int64(n), true
	case int8:
		return int64(n), true
	case int16:
		return int64(n), true
	case int32:
		return int64(n), true
	case int64:
		return n, true
	case uint:
		return int64(n), true // #nosec G115 -- same wrap-around as the float64 conversion
	case uint8:
		return int64(n), true
	case uint16:
		return int64(n), true
	case uint32:
		return int64(n), true
	case uint64:
		return int64(n), true // #nosec G115 -- same wrap-around as the float64 conversion
	}
	return 0, false
}

// JSONNumberToUint64 is JSONNumberToInt64 for unsigned values. Integer json.Numbers up to
// math.MaxUint64 are parsed exactly; anything else uses the uint64(float64) conversion
// json.Unmarshal callers used before.
func JSONNumberToUint64(value any) (uint64, bool) {
	switch n := value.(type) {
	case json.Number:
		if v, err := strconv.ParseUint(n.String(), 10, 64); err == nil {
			return v, true
		}
		f, err := n.Float64()
		if err != nil && !isRangeError(err) {
			return 0, false
		}
		return uint64(f), true
	case float64:
		return uint64(n), true
	case uint:
		return uint64(n), true
	case uint8:
		return uint64(n), true
	case uint16:
		return uint64(n), true
	case uint32:
		return uint64(n), true
	case uint64:
		return n, true
	}
	if v, ok := JSONNumberToInt64(value); ok {
		return uint64(v), true // #nosec G115 -- Go integer conversion, as for any int value
	}
	return 0, false
}

func isRangeError(err error) bool {
	numErr, ok := err.(*strconv.NumError)
	return ok && numErr.Err == strconv.ErrRange
}
