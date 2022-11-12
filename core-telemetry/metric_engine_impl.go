package metrics

import (
	"time"

	"google.golang.org/genproto/googleapis/rpc/code"
)

func (me *MetricsEngine) RecordStandardDtxMetrics(op, dest string, status code.Code, start time.Time) {
	metricName := FormatMetricName(&MetricName{
		ServiceName:     *me.ServiceName,
		OperationName:   op,
		IsDistributedTx: true,
		IsDatabaseTx:    false,
		IsError:         false,
	}, Latency)
	me.RecordDtxMetric(me.Metrics.RequestCountMetric, *metricName, dest, status, time.Since(start))
}

func (me *MetricsEngine) RecordRequestCountMetric(operation, destination string, status code.Code) {
	metricName := FormatMetricName(&MetricName{
		ServiceName:     *me.ServiceName,
		OperationName:   operation,
		IsDistributedTx: false,
		IsDatabaseTx:    false,
		IsError:         false,
	}, Count)
	me.RecordCounterMetric(me.Metrics.RequestCountMetric, *metricName, status)
}

func (me *MetricsEngine) RecordRequestLatencyMetric(operationName, destination string, start time.Time) {
	metricName := FormatMetricName(&MetricName{
		ServiceName:     *me.ServiceName,
		OperationName:   operationName,
		IsDistributedTx: false,
		IsDatabaseTx:    false,
		IsError:         false,
	}, Latency)
	me.RecordLatencyMetric(me.Metrics.RequestLatencyMetric, *metricName, destination, time.Since(start))
}

func (me *MetricsEngine) RecordErrorCountMetric(operation, destination string) {
	metricName := FormatMetricName(&MetricName{
		ServiceName:     *me.ServiceName,
		OperationName:   operation,
		IsDistributedTx: false,
		IsDatabaseTx:    false,
		IsError:         true,
	}, ErrorCount)
	me.RecordCounterMetric(me.Metrics.ErrorCountMetric, *metricName, code.Code_INTERNAL)
}

func (me *MetricsEngine) RecordRequestStatusSummaryMetric(operation, destination string) {
	metricName := FormatMetricName(&MetricName{
		ServiceName:     *me.ServiceName,
		OperationName:   operation,
		IsDistributedTx: false,
		IsDatabaseTx:    false,
		IsError:         false,
	}, MetricSummary)
	me.RecordSummaryMetric(me.Metrics.RequestStatusSummaryMetric, *metricName, destination)
}

func (me *MetricsEngine) RecordStandardMetrics(op string, isOperationSuccessful bool) {
	metricName := FormatMetricName(&MetricName{
		ServiceName:     *me.ServiceName,
		OperationName:   op,
		IsDistributedTx: false,
		IsDatabaseTx:    false,
		IsError:         isOperationSuccessful,
	}, Count)
	me.RecordOpCounterMetric(me.Metrics.GrpcRequestCounter, *metricName)
}
