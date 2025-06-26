package fil_parser

import (
	"time"

	metrics2 "github.com/zondax/fil-parser/metrics"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/golem/pkg/metrics"
	golemBackoff "github.com/zondax/golem/pkg/zhttpclient/backoff"
)

// FilecoinParserOptions contains configuration options for the Filecoin parser.
type FilecoinParserOptions struct {
	// metrics is the metrics client used to track parser metrics and statistics.
	metrics metrics2.MetricsClient
	config  parser.Config
	backoff *golemBackoff.BackOff
}

// Option is a function type that modifies FilecoinParserOptions.
// It is used to provide a flexible way to configure the Filecoin parser
// through functional options pattern.
type Option func(*FilecoinParserOptions)

// WithMetrics returns an Option that configures the metrics client for the Filecoin parser.
func WithMetrics(metrics metrics.TaskMetrics) Option {
	return func(o *FilecoinParserOptions) {
		o.metrics = metrics2.NewMetricsClient(metrics)
	}
}

// WithConfig returns an Option that configures the parser config.
func WithConfig(config parser.Config) Option {
	return func(o *FilecoinParserOptions) {
		o.config = config
	}
}

func WithBackoff(maxRetries int, maxWaitBeforeRetrySeconds int) Option {
	return func(o *FilecoinParserOptions) {
		b := golemBackoff.New().
			WithMaxAttempts(maxRetries).
			WithMaxDuration(time.Duration(maxWaitBeforeRetrySeconds) * time.Second).
			WithInitialDuration(time.Duration(maxWaitBeforeRetrySeconds) * time.Second)

		o.backoff = b
	}
}

func DefaultBackoff() *golemBackoff.BackOff {
	return golemBackoff.New().
		WithMaxAttempts(1).
		WithMaxDuration(1 * time.Second).
		WithInitialDuration(1 * time.Second)
}
