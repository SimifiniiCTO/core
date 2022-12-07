// Copyright (C) Simfiny, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

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

	// Tracks the number of db operations performed
	DbOperationCounter *Metric

	// Tracks the latency of various db operations
	DbOperationLatency *Metric

	// Tracks the number of grpc requests partitioned by name and status code
	// used for monitoring and alerting (RED method)
	GrpcRequestCounter *Metric

	// Tracks the latency associated with grpc requests partitioned by service name, target name,
	// status code, and latency
	GrpcRequestLatency *Metric
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
		DbOperationCounter:         newDbOperationCounter(*serviceName),
		DbOperationLatency:         newDbOperationLatency(*serviceName),
		GrpcRequestCounter:         newGrpcRequestCounter(*serviceName),
		GrpcRequestLatency:         newGrpcRequestLatency(*serviceName),
	}, nil
}
