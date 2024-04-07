package parser

import (
	"testing"

	"github.com/filecoin-project/go-state-types/exitcode"
)

func TestCheckExitCodeCommonError(t *testing.T) {
	tests := []struct {
		name string
		code exitcode.ExitCode
		want string
	}{
		{
			name: "ErrIllegalArgument",
			code: exitcode.ErrIllegalArgument,
			want: "ErrIllegalArgument(16)",
		},
		{
			name: "ErrNotFound",
			code: exitcode.ErrNotFound,
			want: "ErrNotFound(17)",
		},
		{
			name: "ErrForbidden",
			code: exitcode.ErrForbidden,
			want: "ErrForbidden(18)",
		},
		{
			name: "ErrInsufficientFunds",
			code: exitcode.ErrInsufficientFunds,
			want: "ErrInsufficientFunds(19)",
		},
		{
			name: "ErrIllegalState",
			code: exitcode.ErrIllegalState,
			want: "ErrIllegalState(20)",
		},
		{
			name: "ErrSerialization",
			code: exitcode.ErrSerialization,
			want: "ErrSerialization(21)",
		},
		{
			name: "ErrUnhandledMessage",
			code: exitcode.ErrUnhandledMessage,
			want: "ErrUnhandledMessage(22)",
		},
		{
			name: "ErrUnspecified",
			code: exitcode.ErrUnspecified,
			want: "ErrUnspecified(23)",
		},
		{
			name: "ErrAssertionFailed",
			code: exitcode.ErrAssertionFailed,
			want: "ErrAssertionFailed(24)",
		},
		{
			name: "ErrReadOnly",
			code: exitcode.ErrReadOnly,
			want: "ErrReadOnly(25)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckExitCodeCommonError(tt.code.String()); got != tt.want {
				t.Errorf("CheckCheCommonError() = %v, want %v", got, tt.want)
			}
		})
	}
}
