package server

import (
	"log"

	core_auth_sdk "github.com/SimifiniiCTO/core/core-auth-sdk"
	"github.com/apssouza22/grpc-production-go/grpcutils"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Configurations struct {
	StatsDConnectionAddr        string
	Logger                      *zap.Logger
	Client                      *core_auth_sdk.Client
	ServiceName                 string
	Origins                     []string
	EnableDelayMiddleware       bool
	EnableRandomErrorMiddleware bool
	MinDelay                    int
	MaxDelay                    int
	DelayUnit                   string
	Version                     string
}

type ServiceMiddlewares struct {
	AuthenticationMiddleware *AuthenticationMiddleware
	CorsMiddleware           *CorsMiddleware
	LoggingMiddleware        *LoggingMiddleware
	MetricsMiddleware        *MetricsMiddleware
	RandomDelayMiddleware    *RandomDelayMiddleware
	RandomErrMiddleware      *RandomErrMiddleware
	VersionMiddleware        *VersionMiddleware
	TracingMiddleware        *TracingMiddleware
}

// InitializeMiddleware initializes a middleware object ecompassing every middleware in this library
func InitializeMiddleware(c *Configurations) *ServiceMiddlewares {
	if c == nil {
		log.Fatalf("invalid input argument. configurations cannot be nil")
	}
	var serviceMw ServiceMiddlewares

	serviceMw.AuthenticationMiddleware = NewAuthenticationMiddleware(c.Logger, c.Client, c.ServiceName)
	serviceMw.CorsMiddleware = NewCorsMiddleware(c.Origins)
	serviceMw.LoggingMiddleware = NewLoggingMiddleware(c.Logger)
	serviceMw.MetricsMiddleware = NewMetricsMiddleware(c.StatsDConnectionAddr, c.Logger)
	serviceMw.VersionMiddleware = NewVersionMw(c.Version)
	serviceMw.TracingMiddleware = NewTracingMiddleware(c.ServiceName)

	if c.EnableRandomErrorMiddleware {
		serviceMw.RandomErrMiddleware = NewRandomErrMiddleware(c.Logger)
	}

	if c.EnableDelayMiddleware {
		serviceMw.RandomDelayMiddleware = NewRandomDelayMiddleware(c.MinDelay, c.MaxDelay, c.DelayUnit)
	}

	return &serviceMw
}

// StreamInterceptor returns a set of stream interceptors
func (m *ServiceMiddlewares) StreamInterceptor() []grpc.StreamServerInterceptor {
	streamInterceptors := []grpc.StreamServerInterceptor{
		m.AuthenticationMiddleware.StreamInterceptor(),
		m.LoggingMiddleware.StreamInterceptor(),
		m.VersionMiddleware.StreamInterceptor(),
		m.TracingMiddleware.StreamInterceptor(),
		grpc_ctxtags.StreamServerInterceptor(),
	}

	streamInterceptors = append(streamInterceptors, grpcutils.GetDefaultStreamServerInterceptors()...)

	if m.RandomDelayMiddleware != nil {
		streamInterceptors = append(streamInterceptors, m.RandomDelayMiddleware.StreamInterceptor())
	}

	if m.RandomErrMiddleware != nil {
		streamInterceptors = append(streamInterceptors, m.RandomErrMiddleware.StreamInterceptor())
	}

	return streamInterceptors
}

// UnaryInterceptor returns a set of unary interceptors
func (m *ServiceMiddlewares) UnaryInterceptor() []grpc.UnaryServerInterceptor {
	unaryInterceptors := []grpc.UnaryServerInterceptor{
		m.AuthenticationMiddleware.UnaryInterceptor(),
		m.LoggingMiddleware.UnaryInterceptor(),
		m.VersionMiddleware.UnaryInterceptor(),
		m.TracingMiddleware.UnaryInterceptor(),
		grpc_ctxtags.UnaryServerInterceptor(),
	}

	unaryInterceptors = append(unaryInterceptors, grpcutils.GetDefaultUnaryServerInterceptors()...)

	if m.RandomDelayMiddleware != nil {
		unaryInterceptors = append(unaryInterceptors, m.RandomDelayMiddleware.UnaryInterceptor())
	}

	if m.RandomErrMiddleware != nil {
		unaryInterceptors = append(unaryInterceptors, m.RandomErrMiddleware.UnaryInterceptor())
	}

	return unaryInterceptors
}
