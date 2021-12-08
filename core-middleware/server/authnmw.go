package server

import (
	"context"
	"errors"
	"net/http"
	"strings"

	core_auth_sdk "github.com/SimifiniiCTO/core/core-auth-sdk"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc/grpc-go/status"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type AuthenticationMiddleware struct {
	client      *core_auth_sdk.Client
	logger      *zap.Logger
	serviceName string
}

type contextKey struct {
	name string
}

var ctxKey *contextKey

// NewAuthenticationMiddleware returns an instance of the authentication middleware object
func NewAuthenticationMiddleware(logger *zap.Logger, client *core_auth_sdk.Client, serviceName string) *AuthenticationMiddleware {
	ctxKey = &contextKey{serviceName}

	return &AuthenticationMiddleware{
		client:      client,
		logger:      logger,
		serviceName: serviceName,
	}
}

// grpcAuthFunc serves as the authentication function of choice ran by the middleware
func (mw *AuthenticationMiddleware) grpcAuthFunc(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		mw.logger.Error(err.Error())
		return nil, err
	}

	// decode the token
	decodedToken, err := mw.client.SubjectFrom(token)
	if err != nil {
		mw.logger.Error(err.Error())
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}

	grpc_ctxtags.Extract(ctx).Set("auth.sub", decodedToken)
	ctx = context.WithValue(ctx, ctxKey, decodedToken)
	return ctx, nil
}

// Handler wraps the authentication middleware around http calls
func (mw *AuthenticationMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		w.Header().Set("Content-Type", "application/json")

		ctx := r.Context()
		authorization := r.Header.Get("Authorization")
		token := strings.TrimPrefix(authorization, "Bearer ")
		decodedToken, err := mw.client.SubjectFrom(token)
		if err != nil {
			mw.logger.Error(err.Error())
			next.ServeHTTP(w, r)
			return
		}

		ctx = context.WithValue(ctx, ctxKey, decodedToken)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

// StreamInterceptor returns a stream middleware for authentication purposes
func (mw *AuthenticationMiddleware) StreamInterceptor() grpc.StreamServerInterceptor {
	return grpc_auth.StreamServerInterceptor(mw.grpcAuthFunc)
}

// UnaryInterceptor returns a unary middleware for authentication purposes
func (mw *AuthenticationMiddleware) UnaryInterceptor() grpc.UnaryServerInterceptor {
	return grpc_auth.UnaryServerInterceptor(mw.grpcAuthFunc)
}

// IsAuthenticated returns whether or not the user is authenticated.
// REQUIRES Middleware to have run.
func (mw *AuthenticationMiddleware) IsAuthenticated(ctx context.Context) bool {
	return ctx.Value(ctxKey) != nil
}

// GetTokenFromCtx extracts the token from the context
func (mw *AuthenticationMiddleware) GetTokenFromCtx(ctx context.Context) (string, error) {
	if mw.IsAuthenticated(ctx) {
		token, ok := ctx.Value(ctxKey).(string)
		if !ok {
			return "", status.Errorf(codes.Unknown, "unable to obtain token from ctx key")
		}

		return token, nil
	}

	return "", errors.New("token not found in context")
}

// InjectContextWithMockToken injects a token into the context. Useful for testing
func (mw *AuthenticationMiddleware) InjectContextWithMockToken(ctx context.Context, token string, serviceName string) context.Context {
	ctxKey = &contextKey{serviceName}
	return context.WithValue(ctx, ctxKey, token)
}
