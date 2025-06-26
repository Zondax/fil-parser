package parser

type Config struct {
	FeesAsColumn                  bool
	ConsolidateRobustAddress      bool
	RobustAddressBestEffort       bool
	NodeMaxRetries                int
	NodeMaxWaitBeforeRetrySeconds int
	// linear, exponential default: linear
	NodeRetryStrategy string
	// Height at which to start translating txcids to txhashes
	TxCidTranslationStart int64
}
