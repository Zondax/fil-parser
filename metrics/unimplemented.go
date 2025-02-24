package metrics

import "github.com/zondax/golem/pkg/metrics"

var (
	_ metrics.TaskMetrics = UnimplementedMetricsClient{}
	_ MetricsClient       = UnimplementedMetricsClient{}
)

// UnimplementedMetricsClient is a no-op implementation of the MetricsClient interface.
// It can be embedded in other types to get a default implementation of the interface.
type UnimplementedMetricsClient struct{}

func (u UnimplementedMetricsClient) Start() error {
	return nil
}

func (u UnimplementedMetricsClient) RegisterMetric(_, _ string, _ []string, _ metrics.MetricHandler) error {
	return nil
}

func (u UnimplementedMetricsClient) UpdateMetric(_ string, _ float64, _ ...string) error {
	return nil
}

func (u UnimplementedMetricsClient) IncrementMetric(_ string, _ ...string) error {
	return nil
}

func (u UnimplementedMetricsClient) DecrementMetric(_ string, _ ...string) error {
	return nil
}

func (u UnimplementedMetricsClient) Name() string {
	return ""
}

func (u UnimplementedMetricsClient) Stop() error {
	return nil
}
func (u UnimplementedMetricsClient) RegisterCustomMetrics(_ ...Metric) {}
