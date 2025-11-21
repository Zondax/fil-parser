package evm

import (
	"github.com/zondax/fil-parser/metrics"
	metrics2 "github.com/zondax/golem/pkg/metrics"
	"github.com/zondax/golem/pkg/metrics/collectors"
)

var (
	_ metrics.MetricsClient = &evmMetricsClient{}
	_ metrics2.TaskMetrics  = &evmMetricsClient{}
)

const parserModule = "parser_module"

// Labels const
const (
// errorLabel  = "error"
// txTypeLabel = "txType"
)

type evmMetricsClient struct {
	metrics.MetricsClient
	name string
}

func newClient(metricsClient metrics.MetricsClient, name string) *evmMetricsClient {
	s := &evmMetricsClient{
		MetricsClient: metricsClient,
		name:          name,
	}

	s.registerModuleMetrics(decodeParamErrorMetric, decodeReturnErrorMetric, decodeSelectorErrorMetric, unpackInputsErrorMetric, unpackOutputsErrorMetric)

	return s
}

const (
	decodeParam    = "fil-parser_evm_decode_param_error"
	decodeReturn   = "fil-parser_evm_decode_return_error"
	decodeSelector = "fil-parser_evm_decode_selector_error"
	unpackInputs   = "fil-parser_evm_unpack_inputs_error"
	unpackOutputs  = "fil-parser_evm_unpack_outputs_error"
)

var (
	decodeParamErrorMetric = metrics.Metric{
		Name:    decodeParam,
		Help:    "decode transaction parameters",
		Labels:  []string{},
		Handler: &collectors.Gauge{},
	}
	decodeReturnErrorMetric = metrics.Metric{
		Name:    decodeParam,
		Help:    "decode transaction return values",
		Labels:  []string{},
		Handler: &collectors.Gauge{},
	}
	decodeSelectorErrorMetric = metrics.Metric{
		Name:    decodeParam,
		Help:    "decode evm function selector",
		Labels:  []string{},
		Handler: &collectors.Gauge{},
	}
	unpackInputsErrorMetric = metrics.Metric{
		Name:    decodeParam,
		Help:    "unpack evm function inputs",
		Labels:  []string{},
		Handler: &collectors.Gauge{},
	}
	unpackOutputsErrorMetric = metrics.Metric{
		Name:    decodeParam,
		Help:    "unpack evm function outputs",
		Labels:  []string{},
		Handler: &collectors.Gauge{},
	}
)

func (c *evmMetricsClient) registerModuleMetrics(metrics ...metrics.Metric) {
	commonLabels := []string{parserModule}
	for i := range metrics {
		metrics[i].Labels = append(metrics[i].Labels, commonLabels...)
	}

	c.RegisterCustomMetrics(metrics...)
}

func (c *evmMetricsClient) IncrementMetric(name string, labels ...string) error {
	labels = append(labels, c.name)
	return c.MetricsClient.IncrementMetric(name, labels...)
}

func (c *evmMetricsClient) UpdateMetric(name string, value float64, labels ...string) error {
	labels = append(labels, c.name)
	return c.MetricsClient.UpdateMetric(name, value, labels...)
}

func (c *evmMetricsClient) UpdateDecodeParamMetric() error {
	return c.IncrementMetric(decodeParam)
}

func (c *evmMetricsClient) UpdateDecodeReturnMetric() error {
	return c.IncrementMetric(decodeReturn)
}

func (c *evmMetricsClient) UpdateDecodeSelectorMetric() error {
	return c.IncrementMetric(decodeSelector)
}

func (c *evmMetricsClient) UpdateUnpackInputsMetric() error {
	return c.IncrementMetric(unpackInputs)
}

func (c *evmMetricsClient) UpdateUnpackOutputsMetric() error {
	return c.IncrementMetric(unpackOutputs)
}
