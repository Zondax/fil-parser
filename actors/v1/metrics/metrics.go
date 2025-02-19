package metrics

import (
	"github.com/zondax/fil-parser/metrics"
	"github.com/zondax/golem/pkg/metrics/collectors"
)

const (
	deserializeRawParams = "parsing_actors_rawParams_deserialize_error"
	deserializeRawReturn = "parsing_actors_rawReturn_deserialize_error"
)

var (
	deserializeRawParamsMetric = metrics.Metric{
		Name:    deserializeRawParams,
		Help:    "Parsing metadata error",
		Labels:  []string{"method", "Error"},
		Handler: &collectors.Gauge{},
	}

	deserializeRawReturnMetric = metrics.Metric{
		Name:    deserializeRawReturn,
		Help:    "Parsing metadata error",
		Labels:  []string{"method", "Error"},
		Handler: &collectors.Gauge{},
	}
)

func (c *ActorsMetricsClient) UpdateDeserializeRawParamsErrorMetric(method string, err error) error {
	return c.IncrementMetric(deserializeRawParams, method, err.Error())
}

func (c *ActorsMetricsClient) UpdateDeserializeRawReturnErrorMetric(method string, err error) error {
	return c.IncrementMetric(deserializeRawReturn, method, err.Error())
}
