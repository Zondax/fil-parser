package metrics

import "github.com/zondax/golem/pkg/metrics"

var (
	_ metrics.TaskMetrics = &Client{}
	_ MetricsClient       = &Client{}
)

// MetricsClient extends the TaskMetrics interface with additional functionality for registering custom metrics.
// It provides a way to register custom metrics with labels and handlers while maintaining the base TaskMetrics capabilities.
type MetricsClient interface {
	metrics.TaskMetrics
	RegisterCustomMetrics(customMetrics ...Metric)
}

type Metric struct {
	Name    string
	Help    string
	Labels  []string
	Handler metrics.MetricHandler
}

type Client struct {
	metrics.TaskMetrics
}

func NewMetricsClient(taskMetrics metrics.TaskMetrics) *Client {
	return &Client{
		TaskMetrics: taskMetrics,
	}
}

func NewNoopMetricsClient() *Client {
	return &Client{
		TaskMetrics: metrics.NewNoopMetrics(),
	}
}

func (c *Client) RegisterCustomMetrics(customMetrics ...Metric) {
	for _, metric := range customMetrics {
		_ = c.RegisterMetric(metric.Name, metric.Help, metric.Labels, metric.Handler)
	}
}

func (c *Client) IncrementMetric(name string, labels ...string) error {
	return c.TaskMetrics.IncrementMetric(name, labels...)
}
