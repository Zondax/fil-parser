package metrics

import (
	"github.com/zondax/fil-parser/metrics"
	"github.com/zondax/golem/pkg/metrics/collectors"
)

const (
	parsingMetadata        = "filParser_parser_tx_metadata_error"
	parsingMethodName      = "filParser_parser_tx_methodName_error"
	parsingActorName       = "filParser_helper_actor_name_error"
	blockCidFromMsgCid     = "filParser_apptools_block_cid_from_msg_cid_error"
	buildCidFromMsgTrace   = "filParser_tools_build_cid_from_msg_trace_error"
	getBlockMiner          = "filParser_get_block_miner_error"
	jsonMarshal            = "filParser_json_marshal_error"
	translateTxCidToTxHash = "filParser_translate_tx_cid_to_tx_hash_error"
)

var (
	parsingMetadataErrorMetric = metrics.Metric{
		Name:    parsingMetadata,
		Help:    "Parsing metadata error",
		Labels:  []string{"actor", "txType", "error"},
		Handler: &collectors.Gauge{},
	}

	parsingMethodNameMetric = metrics.Metric{
		Name:    parsingMethodName,
		Help:    "Parsing method name",
		Labels:  []string{"code", "error"},
		Handler: &collectors.Gauge{},
	}

	parsingActorNameMetric = metrics.Metric{
		Name:    parsingActorName,
		Help:    "Get actor name from address",
		Labels:  []string{"code", "error"},
		Handler: &collectors.Gauge{},
	}

	parsingBlockCidFromMsgCidMetric = metrics.Metric{
		Name:    blockCidFromMsgCid,
		Help:    "Get block cid from message cid",
		Labels:  []string{"txType", "error"},
		Handler: &collectors.Gauge{},
	}

	parsingBuildCidFromMsgTraceMetric = metrics.Metric{
		Name:    buildCidFromMsgTrace,
		Help:    "Build Cid From Message Trace",
		Labels:  []string{"txType", "error"},
		Handler: &collectors.Gauge{},
	}

	parsingGetBlockMinerMetric = metrics.Metric{
		Name:    getBlockMiner,
		Help:    "Get Block Miner",
		Labels:  []string{"code", "txType"},
		Handler: &collectors.Gauge{},
	}

	parsingJsonMarshalMetric = metrics.Metric{
		Name:    jsonMarshal,
		Help:    "Error while marshalling json",
		Labels:  []string{"kind", "txType"},
		Handler: &collectors.Gauge{},
	}

	parsingTranslateTxCidToTxHashMetric = metrics.Metric{
		Name:    translateTxCidToTxHash,
		Help:    "Error while translate tx cid to tx hash",
		Labels:  []string{"error"},
		Handler: &collectors.Gauge{},
	}
)

func (c *ParserMetricsClient) UpdateMetadataErrorMetric(actor, txType string, err error) error {
	return c.IncrementMetric(parsingMetadata, actor, txType, err.Error())
}

func (c *ParserMetricsClient) UpdateMethodNameErrorMetric(code string, err error) error {
	return c.IncrementMetric(parsingMethodName, code, err.Error())
}

func (c *ParserMetricsClient) UpdateActorNameErrorMetric(code string, err error) error {
	return c.IncrementMetric(parsingActorName, code, err.Error())
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

func (c *ParserMetricsClient) UpdateParsingJsonMarshalMetric(kind, txType string) error {
	return c.IncrementMetric(jsonMarshal, kind, txType)
}

func (c *ParserMetricsClient) UpdateParsingTranslateTxCidToTxHashMetric(err error) error {
	return c.IncrementMetric(translateTxCidToTxHash, err.Error())
}
