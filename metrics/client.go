package metrics

import "github.com/zondax/golem/pkg/metrics"

var (
	_ metrics.TaskMetrics = &Client{}
	_ MetricsClient       = &Client{}
)

const metricLabelComponent = "component"

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
	component string
}

func NewMetricsClient(taskMetrics metrics.TaskMetrics, component string) *Client {
	return &Client{
		TaskMetrics: taskMetrics,
		component:   component,
	}
}

func NewNoopMetricsClient() *Client {
	return &Client{
		TaskMetrics: metrics.NewNoopMetrics(),
		component:   "",
	}
}

func (c *Client) RegisterCustomMetrics(customMetrics ...Metric) {
	for _, metric := range customMetrics {
		// TODO: do something with the errors
		metric.Labels = append(metric.Labels, metricLabelComponent)
		_ = c.RegisterMetric(metric.Name, metric.Help, metric.Labels, metric.Handler)
	}
}

func (c *Client) IncrementMetric(name string, labels ...string) error {
	labels = append(labels, c.component)
	return c.TaskMetrics.IncrementMetric(name, labels...)
}
