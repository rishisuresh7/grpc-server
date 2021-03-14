package interceptor

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func NewInterceptor(secret string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, _ := metadata.FromIncomingContext(ctx)
		token, exists := md[":authority"]
		if !exists {
			return nil, status.Errorf(codes.Unauthenticated, "NewInterceptor: no authority header provided")
		}

		if token[0] != secret {
			return nil, status.Errorf(codes.Unauthenticated, "NewInterceptor: invalid token")
		}

		return handler(ctx, req)
	}
}
