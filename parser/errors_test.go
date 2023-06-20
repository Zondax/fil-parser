package parser

import (
	"github.com/filecoin-project/go-state-types/exitcode"
	"testing"
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
			want: "ErrIllegalArgument",
		},
		{
			name: "ErrNotFound",
			code: exitcode.ErrNotFound,
			want: "ErrNotFound",
		},
		{
			name: "ErrForbidden",
			code: exitcode.ErrForbidden,
			want: "ErrForbidden",
		},
		{
			name: "ErrInsufficientFunds",
			code: exitcode.ErrInsufficientFunds,
			want: "ErrInsufficientFunds",
		},
		{
			name: "ErrIllegalState",
			code: exitcode.ErrIllegalState,
			want: "ErrIllegalState",
		},
		{
			name: "ErrSerialization",
			code: exitcode.ErrSerialization,
			want: "ErrSerialization",
		},
		{
			name: "ErrUnhandledMessage",
			code: exitcode.ErrUnhandledMessage,
			want: "ErrUnhandledMessage",
		},
		{
			name: "ErrUnspecified",
			code: exitcode.ErrUnspecified,
			want: "ErrUnspecified",
		},
		{
			name: "ErrAssertionFailed",
			code: exitcode.ErrAssertionFailed,
			want: "ErrAssertionFailed",
		},
		{
			name: "ErrReadOnly",
			code: exitcode.ErrReadOnly,
			want: "ErrReadOnly",
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
