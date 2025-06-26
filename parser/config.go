package parser

import "github.com/cenkalti/backoff/v4"

type Config struct {
	FeesAsColumn                  bool
	ConsolidateRobustAddress      bool
	RobustAddressBestEffort       bool
	NodeMaxRetries                int
	NodeMaxWaitBeforeRetrySeconds int64
	// linear, exponential default: linear
	NodeRetryStrategy string
	// Height at which to start translating txcids to txhashes
	TxCidTranslationStart int64

	backoff backoff.BackOff
}
