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
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		var txn *newrelic.Transaction
		ctx = rkgrpcmid.WrapContextForServer(ctx)
		method, path, _, _ := rkgrpcmid.GetGwInfo(rkgrpcctx.GetIncomingHeaders(ctx))

		value := ctx.Value(TransactionContextKey)
		if value == nil {
			txn = c.Engine.Client.StartTransaction(method + " " + path)
		} else {
			txn = value.(*newrelic.Transaction)
		}

		defer txn.End()

		ctx = context.WithValue(ctx, TransactionContextKey, txn)
		return handler(ctx, req)
	}
}

// TxnStreamServerInterceptor Create new stream server interceptor and passes txn via context.
func (c *Telemetry) TxnStreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		// Before invoking
		var txn *newrelic.Transaction

		wrappedStream := rkgrpcctx.WrapServerStream(stream)
		wrappedStream.WrappedContext = rkgrpcmid.WrapContextForServer(wrappedStream.WrappedContext)

		method, path, _, _ := rkgrpcmid.GetGwInfo(rkgrpcctx.GetIncomingHeaders(wrappedStream.WrappedContext))
		value := wrappedStream.WrappedContext.Value(TransactionContextKey)

		if value == nil {
			txn = c.Engine.Client.StartTransaction(method + " " + path)
			defer txn.End()
		} else {
			txn = value.(*newrelic.Transaction)
			txn.End()
		}

		wrappedStream.WrappedContext = context.WithValue(wrappedStream.WrappedContext, TransactionContextKey, txn)
		return handler(srv, wrappedStream)
	}
}
