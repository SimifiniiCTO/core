package server

import (
	"net/http"
	"strings"
	"time"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
)

type LoggingMiddleware struct {
	logger *zap.Logger
}

// NewLoggingMiddleware returns a new instance of the logging middleware
func NewLoggingMiddleware(logger *zap.Logger) *LoggingMiddleware {
	return &LoggingMiddleware{logger: logger}
}

// LoggingMiddleware runs the logging middleware
func (m *LoggingMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.logger.Info(
			"request started",
			zap.Any("proto", r.Proto),
			zap.Any("uri", r.RequestURI),
			zap.Any("method", r.Method),
			zap.Any("remote", r.RemoteAddr),
			zap.Any("user-agent", r.UserAgent()),
		)
		next.ServeHTTP(w, r)
	})
}

// StreamInterceptor returns a streaming middleware for logging purposes
func (m *LoggingMiddleware) StreamInterceptor() grpc.StreamServerInterceptor {
	opts := m.defaultConfiguration()
	stream := grpc_zap.StreamServerInterceptor(m.logger, opts...)
	return stream
}

// UnaryInterceptor returns a unary middleware for logging purposes
func (m *LoggingMiddleware) UnaryInterceptor() grpc.UnaryServerInterceptor {
	opts := m.defaultConfiguration()
	unary := grpc_zap.UnaryServerInterceptor(m.logger, opts...)
	return unary
}

// defaultConfiguration provides default configurations for the logging middleware
func (m *LoggingMiddleware) defaultConfiguration() []grpc_zap.Option {
	opts := []grpc_zap.Option{
		grpc_zap.WithDurationField(func(duration time.Duration) zapcore.Field {
			return zap.Int64("grpc.time_ns", duration.Nanoseconds())
		}),
		grpc_zap.WithDecider(func(fullMethodName string, err error) bool {
			// will not log gRPC calls if it was a call to healthcheck and no error was raised
			if err == nil && strings.Contains(fullMethodName, "healthcheck") {
				return false
			}
			// by default everything will be logged
			return true
		}),
	}
	return opts
}
