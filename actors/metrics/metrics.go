package metrics

import (
	"regexp"

	"github.com/zondax/fil-parser/metrics"
	"github.com/zondax/golem/pkg/metrics/collectors"
)

const actorMethod = "fil-parser_actors_method_error"

// Metrics labels
const (
	actorLabel  = "actor"
	methodLabel = "method"
	errorLabel  = "error"
)

// byteArrayTooLargeRegex matches error messages of the form "byte array too large (N)" where N is any number.
// It is used to normalize these errors by stripping out the specific size numbers and reduce cardinality.
// This error is commonly present on invokeContract.
var byteArrayTooLargeRegex = regexp.MustCompile(`byte array too large \(\d+\)`)

// metrics
var (
	parseActorMethodErrorMetric = metrics.Metric{
		Name:    actorMethod,
		Help:    "Parsing actor method",
		Labels:  []string{actorLabel, methodLabel, errorLabel}, // TODO: method for txType?
		Handler: &collectors.Gauge{},
	}
)

func (c *ActorsMetricsClient) UpdateActorMethodErrorMetric(actor, method string, err error) error {
	errString := err.Error()

	// Strip out numbers from "byte array too large" errors
	if byteArrayTooLargeRegex.MatchString(errString) {
		errString = "byte array too large"
	}

	return c.IncrementMetric(actorMethod, actor, method, errString)
}
