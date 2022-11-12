package metrics

import "fmt"

type ServiceMetrics struct {
	// Tracks the number of request serviced by the service partitioned by name and status code
	RequestCountMetric *Metric

	// Tracks the latency associated with a various requests partitioned by service name, target name,
	// status code, and latency
	RequestLatencyMetric *Metric

	// Tracks the number of errors encountered by the service
	ErrorCountMetric *Metric

	// Tracks the status of all requests serviced by the service
	RequestStatusSummaryMetric *Metric
}

var (
	ErrInvalidServiceName error = fmt.Errorf("invalid input argument, service name cannot be nil")
)

// newServiceMetrics creates a new ServiceMetrics struct and initializes all of its fields with new metrics
// that can be emitted
func newServiceMetrics(serviceName *string) (*ServiceMetrics, error) {
	if serviceName == nil {
		return nil, ErrInvalidServiceName
	}

	return &ServiceMetrics{
		RequestCountMetric:         newRequestCountMetric(serviceName),
		RequestLatencyMetric:       newRequestLatencyMetric(serviceName),
		ErrorCountMetric:           newErrorCountMetric(serviceName),
		RequestStatusSummaryMetric: newRequestStatusSummaryMetric(serviceName),
	}, nil
}
