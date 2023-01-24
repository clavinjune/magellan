package authentication

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	authenticationv1 "github.com/clavinjune/magellan/api/proto/magellan/authentication/v1"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc/metadata"
)

type handler struct {
	_ struct{}
	*authenticationv1.UnimplementedServiceServer
}

func (h *handler) Login(ctx context.Context, request *authenticationv1.ServiceLoginRequest) (*authenticationv1.ServiceLoginResponse, error) {
	if request.LoginType == authenticationv1.LoginType_LOGIN_TYPE_UNSPECIFIED {
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

	return &authenticationv1.ServiceLoginResponse{
		Message: request.LoginType.String(),
	}, nil
}

func (h *handler) Logout(context.Context, *authenticationv1.ServiceLogoutRequest) (*authenticationv1.ServiceLogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
