package tools

import (
	"errors"

	"github.com/filecoin-project/go-state-types/manifest"
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

var actorSpecificErrors = map[string]map[string]string{
	// Paych errors: https://github.com/filecoin-project/builtin-actors/blob/2f040c122bf672314450fee16aa45e8420f40dde/actors/paych/src/lib.rs#L42
	manifest.PaychKey: {
		"32": "ErrChannelStateUpdateAfterSettled",
	},
	// EVM errors: https://github.com/filecoin-project/builtin-actors/blob/2f040c122bf672314450fee16aa45e8420f40dde/actors/evm/src/lib.rs#L35
	manifest.EvmKey: {
		"33": "ErrEVMContractReverted",
		"34": "ErrEVMContractInvalidInstruction",
		"35": "ErrEVMContractUndefinedInstruction",
		"36": "ErrEVMContractStackUnderflow",
		"37": "ErrEVMContractStackOverflow",
		"38": "ErrEVMContractIllegalMemoryAccess",
		"39": "ErrEVMContractBadJumpDest",
		"40": "ErrEVMContractSelfDestructFailed",
	},

	// Miner errors: https://github.com/filecoin-project/builtin-actors/blob/2f040c122bf672314450fee16aa45e8420f40dde/actors/miner/src/lib.rs#L171
	manifest.MinerKey: {
		"1000": "ErrBalanceInvariantsBroken",
		"1001": "ErrNotificationSendFailed",
		"1002": "ErrNotificationReceiverAborted",
		"1003": "ErrNotificationResponseInvalid",
		"1004": "ErrNotificationRejected",
	},

	// Market errors: https://github.com/filecoin-project/builtin-actors/blob/2f040c122bf672314450fee16aa45e8420f40dde/actors/market/src/lib.rs#L40
	manifest.MarketKey: {
		"32": "ErrDealExpired",
		"33": "ErrDealNotActivated",
	},

	// Power errors: https://github.com/filecoin-project/builtin-actors/blob/2f040c122bf672314450fee16aa45e8420f40dde/actors/power/src/lib.rs#L60
	manifest.PowerKey: {
		"32": "ErrTooManyProveCommits",
	},
}

// CheckExitCodeError given an ExitCode.String() checks if is a common error
func CheckExitCodeError(actorName, code string) string {
	if errs, ok := actorSpecificErrors[actorName]; ok {
		if e, ok := errs[code]; ok {
			return e
		}
	}
	e, ok := commonErrors[code]
	if !ok {
		return code
	}
	return e
}
