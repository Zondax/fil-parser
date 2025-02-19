package metrics

import (
	"github.com/zondax/fil-parser/metrics"
)

type ParserMetricsClient struct {
	metrics.MetricsClient
}

func NewClient(metricsClient metrics.MetricsClient) *ParserMetricsClient {
	s := &ParserMetricsClient{
		MetricsClient: metrics.NewMetricsClient(metricsClient),
	}

	s.registerModuleMetrics()

	return s
}

func (c *ParserMetricsClient) registerModuleMetrics() {
	c.RegisterCustomMetrics(parsingMetadataErrorMetric, parsingMethodNameMetric, parsingActorNameMetric, parsingBlockCidFromMsgCidMetric)
}
