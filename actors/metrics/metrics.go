package metrics

import (
	"github.com/zondax/fil-parser/metrics"
	"github.com/zondax/golem/pkg/metrics/collectors"
)

const (
	actorMethod = "filParser_actors_method_error"
)

var (
	parseActorMethodErrorMetric = metrics.Metric{
		Name:    actorMethod,
		Help:    "Parsing actor method",
		Labels:  []string{"actor", "method", "Error"},
		Handler: &collectors.Gauge{},
	}
)

func (c *ActorsMetricsClient) UpdateActorMethodErrorMetric(actor, method string, err error) error {
	return c.IncrementMetric(actorMethod, actor, method, err.Error())
}
