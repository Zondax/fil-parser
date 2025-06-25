package tools

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

func TestGetExitcodeStatus(t *testing.T) {
	tests := []struct {
		name     string
		exitCode exitcode.ExitCode
		want     string
	}{
		{
			name:     "Ok",
			exitCode: exitcode.Ok,
			want:     "Ok",
		},
		{
			name:     "SysErrSenderInvalid",
			exitCode: exitcode.SysErrSenderInvalid,
			want:     "SysErrSenderInvalid",
		},
		{
			name:     "SysErrSenderStateInvalid",
			exitCode: exitcode.SysErrSenderStateInvalid,
			want:     "SysErrSenderStateInvalid",
		},
		{
			name:     "SysErrIllegalInstruction",
			exitCode: exitcode.SysErrIllegalInstruction,
			want:     "SysErrIllegalInstruction",
		},
		{
			name:     "SysErrInvalidReceiver",
			exitCode: exitcode.SysErrInvalidReceiver,
			want:     "SysErrInvalidReceiver",
		},
		{
			name:     "SysErrInsufficientFunds",
			exitCode: exitcode.SysErrInsufficientFunds,
			want:     "SysErrInsufficientFunds",
		},
		{
			name:     "SysErrOutOfGas",
			exitCode: exitcode.SysErrOutOfGas,
			want:     "SysErrOutOfGas",
		},
		{
			name:     "SysErrIllegalExitCode",
			exitCode: exitcode.SysErrIllegalExitCode,
			want:     "SysErrIllegalExitCode",
		},
		{
			name:     "SysErrFatal",
			exitCode: exitcode.SysErrFatal,
			want:     "SysFatal",
		},
		{
			name:     "SysErrMissingReturn",
			exitCode: exitcode.SysErrMissingReturn,
			want:     "SysErrMissingReturn",
		},
		{
			name:     "SysErrReserved1",
			exitCode: exitcode.SysErrReserved1,
			want:     "SysErrReserved1",
		},
		{
			name:     "SysErrReserved2",
			exitCode: exitcode.SysErrReserved2,
			want:     "SysErrReserved2",
		},
		{
			name:     "SysErrReserved3",
			exitCode: exitcode.SysErrReserved3,
			want:     "SysErrReserved3",
		},
		{
			name:     "SysErrReserved4",
			exitCode: exitcode.SysErrReserved4,
			want:     "SysErrReserved4",
		},
		{
			name:     "SysErrReserved5",
			exitCode: exitcode.SysErrReserved5,
			want:     "SysErrReserved5",
		},
		{
			name:     "SysErrReserved6",
			exitCode: exitcode.SysErrReserved6,
			want:     "SysErrReserved6",
		},
		{
			name:     "ErrIllegalArgument",
			exitCode: exitcode.ErrIllegalArgument,
			want:     "ErrIllegalArgument",
		},
		{
			name:     "ErrNotFound",
			exitCode: exitcode.ErrNotFound,
			want:     "ErrNotFound",
		},
		{
			name:     "ErrForbidden",
			exitCode: exitcode.ErrForbidden,
			want:     "ErrForbidden",
		},
		{
			name:     "ErrInsufficientFunds",
			exitCode: exitcode.ErrInsufficientFunds,
			want:     "ErrInsufficientFunds",
		},
		{
			name:     "ErrIllegalState",
			exitCode: exitcode.ErrIllegalState,
			want:     "ErrIllegalState",
		},
		{
			name:     "ErrSerialization",
			exitCode: exitcode.ErrSerialization,
			want:     "ErrSerialization",
		},
		{
			name:     "ErrUnhandledMessage",
			exitCode: exitcode.ErrUnhandledMessage,
			want:     "ErrUnhandledMessage",
		},
		{
			name:     "ErrUnspecified",
			exitCode: exitcode.ErrUnspecified,
			want:     "ErrUnspecified",
		},
		{
			name:     "ErrAssertionFailed",
			exitCode: exitcode.ErrAssertionFailed,
			want:     "ErrAssertionFailed",
		},
		{
			name:     "ErrReadOnly",
			exitCode: exitcode.ErrReadOnly,
			want:     "ErrReadOnly",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetExitCodeStatus(tt.exitCode); got != tt.want {
				t.Errorf("GetEexitcode Status() = %v, want %v", got, tt.want)
			}
		})
	}
}
