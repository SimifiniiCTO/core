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
	me.recordDtxMetric(me.Metrics.RequestCountMetric, *metricName, dest, status, time.Since(start))
}

func (me *MetricsEngine) RecordRequestCountMetric(operation string, status code.Code) {
	metricName := FormatMetricName(&MetricName{
		ServiceName:     *me.ServiceName,
		OperationName:   operation,
		IsDistributedTx: false,
		IsDatabaseTx:    false,
		IsError:         false,
	}, Count)
	me.recordCounterMetric(me.Metrics.RequestCountMetric, *metricName, status)
}

func (me *MetricsEngine) RecordRequestLatencyMetric(operationName, destination string, start time.Time) {
	metricName := FormatMetricName(&MetricName{
		ServiceName:     *me.ServiceName,
		OperationName:   operationName,
		IsDistributedTx: false,
		IsDatabaseTx:    false,
		IsError:         false,
	}, Latency)
	me.recordLatencyMetric(me.Metrics.RequestLatencyMetric, *metricName, destination, time.Since(start))
}

func (me *MetricsEngine) RecordErrorCountMetric(operation string) {
	metricName := FormatMetricName(&MetricName{
		ServiceName:     *me.ServiceName,
		OperationName:   operation,
		IsDistributedTx: false,
		IsDatabaseTx:    false,
		IsError:         true,
	}, ErrorCount)
	me.recordCounterMetric(me.Metrics.ErrorCountMetric, *metricName, code.Code_INTERNAL)
}

func (me *MetricsEngine) RecordRequestStatusSummaryMetric(operation, destination string) {
	metricName := FormatMetricName(&MetricName{
		ServiceName:     *me.ServiceName,
		OperationName:   operation,
		IsDistributedTx: false,
		IsDatabaseTx:    false,
		IsError:         false,
	}, MetricSummary)
	me.recordSummaryMetric(me.Metrics.RequestStatusSummaryMetric, *metricName, destination)
}

func (me *MetricsEngine) RecordStandardMetrics(op string, isOperationSuccessful bool) {
	metricName := FormatMetricName(&MetricName{
		ServiceName:     *me.ServiceName,
		OperationName:   op,
		IsDistributedTx: false,
		IsDatabaseTx:    false,
		IsError:         isOperationSuccessful,
	}, Count)
	me.recordOpCounterMetric(me.Metrics.GrpcRequestCounter, *metricName)
}
