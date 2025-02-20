package multisig

import (
	"github.com/zondax/fil-parser/metrics"
	metrics2 "github.com/zondax/golem/pkg/metrics"
	"github.com/zondax/golem/pkg/metrics/collectors"
)

var (
	_ metrics.MetricsClient = &multisigMetricsClient{}
	_ metrics2.TaskMetrics  = &multisigMetricsClient{}
)

const parserModule = "parser_module"

type multisigMetricsClient struct {
	metrics.MetricsClient
	name string
}

func newClient(metricsClient metrics.MetricsClient, name string) *multisigMetricsClient {
	s := &multisigMetricsClient{
		MetricsClient: metrics.NewMetricsClient(metricsClient),
		name:          name,
	}

	s.registerModuleMetrics(
		actorNameFromAddressMetric, parseTxMetadataMetric, parseMultisigMetadataMetric, marshalMultisigMetadataMetric,
	)

	return s
}

func (c *multisigMetricsClient) registerModuleMetrics(metrics ...metrics.Metric) {
	commonLabels := []string{parserModule}
	for i := range metrics {
		metrics[i].Labels = append(metrics[i].Labels, commonLabels...)
	}

	c.RegisterCustomMetrics(metrics...)
}

func (c *multisigMetricsClient) IncrementMetric(name string, labels ...string) error {
	labels = append(labels, c.name)
	return c.MetricsClient.IncrementMetric(name, labels...)
}

const (
	actorNameFromAddress    = "fil-parser_multisig_actor_name_from_address"
	parseTxMetadata         = "fil-parser_multisig_parse_multisig_metadata"
	parseMultisigMetadata   = "fil-parser_multisig_parse_multisig_metadata"
	marshalMultisigMetadata = "fil-parser_multisig_marshal_metadata"
)

var (
	actorNameFromAddressMetric = metrics.Metric{
		Name:    actorNameFromAddress,
		Help:    "get actor name from address",
		Labels:  []string{"error"},
		Handler: &collectors.Gauge{},
	}

	parseTxMetadataMetric = metrics.Metric{
		Name:    parseTxMetadata,
		Help:    "error parsing tx metadata",
		Labels:  []string{"txType"},
		Handler: &collectors.Gauge{},
	}

	parseMultisigMetadataMetric = metrics.Metric{
		Name:    parseMultisigMetadata,
		Help:    "error parsing multisig metadata",
		Labels:  []string{"txType"},
		Handler: &collectors.Gauge{},
	}

	marshalMultisigMetadataMetric = metrics.Metric{
		Name:    marshalMultisigMetadata,
		Help:    "error marshaling multisig metadata",
		Labels:  []string{"txType"},
		Handler: &collectors.Gauge{},
	}
)

func (c *multisigMetricsClient) UpdateActorNameFromAddressMetric(err error) error {
	return c.IncrementMetric(actorNameFromAddress, err.Error())
}

func (c *multisigMetricsClient) UpdateParseTxMetadataMetric(txType string) error {
	return c.IncrementMetric(parseTxMetadata, txType)
}

func (c *multisigMetricsClient) UpdateParseMultisigMetadataMetric(txType string) error {
	return c.IncrementMetric(parseMultisigMetadata, txType)
}

func (c *multisigMetricsClient) UpdateMarshalMultisigMetadataMetric(txType string) error {
	return c.IncrementMetric(marshalMultisigMetadata, txType)
}
