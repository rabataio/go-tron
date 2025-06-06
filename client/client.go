package client

import (
	"context"
	"time"

	timeoutmiddleware "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/timeout"
	pbapi "github.com/rabataio/go-tron/proto/api"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

const (
	defaultRateLimit = 15
	defaultBurst     = 1
)

func tokenInterceptor(token string) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply any,
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		ctx = metadata.AppendToOutgoingContext(ctx, "TRON-PRO-API-KEY", token)

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func rateLimitInterceptor(limiter *rate.Limiter) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply any,
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		if err := limiter.Wait(ctx); err != nil {
			return err
		}

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func NewConnection(endpoint string, token string, timeout time.Duration) (*grpc.ClientConn, error) {
	limiter := rate.NewLimiter(rate.Limit(defaultRateLimit), defaultBurst)

	connection, err := grpc.NewClient(
		endpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			rateLimitInterceptor(limiter),
			tokenInterceptor(token),
			timeoutmiddleware.UnaryClientInterceptor(timeout),
		),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(32*1024*1024)),
	)
	if err != nil {
		return nil, err
	}

	return connection, nil
}

func NewClient(connection *grpc.ClientConn) pbapi.WalletClient {
	return pbapi.NewWalletClient(connection)
}
