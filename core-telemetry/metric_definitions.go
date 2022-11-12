package metrics

import (
	"fmt"
)

// > This function creates a new metric that tracks the number of requests serviced by the service
// partitioned by name and status code
func newRequestCountMetric(serviceName *string) *Metric {
	return &Metric{
		MetricName:  fmt.Sprintf("%s.grpc.request.count", *serviceName),
		ServiceName: *serviceName,
		Help:        "Tracks the number of request serviced by the service partitioned by name and status code",
		Subsystem:   GrpcSubSystem,
		Namespace:   RequestNamespace,
	}
}

// `newRequestLatencyMetric` creates a new `Metric` struct with the `MetricName` set to
// `<serviceName>.grpc.request.latency`, the `ServiceName` set to `<serviceName>`, the `Help` set to
// `Tracks the latency associated with various requests partitioned by service name, target name,
// status code, and latency`, the `Subsystem` set to `grpc`, and the `Namespace` set to `request`
func newRequestLatencyMetric(serviceName *string) *Metric {
	return &Metric{
		MetricName:  fmt.Sprintf("%s.grpc.request.latency", *serviceName),
		ServiceName: *serviceName,
		Help:        "Tracks the latency associated with various requests partitioned by service name, target name, status code, and latency",
		Subsystem:   GrpcSubSystem,
		Namespace:   RequestNamespace,
	}
}

// > This function creates a new metric that tracks the number of errors encountered by the service
func newErrorCountMetric(serviceName *string) *Metric {
	return &Metric{
		MetricName:  fmt.Sprintf("%s.error.count", *serviceName),
		ServiceName: *serviceName,
		Help:        "Tracks the number of errors encountered by the service",
		Subsystem:   ErrorSubSystem,
		Namespace:   ServiceNamespace,
	}
}

// `newRequestStatusSummaryMetric` creates a new `Metric` struct with the `MetricName` set to
// `serviceName.grpc.request.summary`, the `ServiceName` set to `serviceName`, the `Help` set to
// `"Tracks the status of all requests serviced by the service"`, the `Subsystem` set to
// `RequestNamespace`, and the `Namespace` set to `RequestNamespace`
func newRequestStatusSummaryMetric(serviceName *string) *Metric {
	return &Metric{
		MetricName:  fmt.Sprintf("%s.grpc.request.summary", *serviceName),
		ServiceName: *serviceName,
		Help:        "Tracks the status of all requests serviced by the service",
		Subsystem:   Subsystem(RequestNamespace),
		Namespace:   RequestNamespace,
	}
}
