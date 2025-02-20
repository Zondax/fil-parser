package metrics

import "github.com/zondax/golem/pkg/metrics"

var (
	_ metrics.TaskMetrics = &Client{}
	_ MetricsClient       = &Client{}
)

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
	return &Client{taskMetrics}
}

func (c Client) RegisterCustomMetrics(customMetrics ...Metric) {
	for _, metric := range customMetrics {
		// TODO: do something with the errors
		_ = c.RegisterMetric(metric.Name, metric.Help, metric.Labels, metric.Handler)
	}
}
