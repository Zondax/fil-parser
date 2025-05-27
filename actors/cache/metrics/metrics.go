package metrics

import (
	"strconv"

	"github.com/zondax/fil-parser/metrics"
	metrics2 "github.com/zondax/golem/pkg/metrics"
	"github.com/zondax/golem/pkg/metrics/collectors"
)

var (
	_ metrics.MetricsClient = &ActorsCacheMetricsClient{}
	_ metrics2.TaskMetrics  = &ActorsCacheMetricsClient{}
)

const parserModule = "parser_module"

const (
	nodeApiCall = "fil-parser_node_api_call_error"

	// Metrics labels
	requestNameLabel = "requestName"
	successLabel     = "success"
	isRetryLabel     = "isRetry"
	isRetriableLabel = "isRetriable"
)

// metrics
var (
	nodeApiCallMetric = metrics.Metric{
		Name:    nodeApiCall,
		Help:    "Node API call",
		Labels:  []string{requestNameLabel, successLabel, isRetryLabel, isRetriableLabel},
		Handler: &collectors.Gauge{},
	}
)

type ActorsCacheMetricsClient struct {
	metrics.MetricsClient
	name string
}

func NewClient(metricsClient metrics.MetricsClient, name string) *ActorsCacheMetricsClient {
	s := &ActorsCacheMetricsClient{
		MetricsClient: metricsClient,
		name:          name,
	}

	s.registerModuleMetrics(nodeApiCallMetric)

	return s
}

func (c *ActorsCacheMetricsClient) registerModuleMetrics(metrics ...metrics.Metric) {
	commonLabels := []string{parserModule}
	for i := range metrics {
		metrics[i].Labels = append(metrics[i].Labels, commonLabels...)
	}

	c.RegisterCustomMetrics(metrics...)
}

func (c *ActorsCacheMetricsClient) IncrementMetric(name string, labels ...string) error {
	labels = append(labels, c.name)
	return c.MetricsClient.IncrementMetric(name, labels...)
}

func (c *ActorsCacheMetricsClient) UpdateNodeApiCallMetric(requestName string, success, isRetry, isRetriable bool) error {
	labels := []string{requestName, strconv.FormatBool(success), strconv.FormatBool(isRetry), strconv.FormatBool(isRetriable)}
	return c.IncrementMetric(nodeApiCall, labels...)
}
