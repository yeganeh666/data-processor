package grpcext

import (
	"context"

	"IofIPOS/shared/contextext"
	"IofIPOS/shared/errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// CheckStaticToken checks all requests have provided valid token
func CheckStaticToken(validToken string) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		token, ok := contextext.GetToken(ctx)
		if !ok || token != validToken {
			return nil, errors.New(ctx, codes.Unauthenticated)
		}

		return handler(ctx, req)
	}
}
