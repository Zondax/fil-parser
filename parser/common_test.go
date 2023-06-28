package parser

import (
	"testing"

	"github.com/filecoin-project/go-state-types/builtin/v11/datacap"
	"github.com/filecoin-project/go-state-types/exitcode"
)

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

func TestParseParams(t *testing.T) {
	tests := []struct {
		name     string
		metadata map[string]interface{}
		want     string
	}{
		{
			name:     "nil metadata",
			metadata: nil,
			want:     "",
		},
		{
			name: "empty params",
			metadata: map[string]interface{}{
				ParamsKey: "",
			},
			want: "",
		},
		{
			name: "string in params",
			metadata: map[string]interface{}{
				ParamsKey: "f1ljefareoomkuplzvk5zkk3cjeq25fjdbs2gwzea",
			},
			want: "f1ljefareoomkuplzvk5zkk3cjeq25fjdbs2gwzea",
		},
		{
			name: "struct in params",
			metadata: map[string]interface{}{
				ParamsKey: Propose{
					To:     "f01656666",
					Value:  "0",
					Method: "",
					Params: nil,
				},
			},
			want: "{\"To\":\"f01656666\",\"Value\":\"0\",\"Method\":\"\",\"Params\":null}",
		},
		{
			name: "cbor",
			metadata: map[string]interface{}{
				ParamsKey: datacap.TransferParams{},
			},
			want: "{\"To\":\"\\u003cempty\\u003e\",\"Amount\":\"\\u003cnil\\u003e\",\"OperatorData\":null}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseParams(tt.metadata); got != tt.want {
				t.Errorf("ParseParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseReturn(t *testing.T) {
	tests := []struct {
		name     string
		metadata map[string]interface{}
		want     string
	}{
		{
			name:     "nil metadata",
			metadata: nil,
			want:     "",
		},
		{
			name: "empty params",
			metadata: map[string]interface{}{
				ReturnKey: "",
			},
			want: "",
		},
		{
			name: "string in params",
			metadata: map[string]interface{}{
				ReturnKey: "f1ljefareoomkuplzvk5zkk3cjeq25fjdbs2gwzea",
			},
			want: "f1ljefareoomkuplzvk5zkk3cjeq25fjdbs2gwzea",
		},
		{
			name: "struct in params",
			metadata: map[string]interface{}{
				ReturnKey: Propose{
					To:     "f01656666",
					Value:  "0",
					Method: "",
					Params: nil,
				},
			},
			want: "{\"To\":\"f01656666\",\"Value\":\"0\",\"Method\":\"\",\"Params\":null}",
		},
		{
			name: "cbor",
			metadata: map[string]interface{}{
				ReturnKey: datacap.TransferParams{},
			},
			want: "{\"To\":\"\\u003cempty\\u003e\",\"Amount\":\"\\u003cnil\\u003e\",\"OperatorData\":null}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseReturn(tt.metadata); got != tt.want {
				t.Errorf("ParseParams() = %v, want %v", got, tt.want)
			}
		})
	}
}
