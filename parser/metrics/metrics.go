package metrics

import (
	"github.com/zondax/fil-parser/metrics"
	"github.com/zondax/golem/pkg/metrics/collectors"
)

const (
	// parser metrics
	parseMetadata          = "fil-parser_parser_parse_tx_metadata_error"
	parseMethodName        = "fil-parser_parser_parse_tx_methodName_error"
	blockCidFromMsgCid     = "fil-parser_parser_apptools_block_cid_from_msg_cid_error"
	buildCidFromMsgTrace   = "fil-parser_parser_tools_build_cid_from_msg_trace_error"
	getBlockMiner          = "fil-parser_parser_get_block_miner_error"
	jsonMarshal            = "fil-parser_parser_json_marshal_error"
	translateTxCidToTxHash = "fil-parser_parser_translate_tx_cid_to_tx_hash_error"

	// helper metrics
	parseActorName = "fil-parser_helper_actor_name_error"
	parseAddress   = "fil-parser_helper_address_error"

	parseNativeEventsLog = "fil-parser_parse_native_events_error"
)

var (
	parsingMetadataErrorMetric = metrics.Metric{
		Name:    parseMetadata,
		Help:    "parsing metadata error",
		Labels:  []string{"actor", "txType", "error"},
		Handler: &collectors.Gauge{},
	}

	parsingMethodNameMetric = metrics.Metric{
		Name:    parseMethodName,
		Help:    "parsing method name",
		Labels:  []string{"code", "error"},
		Handler: &collectors.Gauge{},
	}

	parsingBlockCidFromMsgCidMetric = metrics.Metric{
		Name:    blockCidFromMsgCid,
		Help:    "get block cid from message cid",
		Labels:  []string{"txType", "error"},
		Handler: &collectors.Gauge{},
	}

	parsingBuildCidFromMsgTraceMetric = metrics.Metric{
		Name:    buildCidFromMsgTrace,
		Help:    "build Cid From Message Trace",
		Labels:  []string{"txType", "error"},
		Handler: &collectors.Gauge{},
	}

	parsingGetBlockMinerMetric = metrics.Metric{
		Name:    getBlockMiner,
		Help:    "get Block Miner",
		Labels:  []string{"code", "txType"},
		Handler: &collectors.Gauge{},
	}

	parsingJsonMarshalMetric = metrics.Metric{
		Name:    jsonMarshal,
		Help:    "error while marshalling json",
		Labels:  []string{"kind", "txType"},
		Handler: &collectors.Gauge{},
	}

	parsingTranslateTxCidToTxHashMetric = metrics.Metric{
		Name:    translateTxCidToTxHash,
		Help:    "error while translate tx cid to tx hash",
		Labels:  []string{"error"},
		Handler: &collectors.Gauge{},
	}

	parsingActorNameMetric = metrics.Metric{
		Name:    parseActorName,
		Help:    "get actor name from address",
		Labels:  []string{"code", "error"},
		Handler: &collectors.Gauge{},
	}

	parsingAddress = metrics.Metric{
		Name:    parseAddress,
		Help:    "parse address",
		Labels:  []string{"address", "error"},
		Handler: &collectors.Gauge{},
	}

	parsingParseNativeEventsLogMetric = metrics.Metric{
		Name:    parseNativeEventsLog,
		Help:    "parse native log",
		Labels:  []string{"error"},
		Handler: &collectors.Gauge{},
	}
)

func (c *ParserMetricsClient) UpdateMetadataErrorMetric(actor, txType string, err error) error {
	return c.IncrementMetric(parseMetadata, actor, txType, err.Error())
}

func (c *ParserMetricsClient) UpdateMethodNameErrorMetric(code string, err error) error {
	return c.IncrementMetric(parseMethodName, code, err.Error())
}

func (c *ParserMetricsClient) UpdateActorNameErrorMetric(code string, err error) error {
	return c.IncrementMetric(parseActorName, code, err.Error())
}

func (c *ParserMetricsClient) UpdateParseAddressErrorMetric(code, err string) error {
	return c.IncrementMetric(parseAddress, code, err)
}

func (c *ParserMetricsClient) UpdateBlockCidFromMsgCidMetric(txType string, err error) error {
	return c.IncrementMetric(blockCidFromMsgCid, txType, err.Error())
}

func (c *ParserMetricsClient) UpdateBuildCidFromMsgTraceMetric(txType string, err error) error {
	return c.IncrementMetric(buildCidFromMsgTrace, txType, err.Error())
}

func (c *ParserMetricsClient) UpdateGetBlockMinerMetric(code, txType string) error {
	return c.IncrementMetric(getBlockMiner, code, txType)
}

func (c *ParserMetricsClient) UpdateJsonMarshalMetric(kind, txType string) error {
	return c.IncrementMetric(jsonMarshal, kind, txType)
}

func (c *ParserMetricsClient) UpdateTranslateTxCidToTxHashMetric(err error) error {
	return c.IncrementMetric(translateTxCidToTxHash, err.Error())
}

func (c *ParserMetricsClient) UpdateParseNativeEventsLogsMetric(err error) error {
	return c.IncrementMetric(parseNativeEventsLog, err.Error())
}
