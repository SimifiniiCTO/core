package server

import (
	"net/http"

	"google.golang.org/grpc"
	grpctrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/google.golang.org/grpc"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type TracingMiddleware struct {
	ServiceName string
}

// NewTracingMiddleware returns a new instance of the tracing middleware
func NewTracingMiddleware(serviceName string) *TracingMiddleware {
	return &TracingMiddleware{serviceName}
}

// TracerMiddleware defines a tracing middleware to be used by an application
func (mw *TracingMiddleware) TracerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		parentSpan, _ := tracer.Extract(tracer.HTTPHeadersCarrier(r.Header))
		serverSpan := tracer.StartSpan(r.URL.Path,
			tracer.ChildOf(parentSpan))
		defer serverSpan.Finish()

		r = r.WithContext(tracer.ContextWithSpan(r.Context(), serverSpan))
		next.ServeHTTP(w, r)
	})
}

// GrpcStreamInterceptorMiddleware returns a grpc server stream interceptor to be used as a middleware
func (mw *TracingMiddleware) StreamInterceptor() grpc.StreamServerInterceptor {
	return grpctrace.StreamServerInterceptor(defaultGrpcOption(mw.ServiceName)...)
}

// GrpcUnaryInterceptorMiddleware returns a grpc server stream interceptor to be used as a middleware
func (mw *TracingMiddleware) UnaryInterceptor() grpc.UnaryServerInterceptor {
	return grpctrace.UnaryServerInterceptor(defaultGrpcOption(mw.ServiceName)...)
}

// defaultGrpcOption returns the default grpc configuration options
func defaultGrpcOption(serviceName string) []grpctrace.Option {
	return []grpctrace.Option{
		grpctrace.WithServiceName(serviceName),
		grpctrace.WithStreamCalls(true),
		grpctrace.WithAnalytics(true),
		grpctrace.WithMetadataTags(),
		grpctrace.WithAnalytics(true),
		grpctrace.WithRequestTags(),
	}
}
