package miner

import (
	"github.com/zondax/fil-parser/metrics"
	metrics2 "github.com/zondax/golem/pkg/metrics"
	"github.com/zondax/golem/pkg/metrics/collectors"
)

var (
	_ metrics.MetricsClient = &minerMetricsClient{}
	_ metrics2.TaskMetrics  = &minerMetricsClient{}
)

const parserModule = "parser_module"

// Labels const
const (
// errorLabel  = "error"
// txTypeLabel = "txType"
)

type minerMetricsClient struct {
	metrics.MetricsClient
	name string
}

func newClient(metricsClient metrics.MetricsClient, name string) *minerMetricsClient {
	s := &minerMetricsClient{
		MetricsClient: metricsClient,
		name:          name,
	}

	s.registerModuleMetrics(actorNameFromAddressMetric)
	s.registerModuleMetrics(processedMinerInfoTotalMetric)
	s.registerModuleMetrics(processedMinerSectorsTotalMetric)

	return s
}

const (
	actorNameFromAddress       = "fil-parser_miner_actor_name_from_address"
	processedMinerEventsTotal  = "fil-parser_miner_processed_events_total"
	processedMinerSectorsTotal = "fil-parser_miner_processed_sectors_total"
)

var (
	actorNameFromAddressMetric = metrics.Metric{
		Name:    actorNameFromAddress,
		Help:    "get actor name from address",
		Labels:  []string{},
		Handler: &collectors.Gauge{},
	}

	// Register metrics for processed miner events
	// Usage: Increment counter when a block has been processed based on the amount of events successfully processed by type
	processedMinerInfoTotalMetric = metrics.Metric{
		Name:    processedMinerEventsTotal,
		Help:    "Total number of processed miner info in indexer",
		Labels:  []string{"tx_type"},
		Handler: &collectors.Gauge{},
	}

	// Register metrics for processed miner events
	// Usage: Increment counter when a block has been processed based on the amount of events successfully processed by type
	processedMinerSectorsTotalMetric = metrics.Metric{
		Name:    processedMinerSectorsTotal,
		Help:    "Total number of processed miner events in indexer",
		Labels:  []string{"tx_type"},
		Handler: &collectors.Gauge{},
	}
)

func (c *minerMetricsClient) registerModuleMetrics(metrics ...metrics.Metric) {
	commonLabels := []string{parserModule}
	for i := range metrics {
		metrics[i].Labels = append(metrics[i].Labels, commonLabels...)
	}

	c.RegisterCustomMetrics(metrics...)
}

func (c *minerMetricsClient) IncrementMetric(name string, labels ...string) error {
	labels = append(labels, c.name)
	return c.MetricsClient.IncrementMetric(name, labels...)
}

func (c *minerMetricsClient) UpdateActorNameFromAddressMetric() error {
	return c.IncrementMetric(actorNameFromAddress)
}

func (c *minerMetricsClient) UpdateProcessedMinerInfoTotalMetric(txType string) error {
	return c.IncrementMetric(processedMinerEventsTotal, txType)
}

func (c *minerMetricsClient) UpdateProcessedMinerSectorsTotalMetric(txType string) error {
	return c.IncrementMetric(processedMinerSectorsTotal, txType)
}
