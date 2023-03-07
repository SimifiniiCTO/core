// Copyright (C) Simfiny, Inc. 2022-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package metrics

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"
	rkgrpcctx "github.com/rookie-ninja/rk-grpc/interceptor/context"
	rkgrpcmid "github.com/rookie-ninja/rk-grpc/v2/middleware"
	"google.golang.org/grpc"
)

// TxnUnaryServerInterceptor Create new unary server interceptor and passes txn via context.
func (c *Telemetry) TxnUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if !c.Engine.enabled {
			// telemetry is disabled, execute handler directly
			return handler(ctx, req)
		}

		// add middleware to context
		ctx = rkgrpcmid.WrapContextForServer(ctx)

		// extract request metadata
		method, path, _, _ := rkgrpcmid.GetGwInfo(rkgrpcctx.GetIncomingHeaders(ctx))

		// create New Relic transaction
		txn := newrelic.FromContext(ctx)
		if txn == nil {
			txn = c.Engine.Client.StartTransaction(method + " " + path)
			defer txn.End()
			ctx = newrelic.NewContext(ctx, txn)
		}

		// execute handler and return response
		return handler(ctx, req)
	}
}

// TxnStreamServerInterceptor Create new stream server interceptor and passes txn via context.
func (c *Telemetry) TxnStreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		// Before invoking
		if !c.Engine.enabled {
			// telemetry is disabled, execute handler directly
			return handler(srv, stream)
		}

		wrappedStream := rkgrpcctx.WrapServerStream(stream)
		wrappedStream.WrappedContext = rkgrpcmid.WrapContextForServer(wrappedStream.WrappedContext)
		
		// attempt to get txn from context
		method, path, _, _ := rkgrpcmid.GetGwInfo(rkgrpcctx.GetIncomingHeaders(wrappedStream.WrappedContext))
		txn := newrelic.FromContext(wrappedStream.WrappedContext)
		if txn == nil {
			// create a new txn if one does not already exist
			txn = c.Engine.Client.StartTransaction(method + " " + path)
		}

		defer txn.End()

		wrappedStream.WrappedContext = newrelic.NewContext(wrappedStream.WrappedContext, txn)
		return handler(srv, wrappedStream)
	}
}
