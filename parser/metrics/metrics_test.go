package metrics

import (
	"regexp"
	"testing"
)

func TestRegexRules(t *testing.T) {
	tests := []struct {
		name string
		expr *regexp.Regexp
		err  string
	}{
		{
			name: "resolution lookup failed",
			expr: errResolutionLookupPattern,
			err:  "resolution lookup failed (f410fh6hh4552jjvzdhfcni57ro73ohko5gn5wvjhkbq): resolve address f410fh6hh4552jjvzdhfcni57ro73ohko5gn5wvjhkbq: actor not found",
		},
		{
			name: "address flaged as bad",
			expr: errBadAddressPattern,
			err:  "address f410fh6hh4552jjvzdhfcni57ro73ohko5gn5wvjhkbq is flagged as bad",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.expr.MatchString(tt.err) {
				t.Error("err did not match", tt.err)
			}
		})
	}
}
