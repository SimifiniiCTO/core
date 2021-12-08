package server

import (
	"context"
	"math/rand"
	"net/http"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RandomErrMiddleware struct {
	Logger *zap.Logger
}

// NewRandomErrMiddleware returns a new instance of the random errors middleware
func NewRandomErrMiddleware(logger *zap.Logger) *RandomErrMiddleware {
	return &RandomErrMiddleware{Logger: logger}
}

// Handler runs the random error middleware
func (mw *RandomErrMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if mw.RandomError() {
			errors := []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusConflict}
			w.WriteHeader(errors[rand.Intn(len(errors))])
			mw.Logger.Error("Internal Server Error")
			return
		}
		next.ServeHTTP(w, r)
	})
}

// RandomError adds a random error to a set of requests
func (mw *RandomErrMiddleware) RandomError() bool {
	rand.Seed(time.Now().Unix())
	if rand.Int31n(3) == 0 {
		return true
	}
	return false
}

// UnaryServerInterceptor returns a new unary server interceptor that randomly fails requests
func (mw *RandomErrMiddleware) UnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if mw.RandomError() {
			return nil, status.Error(codes.Internal, "Internal Server Error")
		}
		return handler(ctx, req)
	}
}

// StreamServerInterceptor returns a new unary server interceptor that randomly fails requests
func (mw *RandomErrMiddleware) StreamInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if mw.RandomError() {
			return status.Error(codes.Internal, "Internal Server Error")
		}
		return handler(srv, stream)
	}
}
