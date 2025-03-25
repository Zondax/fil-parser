package metrics

import "github.com/zondax/golem/pkg/metrics"

var (
	_ metrics.TaskMetrics = NoopMetricsClient{}
	_ MetricsClient       = NoopMetricsClient{}
)

// NoopMetricsClient is a no-op implementation of the MetricsClient interface.
// It can be embedded in other types to get a default implementation of the interface.
type NoopMetricsClient struct{}

func NewNoopMetricsClient() *NoopMetricsClient {
	return &NoopMetricsClient{}
}

func (u NoopMetricsClient) Start() error {
	return nil
}

func (u NoopMetricsClient) RegisterMetric(_, _ string, _ []string, _ metrics.MetricHandler) error {
	return nil
}

func (u NoopMetricsClient) UpdateMetric(_ string, _ float64, _ ...string) error {
	return nil
}

func (u NoopMetricsClient) IncrementMetric(_ string, _ ...string) error {
	return nil
}

func (u NoopMetricsClient) DecrementMetric(_ string, _ ...string) error {
	return nil
}

func (u NoopMetricsClient) Name() string {
	return ""
}

func (u NoopMetricsClient) Stop() error {
	return nil
}

func (u NoopMetricsClient) RegisterCustomMetrics(_ ...Metric) {}

func (u NoopMetricsClient) AppName() string {
	return ""
}
