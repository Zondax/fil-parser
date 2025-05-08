package fil_parser

import (
	"github.com/zondax/fil-parser/actors"
	actorsV1 "github.com/zondax/fil-parser/actors/v1"
	actorsV2 "github.com/zondax/fil-parser/actors/v2"
	metrics2 "github.com/zondax/fil-parser/metrics"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/golem/pkg/logger"
	"github.com/zondax/golem/pkg/metrics"
)

// FilecoinParserOptions contains configuration options for the Filecoin parser.
type FilecoinParserOptions struct {
	// metrics is the metrics client used to track parser metrics and statistics.
	metrics                metrics2.MetricsClient
	config                 parser.Config
	actorParserConstructor actors.ActorParserConstructor
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

// WithActorParserV1 returns an Option that configures the Filecoin parser to use the V1 actor parser.
// It sets the actorParserConstructor to create a new instance of the V1 actor parser with the provided parameters.
func WithActorParserV1() Option {
	return func(o *FilecoinParserOptions) {
		o.actorParserConstructor = func(network string, helper *helper.Helper, logger *logger.Logger, metrics metrics2.MetricsClient) actors.ActorParserInterface {
			return actorsV1.NewActorParser(helper, logger, metrics)
		}
	}
}

// WithActorParserV2 returns an Option that configures the Filecoin parser to use the V2 actor parser.
// It sets the actorParserConstructor to create a new instance of the V2 actor parser with the provided parameters.
func WithActorParserV2() Option {
	return func(o *FilecoinParserOptions) {
		o.actorParserConstructor = actorsV2.NewActorParser
	}
}
