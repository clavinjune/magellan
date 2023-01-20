package authentication

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	magellanv1 "github.com/clavinjune/magellan/api/proto/magellan/v1"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc/metadata"
)

type handler struct {
	_ struct{}
	*magellanv1.UnimplementedAuthenticationServiceServer
}

func (h *handler) Login(ctx context.Context, request *magellanv1.AuthenticationServiceLoginRequest) (*magellanv1.AuthenticationServiceLoginResponse, error) {
	if request.AuthenticationType == magellanv1.AuthenticationType_AUTHENTICATION_TYPE_UNSPECIFIED {
		return nil, status.Error(codes.Unauthenticated, "Authentication Type is unspecified")
	}

	md, _ := metadata.FromIncomingContext(ctx)
	slog.LogAttrs(slog.LevelInfo, "search",
		slog.Any("metadata", md),
		slog.Any("request", request),
	)

	grpc.SendHeader(ctx,
		metadata.Pairs("testing", "bruh"),
	)

	return &magellanv1.AuthenticationServiceLoginResponse{
		Message: request.AuthenticationType.String(),
	}, nil
}
