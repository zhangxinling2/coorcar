package auth

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	Logger *zap.Logger
	authpb.UnimplementedAuthServiceServer
	OpenIdResolve OpenIdResolve
}
type OpenIdResolve interface {
	Resolve(code string) (string, error)
}

func (s *Service) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	s.Logger.Info("received code ", zap.String("code", req.Code))
	openId, err := s.OpenIdResolve.Resolve(req.Code)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "cannot resolve openid: %v", err)
	}
	return &authpb.LoginResponse{
		AccessToken: "token for openId " + openId,
		ExpireIn:    7200,
	}, nil
}
