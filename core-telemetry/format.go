package metrics

import "fmt"

type MetricName struct {
	ServiceName     string
	OperationName   string
	IsDistributedTx bool
	IsDatabaseTx    bool
	IsError         bool
}

func FormatMetricName(m *MetricName, metricType MetricType) *string {
	metric := m.ServiceName
	if m.IsDistributedTx {
		metric = fmt.Sprintf("%s.dtx", metric)
	}

	if m.IsDatabaseTx {
		metric = fmt.Sprintf("%s.db", metric)
	}

	metric = fmt.Sprintf("%s.op.%s", metric, m.OperationName)
	if m.IsError {
		metric = fmt.Sprintf("%s.error", metric)
	}

	metric = AppendSuffix(metric, metricType)

	return &metric
}

func AppendSuffix(metricName string, metricType MetricType) string {
	return fmt.Sprintf("%s.%s", metricName, metricType)
}
