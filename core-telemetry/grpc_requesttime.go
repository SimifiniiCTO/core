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

type RequestTimeKey struct{}

// RequestCountUnaryServerInterceptor Create new unary server interceptor to capture the time of requests.
func (c *Telemetry) RequestTimeUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctx = context.WithValue(ctx, RequestTimeKey{}, time.Now().Format(time.RFC3339))
		ctx = rkgrpcmid.WrapContextForServer(ctx)
		return handler(ctx, req)
	}
}

// RequestTimeStreamServerInterceptor Create new stream server interceptor to capture the time of requests.
func (c *Telemetry) RequestTimeStreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		// Before invoking
		wrappedStream := rkgrpcctx.WrapServerStream(stream)
		wrappedStream.WrappedContext = rkgrpcmid.WrapContextForServer(wrappedStream.WrappedContext)
		wrappedStream.WrappedContext = context.WithValue(wrappedStream.WrappedContext, RequestTimeKey{}, time.Now().Format(time.RFC3339))
		return handler(srv, wrappedStream)
	}
}
