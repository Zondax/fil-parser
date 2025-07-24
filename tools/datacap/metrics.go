package datacap

import (
	"github.com/zondax/fil-parser/metrics"
	metrics2 "github.com/zondax/golem/pkg/metrics"
	"github.com/zondax/golem/pkg/metrics/collectors"
)

var (
	_ metrics.MetricsClient = &datacapMetricsClient{}
	_ metrics2.TaskMetrics  = &datacapMetricsClient{}
)

const parserModule = "parser_module"

// Labels const
const (
// errorLabel  = "error"
// txTypeLabel = "txType"
)

type datacapMetricsClient struct {
	metrics.MetricsClient
	name string
}

func newClient(metricsClient metrics.MetricsClient, name string) *datacapMetricsClient {
	s := &datacapMetricsClient{
		MetricsClient: metricsClient,
		name:          name,
	}

	s.registerModuleMetrics(actorNameFromAddressMetric)

	return s
}

const (
	actorNameFromAddress = "fil-parser_datacap_actor_name_from_address"
)

var (
	actorNameFromAddressMetric = metrics.Metric{
		Name:    actorNameFromAddress,
		Help:    "get actor name from address",
		Labels:  []string{},
		Handler: &collectors.Gauge{},
	}
)

func (c *datacapMetricsClient) registerModuleMetrics(metrics ...metrics.Metric) {
	commonLabels := []string{parserModule}
	for i := range metrics {
		metrics[i].Labels = append(metrics[i].Labels, commonLabels...)
	}

	c.RegisterCustomMetrics(metrics...)
}

func (c *datacapMetricsClient) IncrementMetric(name string, labels ...string) error {
	labels = append(labels, c.name)
	return c.MetricsClient.IncrementMetric(name, labels...)
}

func (c *datacapMetricsClient) UpdateMetric(name string, value float64, labels ...string) error {
	labels = append(labels, c.name)
	return c.MetricsClient.UpdateMetric(name, value, labels...)
}

func (c *datacapMetricsClient) UpdateActorNameFromAddressMetric() error {
	return c.IncrementMetric(actorNameFromAddress)
}
