package metrics

import (
	"github.com/zondax/fil-parser/metrics"
	metrics2 "github.com/zondax/golem/pkg/metrics"
)

var (
	_ metrics.MetricsClient = &ParserMetricsClient{}
	_ metrics2.TaskMetrics  = &ParserMetricsClient{}
)

const parserModule = "parser_module"

const (
	MetadataValue     = "metadata"
	FeesMetadataValue = "feesMetadata"
)

type ParserMetricsClient struct {
	metrics.MetricsClient
	name string
}

func NewNoopMetricsClient() *ParserMetricsClient {
	return &ParserMetricsClient{
		MetricsClient: metrics.NewNoopMetricsClient(),
	}
}

func NewClient(metricsClient metrics.MetricsClient, name string) *ParserMetricsClient {
	s := &ParserMetricsClient{
		MetricsClient: metricsClient,
		name:          name,
	}

	s.registerModuleMetrics(
		parsingMetadataErrorMetric, parsingMethodNameMetric, parsingActorNameMetric, parsingBlockCidFromMsgCidMetric,
		parsingBuildCidFromMsgTraceMetric, parsingGetBlockMinerMetric, parsingJsonMarshalMetric, parsingAddressMetric,
		getEvmSelectorSigMetric, parsingTranslateTxCidToTxHashMetric, parsingParseNativeEventsLogMetric, parsingParseEthLogMetric,
		parsingMismatchExitCodeMetric, parsingTraceWithoutMessageMetric, parsingTraceWithoutExecutionTraceMetric, parsingParseTraceMetric,
	)

	return s
}

func (c *ParserMetricsClient) registerModuleMetrics(metrics ...metrics.Metric) {
	commonLabels := []string{parserModule}
	for i := range metrics {
		metrics[i].Labels = append(metrics[i].Labels, commonLabels...)
	}

	c.RegisterCustomMetrics(metrics...)
}

func (c *ParserMetricsClient) IncrementMetric(name string, labels ...string) error {
	labels = append(labels, c.name)
	return c.MetricsClient.IncrementMetric(name, labels...)
}
