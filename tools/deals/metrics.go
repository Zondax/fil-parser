package deals

import (
	"github.com/zondax/fil-parser/metrics"
	metrics2 "github.com/zondax/golem/pkg/metrics"
	"github.com/zondax/golem/pkg/metrics/collectors"
)

var (
	_ metrics.MetricsClient = &dealsMetricsClient{}
	_ metrics2.TaskMetrics  = &dealsMetricsClient{}
)

const parserModule = "parser_module"

type dealsMetricsClient struct {
	metrics.MetricsClient
	name string
}

func newClient(metricsClient metrics.MetricsClient, name string) *dealsMetricsClient {
	s := &dealsMetricsClient{
		MetricsClient: metricsClient,
		name:          name,
	}

	s.registerModuleMetrics(actorNameFromAddressMetric, skippedActivationsMetric, skippedTerminationsMetric, skippedSectorMappingsMetric)

	return s
}

const (
	actorNameFromAddress  = "fil-parser_deals_actor_name_from_address"
	skippedActivations    = "fil-parser_deals_skipped_activations"
	skippedTerminations   = "fil-parser_deals_skipped_terminations"
	skippedSectorMappings = "fil-parser_deals_skipped_sector_mappings"
)

var (
	actorNameFromAddressMetric = metrics.Metric{
		Name:    actorNameFromAddress,
		Help:    "get actor name from address",
		Labels:  []string{},
		Handler: &collectors.Gauge{},
	}
	skippedActivationsMetric = metrics.Metric{
		Name:    skippedActivations,
		Help:    "transactions whose deal activations were skipped because their metadata was missing or unusable",
		Labels:  []string{"tx_type"},
		Handler: &collectors.Gauge{},
	}
	skippedSectorMappingsMetric = metrics.Metric{
		Name:    skippedSectorMappings,
		Help:    "transactions whose sector to deal mappings were skipped because their metadata was missing or unusable",
		Labels:  []string{"tx_type"},
		Handler: &collectors.Gauge{},
	}
	skippedTerminationsMetric = metrics.Metric{
		Name:    skippedTerminations,
		Help:    "transactions whose deal terminations were skipped because their metadata was missing or unusable",
		Labels:  []string{"tx_type"},
		Handler: &collectors.Gauge{},
	}
)

func (c *dealsMetricsClient) registerModuleMetrics(metrics ...metrics.Metric) {
	commonLabels := []string{parserModule}
	for i := range metrics {
		metrics[i].Labels = append(metrics[i].Labels, commonLabels...)
	}

	c.RegisterCustomMetrics(metrics...)
}

func (c *dealsMetricsClient) IncrementMetric(name string, labels ...string) error {
	labels = append(labels, c.name)
	return c.MetricsClient.IncrementMetric(name, labels...)
}

func (c *dealsMetricsClient) UpdateMetric(name string, value float64, labels ...string) error {
	labels = append(labels, c.name)
	return c.MetricsClient.UpdateMetric(name, value, labels...)
}

func (c *dealsMetricsClient) UpdateActorNameFromAddressMetric() error {
	return c.IncrementMetric(actorNameFromAddress)
}

func (c *dealsMetricsClient) UpdateSkippedActivationsMetric(txType string) error {
	return c.IncrementMetric(skippedActivations, txType)
}

func (c *dealsMetricsClient) UpdateSkippedTerminationsMetric(txType string) error {
	return c.IncrementMetric(skippedTerminations, txType)
}

func (c *dealsMetricsClient) UpdateSkippedSectorMappingsMetric(txType string) error {
	return c.IncrementMetric(skippedSectorMappings, txType)
}
