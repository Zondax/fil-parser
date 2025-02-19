package metrics

import (
	"github.com/zondax/fil-parser/metrics"
	"github.com/zondax/golem/pkg/metrics/collectors"
)

const (
	parsingMetadata   = "parsing_tx_metadata_error"
	parsingMethodName = "parsing_tx_methodName_error"
)

var (
	parsingMetadataErrorMetric = metrics.Metric{
		Name:    parsingMetadata,
		Help:    "Parsing metadata error",
		Labels:  []string{"type", "Error"},
		Handler: &collectors.Gauge{},
	}

	parsingMethodNameMetric = metrics.Metric{
		Name:    parsingMethodName,
		Help:    "Parsing method name",
		Labels:  []string{"Error"},
		Handler: &collectors.Gauge{},
	}
)

func (c *ParserMetricsClient) UpdateMetadataErrorMetric(txType string, err error) error {
	return c.IncrementMetric(parsingMetadata, txType, err.Error())
}

func (c *ParserMetricsClient) UpdateMethodNameErrorMetric(err error) error {
	return c.IncrementMetric(parsingMethodName, err.Error())
}
