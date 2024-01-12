package auth

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"

	"go.uber.org/zap"
)

type Service struct {
	Logger *zap.Logger
	authpb.UnimplementedAuthServiceServer
}

func (s *Service) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	s.Logger.Info("received code ", zap.String("code", req.Code))
	return &authpb.LoginResponse{
		AccessToken: "token for" + req.Code,
		ExpireIn:    7200,
	}, nil
}
