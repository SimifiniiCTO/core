// Copyright (C) Simfiny, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package metrics

import (
	"context"
	"time"

	rkgrpcctx "github.com/rookie-ninja/rk-grpc/interceptor/context"
	rkgrpcmid "github.com/rookie-ninja/rk-grpc/v2/middleware"
	"google.golang.org/grpc"
)

// RequestLatencyUnaryServerInterceptor Create new unary server interceptor to capture request latency.
func (c *Telemetry) RequestLatencyUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()

		ctx = rkgrpcmid.WrapContextForServer(ctx)
		method, path, _, _ := rkgrpcmid.GetGwInfo(rkgrpcctx.GetIncomingHeaders(ctx))
		resp, err := handler(ctx, req)

		c.Engine.recordLatencyMetric(c.Metrics.RequestLatencyMetric, method, path, time.Since(start))
		return resp, err
	}
}

// RequestLatencyStreamServerInterceptor Create new stream server interceptor to capture request latency.
func (c *Telemetry) RequestLatencyStreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		// Before invoking
		wrappedStream := rkgrpcctx.WrapServerStream(stream)
		wrappedStream.WrappedContext = rkgrpcmid.WrapContextForServer(wrappedStream.WrappedContext)

		start := time.Now()

		method, path, _, _ := rkgrpcmid.GetGwInfo(rkgrpcctx.GetIncomingHeaders(wrappedStream.WrappedContext))
		err := handler(srv, wrappedStream)

		c.Engine.recordLatencyMetric(c.Metrics.RequestLatencyMetric, method, path, time.Since(start))
		return err
	}
}
