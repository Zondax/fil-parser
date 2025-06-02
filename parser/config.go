package parser

type Config struct {
	FeesAsColumn                  bool
	ConsolidateRobustAddress      bool
	RobustAddressBestEffort       bool
	NodeMaxRetries                int
	NodeMaxWaitBeforeRetrySeconds int64
	// linear, exponential default: linear
	NodeRetryStrategy      string
	TxHashTranslationStart int64
}
