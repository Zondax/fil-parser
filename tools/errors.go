package tools

import (
	"errors"
)

var (
	ErrUnknownMethod = errors.New("not known method")
	ErrBlockHash     = errors.New("unable to get block hash")
	ErrNotValidActor = errors.New("not a valid actor")
	ErrNotKnownActor = errors.New("actor is unknown")
)

type errorsMap map[string]string

// https://github.com/filecoin-project/ref-fvm/blob/4eae3b6e8d1858abfdb82956dc8cbf082a0cac66/shared/src/error/mod.rs
var commonErrors = errorsMap{

	// System errors
	"1":  "SysErrSenderInvalid",
	"2":  "SysErrSenderStateInvalid",
	"3":  "SysErrReserved1",
	"4":  "SysErrIllegalInstruction",
	"5":  "SysErrInvalidReceiver",
	"6":  "SysErrInsufficientFunds",
	"7":  "SysErrOutOfGas",
	"8":  "SysErrReserved2",
	"9":  "SysErrIllegalExitCode",
	"10": "SysFatal",
	"11": "SysErrMissingReturn",
	"12": "SysErrReserved3",
	"13": "SysErrReserved4",
	"14": "SysErrReserved5",
	"15": "SysErrReserved6",

	// Common errors
	"16": "ErrIllegalArgument",
	"17": "ErrNotFound",
	"18": "ErrForbidden",
	"19": "ErrInsufficientFunds",
	"20": "ErrIllegalState",
	"21": "ErrSerialization",
	"22": "ErrUnhandledMessage",
	"23": "ErrUnspecified",
	"24": "ErrAssertionFailed",
	"25": "ErrReadOnly",
	"26": "ErrNotPayable",
}

// CheckExitCodeError given an ExitCode.String() checks if is a common error
func CheckExitCodeError(code string) string {
	e, ok := commonErrors[code]
	if !ok {
		return code
	}
	return e
}
