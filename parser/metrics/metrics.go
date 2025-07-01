package metrics

import (
	"regexp"
	"strconv"

	"github.com/zondax/fil-parser/metrics"
	"github.com/zondax/golem/pkg/metrics/collectors"
)

// Metrics names
const (
	// parser metrics
	parseMetadata              = "fil-parser_parser_parse_tx_metadata_error"
	parseMethodName            = "fil-parser_parser_parse_tx_methodName_error"
	blockCidFromMsgCid         = "fil-parser_parser_apptools_block_cid_from_msg_cid_error"
	buildCidFromMsgTrace       = "fil-parser_parser_tools_build_cid_from_msg_trace_error"
	getBlockMiner              = "fil-parser_parser_get_block_miner_error"
	jsonMarshal                = "fil-parser_parser_json_marshal_error"
	translateTxCidToTxHash     = "fil-parser_parser_translate_tx_cid_to_tx_hash_error"
	mismatchExitCode           = "fil-parser_parser_mismatch_exit_code_error"
	traceWithoutMessage        = "fil-parser_parser_trace_without_message_error"
	traceWithoutExecutionTrace = "fil-parser_parser_trace_without_execution_trace_error"
	parseTraceError            = "fil-parser_parser_parse_trace_error"

	// helper metrics
	parseActorName    = "fil-parser_helper_actor_name_error"
	parseAddress      = "fil-parser_helper_address_error"
	getEvmSelectorSig = "fil-parser_helper_get_evm_selector_sig_error"

	parseNativeEventsLog = "fil-parser_parser_parse_native_events_error"

	parseEthLog = "fil-parser_parser_parse_eth_error"
)

// Metrics labels
const (
	// errorLabel   = "error"
	actorLabel          = "actor"
	txTypeLabel         = "txType"
	codeLabel           = "code"
	kindLabel           = "kind"
	addressLabel        = "address"
	subcallSuccessLabel = "subcallSuccess"
	mainSuccessLabel    = "mainSuccess"
)

// Patterns to normalize error messages
// TODO: this is a hack to reduce metric cardinality, we should find a better solution in the future
var (
	// get evm selector rules
	// errFrom4BytesPattern  = regexp.MustCompile(`error from 4bytes: .*`)
	// errSigNotFoundPattern = regexp.MustCompile(`signature not found: .*`)
	// errCacheStorePattern  = regexp.MustCompile(`error adding selector_sig to cache: .*`)

	// miner
	errBlockMinedByNotFoundPattern = regexp.MustCompile(`could not find block mined by miner '[^']+'`)

	// helper actor name rules
	errResolutionLookupPattern = regexp.MustCompile(`resolution lookup failed \([^)]+\): resolve address [^:]+: actor not found`)
	errBadAddressPattern       = regexp.MustCompile(`address [^ ]+ is flagged as bad`)
)

var (
	parsingMetadataErrorMetric = metrics.Metric{
		Name:    parseMetadata,
		Help:    "parsing metadata error",
		Labels:  []string{actorLabel, txTypeLabel, subcallSuccessLabel, mainSuccessLabel},
		Handler: &collectors.Gauge{},
	}

	parsingMethodNameMetric = metrics.Metric{
		Name:    parseMethodName,
		Help:    "parsing method name",
		Labels:  []string{actorLabel, codeLabel, subcallSuccessLabel, mainSuccessLabel},
		Handler: &collectors.Gauge{},
	}

	parsingBlockCidFromMsgCidMetric = metrics.Metric{
		Name:    blockCidFromMsgCid,
		Help:    "get block cid from message cid",
		Labels:  []string{txTypeLabel},
		Handler: &collectors.Gauge{},
	}

	parsingBuildCidFromMsgTraceMetric = metrics.Metric{
		Name:    buildCidFromMsgTrace,
		Help:    "build Cid From Message Trace",
		Labels:  []string{txTypeLabel},
		Handler: &collectors.Gauge{},
	}

	parsingGetBlockMinerMetric = metrics.Metric{
		Name:    getBlockMiner,
		Help:    "get Block Miner",
		Labels:  []string{codeLabel, txTypeLabel},
		Handler: &collectors.Gauge{},
	}

	parsingJsonMarshalMetric = metrics.Metric{
		Name:    jsonMarshal,
		Help:    "error while marshalling json",
		Labels:  []string{kindLabel, txTypeLabel},
		Handler: &collectors.Gauge{},
	}

	parsingTranslateTxCidToTxHashMetric = metrics.Metric{
		Name:    translateTxCidToTxHash,
		Help:    "error while translate tx cid to tx hash",
		Labels:  []string{},
		Handler: &collectors.Gauge{},
	}

	parsingMismatchExitCodeMetric = metrics.Metric{
		Name:    mismatchExitCode,
		Help:    "mismatch exit code",
		Labels:  []string{},
		Handler: &collectors.Gauge{},
	}

	parsingTraceWithoutMessageMetric = metrics.Metric{
		Name:    traceWithoutMessage,
		Help:    "trace without message",
		Labels:  []string{},
		Handler: &collectors.Gauge{},
	}

	parsingTraceWithoutExecutionTraceMetric = metrics.Metric{
		Name:    traceWithoutExecutionTrace,
		Help:    "trace without execution trace",
		Labels:  []string{},
		Handler: &collectors.Gauge{},
	}

	parsingParseTraceMetric = metrics.Metric{
		Name:    parseTraceError,
		Help:    "error parsing trace",
		Labels:  []string{},
		Handler: &collectors.Gauge{},
	}

	parsingActorNameMetric = metrics.Metric{
		Name:    parseActorName,
		Help:    "get actor name from address",
		Labels:  []string{codeLabel},
		Handler: &collectors.Gauge{},
	}

	parsingAddressMetric = metrics.Metric{
		Name:    parseAddress,
		Help:    "parse address",
		Labels:  []string{addressLabel},
		Handler: &collectors.Gauge{},
	}

	getEvmSelectorSigMetric = metrics.Metric{
		Name:    getEvmSelectorSig,
		Help:    "get evm selector signature",
		Labels:  []string{},
		Handler: &collectors.Gauge{},
	}

	parsingParseNativeEventsLogMetric = metrics.Metric{
		Name:    parseNativeEventsLog,
		Help:    "parse native log",
		Labels:  []string{},
		Handler: &collectors.Gauge{},
	}

	parsingParseEthLogMetric = metrics.Metric{
		Name:    parseEthLog,
		Help:    "parse eth log",
		Labels:  []string{},
		Handler: &collectors.Gauge{},
	}
)

