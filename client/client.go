package client

import (
	"context"
	"time"

	timeoutmiddleware "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/timeout"
	pbapi "github.com/rabataio/go-tron/proto/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func unaryInterceptor(token string) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply any,
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		md := metadata.New(map[string]string{
			"TRON-PRO-API-KEY": token,
		})
		ctx = metadata.NewOutgoingContext(ctx, md)

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func NewConnection(endpoint string, token string, timeout time.Duration) (*grpc.ClientConn, error) {
	connection, err := grpc.NewClient(
		endpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(unaryInterceptor(token)),
		grpc.WithUnaryInterceptor(timeoutmiddleware.UnaryClientInterceptor(timeout)),
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
