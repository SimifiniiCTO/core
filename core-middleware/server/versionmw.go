package server

import (
	"context"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type VersionMiddleware struct {
	version string
}

// NewVersionMw returns a new instance of the version errors middleware
func NewVersionMw(version string) *VersionMiddleware {
	return &VersionMiddleware{}
}

// Handler runs the version middleware
func (mw *VersionMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("X-API-Version", mw.version)
		r.Header.Set("X-API-Revision", mw.version)

		next.ServeHTTP(w, r)
	})
}

// UnaryServerInterceptor returns a new unary server interceptor that add security header
//
// Invalid messages will be rejected with `Internal` before reaching any userspace handlers.
func (mw *VersionMiddleware) UnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if err := mw.setVersion(ctx); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		return handler(ctx, req)
	}
}

// StreamServerInterceptor returns a new unary server interceptor that adds a version header
func (mw *VersionMiddleware) StreamInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		err := mw.setVersion(stream.Context())
		if err != nil {
			return err
		}
		return handler(srv, stream)
	}
}

// setVersion sets a version as part of the ctx parameters
func (mw *VersionMiddleware) setVersion(ctx context.Context) error {
	headers := map[string]string{
		"X-API-Version":  mw.version, // http://stackoverflow.com/a/3146618/244009
		"X-API-Revision": mw.version,
	}
	return grpc.SetHeader(ctx, metadata.New(headers))
}
