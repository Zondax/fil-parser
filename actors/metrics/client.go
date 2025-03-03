package metrics

import (
	"github.com/zondax/fil-parser/metrics"
	metrics2 "github.com/zondax/golem/pkg/metrics"
)

var (
	_ metrics.MetricsClient = &ActorsMetricsClient{}
	_ metrics2.TaskMetrics  = &ActorsMetricsClient{}
)

const parserModule = "parser_module"

type ActorsMetricsClient struct {
	metrics.MetricsClient
	name string
}

func NewClient(metricsClient metrics.MetricsClient, name string) *ActorsMetricsClient {
	s := &ActorsMetricsClient{
		MetricsClient: metricsClient,
		name:          name,
	}

	s.registerModuleMetrics(parseActorMethodErrorMetric, parseMultisigProposeMetric)

	return s
}

func (c *ActorsMetricsClient) registerModuleMetrics(metrics ...metrics.Metric) {
	commonLabels := []string{parserModule}
	for i := range metrics {
		metrics[i].Labels = append(metrics[i].Labels, commonLabels...)
	}

	c.RegisterCustomMetrics(metrics...)
}

func (c *ActorsMetricsClient) IncrementMetric(name string, labels ...string) error {
	labels = append(labels, c.name)
	return c.MetricsClient.IncrementMetric(name, labels...)
}
