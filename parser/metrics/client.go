package metrics

import (
	"github.com/zondax/fil-parser/metrics"
)

type ParserMetricsClient struct {
	metrics.MetricsClient
	version string
}

func NewClient(metricsClient metrics.MetricsClient, version string) *ParserMetricsClient {
	s := &ParserMetricsClient{
		MetricsClient: metrics.NewMetricsClient(metricsClient),
		version:       version,
	}

	s.registerModuleMetrics()

	return s
}

func (c *ParserMetricsClient) registerModuleMetrics() {
	c.RegisterCustomMetrics(parsingMetadataErrorMetric, parsingMethodNameMetric)
}
