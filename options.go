package fil_parser

import (
	metrics2 "github.com/zondax/fil-parser/metrics"
	"github.com/zondax/golem/pkg/metrics"
)

type FilecoinParserOptions struct {
	metrics metrics2.MetricsClient
}

type Option func(*FilecoinParserOptions)

func WithMetrics(metrics metrics.TaskMetrics) Option {
	return func(o *FilecoinParserOptions) {
		o.metrics = metrics2.NewMetricsClient(metrics)
	}
}
