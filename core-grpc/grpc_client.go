package core_grpc

import (
	"context"
	"crypto/x509"
	"log"
	"time"

	"github.com/SimifiniiCTO/core/core-middleware/server"
	tlscert "github.com/SimifiniiCTO/core/core-tlsCert"
	grpcclient "github.com/apssouza22/grpc-production-go/client"
	"github.com/apssouza22/grpc-production-go/grpcutils"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/metadata"
)

type GrpcClient struct {
	Logger *zap.Logger
	Conn   *grpc.ClientConn
}

type GrpcClientConfigurations struct {
	GrpcServerConnectionAddr string
	Certificate              *x509.CertPool
	Logger                   *zap.Logger
	ServiceConfigs           *server.Configurations
	EnableTls                bool
}

// NewGrpcClient initializes a new GRPC client connection
// usage:
//  gc := NewGrpcClient(c *GrpcClientConfigurations)
// defer gc.Conn.Close()
//
func (grpcClientInstance *GrpcClient) NewGrpcClient(c *GrpcClientConfigurations) *GrpcClient {
	clientBuilder := grpcclient.GrpcConnBuilder{}
	clientBuilder.WithInsecure()
	clientBuilder.WithContext(context.Background())
	if c.EnableTls {
		if c.Certificate != nil {
			clientBuilder.WithClientTransportCredentials(false, c.Certificate)
		} else {
			clientBuilder.WithClientTransportCredentials(false, tlscert.CertPool)
		}
	}

	cc, err := clientBuilder.GetConn(c.GrpcServerConnectionAddr)
	if err != nil {
		log.Fatal(err.Error())
	}

	ConfigureClientGrpcMiddlewares(&clientBuilder)

	return &GrpcClient{
		Logger: c.Logger,
		Conn:   cc,
	}
}

// SendRpcRequest performs an rpc request against a downstream server
func (grpcClientInstance *GrpcClient) SendRpcRequest(ctx context.Context, ctxPairs []string, rpcOp func(ctx context.Context,
	param ...interface{}) (response interface{}, err error),
	requestParams ...interface{}) (rpcResponse interface{}, err error) {
	md := metadata.Pairs(ctxPairs...)
	ctx = metadata.NewOutgoingContext(ctx, md)
	timeout := time.Minute * 1
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	{
		healthClient := grpc_health_v1.NewHealthClient(grpcClientInstance.Conn)
		response, err := healthClient.Check(ctx, &grpc_health_v1.HealthCheckRequest{})
		if err != nil {
			grpcClientInstance.Logger.Error(err.Error())
			return nil, err
		}
		grpcClientInstance.Logger.Info("successfully obtained response from health client", zap.Any("Response", response))
	}

	rpcResponse, err = rpcOp(ctx, requestParams...)
	if err != nil {
		grpcClientInstance.Logger.Error(err.Error())
		return nil, err
	}

	grpcClientInstance.Logger.Info("successfully obtained response from rpc server", zap.Any("Response", rpcResponse))

	return rpcResponse, err
}

func ConfigureClientGrpcMiddlewares(builder *grpcclient.GrpcConnBuilder) {
	grpcUnaryInterceptors := grpcutils.GetDefaultUnaryClientInterceptors()
	grpcStreamInterceptors := grpcutils.GetDefaultStreamClientInterceptors()

	opts := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffExponential(20 * time.Millisecond)),
		grpc_retry.WithPerRetryTimeout(300 * time.Millisecond),
		grpc_retry.WithMax(3),
		grpc_retry.WithCodes(codes.NotFound, codes.Aborted, codes.Unknown, codes.Unavailable),
	}

	grpcUnaryInterceptors = append(grpcUnaryInterceptors, grpc_retry.UnaryClientInterceptor(opts...))
	grpcStreamInterceptors = append(grpcStreamInterceptors, grpc_retry.StreamClientInterceptor(opts...))
	builder.WithStreamInterceptors(grpcStreamInterceptors)
	builder.WithUnaryInterceptors(grpcUnaryInterceptors)
}
