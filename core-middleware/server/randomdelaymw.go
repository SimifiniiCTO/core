package server

import (
	"context"
	"math/rand"
	"net/http"
	"time"

	"google.golang.org/grpc"
)

type RandomDelayMiddleware struct {
	min  int
	max  int
	unit string
}

// NewRandomDelayMiddleware returns an instance of the random delay middleware
func NewRandomDelayMiddleware(minDelay, maxDelay int, delayUnit string) *RandomDelayMiddleware {
	return &RandomDelayMiddleware{
		min:  minDelay,
		max:  maxDelay,
		unit: delayUnit,
	}
}

// DelayMiddleware runs the random delay middleware
func (m *RandomDelayMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.Delay()
		next.ServeHTTP(w, r)
	})
}

// Delay adds a random delay to the request
func (m *RandomDelayMiddleware) Delay() {
	var unit time.Duration
	rand.Seed(time.Now().Unix())
	switch m.unit {
	case "s":
		unit = time.Second
	case "ms":
		unit = time.Millisecond
	default:
		unit = time.Second
	}

	delay := rand.Intn(m.max-m.min) + m.min
	time.Sleep(time.Duration(delay) * unit)
}

// UnaryInterceptor returns a new unary server interceptor that adds a delay to the request
func (m *RandomDelayMiddleware) UnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		m.Delay()
		return handler(ctx, req)
	}
}

// StreamInterceptor returns a new unary server interceptor that add a delay to the request
func (m *RandomDelayMiddleware) StreamInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		m.Delay()
		return handler(srv, stream)
	}
}
