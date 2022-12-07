// Copyright (C) Simfiny, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

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
		Subsystem:   Subsystem(GrpcSubSystem),
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
		Subsystem:   Subsystem(GrpcSubSystem),
		Namespace:   RequestNamespace,
	}
}

// > This function creates a new metric that tracks the number of errors encountered by the service
func newErrorCountMetric(serviceName *string) *Metric {
	return &Metric{
		MetricName:  fmt.Sprintf("%s.error.count", *serviceName),
		ServiceName: *serviceName,
		Help:        "Tracks the number of errors encountered by the service",
		Subsystem:   Subsystem(ErrorSubSystem),
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

// NewDbOperationCounter instantiates a new metric around tracking the number of db requests made by the service
func newDbOperationCounter(serviceName string) *Metric {
	return &Metric{
		MetricName:  fmt.Sprintf("%s.db.operation.counter", serviceName),
		ServiceName: serviceName,
		Help:        "Tracks the number of db tx processed by the service",
		Subsystem:   Subsystem(DbSubSystem),
		Namespace:   DatabaseNamespace,
	}
}

// NewGrpcRequestLatency instantiates a new metric object around tracking the latency associated with various gRPC operations
func newDbOperationLatency(serviceName string) *Metric {
	return &Metric{
		MetricName:  fmt.Sprintf("%s.db.operation.latency", serviceName),
		ServiceName: serviceName,
		Help:        "Tracks the latency of all db tx performed by the service.",
		Subsystem:   Subsystem(DbSubSystem),
		Namespace:   DatabaseNamespace,
	}
}

// NewGrpcRequestCounter instantiates a new metric around tracking the number of grpc requests made by the service
func newGrpcRequestCounter(serviceName string) *Metric {
	return &Metric{
		MetricName:  fmt.Sprintf("%s.grpc.request.counter", serviceName),
		ServiceName: serviceName,
		Help:        "Tracks the number of grpc requests processed by the service. Partitioned by status code and operation",
		Subsystem:   Subsystem(GrpcSubSystem),
		Namespace:   RequestNamespace,
	}
}

// NewGrpcRequestLatency instantiates a new metric object around tracking the latency associated with various gRPC operations
func newGrpcRequestLatency(serviceName string) *Metric {
	return &Metric{
		MetricName:  fmt.Sprintf("%s.grpc.request.latency", serviceName),
		ServiceName: serviceName,
		Help:        "Tracks the latency of all outgoing grpc requests initiated by the service. Partitioned by status code and operation",
		Subsystem:   Subsystem(GrpcSubSystem),
		Namespace:   RequestNamespace,
	}
}
