package metrics

import (
	"github.com/zondax/fil-parser/metrics"
	"github.com/zondax/golem/pkg/metrics/collectors"
)

const (
	actorMethod     = "fil-parser_actors_method_error"
	multisigPropose = "fil-parser_multisig_propose_error"

	// Metrics labels
	actorLabel     = "actor"
	methodLabel    = "method"
	methodNumLaber = "methodNum"
	// errorLabel  = "error"
)

// byteArrayTooLargeRegex matches error messages of the form "byte array too large (N)" where N is any number.
// It is used to normalize these errors by stripping out the specific size numbers and reduce cardinality.
// This error is commonly present on invokeContract.
// var byteArrayTooLargeRegex = regexp.MustCompile(`byte array too large \(\d+\)`)

// metrics
var (
	parseActorMethodErrorMetric = metrics.Metric{
		Name:    actorMethod,
		Help:    "Parsing actor method",
		Labels:  []string{actorLabel, methodLabel}, // TODO: method for txType?
		Handler: &collectors.Gauge{},
	}

	parseMultisigProposeMetric = metrics.Metric{
		Name:    multisigPropose,
		Help:    "Parsing multisig propose method",
		Labels:  []string{actorLabel, methodLabel, methodNumLaber}, // TODO: method for txType?
		Handler: &collectors.Gauge{},
	}
)

func (c *ActorsMetricsClient) UpdateActorMethodErrorMetric(actor, method string) error {
	// TODO: remove once errors are normalize
	// errString := err.Error()
	// // Strip out numbers from "byte array too large" errors
	// if byteArrayTooLargeRegex.MatchString(errString) {
	// 	errString = "byte array too large"
	// }

	return c.IncrementMetric(actorMethod, actor, method)
}

func (c *ActorsMetricsClient) UpdateMultisigProposeMetric(actor, method, methodNum string) error {
	return c.IncrementMetric(actorMethod, actor, method, methodNum)
}
