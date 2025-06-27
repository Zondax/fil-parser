package parser

type Config struct {
	FeesAsColumn             bool
	ConsolidateRobustAddress bool
	RobustAddressBestEffort  bool
	// NodeMaxRetries is the maximum number of retries for a node API call.
	NodeMaxRetries int
	// NodeMaxWaitBeforeRetrySeconds is the maximum wait time before retrying a node API call. (linear strategy)
	NodeMaxWaitBeforeRetrySeconds int
	// TxCidTranslationStart is the height at which to start translating txcids to txhashes
	TxCidTranslationStart int64
}
