package parser

import "errors"

var (
	ErrUnknownMethod = errors.New("not known method")
	ErrBlockHash     = errors.New("unable to get block hash")
	ErrNotValidActor = errors.New("not a valid actor")
	ErrNotKnownActor = errors.New("actor is unknown")
)

type errorsMap map[string]string

var commonErrors = errorsMap{
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
}

// CheckExitCodeCommonError given an ExitCode.String() checks if is a common error
func CheckExitCodeCommonError(code string) string {
	e, ok := commonErrors[code]
	if !ok {
		return code
	}
	return e
}
