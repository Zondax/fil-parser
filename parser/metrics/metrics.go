package metrics

import (
	"github.com/zondax/fil-parser/metrics"
	"github.com/zondax/golem/pkg/metrics/collectors"
)

const (
	parsingMetadata           = "filParser_parser_tx_metadata_error"
	parsingMethodName         = "filParser_parser_tx_methodName_error"
	parsingActorName          = "filParser_helper_actor_name_error"
	parsingBlockCidFromMsgCid = "filParser_apptools_block_cid_from_msg_cid_error"
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
		Labels:  []string{"error"},
		Handler: &collectors.Gauge{},
	}

	parsingActorNameMetric = metrics.Metric{
		Name:    parsingActorName,
		Help:    "Get actor name from address",
		Labels:  []string{"code", "error"},
		Handler: &collectors.Gauge{},
	}

	parsingBlockCidFromMsgCidMetric = metrics.Metric{
		Name:    parsingBlockCidFromMsgCid,
		Help:    "Get block cid from message cid",
		Labels:  []string{"txType", "error"},
		Handler: &collectors.Gauge{},
	}
)

func (c *ParserMetricsClient) UpdateMetadataErrorMetric(actor, txType string, err error) error {
	return c.IncrementMetric(parsingMetadata, actor, txType, err.Error())
}

func (c *ParserMetricsClient) UpdateMethodNameErrorMetric(err error) error {
	return c.IncrementMetric(parsingMethodName, err.Error())
}

func (c *ParserMetricsClient) UpdateActorNameErrorMetric(code string, err error) error {
	return c.IncrementMetric(parsingActorName, code, err.Error())
}

func (c *ParserMetricsClient) UpdateBlockCidFromMsgCidMetric(txType string, err error) error {
	return c.IncrementMetric(parsingBlockCidFromMsgCid, txType, err.Error())
}
