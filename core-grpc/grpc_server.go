package core_grpc

import (
	"crypto/tls"

	"github.com/SimifiniiCTO/core/core-middleware/server"
	grpcserver "github.com/apssouza22/grpc-production-go/server"
	tlscert "github.com/yoanyombapro1234/FeelGuuds_Core/core/core-tlsCert"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	Logger  *zap.Logger
	Server  grpcserver.GrpcServer
	Address string
}

type GrpcServerConfigurations struct {
	GrpcServerConnectionAddr string
	Certificate              *tls.Certificate
	Logger                   *zap.Logger
	ServiceConfigs           *server.Configurations
	EnableTls                bool
}

// NewGrpcService Initializes a new instance of a grpc service
func NewGrpcService(c *GrpcServerConfigurations) *GrpcServer {
	serverBuilder := grpcserver.GrpcServerBuilder{}

	AddInterceptors(&serverBuilder, c.ServiceConfigs)
	serverBuilder.EnableReflection(true)

	if c.EnableTls {
		if c.Certificate != nil {
			serverBuilder.SetTlsCert(&tlscert.Cert)
		} else {
			serverBuilder.SetTlsCert(c.Certificate)
		}
	}

	s := serverBuilder.Build()

	return &GrpcServer{Address: c.GrpcServerConnectionAddr, Logger: c.Logger, Server: s}
}

// StartGrpcServer starts a grpc service
// usage:
//  s := NewGrpcService(logger *zap.Logger, enableTls bool, cert *tls.Certificate)
//  s.StartGrpcServer(addr string, fn func(server *grpc.Server))
func (grpcSrvInstance *GrpcServer) StartGrpcServer(fn func(server *grpc.Server)) {
	s := grpcSrvInstance.Server
	l := grpcSrvInstance.Logger
	addr := grpcSrvInstance.Address
	s.RegisterService(fn)

	err := s.Start(addr)
	if err != nil {
		l.Fatal(err.Error())
	}

	grpcSrvInstance.AwaitTermination()
}

// AwaitTermination Shuts down grpc server
func (grpcSrvInstance *GrpcServer) AwaitTermination() {
	s := grpcSrvInstance.Server
	l := grpcSrvInstance.Logger

	s.AwaitTermination(func() {
		l.Info("shutting down grpc server")
	})
}

// AddInterceptors adds default rpc interceptors to grpc service instance
func AddInterceptors(s *grpcserver.GrpcServerBuilder, configurations *server.Configurations) {
	mw := server.InitializeMiddleware(configurations)
	s.SetUnaryInterceptors(mw.UnaryInterceptor())
	s.SetStreamInterceptors(mw.StreamInterceptor())
}
