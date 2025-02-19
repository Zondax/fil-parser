package metrics

import (
	"github.com/zondax/fil-parser/metrics"
)

type ActorsMetricsClient struct {
	metrics.MetricsClient
}

func NewClient(metricsClient metrics.MetricsClient) *ActorsMetricsClient {
	s := &ActorsMetricsClient{
		MetricsClient: metricsClient,
	}

	s.registerModuleMetrics()

	return s
}

func (c *ActorsMetricsClient) registerModuleMetrics() {
	c.RegisterCustomMetrics(parseActorMethodErrorMetric)
}