func (c *ParserMetricsClient) UpdateMetadataErrorMetric(actor, txType string, subcallSuccess, mainSuccess bool) error {
	// TODO: remove once errors are normalize
	// errMsg := err.Error()
	// switch {
	// case errResolutionLookupPattern.MatchString(errMsg):
	// 	errMsg = "resolution lookup failed: actor not found"
	// case errBadAddressPattern.MatchString(errMsg):
	// 	errMsg = "address is flagged as bad"
	// }

	subcallStatusStr := strconv.FormatBool(subcallSuccess)
	mainStatusStr := strconv.FormatBool(mainSuccess)

	return c.IncrementMetric(parseMetadata, actor, txType, subcallStatusStr, mainStatusStr)
}

func (c *ParserMetricsClient) UpdateMethodNameErrorMetric(actorName, code string, subcallSuccess, mainSuccess bool) error {
	subcallStatusStr := strconv.FormatBool(subcallSuccess)
	mainStatusStr := strconv.FormatBool(mainSuccess)

	return c.IncrementMetric(parseMethodName, actorName, code, subcallStatusStr, mainStatusStr)
}

func (c *ParserMetricsClient) UpdateActorNameErrorMetric(code string) error {
	// TODO: remove once errors are normalize
	// errMsg := err.Error()
	// switch {
	// case errResolutionLookupPattern.MatchString(errMsg):
	// 	errMsg = "resolution lookup failed: actor not found"
	// case errBadAddressPattern.MatchString(errMsg):
	// 	errMsg = "address is flagged as bad"
	// }

	return c.IncrementMetric(parseActorName, code)
}

func (c *ParserMetricsClient) UpdateParseAddressErrorMetric(code string) error {
	return c.IncrementMetric(parseAddress, code)
}

func (c *ParserMetricsClient) UpdateGetEvmSelectorSigMetric() error {
	// TODO: remove once errors are normalize
	// errMsg := err.Error()
	// switch {
	// case errFrom4BytesPattern.MatchString(errMsg):
	// 	errMsg = "error from 4bytes"
	// case errSigNotFoundPattern.MatchString(errMsg):
	// 	errMsg = "signature not found"
	// case errCacheStorePattern.MatchString(errMsg):
	// 	errMsg = "error adding selector_sig to cache"
	// }

	return c.IncrementMetric(getEvmSelectorSig)
}

func (c *ParserMetricsClient) UpdateBlockCidFromMsgCidMetric(txType string) error {
	// TODO: remove once errors are normalize
	// errMsg := err.Error()
	// if errBlockMinedByNotFoundPattern.MatchString(errMsg) {
	// 	errMsg = "block miner not found"
	// }

	return c.IncrementMetric(blockCidFromMsgCid, txType)
}

func (c *ParserMetricsClient) UpdateBuildCidFromMsgTraceMetric(txType string) error {
	return c.IncrementMetric(buildCidFromMsgTrace, txType)
}

func (c *ParserMetricsClient) UpdateGetBlockMinerMetric(code, txType string) error {
	return c.IncrementMetric(getBlockMiner, code, txType)
}

func (c *ParserMetricsClient) UpdateJsonMarshalMetric(kind, txType string) error {
	return c.IncrementMetric(jsonMarshal, kind, txType)
}

func (c *ParserMetricsClient) UpdateTranslateTxCidToTxHashMetric() error {
	return c.IncrementMetric(translateTxCidToTxHash)
}

func (c *ParserMetricsClient) UpdateParseNativeEventsLogsMetric() error {
	return c.IncrementMetric(parseNativeEventsLog)
}

func (c *ParserMetricsClient) UpdateParseEthLogMetric() error {
	return c.IncrementMetric(parseEthLog)
}

func (c *ParserMetricsClient) UpdateTraceWithoutMessageMetric() error {
	return c.IncrementMetric(traceWithoutMessage)
}

func (c *ParserMetricsClient) UpdateMismatchExitCodeMetric() error {
	return c.IncrementMetric(mismatchExitCode)
}

func (c *ParserMetricsClient) UpdateTraceWithoutExecutionTraceMetric() error {
	return c.IncrementMetric(traceWithoutExecutionTrace)
}

func (c *ParserMetricsClient) UpdateParseTraceMetric() error {
	return c.IncrementMetric(parseTraceError)
}
