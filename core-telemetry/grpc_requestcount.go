package metrics

import (
	"context"
	"fmt"

	rkgrpcctx "github.com/rookie-ninja/rk-grpc/interceptor/context"
	rkgrpcmid "github.com/rookie-ninja/rk-grpc/v2/middleware"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// RequestCountUnaryServerInterceptor Create new unary server interceptor to capture number of requests.
func (c *Telemetry) RequestCountUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		var opStatusCode codes.Code = codes.OK

		ctx = rkgrpcmid.WrapContextForServer(ctx)
		method, path, _, _ := rkgrpcmid.GetGwInfo(rkgrpcctx.GetIncomingHeaders(ctx))
		resp, err := handler(ctx, req)

		if err != nil {
			opStatusCode = codes.Internal
		}

		c.Engine.RecordRequestCountMetric(fmt.Sprintf("%s-%s", method, path), *c.Engine.ServiceName, code.Code(opStatusCode))
		return resp, err
	}
}

// RequestCountStreamServerInterceptor Create new stream server interceptor to capture the number of requests.
func (c *Telemetry) RequestCountStreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		// Before invoking
		wrappedStream := rkgrpcctx.WrapServerStream(stream)
		wrappedStream.WrappedContext = rkgrpcmid.WrapContextForServer(wrappedStream.WrappedContext)

		var opStatusCode codes.Code = codes.OK

		method, path, _, _ := rkgrpcmid.GetGwInfo(rkgrpcctx.GetIncomingHeaders(wrappedStream.WrappedContext))
		err := handler(srv, wrappedStream)

		if err != nil {
			opStatusCode = codes.Internal
		}

		c.Engine.RecordRequestCountMetric(fmt.Sprintf("%s-%s", method, path), *c.Engine.ServiceName, code.Code(opStatusCode))
		return err
	}
}