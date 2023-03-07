// Copyright (C) Simfiny, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package metrics

import (
	"context"
	"fmt"

	"github.com/newrelic/go-agent/v3/newrelic"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Telemetry struct {
	Engine *MetricsEngine
}

type ServiceTelemetry interface {
	// ErrorCountUnaryServerInterceptor Create new unary server interceptor to capture number of errors.
	ErrorCountUnaryServerInterceptor() grpc.UnaryServerInterceptor
	// ErrorCountStreamServerInterceptor Create new stream server interceptor to capture error count over time.
	ErrorCountStreamServerInterceptor() grpc.StreamServerInterceptor
	// RequestLatencyUnaryServerInterceptor Create new unary server interceptor to capture request latency over time.
	RequestLatencyUnaryServerInterceptor() grpc.UnaryServerInterceptor
	// RequestLatencyStreamServerInterceptor Create new stream server interceptor to capture request latency over time.
	RequestLatencyStreamServerInterceptor() grpc.StreamServerInterceptor
	// RequestCountUnaryServerInterceptor Create new unary server interceptor to capture request count over time.
	RequestCountUnaryServerInterceptor() grpc.UnaryServerInterceptor
	// RequestCountStreamServerInterceptor Create new stream server interceptor to capture request count over time.
	RequestCountStreamServerInterceptor() grpc.StreamServerInterceptor
	// RequestTimeUnaryServerInterceptor Create new unary server interceptor to capture request time.
	RequestTimeUnaryServerInterceptor() grpc.UnaryServerInterceptor
	// RequestTimeStreamServerInterceptor Create new stream server interceptor to capture request time.
	RequestTimeStreamServerInterceptor() grpc.StreamServerInterceptor
	// TxnUnaryServerInterceptor Create new unary server interceptor to propagate the transaction across requests automatically.
	TxnUnaryServerInterceptor() grpc.UnaryServerInterceptor
	// TxnStreamServerInterceptor Create new stream server interceptor to propagate the transaction across requests automatically.
	TxnStreamServerInterceptor() grpc.StreamServerInterceptor
	// NewTransaction Create new transaction for the given service call
	NewTransaction(ctx context.Context, meta *TxMetadata) *newrelic.Transaction
	// New Segment Create new segment for the given service call
	NewSegment(txn *newrelic.Transaction, meta *TxMetadata) *newrelic.Segment
	// TxFromContext returns a transaction from the context if it exists, otherwise it creates a new transaction
	TxFromContext(ctx context.Context, meta *TxMetadata, sdk *newrelic.Application) *newrelic.Transaction
	// SpanName returns a formatted span name
	SpanName(meta *TxMetadata) string
	// TxName returns a formatted transaction name
	TxName(meta *TxMetadata) string
}

type TelemetryParams struct {
	ServiceName string
	// The version of the service actively deployed
	Version string
	// The service P.O.
	PointOfContact string
	// A link to documentation around the service's functionality and uses
	DocumentationLink string
	// The environment in which the service is actively running and deployed in
	Environment string
	// license key for interactions with the new-relic platform
	NewRelicLicenseKey string
	// logger instance used by the newrelic client
	Logger *zap.Logger
	// NewRelic telemetry client
	Client *newrelic.Application
	// Enable metric reporting
	Enabled bool
}

func New(params *TelemetryParams) (*Telemetry, error) {
	if params == nil {
		return nil, fmt.Errorf("invalid input argument. params cannot be nil")
	}

	engine, err := newMetricsEngine(params)
	if err != nil {
		return nil, err
	}

	return &Telemetry{
		Engine: engine,
	}, nil

}
