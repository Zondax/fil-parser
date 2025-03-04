package fil_parser

import (
	metrics2 "github.com/zondax/fil-parser/metrics"
	"github.com/zondax/golem/pkg/metrics"
)

// FilecoinParserOptions contains configuration options for the Filecoin parser.
type FilecoinParserOptions struct {
	// metrics is the metrics client used to track parser metrics and statistics.
	metrics metrics2.MetricsClient
}

// Option is a function type that modifies FilecoinParserOptions.
// It is used to provide a flexible way to configure the Filecoin parser
// through functional options pattern.
type Option func(*FilecoinParserOptions)

// WithMetrics returns an Option that configures the metrics client for the Filecoin parser.
func WithMetrics(metrics metrics.TaskMetrics, component string) Option {
	return func(o *FilecoinParserOptions) {
		o.metrics = metrics2.NewMetricsClient(metrics, component)
	}
}
