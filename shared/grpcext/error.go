package grpcext

import (
	"IofIPOS/shared/errors"
	"context"
	"google.golang.org/grpc"
)

func WrapErrors() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		res, err := handler(ctx, req)
		if err != nil {
			return nil, errors.Cast(ctx, err)
		}
		return res, nil
	}
}
