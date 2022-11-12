package metrics

import (
	"time"

	"google.golang.org/genproto/googleapis/rpc/code"
)

func (me *MetricsEngine) RecordStandardDtxMetrics(op, dest string, status code.Code, start time.Time) {
	me.RecordDtxMetric(me.Metrics.RequestCountMetric, op, dest, status, time.Since(start))
}

func (me *MetricsEngine) RecordRequestCountMetric(operation, destination string, status code.Code) {
	me.RecordCounterMetric(me.Metrics.RequestCountMetric, operation, status)
}

func (me *MetricsEngine) RecordRequestLatencyMetric(operationName, destination string, start time.Time) {
	me.RecordLatencyMetric(me.Metrics.RequestLatencyMetric, operationName, destination, time.Since(start))
}

func (me *MetricsEngine) RecordErrorCountMetric(operation, destination string) {
	me.RecordCounterMetric(me.Metrics.ErrorCountMetric, operation, code.Code_INTERNAL)
}

func (me *MetricsEngine) RecordRequestStatusSummaryMetric(operation, destination string) {
	me.RecordSummaryMetric(me.Metrics.RequestStatusSummaryMetric, operation, destination)
}

func (me *MetricsEngine) RecordStandardMetrics(op string, isOperationSuccessful bool) {
	me.RecordOpCounterMetric(me.Metrics.GrpcRequestCounter, op)
}
